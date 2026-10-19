package providers

import "context"

// Adapter owns provider I/O and returns validated decisions, never wire responses.
// Evaluate returns exactly one result per question ID, or an error and no results.
// Authentication checks never infer, retry, or expose credentials or response bodies.
type Adapter interface {
	Evaluate(context.Context, Request) (map[string]Result, error)
	CheckAuthentication(context.Context) (Authentication, error)
}

type Authentication struct {
	Detail      string
	Hint        string
	KeyRejected bool
}

// Request contains decisions over one state; CLI settings and file I/O stay with the caller.
type Request struct {
	State       any
	Model       string
	Questions   map[string]Question
	LongContext bool
}

type Question struct {
	Type         string `json:"type"`
	Instructions any    `json:"instructions"`
	Criteria     any    `json:"criteria,omitempty"`
}

type Result struct {
	Primitive     string             `json:"-"`
	Answer        any                `json:"answer"`
	Legend        map[string]any     `json:"legend,omitempty"`
	Probabilities map[string]float64 `json:"probabilities,omitempty"`
	Confidence    *float64           `json:"confidence,omitempty"`
	Model         string             `json:"model"`
}
