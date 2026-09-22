package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"unicode/utf8"
)

type batchOptions struct {
	statePath string
	inputPath string
	flags     config
}

type batchInputLine struct {
	ID         string          `json:"id"`
	Type       primitive       `json:"type"`
	Question   string          `json:"question"`
	Conditions json.RawMessage `json:"conditions"`
	Levels     json.RawMessage `json:"levels"`
}

type batchItem struct {
	id         string
	primitive  primitive
	question   string
	conditions []string
	levels     []string
}

type batchOutput struct {
	ID            string             `json:"id"`
	Type          primitive          `json:"type"`
	Answer        any                `json:"answer"`
	Legend        map[string]string  `json:"legend,omitempty"`
	Probabilities map[string]float64 `json:"probabilities,omitempty"`
	Confidence    *float64           `json:"confidence,omitempty"`
	Model         string             `json:"model"`
}

func parseBatch(args []string) (batchOptions, error) {
	var o batchOptions
	var model, timeout string
	f := flag.NewFlagSet("jev batch", flag.ContinueOnError)
	f.SetOutput(io.Discard)
	f.StringVar(&o.statePath, "state-file", "", "shared state file")
	f.StringVar(&o.inputPath, "input", "", "JSONL question file")
	f.StringVar(&model, "model", "jev-latest", "model")
	f.StringVar(&timeout, "timeout", "30s", "timeout")
	if err := f.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return o, flag.ErrHelp
		}
		return o, errors.New("invalid options; see jev batch --help")
	}
	if f.NArg() != 0 {
		return o, errors.New("unexpected positional arguments; see jev batch --help")
	}
	if strings.TrimSpace(o.statePath) == "" {
		return o, errors.New("--state-file requires a path; see jev batch --help")
	}
	if strings.TrimSpace(o.inputPath) == "" {
		return o, errors.New("--input requires a path; see jev batch --help")
	}
	f.Visit(func(v *flag.Flag) {
		switch v.Name {
		case "model":
			o.flags.Model = &model
		case "timeout":
			o.flags.Timeout = &timeout
		}
	})
	if err := o.flags.resolve("").validate(); err != nil {
		return o, fmt.Errorf("%w; see jev batch --help", err)
	}
	return o, nil
}

func (a application) executeBatch(ctx context.Context, o batchOptions) error {
	path, err := a.configPath()
	if err != nil {
		return err
	}
	c, err := readConfig(path)
	if err != nil {
		return err
	}
	c.override(o.flags)
	s := c.resolve(a.envToken)
	s.Threshold = nil
	s.MinScore = nil
	s.JSON = false
	s.Pick = nil
	if err := s.validate(); err != nil {
		return fmt.Errorf("%w; see jev batch --help", err)
	}
	if s.APIKey == "" {
		return errors.New("API key missing; set TYPESAFE_API_KEY or use jev config set api_key --stdin")
	}
	state, err := readBatchState(o.statePath)
	if err != nil {
		return err
	}
	items, err := readBatchItems(o.inputPath)
	if err != nil {
		return err
	}
	results, err := evaluateBatch(ctx, state, items, s, a.transport)
	if err != nil {
		return err
	}
	return writeBatchResults(a.out, items, results)
}

func readBatchState(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read state file: %w", err)
	}
	if !utf8.Valid(data) {
		return "", errors.New("state file must be valid UTF-8")
	}
	state := string(data)
	if strings.TrimSpace(state) == "" {
		return "", errors.New("state file must not be empty")
	}
	return state, nil
}

func readBatchItems(path string) ([]batchItem, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open batch input: %w", err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	items := make([]batchItem, 0)
	seen := make(map[string]bool)
	for line := 1; ; line++ {
		data, readErr := reader.ReadBytes('\n')
		if len(data) > 0 {
			if !utf8.Valid(data) {
				return nil, fmt.Errorf("batch input line %d must be valid UTF-8", line)
			}
			if len(bytes.TrimSpace(data)) > 0 {
				item, err := decodeBatchItem(data)
				if err != nil {
					return nil, fmt.Errorf("batch input line %d: %w", line, err)
				}
				if seen[item.id] {
					return nil, fmt.Errorf("batch input line %d: duplicate id %q", line, item.id)
				}
				seen[item.id] = true
				items = append(items, item)
			}
		}
		if errors.Is(readErr, io.EOF) {
			break
		}
		if readErr != nil {
			return nil, fmt.Errorf("read batch input: %w", readErr)
		}
	}
	if len(items) == 0 {
		return nil, errors.New("batch input must contain at least one question")
	}
	return items, nil
}

