# Delegation Rubric

## Independent Workstream Test

Treat a workstream as independent only when it has:

- a separate question, module, test bundle, platform, runtime, log slice, review lens, or implementation ownership
- enough inputs to start now
- a deliverable the main agent can judge without waiting for another unfinished lane
- a validation path or evidence source

Do not treat file count, package count, or broad wording as independence by itself.

## Parallelism Benefit Test

Parallelism is worthwhile when at least one benefit is concrete:

- time: independent exploration or verification can run while the main agent works elsewhere
- quality: separate context reduces missed issues across modules or review lenses
- comparison: options can be evaluated against the same criteria
- isolation: logs, tests, or modules can be investigated without mixing hypotheses

Choose DIRECT when the split mostly creates coordination, merge, or re-check work.

## Shared-State And Write Conflict Test

Before spawning, list possible shared surfaces:

- files and directories
- public interfaces, schemas, API contracts
- config and lockfiles
- generated output
- databases, ports, temp directories, build folders
- fixtures, emulators, external accounts
- current working tree state

If a shared surface is writable or mutable and cannot be isolated, prefer PARALLEL_READ or DIRECT.

## Decision Tree

1. Did the user forbid subagents?
   - yes: DIRECT
2. Are there two or more meaningful independent workstreams?
   - no: DIRECT
3. Can each workstream get a clear scope, output, and evidence requirement?
   - no: DIRECT
4. Can the workstreams start without each other's unfinished results?
   - no: DIRECT or PARALLEL_READ before later sequential work
5. Are writes required now?
   - no: PARALLEL_READ
6. Are writable file sets fully disjoint and shared contracts fixed?
   - yes: PARALLEL_WRITE
   - no: PARALLEL_READ, then main-agent integration

When unsure between PARALLEL_READ and PARALLEL_WRITE, choose PARALLEL_READ.

## Recommended Agent Count

| Task shape | Default count | Notes |
| --- | ---: | --- |
| Two independent modules, tests, or review lenses | 2 | Keep scopes separate. |
| Security/correctness/tests/performance review | 3-4 | Use perspectives only when outputs differ. |
| Three independent test bundles | 3 | Investigate in parallel; write only after overlap check. |
| Migration across API, batch, mobile, admin | 3-4 | Cap at 4 unless user asks. |
| Technical option comparison | 2-3 | Assign options or criteria. |
| Whole-repo lint failures | 2-3 only if causes split by package | Use DIRECT for shared config cause. |

Do not create agents to fill a quota.

## Boundary Cases

### Whole-repo lint failures

Use PARALLEL_READ if independent packages have different failure patterns. Use PARALLEL_WRITE only when file ownership is disjoint and no shared lint config, type root, or lockfile change is needed. Use DIRECT when one shared config or type error drives all failures.

### Frontend, API, Data Feature

Use PARALLEL_READ until the shared API contract is fixed. Then assign disjoint ownership only if the contract is stable. If the interface is still changing, implement sequentially.

### Dependency Upgrade

Parallelize compatibility research by package or runtime. Keep lockfile, package manager install, core config, and full verification with the main agent or one writer.

### Perspective Review On Same Code

Parallel review can be valid even on the same files when review lenses and deliverables differ, such as security versus performance. It is read-only unless one owner is assigned to implement.
