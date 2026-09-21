# alias.codex maintenance

Use this reference only for an explicit `alias.codex` install, uninstall, or doctor request.

## Contract

The expected global alias dispatches to `$HOME/.local/bin/git-codex`. The CLI owns its exact value, binary verification, and preservation rules. Alias maintenance does not modify PATH or shell profiles.

## Install

```sh
git codex install
git codex install --force
```

Install copies the current executable to `$HOME/.local/bin/git-codex`, configures the expected global alias, and runs its doctor checks. A different existing binary or alias is preserved unless `--force` is explicitly authorized.

## Uninstall

```sh
git codex uninstall
git codex uninstall --force
```

Uninstall removes the expected alias only and preserves `$HOME/.local/bin/git-codex`. A different existing alias is preserved unless `--force` is explicitly authorized.

## Doctor

```sh
git codex doctor
```

Doctor is read-only. It verifies the expected alias and binary, then reports the alias value, binary path, and version. Do not repair configuration unless the user separately requests install or uninstall.
