package cli

import (
	"time"

	cfg "opnay/jev/config"
	"opnay/jev/providers"
	jevprovider "opnay/jev/providers/jev"
)

func (a App) newProvider(s cfg.Settings) providers.Adapter {
	timeout, _ := time.ParseDuration(s.Timeout)
	return jevprovider.New(s.APIKey, timeout, a.transport)
}
