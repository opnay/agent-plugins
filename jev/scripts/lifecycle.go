package main

import (
	"debug/buildinfo"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type lifecycleOptions struct {
	target string
	force  bool
}

func parseLifecycle(command string, args []string) (lifecycleOptions, error) {
	var o lifecycleOptions
	var directory string
	f := flag.NewFlagSet(command, flag.ContinueOnError)
	f.SetOutput(io.Discard)
	f.StringVar(&directory, "dir", "", "installation directory (default: ~/.local/bin)")
	if command == "install" {
		f.BoolVar(&o.force, "force", false, "replace an existing regular file")
	}
	if err := f.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return o, flag.ErrHelp
		}
		return o, errors.New("invalid maintenance options; see jev --help")
	}
	if f.NArg() != 0 {
		return o, errors.New("use --dir for an installation directory; see jev --help")
	}
	dirSet := false
	f.Visit(func(v *flag.Flag) {
		if v.Name == "dir" {
			dirSet = true
		}
	})
	if dirSet && directory == "" {
		return o, errors.New("--dir must not be empty")
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

func (a application) maintain(command string, args []string) error {
	o, err := parseLifecycle(command, args)
	if errors.Is(err, flag.ErrHelp) {
		_, err = io.WriteString(a.out, usage)
		return err
	}
	if err != nil {
		return err
	}
	if command == "doctor" {
		return a.doctor(o.target)
	}
	var message string
	if command == "install" {
		source := a.executable
		if source == "" {
			source, err = os.Executable()
			if err != nil {
				return fmt.Errorf("locate running executable: %w", err)
			}
		}
		message, err = installExecutable(source, o.target, o.force)
	} else {
		message, err = uninstallExecutable(o.target)
	}
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(a.out, message)
	return err
}

func regularTarget(path string) (os.FileInfo, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return nil, err
	}
	if !info.Mode().IsRegular() {
		return nil, errors.New("refusing a symlink or non-regular target")
	}
	return info, nil
}

func identifyJev(path string) error {
	info, err := buildinfo.ReadFile(path)
	if err != nil || info.Path != "opnay/jev" {
		return errors.New("target is not an identifiable Jev binary; leaving it untouched")
	}
	return nil
}
