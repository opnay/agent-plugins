---
name: jev
description: Explain TypeSafe Jev and use its hosted Noul, Choice, Score, and batch CLI; interpret results or manage Jev settings and the executable when requested. Use for explicit Jev or TypeSafe questions and CLI operations, not routine work routing.
---

# Jev

Own Jev information, primitive selection, CLI inputs, result interpretation, and requested configuration or executable maintenance. `$jev:jev-use` owns advisory judgments during routine work. This skill runs without sibling context or development specs.

The CLI sends questions, context, criteria, and batch state to TypeSafe's hosted API. Include only content authorized for that external call. Preserve related evidence together; do not present a judgment as evidence for facts absent from the input.

## Choose A Primitive

| Command | Input | Answer |
|---|---|---|
| `noul` | One yes/no proposition | Probability of truth, 0–1 |
| `choice` | Discrete alternatives | Highest-probability label |
| `score` | Ordered descriptive levels | Weighted score, 0 through level count minus 1 |

Put the question and material context in `--if`:

```sh
jev noul --if "Is this requirement mandatory? Context: ..."
jev choice --if "How does the evidence relate to the claim? Context: ..." \
  --conditions supported,conflicting,undetermined
jev score --if "How severe is this issue? Context: ..." \
  --level "No functional impact" \
  --level "Degraded with a workaround" \
  --level "Blocked with no workaround"
```

Choice takes 2–255 unique comma-separated labels; a label cannot contain commas. Score takes 2–10 unique `--level` values in ascending order; level text may contain commas. CLI, config, and JSON-owned keys use `snake_case`. Choice labels are free text; prefer `snake_case` when code consumes them. Legacy `jev --if ... --conditions ...` is a Choice call.

## Batch Shared Context

For two or more independent questions about one nonempty UTF-8 state, use:

```sh
jev batch --state-file target.md --input questions.jsonl
```

```jsonl
{"id":"S001","type":"choice","question":"Which action applies?","conditions":["proceed","block","unexpected_action"]}
{"id":"S002","type":"noul","question":"Is the proposition true?"}
{"id":"S003","type":"score","question":"How severe is the risk?","levels":["none","recoverable","destructive"]}
```

Each nonblank row has a unique nonempty `id`, a nonempty `question`, exact lowercase `type`, and the matching primitive's criteria. Unknown keys are rejected. IDs and criteria have no surrounding whitespace. Choice needs 2–255 distinct nonempty conditions; JSON labels may contain commas. Score needs 2–10 distinct nonempty levels in ascending order.

Keep different states separate. Batch validates all input before one API request and emits complete JSONL in input order. It has no chunking, retry, resume, or partial success; threshold, min_score, json, and pick options or saved values do not apply.

## Use Results

Single commands print only `answer` by default. JSON fields are:

- Noul: `answer`, `model`.
- Choice: `answer`, `probabilities`, `confidence`, `model`.
- Score: `answer`, `legend`, `probabilities`, `confidence`, `model`.

Use `--json --pick answer,probabilities` to select available top-level fields; `--pick ''` clears a saved restriction. Score can fall between levels. `--threshold` tests the selected Choice probability or the Noul truth probability; Score uses `--min-score`. Confidence is separate from these criteria.

Single-command exits:

- `0`: use the successful stdout result once.
- `2`: criterion not met; empty stdout and boundary details on stderr.
- `1`: input, config, authentication, API, response, or output error.

Batch uses `0` for complete JSONL and `1` for failure; it never uses `2`. Failures before output begins leave stdout empty; output-device partial writes cannot be rolled back. Report failures without lowering criteria, changing settings, or repeating a successful call to check connectivity.

## Configuration

`~/.agents/jev.toml` stores defaults. Explicit flags override file values, which override built-ins. Nonempty `TYPESAFE_API_KEY` overrides file `api_key`. Built-ins: no key or result criteria, `model = "jev-latest"`, `timeout = "30s"`, `json = false`, unrestricted pick. Reads do not create files; successful config mutations are silent and offline.

```sh
jev config
jev config path
jev config set threshold 0.8
jev config set min_score 1.5
jev config set model jev-latest
jev config set timeout 45s
jev config set json true
jev config set pick answer,model
jev config unset threshold
jev config set api_key --stdin
```

Change settings only within the request. Pass an actual token from an authorized local source through stdin, never a conversation message or command argument. Config reads mask the complete key; storage is plaintext mode 0600. Symlink files are refused, and invalid TOML requires manual repair. Batch uses only saved key, model, and timeout.

## Install And Diagnose

Plugin installation does not put `jev` on PATH. From the plugin's `scripts/` directory, use `go run . install` for first installation or `go run . install --force` after updating source. `--dir <directory>` selects another target; the default is `~/.local/bin/jev`.

For requested maintenance or a reported setup problem:

```sh
jev install [--force] [--dir <directory>]
jev uninstall [--dir <directory>]
jev doctor [--dir <directory>]
```

`install` copies the running binary; it does not build or download new source. Replacing a different regular file requires `--force`; symlink and non-regular targets are refused. `uninstall` removes only an identifiable Jev regular file and succeeds if absent. Both preserve configuration, keys, directories, and shell settings and make no API calls.

`doctor` checks binary identity, PATH selection, config validity and permissions, and key presence and source. With valid settings and a key, its default performs one authenticated `GET https://api.typesafe.ai/v1/models` using the saved timeout. It makes no inference request, follows no redirects, retries nothing, and repairs no state. Invalid config or a missing key skips the API check.

Read the `OK / FAIL / SKIP` rows, fixes, and summary. HTTP 401 means the key was rejected; 403 means access denied. Network, timeout, rate-limit, and server failures leave key validity unknown. If `TYPESAFE_API_KEY` is selected, updating only the file key will not change authentication. Never print key values or API error bodies. A successful lookup proves authentication and API access, not model accuracy or inference availability.

Maintenance success uses stdout and exit `0`; problems use empty stdout, stderr, and exit `1`. Do not run doctor automatically before evaluation or after installation, or issue a separate API call to recheck a successful result.
