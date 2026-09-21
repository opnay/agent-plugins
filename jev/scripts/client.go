package main

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
)

const endpoint = "https://api.typesafe.ai/v1/systemone"

type choiceQuestion struct {
	Type         string         `json:"type"`
	Instructions string         `json:"instructions"`
	Criteria     map[string]any `json:"criteria"`
}

type evaluationRequest struct {
	State     string                    `json:"state"`
	Model     string                    `json:"model"`
	Questions map[string]choiceQuestion `json:"questions"`
}

type choiceAnswer struct {
	Type          string              `json:"type"`
	Choice        string              `json:"choice"`
	Probabilities map[string]*float64 `json:"probabilities"`
	Confidence    *float64            `json:"confidence"`
}

func evaluate(ctx context.Context, o evaluationOptions, s settings, transport http.RoundTripper) (result, error) {
	criteria := make(map[string]any, len(o.choices))
	for _, choice := range o.choices {
		criteria[choice] = nil
	}
	payload := evaluationRequest{State: o.question, Model: s.Model,
		Questions: map[string]choiceQuestion{"result": {
			Type: "choice", Instructions: "Select the option that best answers the question in state.", Criteria: criteria,
		}}}
	body, err := json.Marshal(payload)
	if err != nil {
		return result{}, errors.New("could not encode API request")
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(body))
	if err != nil {
		return result{}, errors.New("could not create API request")
	}
	req.Header.Set("Authorization", "Bearer "+s.APIKey)
	req.Header.Set("Content-Type", "application/json")
	timeout, _ := time.ParseDuration(s.Timeout) // Validated before network access.
	client := &http.Client{Transport: transport, Timeout: timeout,
		CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }}
	response, err := client.Do(req)
	if err != nil {
		if ctx.Err() != nil {
			return result{}, errors.New("API request canceled")
		}
		if errors.Is(err, context.DeadlineExceeded) {
			return result{}, errors.New("API request timed out")
		}
		return result{}, errors.New("API request failed; check network access and authentication configuration")
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return result{}, fmt.Errorf("API returned HTTP %d (%s)", response.StatusCode, http.StatusText(response.StatusCode))
	}
	data, err := io.ReadAll(io.LimitReader(response.Body, 4*1024*1024+1))
	if err != nil {
		return result{}, errors.New("could not read API response (connection interrupted or request timed out)")
	}
	if len(data) > 4*1024*1024 {
		return result{}, errors.New("API response exceeds 4 MiB")
	}
	var envelope struct {
		Model   string                  `json:"model"`
		Answers map[string]choiceAnswer `json:"answers"`
	}
	if err := json.Unmarshal(data, &envelope); err != nil {
		return result{}, errors.New("invalid JSON in API response")
	}
	answer, ok := envelope.Answers["result"]
	if !ok || strings.TrimSpace(envelope.Model) == "" || answer.Type != "choice" ||
		answer.Confidence == nil || !probability(*answer.Confidence) {
		return result{}, errors.New("API response is missing valid model or Choice fields")
	}
	probabilities := make(map[string]float64, len(o.choices))
	if len(answer.Probabilities) != len(o.choices) {
		return result{}, errors.New("API probabilities do not match requested choices")
	}
	for _, choice := range o.choices {
		p := answer.Probabilities[choice]
		if p == nil || !probability(*p) {
			return result{}, errors.New("API response contains a missing or invalid probability")
		}
		probabilities[choice] = *p
	}
	selected, ok := probabilities[answer.Choice]
	if !ok {
		return result{}, errors.New("API selected an unknown choice")
	}
	for _, p := range probabilities {
		if p > selected {
			return result{}, errors.New("API choice is not a highest-probability option")
		}
	}
	return result{Answer: answer.Choice, Probabilities: probabilities,
		Confidence: *answer.Confidence, Model: envelope.Model}, nil
}
