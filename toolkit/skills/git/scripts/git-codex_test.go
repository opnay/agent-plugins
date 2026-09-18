package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

var testBinary string

func TestMain(m *testing.M) {
	root, err := os.MkdirTemp("", "git-codex-test-bin-")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	testBinary = filepath.Join(root, "git-codex")
	build := exec.Command("go", "build", "-o", testBinary, "git-codex.go")
	if output, err := build.CombinedOutput(); err != nil {
		fmt.Fprintf(os.Stderr, "build: %v\n%s", err, output)
		os.RemoveAll(root)
		os.Exit(1)
	}
	code := m.Run()
	os.RemoveAll(root)
	os.Exit(code)
}

func testEnv() []string {
	var env []string
	for _, entry := range os.Environ() {
		if !strings.HasPrefix(entry, "GIT_") {
			env = append(env, entry)
		}
	}
	return append(env, "GIT_CONFIG_NOSYSTEM=1", "GIT_CONFIG_GLOBAL="+os.DevNull)
}

func execute(t *testing.T, dir, input, command string, args ...string) (string, string, int) {
	t.Helper()
	cmd := exec.Command(command, args...)
	cmd.Dir, cmd.Env = dir, testEnv()
	cmd.Stdin = strings.NewReader(input)
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	err := cmd.Run()
	if err == nil {
		return stdout.String(), stderr.String(), 0
	}
	if exit, ok := err.(*exec.ExitError); ok {
		return stdout.String(), stderr.String(), exit.ExitCode()
	}
	t.Fatalf("execute %s: %v", command, err)
	return "", "", -1
}

func git(t *testing.T, dir string, args ...string) string {
	t.Helper()
	out, stderr, code := execute(t, dir, "", "git", args...)
	if code != 0 {
		t.Fatalf("git %v: status=%d stdout=%q stderr=%q", args, code, out, stderr)
	}
	return strings.TrimSuffix(out, "\n")
}

func repository(t *testing.T) string {
	t.Helper()
	dir := filepath.Join(t.TempDir(), "repository with spaces")
	git(t, "", "init", "--quiet", "--initial-branch=main", dir)
	git(t, dir, "config", "user.name", "Git Codex Test")
	git(t, dir, "config", "user.email", "git-codex@example.invalid")
	git(t, dir, "config", "commit.gpgsign", "false")
	return dir
}

func writeFixture(t *testing.T, path, text string, mode os.FileMode) {
	t.Helper()
	if err := os.WriteFile(path, []byte(text), mode); err != nil {
		t.Fatal(err)
	}
}

func create(t *testing.T, dir, text string) (string, string) {
	t.Helper()
	id, stderr, code := execute(t, dir, text, testBinary, "message", "create", "--stdin")
	if code != 0 || stderr != "" {
		t.Fatalf("create: status=%d stdout=%q stderr=%q", code, id, stderr)
	}
	id = strings.TrimSuffix(id, "\n")
	if !strings.HasPrefix(id, "MSG-") || filepath.Base(id) != id || strings.Contains(id, "\n") {
		t.Fatalf("expected short message ID, got %q", id)
	}
	return id, filepath.Join(git(t, dir, "rev-parse", "--absolute-git-dir"), id)
}

func TestQuotedHeredoc(t *testing.T) {
	repo := repository(t)
	script := "\"$1\" message create --stdin <<'EOF'\ndocs: literal input\n\n- $TOKEN $(printf expanded)\nEOF\n"
	out, stderr, code := execute(t, repo, "", "bash", "-c", script, "heredoc-test", testBinary)
	if code != 0 || stderr != "" {
		t.Fatalf("heredoc: status=%d stdout=%q stderr=%q", code, out, stderr)
	}
	path := filepath.Join(repo, ".git", strings.TrimSpace(out))
	content, err := os.ReadFile(path)
	if err != nil || string(content) != "docs: literal input\n\n- $TOKEN $(printf expanded)\n" {
		t.Fatalf("literal input changed: %q, %v", content, err)
	}
	info, err := os.Stat(path)
	if err != nil || info.Mode().Perm() != 0600 {
		t.Fatalf("private message mode: %v, %v", info, err)
	}
}

