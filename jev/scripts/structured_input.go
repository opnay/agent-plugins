package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"

	"opnay/jev/providers"
)

func (o *evaluationOptions) readInput() error {
	data, err := os.ReadFile(o.inputPath)
	if err != nil {
		return fmt.Errorf("read input file: %w", err)
	}
	value, err := providers.DecodeJSON(data)
	if err != nil {
		return err
	}
	fields, ok := value.(map[string]any)
	if !ok {
		return errors.New("input must be a JSON object")
	}
	for key := range fields {
		if key != "state" && key != "instructions" && key != "criteria" {
			return errors.New("input accepts only state, instructions, and criteria")
		}
	}
	state, exists := fields["state"]
	if !exists {
		return errors.New("input requires state")
	}
	if err := providers.ValidateEntry(state, false); err != nil {
		return fmt.Errorf("state: %w", err)
	}
	if err := o.setStructuredQuestion(fields["instructions"], fields["criteria"]); err != nil {
		return err
	}
	o.state = state
	return nil
}

func (o *evaluationOptions) setStructuredQuestion(instructions, criteria any) error {
	if err := providers.ValidateEntry(instructions, true); err != nil {
		return fmt.Errorf("instructions: %w", err)
	}
	if err := providers.ValidateJSONContent(criteria, 0); err != nil {
		return fmt.Errorf("criteria: %w", err)
	}
	switch o.primitive {
	case primitiveNoul:
		if criteria != nil {
			values, ok := criteria.(map[string]any)
			if !ok {
				return errors.New("Noul criteria must be an object or null")
			}
			for key, value := range values {
				if key != "true" && key != "false" {
					return errors.New("Noul criteria accepts only true and false")
				}
				if err := providers.ValidateEntry(value, true); err != nil {
					return fmt.Errorf("Noul criteria description: %w", err)
				}
			}
		}
	case primitiveChoice:
		values, ok := criteria.(map[string]any)
		if !ok || len(values) < 2 || len(values) > 255 {
			return errors.New("Choice criteria requires an object with 2 to 255 labels")
		}
		for label, value := range values {
			if label == "" || label != strings.TrimSpace(label) {
				return errors.New("Choice labels must be nonempty without surrounding whitespace")
			}
			if err := providers.ValidateEntry(value, true); err != nil {
				return fmt.Errorf("Choice criteria description: %w", err)
			}
			o.choices = append(o.choices, label)
		}
		sort.Strings(o.choices)
	case primitiveScore:
		values, ok := criteria.([]any)
		if !ok || len(values) < 2 || len(values) > 10 {
			return errors.New("Score criteria requires 2 to 10 ordered levels")
		}
		for i, value := range values {
			if err := providers.ValidateEntry(value, false); err != nil {
				return fmt.Errorf("Score level: %w", err)
			}
			for _, previous := range values[:i] {
				if providers.JSONValuesEqual(previous, value) {
					return errors.New("Score criteria contains a duplicate level")
				}
			}
		}
		o.levels = values
	default:
		return errors.New("unknown primitive")
	}
	o.payload = &providers.Question{Type: string(o.primitive), Instructions: instructions, Criteria: criteria}
	return nil
}
