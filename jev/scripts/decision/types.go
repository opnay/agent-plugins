package decision

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
