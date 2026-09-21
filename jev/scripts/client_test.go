package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestAPIRequestContract(t *testing.T) {
	a, _, errOut := harness(t)
	calls := 0
	a.transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		calls++
		if r.Method != "POST" || r.URL.String() != endpoint || r.Header.Get("Authorization") != "Bearer test-token" || r.Header.Get("Content-Type") != "application/json" {
			t.Fatal("wrong endpoint or headers")
		}
		data, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		var request evaluationRequest
		if err := json.Unmarshal(data, &request); err != nil {
			t.Fatal(err)
		}
		q := request.Questions["result"]
		if request.State != evaluationArgs()[1] || request.Model != "chosen" || len(request.Questions) != 1 || q.Type != "choice" ||
			q.Instructions != "Select the option that best answers the question in state." || len(q.Criteria) != 2 {
			t.Fatalf("wrong request: %s", data)
		}
		for _, name := range []string{"지지됨", "충돌함"} {
			v, exists := q.Criteria[name]
			if !exists || v != nil {
				t.Fatal("criteria must contain null descriptions")
			}
		}
		if strings.Contains(string(data), "threshold") || strings.Contains(string(data), "test-token") {
			t.Fatal("CLI options or token in body")
		}
		return response(200, fixture), nil
	})
	if code := a.run(context.Background(), evaluationArgs("--model", "chosen", "--threshold", "0.8")); code != 0 || calls != 1 {
		t.Fatalf("code=%d calls=%d stderr=%s", code, calls, errOut)
	}
}

func TestAPIErrorsAreSafeAndNotRetried(t *testing.T) {
	for _, status := range []int{301, 401, 422, 429, 500, 529} {
		t.Run(fmt.Sprint(status), func(t *testing.T) {
			a, out, errOut := harness(t)
			calls := 0
			a.transport = roundTripFunc(func(*http.Request) (*http.Response, error) {
				calls++
				r := response(status, "sensitive echoed body: test-token")
				r.Header.Set("Location", "https://elsewhere.invalid")
				return r, nil
			})
			if code := a.run(context.Background(), evaluationArgs("--json")); code != 1 || out.Len() != 0 || calls != 1 || strings.Contains(errOut.String(), "test-token") {
				t.Fatalf("code=%d calls=%d stdout=%q stderr=%q", code, calls, out, errOut)
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
		strings.Replace(fixture, `"지지됨":0.2,`, "", 1),
	} {
		a, out, errOut := harness(t)
		a.transport = roundTripFunc(func(*http.Request) (*http.Response, error) { return response(200, body), nil })
		if code := a.run(context.Background(), evaluationArgs()); code != 1 || out.Len() != 0 || strings.Contains(errOut.String(), "test-token") {
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
		if code := a.run(context.Background(), evaluationArgs("--timeout", "1ms")); code != 1 || out.Len() != 0 || strings.Contains(errOut.String(), "test-token") {
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
	if code := a.run(context.Background(), evaluationArgs()); code != 0 || out.String() != "충돌함\n" {
		t.Fatalf("tie: %d %s %s", code, out, errOut)
	}
}