func decodeBatchItem(data []byte) (batchItem, error) {
	if err := validateBatchKeys(data); err != nil {
		return batchItem{}, err
	}
	var input batchInputLine
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&input); err != nil {
		return batchItem{}, errors.New("invalid JSON object or field type")
	}
	if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return batchItem{}, errors.New("line must contain exactly one JSON object")
	}

	item := batchItem{id: input.ID, primitive: input.Type, question: input.Question}
	if strings.TrimSpace(item.id) == "" {
		return batchItem{}, errors.New("id must not be empty")
	}
	if item.id != strings.TrimSpace(item.id) {
		return batchItem{}, errors.New("id must not have surrounding whitespace")
	}
	if strings.TrimSpace(item.question) == "" {
		return batchItem{}, errors.New("question must not be empty")
	}

	hasConditions := len(input.Conditions) > 0
	hasLevels := len(input.Levels) > 0
	switch item.primitive {
	case primitiveNoul:
		if hasConditions || hasLevels {
			return batchItem{}, errors.New("noul does not accept conditions or levels")
		}
	case primitiveChoice:
		if !hasConditions || hasLevels {
			return batchItem{}, errors.New("choice requires conditions and does not accept levels")
		}
		if err := decodeBatchList(input.Conditions, &item.conditions, 2, 255, "conditions"); err != nil {
			return batchItem{}, err
		}
	case primitiveScore:
		if !hasLevels || hasConditions {
			return batchItem{}, errors.New("score requires levels and does not accept conditions")
		}
		if err := decodeBatchList(input.Levels, &item.levels, 2, 10, "levels"); err != nil {
			return batchItem{}, err
		}
	default:
		return batchItem{}, errors.New("type must be noul, choice, or score")
	}
	return item, nil
}

func validateBatchKeys(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil || fields == nil {
		return errors.New("invalid JSON object or field type")
	}
	allowed := map[string]bool{"id": true, "type": true, "question": true, "conditions": true, "levels": true}
	for key := range fields {
		if !allowed[key] {
			return fmt.Errorf("unknown field %q", key)
		}
	}
	return nil
}

func decodeBatchList(raw json.RawMessage, target *[]string, min, max int, name string) error {
	if bytes.Equal(bytes.TrimSpace(raw), []byte("null")) {
		return fmt.Errorf("%s must be an array of strings", name)
	}
	if err := json.Unmarshal(raw, target); err != nil {
		return fmt.Errorf("%s must be an array of strings", name)
	}
	if len(*target) < min || len(*target) > max {
		return fmt.Errorf("%s requires %d to %d distinct values", name, min, max)
	}
	for _, value := range *target {
		if value != strings.TrimSpace(value) {
			return fmt.Errorf("%s values must not have surrounding whitespace", name)
		}
	}
	if err := distinctNonempty(*target, name+" contains an empty or duplicate value"); err != nil {
		return err
	}
	return nil
}

func (i batchItem) evaluationOptions() evaluationOptions {
	return evaluationOptions{primitive: i.primitive, choices: i.conditions, levels: i.levels}
}

func (i batchItem) questionPayload() evaluationQuestion {
	payload := i.evaluationOptions().questionPayload()
	payload.Instructions = i.question
	return payload
}

func evaluateBatch(ctx context.Context, state string, items []batchItem, s settings, transport http.RoundTripper) ([]result, error) {
	questions := make(map[string]evaluationQuestion, len(items))
	for _, item := range items {
		questions[item.id] = item.questionPayload()
	}
	envelope, err := requestEvaluation(ctx, evaluationRequest{State: state, Model: s.Model, Questions: questions}, s, transport)
	if err != nil {
		return nil, err
	}
	if strings.TrimSpace(envelope.Model) == "" || len(envelope.Answers) != len(items) {
		return nil, errors.New("API response is missing a valid model or complete batch answers")
	}
	results := make([]result, len(items))
	for index, item := range items {
		raw, ok := envelope.Answers[item.id]
		if !ok {
			return nil, fmt.Errorf("API response is missing batch answer %q", item.id)
		}
		parsed, err := parseAnswer(raw, envelope.Model, item.evaluationOptions())
		if err != nil {
			return nil, fmt.Errorf("API response for %q: %w", item.id, err)
		}
		results[index] = parsed
	}
	return results, nil
}

func writeBatchResults(out io.Writer, items []batchItem, results []result) error {
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	for index, item := range items {
		result := results[index]
		value := batchOutput{ID: item.id, Type: item.primitive, Answer: result.Answer, Legend: result.Legend,
			Probabilities: result.Probabilities, Confidence: result.Confidence, Model: result.Model}
		if err := encoder.Encode(value); err != nil {
			return errors.New("could not encode batch result")
		}
	}
	written, err := out.Write(buffer.Bytes())
	if err == nil && written != buffer.Len() {
		return io.ErrShortWrite
	}
	return err
}
