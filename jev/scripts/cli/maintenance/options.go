package maintenance

import (
	"errors"
	"flag"
	"io"
	"os"
	"path/filepath"
)

type lifecycleOptions struct {
	target   string
	force    bool
	provider *string
}

func parseLifecycle(command string, args []string) (lifecycleOptions, error) {
	var o lifecycleOptions
	var directory, provider string
	f := flag.NewFlagSet(command, flag.ContinueOnError)
	f.SetOutput(io.Discard)
	f.StringVar(&directory, "dir", "", "installation directory (default: ~/.local/bin)")
	if command == "doctor" {
		f.StringVar(&provider, "provider", "jev", "decision provider")
	}
	if command == "install" {
		f.BoolVar(&o.force, "force", false, "replace an existing regular file")
	}
	if err := f.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return o, flag.ErrHelp
		}
		return o, errors.New("invalid maintenance options; see jev " + command + " --help")
	}
	if f.NArg() != 0 {
		return o, errors.New("use --dir for an installation directory; see jev " + command + " --help")
	}
	dirSet := false
	f.Visit(func(v *flag.Flag) {
		if v.Name == "provider" {
			o.provider = &provider
		}
		if v.Name == "dir" {
			dirSet = true
		}
	})
	if dirSet && directory == "" {
		return o, errors.New("--dir must not be empty; see jev " + command + " --help")
	}
	if directory == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return o, errors.New("cannot resolve user home for installation")
		}
		directory = filepath.Join(home, ".local", "bin")
	}
	target, err := filepath.Abs(filepath.Join(directory, "jev"))
	o.target = target
	return o, err
}
