package jev

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"opnay/jev/providers"
)

const modelsEndpoint = "https://api.typesafe.ai/v1/models"

func (a *Adapter) CheckAuthentication(ctx context.Context) (providers.Authentication, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, modelsEndpoint, nil)
	if err != nil {
		return providers.Authentication{}, errors.New("could not create authentication check")
	}
	req.Header.Set("Authorization", "Bearer "+a.apiKey)
	req.Header.Set("Accept", "application/json")
	client := &http.Client{Transport: a.transport, Timeout: a.timeout,
		CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }}
	response, err := client.Do(req)
	if err != nil {
		if ctx.Err() != nil {
			return providers.Authentication{Hint: "run doctor again when ready"}, errors.New("authentication check canceled; key validity unknown")
		}
		if errors.Is(err, context.DeadlineExceeded) {
			return providers.Authentication{Hint: "check connectivity or increase request timeout"}, errors.New("authentication check timed out; key validity unknown")
		}
		return providers.Authentication{Hint: "check network access, DNS, proxy, and TLS settings"}, errors.New("cannot reach TypeSafe API; key validity unknown")
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		return providers.Authentication{Detail: "authenticated (GET " + modelsEndpoint + ")"}, nil
	}
	detail := fmt.Sprintf("HTTP %d (%s)", response.StatusCode, http.StatusText(response.StatusCode))
	hint := "check TypeSafe service status; authentication was not confirmed"
	switch response.StatusCode {
	case http.StatusUnauthorized:
		detail += "; API key rejected (invalid or expired)"
		hint = "replace the rejected API key"
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
	return providers.Authentication{Hint: hint, KeyRejected: response.StatusCode == http.StatusUnauthorized}, errors.New(detail)
}
