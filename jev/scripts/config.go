package main

import (
	"errors"
	"math"
	"strings"
	"time"
)

// Pointers distinguish omission from explicit zero, false, or an empty list.
type config struct {
	APIKey    string    `toml:"api_key,omitempty"`
	Threshold *float64  `toml:"threshold,omitempty"`
	Model     *string   `toml:"model,omitempty"`
	Timeout   *string   `toml:"timeout,omitempty"`
	JSON      *bool     `toml:"json,omitempty"`
	Pick      *[]string `toml:"pick,omitempty"`
}

type settings struct {
	APIKey    string   `toml:"api_key"`
	Threshold *float64 `toml:"threshold,omitempty"`
	Model     string   `toml:"model"`
	Timeout   string   `toml:"timeout"`
	JSON      bool     `toml:"json"`
	Pick      []string `toml:"pick"`
}

func (c config) resolve(envToken string) settings {
	s := settings{APIKey: strings.TrimSpace(c.APIKey), Threshold: c.Threshold,
		Model: "jev-latest", Timeout: "30s", Pick: []string{}}
	if c.Model != nil {
		s.Model = *c.Model
	}
	if c.Timeout != nil {
		s.Timeout = *c.Timeout
	}
	if c.JSON != nil {
		s.JSON = *c.JSON
	}
	if c.Pick != nil {
		s.Pick = *c.Pick
	}
	if token := strings.TrimSpace(envToken); token != "" {
		s.APIKey = token
	}
	return s
}

func (c *config) override(flags config) {
	if flags.Threshold != nil {
		c.Threshold = flags.Threshold
	}
	if flags.Model != nil {
		c.Model = flags.Model
	}
	if flags.Timeout != nil {
		c.Timeout = flags.Timeout
	}
	if flags.JSON != nil {
		c.JSON = flags.JSON
	}
	if flags.Pick != nil {
		c.Pick = flags.Pick
	}
}

func (s settings) validate() error {
	if s.Threshold != nil && !probability(*s.Threshold) {
		return errors.New("threshold must be a finite number between 0 and 1")
	}
	if strings.TrimSpace(s.Model) == "" {
		return errors.New("model must not be empty")
	}
	if duration, err := time.ParseDuration(s.Timeout); err != nil || duration <= 0 {
		return errors.New("timeout must be a positive duration, such as 30s")
	}
	if strings.ContainsAny(s.APIKey, "\r\n") {
		return errors.New("API key must be a single line")
	}
	return validatePick(s.Pick)
}

func probability(value float64) bool {
	return !math.IsNaN(value) && !math.IsInf(value, 0) && value >= 0 && value <= 1
}

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
