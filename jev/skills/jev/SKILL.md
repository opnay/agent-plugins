---
name: jev
description: Explain TypeSafe Jev and use its hosted Noul, Choice, Score, and batch CLI; interpret results or manage Jev settings and the executable when requested. Use for explicit Jev or TypeSafe questions and CLI operations, not routine work routing.
---

# Jev CLI

Jev is a hosted structured-judgment service. The CLI sends the question, relevant context, choice labels or score levels, and batch state to TypeSafe. Send only content authorized for that external call. This skill explains the primitives and CLI contract; `$jev:jev-use` owns when a routine task should ask Jev for an advisory next action.

## Select The Primitive

- `noul`: probability from 0 to 1 that one proposition is true.
- `choice`: highest-probability label among discrete alternatives.
- `score`: probability-weighted value across ordered descriptive levels, from 0 through `level count - 1`.

Include the question and all material context in `--if`. Keep related evidence together rather than treating isolated files as a complete review. Do not report a judgment as if it establishes facts absent from the input.

```sh
jev noul --if "Is this statement mandatory? Context: ..."

jev choice --if "How does the evidence relate to the claim? Context: ..." \
  --conditions supported,conflicting,undetermined

jev score --if "How severe is this issue? Context: ..." \
  --level "No functional impact" \
  --level "Degraded with a workaround" \
  --level "Blocked with no workaround"
```

Choice accepts 2–255 unique, comma-separated labels in a single command; labels cannot contain commas. Score accepts 2–10 distinct `--level` values in ascending order, including commas in level text. Prefer `snake_case` for labels consumed by code. CLI, configuration, and JSON-owned keys use `snake_case`; natural-language labels do not have that restriction. The legacy `jev --if ... --conditions ...` is a single Choice call.

## Batch Questions With One State

Use `jev batch` for at least two independent questions that share exactly one nonempty UTF-8 state. Do not combine different states to reduce calls.

```sh
jev batch --state-file target.md --input questions.jsonl
```

Each nonblank JSONL row needs a unique nonempty `id`, a nonempty `question`, `type` (`noul`, `choice`, or `score`), and the corresponding `conditions` or `levels`:

```jsonl
{"id":"S001","type":"choice","question":"Which action applies?","conditions":["proceed","block","unexpected_action"]}
{"id":"S002","type":"noul","question":"Is the proposition true?"}
{"id":"S003","type":"score","question":"How severe is the risk?","levels":["none","recoverable","destructive"]}
```

Use exact lowercase keys and no unknown keys. IDs, conditions, and levels have no surrounding whitespace. Choice needs 2–255 distinct nonempty conditions; JSON labels may contain commas. Score needs 2–10 distinct nonempty levels in ascending order. The CLI validates the full input before one request, then prints one full JSON object per row in input order. Batch has no chunking, retry, resume, partial results, threshold, min_score, json, or pick option. It ignores saved single-result criteria and output selection.

## Interpret Output And Exits

A successful single command prints only its primitive answer by default. Use JSON when supporting fields matter:

- Noul: `answer`, `model`.
- Choice: `answer`, `probabilities`, `confidence`, `model`.
- Score: `answer`, `legend`, `probabilities`, `confidence`, `model`.

`--json --pick answer,probabilities` limits output to named top-level fields that exist for the primitive; `--pick ''` clears a saved restriction. Score can return a value between levels. Choice `--threshold` checks the selected label's probability; Noul `--threshold` checks the yes probability; Score uses `--min-score`. API confidence is separate and does not replace these criteria.

For a single command, exit `0` means consume stdout once; exit `2` means the result criterion was not met, with empty stdout and a boundary report on stderr; exit `1` means an input, configuration, API, response, or output error. Do not quietly lower a failed criterion or repeat a successful call as a connection check. Batch exits `0` only for the complete JSONL result and `1` for failure with empty stdout before output begins; it never exits `2`. A partial write to the output device cannot be rolled back.

## Settings And Executable

Settings are in `~/.agents/jev.toml`. Explicit flags override file values, which override built-ins. Nonempty `TYPESAFE_API_KEY` overrides the file's `api_key`. Built-ins are no token, no result criteria, `model = "jev-latest"`, `timeout = "30s"`, `json = false`, and no pick restriction. Reads do not create files; successful mutations are silent and offline.

```sh
jev config
jev config path
jev config set threshold 0.8
jev config set min_score 1.5
jev config set pick answer,model
jev config unset threshold
jev config set api_key --stdin
```

Change settings only when requested. Read a token from an authorized local source through stdin, never a conversation message or command argument. `jev config` masks the token on reads; the config file itself is plaintext mode 0600. Symlink config files are refused; invalid TOML needs manual repair. Batch uses only saved `api_key`, `model`, and `timeout`.

For first installation, run `go run . install` from the plugin's `scripts/` directory. After a source update, run `go run . install --force`; `--dir <directory>` changes the target. Plugin installation does not put the CLI on PATH. Run `jev install [--force] [--dir <directory>]`, `jev uninstall [--dir <directory>]`, or `jev doctor [--dir <directory>]` only for requested maintenance or a reported setup problem. The default target is `~/.local/bin/jev`. `install` copies the running binary and needs `--force` to replace a different regular file; symlink and non-regular targets are refused. `uninstall` removes only an identifiable Jev binary, succeeds when the target is absent, and preserves configuration, tokens, directories, and shell settings. `doctor` checks local identity, PATH, config validity and permissions, and token presence without calling the API or repairing state; a problem reports on stderr with exit `1`. Maintenance succeeds with short stdout and exit `0`. Do not run doctor automatically before evaluation or after installation.
