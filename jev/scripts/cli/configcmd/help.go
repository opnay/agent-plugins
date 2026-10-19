package configcmd

const configUsage = `Usage:
  jev config
  jev config path
  jev config init
  jev config set <dotted-key> <value>
  jev config set providers.<name>.api_key --stdin
  jev config unset <dotted-key>

Commands:
  config              Show defaults and provider settings; mask every key
  config path         Print ~/.agents/jev.toml without reading it
  config init         Create the embedded template (0600); never overwrite
  config set          Store one validated value; first set starts from the template
  config unset        Remove a saved value; missing values do not create files

Options:
  --help              Show this help

Keys:
  defaults.provider   jev (built-in default); also openrouter or openai
  defaults.timeout    30s
  defaults.json       false
  defaults.pick       No field restriction
  defaults.threshold  Off; Choice selected probability and Noul yes probability
  defaults.min_score  Off; Score weighted result
  providers.<name>.model    jev-latest for Jev; no other model defaults
  providers.<name>.api_key  Stdin only; selected provider environment wins

Providers: jev / openrouter / openai. Only Jev has an execution adapter today.
Auth environments: TYPESAFE_API_KEY / OPENROUTER_API_KEY / OPENAI_API_KEY.
Provider/model flags override the selected provider's file settings; common flags
use [defaults]. No key or model crosses provider boundaries. Config is offline.
Legacy flat keys must be moved to [defaults] and [providers.jev]; no auto-migration.
Reads never create files. init refuses any existing file, including invalid TOML.
Config mutations are silent on success. Secrets are plaintext, mode 0600.
`
