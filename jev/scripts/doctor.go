package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type doctorReport string

func (r doctorReport) Error() string { return string(r) }

func (a application) doctor(ctx context.Context, target string) error {
	var report strings.Builder
	ok, failed, skipped := 0, 0, 0
	record := func(status, name, detail, hint string) {
		fmt.Fprintf(&report, "  %-4s  %-10s %s\n", status, name, detail)
		if hint != "" {
			fmt.Fprintf(&report, "                    Fix: %s\n", hint)
		}
		switch status {
		case "OK":
			ok++
		case "FAIL":
			failed++
		case "SKIP":
			skipped++
		}
	}
	check := func(name, detail, hint string, err error) {
		if err != nil {
			record("FAIL", name, err.Error(), hint)
		} else {
			record("OK", name, detail, "")
		}
	}
	report.WriteString("Jev doctor\n\nLocal setup\n")
	info, binaryErr := regularTarget(target)
	if binaryErr == nil {
		binaryErr = identifyJev(target)
	}
	if binaryErr == nil && info.Mode().Perm()&0111 == 0 {
		binaryErr = errors.New("installed binary is not executable")
	}
	check("Binary", target, "run go run . install from the plugin scripts/ directory; use --dir for a custom target", binaryErr)

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
	check("PATH", resolved, "add the target directory to PATH before other Jev installations", pathErr)

	path, configErr := a.configPath()
	var c config
	if configErr == nil {
		c, configErr = readConfig(path)
	}
	configReadErr := configErr
	s := c.resolve(a.envToken)
	if configErr == nil {
		configErr = s.validate()
	}
	if configErr == nil && len(s.Pick) > 0 && !s.JSON {
		configErr = errors.New("pick requires JSON output")
	}
	configDetail := path + " (built-in defaults; file absent)"
	configHint := "check jev config and edit the config file to repair invalid values"
	if configErr == nil {
		if stat, err := os.Stat(path); err == nil {
			configDetail = path + " (valid, owner-only permissions)"
			if stat.Mode().Perm()&0077 != 0 {
				configErr = errors.New("config must have owner-only permissions (0600)")
				configHint = "set the config file permissions to 0600"
			}
		} else if !errors.Is(err, os.ErrNotExist) {
			configErr = err
		}
	}
	check("Config", configDetail, configHint, configErr)

	report.WriteString("\nAuthentication\n")
	keyHint := "use jev config set api_key --stdin or set TYPESAFE_API_KEY"
	if strings.TrimSpace(a.envToken) != "" {
		keyHint = "update or unset TYPESAFE_API_KEY; it overrides the config file api_key"
	}
	switch {
	case s.APIKey == "" && configReadErr != nil:
		record("SKIP", "API key", "config could not be read; key presence is unknown", "")
	case s.APIKey == "":
		record("FAIL", "API key", "missing", keyHint)
	case strings.ContainsAny(s.APIKey, "\r\n"):
		record("FAIL", "API key", "must be a single line", keyHint)
	case strings.TrimSpace(a.envToken) != "":
		record("OK", "API key", "TYPESAFE_API_KEY (overrides config api_key; value hidden)", "")
	default:
		record("OK", "API key", "config api_key (value hidden)", "")
	}
	switch {
	case configErr != nil:
		record("SKIP", "API", "resolve the config problem before checking authentication", "")
	case s.APIKey == "":
		record("SKIP", "API", "API key required; authentication not checked", "")
	default:
		detail, hint, err := checkDoctorAPI(ctx, s, a.transport, keyHint)
		check("API", detail, hint, err)
	}

	fmt.Fprintf(&report, "\nSummary: %d OK, %d FAIL, %d SKIP\n", ok, failed, skipped)
	report.WriteString("No local changes or inference requests were made.\n")
	// Local paths and diagnostics can also contain the configured token.
	output := report.String()
	for _, token := range []string{c.APIKey, a.envToken, s.APIKey} {
		if token != "" {
			output = strings.ReplaceAll(output, token, "[redacted]")
		}
	}
	if failed > 0 {
		return doctorReport(strings.TrimSuffix(output, "\n"))
	}
	_, err := io.WriteString(a.out, output)
	return err
}
