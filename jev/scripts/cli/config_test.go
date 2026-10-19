package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/pelletier/go-toml/v2"
	cfg "opnay/jev/config"
)

func TestConfigRoundTrip(t *testing.T) {
	a, out, errOut := harness(t)
	a.Env["TYPESAFE_API_KEY"] = ""
	a.In = strings.NewReader("  secret-for-test\n")
	a.transport = roundTripFunc(func(*http.Request) (*http.Response, error) { t.Fatal("config called API"); return nil, nil })
	for _, args := range [][]string{
		{"config", "set", "providers.jev.api_key", "--stdin"},
		{"config", "set", "defaults.threshold", "0.8"},
		{"config", "set", "defaults.min_score", "1.5"},
		{"config", "set", "defaults.json", "false"},
		{"config", "set", "providers.jev.model", "chosen-model"},
		{"config", "set", "defaults.timeout", "45s"},
		{"config", "set", "defaults.pick", "answer,model"},
	} {
		if code := a.Run(context.Background(), args); code != 0 || out.Len() > 0 || errOut.Len() > 0 {
			t.Fatalf("%v: code=%d stdout=%q stderr=%q", args, code, out, errOut)
		}
	}
	c, err := cfg.Read(a.path)
	if err != nil || c.Providers["jev"].APIKey != "secret-for-test" || c.Defaults.Threshold == nil || *c.Defaults.Threshold != .8 ||
		c.Defaults.MinScore == nil || *c.Defaults.MinScore != 1.5 || c.Defaults.JSON == nil || *c.Defaults.JSON ||
		c.Defaults.Pick == nil || !slices.Equal(*c.Defaults.Pick, []string{"answer", "model"}) {
		t.Fatalf("config did not preserve stored settings: %v", err)
	}
	info, err := os.Stat(a.path)
	if err != nil || info.Mode().Perm() != 0600 {
		t.Fatalf("config permissions: %v %v", info, err)
	}
	if code := a.Run(context.Background(), []string{"config"}); code != 0 || strings.Contains(out.String(), "secret-for-test") {
		t.Fatalf("unsafe display: code=%d stdout=%q stderr=%q", code, out, errOut)
	}
	var shown cfg.Config
	if err := toml.Unmarshal(out.Bytes(), &shown); err != nil || shown.Providers["jev"].APIKey != "********" || *shown.Providers["jev"].Model != "chosen-model" || *shown.Defaults.Timeout != "45s" ||
		!slices.Equal(*shown.Defaults.Pick, []string{"answer", "model"}) {
		t.Fatalf("display did not reflect settings: %v", err)
	}
	for _, key := range []string{"defaults.threshold", "defaults.min_score", "providers.jev.model", "defaults.timeout", "defaults.json", "defaults.pick", "providers.jev.api_key"} {
		out.Reset()
		if code := a.Run(context.Background(), []string{"config", "unset", key}); code != 0 || out.Len() != 0 {
			t.Fatalf("unset %s: code=%d stderr=%s", key, code, errOut)
		}
	}
	c, err = cfg.Read(a.path)
	s := c.Resolve("")
	if err != nil || s.Threshold != nil || s.MinScore != nil || s.JSON || s.Model != "jev-latest" || s.Timeout != "30s" || len(s.Pick) != 0 || s.APIKey != "" {
		t.Fatal("unset did not restore defaults")
	}
}

func TestConfigReadAndNoopDoNotCreateFile(t *testing.T) {
	for _, args := range [][]string{{"config"}, {"config", "path"}, {"config", "unset", "defaults.threshold"}, evaluationArgs()} {
		a, _, errOut := harness(t)
		if code := a.Run(context.Background(), args); code != 0 {
			t.Fatalf("%v: %d %s", args, code, errOut)
		}
		if _, err := os.Stat(filepath.Dir(a.path)); !errors.Is(err, os.ErrNotExist) {
			t.Fatal("read created config directory")
		}
	}
}

func TestConfigPrecedence(t *testing.T) {
	a, out, errOut := harness(t)
	a.Env["TYPESAFE_API_KEY"] = "environment-token"
	threshold, jsonOutput, model, timeout := .99, true, "file-model", "12s"
	c := cfg.Config{Defaults: cfg.Defaults{Threshold: &threshold, JSON: &jsonOutput, Timeout: &timeout, Pick: &[]string{"answer"}}, Providers: map[string]cfg.ProviderConfig{"jev": {APIKey: "file-token", Model: &model}}}
	if err := cfg.Write(a.path, c); err != nil {
		t.Fatal(err)
	}
	wantModel, wantTimeout := "file-model", 12*time.Second
	a.transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		if r.Header.Get("Authorization") != "Bearer environment-token" {
			t.Fatal("environment did not override file token")
		}
		var request struct {
			Model string `json:"model"`
		}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil || request.Model != wantModel {
			t.Fatalf("request model=%q err=%v", request.Model, err)
		}
		deadline, ok := r.Context().Deadline()
		remaining := time.Until(deadline)
		if !ok || remaining > wantTimeout || remaining < wantTimeout-time.Second {
			t.Fatalf("request timeout=%v want=%v", remaining, wantTimeout)
		}
		return response(200, fixture), nil
	})
	args := evaluationArgs("--threshold", "0", "--json=false", "--pick", "")
	if code := a.Run(context.Background(), args); code != 0 || out.String() != "충돌함\n" {
		t.Fatalf("code=%d stdout=%q stderr=%q", code, out, errOut)
	}
	s := c.Resolve(" ")
	if s.APIKey != "file-token" || s.Model != "file-model" || s.Timeout != "12s" {
		t.Fatal("file defaults were lost")
	}
	wantModel, wantTimeout = "flag-model", 9*time.Second
	out.Reset()
	errOut.Reset()
	args = append(args, "--model", "flag-model", "--timeout", "9s")
	if code := a.Run(context.Background(), args); code != 0 || out.String() != "충돌함\n" {
		t.Fatalf("flag overrides: code=%d stdout=%q stderr=%q", code, out, errOut)
	}
}

