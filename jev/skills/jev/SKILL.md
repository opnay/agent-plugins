---
name: jev
description: Evaluate a question with explicit choices through TypeSafe's Jev CLI, manage its settings, or install, uninstall, and diagnose the executable. Use for Jev or TypeSafe requests, not general questions or whole-codebase review.
---

# Jev

Use `jev` for one Choice question. The model is hosted by TypeSafe; questions and choices leave the machine. Send only context authorized for that external call.

## Form the input

Put the question and its necessary context together in `--if`. Supply 2–255 distinct, nonempty labels through `--conditions`, separated by commas. Labels cannot contain commas.

```sh
jev --if "The setting is recommended but optional. Is it mandatory?" \
  --conditions supported,contradicted,undetermined
```

Keep relationships needed for the decision. Per-file answers do not establish a conclusion about a whole system. Clarify ambiguous choice meanings in the input; do not invent missing evidence.

## Select output and use the result

Default output is the API's highest-probability answer and a newline. Preserve its tie choice and scores.

- `--json`: an object containing `answer`, `probabilities`, `confidence`, and the actual response `model`.
- `--json --pick answer,probabilities`: only named top-level fields. A single field still produces an object.
- `--pick ''`: clear a configured field restriction. Nonempty pick requires JSON mode. Unknown, empty, or duplicate fields fail.
- `--threshold 0.8`: require `probabilities[answer] >= 0.8`. Without a threshold, no filter applies.

API `confidence` is a separate distribution-derived value, not the threshold measure. Scores do not guarantee correctness.

Exit `0` means consume stdout. Exit `2` means threshold not met: stdout is empty and stderr identifies the choice, probability, and threshold. Report the shortfall without silently lowering the threshold. Exit `1` means an execution error, not a classification; explain the reported input, setup, API, or output failure.

JSON mode does not change failure streams. The CLI does not retry automatically. A successful result needs no second API call or connection probe. Address installation or authentication when execution reports a setup problem; do not make doctor a prerequisite for ordinary evaluation.

## Configure defaults

Settings live in `~/.agents/jev.toml`. Explicit CLI flags override file values, which override built-in defaults. Nonempty `TYPESAFE_API_KEY` overrides the file's `api_key`.

```sh
jev config
jev config path
jev config set threshold 0.8
jev config unset threshold
jev config set api_key --stdin
```

Config commands are offline and need no token. `config` shows effective settings with the whole token masked. `path` works even with invalid TOML. `set` stores a value; `unset` restores fallback behavior. Successful mutations print nothing, and reads do not create files.

Keys and defaults: no `api_key`, `threshold` off, `model = "jev-latest"`, `timeout = "30s"`, `json = false`, and no `pick` restriction. Set pick with `jev config set pick answer,model`; TOML stores an array. Enable JSON before evaluating with a nonempty pick.

Override a call with `--model`, `--timeout`, `--json[=false]`, `--pick`, or `--threshold`. Explicit `--threshold 0` overrides a saved positive threshold, and `--json=false` overrides saved JSON mode.

Change settings only when requested. Read tokens from an authorized local source through stdin, not conversation text or command arguments. The file is plaintext with mode 0600, not encrypted storage. Writes preserve other values but not comments or formatting. Symlink config files are refused; invalid TOML needs manual repair.

## Install, remove, or diagnose

For initial installation, run from the plugin root:

```sh
cd scripts
go run . install
```

Go 1.23 or newer is required for this source bootstrap. After a source update, use `go run . install --force` in that directory. Use `--dir <directory>` for another location. No shell wrapper is needed, and plugin installation alone does not put the executable on PATH.

Use `jev install [--force]`, `jev uninstall`, or `jev doctor` for requested maintenance. All accept `--dir <directory>`; the default target is `~/.local/bin/jev`. An explicit directory must not be empty.

- `install` copies the running binary without downloading or building a newer version. Identical installations succeed unchanged. Replacing a different regular file requires `--force`; symlinks and non-regular targets are refused.
- `uninstall` removes only an identifiable Jev binary. A missing target succeeds. Config, tokens, the installation directory, other files, and shell settings remain. Restore the executable with `go run . install` from the source directory.
- `doctor` reads installation identity, PATH resolution, config validity and owner-only permissions, and token presence. It makes no API call or repairs; a configured token does not establish valid authentication.

Maintenance is offline and never edits shell configuration. Success uses a short stdout report and exit `0`; failure uses empty stdout, stderr, and exit `1`. Doctor groups its checks in one report. Run it for requested diagnosis or a reported setup problem, not automatically after installation or before inference.
