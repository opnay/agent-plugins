package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func inputFile(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "request.json")
	if err := os.WriteFile(path, []byte(content), 0600); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestStructuredSingleRequestAndResult(t *testing.T) {
	tests := []struct {
		name, kind, input, fixture, answer string
	}{
		{"noul", "noul", `{"state":{"id":9007199254740993,"text":"notice"},"instructions":{"question":"Is it urgent?","scope":["text"]},"criteria":{"true":{"description":"time-sensitive"},"false":["routine"]}}`, noulFixture, `0.87`},
		{"choice", "choice", `{"state":[{"text":"notice"}],"instructions":["Select the handler"],"criteria":{"supported":{"when":"evidence agrees"},"conflicting":["evidence disagrees"]}}`, choiceCommandFixture, `"conflicting"`},
		{"score", "score", `{"state":"notice","instructions":{"question":"How severe?"},"criteria":[{"impact":"none","weight":1.0},["blocked",{"weight":9007199254740993}]]}`, `{"model":"jev-fixture","answers":{"result":{"type":"score","score":0.75,"legend":{"1":["blocked",{"weight":9007199254740993}],"0":{"weight":1,"impact":"none"}},"probabilities":{"0":0.25,"1":0.75},"confidence":0.5}}}`, `0.75`},
		{"optional instructions and criteria", "noul", `{"state":["notice"]}`, noulFixture, `0.87`},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			a, out, errOut := harness(t)
			calls := 0
			a.transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
				calls++
				data, err := io.ReadAll(r.Body)
				if err != nil {
					t.Fatal(err)
				}
				var request map[string]json.RawMessage
				if err := json.Unmarshal(data, &request); err != nil {
					t.Fatal(err)
				}
				input, err := decodeJSON([]byte(tc.input))
				if err != nil {
					t.Fatal(err)
				}
				fields := input.(map[string]any)
				state, err := decodeJSON(request["state"])
				if err != nil || !reflect.DeepEqual(state, fields["state"]) {
					t.Fatalf("state was not preserved: %s", data)
				}
				var questions map[string]json.RawMessage
				if err := json.Unmarshal(request["questions"], &questions); err != nil {
					t.Fatal(err)
				}
				question, err := decodeJSON(questions["result"])
				if err != nil {
					t.Fatal(err)
				}
				want := map[string]any{"type": tc.kind, "instructions": fields["instructions"]}
				if fields["criteria"] != nil {
					want["criteria"] = fields["criteria"]
				}
				if !reflect.DeepEqual(question, want) {
					t.Fatalf("question was not preserved: %s", data)
				}
				return response(200, tc.fixture), nil
			})
			if code := a.run(context.Background(), []string{tc.kind, "--input", inputFile(t, tc.input), "--json"}); code != 0 || calls != 1 {
				t.Fatalf("code=%d calls=%d stderr=%s", code, calls, errOut)
			}
			var result map[string]json.RawMessage
			if err := json.Unmarshal(out.Bytes(), &result); err != nil {
				t.Fatal(err)
			}
			if string(result["answer"]) != tc.answer {
				t.Fatalf("wrong answer: %s", out)
			}
			if tc.kind == "score" {
				legend, err := decodeJSON(result["legend"])
				if err != nil {
					t.Fatal(err)
				}
				want, _ := decodeJSON([]byte(`{"0":{"impact":"none","weight":1},"1":["blocked",{"weight":9007199254740993}]}`))
				if !jsonValuesEqual(legend, want) {
					t.Fatalf("structured legend lost: %s", out)
				}
			}
		})
	}
}

