package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pelletier/go-toml/v2"
)

func TestConfigRoundTrip(t *testing.T) {
	a, out, errOut := harness(t)
	a.envToken = ""
	a.in = strings.NewReader("  secret-for-test\n")
	a.transport = roundTripFunc(func(*http.Request) (*http.Response, error) { t.Fatal("config called API"); return nil, nil })
	for _, args := range [][]string{
		{"config", "set", "api_key", "--stdin"},
		{"config", "set", "threshold", "0.8"},
		{"config", "set", "json", "false"},
		{"config", "set", "model", "chosen-model"},
		{"config", "set", "timeout", "45s"},
		{"config", "set", "pick", "answer,model"},
	} {
		if code := a.run(context.Background(), args); code != 0 || out.Len() > 0 || errOut.Len() > 0 {
			t.Fatalf("%v: code=%d stdout=%q stderr=%q", args, code, out, errOut)
		}
	}
	c, err := readConfig(a.path)
	if err != nil || c.APIKey != "secret-for-test" || c.Threshold == nil || *c.Threshold != .8 || c.JSON == nil || *c.JSON {
		t.Fatalf("config did not preserve stored settings: %v", err)
	}
	info, err := os.Stat(a.path)
	if err != nil || info.Mode().Perm() != 0600 {
		t.Fatalf("config permissions: %v %v", info, err)
	}
	if code := a.run(context.Background(), []string{"config"}); code != 0 || strings.Contains(out.String(), "secret-for-test") {
		t.Fatalf("unsafe display: code=%d stdout=%q stderr=%q", code, out, errOut)
	}
	var shown settings
	if err := toml.Unmarshal(out.Bytes(), &shown); err != nil || shown.APIKey != "********" || shown.Model != "chosen-model" || shown.Timeout != "45s" {
		t.Fatalf("display did not reflect settings: %v", err)
	}
	for _, key := range []string{"threshold", "model", "timeout", "json", "pick", "api_key"} {
		out.Reset()
		if code := a.run(context.Background(), []string{"config", "unset", key}); code != 0 || out.Len() != 0 {
			t.Fatalf("unset %s: code=%d stderr=%s", key, code, errOut)
		}
	}
	c, err = readConfig(a.path)
	s := c.resolve("")
	if err != nil || s.Threshold != nil || s.JSON || s.Model != "jev-latest" || s.Timeout != "30s" || len(s.Pick) != 0 || s.APIKey != "" {
		t.Fatal("unset did not restore defaults")
	}
}

func TestConfigReadAndNoopDoNotCreateFile(t *testing.T) {
	for _, args := range [][]string{{"config"}, {"config", "path"}, {"config", "unset", "threshold"}, evaluationArgs()} {
		a, _, errOut := harness(t)
		if code := a.run(context.Background(), args); code != 0 {
			t.Fatalf("%v: %d %s", args, code, errOut)
		}
		if _, err := os.Stat(filepath.Dir(a.path)); !errors.Is(err, os.ErrNotExist) {
			t.Fatal("read created config directory")
		}
	}
}

func TestConfigPrecedence(t *testing.T) {
	a, out, errOut := harness(t)
	a.envToken = "environment-token"
	threshold, jsonOutput, model, timeout := .99, true, "file-model", "12s"
	c := config{APIKey: "file-token", Threshold: &threshold, JSON: &jsonOutput, Model: &model, Timeout: &timeout, Pick: &[]string{"answer"}}
	if err := writeConfig(a.path, c); err != nil {
		t.Fatal(err)
	}
	a.transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		if r.Header.Get("Authorization") != "Bearer environment-token" {
			t.Fatal("environment did not override file token")
		}
		return response(200, fixture), nil
	})
	args := evaluationArgs("--threshold", "0", "--json=false", "--pick", "")
	if code := a.run(context.Background(), args); code != 0 || out.String() != "충돌함\n" {
		t.Fatalf("code=%d stdout=%q stderr=%q", code, out, errOut)
	}
	s := c.resolve(" ")
	if s.APIKey != "file-token" || s.Model != "file-model" || s.Timeout != "12s" {
		t.Fatal("file defaults were lost")
	}
	flags, err := parseEvaluation(evaluationArgs("--model", "flag-model", "--timeout", "9s"))
	if err != nil {
		t.Fatal(err)
	}
	c.override(flags.flags)
	s = c.resolve("")
	if s.Model != "flag-model" || s.Timeout != "9s" {
		t.Fatal("flags did not override config")
	}
}

