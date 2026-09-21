// git-codex provides the bounded message-file and commit lifecycle for toolkit:git.
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"unicode/utf8"
)

const version = "0.1.0"
const prefix = "MSG-"
const aliasKey = "alias.codex"
const aliasValue = `!f() { "$HOME/.local/bin/git-codex" "$@"; }; f`

func fail(code, message string) { fmt.Fprintf(os.Stderr, "git-codex: %s: %s\n", code, message) }
func usage() {
	fmt.Println("Usage:\n  git codex --help\n  git codex --version\n  git codex install [--force]\n  git codex uninstall [--force]\n  git codex doctor\n  git codex message create [--stdin]\n  git codex message validate <message>\n  git codex commit <message>\n\nMessages are stored in the current Git directory. Use the returned MSG-... ID\nfrom any directory in the same repository/worktree. --stdin reads through EOF\nand validates before returning the ID; without it, create allocates an empty file.")
}

func gitRoot() (string, error) {
	b, err := exec.Command("git", "rev-parse", "--absolute-git-dir").Output()
	if err != nil {
		return "", fmt.Errorf("repository_unavailable: %w", err)
	}
	root, err := filepath.EvalSymlinks(strings.TrimSuffix(string(b), "\n"))
	if err != nil {
		return "", err
	}
	if err := safeDir(root); err != nil {
		return "", fmt.Errorf("unsafe Git directory: %w", err)
	}
	return root, nil
}

func messagePath(message string) (string, error) {
	if filepath.Base(message) == message && strings.HasPrefix(message, prefix) {
		root, err := gitRoot()
		if err != nil {
			return "", err
		}
		return filepath.Join(root, message), nil
	}
	return filepath.Abs(message)
}

func safeFile(path string) (syscall.Stat_t, error) {
	var zero syscall.Stat_t
	if !filepath.IsAbs(path) {
		return zero, fmt.Errorf("path_invalid")
	}
	name := filepath.Base(path)
	if !strings.HasPrefix(name, prefix) || len(name) == len(prefix) {
		return zero, fmt.Errorf("path_invalid")
	}
	root, err := gitRoot()
	if err != nil {
		return zero, err
	}
	parent, err := filepath.EvalSymlinks(filepath.Dir(path))
	if err != nil || parent != root {
		return zero, fmt.Errorf("path_invalid")
	}
	info, err := os.Lstat(path)
	if err != nil || !info.Mode().IsRegular() || info.Mode().Perm()&0077 != 0 || info.Mode().Perm()&0400 == 0 {
		return zero, fmt.Errorf("file_unsafe")
	}
	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok || stat.Uid != uint32(os.Geteuid()) {
		return zero, fmt.Errorf("file_unsafe")
	}
	return *stat, nil
}
func validate(path string) (syscall.Stat_t, error) {
	stat, err := safeFile(path)
	if err != nil {
		return stat, err
	}
	f, err := os.OpenFile(path, os.O_RDONLY|syscall.O_NOFOLLOW, 0)
	if err != nil {
		return stat, fmt.Errorf("file_unsafe")
	}
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		return stat, fmt.Errorf("file_unsafe")
	}
	opened, ok := info.Sys().(*syscall.Stat_t)
	if !ok || !same(stat, *opened) {
		return stat, fmt.Errorf("file_unsafe")
	}
	b, err := io.ReadAll(f)
	if err != nil || !utf8.Valid(b) || bytes.IndexByte(b, 0) >= 0 || bytes.Contains(b, []byte(`\n`)) {
		return stat, fmt.Errorf("message_invalid")
	}
	normalized := bytes.ReplaceAll(b, []byte("\r\n"), []byte("\n"))
	if bytes.ContainsRune(normalized, '\r') {
		return stat, fmt.Errorf("message_invalid")
	}
	lines := bytes.Split(normalized, []byte("\n"))
	if len(bytes.TrimSpace(lines[0])) == 0 {
		return stat, fmt.Errorf("message_invalid")
	}
	if len(lines) > 1 && len(bytes.TrimSpace(bytes.Join(lines[1:], nil))) > 0 && len(lines[1]) != 0 {
		return stat, fmt.Errorf("message_invalid")
	}
	return stat, nil
}
func same(a, b syscall.Stat_t) bool { return a.Dev == b.Dev && a.Ino == b.Ino }

func removeMessage(path string, expected syscall.Stat_t) error {
	current, err := safeFile(path)
	if err != nil || !same(expected, current) {
		return fmt.Errorf("message identity or safety changed")
	}
	return os.Remove(path)
}

