package main

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestDoctorLocalChecksAreReadOnly(t *testing.T) {
	source := buildTestBinary(t)
	for _, scenario := range []string{"healthy", "no-config", "missing-binary", "missing-path", "shadowed-path", "no-token", "invalid-config", "public-config", "pick-without-json", "not-executable"} {
		t.Run(scenario, func(t *testing.T) {
			a, out, errOut := harness(t)
			a.envToken = ""
			calls := 0
			a.transport = roundTripFunc(func(*http.Request) (*http.Response, error) { calls++; return response(200, ""), nil })
			target := filepath.Join(t.TempDir(), "jev")
			if _, err := installExecutable(source, target, false); err != nil {
				t.Fatal(err)
			}
			a.lookupPath = func(string) (string, error) { return target, nil }
			if err := writeConfig(a.path, config{APIKey: "never-print-this-token"}); err != nil {
				t.Fatal(err)
			}
			switch scenario {
			case "no-config":
				if err := os.Remove(a.path); err != nil {
					t.Fatal(err)
				}
				a.envToken = "environment-secret"
			case "missing-binary":
				if err := os.Remove(target); err != nil {
					t.Fatal(err)
				}
			case "missing-path":
				a.lookupPath = func(string) (string, error) { return "", exec.ErrNotFound }
			case "shadowed-path":
				shadow := filepath.Join(t.TempDir(), "jev")
				if err := os.WriteFile(shadow, []byte("shadow"), 0755); err != nil {
					t.Fatal(err)
				}
				a.lookupPath = func(string) (string, error) { return shadow, nil }
			case "no-token":
				if err := writeConfig(a.path, config{}); err != nil {
					t.Fatal(err)
				}
			case "invalid-config":
				if err := os.WriteFile(a.path, []byte(`api_key = "never-print-this-token`), 0600); err != nil {
					t.Fatal(err)
				}
			case "public-config":
				if err := os.Chmod(a.path, 0644); err != nil {
					t.Fatal(err)
				}
			case "pick-without-json":
				if err := writeConfig(a.path, config{APIKey: "never-print-this-token", Pick: &[]string{"answer"}}); err != nil {
					t.Fatal(err)
				}
			case "not-executable":
				if err := os.Chmod(target, 0600); err != nil {
					t.Fatal(err)
				}
			}
			before, beforeErr := os.ReadFile(a.path)
			beforeInfo, _ := os.Stat(a.path)
			code := a.run(context.Background(), []string{"doctor", "--dir", filepath.Dir(target)})
			want := 1
			if scenario == "healthy" || scenario == "no-config" {
				want = 0
			}
			if code != want || want == 1 && out.Len() > 0 || want == 0 && errOut.Len() > 0 {
				t.Fatalf("code=%d stdout=%s stderr=%s", code, out, errOut)
			}
			report := out.String() + errOut.String()
			if strings.Contains(report, "never-print-this-token") || strings.Contains(report, "environment-secret") {
				t.Fatal("doctor leaked token")
			}
			if !strings.Contains(report, "No local changes or inference requests were made.") {
				t.Fatal("doctor omitted diagnostic scope")
			}
			wantCalls := 1
			if scenario == "no-token" || scenario == "invalid-config" || scenario == "public-config" || scenario == "pick-without-json" {
				wantCalls = 0
			}
			if calls != wantCalls {
				t.Fatalf("API calls=%d want=%d", calls, wantCalls)
			}
			after, afterErr := os.ReadFile(a.path)
			if errors.Is(beforeErr, os.ErrNotExist) {
				if !errors.Is(afterErr, os.ErrNotExist) {
					t.Fatal("doctor created config")
				}
			} else {
				afterInfo, err := os.Stat(a.path)
				if beforeErr != nil || afterErr != nil || err != nil || !bytes.Equal(before, after) || beforeInfo.Mode() != afterInfo.Mode() {
					t.Fatal("doctor mutated config")
				}
			}
		})
	}
}

func doctorHarness(t *testing.T, source string) (application, *bytes.Buffer, *bytes.Buffer, string) {
	t.Helper()
	a, out, errOut := harness(t)
	target := filepath.Join(t.TempDir(), "jev")
	if _, err := installExecutable(source, target, false); err != nil {
		t.Fatal(err)
	}
	a.lookupPath = func(string) (string, error) { return target, nil }
	return a, out, errOut, filepath.Dir(target)
}

