package evaluation

import (
	"context"
	"errors"
	"flag"
	"io"

	cfg "opnay/jev/config"
	"opnay/jev/providers"
)

type Command struct {
	Out         io.Writer
	ConfigPath  func() (string, error)
	Resolve     func(cfg.Config) cfg.Settings
	NewProvider func(cfg.Settings) providers.Adapter
}

func (a Command) Run(ctx context.Context, args []string) (int, error) {
	if len(args) == 0 || args[0] != "batch" {
		return a.executeEvaluation(ctx, args)
	}
	o, err := parseBatch(args[1:])
	if errors.Is(err, flag.ErrHelp) {
		_, err = io.WriteString(a.Out, batchUsage)
		return exitCode(err), err
	}
	if err != nil {
		return 1, err
	}
	err = a.executeBatch(ctx, o)
	return exitCode(err), err
}

func exitCode(err error) int {
	if err != nil {
		return 1
	}
	return 0
}
