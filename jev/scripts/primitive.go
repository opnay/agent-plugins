package main

import "fmt"

type primitive string

const (
	primitiveNoul   primitive = "noul"
	primitiveChoice primitive = "choice"
	primitiveScore  primitive = "score"
)

func (p primitive) fields() map[string]bool {
	fields := map[string]bool{"answer": true, "model": true}
	if p == primitiveChoice || p == primitiveScore {
		fields["probabilities"] = true
		fields["confidence"] = true
	}
	if p == primitiveScore {
		fields["legend"] = true
	}
	return fields
}

func (p primitive) validatePick(fields []string) error {
	available := p.fields()
	for _, field := range fields {
		if !available[field] {
			return fmt.Errorf("pick field %q is not available for %s", field, p)
		}
	}
	return nil
}
