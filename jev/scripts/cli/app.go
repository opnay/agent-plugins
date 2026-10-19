package cli

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"opnay/jev/cli/configcmd"
	"opnay/jev/cli/evaluation"
	"opnay/jev/cli/maintenance"
	cfg "opnay/jev/config"
)

type App struct {
	In         io.Reader
	Out        io.Writer
	ErrOut     io.Writer
	path       string // Tests supply a temporary path, never the user's config.
	Env        map[string]string
	transport  http.RoundTripper
	executable string // Optional test source; production uses os.Executable.
	lookupPath func(string) (string, error)
}

func (a App) configPath() (string, error) {
	if a.path != "" {
		return a.path, nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", errors.New("cannot resolve user home for config")
	}
	return filepath.Join(home, ".agents", "jev.toml"), nil
}

func (a App) Run(ctx context.Context, args []string) int {
	code, err := a.execute(ctx, args)
	if err != nil {
		if _, report := err.(maintenance.Report); report {
			fmt.Fprintln(a.ErrOut, err)
		} else {
			fmt.Fprintln(a.ErrOut, "jev:", err)
		}
	}
	return code
}

func (a App) execute(ctx context.Context, args []string) (int, error) {
	if len(args) == 0 || (len(args) == 1 && (args[0] == "--help" || args[0] == "-h")) {
		_, err := io.WriteString(a.Out, usage)
		return exitCode(err), err
	}
	if args[0] == "config" {
		command := configcmd.Command{In: a.In, Out: a.Out, Env: a.Env, ConfigPath: a.configPath}
		err := command.Run(args[1:])
		return exitCode(err), err
	}
	if args[0] == "install" || args[0] == "uninstall" || args[0] == "doctor" {
		command := maintenance.Command{
			Out: a.Out, Env: a.Env,
			ConfigPath: a.configPath, Resolve: a.resolve, NewProvider: a.newProvider,
			Executable: a.executable, LookupPath: a.lookupPath,
		}
		err := command.Run(ctx, args[0], args[1:])
		return exitCode(err), err
	}
	command := evaluation.Command{Out: a.Out, ConfigPath: a.configPath, Resolve: a.resolve, NewProvider: a.newProvider}
	return command.Run(ctx, args)
}

func exitCode(err error) int {
	if err != nil {
		return 1
	}
	return 0
}

func (a App) resolve(c cfg.Config) cfg.Settings {
	return c.Resolve(a.Env[cfg.ProviderEnvironment(c.Provider())])
}
