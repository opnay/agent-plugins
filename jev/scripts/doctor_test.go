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

func TestDoctorIsOfflineAndReadOnly(t *testing.T) {
	source := buildTestBinary(t)
	for _, scenario := range []string{"healthy", "no-config", "missing-binary", "missing-path", "shadowed-path", "no-token", "invalid-config", "public-config", "pick-without-json", "not-executable"} {
		t.Run(scenario, func(t *testing.T) {
			a, out, errOut := harness(t)
			a.envToken = ""
			a.transport = roundTripFunc(func(*http.Request) (*http.Response, error) { t.Fatal("doctor called API"); return nil, nil })
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
			if !strings.Contains(report, "Offline checks only") {
				t.Fatal("doctor omitted diagnostic scope")
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