func TestCreateRejectsInvalidInputAndCleans(t *testing.T) {
	inputs := map[string]string{
		"empty": "", "blank subject": "\nbody\n", "whitespace": " \n",
		"two lines without delimiter": "docs: subject\nbody",
		"body without delimiter":      "docs: subject\nbody\n",
		"literal escape":              "docs: subject\\nbody\n", "nul": "docs: bad\x00\n",
		"invalid utf8": "docs: bad\xff\n", "lone cr": "docs: bad\rtext\n",
		"trailing lone cr": "docs: bad\r",
	}
	for name, input := range inputs {
		t.Run(name, func(t *testing.T) {
			repo := repository(t)
			out, stderr, code := execute(t, repo, input, testBinary, "message", "create", "--stdin")
			if code != 1 || out != "" || !strings.Contains(stderr, "message_invalid") {
				t.Fatalf("invalid input: status=%d stdout=%q stderr=%q", code, out, stderr)
			}
			files, err := filepath.Glob(filepath.Join(repo, ".git", "MSG-*"))
			if err != nil || len(files) != 0 {
				t.Fatalf("failed create left messages: %v, %v", files, err)
			}
		})
	}
}

func TestCRLFMessage(t *testing.T) {
	repo := repository(t)
	text := "docs: CRLF message\r\n\r\n- body\r\n"
	_, path := create(t, repo, text)
	content, err := os.ReadFile(path)
	if err != nil || string(content) != text {
		t.Fatalf("CRLF content changed: %q, %v", content, err)
	}
}

func TestAllocationAndExplicitValidation(t *testing.T) {
	repo := repository(t)
	out, stderr, code := execute(t, repo, "ignored stdin", testBinary, "message", "create")
	if code != 0 || stderr != "" {
		t.Fatalf("allocate: %d %q %q", code, out, stderr)
	}
	id := strings.TrimSpace(out)
	path := filepath.Join(repo, ".git", id)
	if content, err := os.ReadFile(path); err != nil || len(content) != 0 {
		t.Fatalf("expected empty allocation: %q, %v", content, err)
	}
	_, _, code = execute(t, repo, "", testBinary, "message", "validate", id)
	if code != 1 {
		t.Fatalf("empty allocation validated: %d", code)
	}
	writeFixture(t, path, "docs: filled later\n\n- detail\n", 0600)
	for _, ref := range []string{id, filepath.Join(".git", id), path} {
		_, stderr, code = execute(t, repo, "", testBinary, "message", "validate", ref)
		if code != 0 {
			t.Fatalf("validate %q: %d %s", ref, code, stderr)
		}
	}
}

func TestGitDirectoryDiscovery(t *testing.T) {
	for _, layout := range []string{"nested", "linked worktree", "separate git directory", "bare"} {
		t.Run(layout, func(t *testing.T) {
			repo := repository(t)
			dir := repo
			switch layout {
			case "linked worktree":
				git(t, repo, "commit", "--quiet", "--allow-empty", "-m", "initial")
				dir = filepath.Join(t.TempDir(), "linked")
				git(t, repo, "worktree", "add", "--quiet", "--detach", dir)
			case "separate git directory":
				dir = filepath.Join(t.TempDir(), "work")
				git(t, "", "init", "--quiet", "--separate-git-dir", filepath.Join(t.TempDir(), "metadata"), dir)
			case "bare":
				dir = t.TempDir()
				git(t, dir, "init", "--quiet", "--bare")
			}
			root := dir
			if layout != "bare" {
				dir = filepath.Join(dir, "deep", "inside")
				if err := os.MkdirAll(dir, 0755); err != nil {
					t.Fatal(err)
				}
			}
			id, path := create(t, dir, "docs: from nested directory\n")
			if _, err := os.Stat(path); err != nil {
				t.Fatal(err)
			}
			_, stderr, code := execute(t, root, "", testBinary, "message", "validate", id)
			if code != 0 {
				t.Fatalf("resolve same ID from root: %d %s", code, stderr)
			}
			if layout == "linked worktree" {
				_, _, code = execute(t, repo, "", testBinary, "message", "validate", path)
				if code == 0 {
					t.Fatal("accepted another worktree's message")
				}
			}
		})
	}
}

