package decision

import (
	"errors"
	"fmt"
)

type Primitive string

const (
	Noul   = "noul"
	Choice = "choice"
	Score  = "score"
)

func (p Primitive) fields() map[string]bool {
	fields := map[string]bool{"answer": true, "model": true}
	if p == Choice || p == Score {
		fields["probabilities"] = true
		fields["confidence"] = true
	}
	if p == Score {
		fields["legend"] = true
	}
	return fields
}

func (p Primitive) ValidatePick(fields []string) error {
	available := p.fields()
	for _, field := range fields {
		if !available[field] {
			return fmt.Errorf("pick field %q is not available for %s", field, p)
		}
	}
	return nil
}

func ValidatePick(fields []string) error {
	seen := make(map[string]bool)
	for _, field := range fields {
		switch field {
		case "answer", "legend", "probabilities", "confidence", "model":
		default:
			return errors.New("pick fields must be answer, legend, probabilities, confidence, or model")
		}
		if seen[field] {
			return errors.New("pick contains a duplicate field")
		}
		seen[field] = true
	}
	return nil
}
