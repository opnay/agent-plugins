package configcmd

import (
	"errors"
	"fmt"
	"io"

	"github.com/pelletier/go-toml/v2"
	cfg "opnay/jev/config"
)

type Command struct {
	In         io.Reader
	Out        io.Writer
	Env        map[string]string
	ConfigPath func() (string, error)
}

func (a Command) Run(args []string) error {
	if len(args) == 1 && (args[0] == "--help" || args[0] == "-h") {
		_, err := io.WriteString(a.Out, configUsage)
		return err
	}
	path, err := a.ConfigPath()
	if err != nil {
		return err
	}
	if len(args) == 1 && args[0] == "path" {
		_, err = fmt.Fprintln(a.Out, path)
		return err
	}
	if len(args) == 1 && args[0] == "init" {
		return cfg.Init(path)
	}
	if len(args) != 0 && !(len(args) == 3 && args[0] == "set") && !(len(args) == 2 && args[0] == "unset") {
		return errors.New("invalid config command; see jev config --help")
	}
	// Parse and validate command input before reading the destination file.
	var pending cfg.Config
	if len(args) > 0 {
		if _, err := cfg.Key(args[1]); err != nil {
			return configHelpError(err)
		}
		if args[0] == "set" {
			if err := pending.Set(args[1], args[2], a.In); err != nil {
				return configHelpError(err)
			}
			if err := pending.Validate(); err != nil {
				return configHelpError(err)
			}
		}
	}
	exists, err := cfg.FileStatus(path)
	if err != nil {
		return err
	}
	c, err := cfg.Read(path)
	if err != nil {
		return err
	}
	if len(args) == 0 {
		shown := c.Display(a.Env)
		data, err := toml.Marshal(shown)
		if err != nil {
			return errors.New("could not encode config display")
		}
		_, err = a.Out.Write(data)
		return err
	}
	if args[0] == "set" {
		if !exists {
			c, err = cfg.Template()
			if err != nil {
				return err
			}
		}
		c.ApplySet(args[1], pending)
		if !exists {
			return cfg.Create(path, c)
		}
		return cfg.Write(path, c)
	}
	changed, err := c.Unset(args[1])
	if err != nil {
		return configHelpError(err)
	}
	if !changed {
		return nil
	}
	return cfg.Write(path, c)
}

func configHelpError(err error) error { return fmt.Errorf("%w; see jev config --help", err) }
