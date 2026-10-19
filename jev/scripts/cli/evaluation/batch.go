package evaluation

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"

	cfg "opnay/jev/config"
	"opnay/jev/decision"
)

type batchOptions struct {
	statePath   string
	inputPath   string
	flags       cfg.Flags
	stateFormat string
	longContext bool
}

type batchInputLine struct {
	ID           string             `json:"id"`
	Type         decision.Primitive `json:"type"`
	Question     string             `json:"question"`
	Instructions json.RawMessage    `json:"instructions"`
	Criteria     json.RawMessage    `json:"criteria"`
	Conditions   json.RawMessage    `json:"conditions"`
	Levels       json.RawMessage    `json:"levels"`
}

type batchItem struct {
	id      string
	options evaluationOptions
}

type batchOutput struct {
	ID            string             `json:"id"`
	Type          decision.Primitive `json:"type"`
	Answer        any                `json:"answer"`
	Legend        map[string]any     `json:"legend,omitempty"`
	Probabilities map[string]float64 `json:"probabilities,omitempty"`
	Confidence    *float64           `json:"confidence,omitempty"`
	Model         string             `json:"model"`
}

func parseBatch(args []string) (batchOptions, error) {
	var o batchOptions
	var model, timeout, provider string
	f := flag.NewFlagSet("jev batch", flag.ContinueOnError)
	f.SetOutput(io.Discard)
	f.StringVar(&o.statePath, "state-file", "", "shared state file")
	f.StringVar(&o.inputPath, "input", "", "JSONL question file")
	f.StringVar(&o.stateFormat, "state-format", "text", "shared state format: text or json")
	f.BoolVar(&o.longContext, "long-context", false, "enable expanded context for this request")
	f.StringVar(&provider, "provider", "jev", "decision provider")
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
	if o.stateFormat != "text" && o.stateFormat != "json" {
		return o, errors.New("--state-format must be text or json; see jev batch --help")
	}
	f.Visit(func(v *flag.Flag) {
		switch v.Name {
		case "provider":
			o.flags.Provider = &provider
		case "model":
			o.flags.Model = &model
		case "timeout":
			o.flags.Timeout = &timeout
		}
	})
	if err := o.flags.Validate(); err != nil {
		return o, fmt.Errorf("%w; see jev batch --help", err)
	}
	return o, nil
}

func (a Command) executeBatch(ctx context.Context, o batchOptions) error {
	path, err := a.ConfigPath()
	if err != nil {
		return err
	}
	c, err := cfg.Read(path)
	if err != nil {
		return err
	}
	c.Override(o.flags)
	s := a.Resolve(c)
	s.Threshold = nil
	s.MinScore = nil
	s.JSON = false
	s.Pick = nil
	if err := s.Validate(); err != nil {
		return fmt.Errorf("%w; see jev batch --help", err)
	}
	if s.APIKey == "" {
		return errors.New("API key missing; set TYPESAFE_API_KEY or use jev config set providers.jev.api_key --stdin")
	}
	state, err := readBatchState(o.statePath, o.stateFormat)
	if err != nil {
		return err
	}
	items, err := readBatchItems(o.inputPath)
	if err != nil {
		return err
	}
	results, err := a.evaluateBatch(ctx, state, items, s, o.longContext)
	if err != nil {
		return err
	}
	return writeBatchResults(a.Out, items, results)
}

func readBatchState(path, format string) (any, error) {
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
	if format == "json" {
		value, err := decision.DecodeJSON(data)
		if err != nil {
			return nil, fmt.Errorf("state: %w", err)
		}
		if err := decision.ValidateEntry(value, false); err != nil {
			return nil, fmt.Errorf("state: %w", err)
		}
		return value, nil
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

	item := batchItem{id: input.ID, options: evaluationOptions{primitive: input.Type}}
	if strings.TrimSpace(item.id) == "" {
		return batchItem{}, errors.New("id must not be empty")
	}
	if item.id != strings.TrimSpace(item.id) {
		return batchItem{}, errors.New("id must not have surrounding whitespace")
	}
	var fields map[string]json.RawMessage
	_ = json.Unmarshal(data, &fields)
	_, hasQuestion := fields["question"]
	hasInstructions := len(input.Instructions) > 0
	if hasQuestion == hasInstructions {
		return batchItem{}, errors.New("use exactly one of question or instructions")
	}
	var instructions any = input.Question
	if hasInstructions {
		value, err := decision.DecodeJSON(input.Instructions)
		if err != nil {
			return batchItem{}, fmt.Errorf("instructions: %w", err)
		}
		instructions = value
	} else if strings.TrimSpace(input.Question) == "" {
		return batchItem{}, errors.New("question must not be empty")
	}
	hasConditions := len(input.Conditions) > 0
	hasLevels := len(input.Levels) > 0
	if len(input.Criteria) > 0 {
		if hasConditions || hasLevels {
			return batchItem{}, errors.New("criteria cannot be combined with conditions or levels")
		}
		criteria, err := decision.DecodeJSON(input.Criteria)
		if err != nil {
			return batchItem{}, fmt.Errorf("criteria: %w", err)
		}
		if err := item.options.setStructuredQuestion(instructions, criteria); err != nil {
			return batchItem{}, err
		}
	} else {
		switch input.Type {
		case decision.Noul:
			if hasConditions || hasLevels {
				return batchItem{}, errors.New("noul does not accept conditions or levels")
			}
		case decision.Choice:
			if !hasConditions || hasLevels {
				return batchItem{}, errors.New("choice requires conditions and does not accept levels")
			}
			if err := decodeBatchList(input.Conditions, &item.options.choices, 2, 255, "conditions"); err != nil {
				return batchItem{}, err
			}
		case decision.Score:
			if !hasLevels || hasConditions {
				return batchItem{}, errors.New("score requires levels and does not accept conditions")
			}
			var levels []string
			if err := decodeBatchList(input.Levels, &levels, 2, 10, "levels"); err != nil {
				return batchItem{}, err
			}
			for _, level := range levels {
				item.options.levels = append(item.options.levels, level)
			}
		default:
			return batchItem{}, errors.New("type must be noul, choice, or score")
		}
		if err := decision.ValidateEntry(instructions, true); err != nil {
			return batchItem{}, fmt.Errorf("instructions: %w", err)
		}
		payload := item.options.questionPayload()
		payload.Instructions = instructions
		item.options.payload = &payload
	}
	return item, nil
}

func validateBatchKeys(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil || fields == nil {
		return errors.New("invalid JSON object or field type")
	}
	allowed := map[string]bool{"id": true, "type": true, "question": true, "conditions": true, "levels": true, "instructions": true, "criteria": true}
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

func (a Command) evaluateBatch(ctx context.Context, state any, items []batchItem, s cfg.Settings, longContext bool) ([]decision.Result, error) {
	questions := make(map[string]decision.Question, len(items))
	for _, item := range items {
		questions[item.id] = item.options.questionPayload()
	}
	answers, err := a.NewProvider(s).Evaluate(ctx, decision.Request{State: state, Model: s.Model, Questions: questions, LongContext: longContext})
	if err != nil {
		return nil, err
	}
	results := make([]decision.Result, len(items))
	for i, item := range items {
		results[i] = answers[item.id]
	}
	return results, nil
}

func writeBatchResults(out io.Writer, items []batchItem, results []decision.Result) error {
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	for index, item := range items {
		result := results[index]
		value := batchOutput{ID: item.id, Type: item.options.primitive, Answer: result.Answer, Legend: result.Legend,
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
