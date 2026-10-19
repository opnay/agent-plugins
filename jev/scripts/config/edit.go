package config

import (
	"errors"
	"io"
	"strconv"
	"strings"
)

// Return the provider name, or an empty string for a common defaults key.
func Key(key string) (string, error) {
	parts := strings.Split(key, ".")
	if len(parts) == 2 && parts[0] == "defaults" {
		switch parts[1] {
		case "provider", "threshold", "min_score", "timeout", "json", "pick":
			return "", nil
		}
	}
	if len(parts) == 3 && parts[0] == "providers" && ProviderEnvironment(parts[1]) != "" && (parts[2] == "model" || parts[2] == "api_key") {
		return parts[1], nil
	}
	return "", errors.New("use defaults.provider/threshold/min_score/timeout/json/pick or providers.<jev|openrouter|openai>.model/api_key")
}

func (c *Config) ApplySet(key string, pending Config) {
	name, _ := Key(key)
	if name != "" {
		if c.Providers == nil {
			c.Providers = make(map[string]ProviderConfig)
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

func (c *Config) Set(key, value string, in io.Reader) error {
	name, err := Key(key)
	if err != nil {
		return err
	}
	if name != "" {
		if c.Providers == nil {
			c.Providers = make(map[string]ProviderConfig)
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
		fields := SplitList(value)
		d.Pick = &fields
	}
	return nil
}

func (c *Config) Unset(key string) (bool, error) {
	name, err := Key(key)
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
