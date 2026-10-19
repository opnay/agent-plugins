package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"os"
	"reflect"
	"sort"
	"strings"
	"unicode/utf8"
)

// UseNumber preserves numeric literals when structured input is sent to the API.
func decodeJSON(data []byte) (any, error) {
	if !utf8.Valid(data) {
		return nil, errors.New("JSON must be valid UTF-8")
	}
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	var value any
	if err := decoder.Decode(&value); err != nil {
		return nil, errors.New("invalid JSON")
	}
	if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return nil, errors.New("expected exactly one JSON value")
	}
	return value, nil
}

func validateEntry(value any, nullable bool) error {
	switch v := value.(type) {
	case nil:
		if !nullable {
			return errors.New("expected a string, object, or array")
		}
	case string:
		if strings.TrimSpace(v) == "" {
			return errors.New("text must not be empty")
		}
	case map[string]any, []any:
	default:
		return errors.New("expected a string, object, or array")
	}
	return validateJSONContent(value, 0)
}

func validateJSONContent(value any, depth int) error {
	switch v := value.(type) {
	case map[string]any:
		if depth >= 64 {
			return errors.New("JSON nesting must not exceed 64 levels")
		}
		for _, child := range v {
			if err := validateJSONContent(child, depth+1); err != nil {
				return err
			}
		}
	case []any:
		if depth >= 64 {
			return errors.New("JSON nesting must not exceed 64 levels")
		}
		for _, child := range v {
			if err := validateJSONContent(child, depth+1); err != nil {
				return err
			}
		}
	case json.Number:
		f, err := v.Float64()
		if err != nil || !finite(f) {
			return errors.New("JSON numbers must be finite and within float64 range")
		}
	}
	return nil
}

func (o *evaluationOptions) readInput() error {
	data, err := os.ReadFile(o.inputPath)
	if err != nil {
		return fmt.Errorf("read input file: %w", err)
	}
	value, err := decodeJSON(data)
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
	if err := validateEntry(state, false); err != nil {
		return fmt.Errorf("state: %w", err)
	}
	if err := o.setStructuredQuestion(fields["instructions"], fields["criteria"]); err != nil {
		return err
	}
	o.state = state
	return nil
}

func (o *evaluationOptions) setStructuredQuestion(instructions, criteria any) error {
	if err := validateEntry(instructions, true); err != nil {
		return fmt.Errorf("instructions: %w", err)
	}
	if err := validateJSONContent(criteria, 0); err != nil {
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
				if err := validateEntry(value, true); err != nil {
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
			if err := validateEntry(value, true); err != nil {
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
			if err := validateEntry(value, false); err != nil {
				return fmt.Errorf("Score level: %w", err)
			}
			for _, previous := range values[:i] {
				if jsonValuesEqual(previous, value) {
					return errors.New("Score criteria contains a duplicate level")
				}
			}
		}
		o.levels = values
	default:
		return errors.New("unknown primitive")
	}
	o.payload = &evaluationQuestion{Type: o.primitive, Instructions: instructions, Criteria: criteria}
	return nil
}

// Compare structured legends by JSON value, independent of object order or number notation.
func jsonValuesEqual(a, b any) bool {
	switch x := a.(type) {
	case json.Number:
		y, ok := b.(json.Number)
		if !ok {
			return false
		}
		left, leftOK := new(big.Rat).SetString(string(x))
		right, rightOK := new(big.Rat).SetString(string(y))
		return leftOK && rightOK && left.Cmp(right) == 0
	case map[string]any:
		y, ok := b.(map[string]any)
		if !ok || len(x) != len(y) {
			return false
		}
		for key, value := range x {
			other, exists := y[key]
			if !exists || !jsonValuesEqual(value, other) {
				return false
			}
		}
		return true
	case []any:
		y, ok := b.([]any)
		if !ok || len(x) != len(y) {
			return false
		}
		for i, value := range x {
			if !jsonValuesEqual(value, y[i]) {
				return false
			}
		}
		return true
	default:
		return reflect.DeepEqual(a, b)
	}
}
