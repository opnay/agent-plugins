package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type result struct {
	Answer        string             `json:"answer"`
	Probabilities map[string]float64 `json:"probabilities"`
	Confidence    float64            `json:"confidence"`
	Model         string             `json:"model"`
}

func validatePick(fields []string) error {
	seen := make(map[string]bool)
	for _, field := range fields {
		switch field {
		case "answer", "probabilities", "confidence", "model":
		default:
			return errors.New("pick fields must be answer, probabilities, confidence, or model")
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
		all := map[string]any{"answer": r.Answer, "probabilities": r.Probabilities,
			"confidence": r.Confidence, "model": r.Model}
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
