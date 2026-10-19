package jev

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"opnay/jev/decision"
	"opnay/jev/providers"
)

const endpoint = "https://api.typesafe.ai/v1/systemone"

type Adapter struct {
	apiKey    string
	timeout   time.Duration
	transport http.RoundTripper
}

var _ providers.Adapter = (*Adapter)(nil)

func New(apiKey string, timeout time.Duration, transport http.RoundTripper) *Adapter {
	return &Adapter{apiKey: apiKey, timeout: timeout, transport: transport}
}

type evaluationRequest struct {
	State       any                          `json:"state"`
	Model       string                       `json:"model"`
	Questions   map[string]decision.Question `json:"questions"`
	LongContext bool                         `json:"-"`
}

type evaluationEnvelope struct {
	Model   string                     `json:"model"`
	Answers map[string]json.RawMessage `json:"answers"`
}

func (a *Adapter) Evaluate(ctx context.Context, request decision.Request) (map[string]decision.Result, error) {
	payload := evaluationRequest{
		State: request.State, Model: request.Model,
		Questions: request.Questions, LongContext: request.LongContext,
	}
	envelope, err := a.requestEvaluation(ctx, payload)
	if err != nil {
		return nil, err
	}
	if strings.TrimSpace(envelope.Model) == "" || len(envelope.Answers) != len(request.Questions) {
		return nil, errors.New("API response is missing a valid model or complete answers")
	}
	results := make(map[string]decision.Result, len(request.Questions))
	for id, question := range request.Questions {
		raw, ok := envelope.Answers[id]
		if !ok {
			return nil, fmt.Errorf("API response is missing answer %q", id)
		}
		answer, err := parseAnswer(raw, envelope.Model, question)
		if err != nil {
			return nil, fmt.Errorf("API response for %q: %w", id, err)
		}
		results[id] = answer
	}
	return results, nil
}

func (a *Adapter) requestEvaluation(ctx context.Context, payload evaluationRequest) (evaluationEnvelope, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return evaluationEnvelope{}, errors.New("could not encode API request")
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(body))
	if err != nil {
		return evaluationEnvelope{}, errors.New("could not create API request")
	}
	req.Header.Set("Authorization", "Bearer "+a.apiKey)
	req.Header.Set("Content-Type", "application/json")
	if payload.LongContext {
		req.Header.Set("TypeSafe-Danger-Zone-Long-Context", "true")
	}
	client := &http.Client{Transport: a.transport, Timeout: a.timeout,
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
