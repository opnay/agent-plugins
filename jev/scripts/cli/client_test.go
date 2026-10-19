package cli

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestAPIRequestContract(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		fixture     string
		state       string
		kind        string
		instruction string
		criteria    any
	}{
		{"noul", noulArgs(), noulFixture, "Is this statement true?", "noul", "Answer whether the proposition or question in state is true.", nil},
		{"choice", choiceArgs(), choiceCommandFixture, "Which relation applies?", "choice", "Select the option that best answers the question in state.", map[string]any{"supported": nil, "conflicting": nil}},
		{"score", scoreArgs(), scoreFixture, "How severe is this?", "score", "Rate the state against the ordered criteria.", []any{"low, minor", "medium", "high"}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			a, _, errOut := harness(t)
			calls := 0
			a.transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
				calls++
				if r.Method != "POST" || r.URL.String() != "https://api.typesafe.ai/v1/systemone" || r.Header.Get("Authorization") != "Bearer test-token" || r.Header.Get("Content-Type") != "application/json" {
					t.Fatal("wrong endpoint or headers")
				}
				data, err := io.ReadAll(r.Body)
				if err != nil {
					t.Fatal(err)
				}
				var request any
				if err := json.Unmarshal(data, &request); err != nil {
					t.Fatal(err)
				}
				question := map[string]any{"type": tc.kind, "instructions": tc.instruction}
				if tc.criteria != nil {
					question["criteria"] = tc.criteria
				}
				want := map[string]any{"state": tc.state, "model": "chosen", "questions": map[string]any{"result": question}}
				if !reflect.DeepEqual(request, want) {
					t.Fatalf("wrong request: %s", data)
				}
				return response(200, tc.fixture), nil
			})
			args := append(append([]string{}, tc.args...), "--model", "chosen")
			if code := a.Run(context.Background(), args); code != 0 || calls != 1 {
				t.Fatalf("code=%d calls=%d stderr=%s", code, calls, errOut)
			}
		})
	}
}

func TestAPIErrorsAreSafeAndNotRetried(t *testing.T) {
	generic := func(status int) string {
		return fmt.Sprintf("API returned HTTP %d (%s)", status, http.StatusText(status))
	}
	for _, tc := range []struct {
		name       string
		status     int
		body, want string
	}{
		{"redirect", 301, `{"detail":"test-token"}`, generic(301)},
		{"unauthorized", 401, `{"detail":"test-token"}`, generic(401)},
		{"validation", 422, `{"detail":[{"loc":["body","questions","test-token","criteria",1,"test-token"],"type":"string_type","msg":"test-token","input":"test-token","ctx":{"secret":"test-token"}}]}`, "body.questions.<question>.criteria[1].<field>: expected a string"},

		{"discriminated union location", 422, `{"detail":[{"loc":["body","questions","test-token","score","criteria",0],"type":"string_type"}]}`, "body.questions.<question>.score.criteria[0]: expected a string"},
		{"malformed location", 422, `{"detail":[{"loc":["body","state",null],"type":"missing"}]}`, generic(422)},
		{"bounded errors", 422, `{"detail":[` + strings.Repeat(`{"loc":["body","state"],"type":"missing"},`, 5) + `{"loc":["body","model"],"type":"missing"}]}`, "body.state: required field missing"},
		{"missing", 422, `{"detail":[{"loc":["body","state"],"type":"missing"}]}`, "body.state: required field missing"},
		{"unknown validation type", 422, `{"detail":[{"loc":["body","state"],"type":"test-token","msg":"test-token"}]}`, generic(422)},
		{"unknown location", 422, `{"detail":[{"loc":["test-token"],"type":"missing"}]}`, generic(422)},
		{"context root", 400, `{"code":"max_tokens_exceeded","input":"test-token"}`, "context limit exceeded"},
		{"context detail object", 400, `{"detail":{"type":"max_tokens_exceeded","msg":"test-token"}}`, "context limit exceeded"},
		{"context error object", 400, `{"error":{"code":"max_tokens_exceeded"}}`, "context limit exceeded"},
		{"context detail code", 400, `{"detail":"max_tokens_exceeded"}`, "context limit exceeded"},
		{"context substring is not code", 400, `{"detail":"test-token max_tokens_exceeded"}`, generic(400)},
		{"unrelated metadata is not code", 400, `{"metadata":{"code":"max_tokens_exceeded"}}`, generic(400)},
		{"too large", 413, `{"detail":"test-token"}`, "request body exceeds 32 MiB"},
		{"rate limited", 429, `{"detail":"test-token"}`, "rate limit exceeded; try again later"},
		{"overloaded", 529, `{"detail":"test-token"}`, "TypeSafe temporarily overloaded; try again later"},
		{"unknown status", 500, `{"detail":"test-token"}`, generic(500)},
		{"malformed", 422, `{"detail":"test-token`, generic(422)},
		{"HTML", 429, "<html>test-token</html>", generic(429)},
		{"trailing JSON", 400, `{"code":"max_tokens_exceeded"} {}`, generic(400)},
		{"nonobject JSON", 422, `["test-token"]`, generic(422)},
		{"oversized body", 422, `{"detail":[{"loc":["body","state"],"type":"missing","input":"` + strings.Repeat("x", 65536) + `test-token"}]}`, generic(422)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			a, out, errOut := harness(t)
			calls := 0
			a.transport = roundTripFunc(func(*http.Request) (*http.Response, error) {
				calls++
				r := response(tc.status, tc.body)
				r.Header.Set("Location", "https://elsewhere.invalid")
				return r, nil
			})
			if code := a.Run(context.Background(), evaluationArgs("--json")); code != 1 || out.Len() != 0 || calls != 1 || strings.Contains(errOut.String(), "test-token") || !strings.Contains(errOut.String(), tc.want) {
				t.Fatalf("code=%d calls=%d stdout=%q stderr=%q", code, calls, out, errOut)
			}
			if tc.name == "bounded errors" && (strings.Count(errOut.String(), "required field missing") != 5 || strings.Contains(errOut.String(), "body.model")) {
				t.Fatalf("unbounded validation details: %s", errOut)
			}
			if tc.want == generic(tc.status) && strings.TrimSpace(errOut.String()) != "jev: "+tc.want {
				t.Fatalf("unsafe or unexpected fallback: %q", errOut)
			}
		})
	}
}

