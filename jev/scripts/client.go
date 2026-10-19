package main

import (
	"context"
	"net/http"
	"time"

	"opnay/jev/providers"
	jevprovider "opnay/jev/providers/jev"
)

func newProvider(s settings, transport http.RoundTripper) providers.Adapter {
	timeout, _ := time.ParseDuration(s.Timeout)
	return jevprovider.New(s.APIKey, timeout, transport)
}

func evaluate(ctx context.Context, o evaluationOptions, s settings, transport http.RoundTripper) (result, error) {
	request := providers.Request{
		State: o.state, Model: s.Model, LongContext: o.longContext,
		Questions: map[string]providers.Question{"result": o.questionPayload()},
	}
	answers, err := newProvider(s, transport).Evaluate(ctx, request)
	if err != nil {
		return result{}, err
	}
	return answers["result"], nil
}

func (o evaluationOptions) questionPayload() providers.Question {
	if o.payload != nil {
		return *o.payload
	}
	switch o.primitive {
	case primitiveNoul:
		return providers.Question{Type: string(o.primitive), Instructions: "Answer whether the proposition or question in state is true."}
	case primitiveScore:
		return providers.Question{Type: string(o.primitive), Instructions: "Rate the state against the ordered criteria.", Criteria: o.levels}
	default:
		criteria := make(map[string]any, len(o.choices))
		for _, choice := range o.choices {
			criteria[choice] = nil
		}
		return providers.Question{Type: string(o.primitive), Instructions: "Select the option that best answers the question in state.", Criteria: criteria}
	}
}
