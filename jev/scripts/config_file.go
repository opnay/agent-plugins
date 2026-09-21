package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

func configFileStatus(path string) (bool, error) {
	info, err := os.Lstat(path)
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("access config: %w", err)
	}
	if !info.Mode().IsRegular() {
		return false, errors.New("config must be a regular file, not a symlink or directory")
	}
	return true, nil
}

func readConfig(path string) (config, error) {
	var c config
	exists, err := configFileStatus(path)
	if err != nil || !exists {
		return c, err
	}
	f, err := os.Open(path)
	if err != nil {
		return c, fmt.Errorf("open config: %w", err)
	}
	defer f.Close()
	data, err := io.ReadAll(io.LimitReader(f, 1024*1024+1))
	if err != nil {
		return c, fmt.Errorf("read config: %w", err)
	}
	if len(data) > 1024*1024 {
		return c, errors.New("config exceeds 1 MiB")
	}
	if err := toml.NewDecoder(bytes.NewReader(data)).DisallowUnknownFields().Decode(&c); err != nil {
		// TOML decoder diagnostics can contain the secret-bearing source line.
		return config{}, errors.New("invalid TOML, field type, or unknown key in config; edit the file to repair it")
	}
	return c, nil
}

func writeConfig(path string, c config) error {
	if err := c.resolve("").validate(); err != nil {
		return err
	}
	if _, err := configFileStatus(path); err != nil {
		return err
	}
	data, err := toml.Marshal(c)
	if err != nil {
		return errors.New("could not encode config")
	}
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("create config directory: %w", err)
	}
	f, err := os.CreateTemp(dir, ".jev-*.toml")
	if err != nil {
		return fmt.Errorf("create temporary config: %w", err)
	}
	defer os.Remove(f.Name())
	defer f.Close()
	// CreateTemp creates 0600 files, including the first write of the token.
	if _, err := f.Write(data); err != nil {
		return fmt.Errorf("write config: %w", err)
	}
	if err := f.Sync(); err != nil {
		return fmt.Errorf("sync config: %w", err)
	}
	if err := f.Close(); err != nil {
		return fmt.Errorf("close config: %w", err)
	}
	if _, err := configFileStatus(path); err != nil {
		return err
	}
	if err := os.Rename(f.Name(), path); err != nil {
		return fmt.Errorf("replace config: %w", err)
	}
	return nil
}
