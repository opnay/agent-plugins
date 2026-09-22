package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const batchFixture = `{"model":"jev-batch","answers":{"C01":{"type":"choice","choice":"a,b","probabilities":{"a,b":0.75,"other":0.25},"confidence":0.5},"N01":{"type":"noul","noul":0.87},"S01":{"type":"score","score":0.75,"legend":{"0":"low","1":"high"},"probabilities":{"0":0.25,"1":0.75},"confidence":0.4}}}`

func batchFiles(t *testing.T, state, input string) (string, string) {
	t.Helper()
	dir := t.TempDir()
	statePath := filepath.Join(dir, "state.md")
	inputPath := filepath.Join(dir, "questions.jsonl")
	if err := os.WriteFile(statePath, []byte(state), 0600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(inputPath, []byte(input), 0600); err != nil {
		t.Fatal(err)
	}
	return statePath, inputPath
}

func batchArgs(statePath, inputPath string, extra ...string) []string {
	args := []string{"batch", "--state-file", statePath, "--input", inputPath}
	return append(args, extra...)
}

func TestBatchRequestAndOrderedOutput(t *testing.T) {
	state := "shared target\nwith context\n"
	input := strings.Join([]string{
		`{"id":"C01","type":"choice","question":"Choose one","conditions":["a,b","other"]}`,
		`{"id":"N01","type":"noul","question":"Is it true?"}`,
		`{"id":"S01","type":"score","question":"How high?","levels":["low","high"]}`,
	}, "\n") + "\n"
	statePath, inputPath := batchFiles(t, state, input)
	a, out, errOut := harness(t)
	calls := 0
	a.transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		calls++
		data, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		var request evaluationRequest
		if err := json.Unmarshal(data, &request); err != nil {
			t.Fatal(err)
		}
		choice := request.Questions["C01"]
		criteria, ok := choice.Criteria.(map[string]any)
		if request.State != state || request.Model != "chosen" || len(request.Questions) != 3 ||
			choice.Type != primitiveChoice || choice.Instructions != "Choose one" || !ok || len(criteria) != 2 ||
			request.Questions["N01"].Type != primitiveNoul || request.Questions["S01"].Type != primitiveScore {
			t.Fatalf("wrong batch request: %s", data)
		}
		return response(200, batchFixture), nil
	})

	if code := a.run(context.Background(), batchArgs(statePath, inputPath, "--model", "chosen")); code != 0 || calls != 1 || errOut.Len() != 0 {
		t.Fatalf("code=%d calls=%d stdout=%s stderr=%s", code, calls, out, errOut)
	}
	var rows []map[string]json.RawMessage
	scanner := bufio.NewScanner(bytes.NewReader(out.Bytes()))
	for scanner.Scan() {
		var row map[string]json.RawMessage
		if err := json.Unmarshal(scanner.Bytes(), &row); err != nil {
			t.Fatal(err)
		}
		rows = append(rows, row)
	}
	if scanner.Err() != nil || len(rows) != 3 || string(rows[0]["id"]) != `"C01"` ||
		string(rows[1]["id"]) != `"N01"` || string(rows[2]["id"]) != `"S01"` {
		t.Fatalf("unordered or invalid output: %s", out)
	}
	if _, ok := rows[0]["probabilities"]; !ok {
		t.Fatal("choice probabilities missing")
	}
	if _, ok := rows[1]["probabilities"]; ok {
		t.Fatal("noul included choice fields")
	}
	if _, ok := rows[2]["legend"]; !ok {
		t.Fatal("score legend missing")
	}
}

func TestBatchValidatesAllInputBeforeRequest(t *testing.T) {
	valid := `{"id":"A","type":"noul","question":"Valid?"}` + "\n"
	tests := []struct {
		name  string
		state string
		input string
	}{
		{"empty state", " \n", valid},
		{"empty input", "state", "\n"},
		{"malformed JSON", "state", `{`},
		{"unknown field", "state", `{"id":"A","type":"noul","question":"q","extra":true}`},
		{"noncanonical field", "state", `{"ID":"A","TYPE":"noul","QUESTION":"q"}`},
		{"two values", "state", `{"id":"A","type":"noul","question":"q"} {}`},
		{"id whitespace", "state", `{"id":" A ","type":"noul","question":"q"}`},
		{"duplicate id", "state", valid + `{"id":"A","type":"noul","question":"Again?"}`},
		{"unknown type", "state", `{"id":"A","type":"other","question":"q"}`},
		{"choice missing conditions", "state", `{"id":"A","type":"choice","question":"q"}`},
		{"choice null conditions", "state", `{"id":"A","type":"choice","question":"q","conditions":null}`},
		{"choice duplicate conditions", "state", `{"id":"A","type":"choice","question":"q","conditions":["a","a"]}`},
		{"choice condition whitespace", "state", `{"id":"A","type":"choice","question":"q","conditions":[" a","b"]}`},
		{"score with conditions", "state", `{"id":"A","type":"score","question":"q","conditions":["a","b"],"levels":["low","high"]}`},
		{"score level whitespace", "state", `{"id":"A","type":"score","question":"q","levels":["low","high "]}`},
		{"later invalid", "state", valid + `{"id":"B","type":"score","question":"q","levels":["only"]}`},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			statePath, inputPath := batchFiles(t, tc.state, tc.input)
			a, out, errOut := harness(t)
			a.transport = roundTripFunc(func(*http.Request) (*http.Response, error) {
				t.Fatal("invalid batch called API")
				return nil, nil
			})
			if code := a.run(context.Background(), batchArgs(statePath, inputPath)); code != 1 || out.Len() != 0 || errOut.Len() == 0 {
				t.Fatalf("code=%d stdout=%q stderr=%q", code, out, errOut)
			}
		})
	}
}

