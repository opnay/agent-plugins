package evaluation

import (
	"context"

	cfg "opnay/jev/config"
	"opnay/jev/decision"
)

func (a Command) evaluate(ctx context.Context, o evaluationOptions, s cfg.Settings) (decision.Result, error) {
	request := decision.Request{
		State: o.state, Model: s.Model, LongContext: o.longContext,
		Questions: map[string]decision.Question{"result": o.questionPayload()},
	}
	answers, err := a.NewProvider(s).Evaluate(ctx, request)
	if err != nil {
		return decision.Result{}, err
	}
	return answers["result"], nil
}

func (o evaluationOptions) questionPayload() decision.Question {
	if o.payload != nil {
		return *o.payload
	}
	switch o.primitive {
	case decision.Noul:
		return decision.Question{Type: string(o.primitive), Instructions: "Answer whether the proposition or question in state is true."}
	case decision.Score:
		return decision.Question{Type: string(o.primitive), Instructions: "Rate the state against the ordered criteria.", Criteria: o.levels}
	default:
		criteria := make(map[string]any, len(o.choices))
		for _, choice := range o.choices {
			criteria[choice] = nil
		}
		return decision.Question{Type: string(o.primitive), Instructions: "Select the option that best answers the question in state.", Criteria: criteria}
	}
}
