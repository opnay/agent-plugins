package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"time"
)

type defaults struct {
	Provider  *string   `toml:"provider,omitempty"`
	Threshold *float64  `toml:"threshold,omitempty"`
	MinScore  *float64  `toml:"min_score,omitempty"`
	Timeout   *string   `toml:"timeout,omitempty"`
	JSON      *bool     `toml:"json,omitempty"`
	Pick      *[]string `toml:"pick,omitempty"`
}

type providerConfig struct {
	Model  *string `toml:"model,omitempty"`
	APIKey string  `toml:"api_key,omitempty"`
}

type config struct {
	Defaults  defaults                  `toml:"defaults"`
	Providers map[string]providerConfig `toml:"providers,omitempty"`
}

// Request overrides are separate from the persisted schema.
type configFlags struct {
	Provider  *string
	Threshold *float64
	MinScore  *float64
	Model     *string
	Timeout   *string
	JSON      *bool
	Pick      *[]string
}

type settings struct {
	Provider  string
	APIKey    string
	Threshold *float64
	MinScore  *float64
	Model     string
	Timeout   string
	JSON      bool
	Pick      []string
}

func (c config) provider() string {
	if c.Defaults.Provider != nil {
		return *c.Defaults.Provider
	}
	return "jev"
}

func providerEnvironment(name string) string {
	switch name {
	case "jev":
		return "TYPESAFE_API_KEY"
	case "openrouter":
		return "OPENROUTER_API_KEY"
	case "openai":
		return "OPENAI_API_KEY"
	default:
		return ""
	}
}

func (c config) resolve(envToken string) settings {
	d := c.Defaults
	s := settings{Provider: c.provider(), Threshold: d.Threshold, MinScore: d.MinScore, Timeout: "30s", Pick: []string{}}
	p := c.Providers[s.Provider]
	s.APIKey = strings.TrimSpace(p.APIKey)
	if s.Provider == "jev" {
		s.Model = "jev-latest"
	}
	if p.Model != nil {
		s.Model = *p.Model
	}
	if d.Timeout != nil {
		s.Timeout = *d.Timeout
	}
	if d.JSON != nil {
		s.JSON = *d.JSON
	}
	if d.Pick != nil {
		s.Pick = *d.Pick
	}
	if token := strings.TrimSpace(envToken); token != "" {
		s.APIKey = token
	}
	return s
}

func (c *config) override(f configFlags) {
	d := &c.Defaults
	if f.Provider != nil {
		d.Provider = f.Provider
	}
	if f.Threshold != nil {
		d.Threshold = f.Threshold
	}
	if f.MinScore != nil {
		d.MinScore = f.MinScore
	}
	if f.Timeout != nil {
		d.Timeout = f.Timeout
	}
	if f.JSON != nil {
		d.JSON = f.JSON
	}
	if f.Pick != nil {
		d.Pick = f.Pick
	}
	if f.Model != nil {
		if c.Providers == nil {
			c.Providers = make(map[string]providerConfig)
		}
		p := c.Providers[c.provider()]
		p.Model = f.Model
		c.Providers[c.provider()] = p
	}
}

func (f configFlags) validate() error {
	c := config{}
	c.override(f)
	return c.validate()
}

func (c config) validate() error {
	if providerEnvironment(c.provider()) == "" {
		return errors.New("defaults.provider must be jev, openrouter, or openai")
	}
	s := c.resolve("")
	if err := s.validateCommon(); err != nil {
		return err
	}
	for name, p := range c.Providers {
		if providerEnvironment(name) == "" {
			return errors.New("provider name must be jev, openrouter, or openai")
		}
		if p.Model != nil && strings.TrimSpace(*p.Model) == "" {
			return fmt.Errorf("providers.%s.model must not be empty", name)
		}
		if strings.ContainsAny(p.APIKey, "\r\n") {
			return fmt.Errorf("providers.%s.api_key must be a single line", name)
		}
	}
	return nil
}

func (s settings) validate() error {
	if s.Provider != "jev" {
		return errors.New("selected provider is not implemented; currently only jev can execute")
	}
	if strings.TrimSpace(s.Model) == "" {
		return errors.New("model must not be empty")
	}
	return s.validateCommon()
}

func (s settings) validateCommon() error {
	if s.Threshold != nil && !probability(*s.Threshold) {
		return errors.New("threshold must be a finite number between 0 and 1")
	}
	if s.MinScore != nil && (!finite(*s.MinScore) || *s.MinScore < 0) {
		return errors.New("min_score must be a finite non-negative number")
	}
	if duration, err := time.ParseDuration(s.Timeout); err != nil || duration <= 0 {
		return errors.New("timeout must be a positive duration, such as 30s")
	}
	if strings.ContainsAny(s.APIKey, "\r\n") {
		return errors.New("API key must be a single line")
	}
	return validatePick(s.Pick)
}

func probability(value float64) bool { return finite(value) && value >= 0 && value <= 1 }
func finite(value float64) bool      { return !math.IsNaN(value) && !math.IsInf(value, 0) }
func splitList(value string) []string {
	if value == "" {
		return []string{}
	}
	items := strings.Split(value, ",")
	for i := range items {
		items[i] = strings.TrimSpace(items[i])
	}
	return items
}
