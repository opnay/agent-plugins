package main

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

func (a application) configure(args []string) error {
	if len(args) == 1 && (args[0] == "--help" || args[0] == "-h") {
		_, err := io.WriteString(a.out, usageFor("config"))
		return err
	}
	path, err := a.configPath()
	if err != nil {
		return err
	}
	if len(args) == 1 && args[0] == "path" {
		_, err = fmt.Fprintln(a.out, path)
		return err
	}
	if len(args) == 1 && args[0] == "init" {
		return initConfig(path)
	}
	if len(args) != 0 && !(len(args) == 3 && args[0] == "set") && !(len(args) == 2 && args[0] == "unset") {
		return errors.New("invalid config command; see jev config --help")
	}
	// Parse and validate command input before reading the destination file.
	var pending config
	if len(args) > 0 {
		if _, err := configKey(args[1]); err != nil {
			return configHelpError(err)
		}
		if args[0] == "set" {
			if err := pending.set(args[1], args[2], a.in); err != nil {
				return configHelpError(err)
			}
			if err := pending.validate(); err != nil {
				return configHelpError(err)
			}
		}
	}
	exists, err := configFileStatus(path)
	if err != nil {
		return err
	}
	c, err := readConfig(path)
	if err != nil {
		return err
	}
	if len(args) == 0 {
		shown := a.configDisplay(c)
		data, err := toml.Marshal(shown)
		if err != nil {
			return errors.New("could not encode config display")
		}
		_, err = a.out.Write(data)
		return err
	}
	if args[0] == "set" {
		if !exists {
			c, err = templateConfig()
			if err != nil {
				return err
			}
		}
		c.applySet(args[1], pending)
		if !exists {
			return createConfig(path, c)
		}
		return writeConfig(path, c)
	}
	changed, err := c.unset(args[1])
	if err != nil {
		return configHelpError(err)
	}
	if !changed {
		return nil
	}
	return writeConfig(path, c)
}

func (a application) configDisplay(c config) config {
	s := c.resolve("")
	shown := config{Defaults: defaults{Provider: &s.Provider, Timeout: &s.Timeout, JSON: &s.JSON, Pick: &s.Pick, Threshold: s.Threshold, MinScore: s.MinScore}, Providers: map[string]providerConfig{}}
	for _, name := range []string{"jev", "openrouter", "openai"} {
		p, exists := c.Providers[name]
		if name == "jev" && p.Model == nil {
			model := "jev-latest"
			p.Model = &model
			exists = true
		}
		token := strings.TrimSpace(a.env[providerEnvironment(name)])
		if token != "" {
			p.APIKey = token
			exists = true
		}
		if p.APIKey != "" {
			p.APIKey = "********"
		}
		if exists {
			shown.Providers[name] = p
		}
	}
	return shown
}

func configHelpError(err error) error { return fmt.Errorf("%w; see jev config --help", err) }

// Return the provider name, or an empty string for a common defaults key.
func configKey(key string) (string, error) {
	parts := strings.Split(key, ".")
	if len(parts) == 2 && parts[0] == "defaults" {
		switch parts[1] {
		case "provider", "threshold", "min_score", "timeout", "json", "pick":
			return "", nil
		}
	}
	if len(parts) == 3 && parts[0] == "providers" && providerEnvironment(parts[1]) != "" && (parts[2] == "model" || parts[2] == "api_key") {
		return parts[1], nil
	}
	return "", errors.New("use defaults.provider/threshold/min_score/timeout/json/pick or providers.<jev|openrouter|openai>.model/api_key")
}

func (c *config) applySet(key string, pending config) {
	name, _ := configKey(key)
	if name != "" {
		if c.Providers == nil {
			c.Providers = make(map[string]providerConfig)
		}
		p := c.Providers[name]
		value := pending.Providers[name]
		if strings.HasSuffix(key, ".api_key") {
			p.APIKey = value.APIKey
		} else {
			p.Model = value.Model
		}
		c.Providers[name] = p
		return
	}
	switch key {
	case "defaults.provider":
		c.Defaults.Provider = pending.Defaults.Provider
	case "defaults.threshold":
		c.Defaults.Threshold = pending.Defaults.Threshold
	case "defaults.min_score":
		c.Defaults.MinScore = pending.Defaults.MinScore
	case "defaults.timeout":
		c.Defaults.Timeout = pending.Defaults.Timeout
	case "defaults.json":
		c.Defaults.JSON = pending.Defaults.JSON
	case "defaults.pick":
		c.Defaults.Pick = pending.Defaults.Pick
	}
}

func (c *config) set(key, value string, in io.Reader) error {
	name, err := configKey(key)
	if err != nil {
		return err
	}
	if name != "" {
		if c.Providers == nil {
			c.Providers = make(map[string]providerConfig)
		}
		p := c.Providers[name]
		if strings.HasSuffix(key, ".model") {
			p.Model = &value
		} else {
			if value != "--stdin" {
				return errors.New("api_key accepts --stdin only; do not pass tokens as arguments")
			}
			data, err := io.ReadAll(io.LimitReader(in, 65537))
			if err != nil {
				return errors.New("could not read API key from stdin")
			}
			if len(data) > 65536 {
				return errors.New("API key input exceeds 64 KiB")
			}
			p.APIKey = strings.TrimSpace(string(data))
			if p.APIKey == "" {
				return errors.New("API key must not be empty; use config unset to remove it")
			}
		}
		c.Providers[name] = p
		return nil
	}
	d := &c.Defaults
	switch key {
	case "defaults.provider":
		d.Provider = &value
	case "defaults.threshold":
		n, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return errors.New("threshold must be a number between 0 and 1")
		}
		d.Threshold = &n
	case "defaults.min_score":
		n, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return errors.New("min_score must be a non-negative number")
		}
		d.MinScore = &n
	case "defaults.timeout":
		d.Timeout = &value
	case "defaults.json":
		if value != "true" && value != "false" {
			return errors.New("json must be true or false")
		}
		b := value == "true"
		d.JSON = &b
	case "defaults.pick":
		fields := splitList(value)
		d.Pick = &fields
	}
	return nil
}

func (c *config) unset(key string) (bool, error) {
	name, err := configKey(key)
	if err != nil {
		return false, err
	}
	if name != "" {
		p := c.Providers[name]
		changed := false
		if strings.HasSuffix(key, ".model") {
			changed = p.Model != nil
			p.Model = nil
		} else {
			changed = p.APIKey != ""
			p.APIKey = ""
		}
		if changed {
			c.Providers[name] = p
		}
		return changed, nil
	}
	d := &c.Defaults
	switch key {
	case "defaults.provider":
		changed := d.Provider != nil
		d.Provider = nil
		return changed, nil
	case "defaults.threshold":
		changed := d.Threshold != nil
		d.Threshold = nil
		return changed, nil
	case "defaults.min_score":
		changed := d.MinScore != nil
		d.MinScore = nil
		return changed, nil
	case "defaults.timeout":
		changed := d.Timeout != nil
		d.Timeout = nil
		return changed, nil
	case "defaults.json":
		changed := d.JSON != nil
		d.JSON = nil
		return changed, nil
	case "defaults.pick":
		changed := d.Pick != nil
		d.Pick = nil
		return changed, nil
	}
	return false, nil
}
