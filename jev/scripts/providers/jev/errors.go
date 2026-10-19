package jev

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"
)

const maxErrorBodyBytes = 64 * 1024

// Interpret only bounded, recognized fields. Never return server messages or input values.
func evaluationHTTPError(response *http.Response, apiKey string) error {
	message := fmt.Sprintf("API returned HTTP %d (%s)", response.StatusCode, http.StatusText(response.StatusCode))
	data, err := io.ReadAll(io.LimitReader(response.Body, maxErrorBodyBytes+1))
	if err == nil && len(data) <= maxErrorBodyBytes && utf8.Valid(data) {
		var body map[string]json.RawMessage
		if json.Unmarshal(data, &body) == nil && body != nil {
			detail := ""
			switch response.StatusCode {
			case http.StatusBadRequest:
				if contextLimitCode(body) {
					detail = "context limit exceeded; reduce state or questions, or explicitly use --long-context if appropriate"
				}
			case http.StatusRequestEntityTooLarge:
				detail = "request body exceeds 32 MiB; reduce state or questions"
			case http.StatusUnprocessableEntity:
				detail = validationDetails(body["detail"])
			case http.StatusTooManyRequests:
				detail = "rate limit exceeded; try again later (no automatic retry)"
			case 529:
				detail = "TypeSafe temporarily overloaded; try again later (no automatic retry)"
			}
			if detail != "" {
				message += "; " + detail
			}
		}
	}
	for _, token := range []string{apiKey, strings.TrimSpace(apiKey)} {
		if token != "" {
			message = strings.ReplaceAll(message, token, "[redacted]")
		}
	}
	return errors.New(message)
}

func contextLimitCode(body map[string]json.RawMessage) bool {
	for _, field := range []string{"code", "type"} {
		if exactContextCode(body[field]) {
			return true
		}
	}
	for _, field := range []string{"detail", "error"} {
		raw := body[field]
		if exactContextCode(raw) {
			return true
		}
		var nested map[string]json.RawMessage
		if json.Unmarshal(raw, &nested) == nil {
			for _, key := range []string{"code", "type"} {
				if exactContextCode(nested[key]) {
					return true
				}
			}
		}
	}
	return false
}

func exactContextCode(raw json.RawMessage) bool {
	var value string
	return json.Unmarshal(raw, &value) == nil && value == "max_tokens_exceeded"
}

func validationDetails(raw json.RawMessage) string {
	var details []struct {
		Location []json.RawMessage `json:"loc"`
		Type     string            `json:"type"`
	}
	if json.Unmarshal(raw, &details) != nil {
		return ""
	}
	descriptions := map[string]string{
		"missing":             "required field missing",
		"string_type":         "expected a string",
		"list_type":           "expected an array",
		"dict_type":           "expected an object",
		"model_type":          "expected an object",
		"float_type":          "expected a number",
		"float_parsing":       "expected a number",
		"int_type":            "expected an integer",
		"int_parsing":         "expected an integer",
		"literal_error":       "unsupported value",
		"union_tag_invalid":   "unsupported question type",
		"union_tag_not_found": "question type missing",
		"too_short":           "too few entries",
		"too_long":            "too many entries",
		"extra_forbidden":     "unexpected field",
		"finite_number":       "expected a finite number",
	}
	messages := make([]string, 0, 5)
	for i, detail := range details {
		if i == 5 {
			break
		}
		description := descriptions[detail.Type]
		path := validationPath(detail.Location)
		if description != "" && path != "" {
			messages = append(messages, path+": "+description)
		}
	}
	return strings.Join(messages, "; ")
}

func validationPath(location []json.RawMessage) string {
	if len(location) < 2 || len(location) > 16 {
		return ""
	}
	var root string
	if json.Unmarshal(location[0], &root) != nil || root != "body" {
		return ""
	}
	var field string
	if json.Unmarshal(location[1], &field) != nil {
		return ""
	}
	switch field {
	case "state", "model", "questions":
	default:
		return ""
	}
	path := "body." + field
	discriminator := false
	for i, raw := range location[2:] {
		var name string
		if len(raw) > 0 && raw[0] == '"' && json.Unmarshal(raw, &name) == nil {
			switch {
			case field == "questions" && i == 0:
				path += ".<question>"
			case field == "questions" && i == 1 && (name == "noul" || name == "choice" || name == "score"):
				path += "." + name
				discriminator = true
			case field == "questions" && (i == 1 || (i == 2 && discriminator)):
				switch name {
				case "type", "instructions", "criteria":
					path += "." + name
				default:
					path += ".<field>"
				}
			default:
				path += ".<field>"
			}
			continue
		}
		var index *int
		if json.Unmarshal(raw, &index) != nil || index == nil || *index < 0 || *index > 1000000 {
			return ""
		}
		path += "[" + strconv.Itoa(*index) + "]"
	}
	return path
}
