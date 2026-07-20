# Task And Dispatch Contract

Use this reference before spawning subagents.

## Task Packet

```text
Objective:
- <exact question or implementation goal>

Scope:
- Include: <directories, modules, files, or feature area>
- Exclude: <out-of-scope items>

Access mode:
- read-only | write-enabled

Ownership:
- none | <exact writable files or directories>

Inputs:
- User goal: <required outcome>
- Required artifacts: <code, logs, tests, contracts, or data>

Constraints:
- Do not spawn subagents.
- Do not expand beyond Scope.
- Do not edit outside Ownership.
- Preserve <behavior, API, security, performance, compatibility>.
- Commands allowed: <commands or none>
- Shared-state limits: <ports, DBs, temp dirs, lockfiles, generated files>

Deliverable:
- <specific result the main agent can integrate>

Evidence:
- Cite paths, symbols, commands, and summarized outputs.
- Separate confirmed facts from inference.

Completion criteria:
- <observable completed, blocked, or inconclusive condition>
```

## Packet Checks

- Name real scope; avoid `the whole repo` unless the lane owns a distinct lens.
- Use `Ownership: none` for read-only work.
- Assign one writer per writable file.
- Preserve the original goal and explicit user constraints.
- Give enough inputs to start without another unfinished lane.
- Do not request raw logs, work diaries, or hidden reasoning.

## DispatchManifest

```yaml
mode: PARALLEL_READ | PARALLEL_WRITE
assignments:
  - agent_id: <spawn result id>
    role: explorer | worker | default
    task_packet: <complete packet>
ownership:
  agent_owned: <per-agent writable surfaces or none>
  main_owned: <shared contracts, files, and integration surfaces>
required_results:
  - <agent id>
main_owned_work:
  - <decision, shared edit, or integration task>
follow_up_used: false
whole_result_verification:
  - <command or evidence check>
```

Do not fabricate agent IDs. Do not hand off an incomplete manifest.
