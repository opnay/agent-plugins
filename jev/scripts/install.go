package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func installExecutable(source, target string, force bool) (string, error) {
	sourceInfo, err := os.Stat(source)
	if err != nil {
		return "", fmt.Errorf("read running executable: %w", err)
	}
	if !sourceInfo.Mode().IsRegular() {
		return "", errors.New("running executable must be a regular file")
	}
	targetInfo, err := regularTarget(target)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return "", err
	}
	if targetInfo != nil {
		if os.SameFile(sourceInfo, targetInfo) {
			return "Already installed " + target, nil
		}
		if targetInfo.Size() == sourceInfo.Size() && targetInfo.Mode().Perm() == 0755 {
			a, err := os.ReadFile(source)
			if err != nil {
				return "", err
			}
			b, err := os.ReadFile(target)
			if err != nil {
				return "", err
			}
			if bytes.Equal(a, b) {
				return "Already installed " + target, nil
			}
		}
		if !force {
			return "", errors.New("target exists; use --force to replace this file")
		}
	}
	if err := copyExecutable(source, target, force); err != nil {
		return "", err
	}
	return "Installed " + target + "\nAdd its directory to PATH if needed; shell configuration was not changed.", nil
}

func copyExecutable(source, target string, force bool) error {
	src, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("open running executable: %w", err)
	}
	defer src.Close()
	dir := filepath.Dir(target)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("create installation directory: %w", err)
	}
	tmp, err := os.CreateTemp(dir, ".jev-install.*")
	if err != nil {
		return fmt.Errorf("create temporary executable: %w", err)
	}
	defer os.Remove(tmp.Name())
	defer tmp.Close()
	if _, err := io.Copy(tmp, src); err != nil {
		return fmt.Errorf("copy executable: %w", err)
	}
	if err := tmp.Chmod(0755); err != nil {
		return fmt.Errorf("set executable permissions: %w", err)
	}
	if err := tmp.Sync(); err != nil {
		return fmt.Errorf("sync executable: %w", err)
	}
	if err := tmp.Close(); err != nil {
		return fmt.Errorf("close executable: %w", err)
	}
	if !force {
		// Publish only if the target is still absent; never clobber a concurrent install.
		if err := os.Link(tmp.Name(), target); err != nil {
			return fmt.Errorf("publish executable without replacing an existing target: %w", err)
		}
		return nil
	}
	if _, err := regularTarget(target); err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	if err := os.Rename(tmp.Name(), target); err != nil {
		return fmt.Errorf("replace executable: %w", err)
	}
	return nil
}