func TestMalformedAPIResponses(t *testing.T) {
	for _, body := range []string{
		"not json test-token", `{}`, strings.Replace(fixture, `"choice":"충돌함"`, `"choice":"unknown"`, 1),
		strings.Replace(fixture, `"choice":"충돌함"`, `"choice":"지지됨"`, 1),
		strings.Replace(fixture, `"confidence":0.12`, `"confidence":null`, 1),
		strings.Replace(fixture, `"confidence":0.12`, `"confidence":2`, 1),
		strings.Replace(fixture, `"충돌함":0.8`, `"충돌함":null`, 1),
		strings.Replace(fixture, `"충돌함":0.8`, `"충돌함":-1`, 1),
		strings.Replace(fixture, `"지지됨":0.2`, `"지지됨":0.1`, 1),
		strings.Replace(fixture, `"지지됨":0.2,`, "", 1),
	} {
		a, out, errOut := harness(t)
		a.transport = roundTripFunc(func(*http.Request) (*http.Response, error) { return response(200, body), nil })
		if code := a.Run(context.Background(), evaluationArgs()); code != 1 || out.Len() != 0 || strings.Contains(errOut.String(), "test-token") {
			t.Fatalf("code=%d stdout=%q stderr=%q", code, out, errOut)
		}
	}
}

func TestPrimitiveMalformedResponses(t *testing.T) {
	tests := []struct {
		args []string
		body string
	}{
		{noulArgs(), strings.Replace(noulFixture, `"noul":0.87`, `"noul":null`, 1)},
		{noulArgs(), strings.Replace(noulFixture, `"noul":0.87`, `"noul":-0.1`, 1)},
		{noulArgs(), strings.Replace(noulFixture, `"type":"noul"`, `"type":"choice"`, 1)},
		{scoreArgs(), strings.Replace(scoreFixture, `"score":1.2`, `"score":2.1`, 1)},
		{scoreArgs(), strings.Replace(scoreFixture, `"score":1.2`, `"score":1.1`, 1)},
		{scoreArgs(), strings.Replace(strings.Replace(scoreFixture, `"score":1.2`, `"score":1.0`, 1), `"2":0.3`, `"2":0.2`, 1)},
		{scoreArgs(), strings.Replace(scoreFixture, `"0":"low, minor"`, `"0":"wrong"`, 1)},
		{scoreArgs(), strings.Replace(scoreFixture, `"confidence":0.5`, `"confidence":null`, 1)},
		{scoreArgs(), strings.Replace(scoreFixture, `"0":0.1,`, "", 1)},
	}
	for _, tc := range tests {
		a, out, errOut := harness(t)
		a.transport = roundTripFunc(func(*http.Request) (*http.Response, error) { return response(200, tc.body), nil })
		if code := a.Run(context.Background(), tc.args); code != 1 || out.Len() != 0 || errOut.Len() == 0 {
			t.Fatalf("code=%d stdout=%q stderr=%q", code, out, errOut)
		}
	}
}

func TestTransportFailureAndTimeout(t *testing.T) {
	for _, timeout := range []bool{false, true} {
		a, out, errOut := harness(t)
		a.transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
			if timeout {
				<-r.Context().Done()
				return nil, r.Context().Err()
			}
			return nil, errors.New("sensitive transport error test-token")
		})
		if code := a.Run(context.Background(), evaluationArgs("--timeout", "1ms")); code != 1 || out.Len() != 0 || strings.Contains(errOut.String(), "test-token") {
			t.Fatalf("code=%d stdout=%q stderr=%q", code, out, errOut)
		}
		if timeout && !strings.Contains(errOut.String(), "timed out") {
			t.Fatalf("lost timeout: %s", errOut)
		}
	}
}

func TestAPITiePreservesChoice(t *testing.T) {
	a, out, errOut := harness(t)
	body := strings.ReplaceAll(strings.ReplaceAll(fixture, ":0.2", ":0.5"), ":0.8", ":0.5")
	a.transport = roundTripFunc(func(*http.Request) (*http.Response, error) { return response(200, body), nil })
	if code := a.Run(context.Background(), evaluationArgs()); code != 0 || out.String() != "충돌함\n" {
		t.Fatalf("tie: %d %s %s", code, out, errOut)
	}
}
