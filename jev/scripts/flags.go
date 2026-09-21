package main

import (
	"errors"
	"flag"
	"io"
	"strings"
)

const usage = `Usage:
  jev --if <question and context> --conditions <a,b,...> [options]
  jev config
  jev config path
  jev config set <key> <value>
  jev config set api_key --stdin
  jev config unset <key>
  jev install [--force] [--dir <directory>]
  jev uninstall [--dir <directory>]
  jev doctor [--dir <directory>]

Options:
  --threshold <0..1>   Minimum probability of the selected answer (default: off)
  --model <name>       Model (default: jev-latest)
  --timeout <duration> Request timeout (default: 30s)
  --json[=false]      Output JSON (default: false)
  --pick <fields>      JSON fields: answer,probabilities,confidence,model
  --help              Show this help

Config: ~/.agents/jev.toml; flags > file > defaults.
Auth: TYPESAFE_API_KEY > config api_key. Token input is stdin only.
Exit: 0 success, 1 error, 2 threshold not met. Failures go to stderr.
Questions and choices are sent to the hosted TypeSafe API.
Maintenance is offline. Default install directory: ~/.local/bin.
Uninstall preserves config and tokens; doctor does not verify API authentication.
`

type evaluationOptions struct {
	question string
	choices  []string
	flags    config
}

func parseEvaluation(args []string) (evaluationOptions, error) {
	var o evaluationOptions
	var conditions, pick, model, timeout string
	var threshold float64
	var jsonOutput bool
	f := flag.NewFlagSet("jev", flag.ContinueOnError)
	f.SetOutput(io.Discard)
	f.StringVar(&o.question, "if", "", "question and context")
	f.StringVar(&conditions, "conditions", "", "comma-separated choices")
	f.Float64Var(&threshold, "threshold", 0, "minimum answer probability")
	f.StringVar(&model, "model", "jev-latest", "model")
	f.StringVar(&timeout, "timeout", "30s", "timeout")
	f.BoolVar(&jsonOutput, "json", false, "JSON output")
	f.StringVar(&pick, "pick", "", "JSON fields")
	if err := f.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return o, flag.ErrHelp
		}
		return o, errors.New("invalid options; see jev --help")
	}
	if f.NArg() != 0 {
		return o, errors.New("unexpected positional arguments; see jev --help")
	}
	if strings.TrimSpace(o.question) == "" {
		return o, errors.New("--if requires a question and its context")
	}
	o.choices = splitList(conditions)
	if len(o.choices) < 2 || len(o.choices) > 255 {
		return o, errors.New("--conditions requires 2 to 255 distinct choices")
	}
	seen := make(map[string]bool)
	for _, choice := range o.choices {
		if choice == "" || seen[choice] {
			return o, errors.New("--conditions contains an empty or duplicate choice")
		}
		seen[choice] = true
	}
	f.Visit(func(v *flag.Flag) {
		switch v.Name {
		case "threshold":
			o.flags.Threshold = &threshold
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
