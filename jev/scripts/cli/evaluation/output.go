package evaluation

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	cfg "opnay/jev/config"
	"opnay/jev/decision"
)

func writeResult(out io.Writer, r decision.Result, s cfg.Settings) error {
	if !s.JSON {
		_, err := fmt.Fprintln(out, r.Answer)
		return err
	}
	var value any = r
	if len(s.Pick) > 0 {
		all := resultFields(r)
		selected := make(map[string]any, len(s.Pick))
		for _, key := range s.Pick {
			selected[key] = all[key]
		}
		value = selected
	}
	data, err := json.Marshal(value)
	if err != nil {
		return errors.New("could not encode result")
	}
	_, err = out.Write(append(data, '\n'))
	return err
}

func resultFields(r decision.Result) map[string]any {
	fields := map[string]any{"answer": r.Answer, "model": r.Model}
	if r.Primitive == decision.Choice || r.Primitive == decision.Score {
		fields["probabilities"] = r.Probabilities
		fields["confidence"] = *r.Confidence
	}
	if r.Primitive == decision.Score {
		fields["legend"] = r.Legend
	}
	return fields
}
