---
name: jev
description: Evaluate one or many yes/no propositions, discrete choices, or ordered scales through TypeSafe's Jev CLI; manage its settings or executable when requested. Use for Jev or TypeSafe requests, not general questions or whole-codebase review.
---

# Jev

Use the hosted Jev CLI for structured judgments. Questions, shared state, criteria, and levels leave the machine for TypeSafe. Send only content authorized for that external call.

## Choose Single Or Batch Evaluation

Use a single command for one judgment:

- `noul`: whether one proposition is true; answer is the probability of yes from 0 to 1.
- `choice`: which supplied label applies; answer is the highest-probability label selected by the API.
- `score`: where the input lies on ordered descriptive levels; answer is a probability-weighted value from 0 through `level count - 1`.

```sh
jev noul --if "Is the statement mandatory? Context: ..."

jev choice --if "How does the evidence relate to the claim? Context: ..." \
  --conditions supported,conflicting,undetermined

jev score --if "How severe is the issue? Context: ..." \
  --level "No functional impact" \
  --level "Degraded with a workaround" \
  --level "Blocked with no workaround"
```

Use `jev batch` when two or more independent questions share one exact state. Do not combine questions that require different states merely to reduce calls.

```sh
jev batch --state-file target.md --input questions.jsonl
```

Each nonblank JSONL row needs a unique `id`, a `type`, and a `question`:

```jsonl
{"id":"S001","type":"choice","question":"Which action applies?","conditions":["proceed","block","unexpected_action"]}
{"id":"S002","type":"noul","question":"Is this behavior explicitly authorized?"}
{"id":"S003","type":"score","question":"How severe is the risk?","levels":["none","recoverable","destructive"]}
```

Batch rules:

- The state file must contain nonempty UTF-8 text. Its exact content is shared by every question.
- Keys use the exact lowercase names shown above. Noncanonical or unknown keys fail validation.
- IDs and questions must be nonempty strings, and IDs must be unique.
- IDs, conditions, and levels must not have surrounding whitespace; the CLI preserves validated strings exactly.
- `type` is `noul`, `choice`, or `score`.
- Choice uses 2-255 distinct nonempty `conditions`. JSON labels may contain commas.
- Score uses 2-10 distinct nonempty `levels` in ascending order.
- The CLI validates the complete state and JSONL before one API request.
- Success emits one full JSON object per input row, in input order.
- Batch does not chunk, retry, resume, or return partial results.
- Batch does not accept or apply `threshold`, `min_score`, `json`, or `pick`.

The legacy form `jev --if ... --conditions ...` remains a single Choice call.

## Use Single Results

Single commands print the primitive answer and one newline by default. Use JSON when supporting fields are needed:

- Noul: `answer`, `model`.
- Choice: `answer`, `probabilities`, `confidence`, `model`.
- Score: `answer`, `legend`, `probabilities`, `confidence`, `model`.

`--json --pick answer,probabilities` keeps named top-level fields. A field must exist for the selected primitive. `--pick ''` clears a saved restriction. CLI, config, and JSON-owned keys use `snake_case`.

Choice accepts 2-255 comma-separated labels; labels cannot contain commas in the single-command form. Prefer `snake_case` when code consumes them. Score accepts 2-10 repeated `--level` values in ascending order; levels may contain commas.

For Choice, `--threshold 0.8` checks the selected label's probability. For Noul, it checks the probability of yes. For Score, use `--min-score 1.5`. API confidence is separate and is not a result criterion.

Single-command exits:

- `0`: consume stdout.
- `2`: result criterion not met; stdout is empty and stderr reports the result and boundary.
- `1`: input, configuration, API, response, or output error.

Batch uses exit `0` only for a complete JSONL result and exit `1` for any failure. It never uses exit `2`. A successful single or batch result needs no repeated API call or connection probe.

## Configure Defaults

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

Config commands are offline. Successful mutations print nothing. Reads do not create files. Built-ins are no token, no result criteria, `model = "jev-latest"`, `timeout = "30s"`, `json = false`, and no pick restriction.

Batch applies only saved `api_key`, `model`, and `timeout`; it ignores saved single-result criteria and output selection.

Change settings only when requested. Read tokens from an authorized local source through stdin, never conversation text or command arguments. The config is plaintext mode 0600. Symlink config files are refused; invalid TOML needs manual repair.

## Maintain The Executable

For initial installation, run from the plugin root:

```sh
cd scripts
go run . install
```

After a source update, use `go run . install --force`. Use `--dir <directory>` for another location. Plugin installation alone does not put the executable on PATH.

Run `jev install`, `jev uninstall`, or `jev doctor` only for requested maintenance or a reported setup problem.

- `install` copies the running binary. Replacing a different regular file requires `--force`; symlinks and non-regular targets are refused.
- `uninstall` removes only an identifiable Jev binary and preserves config, tokens, directories, adjacent files, and shell settings.
- `doctor` reads binary identity, PATH resolution, config validity and permissions, and token presence. It neither calls the API nor repairs state.

Maintenance uses exit `0` with a short stdout result or exit `1` with stderr. Do not run doctor automatically after installation or before evaluation.
