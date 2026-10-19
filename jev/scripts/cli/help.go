package cli

const usage = `Usage:
  jev <command> [options]

Commands:
  noul       Estimate the probability that a statement is true
  choice     Select an option from named choices
  score      Rate a state against ordered levels
  batch      Evaluate multiple questions against one shared state
  config     View and edit defaults and provider settings
  install    Install the Jev executable
  uninstall  Remove the installed Jev executable
  doctor     Check local setup and API authentication

Shared evaluation options (noul, choice, score, batch):
  --provider <name>     Decision provider; also supported by doctor
  --model <name>        Decision model
  --timeout <duration>  Request timeout
  --long-context       Enable expanded context

Run jev <command> --help for command usage and options.
`
