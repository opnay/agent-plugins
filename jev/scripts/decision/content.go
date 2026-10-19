package decision

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"math"
	"math/big"
	"reflect"
	"strings"
	"unicode/utf8"
)

// UseNumber preserves numeric literals when structured input is sent to the API.
func DecodeJSON(data []byte) (any, error) {
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

func ValidateEntry(value any, nullable bool) error {
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
	return ValidateJSONContent(value, 0)
}

func ValidateJSONContent(value any, depth int) error {
	switch v := value.(type) {
	case map[string]any:
		if depth >= 64 {
			return errors.New("JSON nesting must not exceed 64 levels")
		}
		for _, child := range v {
			if err := ValidateJSONContent(child, depth+1); err != nil {
				return err
			}
		}
	case []any:
		if depth >= 64 {
			return errors.New("JSON nesting must not exceed 64 levels")
		}
		for _, child := range v {
			if err := ValidateJSONContent(child, depth+1); err != nil {
				return err
			}
		}
	case json.Number:
		f, err := v.Float64()
		if err != nil || (math.IsNaN(f) || math.IsInf(f, 0)) {
			return errors.New("JSON numbers must be finite and within float64 range")
		}
	}
	return nil
}

// Compare structured legends by JSON value, independent of object order or number notation.
func JSONValuesEqual(a, b any) bool {
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
			if !exists || !JSONValuesEqual(value, other) {
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
			if !JSONValuesEqual(value, y[i]) {
				return false
			}
		}
		return true
	default:
		return reflect.DeepEqual(a, b)
	}
}