func TestStructuredInputErrorsBeforeRequest(t *testing.T) {
	tests := []struct {
		name, kind, input string
		flags             []string
	}{
		{"malformed", "noul", `{"state":`, nil},
		{"two values", "noul", `{"state":"text"} {}`, nil},
		{"not object", "noul", `[]`, nil},
		{"unknown field", "noul", `{"state":"text","assets":{}}`, nil},
		{"noncanonical field", "noul", `{"State":"text"}`, nil},
		{"missing state", "noul", `{}`, nil},
		{"null state", "noul", `{"state":null}`, nil},
		{"scalar state", "noul", `{"state":true}`, nil},
		{"empty state", "noul", `{"state":"  "}`, nil},
		{"bad utf8", "noul", "{\"state\":\"\xff\"}", nil},
		{"numeric overflow", "noul", `{"state":{"value":1e400}}`, nil},
		{"deep state", "noul", `{"state":` + strings.Repeat("[", 65) + `0` + strings.Repeat("]", 65) + `}`, nil},
		{"scalar instructions", "noul", `{"state":"text","instructions":1}`, nil},
		{"noul keys", "noul", `{"state":"text","criteria":{"yes":"yes"}}`, nil},
		{"noul description type", "noul", `{"state":"text","criteria":{"true":true}}`, nil},
		{"choice missing criteria", "choice", `{"state":"text"}`, nil},
		{"choice array", "choice", `{"state":"text","criteria":["a","b"]}`, nil},
		{"choice label", "choice", `{"state":"text","criteria":{" a":null,"b":null}}`, nil},
		{"choice description", "choice", `{"state":"text","criteria":{"a":1,"b":null}}`, nil},
		{"score duplicate", "score", `{"state":"text","criteria":[{"n":1},{"n":1.0}]}`, nil},
		{"score null", "score", `{"state":"text","criteria":[null,"high"]}`, nil},
		{"score empty", "score", `{"state":"text","criteria":[" ","high"]}`, nil},
		{"mixed if", "noul", `{"state":"text"}`, []string{"--if", ""}},
		{"mixed conditions", "choice", `{"state":"text","criteria":{"a":null,"b":null}}`, []string{"--conditions", "a,b"}},
		{"mixed level", "score", `{"state":"text","criteria":["low","high"]}`, []string{"--level", "low"}},
		{"score bound", "score", `{"state":"text","criteria":["low","high"]}`, []string{"--min-score", "2"}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			a, out, errOut := harness(t)
			a.transport = roundTripFunc(func(*http.Request) (*http.Response, error) { t.Fatal("invalid input reached API"); return nil, nil })
			args := append([]string{tc.kind, "--input", inputFile(t, tc.input)}, tc.flags...)
			if code := a.run(context.Background(), args); code != 1 || out.Len() != 0 || errOut.Len() == 0 {
				t.Fatalf("code=%d stdout=%q stderr=%q", code, out, errOut)
			}
		})
	}
}

func TestLongContextHeaders(t *testing.T) {
	for _, batch := range []bool{false, true} {
		for _, flag := range []string{"", "--long-context", "--long-context=false"} {
			t.Run(strings.Join([]string{map[bool]string{false: "single", true: "batch"}[batch], flag}, "/"), func(t *testing.T) {
				a, _, errOut := harness(t)
				args := noulArgs()
				fixture := noulFixture
				if batch {
					state, input := batchFiles(t, "text", `{"id":"result","type":"noul","question":"True?"}`)
					args = batchArgs(state, input)
				}
				if flag != "" {
					args = append(args, flag)
				}
				a.transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
					want := ""
					if flag == "--long-context" {
						want = "true"
					}
					if got := r.Header.Get("TypeSafe-Danger-Zone-Long-Context"); got != want {
						t.Fatalf("header=%q want=%q", got, want)
					}
					return response(200, fixture), nil
				})
				if code := a.run(context.Background(), args); code != 0 {
					t.Fatalf("code=%d stderr=%s", code, errOut)
				}
			})
		}
	}
}