func TestPrimitiveConfigCriteriaScope(t *testing.T) {
	threshold, minScore := .9, 1.3
	tests := []struct {
		name    string
		stored  cfg.Config
		args    []string
		fixture string
		code    int
	}{
		{"threshold applies to noul", cfg.Config{Defaults: cfg.Defaults{Threshold: &threshold}}, noulArgs(), noulFixture, 2},
		{"threshold ignored by score", cfg.Config{Defaults: cfg.Defaults{Threshold: &threshold}}, scoreArgs(), scoreFixture, 0},
		{"min_score applies to score", cfg.Config{Defaults: cfg.Defaults{MinScore: &minScore}}, scoreArgs(), scoreFixture, 2},
		{"min_score ignored by choice", cfg.Config{Defaults: cfg.Defaults{MinScore: &minScore}}, evaluationArgs(), choiceFixture, 0},
		{"noul flag overrides file", cfg.Config{Defaults: cfg.Defaults{Threshold: &threshold}}, noulArgs("--threshold", "0"), noulFixture, 0},
		{"score flag overrides file", cfg.Config{Defaults: cfg.Defaults{MinScore: &minScore}}, scoreArgs("--min-score", "0"), scoreFixture, 0},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			a, out, errOut := harness(t)
			if err := cfg.Write(a.path, tc.stored); err != nil {
				t.Fatal(err)
			}
			a.transport = roundTripFunc(func(*http.Request) (*http.Response, error) { return response(200, tc.fixture), nil })
			if code := a.Run(context.Background(), tc.args); code != tc.code {
				t.Fatalf("code=%d stdout=%q stderr=%q", code, out, errOut)
			}
		})
	}
}

func TestConfigFailuresDoNotLeakOrReplace(t *testing.T) {
	a, out, errOut := harness(t)
	threshold, minScore, jsonOutput := .8, 1.5, true
	model, timeout := "retained-model", "45s"
	pick := []string{"answer", "model"}
	if err := cfg.Write(a.path, cfg.Config{Defaults: cfg.Defaults{Threshold: &threshold, MinScore: &minScore, JSON: &jsonOutput, Timeout: &timeout, Pick: &pick}, Providers: map[string]cfg.ProviderConfig{"jev": {APIKey: "retained-secret", Model: &model}}}); err != nil {
		t.Fatal(err)
	}
	before, err := os.ReadFile(a.path)
	if err != nil {
		t.Fatal(err)
	}
	beforeInfo, err := os.Stat(a.path)
	if err != nil {
		t.Fatal(err)
	}
	for _, args := range [][]string{
		{"config", "set", "providers.jev.api_key", "argument-secret"}, {"config", "set", "defaults.threshold", "NaN"},
		{"config", "set", "defaults.min_score", "-1"}, {"config", "set", "defaults.min_score", "NaN"},
		{"config", "set", "defaults.timeout", "0s"}, {"config", "set", "defaults.json", "maybe"},
		{"config", "set", "defaults.pick", "unknown"}, {"config", "unset", "unknown"},
	} {
		out.Reset()
		errOut.Reset()
		if code := a.Run(context.Background(), args); code != 1 || out.Len() != 0 || strings.Contains(errOut.String(), "secret") {
			t.Fatalf("unsafe failure: code=%d stdout=%q stderr=%q", code, out, errOut)
		}
		after, err := os.ReadFile(a.path)
		if err != nil || !bytes.Equal(before, after) {
			t.Fatal("failed write changed config")
		}
		afterInfo, err := os.Stat(a.path)
		if err != nil || afterInfo.Mode() != beforeInfo.Mode() {
			t.Fatal("failed write changed config permissions")
		}
	}
	for _, raw := range []string{`api_key = "source-secret`, `threshold = "source-secret"`, `unknown = "source-secret"`} {
		if err := os.WriteFile(a.path, []byte(raw), 0600); err != nil {
			t.Fatal(err)
		}
		out.Reset()
		errOut.Reset()
		if code := a.Run(context.Background(), []string{"config"}); code != 1 || out.Len() != 0 || strings.Contains(errOut.String(), "source-secret") {
			t.Fatalf("TOML failure leaked source: %d %q %q", code, out, errOut)
		}
	}
	if code := a.Run(context.Background(), []string{"config", "path"}); code != 0 {
		t.Fatal("path depends on valid config")
	}
}

func TestConfigPreservesExplicitEmptyPick(t *testing.T) {
	a, _, errOut := harness(t)
	for _, args := range [][]string{
		{"config", "set", "defaults.pick", ""},
		{"config", "set", "defaults.threshold", "0"},
	} {
		if code := a.Run(context.Background(), args); code != 0 {
			t.Fatalf("%v: %d %s", args, code, errOut)
		}
		c, err := cfg.Read(a.path)
		if err != nil || c.Defaults.Pick == nil || len(*c.Defaults.Pick) != 0 {
			t.Fatalf("explicit empty pick not preserved: %v", err)
		}
	}
	if code := a.Run(context.Background(), []string{"config", "unset", "defaults.pick"}); code != 0 {
		t.Fatalf("unset: %d %s", code, errOut)
	}
	c, err := cfg.Read(a.path)
	if err != nil || c.Defaults.Pick != nil {
		t.Fatal("unset did not remove pick")
	}
}
