package main

import (
	"errors"
	"fmt"
	"os"
)

func uninstallExecutable(target string) (string, error) {
	if _, err := regularTarget(target); errors.Is(err, os.ErrNotExist) {
		return "Not installed " + target, nil
	} else if err != nil {
		return "", err
	}
	if err := identifyJev(target); err != nil {
		return "", err
	}
	if err := os.Remove(target); err != nil {
		return "", fmt.Errorf("remove installed Jev binary: %w", err)
	}
	return "Uninstalled " + target + "\nConfig, tokens, installation directory, and shell configuration were preserved.", nil
}