func TestConfigFailuresDoNotLeakOrReplace(t *testing.T) {
	a, out, errOut := harness(t)
	if err := writeConfig(a.path, config{APIKey: "retained-secret"}); err != nil {
		t.Fatal(err)
	}
	for _, args := range [][]string{
		{"config", "set", "api_key", "argument-secret"}, {"config", "set", "threshold", "NaN"},
		{"config", "set", "timeout", "0s"}, {"config", "set", "json", "maybe"},
		{"config", "set", "pick", "unknown"}, {"config", "unset", "unknown"},
	} {
		out.Reset()
		errOut.Reset()
		if code := a.run(context.Background(), args); code != 1 || out.Len() != 0 || strings.Contains(errOut.String(), "secret") {
			t.Fatalf("unsafe failure: code=%d stdout=%q stderr=%q", code, out, errOut)
		}
		c, err := readConfig(a.path)
		if err != nil || c.APIKey != "retained-secret" || c.Threshold != nil {
			t.Fatal("failed write changed config")
		}
	}
	for _, raw := range []string{`api_key = "source-secret`, `threshold = "source-secret"`, `unknown = "source-secret"`} {
		if err := os.WriteFile(a.path, []byte(raw), 0600); err != nil {
			t.Fatal(err)
		}
		out.Reset()
		errOut.Reset()
		if code := a.run(context.Background(), []string{"config"}); code != 1 || out.Len() != 0 || strings.Contains(errOut.String(), "source-secret") {
			t.Fatalf("TOML failure leaked source: %d %q %q", code, out, errOut)
		}
	}
	if code := a.run(context.Background(), []string{"config", "path"}); code != 0 {
		t.Fatal("path depends on valid config")
	}
}

func TestConfigRefusesSymlinksAndDirectories(t *testing.T) {
	dir := t.TempDir()
	target := filepath.Join(dir, "target")
	if err := os.WriteFile(target, []byte(`api_key = "unchanged"`), 0600); err != nil {
		t.Fatal(err)
	}
	link := filepath.Join(dir, "link")
	if err := os.Symlink(target, link); err != nil {
		t.Fatal(err)
	}
	for _, path := range []string{link, dir} {
		if _, err := readConfig(path); err == nil {
			t.Fatal("read non-regular config")
		}
		if err := writeConfig(path, config{}); err == nil {
			t.Fatal("replaced non-regular config")
		}
	}
	data, err := os.ReadFile(target)
	if err != nil || string(data) != `api_key = "unchanged"` {
		t.Fatal("symlink target changed")
	}
}

func TestConfigPreservesExplicitEmptyPick(t *testing.T) {
	a, _, errOut := harness(t)
	for _, args := range [][]string{
		{"config", "set", "pick", ""},
		{"config", "set", "threshold", "0"},
	} {
		if code := a.run(context.Background(), args); code != 0 {
			t.Fatalf("%v: %d %s", args, code, errOut)
		}
		c, err := readConfig(a.path)
		if err != nil || c.Pick == nil || len(*c.Pick) != 0 {
			t.Fatalf("explicit empty pick not preserved: %v", err)
		}
	}
	if code := a.run(context.Background(), []string{"config", "unset", "pick"}); code != 0 {
		t.Fatalf("unset: %d %s", code, errOut)
	}
	c, err := readConfig(a.path)
	if err != nil || c.Pick != nil {
		t.Fatal("unset did not remove pick")
	}
}
