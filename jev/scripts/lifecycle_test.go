package main

import (
	"bytes"
	"context"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func buildTestBinary(t *testing.T) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "jev")
	if output, err := exec.Command("go", "build", "-trimpath", "-o", path, ".").CombinedOutput(); err != nil {
		t.Fatalf("build test binary: %v: %s", err, output)
	}
	return path
}

func TestGoRunBootstrap(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "bin with spaces")
	cmd := exec.Command("go", "run", ".", "install", "--dir", dir)
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("go run install: %v: %s", err, output)
	}
	target := filepath.Join(dir, "jev")
	if err := identifyJev(target); err != nil {
		t.Fatal(err)
	}
	if output, err := exec.Command(target, "--help").CombinedOutput(); err != nil || !strings.Contains(string(output), "Usage:") {
		t.Fatalf("installed binary: %v: %s", err, output)
	}
}

func TestMaintenanceLifecycle(t *testing.T) {
	a, out, errOut := harness(t)
	a.executable = buildTestBinary(t)
	a.envToken = "" // Neither install nor uninstall requires authentication.
	dir := filepath.Join(t.TempDir(), "bin with spaces")
	target := filepath.Join(dir, "jev")
	if err := writeConfig(a.path, config{APIKey: "preserved-token"}); err != nil {
		t.Fatal(err)
	}
	before, err := os.ReadFile(a.path)
	if err != nil {
		t.Fatal(err)
	}
	run := func(args ...string) {
		t.Helper()
		out.Reset()
		errOut.Reset()
		if code := a.run(context.Background(), args); code != 0 || errOut.Len() != 0 {
			t.Fatalf("%v: code=%d stdout=%q stderr=%q", args, code, out, errOut)
		}
	}
	run("install", "--dir", dir)
	info, err := os.Stat(target)
	if err != nil || info.Mode().Perm() != 0755 {
		t.Fatalf("installed permissions: %v %v", info, err)
	}
	run("install", "--dir", dir)
	if !strings.Contains(out.String(), "Already installed") {
		t.Fatal("reinstall was not idempotent")
	}
	run("install", "--dir", filepath.Dir(a.executable)) // Already running at target.
	marker := filepath.Join(dir, "keep.txt")
	if err := os.WriteFile(marker, []byte("keep"), 0600); err != nil {
		t.Fatal(err)
	}
	// Exercise main dispatch and removal of the currently running installed binary.
	if output, err := exec.Command(target, "uninstall", "--dir", dir).CombinedOutput(); err != nil || !strings.Contains(string(output), "Uninstalled") {
		t.Fatalf("self uninstall: %v %s", err, output)
	}
	if _, err := os.Stat(target); !errors.Is(err, os.ErrNotExist) {
		t.Fatal("binary was not removed")
	}
	run("uninstall", "--dir", dir)
	if !strings.Contains(out.String(), "Not installed") {
		t.Fatal("missing uninstall must succeed")
	}
	after, err := os.ReadFile(a.path)
	if err != nil || !bytes.Equal(before, after) {
		t.Fatal("uninstall changed config")
	}
	if data, err := os.ReadFile(marker); err != nil || string(data) != "keep" {
		t.Fatal("uninstall changed adjacent files")
	}
}

func TestMaintenanceProtectsTargets(t *testing.T) {
	a, out, errOut := harness(t)
	a.executable = buildTestBinary(t)
	for _, kind := range []string{"foreign", "symlink", "directory"} {
		t.Run(kind, func(t *testing.T) {
			dir := t.TempDir()
			target := filepath.Join(dir, "jev")
			switch kind {
			case "foreign":
				if err := os.WriteFile(target, []byte("unrelated file"), 0755); err != nil {
					t.Fatal(err)
				}
			case "symlink":
				if err := os.Symlink(a.executable, target); err != nil {
					t.Fatal(err)
				}
			case "directory":
				if err := os.Mkdir(target, 0700); err != nil {
					t.Fatal(err)
				}
			}
			for _, command := range []string{"install", "uninstall"} {
				out.Reset()
				errOut.Reset()
				if code := a.run(context.Background(), []string{command, "--dir", dir}); code != 1 || out.Len() > 0 {
					t.Fatalf("%s replaced %s: code=%d %s", command, kind, code, errOut)
				}
				if _, err := os.Lstat(target); err != nil {
					t.Fatal("target removed")
				}
			}
			out.Reset()
			errOut.Reset()
			code := a.run(context.Background(), []string{"install", "--force", "--dir", dir})
			if kind == "foreign" {
				if code != 0 {
					t.Fatalf("explicit regular file replacement failed: %s", errOut)
				}
				if err := identifyJev(target); err != nil {
					t.Fatal(err)
				}
			} else if code != 1 || out.Len() > 0 {
				t.Fatalf("force replaced non-regular target: %d %s", code, errOut)
			}
		})
	}
}

func TestMaintenanceOptionsAndHelp(t *testing.T) {
	for _, command := range []string{"install", "uninstall", "doctor"} {
		a, out, errOut := harness(t)
		a.path = t.TempDir() // Help does not load config.
		if code := a.run(context.Background(), []string{command, "--help"}); code != 0 || !strings.Contains(out.String(), "Usage:\n  jev "+command) {
			t.Fatalf("help %s: %d %s", command, code, errOut)
		}
		for _, extra := range [][]string{{"--unknown"}, {"unexpected"}, {"--dir"}} {
			out.Reset()
			errOut.Reset()
			if code := a.run(context.Background(), append([]string{command}, extra...)); code != 1 || out.Len() > 0 {
				t.Fatalf("invalid options accepted: %s %v", command, extra)
			}
		}
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
