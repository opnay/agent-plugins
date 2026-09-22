package main

import (
	"errors"
	"flag"
	"io"
	"strings"
)

type evaluationOptions struct {
	primitive primitive
	question  string
	choices   []string
	levels    []string
	flags     config
}

type repeatedStrings []string

func (v *repeatedStrings) String() string { return strings.Join(*v, ",") }
func (v *repeatedStrings) Set(value string) error {
	*v = append(*v, value)
	return nil
}

func parseEvaluation(args []string) (evaluationOptions, error) {
	o := evaluationOptions{primitive: primitiveChoice}
	if len(args) > 0 {
		switch primitive(args[0]) {
		case primitiveNoul, primitiveChoice, primitiveScore:
			o.primitive = primitive(args[0])
			args = args[1:]
		}
	}
	var conditions, pick, model, timeout string
	var levels repeatedStrings
	var threshold, minScore float64
	var jsonOutput bool
	f := flag.NewFlagSet("jev "+string(o.primitive), flag.ContinueOnError)
	f.SetOutput(io.Discard)
	f.StringVar(&o.question, "if", "", "question and context")
	if o.primitive == primitiveChoice {
		f.StringVar(&conditions, "conditions", "", "comma-separated choices")
	}
	if o.primitive == primitiveScore {
		f.Var(&levels, "level", "ordered score level; repeat for each level")
		f.Float64Var(&minScore, "min-score", 0, "minimum score")
	} else {
		f.Float64Var(&threshold, "threshold", 0, "minimum result probability")
	}
	f.StringVar(&model, "model", "jev-latest", "model")
	f.StringVar(&timeout, "timeout", "30s", "timeout")
	f.BoolVar(&jsonOutput, "json", false, "JSON output")
	f.StringVar(&pick, "pick", "", "JSON fields")
	if err := f.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return o, flag.ErrHelp
		}
		return o, errors.New("invalid options; see jev " + string(o.primitive) + " --help")
	}
	if f.NArg() != 0 {
		return o, errors.New("unexpected positional arguments; see jev " + string(o.primitive) + " --help")
	}
	help := "; see jev " + string(o.primitive) + " --help"
	if strings.TrimSpace(o.question) == "" {
		return o, errors.New("--if requires a question and its context" + help)
	}
	if err := o.validateCriteria(conditions, levels); err != nil {
		return o, errors.New(err.Error() + help)
	}
	f.Visit(func(v *flag.Flag) {
		switch v.Name {
		case "threshold":
			o.flags.Threshold = &threshold
		case "min-score":
			o.flags.MinScore = &minScore
		case "model":
			o.flags.Model = &model
		case "timeout":
			o.flags.Timeout = &timeout
		case "json":
			o.flags.JSON = &jsonOutput
		case "pick":
			fields := splitList(pick)
			o.flags.Pick = &fields
		}
	})
	return o, nil
}

func (o *evaluationOptions) validateCriteria(conditions string, levels []string) error {
	switch o.primitive {
	case primitiveNoul:
		return nil
	case primitiveChoice:
		o.choices = splitList(conditions)
		if len(o.choices) < 2 || len(o.choices) > 255 {
			return errors.New("--conditions requires 2 to 255 distinct choices")
		}
		return distinctNonempty(o.choices, "--conditions contains an empty or duplicate choice")
	case primitiveScore:
		o.levels = append([]string(nil), levels...)
		if len(o.levels) < 2 || len(o.levels) > 10 {
			return errors.New("--level requires 2 to 10 ordered levels")
		}
		for i := range o.levels {
			o.levels[i] = strings.TrimSpace(o.levels[i])
		}
		return distinctNonempty(o.levels, "--level contains an empty or duplicate level")
	default:
		return errors.New("unknown primitive")
	}
}

func distinctNonempty(values []string, message string) error {
	seen := make(map[string]bool, len(values))
	for _, value := range values {
		if value == "" || seen[value] {
			return errors.New(message)
		}
		seen[value] = true
	}
	return nil
}
