---
name: jev
description: Explain and operate TypeSafe Jev's Noul, Choice, Score, and batch CLI, including structured JSON inputs, provider settings, and opt-in long context, and apply basic judgment patterns. Use for Jev or TypeSafe requests and bounded judgments that can change the next action. Jev usage, judgment framing, uncertain evidence, shared-state batch
---

# Jev

Help agents frame judgments, use the CLI, interpret results, and perform requested setup or maintenance. Domain policy, model policy, decomposition, agent execution, and integration belong to the current task and applicable instructions.

## Frame A Useful Judgment

Inspect the request, rules, and available evidence first. Use Jev when a bounded judgment can change the next action. If direct inspection or deterministic rules settle the answer, proceed from those facts.

Include the goal, confirmed facts, material unknowns, constraints, criterion, and permitted candidates. Add uncertainty or recovery options such as `insufficient`, `not_shown`, `inspect_more`, or `ask_user` when needed. Preserve relationships across source material.

Questions, state, and criteria go to TypeSafe's hosted API. Send only authorized, necessary content. Advice cannot provide missing facts, user preferences, approval, authority, or proof of quality. Check it against original evidence and current instructions before acting.

Read [Built-In Use Cases](references/use-cases.md) for small A/B choices, clarification, broad and narrow evidence questions, or model selection. These examples apply the common pattern; they do not establish execution or model policy. This skill and its reference require no sibling skill or development spec.

## Select The Primitive And Input

| Command | Judgment | Answer |
| --- | --- | --- |
| `jev noul` | Whether a proposition is true | Probability of yes, 0–1 |
| `jev choice` | One of defined alternatives | Highest-probability label |
| `jev score` | Degree along ordered levels | Weighted score, 0 through level count minus 1 |

For simple input, place the question and relevant context together in `--if`:

```sh
jev noul --if "Does the requirement mandate this? Context: ..."
jev choice --if "How does this evidence relate to the claim? Context: ..." \
  --conditions supported,contradicted,insufficient
jev score --if "What is the impact? Context: ..." \
  --level "No functional impact" --level "Workaround available" --level "Blocked"
```

Choice accepts 2–255 unique comma-separated labels; labels cannot contain commas in this form. Score accepts 2–10 unique ascending `--level` descriptions, including commas. CLI/config/JSON-owned keys use `snake_case`. Labels are free text; prefer `snake_case` for machine consumers. Legacy `jev --if ... --conditions ...` still invokes Choice.

For structured input, use `jev <primitive> --input request.json`. Do not combine it with `--if`, `--conditions`, or `--level`. Output and result-filter flags still apply.

```json
{
  "state": {"message": "I was charged twice."},
  "instructions": {"question": "Which team should handle message?"},
  "criteria": {
    "billing": {"description": "Payments and refunds"},
    "support": ["Technical problems", "Product usage"]
  }
}
```

The file accepts only these keys:

- `state`: required string/object/array; structure and numeric literals are preserved.
- `instructions`: optional string/object/array/null; omitted means null.
- Choice `criteria`: 2–255 nonempty labels without surrounding whitespace, mapped to string/object/array/null descriptions. JSON labels may contain commas.
- Score `criteria`: 2–10 distinct ascending string/object/array levels; JSON `legend` preserves their structure.
- Noul `criteria`: optional object containing `true` and/or `false` descriptions, or null. Descriptions accept string/object/array/null.

Blank strings, wrong types, unknown keys, invalid UTF-8/JSON, non-finite or out-of-range numbers, and nesting beyond 64 levels are rejected before transmission. Errors do not echo input contents. Read the selected command's `--help` for details.

## Batch Independent Questions

Use one call for independent questions sharing one state:

```sh
jev batch --state-file evidence.md --input questions.jsonl
jev batch --state-file state.json --state-format json --input questions.jsonl
```

Default `--state-format text` sends nonempty UTF-8 content unchanged. `json` decodes one string/object/array and preserves its numbers. Format is never inferred.

Each nonblank JSONL row needs a unique nonempty `id`, exact lowercase `type`, and exactly one of `question` or `instructions`:

```jsonl
{"id":"relation","type":"choice","question":"How does the evidence relate to the claim?","conditions":["supported","contradicted","insufficient"]}
{"id":"required","type":"noul","instructions":{"question":"Is this mandatory?"},"criteria":{"true":"Explicit obligation","false":"Advice or no obligation"}}
{"id":"impact","type":"score","instructions":["Rate functional impact"],"criteria":[{"impact":"none"},{"impact":"recoverable"},{"impact":"blocked"}]}
```

Legacy `question` is nonempty text; `instructions` accepts string/object/array/null. Legacy Choice `conditions` and Score `levels` are distinct nonempty string arrays with their primitive's limits. IDs and those values cannot have surrounding whitespace. Unknown keys are rejected.

Use `criteria` instead of `conditions`/`levels` for structured rubrics; it follows the single-input rules. A legacy question can use criteria. Never combine question with instructions, or criteria with conditions/levels. Noul forbids conditions and levels, Choice forbids levels, and Score forbids conditions.

Questions in one request cannot see each other's answers. Ask dependent questions sequentially after inspecting the preceding answer, and keep different states in separate calls. Use distinct action-relevant questions within API limits and reasonable transmission cost, without an arbitrary small question cap. Paraphrased duplicates are not independent evidence.

