package main

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

//go:embed config.template.toml
var configTemplate string

func templateConfig() (config, error) { return decodeConfig([]byte(configTemplate)) }

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
	return decodeConfig(data)
}

func decodeConfig(data []byte) (config, error) {
	var fields map[string]any
	if err := toml.Unmarshal(data, &fields); err != nil {
		return config{}, errors.New("invalid TOML; edit the config file to repair it")
	}
	for key := range fields {
		if key != "defaults" && key != "providers" {
			return config{}, errors.New("unsupported config layout; move common keys to [defaults] and model/api_key to [providers.jev]; do not mix old and new layouts")
		}
	}
	var c config
	if err := toml.NewDecoder(bytes.NewReader(data)).DisallowUnknownFields().Decode(&c); err != nil {
		return config{}, errors.New("invalid TOML, field type, or unknown key in config; edit the file to repair it")
	}
	if err := c.validate(); err != nil {
		return config{}, err
	}
	return c, nil
}

func encodeInitialConfig(c config) ([]byte, error) {
	data, err := toml.Marshal(c)
	if err != nil {
		return nil, errors.New("could not encode config")
	}
	// Keep template comments with their table so optional settings remain usable.
	comments := make(map[string][]string)
	section := ""
	for _, line := range strings.Split(configTemplate, "\n") {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "[") {
			section = trimmed
		}
		if strings.HasPrefix(trimmed, "#") {
			comments[section] = append(comments[section], line)
		}
	}
	var lines []string
	section = ""
	for _, line := range strings.Split(string(data), "\n") {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "[") {
			lines = append(lines, comments[section]...)
			delete(comments, section)
			section = trimmed
		}
		lines = append(lines, line)
	}
	lines = append(lines, comments[section]...)
	return []byte(strings.Join(lines, "\n") + "\n"), nil
}

func initConfig(path string) error {
	if _, err := templateConfig(); err != nil {
		return err
	}
	return storeConfig(path, []byte(configTemplate), true)
}

func createConfig(path string, c config) error {
	if err := c.validate(); err != nil {
		return err
	}
	data, err := encodeInitialConfig(c)
	if err != nil {
		return err
	}
	return storeConfig(path, data, true)
}

func writeConfig(path string, c config) error {
	if err := c.validate(); err != nil {
		return err
	}
	exists, err := configFileStatus(path)
	if err != nil {
		return err
	}
	var data []byte
	if !exists {
		data, err = encodeInitialConfig(c)
	} else {
		data, err = toml.Marshal(c)
	}
	if err != nil {
		return errors.New("could not encode config")
	}
	return storeConfig(path, data, !exists)
}

func storeConfig(path string, data []byte, createOnly bool) error {
	exists, err := configFileStatus(path)
	if err != nil {
		return err
	}
	if createOnly && exists {
		return errors.New("config already exists; leaving it unchanged")
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
	if createOnly {
		// Linking a fully written temp file publishes it atomically without replacing a race winner.
		if err := os.Link(f.Name(), path); err != nil {
			return errors.New("could not create config without replacing an existing file")
		}
		return nil
	}
	if err := os.Rename(f.Name(), path); err != nil {
		return fmt.Errorf("replace config: %w", err)
	}
	return nil
}
