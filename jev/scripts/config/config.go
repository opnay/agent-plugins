package config

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"time"

	"opnay/jev/decision"
)

type Defaults struct {
	Provider  *string   `toml:"provider,omitempty"`
	Threshold *float64  `toml:"threshold,omitempty"`
	MinScore  *float64  `toml:"min_score,omitempty"`
	Timeout   *string   `toml:"timeout,omitempty"`
	JSON      *bool     `toml:"json,omitempty"`
	Pick      *[]string `toml:"pick,omitempty"`
}

type ProviderConfig struct {
	Model  *string `toml:"model,omitempty"`
	APIKey string  `toml:"api_key,omitempty"`
}

type Config struct {
	Defaults  Defaults                  `toml:"defaults"`
	Providers map[string]ProviderConfig `toml:"providers,omitempty"`
}

// Request overrides are separate from the persisted schema.
type Flags struct {
	Provider  *string
	Threshold *float64
	MinScore  *float64
	Model     *string
	Timeout   *string
	JSON      *bool
	Pick      *[]string
}

type Settings struct {
	Provider  string
	APIKey    string
	Threshold *float64
	MinScore  *float64
	Model     string
	Timeout   string
	JSON      bool
	Pick      []string
}

func (f Flags) Validate() error {
	c := Config{}
	c.Override(f)
	return c.Validate()
}

func (c Config) Validate() error {
	if ProviderEnvironment(c.Provider()) == "" {
		return errors.New("defaults.provider must be jev, openrouter, or openai")
	}
	s := c.Resolve("")
	if err := s.validateCommon(); err != nil {
		return err
	}
	for name, p := range c.Providers {
		if ProviderEnvironment(name) == "" {
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

func (s Settings) Validate() error {
	if s.Provider != "jev" {
		return errors.New("selected provider is not implemented; currently only jev can execute")
	}
	if strings.TrimSpace(s.Model) == "" {
		return errors.New("model must not be empty")
	}
	return s.validateCommon()
}

func (s Settings) validateCommon() error {
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
	return decision.ValidatePick(s.Pick)
}

func probability(value float64) bool { return finite(value) && value >= 0 && value <= 1 }
func finite(value float64) bool      { return !math.IsNaN(value) && !math.IsInf(value, 0) }
func SplitList(value string) []string {
	if value == "" {
		return []string{}
	}
	items := strings.Split(value, ",")
	for i := range items {
		items[i] = strings.TrimSpace(items[i])
	}
	return items
}
