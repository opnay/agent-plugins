package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"path/filepath"
	"strings"
	"testing"
)

const fixture = `{"model":"jev-fixture","answers":{"result":{"type":"choice","choice":"충돌함","probabilities":{"지지됨":0.2,"충돌함":0.8},"confidence":0.12}}}`

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func response(status int, body string) *http.Response {
	return &http.Response{StatusCode: status, Header: make(http.Header), Body: io.NopCloser(strings.NewReader(body))}
}

func harness(t *testing.T) (application, *bytes.Buffer, *bytes.Buffer) {
	t.Helper()
	out, errOut := new(bytes.Buffer), new(bytes.Buffer)
	return application{
		in: strings.NewReader(""), out: out, errOut: errOut,
		path: filepath.Join(t.TempDir(), ".agents", "jev.toml"), envToken: "test-token",
		transport: roundTripFunc(func(*http.Request) (*http.Response, error) { return response(200, fixture), nil }),
	}, out, errOut
}

func evaluationArgs(extra ...string) []string {
	return append([]string{"--if", "권장하지만 필수는 아니다. 반드시 사용해야 하는가?", "--conditions", "지지됨,충돌함"}, extra...)
}

func TestEvaluationOutputAndThreshold(t *testing.T) {
	tests := []struct {
		name string
		args []string
		code int
		want string
	}{
		{"plain", nil, 0, "충돌함\n"},
		{"equal threshold uses probability not confidence", []string{"--threshold", "0.8"}, 0, "충돌함\n"},
		{"zero", []string{"--threshold", "0"}, 0, "충돌함\n"},
		{"threshold failure", []string{"--threshold", "0.81"}, 2, ""},
		{"JSON threshold failure", []string{"--threshold", "1", "--json", "--pick", "answer"}, 2, ""},
		{"pick one", []string{"--json", "--pick", "answer"}, 0, "{\"answer\":\"충돌함\"}\n"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			a, out, errOut := harness(t)
			if code := a.run(context.Background(), evaluationArgs(tc.args...)); code != tc.code || out.String() != tc.want {
				t.Fatalf("code=%d stdout=%q stderr=%q", code, out, errOut)
			}
			if tc.code == 0 && errOut.Len() != 0 || tc.code == 2 && !strings.Contains(errOut.String(), "threshold not met") {
				t.Fatalf("unexpected stderr: %q", errOut)
			}
		})
	}
}

func TestJSONFields(t *testing.T) {
	for _, pick := range []string{"", "answer,probabilities", "confidence,model"} {
		t.Run(pick, func(t *testing.T) {
			a, out, errOut := harness(t)
			if code := a.run(context.Background(), evaluationArgs("--json", "--pick", pick)); code != 0 {
				t.Fatalf("code=%d stderr=%s", code, errOut)
			}
			var fields map[string]json.RawMessage
			if err := json.Unmarshal(out.Bytes(), &fields); err != nil {
				t.Fatal(err)
			}
			expected := []string{"answer", "probabilities", "confidence", "model"}
			if pick != "" {
				expected = splitList(pick)
			}
			if len(fields) != len(expected) {
				t.Fatalf("fields=%s", out)
			}
			for _, key := range expected {
				if _, ok := fields[key]; !ok {
					t.Fatalf("missing %s", key)
				}
			}
			if raw, ok := fields["confidence"]; ok && string(raw) != "0.12" {
				t.Fatal("confidence was not preserved")
			}
		})
	}
}

func TestInvalidOptionsDoNotCallAPI(t *testing.T) {
	tests := [][]string{
		{"--if", "q"}, {"--conditions", "a,b"}, {"--if", " ", "--conditions", "a,b"},
		{"--if", "q", "--conditions", "a,a"}, {"--if", "q", "--conditions", "a,"},
		{"--if", "q", "--conditions", "a"},
		evaluationArgs("--threshold", "NaN"), evaluationArgs("--threshold", "Inf"),
		evaluationArgs("--threshold", "-1"), evaluationArgs("--threshold", "1.1"),
		evaluationArgs("--timeout", "0s"), evaluationArgs("--timeout", "bad"),
		evaluationArgs("--model", ""), evaluationArgs("--pick", "answer"),
		evaluationArgs("--json", "--pick", "answer,answer"),
		evaluationArgs("--json", "--pick", "answer,"), evaluationArgs("--json", "--pick", "unknown"),
		evaluationArgs("extra"), evaluationArgs("--unknown"),
	}
	for _, args := range tests {
		t.Run(strings.Join(args, " "), func(t *testing.T) {
			a, out, errOut := harness(t)
			a.transport = roundTripFunc(func(*http.Request) (*http.Response, error) { t.Fatal("unexpected API call"); return nil, nil })
			if code := a.run(context.Background(), args); code != 1 || out.Len() != 0 || errOut.Len() == 0 {
				t.Fatalf("code=%d stdout=%q stderr=%q", code, out, errOut)
			}
		})
	}
}

func TestMissingAuth(t *testing.T) {
	a, out, errOut := harness(t)
	a.envToken = ""
	if code := a.run(context.Background(), evaluationArgs()); code != 1 || out.Len() != 0 || !strings.Contains(errOut.String(), "API key missing") {
		t.Fatalf("code=%d stdout=%s stderr=%s", code, out, errOut)
	}
}

func TestHelpDoesNotReadConfig(t *testing.T) {
	for _, args := range [][]string{nil, {"--help"}, {"config", "--help"}} {
		a, out, errOut := harness(t)
		a.path = t.TempDir() // Not a config file.
		if code := a.run(context.Background(), args); code != 0 || !strings.Contains(out.String(), "Usage:") || errOut.Len() > 0 {
			t.Fatalf("code=%d stdout=%s stderr=%s", code, out, errOut)
		}
	}
}