func TestBatchResponseFailureIsAtomic(t *testing.T) {
	statePath, inputPath := batchFiles(t, "state", strings.Join([]string{
		`{"id":"C01","type":"choice","question":"Choose one","conditions":["a,b","other"]}`,
		`{"id":"N01","type":"noul","question":"Is it true?"}`,
		`{"id":"S01","type":"score","question":"How high?","levels":["low","high"]}`,
	}, "\n"))
	tests := []string{
		strings.Replace(batchFixture, `,"S01":{"type":"score","score":0.75,"legend":{"0":"low","1":"high"},"probabilities":{"0":0.25,"1":0.75},"confidence":0.4}`, "", 1),
		strings.Replace(batchFixture, `"type":"noul"`, `"type":"choice"`, 1),
		strings.Replace(batchFixture, `"model":"jev-batch"`, `"model":""`, 1),
		strings.Replace(batchFixture, `"answers":{`, `"answers":{"extra":{"type":"noul","noul":0.5},`, 1),
	}
	for _, fixture := range tests {
		a, out, errOut := harness(t)
		a.transport = roundTripFunc(func(*http.Request) (*http.Response, error) { return response(200, fixture), nil })
		if code := a.run(context.Background(), batchArgs(statePath, inputPath)); code != 1 || out.Len() != 0 || errOut.Len() == 0 {
			t.Fatalf("code=%d stdout=%q stderr=%q", code, out, errOut)
		}
	}
}

func TestBatchIgnoresSingleEvaluationSettings(t *testing.T) {
	statePath, inputPath := batchFiles(t, "state", `{"id":"N01","type":"noul","question":"Is it true?"}`)
	a, out, errOut := harness(t)
	threshold, minScore, jsonOutput := 1.0, 9.0, true
	pick := []string{"legend"}
	model, timeout := "saved-model", "2s"
	if err := writeConfig(a.path, config{Threshold: &threshold, MinScore: &minScore, JSON: &jsonOutput, Pick: &pick, Model: &model, Timeout: &timeout}); err != nil {
		t.Fatal(err)
	}
	a.transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		var request evaluationRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			t.Fatal(err)
		}
		if request.Model != model {
			t.Fatalf("model=%q", request.Model)
		}
		return response(200, `{"model":"jev-batch","answers":{"N01":{"type":"noul","noul":0.2}}}`), nil
	})
	if code := a.run(context.Background(), batchArgs(statePath, inputPath)); code != 0 || errOut.Len() != 0 || !strings.Contains(out.String(), `"answer":0.2`) {
		t.Fatalf("code=%d stdout=%q stderr=%q", code, out, errOut)
	}
}

func TestBatchAcceptsValidResponseLargerThanRequest(t *testing.T) {
	statePath, inputPath := batchFiles(t, "state", `{"id":"N01","type":"noul","question":"Is it true?"}`)
	a, out, errOut := harness(t)
	a.transport = roundTripFunc(func(*http.Request) (*http.Response, error) {
		body, err := json.Marshal(evaluationEnvelope{
			Model: strings.Repeat("m", 5*1024*1024),
			Answers: map[string]json.RawMessage{
				"N01": json.RawMessage(`{"type":"noul","noul":0.75}`),
			},
		})
		if err != nil {
			t.Fatal(err)
		}
		if len(body) <= 4*1024*1024 {
			t.Fatalf("fixture is only %d bytes", len(body))
		}
		return response(200, string(body)), nil
	})

	if code := a.run(context.Background(), batchArgs(statePath, inputPath)); code != 0 || errOut.Len() != 0 {
		t.Fatalf("code=%d stdout-bytes=%d stderr=%q", code, out.Len(), errOut)
	}
}
