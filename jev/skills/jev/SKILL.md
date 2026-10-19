---
name: jev
description: Explain and operate TypeSafe Jev's Noul, Choice, Score, and batch CLI, and apply basic patterns for framing judgments, using shared context, and interpreting results. Use for Jev or TypeSafe requests and bounded task judgments where Jev can change the next action. Jev usage, judgment framing, uncertain evidence, shared-state batch
---

# Jev

Provide Jev information, CLI operation, and basic patterns for using advisory judgments. Small choices, clarification decisions, evidence assessment, and model selection are built-in use cases. The current task and applicable instructions own domain policy, model policy, work decomposition, agent execution, and result integration.

## Frame A Useful Judgment

First inspect the request, rules, and available evidence. Proceed directly when they settle the answer. Use Jev when a bounded judgment can change the next action; do not ask it to guess facts that a quick local check can establish.

Supply the goal, confirmed facts, material unknowns, constraints, decision criterion, and permitted candidates. Include genuine recovery or uncertainty options such as `inspect_more`, `ask_user`, `insufficient`, or `not_shown` when applicable. Preserve related evidence together.

Questions, context, criteria, and batch state go to TypeSafe's hosted API. Send only authorized, necessary information. Jev cannot supply missing facts, user preferences, approval, authority, or proof of quality. Check its advice against original evidence and current instructions before acting.

Read [Built-In Use Cases](references/use-cases.md) for small choices, clarification decisions, broad and narrow evidence questions, or model selection. These examples apply the common pattern; they do not define task execution or model policy. This skill and its bundled reference require no sibling skill or development spec.

## Select The Primitive

| Command | Question | Answer |
|---|---|---|
| `jev noul` | One yes/no proposition | Probability of truth, 0–1 |
| `jev choice` | Discrete alternatives | Highest-probability label |
| `jev score` | Ordered descriptive levels | Weighted score, 0 through level count minus 1 |

Put the question and required context in `--if`:

```sh
jev noul --if "Does the supplied requirement make this mandatory? Context: ..."
jev choice --if "How does this evidence relate to the claim? Context: ..." \
  --conditions supported,contradicted,insufficient
jev score --if "What is the impact of this observed failure? Context: ..." \
  --level "No functional impact" \
  --level "Degraded with a workaround" \
  --level "Blocked without a workaround"
```

Choice accepts 2–255 unique comma-separated labels; labels cannot contain commas. Score accepts 2–10 unique `--level` values in ascending order; natural-language levels may contain commas. CLI, config, and JSON-owned keys use `snake_case`. Labels are free text; prefer `snake_case` for machine consumers. Legacy `jev --if ... --conditions ...` invokes Choice.

## Batch Independent Questions

Use batch for two or more independent questions sharing exactly one nonempty UTF-8 state:

```sh
jev batch --state-file evidence.md --input questions.jsonl
```

```jsonl
{"id":"relation","type":"choice","question":"How does this evidence relate to the claim?","conditions":["supported","contradicted","insufficient"]}
{"id":"requirement","type":"noul","question":"Does the supplied requirement mandate this behavior?"}
{"id":"impact","type":"score","question":"What impact does the observed failure have?","levels":["none","recoverable","blocked"]}
```

Each nonblank row needs a unique nonempty `id`, nonempty `question`, exact lowercase `type`, and matching `conditions` or `levels` where required. Unknown keys are rejected. IDs and criteria have no surrounding whitespace. Choice requires 2–255 distinct nonempty conditions; JSON conditions may contain commas. Score requires 2–10 distinct nonempty ascending levels.

Ask dependent questions sequentially, after inspecting the preceding result. Keep different states in separate calls. Do not impose a small fixed question count: use distinct action-relevant questions within current API request and context limits and reasonable transmission cost. Rephrased duplicates do not create independent evidence.

Batch validates the full input, sends one request, and emits complete JSONL in input order. It does not chunk, retry, resume, or provide partial success. Single-command `threshold`, `min_score`, `json`, and `pick` flags and saved values do not apply to batch.

## Interpret Results And Failures

Single commands print `answer` by default. Available JSON fields:

- Noul: `answer`, `model`.
- Choice: `answer`, `probabilities`, `confidence`, `model`.
- Score: `answer`, `legend`, `probabilities`, `confidence`, `model`.

Use `--json --pick answer,probabilities` for available top-level fields; `--pick ''` clears a saved restriction. Noul's answer is truth probability, Choice's answer is a label, and Score's answer can fall between levels. `--threshold` tests the selected Choice probability or Noul truth probability. Score uses `--min-score`. Confidence is separate from these criteria and does not prove correctness.

| Mode | Exit | Meaning |
|---|---|---|
| Single | `0` | Successful stdout result; consume once |
| Single | `2` | Result criterion not met; empty stdout, boundary details on stderr |
| Either | `1` | Input, config, authentication, API, response, or output failure |
| Batch | `0` | Complete JSONL result |

Batch never returns exit `2`. For either mode, exit `1` reports the reason on stderr. Failure before output starts leaves stdout empty; output-device partial writes may leave output and cannot be rolled back.

If the CLI is absent, a call fails, or a criterion is not met, do not automatically install tools, change credentials or settings, lower criteria, or repeat the same call. Continue with established defaults or local judgment when sufficient; ask the user only for a decision they must make. Distinguish actual Jev advice, failure, and local fallback. Do not repeat a successful call or run a separate connection check to validate it.

## Configure When Requested

Settings live in `~/.agents/jev.toml`. Explicit flags override file values, which override built-ins. A nonempty `TYPESAFE_API_KEY` overrides file `api_key`. Built-ins are no key or result criteria, `model = "jev-latest"`, `timeout = "30s"`, `json = false`, and unrestricted pick. Reads do not create files. Batch uses saved key, model, and timeout only.

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

Config mutations are offline and silent on success. Obtain an actual token only from an authorized local source and pass it through stdin, never conversation text or command arguments. Reads mask the entire key; storage is plaintext mode 0600. Symlink files are refused; invalid TOML requires manual repair. Preserve the requested scope of configuration changes.

## Install Or Diagnose When Requested

Plugin installation does not install the binary on PATH. From the plugin's `scripts/` directory, use `go run . install` for first installation or `go run . install --force` after source updates. Default destination: `~/.local/bin/jev`; use `--dir <directory>` for another target.

```sh
jev install [--force] [--dir <directory>]
jev uninstall [--dir <directory>]
jev doctor [--dir <directory>]
```

`install` copies the running binary; it does not build or download source. Replacing another regular file requires `--force`; symlink and non-regular targets are refused. `uninstall` removes only an identifiable Jev regular file and succeeds when absent. Both preserve settings, keys, directories, and shell configuration and make no API calls.

`doctor` checks binary identity, PATH selection, settings validity and permissions, and key presence and source. With valid settings and a key, it makes one authenticated `GET https://api.typesafe.ai/v1/models` using the saved timeout. Invalid settings or a missing key skips this lookup. It makes no inference request, follows no redirects, retries nothing, and repairs no state.

Read `OK / FAIL / SKIP`, suggested fixes, and the summary. HTTP 401 means a rejected key; 403 means denied access. Network, timeout, rate-limit, and server errors leave key validity unknown. If the environment key wins, changing only the file key will not change authentication. Never print key values or API error bodies. A successful lookup proves authentication and API access, not inference availability or model accuracy.

Maintenance success uses stdout and exit `0`; problems use empty stdout, stderr, and exit `1`. Do not run doctor automatically before evaluation or after installation.
