package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type application struct {
	in         io.Reader
	out        io.Writer
	errOut     io.Writer
	path       string // Tests supply a temporary path, never the user's config.
	envToken   string
	transport  http.RoundTripper
	executable string // Optional test source; production uses os.Executable.
	lookupPath func(string) (string, error)
}

func (a application) configPath() (string, error) {
	if a.path != "" {
		return a.path, nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", errors.New("cannot resolve user home for config")
	}
	return filepath.Join(home, ".agents", "jev.toml"), nil
}

func (a application) run(ctx context.Context, args []string) int {
	code, err := a.execute(ctx, args)
	if err != nil {
		fmt.Fprintln(a.errOut, "jev:", err)
	}
	return code
}

func (a application) execute(ctx context.Context, args []string) (int, error) {
	if len(args) == 0 || (len(args) == 1 && (args[0] == "--help" || args[0] == "-h")) {
		_, err := io.WriteString(a.out, usage)
		return exitCode(err), err
	}
	if args[0] == "config" {
		err := a.configure(args[1:])
		return exitCode(err), err
	}
	if args[0] == "install" || args[0] == "uninstall" || args[0] == "doctor" {
		err := a.maintain(args[0], args[1:])
		return exitCode(err), err
	}
	o, err := parseEvaluation(args)
	if errors.Is(err, flag.ErrHelp) {
		_, err = io.WriteString(a.out, usage)
		return exitCode(err), err
	}
	if err != nil {
		return 1, err
	}
	path, err := a.configPath()
	if err != nil {
		return 1, err
	}
	c, err := readConfig(path)
	if err != nil {
		return 1, err
	}
	c.override(o.flags)
	s := c.resolve(a.envToken)
	if err := s.validate(); err != nil {
		return 1, err
	}
	if len(s.Pick) > 0 && !s.JSON {
		return 1, errors.New("pick requires JSON output; use --json or --pick ''")
	}
	if s.APIKey == "" {
		return 1, errors.New("API key missing; set TYPESAFE_API_KEY or use jev config set api_key --stdin")
	}
	result, err := evaluate(ctx, o, s, a.transport)
	if err != nil {
		return 1, err
	}
	if s.Threshold != nil && result.Probabilities[result.Answer] < *s.Threshold {
		return 2, fmt.Errorf("threshold not met: %g < %g (answer: %s)",
			result.Probabilities[result.Answer], *s.Threshold, result.Answer)
	}
	err = writeResult(a.out, result, s)
	return exitCode(err), err
}

func exitCode(err error) int {
	if err != nil {
		return 1
	}
	return 0
}
