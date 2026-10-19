package maintenance

import (
	"fmt"
	"os"

	"opnay/jev/installation"
)

func (a Command) install(o lifecycleOptions) (string, error) {
	source := a.Executable
	if source == "" {
		var err error
		source, err = os.Executable()
		if err != nil {
			return "", fmt.Errorf("locate running executable: %w", err)
		}
	}
	return installation.Install(source, o.target, o.force)
}
