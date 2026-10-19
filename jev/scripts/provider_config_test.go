package main

import (
	"bytes"
	"context"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"testing"

	"github.com/pelletier/go-toml/v2"
)

func TestConfigInitAndTemplateDefaults(t *testing.T) {
	a, out, errOut := harness(t)
	a.transport = roundTripFunc(func(*http.Request) (*http.Response, error) { t.Fatal("init called API"); return nil, nil })
	if code := a.run(context.Background(), []string{"config", "init"}); code != 0 || out.Len() != 0 {
		t.Fatalf("init=%d %s", code, errOut)
	}
	data, err := os.ReadFile(a.path)
	if err != nil {
		t.Fatal(err)
	}
	var actual config
	if err := toml.Unmarshal(data, &actual); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(actual.resolve(""), (config{}).resolve("")) {
		t.Fatal("template differs from built-in defaults")
	}
	if len(actual.Providers) != 1 || actual.Providers["jev"].APIKey != "" {
		t.Fatal("template has extra providers or credentials")
	}
	info, err := os.Stat(a.path)
	if err != nil || info.Mode().Perm() != 0600 {
		t.Fatalf("mode=%v err=%v", info, err)
	}
	for _, existing := range [][]byte{data, []byte("invalid config")} {
		if err := os.WriteFile(a.path, existing, 0600); err != nil {
			t.Fatal(err)
		}
		out.Reset()
		errOut.Reset()
		if code := a.run(context.Background(), []string{"config", "init"}); code != 1 || out.Len() != 0 {
			t.Fatalf("existing init=%d %s", code, errOut)
		}
		after, _ := os.ReadFile(a.path)
		if !bytes.Equal(after, existing) {
			t.Fatal("init replaced existing file")
		}
	}
}

func TestConcurrentConfigInit(t *testing.T) {
	path := filepath.Join(t.TempDir(), "jev.toml")
	var wg sync.WaitGroup
	results := make(chan error, 8)
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); results <- initConfig(path) }()
	}
	wg.Wait()
	close(results)
	successes := 0
	for err := range results {
		if err == nil {
			successes++
		}
	}
	if successes != 1 {
		t.Fatalf("successful initializers=%d", successes)
	}
	data, err := os.ReadFile(path)
	if err != nil || string(data) != configTemplate {
		t.Fatal("incomplete template")
	}
}

func TestFirstSetAndProviderScopedUnset(t *testing.T) {
	a, _, errOut := harness(t)
	if code := a.run(context.Background(), []string{"config", "set", "providers.openai.model", "configured-model"}); code != 0 {
		t.Fatalf("set=%d %s", code, errOut)
	}
	c, err := readConfig(a.path)
	if err != nil {
		t.Fatal(err)
	}
	if c.Defaults.Provider == nil || *c.Defaults.Provider != "jev" || c.Providers["jev"].Model == nil || *c.Providers["jev"].Model != "jev-latest" || *c.Providers["openai"].Model != "configured-model" {
		t.Fatal("first set lost template or requested value")
	}
	data, _ := os.ReadFile(a.path)
	if !strings.Contains(string(data), "# threshold = 0.8") {
		t.Fatal("template guidance missing")
	}
	// Uncommenting the generated example must still configure defaults, not a provider.
	example := strings.Replace(string(data), "# threshold = 0.8", "threshold = 0.8", 1)
	var edited config
	if err := toml.Unmarshal([]byte(example), &edited); err != nil || edited.Defaults.Threshold == nil || *edited.Defaults.Threshold != 0.8 {
		t.Fatal("first-set template misplaced the default example")
	}
	a.in = strings.NewReader("openai-secret")
	if code := a.run(context.Background(), []string{"config", "set", "providers.openai.api_key", "--stdin"}); code != 0 {
		t.Fatalf("key=%d %s", code, errOut)
	}
	if code := a.run(context.Background(), []string{"config", "unset", "providers.openai.model"}); code != 0 {
		t.Fatalf("unset=%d %s", code, errOut)
	}
	c, err = readConfig(a.path)
	if err != nil || c.Providers["openai"].Model != nil || c.Providers["openai"].APIKey != "openai-secret" || c.Providers["jev"].Model == nil {
		t.Fatal("unset changed adjacent fields")
	}
}

func TestConfigDisplayMasksEveryProvider(t *testing.T) {
	a, out, errOut := harness(t)
	a.env = map[string]string{"TYPESAFE_API_KEY": "jev-env-secret", "OPENROUTER_API_KEY": "router-env-secret", "OPENAI_API_KEY": "openai-env-secret"}
	selected := "openai"
	if err := writeConfig(a.path, config{Defaults: defaults{Provider: &selected}, Providers: map[string]providerConfig{"jev": {APIKey: "jev-file-secret"}, "openai": {APIKey: "openai-file-secret"}, "openrouter": {}}}); err != nil {
		t.Fatal(err)
	}
	a.transport = roundTripFunc(func(*http.Request) (*http.Response, error) { t.Fatal("config called API"); return nil, nil })
	if code := a.run(context.Background(), []string{"config"}); code != 0 {
		t.Fatalf("display=%d %s", code, errOut)
	}
	if strings.Contains(out.String(), "secret") {
		t.Fatalf("key leak: %s", out)
	}
	var shown config
	if err := toml.Unmarshal(out.Bytes(), &shown); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"jev", "openrouter", "openai"} {
		if shown.Providers[name].APIKey != "********" {
			t.Fatalf("unmasked provider %s", name)
		}
	}
	if *shown.Defaults.Provider != "openai" {
		t.Fatal("display changed default")
	}
}

