# Delegation Rubric

## Allowlist

Treat a workstream as independently dispatchable only when it has:

- a separate question, module, test bundle, platform, runtime, log slice, review lens, or implementation ownership
- enough inputs to start now
- a deliverable the main agent can judge without another unfinished lane
- an evidence source or validation path

Do not infer independence from file count, package count, complexity, or broad engineering wording.

## Parallel Benefit

Require at least one concrete benefit:

- time: independent work proceeds concurrently
- quality: separate context reduces omissions across modules or review lenses
- comparison: options use one shared evaluation contract
- isolation: logs, tests, or modules can be investigated without mixing hypotheses

Use `DIRECT` when splitting mainly adds coordination, merge, or re-check work.

## Shared-State Check

List shared surfaces before spawn:

- files, directories, interfaces, schemas, and API contracts
- config, lockfiles, generated output, and build folders
- databases, ports, temp directories, fixtures, and emulators
- external accounts and current working-tree state

If writable or mutable shared state cannot be isolated, use `PARALLEL_READ` or `DIRECT`.

## Decision Tree

1. Did the user forbid subagents? If yes, use `DIRECT`.
2. Are there at least two meaningful independent workstreams? If no, use `DIRECT`.
3. Does each have clear scope, output, evidence, and completion criteria? If no, use `DIRECT`.
4. Can each start without another unfinished result? If no, use `DIRECT` or investigate sequentially.
5. Are writes required now? If no, use `PARALLEL_READ`.
6. Are writable file sets disjoint and shared contracts fixed? If yes, use `PARALLEL_WRITE`; otherwise use `PARALLEL_READ`.

## Agent Count

| Task shape | Default | Limit |
| --- | ---: | --- |
| Two independent modules, tests, or lenses | 2 | Keep scopes separate. |
| Security, correctness, tests, performance review | 3-4 | Use only distinct lenses. |
| Three independent test bundles | 3 | Investigate before write assignment. |
| API, batch, mobile, admin migration | 3-4 | Keep shared contracts main-owned. |
| Technical option comparison | 2-3 | Use one evaluation contract. |

Never create agents to fill capacity.

## Boundary Cases

- Whole-repo lint: use `DIRECT` for one shared config cause; use `PARALLEL_READ` only for independent package patterns.
- Cross-layer feature: keep writes sequential until the shared interface is fixed.
- Dependency upgrade: parallelize compatibility research; keep lockfile and root config with one owner.
- Same-code perspective review: parallel read-only lenses are allowed when outputs differ.