func TestNoRepositoryDoesNotAllocate(t *testing.T) {
	dir := t.TempDir()
	out, stderr, code := execute(t, dir, "docs: outside\n", testBinary, "message", "create", "--stdin")
	if code != 1 || out != "" || !strings.Contains(stderr, "repository_unavailable") {
		t.Fatalf("outside repository: %d %q %q", code, out, stderr)
	}
	entries, err := os.ReadDir(dir)
	if err != nil || len(entries) != 0 {
		t.Fatalf("unexpected allocation: %v %v", entries, err)
	}
}

func TestUnsafeGitDirectoryDoesNotAllocate(t *testing.T) {
	repo := repository(t)
	root := filepath.Join(repo, ".git")
	if err := os.Chmod(root, 0770); err != nil {
		t.Fatal(err)
	}
	out, stderr, code := execute(t, repo, "docs: input\n", testBinary, "message", "create", "--stdin")
	if code != 1 || out != "" || !strings.Contains(stderr, "unsafe Git directory") {
		t.Fatalf("unsafe directory: %d %q %q", code, out, stderr)
	}
	files, err := filepath.Glob(filepath.Join(root, "MSG-*"))
	if err != nil || len(files) != 0 {
		t.Fatalf("unexpected allocation: %v %v", files, err)
	}
}

type inputFunc func([]byte) (int, error)

func (input inputFunc) Read(p []byte) (int, error) { return input(p) }

func TestCreateReadFailureCleanup(t *testing.T) {
	for _, replace := range []bool{false, true} {
		t.Run(fmt.Sprintf("replace=%t", replace), func(t *testing.T) {
			repo := repository(t)
			root := filepath.Join(repo, ".git")
			t.Setenv("GIT_DIR", root)
			t.Setenv("GIT_CONFIG_GLOBAL", os.DevNull)
			var allocated string
			input := inputFunc(func(p []byte) (int, error) {
				files, err := filepath.Glob(filepath.Join(root, "MSG-*"))
				if err != nil || len(files) != 1 {
					t.Fatalf("allocation: %v %v", files, err)
				}
				allocated = files[0]
				if replace {
					if err := os.Rename(allocated, allocated+".saved"); err != nil {
						t.Fatal(err)
					}
					writeFixture(t, allocated, "docs: replacement\n", 0600)
				}
				return copy(p, "partial"), fmt.Errorf("input interrupted")
			})
			id, err := createMessage(input)
			if id != "" || err == nil || !strings.Contains(err.Error(), "input interrupted") {
				t.Fatalf("read failure: %q %v", id, err)
			}
			if replace {
				if !strings.Contains(err.Error(), "cleanup failed, preserved") {
					t.Fatalf("missing preserved-path diagnostic: %v", err)
				}
				content, readErr := os.ReadFile(allocated)
				if readErr != nil || string(content) != "docs: replacement\n" {
					t.Fatalf("replacement was changed: %q %v", content, readErr)
				}
			} else if _, statErr := os.Stat(allocated); !os.IsNotExist(statErr) {
				t.Fatalf("failed input retained its own file: %v", statErr)
			}
		})
	}
}

func TestCommitValidationFailurePreservesInput(t *testing.T) {
	repo := repository(t)
	path := filepath.Join(repo, ".git", "MSG-invalid")
	writeFixture(t, path, "invalid\nbody", 0600)
	_, stderr, code := execute(t, repo, "", testBinary, "commit", "MSG-invalid")
	if code != 1 || !strings.Contains(stderr, "reason=validation_failed attempted=false") {
		t.Fatalf("invalid commit: %d %q", code, stderr)
	}
	// The caller owns pre-commit cleanup; commit only deletes after success.
	if content, err := os.ReadFile(path); err != nil || string(content) != "invalid\nbody" {
		t.Fatalf("commit validation changed its input: %q %v", content, err)
	}
}

