package providers

import (
	"context"

	"opnay/jev/decision"
)

// Adapter owns provider I/O and returns validated decisions, never wire responses.
// Evaluate returns exactly one result per question ID, or an error and no results.
// Authentication checks never infer, retry, or expose credentials or response bodies.
type Adapter interface {
	Evaluate(context.Context, decision.Request) (map[string]decision.Result, error)
	CheckAuthentication(context.Context) (Authentication, error)
}

type Authentication struct {
	Detail      string
	Hint        string
	KeyRejected bool
}