func TestDoctorAuthentication(t *testing.T) {
	source := buildTestBinary(t)
	for _, tc := range []struct {
		name   string
		status int
		code   int
		detail string
		hint   string
	}{
		{"authenticated", 200, 0, "authenticated (GET " + modelsEndpoint + ")", "Summary: 5 OK, 0 FAIL, 0 SKIP"},
		{"rejected", 401, 1, "API key rejected (invalid or expired)", "update or unset TYPESAFE_API_KEY"},
		{"forbidden", 403, 1, "access denied", "check API key permissions and account access"},
		{"rate-limit", 429, 1, "rate limited; key validity unknown", "wait for the rate limit"},
		{"server-error", 500, 1, "key validity unknown", "check TypeSafe service status"},
		{"overloaded", 529, 1, "key validity unknown", "check TypeSafe service status"},
		{"redirect", 302, 1, "key validity unknown", "check TypeSafe service status"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			a, out, errOut, dir := doctorHarness(t, source)
			if err := writeConfig(a.path, config{APIKey: "file-secret"}); err != nil {
				t.Fatal(err)
			}
			a.envToken = "environment-secret"
			calls := 0
			a.transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
				calls++
				if r.Method != http.MethodGet || r.URL.String() != modelsEndpoint || r.Body != nil ||
					r.Header.Get("Authorization") != "Bearer environment-secret" || r.Header.Get("Accept") != "application/json" {
					t.Fatal("wrong doctor authentication request")
				}
				rsp := response(tc.status, "echoed secrets: environment-secret file-secret")
				rsp.Header.Set("Location", "https://elsewhere.invalid")
				return rsp, nil
			})
			if code := a.run(context.Background(), []string{"doctor", "--dir", dir}); code != tc.code || calls != 1 {
				t.Fatalf("code=%d calls=%d stderr=%s", code, calls, errOut)
			}
			report := out.String() + errOut.String()
			for _, want := range []string{"Local setup", "Authentication", "TYPESAFE_API_KEY (overrides config api_key", tc.detail, tc.hint} {
				if !strings.Contains(report, want) {
					t.Fatalf("missing %q: %s", want, report)
				}
			}
			if tc.code == 0 && errOut.Len() != 0 || tc.code == 1 && out.Len() != 0 {
				t.Fatalf("wrong report stream: stdout=%s stderr=%s", out, errOut)
			}
			if strings.Contains(report, "environment-secret") || strings.Contains(report, "file-secret") || strings.Contains(report, "echoed secrets") {
				t.Fatal("doctor exposed an API body or key")
			}
		})
	}
}

func TestDoctorKeySelection(t *testing.T) {
	source := buildTestBinary(t)
	for _, env := range []string{"", "   ", "environment-secret"} {
		t.Run(env, func(t *testing.T) {
			a, out, errOut, dir := doctorHarness(t, source)
			a.envToken = env
			if err := writeConfig(a.path, config{APIKey: "file-secret"}); err != nil {
				t.Fatal(err)
			}
			key, wantSource, wantFix := "file-secret", "config api_key (value hidden)", "use jev config set api_key --stdin"
			if strings.TrimSpace(env) != "" {
				key, wantSource, wantFix = env, "TYPESAFE_API_KEY (overrides config api_key", "update or unset TYPESAFE_API_KEY"
			}
			a.transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
				if r.Header.Get("Authorization") != "Bearer "+key {
					t.Fatal("wrong selected key")
				}
				return response(401, ""), nil
			})
			if code := a.run(context.Background(), []string{"doctor", "--dir", dir}); code != 1 || out.Len() > 0 ||
				!strings.Contains(errOut.String(), wantSource) || !strings.Contains(errOut.String(), wantFix) {
				t.Fatalf("code=%d stdout=%s stderr=%s", code, out, errOut)
			}
		})
	}
}

func TestDoctorNetworkErrors(t *testing.T) {
	source := buildTestBinary(t)
	for _, scenario := range []string{"network", "timeout", "canceled"} {
		t.Run(scenario, func(t *testing.T) {
			a, out, errOut, dir := doctorHarness(t, source)
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			want := "cannot reach TypeSafe API"
			if scenario == "timeout" {
				timeout := "1ms"
				if err := writeConfig(a.path, config{Timeout: &timeout}); err != nil {
					t.Fatal(err)
				}
				want = "authentication check timed out"
			}
			if scenario == "canceled" {
				cancel()
				want = "authentication check canceled"
			}
			calls := 0
			a.transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
				calls++
				if scenario != "network" {
					<-r.Context().Done()
					return nil, r.Context().Err()
				}
				return nil, errors.New("transport error containing test-token")
			})
			if code := a.run(ctx, []string{"doctor", "--dir", dir}); code != 1 || out.Len() != 0 || calls != 1 ||
				!strings.Contains(errOut.String(), want) || !strings.Contains(errOut.String(), "key validity unknown") || strings.Contains(errOut.String(), "test-token") {
				t.Fatalf("code=%d calls=%d stdout=%s stderr=%s", code, calls, out, errOut)
			}
		})
	}
}

func TestDoctorRejectsOfflineAndDoesNotCallAPI(t *testing.T) {
	a, out, errOut := harness(t)
	a.transport = roundTripFunc(func(*http.Request) (*http.Response, error) { t.Fatal("invalid flags called API"); return nil, nil })
	if code := a.run(context.Background(), []string{"doctor", "--offline"}); code != 1 || out.Len() != 0 || !strings.Contains(errOut.String(), "jev doctor --help") {
		t.Fatalf("code=%d stdout=%s stderr=%s", code, out, errOut)
	}
}
