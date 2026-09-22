package main

const usage = `Usage:

Evaluation:
  jev noul --if <question and context> [options]
  jev choice --if <question and context> --conditions <a,b,...> [options]
  jev score --if <question and context> --level <description>... [options]
  jev --if <question and context> --conditions <a,b,...> [options]

Configuration:
  jev config
  jev config path
  jev config set <key> <value>
  jev config set api_key --stdin
  jev config unset <key>

Maintenance:
  jev install [--force] [--dir <directory>]
  jev uninstall [--dir <directory>]
  jev doctor [--dir <directory>]

Common evaluation options:
  --model <name>       Model (default: jev-latest)
  --timeout <duration> Request timeout (default: 30s)
  --json[=false]       Output JSON (default: false)
  --pick <fields>      JSON fields available for the selected primitive
  --help               Show this help

Primitive options:
  Noul:
    --threshold <0..1>       Minimum probability of yes

  Choice:
    --conditions <a,b,...>   2–255 distinct labels
    --threshold <0..1>       Minimum selected-label probability

  Score:
    --level <description>    Repeat for 2–10 ordered levels
    --min-score <number>     Minimum weighted score

Config keys:
  api_key, threshold, min_score, model, timeout, json, pick

Output:
  Noul:   answer, model
  Choice: answer, probabilities, confidence, model
  Score:  answer, legend, probabilities, confidence, model

Exit codes:
  0  Success
  1  Input, configuration, API, response, or output error
  2  Result criterion not met

Config path: ~/.agents/jev.toml; flags > file > defaults.
Auth: TYPESAFE_API_KEY > config api_key. Token input is stdin only.
Questions and criteria are sent to the hosted TypeSafe API.
Maintenance is offline. Default install directory: ~/.local/bin.
`

func usageFor(command string) string {
	switch command {
	case "noul":
		return noulUsage
	case "choice":
		return choiceUsage
	case "score":
		return scoreUsage
	case "config":
		return configUsage
	case "install":
		return installUsage
	case "uninstall":
		return uninstallUsage
	case "doctor":
		return doctorUsage
	default:
		return usage
	}
}
