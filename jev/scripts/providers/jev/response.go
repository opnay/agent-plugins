package jev

import (
	"bytes"
	"encoding/json"
	"errors"
	"math"
	"strconv"

	"opnay/jev/providers"
)

type apiAnswer struct {
	Type          string              `json:"type"`
	Noul          *float64            `json:"noul"`
	Choice        string              `json:"choice"`
	Score         *float64            `json:"score"`
	Legend        map[string]any      `json:"legend"`
	Probabilities map[string]*float64 `json:"probabilities"`
	Confidence    *float64            `json:"confidence"`
}

func parseAnswer(raw json.RawMessage, model string, q providers.Question) (providers.Result, error) {
	var answer apiAnswer
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.UseNumber()
	if err := decoder.Decode(&answer); err != nil || answer.Type != q.Type {
		return providers.Result{}, errors.New("API response type does not match the requested primitive")
	}
	switch q.Type {
	case "noul":
		return parseNoul(answer, model)
	case "choice":
		criteria, ok := q.Criteria.(map[string]any)
		if !ok {
			return providers.Result{}, errors.New("invalid Choice criteria")
		}
		choices := make([]string, 0, len(criteria))
		for label := range criteria {
			choices = append(choices, label)
		}
		return parseChoice(answer, model, choices)
	case "score":
		levels, ok := q.Criteria.([]any)
		if !ok {
			return providers.Result{}, errors.New("invalid Score criteria")
		}
		return parseScore(answer, model, levels)
	default:
		return providers.Result{}, errors.New("unknown response primitive")
	}
}

func parseNoul(answer apiAnswer, model string) (providers.Result, error) {
	if answer.Noul == nil || !probability(*answer.Noul) {
		return providers.Result{}, errors.New("API response is missing a valid Noul value")
	}
	return providers.Result{Primitive: "noul", Answer: *answer.Noul, Model: model}, nil
}

func parseChoice(answer apiAnswer, model string, choices []string) (providers.Result, error) {
	if answer.Confidence == nil || !probability(*answer.Confidence) || len(answer.Probabilities) != len(choices) {
		return providers.Result{}, errors.New("API response is missing valid Choice fields")
	}
	probabilities := make(map[string]float64, len(choices))
	total := 0.0
	for _, choice := range choices {
		p := answer.Probabilities[choice]
		if p == nil || !probability(*p) {
			return providers.Result{}, errors.New("API response contains a missing or invalid probability")
		}
		probabilities[choice] = *p
		total += *p
	}
	if math.Abs(total-1) > 1e-6 {
		return providers.Result{}, errors.New("API Choice probabilities do not sum to 1")
	}
	selected, ok := probabilities[answer.Choice]
	if !ok {
		return providers.Result{}, errors.New("API selected an unknown choice")
	}
	for _, p := range probabilities {
		if p > selected {
			return providers.Result{}, errors.New("API choice is not a highest-probability option")
		}
	}
	confidence := *answer.Confidence
	return providers.Result{Primitive: "choice", Answer: answer.Choice, Probabilities: probabilities,
		Confidence: &confidence, Model: model}, nil
}

func parseScore(answer apiAnswer, model string, levels []any) (providers.Result, error) {
	if err := providers.ValidateJSONContent(answer.Legend, 0); err != nil {
		return providers.Result{}, errors.New("API Score legend contains invalid JSON values")
	}
	max := float64(len(levels) - 1)
	if answer.Score == nil || !finite(*answer.Score) || *answer.Score < 0 || *answer.Score > max ||
		answer.Confidence == nil || !probability(*answer.Confidence) ||
		len(answer.Legend) != len(levels) || len(answer.Probabilities) != len(levels) {
		return providers.Result{}, errors.New("API response is missing valid Score fields")
	}
	legend := make(map[string]any, len(levels))
	probabilities := make(map[string]float64, len(levels))
	weighted := 0.0
	total := 0.0
	for i, level := range levels {
		key := strconv.Itoa(i)
		p := answer.Probabilities[key]
		if !providers.JSONValuesEqual(answer.Legend[key], level) || p == nil || !probability(*p) {
			return providers.Result{}, errors.New("API Score legend or probabilities do not match requested levels")
		}
		legend[key] = level
		probabilities[key] = *p
		total += *p
		weighted += float64(i) * *p
	}
	if math.Abs(total-1) > 1e-6 {
		return providers.Result{}, errors.New("API Score probabilities do not sum to 1")
	}
	if math.Abs(weighted-*answer.Score) > 1e-6 {
		return providers.Result{}, errors.New("API Score is inconsistent with its probabilities")
	}
	confidence := *answer.Confidence
	return providers.Result{Primitive: "score", Answer: *answer.Score, Legend: legend,
		Probabilities: probabilities, Confidence: &confidence, Model: model}, nil
}

func finite(value float64) bool      { return !math.IsNaN(value) && !math.IsInf(value, 0) }
func probability(value float64) bool { return finite(value) && value >= 0 && value <= 1 }
