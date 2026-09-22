package main

const evaluationCommonOptions = `Common options:
  --model <name>       Model (default: jev-latest)
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

` + evaluationExitCodes

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

` + evaluationExitCodes

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

` + evaluationExitCodes
