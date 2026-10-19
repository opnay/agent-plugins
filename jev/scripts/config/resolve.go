package config

import "strings"

func (c Config) Provider() string {
	if c.Defaults.Provider != nil {
		return *c.Defaults.Provider
	}
	return "jev"
}

func ProviderEnvironment(name string) string {
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

func (c Config) Resolve(envToken string) Settings {
	d := c.Defaults
	s := Settings{Provider: c.Provider(), Threshold: d.Threshold, MinScore: d.MinScore, Timeout: "30s", Pick: []string{}}
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

func (c *Config) Override(f Flags) {
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
			c.Providers = make(map[string]ProviderConfig)
		}
		p := c.Providers[c.Provider()]
		p.Model = f.Model
		c.Providers[c.Provider()] = p
	}
}
