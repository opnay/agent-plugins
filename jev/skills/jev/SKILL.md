---
name: jev
description: Explain and operate TypeSafe Jev's Noul, Choice, Score, and batch CLI, including structured JSON inputs and opt-in long context, and apply basic judgment patterns. Use for Jev or TypeSafe requests and bounded judgments that can change the next action. Jev usage, judgment framing, uncertain evidence, shared-state batch
---

# Jev

Help agents use Jev through the CLI: frame judgments, select primitives and input forms, interpret results, and perform requested maintenance. The current task and instructions own domain policy, model policy, decomposition, agent execution, and integration.

## Frame The Judgment

Inspect the request, rules, and observable evidence first. Use Jev when a bounded judgment can change the next action. If direct inspection or deterministic rules settle the answer, proceed from those facts.

Supply the goal, known facts, material gaps, constraints, decision criterion, and permitted candidates. Include uncertainty or recovery outcomes such as `insufficient`, `not_shown`, `inspect_more`, or `ask_user` where useful. Preserve relationships across source material.

Questions, state, and criteria go to TypeSafe's hosted API. Send only authorized, necessary content. Advice cannot provide missing facts, preferences, approval, authority, or proof of quality; check it against original evidence and current instructions.

Read [Built-In Use Cases](references/use-cases.md) for small A/B choices, clarification, uncertain evidence, or model selection. These are examples of the common pattern, not execution or model policy. This skill and its reference work without sibling skills or development specs.

## Choose A Primitive And Input

| Command | Judgment | Answer |
| --- | --- | --- |
| `jev noul` | Whether a proposition is true | Probability of yes, 0–1 |
| `jev choice` | One option from a defined set | Highest-probability label |
| `jev score` | Degree on ordered descriptive levels | Weighted score, 0 through level count minus 1 |

For simple inputs, put the question and needed context in `--if`:

```sh
jev noul --if "Does the requirement mandate this? Context: ..."
jev choice --if "How does the evidence relate to the claim? Context: ..." \
  --conditions supported,contradicted,insufficient
jev score --if "What is the impact? Context: ..." \
  --level "No functional impact" --level "Workaround available" --level "Blocked"
```

Choice takes 2–255 unique comma-separated labels; a label cannot contain a comma in this form. Score takes 2–10 unique ascending `--level` descriptions, which may contain commas. CLI, config, and JSON-owned keys use `snake_case`; labels are free text, with `snake_case` preferred for machine consumers. `jev --if ... --conditions ...` remains a Choice alias.

For structured state, instructions, or criteria, write a JSON file and use:

```sh
jev choice --input request.json --json
```

```json
{
  "state": {"message": "I was charged twice."},
  "instructions": {"question": "Which team should handle message?"},
  "criteria": {
    "billing": {"description": "Payments and refunds", "examples": ["Duplicate charge"]},
    "support": ["Technical problems", "Product usage"]
  }
}
```

The file accepts only `state`, `instructions`, and `criteria`:

- `state`: required string, object, or array; numbers and structure are preserved.
- `instructions`: optional string, object, array, or null; omitted means null.
- Choice `criteria`: 2–255 nonempty labels without surrounding whitespace, mapped to string/object/array/null descriptions. JSON labels may contain commas.
- Score `criteria`: 2–10 distinct ascending string/object/array levels. JSON output preserves their structure in `legend`.
- Noul `criteria`: optional object with `true` and/or `false` descriptions, or null. Descriptions accept string/object/array/null.

Do not mix `--input` with `--if`, `--conditions`, or `--level`. Result filters and output flags still apply. Blank strings, wrong types, unknown keys, invalid UTF-8/JSON, non-finite or out-of-range numbers, and nesting beyond 64 levels are rejected before sending. Input contents are not echoed in errors. Use the primitive's `--help` for its complete input contract.

## Batch Questions Over Shared State

For two or more independent questions sharing one state:

```sh
jev batch --state-file evidence.md --input questions.jsonl
jev batch --state-file state.json --state-format json --input questions.jsonl
```

Default `--state-format text` sends nonempty UTF-8 file content unchanged. `json` sends a decoded string/object/array and preserves its numbers; there is no automatic format detection.

Each nonblank JSONL row needs a unique nonempty `id`, exact lowercase `type`, and either `question` or `instructions`:

```jsonl
{"id":"relation","type":"choice","question":"How does the evidence relate to the claim?","conditions":["supported","contradicted","insufficient"]}
{"id":"required","type":"noul","instructions":{"question":"Is this mandatory?"},"criteria":{"true":"Explicit obligation","false":"Advice or no obligation"}}
{"id":"impact","type":"score","instructions":["Rate functional impact"],"criteria":[{"impact":"none"},{"impact":"recoverable"},{"impact":"blocked"}]}
```

Legacy `question` is nonempty text. `instructions` accepts string/object/array/null. Legacy Choice `conditions` and Score `levels` are distinct nonempty string arrays with the primitive's limits; IDs and those values cannot have surrounding whitespace. Unknown keys are errors.

