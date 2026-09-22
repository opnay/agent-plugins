package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

const endpoint = "https://api.typesafe.ai/v1/systemone"

type evaluationQuestion struct {
	Type         primitive `json:"type"`
	Instructions string    `json:"instructions"`
	Criteria     any       `json:"criteria,omitempty"`
}

type evaluationRequest struct {
	State     string                        `json:"state"`
	Model     string                        `json:"model"`
	Questions map[string]evaluationQuestion `json:"questions"`
}

type evaluationEnvelope struct {
	Model   string                     `json:"model"`
	Answers map[string]json.RawMessage `json:"answers"`
}

func evaluate(ctx context.Context, o evaluationOptions, s settings, transport http.RoundTripper) (result, error) {
	payload := evaluationRequest{State: o.question, Model: s.Model,
		Questions: map[string]evaluationQuestion{"result": o.questionPayload()}}
	envelope, err := requestEvaluation(ctx, payload, s, transport)
	if err != nil {
		return result{}, err
	}
	return parseResponse(envelope, o)
}

func requestEvaluation(ctx context.Context, payload evaluationRequest, s settings, transport http.RoundTripper) (evaluationEnvelope, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return evaluationEnvelope{}, errors.New("could not encode API request")
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(body))
	if err != nil {
		return evaluationEnvelope{}, errors.New("could not create API request")
	}
	req.Header.Set("Authorization", "Bearer "+s.APIKey)
	req.Header.Set("Content-Type", "application/json")
	timeout, _ := time.ParseDuration(s.Timeout)
	client := &http.Client{Transport: transport, Timeout: timeout,
		CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }}
	response, err := client.Do(req)
	if err != nil {
		if ctx.Err() != nil {
			return evaluationEnvelope{}, errors.New("API request canceled")
		}
		if errors.Is(err, context.DeadlineExceeded) {
			return evaluationEnvelope{}, errors.New("API request timed out")
		}
		return evaluationEnvelope{}, errors.New("API request failed; check network access and authentication configuration")
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return evaluationEnvelope{}, fmt.Errorf("API returned HTTP %d (%s)", response.StatusCode, http.StatusText(response.StatusCode))
	}
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return evaluationEnvelope{}, errors.New("could not read API response (connection interrupted or request timed out)")
	}
	var envelope evaluationEnvelope
	if err := json.Unmarshal(data, &envelope); err != nil {
		return evaluationEnvelope{}, errors.New("invalid JSON in API response")
	}
	return envelope, nil
}

func (o evaluationOptions) questionPayload() evaluationQuestion {
	switch o.primitive {
	case primitiveNoul:
		return evaluationQuestion{Type: o.primitive, Instructions: "Answer whether the proposition or question in state is true."}
	case primitiveScore:
		return evaluationQuestion{Type: o.primitive, Instructions: "Rate the state against the ordered criteria.", Criteria: o.levels}
	default:
		criteria := make(map[string]any, len(o.choices))
		for _, choice := range o.choices {
			criteria[choice] = nil
		}
		return evaluationQuestion{Type: o.primitive, Instructions: "Select the option that best answers the question in state.", Criteria: criteria}
	}
}
