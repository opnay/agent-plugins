package maintenance

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMaintenanceTargets(t *testing.T) {
	for _, command := range []string{"install", "uninstall", "doctor"} {
		// Parse only: a regression must never touch the real default installation.
		for _, extra := range [][]string{{"--dir", ""}, {"--dir="}} {
			if _, err := parseLifecycle(command, extra); err == nil {
				t.Fatalf("%s accepted an explicit empty directory", command)
			}
		}
	}
	for _, command := range []string{"uninstall", "doctor"} {
		if _, err := parseLifecycle(command, []string{"--force"}); err == nil {
			t.Fatalf("%s accepted --force", command)
		}
	}
	o, err := parseLifecycle("install", nil)
	home, homeErr := os.UserHomeDir()
	if err != nil || homeErr != nil || o.target != filepath.Join(home, ".local", "bin", "jev") {
		t.Fatal("wrong default target")
	}
}
