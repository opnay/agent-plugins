# Kubernetes Plugin

Use this reference for the bundled experimental `container k8s` plugin. It owns local cluster lifecycle and explanatory boundaries, not general Kubernetes client work.

## Scope

`container k8s` manages local single-node Kubernetes development clusters backed by Apple container VMs. The cluster control-plane runs inside a container using a `kindest/node` image and `kubeadm`; the plugin installs the cluster network and manages its local kubeconfig entry.

The plugin is experimental. Its subcommands, defaults, and options may change between releases. Confirm the installed surface before acting:

```sh
container help
container k8s --help
container k8s <subcommand> --help
```

`container help` discovers plugins, while the dynamically dispatched plugin owns `container k8s --help`. Do not assume `container help k8s` follows the core-command help path.

## Lifecycle

### Create

```sh
container k8s create --name <cluster> --cpus <cpus> --memory <memory>
```

Creation can pull the node image, create and start the cluster container, initialize the control-plane and network, and merge cluster credentials into `~/.kube/config`. The default cluster name is version-dependent; use an explicit name when identity matters.

The optional `--rm` flag removes the cluster container after it stops. Use it only when that disposable lifecycle is intended.

### List and restart

```sh
container k8s list
container k8s start --name <cluster>
```

Starting a stopped cluster can refresh its kubeconfig entry because its container IP may change.

### Load a local image

```sh
container k8s load-image --name <cluster> <image>:<tag>
```

This exports the selected local image and imports it into the cluster's containerd `k8s.io` namespace. Use `--platform <os/arch[/variant]>` only when the local image contains multiple platform variants and the target variant is known.

### Write kubeconfig

```sh
container k8s write-config --name <cluster> --kubeconfig <path>
```

Without an alternate path, the plugin uses `~/.kube/config`. Treat this as a user-file mutation: resolve the target, preserve unrelated contexts, and report which file was changed. The resulting kubeconfig can be consumed by standard Kubernetes clients, but this skill does not install or run those clients.

### Delete

```sh
container k8s delete --name <cluster>
```

Deletion stops and removes the cluster container and removes its entry from the default `~/.kube/config`. It does not guarantee cleanup of alternate kubeconfig files written earlier. Confirm the exact cluster name, inspect `container k8s list`, and report alternate files that may remain. Do not broaden deletion to unrelated containers, images, networks, volumes, or kubeconfig contexts.

## Third-Party CLI Boundary

`kubectl`, Helm, Kustomize, and other Kubernetes clients are outside this skill's execution contract.

- Do not install, authenticate, configure, or execute them.
- Do not make them a prerequisite for `container k8s` create, start, list, load-image, write-config, or delete operations.
- Mention them only to explain that kubeconfig is a handoff surface or that workload management begins outside this skill.
- If the task requires applying manifests, managing pods or services, switching contexts, inspecting workload logs, port forwarding, packaging charts, or other client operations, report that boundary and route the work to a separately authorized owner.

## Operational Limits

- Use the plugin for local development, not production orchestration.
- Plugin availability depends on the installed Apple container release and package contents.
- Creating a cluster may download a large node image and allocate CPU, memory, storage, networking, and kubeconfig state.
- Loading an image copies it into the cluster's containerd store; deleting the source image does not imply the cluster copy was removed.
- Version-specific defaults such as Kubernetes node image tags must come from current plugin help rather than this reference.

## Authoritative Sources

- [Apple container command reference](https://github.com/apple/container/blob/main/docs/command-reference.md#kubernetes-cluster-management)
- [Apple container releases](https://github.com/apple/container/releases)
