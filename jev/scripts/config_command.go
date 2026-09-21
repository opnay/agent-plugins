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
		_, err := io.WriteString(a.out, usage)
		return err
	}
	path, err := a.configPath()
	if err != nil {
		return err
	}
	if len(args) == 1 && args[0] == "path" {
		_, err := fmt.Fprintln(a.out, path)
		return err
	}
	c, err := readConfig(path)
	if err != nil {
		return err
	}
	if len(args) == 0 {
		s := c.resolve(a.envToken)
		if err := s.validate(); err != nil {
			return err
		}
		if s.APIKey != "" {
			s.APIKey = "********"
		}
		data, err := toml.Marshal(s)
		if err != nil {
			return errors.New("could not encode config display")
		}
		_, err = a.out.Write(data)
		return err
	}
	if len(args) == 3 && args[0] == "set" {
		if err := c.set(args[1], args[2], a.in); err != nil {
			return err
		}
		return writeConfig(path, c)
	}
	if len(args) == 2 && args[0] == "unset" {
		changed, err := c.unset(args[1])
		if err != nil || !changed {
			return err
		}
		return writeConfig(path, c)
	}
	return errors.New("invalid config command; see jev --help")
}

func (c *config) set(key, value string, in io.Reader) error {
	switch key {
	case "api_key":
		if value != "--stdin" {
			return errors.New("set api_key accepts --stdin only; do not pass tokens as arguments")
		}
		data, err := io.ReadAll(io.LimitReader(in, 65537))
		if err != nil {
			return errors.New("could not read API key from stdin")
		}
		if len(data) > 65536 {
			return errors.New("API key input exceeds 64 KiB")
		}
		c.APIKey = strings.TrimSpace(string(data))
		if c.APIKey == "" {
			return errors.New("API key must not be empty; use config unset api_key to remove it")
		}
	case "threshold":
		n, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return errors.New("threshold must be a number between 0 and 1")
		}
		c.Threshold = &n
	case "json":
		if value != "true" && value != "false" {
			return errors.New("json must be true or false")
		}
		b := value == "true"
		c.JSON = &b
	case "model":
		c.Model = &value
	case "timeout":
		c.Timeout = &value
	case "pick":
		fields := splitList(value)
		c.Pick = &fields
	default:
		return errors.New("unknown config key; use api_key, threshold, model, timeout, json, or pick")
	}
	return nil
}

func (c *config) unset(key string) (bool, error) {
	switch key {
	case "api_key":
		changed := c.APIKey != ""
		c.APIKey = ""
		return changed, nil
	case "threshold":
		changed := c.Threshold != nil
		c.Threshold = nil
		return changed, nil
	case "model":
		changed := c.Model != nil
		c.Model = nil
		return changed, nil
	case "timeout":
		changed := c.Timeout != nil
		c.Timeout = nil
		return changed, nil
	case "json":
		changed := c.JSON != nil
		c.JSON = nil
		return changed, nil
	case "pick":
		changed := c.Pick != nil
		c.Pick = nil
		return changed, nil
	default:
		return false, errors.New("unknown config key; use api_key, threshold, model, timeout, json, or pick")
	}
}