func TestProviderSelectionAndKeyIsolation(t *testing.T) {
	a, out, errOut := harness(t)
	selected, jevModel, otherModel := "openai", "jev-model", "other-model"
	c := config{Defaults: defaults{Provider: &selected}, Providers: map[string]providerConfig{"jev": {APIKey: "jev-file-secret", Model: &jevModel}, "openai": {APIKey: "other-file-secret", Model: &otherModel}}}
	if err := writeConfig(a.path, c); err != nil {
		t.Fatal(err)
	}
	a.env = map[string]string{"TYPESAFE_API_KEY": "jev-env-secret", "OPENAI_API_KEY": "other-env-secret"}
	calls := 0
	a.transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		calls++
		if r.Header.Get("Authorization") != "Bearer jev-env-secret" {
			t.Fatal("wrong provider key")
		}
		return response(200, noulFixture), nil
	})
	if code := a.run(context.Background(), noulArgs()); code != 1 || calls != 0 || !strings.Contains(errOut.String(), "not implemented") {
		t.Fatalf("unsupported=%d calls=%d %s", code, calls, errOut)
	}
	out.Reset()
	errOut.Reset()
	if code := a.run(context.Background(), noulArgs("--provider", "jev")); code != 0 || calls != 1 {
		t.Fatalf("override=%d calls=%d %s", code, calls, errOut)
	}
	saved, _ := readConfig(a.path)
	if saved.provider() != "openai" {
		t.Fatal("override persisted")
	}
	delete(a.env, "TYPESAFE_API_KEY")
	delete(c.Providers, "jev")
	if err := writeConfig(a.path, c); err != nil {
		t.Fatal(err)
	}
	out.Reset()
	errOut.Reset()
	if code := a.run(context.Background(), noulArgs("--provider", "jev")); code != 1 || calls != 1 || !strings.Contains(errOut.String(), "API key missing") {
		t.Fatalf("cross-key=%d calls=%d %s", code, calls, errOut)
	}
}

func TestProviderOverridesInBatchAndDoctor(t *testing.T) {
	a, out, errOut := harness(t)
	selected := "openrouter"
	if err := writeConfig(a.path, config{Defaults: defaults{Provider: &selected}}); err != nil {
		t.Fatal(err)
	}
	state, input := batchFiles(t, "state", `{"id":"result","type":"noul","question":"True?"}`)
	calls := 0
	a.transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		calls++
		if r.Header.Get("Authorization") != "Bearer test-token" {
			t.Fatal("wrong key")
		}
		return response(200, noulFixture), nil
	})
	if code := a.run(context.Background(), batchArgs(state, input)); code != 1 || calls != 0 {
		t.Fatalf("unsupported batch=%d", code)
	}
	out.Reset()
	errOut.Reset()
	if code := a.run(context.Background(), batchArgs(state, input, "--provider", "jev")); code != 0 || calls != 1 {
		t.Fatalf("batch=%d calls=%d %s", code, calls, errOut)
	}
	out.Reset()
	errOut.Reset()
	if code := a.run(context.Background(), []string{"doctor", "--dir", t.TempDir(), "--provider", "jev"}); code != 1 || calls != 2 || !strings.Contains(errOut.String(), "authenticated") {
		t.Fatalf("doctor=%d calls=%d %s", code, calls, errOut)
	}
	out.Reset()
	errOut.Reset()
	if code := a.run(context.Background(), []string{"doctor", "--dir", t.TempDir()}); code != 1 || calls != 2 || !strings.Contains(errOut.String(), "not implemented") {
		t.Fatalf("unsupported doctor=%d calls=%d %s", code, calls, errOut)
	}
}

func TestConfigRejectsLegacyAndMixedLayouts(t *testing.T) {
	for _, raw := range []string{`api_key="private-token"`, "[default]\nprovider=\"jev\"", "model=\"legacy\"\n[defaults]\nprovider=\"jev\"", "[defaults]\nunknown=\"private-token\"", "[providers.other]\napi_key=\"private-token\""} {
		a, out, errOut := harness(t)
		if err := os.MkdirAll(filepath.Dir(a.path), 0700); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(a.path, []byte(raw), 0600); err != nil {
			t.Fatal(err)
		}
		a.transport = roundTripFunc(func(*http.Request) (*http.Response, error) { t.Fatal("invalid config called API"); return nil, nil })
		if code := a.run(context.Background(), []string{"config", "set", "defaults.timeout", "45s"}); code != 1 || out.Len() != 0 || strings.Contains(errOut.String(), "private-token") {
			t.Fatalf("code=%d stdout=%s stderr=%s", code, out, errOut)
		}
		after, _ := os.ReadFile(a.path)
		if string(after) != raw {
			t.Fatal("invalid config was rewritten")
		}
	}
}
