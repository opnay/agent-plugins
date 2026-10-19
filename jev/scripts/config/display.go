package config

import "strings"

func (c Config) Display(env map[string]string) Config {
	s := c.Resolve("")
	shown := Config{Defaults: Defaults{Provider: &s.Provider, Timeout: &s.Timeout, JSON: &s.JSON, Pick: &s.Pick, Threshold: s.Threshold, MinScore: s.MinScore}, Providers: map[string]ProviderConfig{}}
	for _, name := range []string{"jev", "openrouter", "openai"} {
		p, exists := c.Providers[name]
		if name == "jev" && p.Model == nil {
			model := "jev-latest"
			p.Model = &model
			exists = true
		}
		token := strings.TrimSpace(env[ProviderEnvironment(name)])
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
