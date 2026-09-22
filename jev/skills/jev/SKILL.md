---
name: jev
description: Evaluate a yes/no proposition, discrete choice, or ordered scale through TypeSafe's Jev CLI; manage its settings or executable when requested. Use for Jev or TypeSafe requests, not general questions or whole-codebase review.
---

# Jev

Use the hosted Jev CLI for one structured judgment. Questions, context, options, and levels leave the machine for TypeSafe. Send only content authorized for that external call.

## Choose the primitive

- `noul`: whether one proposition is true. The answer is the probability of yes from 0 to 1.
- `choice`: which member of a discrete set applies. The answer is the API's highest-probability supplied label; preserve its selection when probabilities tie.
- `score`: where the input lies on ordered descriptive levels. The answer is a probability-weighted number from 0 through `level count - 1` and may fall between levels.

Use one primitive for one question. Keep relationships needed for the judgment in `--if`; separate per-file calls do not establish a system-wide conclusion. Do not report a judgment based on facts absent from the supplied input.

```sh
jev noul --if "Is the statement mandatory? Context: ..."

jev choice --if "How does the evidence relate to the claim? Context: ..." \
  --conditions supported,conflicting,undetermined

jev score --if "How severe is the issue? Context: ..." \
  --level "No functional impact" \
  --level "Degraded with a workaround" \
  --level "Blocked with no workaround"
```

Choice accepts 2–255 distinct comma-separated labels; labels cannot contain commas. Labels are free strings, but prefer `snake_case` when code will consume them. Score accepts 2–10 distinct `--level` values in ascending order; levels are natural-language criteria and may contain commas. The legacy form `jev --if ... --conditions ...` remains a Choice call.

## Use output and result criteria

Default stdout is the primitive's answer and one newline. Use JSON only when the task needs supporting fields:

- Noul: `answer`, `model`.
- Choice: `answer`, `probabilities`, `confidence`, `model`.
- Score: `answer`, `legend`, `probabilities`, `confidence`, `model`.

`--json --pick answer,probabilities` keeps named fields in an object. A field must exist for the selected primitive. `--pick ''` clears a saved restriction. CLI, config, and JSON-owned keys use `snake_case`.

For Choice, `--threshold 0.8` requires the selected label's probability to be at least 0.8. For Noul, it requires the probability of yes to be at least 0.8. For Score, use `--min-score 1.5`; its maximum is the highest level index. API confidence is separate and is not used for either criterion.

Exit `0` means consume stdout. Exit `2` means the result criterion was not met: stdout is empty and stderr reports the result and boundary. Do not silently lower the boundary. Exit `1` is an execution error, not a judgment. A successful result needs no second API call or connection probe.

## Configure defaults

Settings live in `~/.agents/jev.toml`. Explicit flags override file values, which override built-ins. Nonempty `TYPESAFE_API_KEY` overrides file `api_key`.

```sh
jev config
jev config path
jev config set threshold 0.8
jev config set min_score 1.5
jev config set pick answer,model
jev config unset threshold
jev config set api_key --stdin
```

Config commands are offline. `config` masks the entire token; `path` does not require a valid file. Successful mutations print nothing. Reads do not create files. Built-ins are no token, no result criteria, `model = "jev-latest"`, `timeout = "30s"`, `json = false`, and no pick restriction.

Saved `threshold` applies to Choice and Noul; saved `min_score` applies to Score. A saved pick incompatible with the selected primitive fails before the API call. Override calls with `--model`, `--timeout`, `--json[=false]`, `--pick`, and the primitive's result criterion.

Change settings only when requested. Read tokens from an authorized local source through stdin, never conversation text or command arguments. The file is plaintext mode 0600. Writes preserve other values but not comments or formatting. Symlink config files are refused; invalid TOML needs manual repair.

## Maintain the executable

For initial installation, run from the plugin root:

```sh
cd scripts
go run . install
```

After a source update, use `go run . install --force`. Use `--dir <directory>` for another location. Plugin installation alone does not put the executable on PATH.

Run `jev install [--force]`, `jev uninstall`, or `jev doctor` only for requested maintenance or a reported setup problem. All accept `--dir`; the default target is `~/.local/bin/jev`.

- `install` copies the running binary. Replacing a different regular file requires `--force`; symlinks and non-regular targets are refused.
- `uninstall` removes only an identifiable Jev binary; a missing target succeeds. It preserves config, tokens, directories, adjacent files, and shell settings.
- `doctor` reads binary identity, PATH resolution, general config validity and permissions, and token presence. It neither calls the API nor repairs state.

Maintenance uses exit `0` with a short stdout result or exit `1` with stderr. Do not run doctor automatically after installation or before evaluation.
