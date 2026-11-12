# Uninstallation

Use this reference when removing the Apple `container` CLI itself. Resource deletion inside a retained installation belongs to the `SKILL.md` cheatsheet and command-specific help.

## Choose the Data Boundary

The official uninstaller requires exactly one mode:

- `-k`: remove the tool and helpers but keep user data for a later reinstall.
- `-d`: remove the tool, helpers, user data, and stored defaults.

Prefer `-k` when the request is only to uninstall the CLI. Use `-d` only when the user explicitly requests removal of local container data and understands that images, containers, volumes, configuration, and metadata stored under the application data root may be lost.

## Inspect Before Removal

```sh
container system status
container system df
container ls --all
container image ls
container volume ls
container network ls
```

Report relevant retained or deleted state before proceeding. Do not use direct recursive deletion as a substitute for the bundled uninstaller.

## Remove the Tool

Stop the services first:

```sh
container system stop
```

Then run exactly one official mode:

```sh
# Keep user data
/usr/local/bin/uninstall-container.sh -k

# Delete user data
/usr/local/bin/uninstall-container.sh -d
```

The script requests administrator privileges for installed files. The `-d` path also deletes the Apple container application data and defaults; it is the destructive branch.

## Verify

```sh
command -v container
pkgutil --pkg-info com.apple.container-installer
```

Successful removal should leave no executable on `PATH` and no registered installer package. With `-k`, do not claim user data was removed. With `-d`, report what the official script removed and that recovery depends on an external backup.

If the bundled uninstaller is missing, stop and inspect the installed package receipt and version. Do not reconstruct its internal deletion commands from memory.
