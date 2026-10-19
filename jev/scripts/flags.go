package main

import (
	"errors"
	"flag"
	"io"
	"strings"
)

type evaluationOptions struct {
	primitive   primitive
	question    string
	inputPath   string
	state       any
	payload     *evaluationQuestion
	longContext bool
	choices     []string
	levels      []any
	flags       config
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
	f.StringVar(&o.inputPath, "input", "", "JSON state, instructions, and criteria file")
	f.BoolVar(&o.longContext, "long-context", false, "enable expanded context for this request")
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
	inputSet, legacySet := false, false
	f.Visit(func(v *flag.Flag) {
		inputSet = inputSet || v.Name == "input"
		legacySet = legacySet || v.Name == "if" || v.Name == "conditions" || v.Name == "level"
	})
	if inputSet {
		if legacySet || strings.TrimSpace(o.inputPath) == "" {
			return o, errors.New("--input requires a path and cannot be combined with --if, --conditions, or --level" + help)
		}
		if err := o.readInput(); err != nil {
			return o, errors.New(err.Error() + help)
		}
	} else {
		if strings.TrimSpace(o.question) == "" {
			return o, errors.New("--if requires a question and its context" + help)
		}
		if err := o.validateCriteria(conditions, levels); err != nil {
			return o, errors.New(err.Error() + help)
		}
		o.state = o.question
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
		if len(levels) < 2 || len(levels) > 10 {
			return errors.New("--level requires 2 to 10 ordered levels")
		}
		for i := range levels {
			levels[i] = strings.TrimSpace(levels[i])
		}
		if err := distinctNonempty(levels, "--level contains an empty or duplicate level"); err != nil {
			return err
		}
		for _, level := range levels {
			o.levels = append(o.levels, level)
		}
		return nil
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