Batch validates all input, sends one request, and emits complete JSONL in input order. It does not chunk, retry, resume, or return partial success. Threshold/min_score/json/pick defaults and single-command flags do not apply to batch.

## Opt Into Long Context

`--long-context` sends `TypeSafe-Danger-Zone-Long-Context: true` for that single or batch request. It defaults off, accepts `--long-context=false`, and is neither saved nor enabled automatically.

The [model documentation](https://docs.typesafe.ai/models.md#long-context), checked 2026-10-19, gives default limits of 64k request tokens and 32k for state plus one question. Long context raises both to 262k but may be slower and less accurate. Check current limits before large calls. More input raises total cost even when token pricing is unchanged.

## Interpret Results And Failures

Single commands print the answer by default. Available JSON fields:

- Noul: `answer`, `model`.
- Choice: `answer`, `probabilities`, `confidence`, `model`.
- Score: `answer`, `legend`, `probabilities`, `confidence`, `model`.

`--json --pick answer,probabilities` selects available top-level fields; `--pick ''` clears a saved restriction. Noul is truth probability, Choice is a label, and Score may lie between levels. Confidence summarizes the distribution and does not prove correctness. `--threshold` tests the chosen Choice probability or Noul answer; Score uses `--min-score`.

| Mode | Exit | Meaning |
| --- | --- | --- |
| Single | `0` | Successful stdout result; consume once |
| Single | `2` | Criterion not met; empty stdout, details on stderr |
| Either | `1` | Input, config, authentication, API, response, or output error |
| Batch | `0` | Complete JSONL result |

Batch never returns exit 2. Errors before output starts leave stdout empty; output-device partial writes cannot be rolled back. Errors use stderr without keys or API error bodies.

On missing CLI, failure, or unmet criteria, do not automatically install, change credentials/settings, lower criteria, or repeat the call. Explain relevant setup when installation or authentication is the cause. Use established defaults or local judgment when sufficient; ask only for a decision the user must make. Distinguish actual advice, failure, and fallback. Never repeat successful calls or run connection checks to validate them.

## Configure When Requested

Settings live in `~/.agents/jev.toml`:

```toml
[defaults]
provider = "jev"
timeout = "30s"
json = false
pick = []

[providers.jev]
model = "jev-latest"
```

`defaults` owns provider, timeout, threshold, min_score, json, and pick. `providers.<name>` owns model and api_key. Provider selection is `--provider` > defaults.provider > built-in jev. Model selection is `--model` > selected provider model > that provider's built-in. Other explicit flags override defaults, then built-ins. Built-ins have no key/result filters, timeout 30s, json false, unrestricted pick; Jev's model is jev-latest.

Authentication uses the selected provider's nonempty environment variable before its file key: Jev uses `TYPESAFE_API_KEY`, OpenRouter uses `OPENROUTER_API_KEY`, and OpenAI uses `OPENAI_API_KEY`. Never reuse another provider's key or model. Only Jev has an execution adapter; OpenRouter/OpenAI settings can be saved but selecting them for execution fails without an API call. Batch uses the chosen provider, model, key, and timeout only.

```sh
jev config
jev config path
jev config init
jev config set defaults.provider jev
jev config set defaults.threshold 0.8
jev config set defaults.min_score 1.5
jev config set defaults.timeout 45s
jev config set defaults.json true
jev config set defaults.pick answer,model
jev config set providers.jev.model jev-latest
jev config set providers.jev.api_key --stdin
jev config unset defaults.threshold
```

`config init` creates the embedded template with mode 0600 and refuses to overwrite any existing file. The first set starts from the same template. Config reads, evaluations, and doctor never create files. Set/unset are offline and silent on success. Dotted keys are required. Legacy flat settings must be moved into defaults and providers.jev; do not auto-migrate a user's file.

Get tokens only from authorized local sources and pass them on stdin, never in conversation or command arguments. All provider keys are fully masked in displays; plaintext storage uses 0600 and atomic writes. Symlink files are refused; invalid TOML needs manual repair. Missing keys/models in unselected providers do not block a selected provider. Preserve the requested scope of settings changes.

## Install Or Diagnose When Requested

Plugin installation does not install a PATH binary. From the plugin's scripts directory, use `go run . install` initially or `go run . install --force` to build and replace from updated source. Default target is ~/.local/bin/jev; another location uses --dir.

```sh
jev install [--force] [--dir <directory>]
jev uninstall [--dir <directory>]
jev doctor [--dir <directory>] [--provider <name>]
```

Install copies the running binary; it does not build/download source. Replacing another regular file needs --force. Symlink/non-regular targets are refused. Uninstall removes only identifiable Jev regular files and succeeds if absent. Both preserve settings, tokens, directories, and shell configuration, and make no API calls.

Doctor checks binary identity, PATH selection, settings validity/permissions, and the selected key's presence/source. Valid Jev settings and a key enable one authenticated GET https://api.typesafe.ai/v1/models using the configured timeout; otherwise the lookup is skipped. It makes no inference request, follows no redirects, retries nothing, and repairs nothing.

Read OK/FAIL/SKIP, fixes, and summary. HTTP 401 means key rejection; 403 means denied access. Network, timeout, rate-limit, and server failures leave validity unknown. If an environment key wins, changing only the file key does not change authentication. Never expose tokens or API error bodies. Successful lookup proves authentication and lookup access, not model accuracy or inference availability.

Maintenance success uses stdout/exit 0; problems use empty stdout and stderr/exit 1. Do not run doctor automatically before evaluation or after installation.
