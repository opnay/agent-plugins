---
name: apple-container
description: Use Apple's container CLI and bundled experimental k8s plugin on Apple silicon macOS to build, run, inspect, and manage OCI containers and local Kubernetes clusters with explicit Docker and third-party CLI boundaries. Apple container, container CLI, container k8s, Apple container Kubernetes plugin, Docker Desktop alternative, Docker CLI migration, Podman migration, macOS OCI container
---

# Apple Container CLI

## Boundary

- Use `container` as the preferred direct local container CLI on Apple silicon macOS.
- Own the bundled experimental `container k8s` plugin only for local cluster lifecycle, image loading, and kubeconfig generation.
- Verify current core or plugin help before relying on a stored flag or subcommand; the cheatsheet is a fast path, not the source of truth.
- Preserve Dockerfiles, OCI image references, and remote CI or Linux runtime contracts unless the task proves they must change.
- Do not assume native Docker Compose or Docker Engine API compatibility. Do not install or switch to Docker, Podman, or a third-party shim without user authorization.
- Do not install or run `kubectl`, Helm, Kustomize, or another third-party Kubernetes CLI. Mention such tools only to explain kubeconfig consumers or the boundary of workload management.

## Preflight

Run only the checks needed for the task:

```sh
uname -m
sw_vers -productVersion
command -v container
container --version
container system status
container system version
```

If the CLI is missing or the host is incompatible, read the installation reference. If the service is stopped and execution is requested, start it with `container system start`.

For a Kubernetes plugin task, also confirm plugin discovery and use plugin-specific help:

```sh
container help
container k8s --help
container k8s <subcommand> --help
```

Do not substitute `container help k8s`; dynamically dispatched plugins own their help surface.

## Cheatsheet

### Help and system

```sh
container help
container help <core-command>
container system start
container system status
container system version
container system logs
container system df
container system stop
```

### Containers

```sh
# Run once and remove after exit
container run --rm <image> <command>

# Run a named background service with a published port
container run -d --name <name> -p <host-port>:<container-port> <image>

# List, inspect, and observe
container ls
container ls --all
container ls --format json
container inspect <name>
container logs <name>
container logs --follow <name>
container stats <name>

# Enter, copy, and stop
container exec -it <name> <command>
container cp <local-path> <name>:<container-path>
container cp <name>:<container-path> <local-path>
container stop <name>
container rm <name>
```

Use `--platform linux/amd64 --rosetta` only when the image or task requires x86-64 emulation. Use `--mount`, `--volume`, `--network`, CPU, and memory flags only after checking `container help run`.

### Images and builds

```sh
container image pull <registry>/<image>:<tag>
container image ls
container image inspect <image>:<tag>
container build -t <image>:<tag> .
container build -f <dockerfile> -t <image>:<tag> <context>
container image tag <source> <target>
container image push <registry>/<image>:<tag>
container image rm <image>:<tag>
```

`container build` accepts Dockerfile and Containerfile inputs and produces OCI images. Keep registry names explicit when ambiguity matters.

### Registry, networks, and volumes

```sh
container registry login --username <user> --password-stdin <registry>
container registry ls
container registry logout <registry>

container network create <network>
container network ls
container network inspect <network>
container network rm <network>

container volume create <volume>
container volume ls
container volume inspect <volume>
container volume rm <volume>
```

### Kubernetes plugin

Treat `container k8s` as an experimental local-development plugin. Confirm it appears under `PLUGINS` in `container help` before use.

```sh
container k8s --help
container k8s create --name <cluster> --cpus <cpus> --memory <memory>
container k8s list
container k8s start --name <cluster>
container k8s load-image --name <cluster> <image>:<tag>
container k8s write-config --name <cluster> --kubeconfig <path>
container k8s delete --name <cluster>
```

`create`, `start`, and `write-config` can create or refresh kubeconfig entries. `delete` removes the cluster container and its entry from the default `~/.kube/config`; do not assume it cleans alternate kubeconfig files. Resolve the exact cluster name and kubeconfig target first. Do not add third-party Kubernetes CLI commands to complete the workflow; read the Kubernetes reference when the task reaches workload-management boundaries.

### Cleanup

Treat `rm`, `delete`, `prune`, and data deletion as destructive. Resolve the exact target, inspect current state, and preserve unrelated resources.

```sh
container ls --all
container image ls
container network ls
container volume ls
container k8s list
container system df
```

Use a resource-specific `prune` only when the user requested that cleanup and its blast radius is understood.

## Compatibility Gate

Before translating Docker-oriented work, inspect for `compose.yaml`, `docker-compose.yml`, `DOCKER_HOST`, `/var/run/docker.sock`, Testcontainers, devcontainers, buildx, or tools that consume the Docker Engine API.

- Translate only commands supported by the installed `container` CLI.
- Do not translate `docker compose` to `container compose` unless a compatible plugin is installed and verified for the requested features.
- Report the unsupported contract and available options instead of silently changing runtimes.
- Keep local macOS tool preference separate from CI, Linux, and production runtime choices.
- Keep `container k8s` local cluster lifecycle separate from production orchestration and third-party Kubernetes client workflows.

## Reference Routing

- For prerequisites, signed package installation, initial service setup, update, or downgrade, read [references/installation.md](references/installation.md). A downgrade starts there and routes to the uninstallation reference when the existing installation must be removed.
- For removing the CLI, choosing whether to preserve user data, or verifying removal, read [references/uninstallation.md](references/uninstallation.md).
- For OCI, per-container VM architecture, Docker and Podman distinctions, or compatibility reasoning, read [references/concepts.md](references/concepts.md).
- For the experimental `container k8s` plugin model, cluster lifecycle, kubeconfig side effects, or third-party Kubernetes CLI boundary, read [references/kubernetes.md](references/kubernetes.md).
