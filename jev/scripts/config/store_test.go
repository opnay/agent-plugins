package config

import (
	"os"
	"path/filepath"
	"sync"
	"testing"
)

func TestConfigRefusesSymlinksAndDirectories(t *testing.T) {
	dir := t.TempDir()
	target := filepath.Join(dir, "target")
	if err := os.WriteFile(target, []byte(`api_key = "unchanged"`), 0600); err != nil {
		t.Fatal(err)
	}
	link := filepath.Join(dir, "link")
	if err := os.Symlink(target, link); err != nil {
		t.Fatal(err)
	}
	for _, path := range []string{link, dir} {
		if _, err := Read(path); err == nil {
			t.Fatal("read non-regular config")
		}
		if err := Write(path, Config{}); err == nil {
			t.Fatal("replaced non-regular config")
		}
	}
	data, err := os.ReadFile(target)
	if err != nil || string(data) != `api_key = "unchanged"` {
		t.Fatal("symlink target changed")
	}
}

func TestConcurrentConfigInit(t *testing.T) {
	path := filepath.Join(t.TempDir(), "jev.toml")
	var wg sync.WaitGroup
	results := make(chan error, 8)
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); results <- Init(path) }()
	}
	wg.Wait()
	close(results)
	successes := 0
	for err := range results {
		if err == nil {
			successes++
		}
	}
	if successes != 1 {
		t.Fatalf("successful initializers=%d", successes)
	}
	data, err := os.ReadFile(path)
	if err != nil || string(data) != configTemplate {
		t.Fatal("incomplete template")
	}
}
