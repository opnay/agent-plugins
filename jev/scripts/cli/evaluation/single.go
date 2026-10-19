package evaluation

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"

	cfg "opnay/jev/config"
	"opnay/jev/decision"
)

func (a Command) executeEvaluation(ctx context.Context, args []string) (int, error) {
	o, err := parseEvaluation(args)
	if errors.Is(err, flag.ErrHelp) {
		_, err = io.WriteString(a.Out, usageFor(string(o.primitive)))
		return exitCode(err), err
	}
	if err != nil {
		return 1, err
	}
	if err := validateEvaluationFlags(o); err != nil {
		return 1, evaluationHelpError(o.primitive, err)
	}
	path, err := a.ConfigPath()
	if err != nil {
		return 1, err
	}
	c, err := cfg.Read(path)
	if err != nil {
		return 1, err
	}
	c.Override(o.flags)
	s := a.Resolve(c)
	if err := s.Validate(); err != nil {
		return 1, evaluationHelpError(o.primitive, err)
	}
	if len(s.Pick) > 0 && !s.JSON {
		return 1, evaluationHelpError(o.primitive, errors.New("pick requires JSON output; use --json or --pick ''"))
	}
	if err := o.primitive.ValidatePick(s.Pick); err != nil {
		return 1, evaluationHelpError(o.primitive, err)
	}
	if o.primitive == decision.Score && s.MinScore != nil && *s.MinScore > float64(len(o.levels)-1) {
		return 1, evaluationHelpError(o.primitive, fmt.Errorf("min_score must not exceed the highest Score level %d", len(o.levels)-1))
	}
	if s.APIKey == "" {
		return 1, errors.New("API key missing; set TYPESAFE_API_KEY or use jev config set providers.jev.api_key --stdin")
	}
	result, err := a.evaluate(ctx, o, s)
	if err != nil {
		return 1, err
	}
	if o.primitive == decision.Choice && s.Threshold != nil {
		answer := result.Answer.(string)
		if result.Probabilities[answer] < *s.Threshold {
			return 2, fmt.Errorf("threshold not met: %g < %g (answer: %s)",
				result.Probabilities[answer], *s.Threshold, answer)
		}
	}
	if o.primitive == decision.Noul && s.Threshold != nil && result.Answer.(float64) < *s.Threshold {
		return 2, fmt.Errorf("threshold not met: %g < %g", result.Answer, *s.Threshold)
	}
	if o.primitive == decision.Score && s.MinScore != nil && result.Answer.(float64) < *s.MinScore {
		return 2, fmt.Errorf("min_score not met: %g < %g", result.Answer, *s.MinScore)
	}
	err = writeResult(a.Out, result, s)
	return exitCode(err), err
}

func evaluationHelpError(p decision.Primitive, err error) error {
	return fmt.Errorf("%w; see jev %s --help", err, p)
}

func validateEvaluationFlags(o evaluationOptions) error {
	if err := o.flags.Validate(); err != nil {
		return err
	}
	if o.flags.Pick != nil {
		if err := o.primitive.ValidatePick(*o.flags.Pick); err != nil {
			return err
		}
		if len(*o.flags.Pick) > 0 && o.flags.JSON != nil && !*o.flags.JSON {
			return errors.New("pick requires JSON output; use --json or --pick ''")
		}
	}
	if o.primitive == decision.Score && o.flags.MinScore != nil && *o.flags.MinScore > float64(len(o.levels)-1) {
		return fmt.Errorf("min_score must not exceed the highest Score level %d", len(o.levels)-1)
	}
	return nil
}
