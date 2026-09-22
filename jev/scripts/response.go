package main

import (
	"encoding/json"
	"errors"
	"math"
	"strconv"
	"strings"
)

type apiAnswer struct {
	Type          primitive           `json:"type"`
	Noul          *float64            `json:"noul"`
	Choice        string              `json:"choice"`
	Score         *float64            `json:"score"`
	Legend        map[string]string   `json:"legend"`
	Probabilities map[string]*float64 `json:"probabilities"`
	Confidence    *float64            `json:"confidence"`
}

func parseResponse(envelope evaluationEnvelope, o evaluationOptions) (result, error) {
	raw, ok := envelope.Answers["result"]
	if !ok || strings.TrimSpace(envelope.Model) == "" {
		return result{}, errors.New("API response is missing a valid model or result")
	}
	return parseAnswer(raw, envelope.Model, o)
}

func parseAnswer(raw json.RawMessage, model string, o evaluationOptions) (result, error) {
	var answer apiAnswer
	if err := json.Unmarshal(raw, &answer); err != nil || answer.Type != o.primitive {
		return result{}, errors.New("API response type does not match the requested primitive")
	}
	switch o.primitive {
	case primitiveNoul:
		return parseNoul(answer, model)
	case primitiveChoice:
		return parseChoice(answer, model, o.choices)
	case primitiveScore:
		return parseScore(answer, model, o.levels)
	default:
		return result{}, errors.New("unknown response primitive")
	}
}

func parseNoul(answer apiAnswer, model string) (result, error) {
	if answer.Noul == nil || !probability(*answer.Noul) {
		return result{}, errors.New("API response is missing a valid Noul value")
	}
	return result{Primitive: primitiveNoul, Answer: *answer.Noul, Model: model}, nil
}

func parseChoice(answer apiAnswer, model string, choices []string) (result, error) {
	if answer.Confidence == nil || !probability(*answer.Confidence) || len(answer.Probabilities) != len(choices) {
		return result{}, errors.New("API response is missing valid Choice fields")
	}
	probabilities := make(map[string]float64, len(choices))
	total := 0.0
	for _, choice := range choices {
		p := answer.Probabilities[choice]
		if p == nil || !probability(*p) {
			return result{}, errors.New("API response contains a missing or invalid probability")
		}
		probabilities[choice] = *p
		total += *p
	}
	if math.Abs(total-1) > 1e-6 {
		return result{}, errors.New("API Choice probabilities do not sum to 1")
	}
	selected, ok := probabilities[answer.Choice]
	if !ok {
		return result{}, errors.New("API selected an unknown choice")
	}
	for _, p := range probabilities {
		if p > selected {
			return result{}, errors.New("API choice is not a highest-probability option")
		}
	}
	confidence := *answer.Confidence
	return result{Primitive: primitiveChoice, Answer: answer.Choice, Probabilities: probabilities,
		Confidence: &confidence, Model: model}, nil
}

func parseScore(answer apiAnswer, model string, levels []string) (result, error) {
	max := float64(len(levels) - 1)
	if answer.Score == nil || !finite(*answer.Score) || *answer.Score < 0 || *answer.Score > max ||
		answer.Confidence == nil || !probability(*answer.Confidence) ||
		len(answer.Legend) != len(levels) || len(answer.Probabilities) != len(levels) {
		return result{}, errors.New("API response is missing valid Score fields")
	}
	legend := make(map[string]string, len(levels))
	probabilities := make(map[string]float64, len(levels))
	weighted := 0.0
	total := 0.0
	for i, level := range levels {
		key := strconv.Itoa(i)
		p := answer.Probabilities[key]
		if answer.Legend[key] != level || p == nil || !probability(*p) {
			return result{}, errors.New("API Score legend or probabilities do not match requested levels")
		}
		legend[key] = level
		probabilities[key] = *p
		total += *p
		weighted += float64(i) * *p
	}
	if math.Abs(total-1) > 1e-6 {
		return result{}, errors.New("API Score probabilities do not sum to 1")
	}
	if math.Abs(weighted-*answer.Score) > 1e-6 {
		return result{}, errors.New("API Score is inconsistent with its probabilities")
	}
	confidence := *answer.Confidence
	return result{Primitive: primitiveScore, Answer: *answer.Score, Legend: legend,
		Probabilities: probabilities, Confidence: &confidence, Model: model}, nil
}
