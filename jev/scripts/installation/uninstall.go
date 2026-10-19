package installation

import (
	"errors"
	"fmt"
	"os"
)

func Uninstall(target string) (string, error) {
	if _, err := RegularTarget(target); errors.Is(err, os.ErrNotExist) {
		return "Not installed " + target, nil
	} else if err != nil {
		return "", err
	}
	if err := Identify(target); err != nil {
		return "", err
	}
	if err := os.Remove(target); err != nil {
		return "", fmt.Errorf("remove installed Jev binary: %w", err)
	}
	return "Uninstalled " + target + "\nConfig, tokens, installation directory, and shell configuration were preserved.", nil
}
