package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type result struct {
	Primitive     primitive          `json:"-"`
	Answer        any                `json:"answer"`
	Legend        map[string]string  `json:"legend,omitempty"`
	Probabilities map[string]float64 `json:"probabilities,omitempty"`
	Confidence    *float64           `json:"confidence,omitempty"`
	Model         string             `json:"model"`
}

func validatePick(fields []string) error {
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

func writeResult(out io.Writer, r result, s settings) error {
	if !s.JSON {
		_, err := fmt.Fprintln(out, r.Answer)
		return err
	}
	var value any = r
	if len(s.Pick) > 0 {
		all := r.fields()
		selected := make(map[string]any, len(s.Pick))
		for _, key := range s.Pick {
			selected[key] = all[key]
		}
		value = selected
	}
	data, err := json.Marshal(value)
	if err != nil {
		return errors.New("could not encode result")
	}
	_, err = out.Write(append(data, '\n'))
	return err
}

func (r result) fields() map[string]any {
	fields := map[string]any{"answer": r.Answer, "model": r.Model}
	if r.Primitive == primitiveChoice || r.Primitive == primitiveScore {
		fields["probabilities"] = r.Probabilities
		fields["confidence"] = *r.Confidence
	}
	if r.Primitive == primitiveScore {
		fields["legend"] = r.Legend
	}
	return fields
}
