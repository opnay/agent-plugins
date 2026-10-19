package jev

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
	"testing/iotest"
	"time"

	"opnay/jev/decision"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func TestSerializedRequestSizeLimit(t *testing.T) {
	const limit = 32 * 1024 * 1024
	// Independent wire fixture: only the state's string contents vary.
	const empty = `{"state":"","model":"jev-latest","questions":{"result":{"type":"noul","instructions":"True?"}}}`
	for _, tc := range []struct {
		name, state string
		allowed     bool
	}{
		{"exact boundary", strings.Repeat("a", limit-len(empty)), true},
		{"one byte over", strings.Repeat("a", limit-len(empty)+1), false},
		{"JSON escaping exceeds limit", strings.Repeat("<", limit/6+1), false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			calls := 0
			adapter := New("test-token", time.Second, roundTripFunc(func(r *http.Request) (*http.Response, error) {
				calls++
				if r.ContentLength != limit {
					t.Fatalf("wire size=%d want=%d", r.ContentLength, limit)
				}
				return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{"model":"jev-latest","answers":{"result":{"type":"noul","noul":0.8}}}`)), Header: make(http.Header)}, nil
			}))
			result, err := adapter.Evaluate(context.Background(), decision.Request{State: tc.state, Model: "jev-latest", Questions: map[string]decision.Question{"result": {Type: "noul", Instructions: "True?"}}})
			if tc.allowed {
				if err != nil || calls != 1 || result["result"].Answer != 0.8 {
					t.Fatalf("boundary rejected: calls=%d err=%v", calls, err)
				}
			} else if err == nil || calls != 0 || result != nil || !strings.Contains(err.Error(), "32 MiB") {
				t.Fatalf("oversize accepted or sent: calls=%d err=%v", calls, err)
			}
		})
	}
}

type observedBody struct {
	reader io.Reader
	bytes  int
	closed bool
}

func (b *observedBody) Read(p []byte) (int, error) {
	n, err := b.reader.Read(p)
	b.bytes += n
	return n, err
}
func (b *observedBody) Close() error { b.closed = true; return nil }

func TestFailedResponseBodyLifecycle(t *testing.T) {
	const missing = `{"detail":[{"loc":["body","state"],"type":"missing"}]}`
	for _, tc := range []struct {
		name, key string
		reader    io.Reader
		want      string
		maxRead   int
	}{
		{"bounded body", "test-token", strings.NewReader(missing + strings.Repeat(" ", 128*1024)), "API returned HTTP 422 (Unprocessable Entity)", 65537},
		{"read failure", "test-token", io.MultiReader(strings.NewReader(missing), iotest.ErrReader(errors.New("test-token"))), "API returned HTTP 422 (Unprocessable Entity)", len(missing)},
		{"key matches schema word", "state", strings.NewReader(missing), "API returned HTTP 422 (Unprocessable Entity); body.[redacted]: required field missing", len(missing)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			body := &observedBody{reader: tc.reader}
			calls := 0
			adapter := New(tc.key, time.Second, roundTripFunc(func(*http.Request) (*http.Response, error) {
				calls++
				return &http.Response{StatusCode: 422, Body: body, Header: make(http.Header)}, nil
			}))
			results, err := adapter.Evaluate(context.Background(), decision.Request{State: "state", Model: "jev-latest", Questions: map[string]decision.Question{"result": {Type: "noul", Instructions: "True?"}}})
			if err == nil || err.Error() != tc.want || results != nil || calls != 1 || !body.closed || body.bytes > tc.maxRead {
				t.Fatalf("calls=%d closed=%v bytes=%d err=%v", calls, body.closed, body.bytes, err)
			}
		})
	}
}
