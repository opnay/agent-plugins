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

const choiceFixture = `{"model":"jev-fixture","answers":{"result":{"type":"choice","choice":"충돌함","probabilities":{"지지됨":0.2,"충돌함":0.8},"confidence":0.12}}}`
const choiceCommandFixture = `{"model":"jev-fixture","answers":{"result":{"type":"choice","choice":"conflicting","probabilities":{"supported":0.2,"conflicting":0.8},"confidence":0.12}}}`
const noulFixture = `{"model":"jev-fixture","answers":{"result":{"type":"noul","noul":0.87}}}`
const scoreFixture = `{"model":"jev-fixture","answers":{"result":{"type":"score","score":1.2,"legend":{"0":"low, minor","1":"medium","2":"high"},"probabilities":{"0":0.1,"1":0.6,"2":0.3},"confidence":0.5}}}`
const fixture = choiceFixture

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

func choiceArgs(extra ...string) []string {
	return append([]string{"choice", "--if", "Which relation applies?", "--conditions", "supported,conflicting"}, extra...)
}

func noulArgs(extra ...string) []string {
	return append([]string{"noul", "--if", "Is this statement true?"}, extra...)
}

func scoreArgs(extra ...string) []string {
	return append([]string{"score", "--if", "How severe is this?", "--level", "low, minor", "--level", "medium", "--level", "high"}, extra...)
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

func TestPrimitiveOutputAndCriteria(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		fixture string
		code    int
		want    string
	}{
		{"explicit choice", choiceArgs(), choiceCommandFixture, 0, "conflicting\n"},
		{"noul", noulArgs(), noulFixture, 0, "0.87\n"},
		{"noul threshold", noulArgs("--threshold", "0.88"), noulFixture, 2, ""},
		{"score", scoreArgs(), scoreFixture, 0, "1.2\n"},
		{"score minimum", scoreArgs("--min-score", "1.21"), scoreFixture, 2, ""},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			a, out, errOut := harness(t)
			a.transport = roundTripFunc(func(*http.Request) (*http.Response, error) { return response(200, tc.fixture), nil })
			if code := a.run(context.Background(), tc.args); code != tc.code || out.String() != tc.want {
				t.Fatalf("code=%d stdout=%q stderr=%q", code, out, errOut)
			}
			if tc.code == 2 && errOut.Len() == 0 {
				t.Fatal("criterion failure omitted stderr")
			}
		})
	}
}