Use `criteria` instead of `conditions` or `levels` for structured rubrics; it follows the single-input rules. A legacy `question` can also use `criteria`. Never combine `question` with `instructions`, or `criteria` with `conditions`/`levels`. Noul forbids conditions and levels, Choice forbids levels, and Score forbids conditions.

Questions in one call cannot depend on each other's answers. Ask dependent questions sequentially after inspecting the preceding answer, and keep different states in separate calls. Use distinct action-relevant questions within current API limits and reasonable transmission cost, without an arbitrary small question cap. Paraphrased duplicates are not independent evidence.

Batch validates all input, sends one request, and emits complete JSONL in input order. It does not chunk, retry, resume, or return partial success. Saved or single-command `threshold`, `min_score`, `json`, and `pick` do not apply to batch.

## Enable Long Context Explicitly

Add `--long-context` to a single or batch call only when it needs the expanded window. It sends `TypeSafe-Danger-Zone-Long-Context: true` for that request; it is off by default, accepts `--long-context=false`, is never enabled automatically, and is not saved in config.

The [model documentation](https://docs.typesafe.ai/models.md#long-context), checked 2026-10-19, gives default limits of 64k tokens per request and 32k for state plus one question. Long context raises both to 262k, but may be slower and less accurate. Check current limits before large requests. More input increases total cost even when the per-token price is unchanged.

## Interpret Results And Failure

Single commands print the answer by default. JSON fields are:

- Noul: `answer`, `model`.
- Choice: `answer`, `probabilities`, `confidence`, `model`.
- Score: `answer`, `legend`, `probabilities`, `confidence`, `model`.

`--json --pick answer,probabilities` selects available top-level fields. `--pick ''` clears a saved restriction. Noul is a truth probability; Choice is a label; Score can lie between levels. Confidence describes the distribution and does not prove correctness. `--threshold` uses the chosen Choice probability or Noul answer; Score uses `--min-score`.

| Mode | Exit | Meaning |
| --- | --- | --- |
| Single | `0` | Successful stdout result; consume once |
| Single | `2` | Criterion not met; empty stdout, details on stderr |
| Either | `1` | Input, config, authentication, API, response, or output error |
| Batch | `0` | Complete JSONL result |

Batch never returns exit `2`. Failures before output starts leave stdout empty. Output-device partial writes cannot be rolled back. Errors use stderr; API bodies and keys are not exposed.

If the CLI is absent, a call fails, or a criterion is unmet, do not automatically install, change credentials/settings, lower criteria, or repeat the call. If the failure indicates missing installation or authentication, explain the relevant setup or repair without changing it automatically. Use established defaults or local judgment when sufficient, and ask only for decisions the user must make. Distinguish actual advice, failure, and fallback. Do not repeat successful calls or run connection checks to validate them.

## Configure And Maintain When Requested

Settings are in `~/.agents/jev.toml`. Explicit flags override file values, then built-ins; nonempty `TYPESAFE_API_KEY` overrides file `api_key`. Built-ins: no key or result criteria, `model = "jev-latest"`, `timeout = "30s"`, `json = false`, unrestricted pick. Batch uses only the saved key, model, and timeout. Reads never create the file.

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

Config mutations are offline and silent on success. Obtain tokens only from an authorized local source and pass them on stdin, never in conversation or command arguments. Reads mask the entire token; storage is plaintext mode 0600. Symlink files are refused; invalid TOML needs manual repair. Keep changes within the requested scope.

Plugin installation does not install a PATH binary. In the plugin's `scripts/` directory, run `go run . install` initially, or `go run . install --force` to build and replace from updated source. Default target: `~/.local/bin/jev`; another directory uses `--dir`.

```sh
jev install [--force] [--dir <directory>]
jev uninstall [--dir <directory>]
jev doctor [--dir <directory>]
```

`install` copies the running binary; it does not download or build source. Replacing another regular file requires `--force`. Symlink and non-regular targets are refused. `uninstall` removes only an identifiable Jev regular file and succeeds if absent. Both preserve settings, tokens, directories, and shell configuration, and make no API calls.

`doctor` checks binary identity, PATH selection, settings validity/permissions, and key presence/source. Valid settings and a key enable one authenticated `GET https://api.typesafe.ai/v1/models` using the saved timeout; otherwise that lookup is skipped. It makes no inference request, follows no redirects, retries nothing, and repairs nothing.

Read `OK / FAIL / SKIP`, fixes, and summary. HTTP 401 means a rejected key; 403 means access denied. Network, timeout, rate-limit, and server failures leave key validity unknown. If the environment key wins, changing only the file key will not change authentication. Never print keys or API error bodies. A successful lookup proves authentication and lookup access, not model accuracy or inference availability.

Maintenance success uses stdout/exit `0`; problems use empty stdout, stderr/exit `1`. Do not run doctor automatically before evaluation or after installation.
