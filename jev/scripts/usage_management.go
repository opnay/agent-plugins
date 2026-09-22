package main

const configUsage = `Usage:
  jev config
  jev config path
  jev config set <key> <value>
  jev config set api_key --stdin
  jev config unset <key>

Commands:
  config              Show effective settings; mask the complete token
  config path         Print ~/.agents/jev.toml without reading it
  config set          Store one validated value; success is silent
  config unset        Remove one stored value; a missing value succeeds

Options:
  --help              Show this help

Keys:
  api_key      No default; stdin only; TYPESAFE_API_KEY takes precedence
  threshold    Off by default; Choice selected probability and Noul yes probability
  min_score    Off by default; Score weighted result
  model        jev-latest
  timeout      30s
  json         false
  pick         No field restriction

Flags override file values, which override built-in defaults. Config commands are offline.
`

const installUsage = `Usage:
  jev install [--force] [--dir <directory>]

Options:
  --force             Replace a different existing regular file
  --dir <directory>   Installation directory (default: ~/.local/bin)
  --help              Show this help

Copies the running Jev binary to <directory>/jev with mode 0755. An identical
target succeeds unchanged. Symlinks and non-regular targets are refused. The
command does not download, build, edit PATH, or change config and tokens.
`

const uninstallUsage = `Usage:
  jev uninstall [--dir <directory>]

Options:
  --dir <directory>   Installation directory (default: ~/.local/bin)
  --help              Show this help

Removes <directory>/jev only when Go build information identifies it as Jev.
A missing target succeeds. Config, tokens, the directory, adjacent files, and
shell settings are preserved.
`

const doctorUsage = `Usage:
  jev doctor [--dir <directory>]

Options:
  --dir <directory>   Installation directory (default: ~/.local/bin)
  --help              Show this help

Checks the Jev binary, PATH selection, general config validity and permissions,
and token presence. Doctor is read-only and offline; it does not authenticate
the token, call the API, validate a model judgment, or repair state.
`
