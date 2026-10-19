package maintenance

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"

	cfg "opnay/jev/config"
	"opnay/jev/installation"
	"opnay/jev/providers"
)

type Command struct {
	Out         io.Writer
	Env         map[string]string
	ConfigPath  func() (string, error)
	Resolve     func(cfg.Config) cfg.Settings
	NewProvider func(cfg.Settings) providers.Adapter
	Executable  string
	LookupPath  func(string) (string, error)
}

func (a Command) Run(ctx context.Context, command string, args []string) error {
	o, err := parseLifecycle(command, args)
	if errors.Is(err, flag.ErrHelp) {
		_, err = io.WriteString(a.Out, usageFor(command))
		return err
	}
	if err != nil {
		return err
	}
	if command == "doctor" {
		return a.doctor(ctx, o.target, o.provider)
	}
	var message string
	if command == "install" {
		message, err = a.install(o)
	} else {
		message, err = installation.Uninstall(o.target)
	}
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(a.Out, message)
	return err
}