func TestPrimitiveJSONFields(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		fixture string
		fields  []string
	}{
		{"noul", noulArgs("--json"), noulFixture, []string{"answer", "model"}},
		{"noul pick", noulArgs("--json", "--pick", "answer"), noulFixture, []string{"answer"}},
		{"score", scoreArgs("--json"), scoreFixture, []string{"answer", "legend", "probabilities", "confidence", "model"}},
		{"score pick", scoreArgs("--json", "--pick", "answer,legend"), scoreFixture, []string{"answer", "legend"}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			a, out, errOut := harness(t)
			a.transport = roundTripFunc(func(*http.Request) (*http.Response, error) { return response(200, tc.fixture), nil })
			if code := a.run(context.Background(), tc.args); code != 0 {
				t.Fatalf("code=%d stderr=%s", code, errOut)
			}
			var fields map[string]json.RawMessage
			if err := json.Unmarshal(out.Bytes(), &fields); err != nil || len(fields) != len(tc.fields) {
				t.Fatalf("fields=%s err=%v", out, err)
			}
			for _, field := range tc.fields {
				if _, ok := fields[field]; !ok {
					t.Fatalf("missing %s", field)
				}
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
		choiceArgs("--level", "low"), choiceArgs("--json", "--pick", "legend"),
		noulArgs("--conditions", "yes,no"), noulArgs("--json", "--pick", "confidence"),
		scoreArgs("--conditions", "a,b"), scoreArgs("--threshold", "0.5"), scoreArgs("--min-score", "3"),
		{"score", "--if", "q", "--level", "only one"},
		{"score", "--if", "q", "--level", "same", "--level", "same"},
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
	categories := []string{
		"Evaluation:", "Configuration:", "Maintenance:", "Common evaluation options:",
		"Primitive options:", "Config keys:", "Output:", "Exit codes:",
	}
	for _, args := range [][]string{nil, {"--help"}} {
		a, out, errOut := harness(t)
		a.path = t.TempDir() // Not a config file.
		a.envToken = ""
		a.transport = roundTripFunc(func(*http.Request) (*http.Response, error) {
			t.Fatal("help called API")
			return nil, nil
		})
		if code := a.run(context.Background(), args); code != 0 || !strings.Contains(out.String(), "Usage:") || errOut.Len() > 0 {
			t.Fatalf("code=%d stdout=%s stderr=%s", code, out, errOut)
		}
		for _, category := range categories {
			if !strings.Contains(out.String(), category) {
				t.Fatalf("%v help omitted %s", args, category)
			}
		}
	}

	tests := []struct {
		command string
		want    []string
		reject  []string
	}{
		{
			"noul",
			[]string{"Usage:\n  jev noul", "Noul options:", "--threshold <0..1>", "Common options:", "Output:", "probability of yes", "JSON:    answer, model", "Exit codes:"},
			[]string{"Choice options:", "Score options:"},
		},
		{
			"choice",
			[]string{"Usage:\n  jev choice", "Choice options:", "--conditions <a,b,...>", "--threshold <0..1>", "Common options:", "Output:", "selected label", "probabilities, confidence", "Exit codes:"},
			[]string{"Noul options:", "Score options:"},
		},
		{
			"score",
			[]string{"Usage:\n  jev score", "Score options:", "--level <description>", "--min-score <number>", "Common options:", "Output:", "weighted score", "legend, probabilities, confidence", "Exit codes:"},
			[]string{"Noul options:", "Choice options:"},
		},
		{
			"batch",
			[]string{"Usage:\n  jev batch", "Batch options:", "--state-file <path>", "sent unchanged", "--input <path>", "blank lines ignored", `"type":"noul"`, `"type":"choice"`, `"type":"score"`, "2–255 distinct", "2–10 distinct", "forbids conditions", "Unknown or differently cased keys", "one API request", "Noul:   id, type, answer, model", "Choice: id, type, answer, probabilities, confidence, model", "Score:  id, type, answer, legend, probabilities, confidence, model", "does not apply threshold", "does not chunk, retry, resume, or return partial results", "encoded before stdout", "leave stdout empty", "partial write cannot be rolled back", "Exit codes:"},
			[]string{"Noul options:", "Choice options:", "Score options:", "--pick"},
		},
		{
			"config",
			[]string{"Usage:\n  jev config", "Commands:", "Options:", "Keys:", "api_key", "threshold", "min_score", "model        jev-latest", "timeout      30s", "json         false", "TYPESAFE_API_KEY takes precedence", "Flags override file values"},
			[]string{"Evaluation:", "Maintenance:"},
		},
		{
			"install",
			[]string{"Usage:\n  jev install", "Options:", "--force", "--dir <directory>", "<directory>/jev", "mode 0755", "does not download, build, edit PATH, or change config and tokens"},
			[]string{"uninstall", "Doctor"},
		},
		{
			"uninstall",
			[]string{"Usage:\n  jev uninstall", "Options:", "--dir <directory>", "only when Go build information identifies it as Jev", "Config, tokens, the directory, adjacent files, and", "shell settings are preserved"},
			[]string{"--force", "Doctor"},
		},
		{
			"doctor",
			[]string{"Usage:\n  jev doctor", "Options:", "--dir <directory>", "PATH selection", "read-only and offline", "does not authenticate", "or repair state"},
			[]string{"--force", "Removes <directory>/jev"},
		},
	}
	for _, tc := range tests {
		t.Run(tc.command, func(t *testing.T) {
			a, out, errOut := harness(t)
			a.path = t.TempDir() // Not a config file.
			a.envToken = ""
			a.transport = roundTripFunc(func(*http.Request) (*http.Response, error) {
				t.Fatal("help called API")
				return nil, nil
			})
			if code := a.run(context.Background(), []string{tc.command, "--help"}); code != 0 || errOut.Len() != 0 {
				t.Fatalf("code=%d stdout=%s stderr=%s", code, out, errOut)
			}
			for _, want := range tc.want {
				if !strings.Contains(out.String(), want) {
					t.Fatalf("help omitted %q:\n%s", want, out)
				}
			}
			for _, reject := range tc.reject {
				if strings.Contains(out.String(), reject) {
					t.Fatalf("help included sibling detail %q:\n%s", reject, out)
				}
			}
		})
	}
}

func TestOptionErrorsPointToSubcommandHelp(t *testing.T) {
	tests := []struct {
		args []string
		want string
	}{
		{[]string{"noul", "--unknown"}, "jev noul --help"},
		{[]string{"noul", "--if", " "}, "jev noul --help"},
		{noulArgs("--threshold", "1.1"), "jev noul --help"},
		{noulArgs("--json", "--pick", "confidence"), "jev noul --help"},
		{[]string{"choice", "unexpected"}, "jev choice --help"},
		{[]string{"choice", "--if", "q", "--conditions", "one"}, "jev choice --help"},
		{choiceArgs("--pick", "answer"), "jev choice --help"},
		{[]string{"score", "--unknown"}, "jev score --help"},
		{[]string{"score", "--if", "q", "--level", "one"}, "jev score --help"},
		{scoreArgs("--min-score", "3"), "jev score --help"},
		{[]string{"batch", "--state-file", "state"}, "jev batch --help"},
		{[]string{"batch", "--input", "input"}, "jev batch --help"},
		{[]string{"batch", "--state-file", "state", "--input", "input", "--json"}, "jev batch --help"},
		{[]string{"batch", "--state-file", "state", "--input", "input", "extra"}, "jev batch --help"},
		{[]string{"config", "unknown"}, "jev config --help"},
		{[]string{"config", "set", "threshold", "bad"}, "jev config --help"},
		{[]string{"config", "set", "unknown", "value"}, "jev config --help"},
		{[]string{"config", "unset", "unknown"}, "jev config --help"},
		{[]string{"install", "--unknown"}, "jev install --help"},
		{[]string{"install", "--dir="}, "jev install --help"},
		{[]string{"uninstall", "unexpected"}, "jev uninstall --help"},
		{[]string{"doctor", "--unknown"}, "jev doctor --help"},
	}
	for _, tc := range tests {
		t.Run(strings.Join(tc.args, " "), func(t *testing.T) {
			a, out, errOut := harness(t)
			if code := a.run(context.Background(), tc.args); code != 1 || out.Len() != 0 || !strings.Contains(errOut.String(), tc.want) {
				t.Fatalf("code=%d stdout=%q stderr=%q", code, out, errOut)
			}
		})
	}
}

func TestConfigArgumentErrorsPrecedeConfigRead(t *testing.T) {
	tests := [][]string{
		{"config", "unknown"},
		{"config", "set", "threshold"},
		{"config", "set", "threshold", "bad"},
		{"config", "set", "unknown", "value"},
		{"config", "unset", "unknown"},
	}
	for _, args := range tests {
		t.Run(strings.Join(args, " "), func(t *testing.T) {
			a, out, errOut := harness(t)
			a.path = t.TempDir() // A directory would fail if config were read first.
			if code := a.run(context.Background(), args); code != 1 || out.Len() != 0 || !strings.Contains(errOut.String(), "jev config --help") {
				t.Fatalf("code=%d stdout=%q stderr=%q", code, out, errOut)
			}
		})
	}
}

func TestEvaluationFlagErrorsPrecedeConfigRead(t *testing.T) {
	tests := [][]string{
		noulArgs("--threshold", "1.1"),
		noulArgs("--json", "--pick", "confidence"),
		choiceArgs("--model", ""),
		choiceArgs("--timeout", "0s"),
		choiceArgs("--json=false", "--pick", "answer"),
		scoreArgs("--min-score", "3"),
	}
	for _, args := range tests {
		t.Run(strings.Join(args, " "), func(t *testing.T) {
			a, out, errOut := harness(t)
			a.path = t.TempDir() // A directory would fail if config were read first.
			want := "jev " + args[0] + " --help"
			if code := a.run(context.Background(), args); code != 1 || out.Len() != 0 || !strings.Contains(errOut.String(), want) {
				t.Fatalf("code=%d stdout=%q stderr=%q", code, out, errOut)
			}
		})
	}
}
