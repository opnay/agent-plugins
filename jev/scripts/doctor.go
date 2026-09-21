package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func (a application) doctor(target string) error {
	var lines []string
	failed := false
	record := func(name, detail string, err error) {
		if err != nil {
			failed = true
			lines = append(lines, name+": ERROR: "+err.Error())
		} else {
			lines = append(lines, name+": "+detail)
		}
	}
	info, binaryErr := regularTarget(target)
	if binaryErr == nil {
		binaryErr = identifyJev(target)
	}
	if binaryErr == nil && info.Mode().Perm()&0111 == 0 {
		binaryErr = errors.New("installed binary is not executable")
	}
	record("binary", target, binaryErr)

	lookup := a.lookupPath
	if lookup == nil {
		lookup = exec.LookPath
	}
	resolved, pathErr := lookup("jev")
	if pathErr == nil {
		resolvedInfo, err := os.Stat(resolved)
		if err != nil {
			pathErr = err
		} else if info == nil || !os.SameFile(info, resolvedInfo) {
			pathErr = fmt.Errorf("PATH selects %s instead of %s", resolved, target)
		}
	}
	record("path", resolved, pathErr)

	path, configErr := a.configPath()
	var c config
	if configErr == nil {
		c, configErr = readConfig(path)
	}
	s := c.resolve(a.envToken)
	if configErr == nil {
		configErr = s.validate()
	}
	if configErr == nil && len(s.Pick) > 0 && !s.JSON {
		configErr = errors.New("pick requires JSON output")
	}
	if configErr == nil {
		if stat, err := os.Stat(path); err == nil {
			if stat.Mode().Perm()&0077 != 0 {
				configErr = errors.New("config must have owner-only permissions (0600)")
			}
		} else if !errors.Is(err, os.ErrNotExist) {
			configErr = err
		}
	}
	record("config", path+" (file values or built-in defaults)", configErr)
	var authErr error
	if s.APIKey == "" {
		authErr = errors.New("API key missing; set TYPESAFE_API_KEY or use jev config set api_key --stdin")
	}
	record("api_key", "configured (not authenticated)", authErr)
	lines = append(lines, "Offline checks only; no API request or automatic repair was performed.")
	report := strings.Join(lines, "\n")
	if failed {
		return errors.New(report)
	}
	_, err := io.WriteString(a.out, report+"\n")
	return err
}