func TestUnsafeMessagePaths(t *testing.T) {
	repo := repository(t)
	_, valid := create(t, repo, "docs: valid\n")
	outside := filepath.Join(t.TempDir(), "MSG-outside")
	writeFixture(t, outside, "docs: outside\n", 0600)
	symlink := filepath.Join(repo, ".git", "MSG-symlink")
	if err := os.Symlink(valid, symlink); err != nil {
		t.Fatal(err)
	}
	public := filepath.Join(repo, ".git", "MSG-public")
	writeFixture(t, public, "docs: public\n", 0644)
	wrongName := filepath.Join(repo, ".git", "message.txt")
	writeFixture(t, wrongName, "docs: wrong filename\n", 0600)
	if err := os.Chmod(public, 0644); err != nil {
		t.Fatal(err)
	}
	other := repository(t)
	_, otherPath := create(t, other, "docs: other repository\n")
	for _, path := range []string{outside, symlink, public, otherPath, wrongName} {
		_, stderr, code := execute(t, repo, "", testBinary, "message", "validate", path)
		if code == 0 {
			t.Fatalf("accepted unsafe path %s: %s", path, stderr)
		}
		if _, err := os.Lstat(path); err != nil {
			t.Fatalf("validation changed input %s: %v", path, err)
		}
	}
}

func TestCommitMessageLifecycle(t *testing.T) {
	for _, reject := range []bool{false, true} {
		t.Run(fmt.Sprintf("reject=%t", reject), func(t *testing.T) {
			repo := repository(t)
			writeFixture(t, filepath.Join(repo, "owned.txt"), "owned\n", 0644)
			writeFixture(t, filepath.Join(repo, "unrelated.txt"), "unrelated\n", 0644)
			git(t, repo, "add", "--", "owned.txt")
			if reject {
				writeFixture(t, filepath.Join(repo, ".git", "hooks", "pre-commit"), "#!/bin/sh\nexit 1\n", 0755)
			}
			message := "feat: scoped change\n\n- verified message input\n"
			id, path := create(t, repo, message)
			out, stderr, code := execute(t, repo, "", testBinary, "commit", id)
			if reject {
				if code != 1 || !strings.Contains(stderr, "reason=commit_failed attempted=true") {
					t.Fatalf("rejected commit: %d %q %q", code, out, stderr)
				}
				if _, err := os.Stat(path); err != nil {
					t.Fatalf("failed commit lost message: %v", err)
				}
				return
			}
			if code != 0 || !strings.Contains(out, "cleanup=deleted") {
				t.Fatalf("commit: %d %q %q", code, out, stderr)
			}
			if _, err := os.Stat(path); !os.IsNotExist(err) {
				t.Fatalf("successful commit kept message: %v", err)
			}
			if stored := git(t, repo, "log", "-1", "--format=%B"); stored != message {
				t.Fatalf("stored message: %q", stored)
			}
			if files := git(t, repo, "ls-tree", "--name-only", "HEAD"); files != "owned.txt" {
				t.Fatalf("commit changed staged scope: %q", files)
			}
		})
	}
}

func TestLinkedWorktreeCommitFromNestedDirectory(t *testing.T) {
	repo := repository(t)
	git(t, repo, "commit", "--quiet", "--allow-empty", "-m", "initial")
	linked := filepath.Join(t.TempDir(), "linked")
	git(t, repo, "worktree", "add", "--quiet", "--detach", linked)
	nested := filepath.Join(linked, "nested")
	if err := os.Mkdir(nested, 0755); err != nil {
		t.Fatal(err)
	}
	writeFixture(t, filepath.Join(nested, "change.txt"), "change\n", 0644)
	git(t, nested, "add", "--", "change.txt")
	id, path := create(t, linked, "feat: linked change\n")
	out, stderr, code := execute(t, nested, "", testBinary, "commit", id)
	if code != 0 || !strings.Contains(out, "cleanup=deleted") {
		t.Fatalf("linked commit: %d %q %q", code, out, stderr)
	}
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Fatalf("linked message was not cleaned: %v", err)
	}
	if subject := git(t, repo, "log", "-1", "--format=%s"); subject != "initial" {
		t.Fatalf("main worktree HEAD changed: %q", subject)
	}
}