func createMessage(input io.Reader) (string, error) {
	root, err := gitRoot()
	if err != nil {
		return "", err
	}
	f, err := os.CreateTemp(root, prefix+"*")
	if err != nil {
		return "", err
	}
	info, err := f.Stat()
	if err != nil {
		f.Close()
		return "", fmt.Errorf("cannot record identity; preserved %s: %w", f.Name(), err)
	}
	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		f.Close()
		return "", fmt.Errorf("cannot record identity; preserved %s", f.Name())
	}
	err = f.Chmod(0600)
	if err == nil && input != nil {
		_, err = io.Copy(f, input)
	}
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	if err == nil {
		var current syscall.Stat_t
		current, err = safeFile(f.Name())
		if err == nil && !same(*stat, current) {
			err = fmt.Errorf("allocated file identity changed")
		}
	}
	if err == nil && input != nil {
		var validated syscall.Stat_t
		validated, err = validate(f.Name())
		if err == nil && !same(*stat, validated) {
			err = fmt.Errorf("allocated file identity changed")
		}
	}
	if err != nil {
		if cleanupErr := removeMessage(f.Name(), *stat); cleanupErr != nil {
			return "", fmt.Errorf("%w; cleanup failed, preserved %s: %v", err, f.Name(), cleanupErr)
		}
		return "", err
	}
	return filepath.Base(f.Name()), nil
}
func result(reason string, attempted bool) {
	fmt.Fprintf(os.Stderr, "git-codex: commit_result: reason=%s attempted=%t\n", reason, attempted)
}
func run(name string, args ...string) error {
	c := exec.Command(name, args...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}
func head() (string, error) {
	b, e := exec.Command("git", "rev-parse", "--verify", "HEAD").Output()
	return strings.TrimSpace(string(b)), e
}
func safeDir(path string) error {
	info, e := os.Lstat(path)
	if e != nil || !info.IsDir() || info.Mode()&os.ModeSymlink != 0 || info.Mode().Perm()&0022 != 0 {
		return fmt.Errorf("unsafe directory")
	}
	s, ok := info.Sys().(*syscall.Stat_t)
	if !ok || s.Uid != uint32(os.Geteuid()) {
		return fmt.Errorf("unsafe directory")
	}
	return nil
}
func installPaths() (string, string, string, error) {
	self, e := os.Executable()
	if e != nil {
		return "", "", "", e
	}
	home := os.Getenv("HOME")
	if home == "" {
		return "", "", "", fmt.Errorf("HOME is required")
	}
	local := filepath.Join(home, ".local")
	bin := filepath.Join(local, "bin")
	target := filepath.Join(bin, "git-codex")
	return self, local, target, nil
}

func installedBinary() (string, error) {
	self, local, target, e := installPaths()
	if e != nil {
		return "", e
	}
	if e = safeDir(local); e != nil {
		return "", e
	}
	if e = safeDir(filepath.Dir(target)); e != nil {
		return "", e
	}
	a, e := os.ReadFile(self)
	if e != nil {
		return "", e
	}
	b, e := os.ReadFile(target)
	if e != nil || !bytes.Equal(a, b) {
		return "", fmt.Errorf("installed binary does not match")
	}
	return target, nil
}

func installBinary(force bool) (string, error) {
	self, local, target, e := installPaths()
	if e != nil {
		return "", e
	}
	bin := filepath.Dir(target)
	for _, d := range []string{local, bin} {
		if _, e = os.Lstat(d); os.IsNotExist(e) {
			if e = os.Mkdir(d, 0755); e != nil {
				return "", e
			}
		}
		if e = safeDir(d); e != nil {
			return "", e
		}
	}
	if info, e := os.Lstat(target); e == nil {
		if !info.Mode().IsRegular() || info.Mode()&os.ModeSymlink != 0 {
			return "", fmt.Errorf("refusing to replace a symlink or directory")
		}
		a, _ := os.ReadFile(self)
		b, _ := os.ReadFile(target)
		if bytes.Equal(a, b) && info.Mode().Perm() == 0755 {
			return target, nil
		}
		if !force {
			return "", fmt.Errorf("existing file is preserved; rerun with --force")
		}
	}
	src, e := os.Open(self)
	if e != nil {
		return "", e
	}
	defer src.Close()
	tmp, e := os.CreateTemp(bin, ".git-codex.*")
	if e != nil {
		return "", e
	}
	if _, e = io.Copy(tmp, src); e != nil {
		tmp.Close()
		os.Remove(tmp.Name())
		return "", e
	}
	if e = tmp.Chmod(0755); e == nil {
		e = tmp.Close()
	}
	if e == nil {
		e = os.Rename(tmp.Name(), target)
	}
	if e != nil {
		return "", e
	}
	return target, nil
}

func aliasValues() ([]string, error) {
	b, e := exec.Command("git", "config", "--global", "--get-all", aliasKey).Output()
	if e != nil {
		if exit, ok := e.(*exec.ExitError); ok && exit.ExitCode() == 1 {
			return nil, nil
		}
		return nil, e
	}
	return strings.FieldsFunc(strings.TrimSpace(string(b)), func(r rune) bool { return r == '\n' || r == '\r' }), nil
}

func expectedAlias(values []string) bool {
	return len(values) > 0 && all(values, func(value string) bool { return value == aliasValue })
}

func all(values []string, predicate func(string) bool) bool {
	for _, value := range values {
		if !predicate(value) {
			return false
		}
	}
	return true
}

func configureAlias(force bool) error {
	values, e := aliasValues()
	if e != nil {
		return e
	}
	if len(values) == 0 {
		return exec.Command("git", "config", "--global", aliasKey, aliasValue).Run()
	}
	if expectedAlias(values) {
		return nil
	}
	if !force {
		return fmt.Errorf("existing alias is preserved; rerun with --force")
	}
	return exec.Command("git", "config", "--global", "--replace-all", aliasKey, aliasValue).Run()
}

func uninstall(force bool) error {
	values, e := aliasValues()
	if e != nil || len(values) == 0 {
		return e
	}
	if !expectedAlias(values) && !force {
		return fmt.Errorf("existing alias is preserved; rerun with --force")
	}
	return exec.Command("git", "config", "--global", "--unset-all", aliasKey).Run()
}

func doctor() error {
	values, e := aliasValues()
	if e != nil {
		return e
	}
	if !expectedAlias(values) {
		return fmt.Errorf("managed alias is missing or differs")
	}
	target, e := installedBinary()
	if e != nil {
		return e
	}
	fmt.Printf("alias=%s\nbinary=%s\nversion=toolkit-git-codex %s\n", aliasValue, target, version)
	return nil
}
func main() {
	a := os.Args[1:]
	if len(a) == 1 && (a[0] == "--help" || a[0] == "-h") {
		usage()
		return
	}
	if len(a) == 1 && a[0] == "--version" {
		fmt.Printf("toolkit-git-codex %s\n", version)
		return
	}
	if len(a) >= 1 && a[0] == "install" {
		if len(a) > 2 {
			usage()
			os.Exit(1)
		}
		force := len(a) == 2 && a[1] == "--force"
		if len(a) == 2 && !force {
			usage()
			os.Exit(1)
		}
		if _, e := installBinary(force); e != nil {
			fail("install_failed", e.Error())
			os.Exit(1)
		}
		if e := configureAlias(force); e != nil {
			fail("install_failed", e.Error())
			os.Exit(1)
		}
		if e := doctor(); e != nil {
			fail("install_failed", e.Error())
			os.Exit(1)
		}
		return
	}
	if len(a) >= 1 && a[0] == "uninstall" {
		if len(a) > 2 {
			usage()
			os.Exit(1)
		}
		force := len(a) == 2 && a[1] == "--force"
		if len(a) == 2 && !force {
			usage()
			os.Exit(1)
		}
		if e := uninstall(force); e != nil {
			fail("uninstall_failed", e.Error())
			os.Exit(1)
		}
		return
	}
	if len(a) == 1 && a[0] == "doctor" {
		if e := doctor(); e != nil {
			fail("doctor_failed", e.Error())
			os.Exit(1)
		}
		return
	}
	if len(a) >= 2 && a[0] == "message" && a[1] == "create" {
		var input io.Reader
		if len(a) == 3 && a[2] == "--stdin" {
			input = os.Stdin
		} else if len(a) != 2 {
			usage()
			os.Exit(1)
		}
		message, err := createMessage(input)
		if err != nil {
			fail("create_failed", err.Error())
			os.Exit(1)
		}
		fmt.Println(message)
		return
	}
	if len(a) == 3 && a[0] == "message" && a[1] == "validate" {
		path, e := messagePath(a[2])
		if e == nil {
			_, e = validate(path)
		}
		if e != nil {
			fail(e.Error(), "message file failed validation")
			os.Exit(1)
		}
		return
	}
	if len(a) == 2 && a[0] == "commit" {
		path, e := messagePath(a[1])
		var stat syscall.Stat_t
		if e == nil {
			stat, e = validate(path)
		}
		if e != nil {
			fail(e.Error(), "message file failed validation")
			result("validation_failed", false)
			os.Exit(1)
		}
		before, e := head()
		unborn := e != nil
		if unborn {
			if e := run("git", "rev-parse", "--is-inside-work-tree"); e != nil {
				result("head_unavailable", false)
				os.Exit(1)
			}
		}
		err := run("git", "commit", "-F", path)
		after, afterErr := head()
		if err == nil && afterErr == nil && (unborn || after != before) {
			if e := removeMessage(path, stat); e != nil {
				fmt.Printf("commit=%s cleanup=failed message_file=%s\n", after, path)
				os.Exit(2)
			}
			fmt.Printf("commit=%s cleanup=deleted\n", after)
			return
		}
		if err != nil && ((!unborn && afterErr == nil && after == before) || (unborn && afterErr != nil)) {
			result("commit_failed", true)
			os.Exit(1)
		}
		result("outcome_unknown", true)
		os.Exit(3)
	}
	usage()
	os.Exit(1)
}
