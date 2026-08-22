# Installation and Updates

Use this reference only when the Apple `container` CLI is missing, needs initial setup, or must be updated or downgraded.

## Requirements

- Apple silicon Mac
- A macOS version supported by the current upstream release
- Administrator access for files installed under `/usr/local`
- Network access for package, kernel, and image downloads

Check the current requirements in the [Apple container README](https://github.com/apple/container#requirements) before changing the machine. Do not infer support from the cheatsheet or an older installed version.

## Inspect Before Installing

```sh
uname -m
sw_vers
command -v container
container --version
container system status
```

If `container` already exists, decide whether the task requires repair, update, or no change. Do not overwrite a working installation merely because installation was mentioned.

## Initial Installation

1. Open the [official GitHub releases](https://github.com/apple/container/releases).
2. Download the signed installer package for the chosen release.
3. Run the package installer and approve its `/usr/local` changes.
4. Start the services and allow the recommended kernel installation when appropriate:

```sh
container system start
container system status
container system version
```

5. When a networked smoke test is authorized, run:

```sh
container run --rm docker.io/library/alpine:latest echo hello
```

The smoke test may download an image and create transient runtime state. Do not run it for a read-only inspection request.

Prefer the signed upstream package. Do not substitute a community package manager, unsigned package, Docker Desktop, Podman, or another runtime unless the user requested that path.

## Update or Reinstall

The signed installer places an update helper at `/usr/local/bin/update-container.sh`.

```sh
container system stop

# Latest release
/usr/local/bin/update-container.sh

# Specific same-or-newer release
/usr/local/bin/update-container.sh -v <version>

# Reinstall the selected release
/usr/local/bin/update-container.sh -v <version> -f

container system start
container system version
```

The helper uses network access, downloads a package, and requests administrator privileges. Compare the installed and target versions before using this path. Report when only an unsigned package is available; do not accept that fallback silently.

## Downgrade

Upstream requires uninstalling the existing release before installing an older one. Do not run the update-only sequence above as a downgrade.

1. Download the target release's signed installer package from the [official release page](https://github.com/apple/container/releases) before removing the current installation.
2. Read [uninstallation.md](uninstallation.md) and choose `-k` to retain user data or `-d` only when the user explicitly requests data deletion.
3. Stop the services and run the selected official uninstaller mode.
4. Install the already-downloaded signed package for the target release.
5. Start the services and verify the actual versions:

```sh
container system start
container system version
```

The installed uninstaller removes packaged helpers, so do not depend on `/usr/local/bin/update-container.sh` remaining available after removal. Preserve the downloaded target package or follow the current upstream manual installation path.

## Failure Routing

- Unsupported host: stop and report the architecture or macOS mismatch.
- Service start failure: inspect `container system logs` and the exact installed component versions.
- Registry or kernel download failure: separate network, credential, certificate, and upstream availability evidence.
- Version mismatch after installation: inspect `command -v container`, `container system version`, and stale service state before reinstalling.
