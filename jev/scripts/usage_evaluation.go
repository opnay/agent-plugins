package main

const structuredInputHelp = `Structured input:
  jev <primitive> --input <path> [options]
  UTF-8 JSON object with state, optional instructions, and criteria:
    state:        nonempty text, object, or array
    instructions: text, object, array, or null (default: null)
    Choice criteria: object of 2–255 labels mapped to text/object/array/null
    Score criteria: 2–10 distinct ordered text/object/array levels
    Noul criteria: optional object with true/false descriptions, or null
  Noul JSON example: {"state":{"message":"Notice"},"instructions":{"question":"Urgent?"},"criteria":{"true":"Time-sensitive","false":"Routine"}}
  --input cannot be combined with --if, --conditions, or --level.
  Unknown keys, invalid UTF-8/JSON, scalar state/instructions, non-finite or
  out-of-range numbers, and nesting beyond 64 levels are rejected before sending.
  Score JSON legend preserves the structured levels.

`

const longContextHelp = `Long context:
  --long-context enables TypeSafe-Danger-Zone-Long-Context: true for this call.
  Default: off (64k tokens per request, 32k state + one question).
  Enabled: 262k tokens for each limit; may be slower and less accurate.
  Not saved in config; never enabled automatically.

`

const evaluationCommonOptions = `Common options:
  --input <path>      JSON state, instructions, and criteria file
  --long-context[=false] Enable expanded context for this request
  --provider <name>    Decision provider (defaults.provider; built-in: jev)
  --model <name>       Model (selected provider default; Jev: jev-latest)
  --timeout <duration> Request timeout (default: 30s)
  --json[=false]       Output JSON (default: false)
  --pick <fields>      Keep named JSON fields
  --help               Show this help
`

const evaluationExitCodes = `Exit codes:
  0  Success
  1  Input, configuration, API, response, or output error
  2  Result criterion not met

Questions and criteria are sent to the hosted TypeSafe API.
`

const noulUsage = `Usage:
  jev noul --if <question and context> [options]

Evaluates whether one proposition is true.

Noul options:
  --if <text>           Question and required context
  --threshold <0..1>   Minimum probability of yes (default: off)

` + evaluationCommonOptions + `
Output:
  Default: probability of yes from 0 to 1
  JSON:    answer, model

Example:
  jev noul --if "Is this setting mandatory? Context: ..." --threshold 0.8

` + structuredInputHelp + longContextHelp + evaluationExitCodes

const choiceUsage = `Usage:
  jev choice --if <question and context> --conditions <a,b,...> [options]
  jev --if <question and context> --conditions <a,b,...> [options]

Selects one label from a discrete set. The second form is retained for compatibility.

Choice options:
  --if <text>             Question and required context
  --conditions <a,b,...>  2–255 distinct labels; labels cannot contain commas
  --threshold <0..1>      Minimum selected-label probability (default: off)

` + evaluationCommonOptions + `
Output:
  Default: selected label
  JSON:    answer, probabilities, confidence, model

Use snake_case labels when code consumes the result. Labels remain free strings.

Example:
  jev choice --if "How does the evidence relate to the claim? Context: ..." \
    --conditions supported,conflicting,undetermined

` + structuredInputHelp + longContextHelp + evaluationExitCodes

const scoreUsage = `Usage:
  jev score --if <question and context> --level <description>... [options]

Rates one dimension against ordered descriptive levels.

Score options:
  --if <text>             Question and required context
  --level <description>   Repeat for 2–10 distinct levels in ascending order
  --min-score <number>    Minimum weighted score (default: off)

` + evaluationCommonOptions + `
Output:
  Default: weighted score from 0 through level count - 1
  JSON:    answer, legend, probabilities, confidence, model

Level descriptions are natural language and may contain commas. A score may fall between levels.

Example:
  jev score --if "How severe is the issue? Context: ..." \
    --level "No functional impact" \
    --level "Degraded, workaround available" \
    --level "Blocked, no workaround" \
    --min-score 1.5

` + structuredInputHelp + longContextHelp + evaluationExitCodes

const batchUsage = `Usage:
  jev batch --state-file <path> --input <path> [options]

Evaluates JSONL questions against one shared state in one API request.

Batch options:
  --state-file <path>  UTF-8 shared state; non-whitespace and sent unchanged
  --input <path>       JSONL questions; blank lines ignored, one row required
  --state-format <text|json> Shared state format (default: text)
  --long-context[=false] Enable expanded context for this request
  --provider <name>    Decision provider (defaults.provider; built-in: jev)
  --model <name>       Model (selected provider default; Jev: jev-latest)
  --timeout <duration> Request timeout (default: 30s)
  --help               Show this help

Each nonblank JSONL row is one object with exact lowercase keys:
  Noul:   {"id":"...","type":"noul","question":"..."}
  Choice: {"id":"...","type":"choice","question":"...","conditions":["a","b"]}
  Score:  {"id":"...","type":"score","question":"...","levels":["low","high"]}

IDs and legacy questions must be nonempty strings; IDs must be unique. Choice
requires 2–255 distinct, nonempty condition strings.
Score requires 2–10 distinct, nonempty level strings in ascending order.
IDs, conditions, and levels cannot have surrounding whitespace.
Noul forbids conditions and levels; Choice forbids levels; Score
forbids conditions. Unknown or differently cased keys are rejected. The
complete state and input are validated before the single API request.

Structured batch input:
  Use instructions (text/object/array/null) instead of question, and criteria
  instead of conditions or levels. question and instructions are mutually exclusive;
  criteria cannot be mixed with conditions or levels. A legacy question can use criteria.
  Choice criteria maps 2–255 labels to text/object/array/null descriptions.
  Score criteria contains 2–10 distinct ordered text/object/array levels.
  Noul criteria is an optional object of true/false descriptions (or null).
  JSON values are limited to 64 nested levels and finite, in-range numbers.
  Example: {"id":"N01","type":"noul","instructions":{"question":"Urgent?"},"criteria":{"true":"Time-sensitive","false":"Routine"}}
  --state-format json decodes state as a JSON string/object/array, preserving
  its structure and numbers. text sends the file unchanged; no auto-detection.

Output:
  One result object per input row as JSONL, preserving input order:
    Noul:   id, type, answer, model
    Choice: id, type, answer, probabilities, confidence, model
    Score:  id, type, answer, legend, probabilities, confidence, model
  Batch does not apply threshold, min_score, json, or pick settings.

` + longContextHelp + `Exit codes:
  0  Complete batch success
  1  Input, file, configuration, API, response, or output error

Questions, criteria, and shared state are sent to the hosted TypeSafe API.
Batch does not chunk, retry, resume, or return partial results.
Complete results are encoded before stdout is written. Input, state/input file,
configuration, API, and response failures leave stdout empty. An output-device
partial write cannot be rolled back.
`
