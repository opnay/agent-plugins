# Concepts and Compatibility

Use this reference for architectural explanations and decisions that cannot be made from the command cheatsheet alone.

## Project Names

- **Apple `container`**: the end-user CLI and its XPC-managed services.
- **Containerization**: the open-source Swift package used for image, process, filesystem, networking, and lightweight VM primitives.
- **OCI image**: the interoperable image format consumed and produced by `container`.

The CLI and Swift package are related but not interchangeable task surfaces. Use the CLI skill for local operation; API or framework development needs separate engineering guidance.

## Runtime Model

On macOS, Linux containers need Linux virtualization. Apple `container` runs each container in a lightweight VM instead of placing all containers in one shared Linux VM.

The practical consequences are:

- VM-level isolation per container
- host directory sharing scoped to the requesting container
- OCI image interoperability with standard registries
- macOS integration through Virtualization.framework, XPC, launchd, Keychain, and unified logging
- a service lifecycle controlled by `container system start` and `container system stop`

See the upstream [technical overview](https://github.com/apple/container/blob/main/docs/technical-overview.md) for current implementation details.

## Docker and Podman Distinctions

| Surface | Apple `container` | Docker Desktop / Docker CLI | Podman on macOS |
| --- | --- | --- | --- |
| Host scope | Apple silicon macOS | On macOS, Docker Desktop provides a local runtime; Docker CLI can target compatible engines | macOS CLI backed by a Linux VM |
| Image format | OCI-compatible | OCI-compatible ecosystem | OCI-compatible ecosystem |
| Local runtime model | Lightweight VM per container | Desktop-managed Linux VM and Docker Engine API | Podman machine VM and Podman API |
| Dockerfile build | Supported through `container build` | Supported | Supported |
| Docker Engine API socket | Do not assume | Native Docker contract | Compatibility API available with limits |
| Compose | Not a baseline CLI contract | Docker Compose | External provider through `podman compose` |

Docker Desktop subscription terms do not make the open-source Docker CLI or Dockerfile format the same licensed product. Treat the user's choice to avoid Docker CLI as an operational policy, not as proof that every Docker-named artifact must be removed.

## Compatibility Decision

Use Apple `container` directly when the work is local image, container, network, volume, registry, or service management and the installed CLI supports the required flags.

Pause direct translation when a project requires:

- Compose orchestration
- Docker Engine API or `/var/run/docker.sock`
- Testcontainers runtime discovery
- devcontainer integration
- buildx-specific builders or cache contracts
- Docker event, stats, or inspect schemas consumed by another tool

For these cases, inspect the consumer's actual protocol. A third-party plugin or compatibility shim may exist, but it must be explicitly selected and verified; its existence does not make it part of the Apple `container` baseline.

## Scope Separation

- Local macOS commands may use `container` while CI or production retains another OCI runtime.
- Dockerfiles can remain because `container build` supports them.
- Image references can remain portable when they use standard registries and OCI-compatible images.
- Version-specific behavior and experimental plugins should be checked with `container help` and the [release page](https://github.com/apple/container/releases), not frozen into permanent assumptions.

## Authoritative Sources

- [Apple container repository](https://github.com/apple/container)
- [Command reference](https://github.com/apple/container/blob/main/docs/command-reference.md)
- [Technical overview](https://github.com/apple/container/blob/main/docs/technical-overview.md)
- [Apple Containerization repository](https://github.com/apple/containerization)
- [Docker Desktop license agreement](https://docs.docker.com/subscription/desktop-license/)
- [Podman installation](https://podman.io/docs/installation)
