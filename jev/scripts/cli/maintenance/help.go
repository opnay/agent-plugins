package maintenance

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
  jev doctor [--dir <directory>] [--provider <name>]

Options:
  --dir <directory>   Installation directory (default: ~/.local/bin)
  --provider <name>   Override defaults.provider (built-in: jev)
  --help              Show this help

Checks the Jev binary, PATH selection, config validity and permissions, and
selected provider API key presence and source. Only Jev is implemented.
With valid Jev config and a key, sends one authenticated
GET https://api.typesafe.ai/v1/models using the configured timeout. No inference,
redirect, retry, or automatic repair.

Reports OK / FAIL / SKIP, problem-specific fixes, and a summary. Key values and
API response bodies are hidden. Success: stdout and exit 0; problems: stderr
and exit 1. This checks API access, not model accuracy or inference availability.
`

func usageFor(command string) string {
	switch command {
	case "install":
		return installUsage
	case "uninstall":
		return uninstallUsage
	default:
		return doctorUsage
	}
}