func TestBatchStructuredQuestionsAndJSONState(t *testing.T) {
	input := strings.Join([]string{
		`{"id":"N01","type":"noul","instructions":{"question":"Is it urgent?"},"criteria":{"true":{"when":"urgent"},"false":null}}`,
		`{"id":"C01","type":"choice","question":"Which handler?","criteria":{"a,b":{"description":"first"},"other":["second"]}}`,
		`{"id":"S01","type":"score","instructions":["How high?"],"criteria":[{"impact":"low"},["high"]]}`,
	}, "\n")
	state := `{"id":9007199254740993,"messages":["notice"]}`
	for _, format := range []string{"text", "json"} {
		t.Run(format, func(t *testing.T) {
			statePath, inputPath := batchFiles(t, state, input)
			a, out, errOut := harness(t)
			calls := 0
			a.transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
				calls++
				data, _ := io.ReadAll(r.Body)
				request, err := decodeJSON(data)
				if err != nil {
					t.Fatal(err)
				}
				fields := request.(map[string]any)
				var wantState any = state
				if format == "json" {
					wantState, _ = decodeJSON([]byte(state))
				}
				if !reflect.DeepEqual(fields["state"], wantState) {
					t.Fatalf("wrong state: %s", data)
				}
				questions := fields["questions"].(map[string]any)
				for _, line := range strings.Split(input, "\n") {
					row, _ := decodeJSON([]byte(line))
					q := row.(map[string]any)
					instructions, exists := q["instructions"]
					if !exists {
						instructions = q["question"]
					}
					want := map[string]any{"type": q["type"], "instructions": instructions, "criteria": q["criteria"]}
					if !reflect.DeepEqual(questions[q["id"].(string)], want) {
						t.Fatalf("wrong question: %s", data)
					}
				}
				fixture := strings.Replace(batchFixture, `"0":"low","1":"high"`, `"0":{"impact":"low"},"1":["high"]`, 1)
				return response(200, fixture), nil
			})
			if code := a.run(context.Background(), batchArgs(statePath, inputPath, "--state-format", format)); code != 0 || calls != 1 {
				t.Fatalf("code=%d calls=%d stderr=%s", code, calls, errOut)
			}
			rows := strings.Split(strings.TrimSpace(out.String()), "\n")
			if len(rows) != 3 || !strings.Contains(rows[0], `"id":"N01"`) || !strings.Contains(rows[2], `"legend":{"0":{"impact":"low"},"1":["high"]}`) {
				t.Fatalf("wrong ordered output: %s", out)
			}
		})
	}
}

func TestBatchStructuredFailuresAreAtomic(t *testing.T) {
	valid := `{"id":"A","type":"noul","instructions":null}` + "\n"
	for _, bad := range []string{
		`{"id":"B","type":"noul","question":"q","instructions":{}}`,
		`{"id":"B","type":"noul","instructions":true}`,
		`{"id":"B","type":"choice","instructions":{},"criteria":{"a":null,"b":null},"conditions":["a","b"]}`,
		`{"id":"B","type":"score","instructions":{},"criteria":["low",null]}`,
		`{"id":"B","type":"noul","instructions":{},"criteria":{"true":{"n":1e400}}}`,
	} {
		a, out, errOut := harness(t)
		state, input := batchFiles(t, `{"text":"notice"}`, valid+bad)
		a.transport = roundTripFunc(func(*http.Request) (*http.Response, error) { t.Fatal("invalid batch reached API"); return nil, nil })
		if code := a.run(context.Background(), batchArgs(state, input, "--state-format", "json")); code != 1 || out.Len() != 0 || errOut.Len() == 0 {
			t.Fatalf("code=%d stdout=%s stderr=%s", code, out, errOut)
		}
	}
	for _, state := range []string{`{`, `null`, `1`, `{} {}`, `{"n":1e400}`} {
		a, out, errOut := harness(t)
		statePath, input := batchFiles(t, state, valid)
		a.transport = roundTripFunc(func(*http.Request) (*http.Response, error) { t.Fatal("invalid state reached API"); return nil, nil })
		if code := a.run(context.Background(), batchArgs(statePath, input, "--state-format", "json")); code != 1 || out.Len() != 0 {
			t.Fatalf("code=%d stdout=%s stderr=%s", code, out, errOut)
		}
	}
}

func TestStructuredScoreRejectsChangedLegend(t *testing.T) {
	for _, legend := range []string{`{"0":{"n":9007199254740992},"1":["high"]}`, `{"0":{"n":9007199254740993},"1":["other"]}`} {
		a, out, errOut := harness(t)
		input := inputFile(t, `{"state":"text","criteria":[{"n":9007199254740993},["high"]]}`)
		a.transport = roundTripFunc(func(*http.Request) (*http.Response, error) {
			return response(200, `{"model":"jev-fixture","answers":{"result":{"type":"score","score":0.75,"legend":`+legend+`,"probabilities":{"0":0.25,"1":0.75},"confidence":0.5}}}`), nil
		})
		if code := a.run(context.Background(), []string{"score", "--input", input}); code != 1 || out.Len() != 0 {
			t.Fatalf("code=%d stdout=%s stderr=%s", code, out, errOut)
		}
	}
}
