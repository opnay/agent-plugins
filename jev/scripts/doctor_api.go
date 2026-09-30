package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const modelsEndpoint = "https://api.typesafe.ai/v1/models"

func checkDoctorAPI(ctx context.Context, s settings, transport http.RoundTripper, keyHint string) (string, string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, modelsEndpoint, nil)
	if err != nil {
		return "", "", errors.New("could not create authentication check")
	}
	req.Header.Set("Authorization", "Bearer "+s.APIKey)
	req.Header.Set("Accept", "application/json")
	timeout, _ := time.ParseDuration(s.Timeout)
	client := &http.Client{Transport: transport, Timeout: timeout,
		CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }}
	response, err := client.Do(req)
	if err != nil {
		if ctx.Err() != nil {
			return "", "run doctor again when ready", errors.New("authentication check canceled; key validity unknown")
		}
		if errors.Is(err, context.DeadlineExceeded) {
			return "", "check connectivity or increase jev config timeout", errors.New("authentication check timed out; key validity unknown")
		}
		return "", "check network access, DNS, proxy, and TLS settings", errors.New("cannot reach TypeSafe API; key validity unknown")
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		return "authenticated (GET " + modelsEndpoint + ")", "", nil
	}
	detail := fmt.Sprintf("HTTP %d (%s)", response.StatusCode, http.StatusText(response.StatusCode))
	hint := "check TypeSafe service status; authentication was not confirmed"
	switch response.StatusCode {
	case http.StatusUnauthorized:
		detail += "; API key rejected (invalid or expired)"
		hint = keyHint
	case http.StatusForbidden:
		detail += "; access denied"
		hint = "check API key permissions and account access in TypeSafe"
	case http.StatusTooManyRequests:
		detail += "; rate limited; key validity unknown"
		hint = "wait for the rate limit to clear, then run doctor again"
	default:
		detail += "; key validity unknown"
	}
	// Never expose the response body or transport error: either can echo secrets.
	return "", hint, errors.New(detail)
}
