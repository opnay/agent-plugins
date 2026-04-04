---
name: teammate-orchestrator
description: Coordinate a teammate-style multi-agent workflow using an append-only event bus with non-blocking orchestration, mention routing (@agent, @all), consumer checkpoints, retries, DLQ, and log compaction. Use when the user wants main-agent relay instead of direct subagent-to-subagent messaging, needs durable action history, or wants resilient long-running team execution.
---

# Teammate Orchestrator

## Overview

Use this skill to run a team workflow where a main orchestrator routes work through a shared event file instead of waiting on individual agents. Keep all actions as events and recover safely after restarts.

## Workflow

1. Prepare runtime files.
2. Start the orchestrator loop (optionally spawn workers).
3. Append `route.request` events to drive team collaboration.
4. Monitor bus, checkpoints, and DLQ outputs.
5. Compact logs when the bus grows.

## Quick Start

Use these commands from the skill folder:

```bash
python3 scripts/team_bus.py reset --bus ./.tmp/chat.jsonl
: > ./.tmp/chat.dlq.jsonl
: > ./.tmp/chat.checkpoints.json

python3 scripts/team_bus.py send \
  --bus ./.tmp/chat.jsonl \
  --from main \
  --to main \
  --type route.request \
  --payload '{"text":"@all summarize status","team":["alpha","beta","gamma"]}'

python3 scripts/orchestrator.py \
  --bus ./.tmp/chat.jsonl \
  --dlq ./.tmp/chat.dlq.jsonl \
  --checkpoint ./.tmp/chat.checkpoints.json \
  --duration 30 \
  --poll 0.2 \
  --watch-mode auto \
  --spawn-workers \
  --workers alpha,beta,gamma \
  --worker-duration 20 \
  --worker-poll 0.2 \
  --max-hops 8 \
  --pending-timeout 5 \
  --max-attempts 3 \
  --retry-base 0.5
```

## Script Help Snapshot

Use this section to inspect script options quickly without opening the runbook.

Raw help commands:

```bash
python3 scripts/team_bus.py --help
python3 scripts/team_bus.py send --help
python3 scripts/team_bus.py list --help
python3 scripts/team_bus.py dlq-list --help
python3 scripts/team_bus.py checkpoint-list --help
python3 scripts/team_bus.py reset --help
python3 scripts/team_bus.py rotate --help
python3 scripts/team_bus.py worker --help
python3 scripts/orchestrator.py --help
```

`scripts/team_bus.py` command options:

- `send`
  - required: `--bus`, `--from`, `--to`, `--type`
  - optional: `--payload` (default: `{}`; JSON object string)
- `list`
  - required: `--bus`
- `dlq-list`
  - required: `--dlq`
- `checkpoint-list`
  - required: `--checkpoint`
- `reset`
  - required: `--bus`
- `rotate`
  - required: `--bus`, `--max-bytes`
  - optional: `--archive-dir` (default derived from bus path), `--no-compress`
- `worker`
  - required: `--name`, `--bus`
  - optional:
    - `--duration` (default: `20.0`)
    - `--poll` (default: `0.2`)
    - `--max-hops` (default: `8`)
    - `--pending-timeout` (default: `5.0`)
    - `--max-attempts` (default: `3`)
    - `--retry-base` (default: `0.5`)
    - `--dlq` (default derived from bus path when omitted)
    - `--checkpoint` (default derived from bus path when omitted)
    - `--checkpoint-consumer` (default: `worker:<name>`)
    - `--no-status-events` (disable `worker.status` start/stop events)

`scripts/orchestrator.py` options:

- required: `--bus`, `--dlq`, `--checkpoint`
- runtime:
  - `--duration` (default: `30.0`)
  - `--poll` (default: `0.5`)
  - `--watch-mode {auto,watch,poll}` (default: `auto`)
  - `--verbose`
- rotation:
  - `--rotate-max-bytes` (default: `0`, disabled when `0`)
  - `--rotate-archive-dir` (default derived from bus path)
- worker control:
  - `--workers` (default: `alpha,beta,gamma`)
  - `--spawn-workers`
  - `--worker-duration` (default: `20.0`)
  - `--worker-poll` (default: `0.2`)
  - `--max-hops` (default: `8`)
  - `--pending-timeout` (default: `5.0`)
  - `--max-attempts` (default: `3`)
  - `--retry-base` (default: `0.5`)

Runtime value guards (applied internally):

- `duration >= 1.0`
- `poll >= 0.05`
- `worker-duration >= 1.0`
- `worker-poll >= 0.05`
- `max-hops >= 1`
- `pending-timeout >= 0.1`
- `max-attempts >= 1`
- `retry-base >= 0.0`

## Routing Rules

- Use `route.request` events addressed to `main` for orchestrator entry.
- Put collaboration intent in payload text with mentions, for example `@alpha`, `@beta`, `@all`.
- Provide `team` in payload for explicit routing boundaries.
- Prefer structured mentions in payload (`mentions`) when deterministic parsing is required.

## Reliability Rules

- Runtime requires a Unix-like OS (Linux/macOS) because file locking uses `fcntl`.
- Enable checkpoints for each consumer to prevent duplicate handling after restart.
- Use `--watch-mode auto` for immediate wake-up on file changes with polling fallback.
- For strict event-driven wake-up, install `watchfiles` in the runtime.
- Keep retries bounded with `--max-attempts` and `--retry-base`.
- Treat `result.timeout` and DLQ entries as first-class failures requiring operator action.
- Keep `max_hops` low enough to prevent mention loops.

## Watch Backend Setup

Install watch backend in the same Python environment:

```bash
python3 -m pip install watchfiles
```

Mode guidance:

- `--watch-mode auto`: recommended default. Tries watch backend first, then falls back to polling.
- `--watch-mode watch`: forces watch attempt first, still falls back to polling if backend is unavailable.
- `--watch-mode poll`: disables file watch and uses polling only.

Operational check:

- Inspect orchestrator summary `wait_backend` (`watchfiles` or `poll`).

## Monitoring and Maintenance

Read current runtime state:

```bash
python3 scripts/team_bus.py list --bus ./.tmp/chat.jsonl
python3 scripts/team_bus.py dlq-list --dlq ./.tmp/chat.dlq.jsonl
python3 scripts/team_bus.py checkpoint-list --checkpoint ./.tmp/chat.checkpoints.json
```

Compact and archive logs:

```bash
python3 scripts/team_bus.py rotate \
  --bus ./.tmp/chat.jsonl \
  --max-bytes 5242880 \
  --archive-dir ./.tmp/chat.archives
```

## References

- Read [references/event-schema.md](references/event-schema.md) for event types and payload contracts.
- Read [references/runbook.md](references/runbook.md) for operational procedures and failure handling.

## Scripts

- `scripts/team_bus.py`: event bus engine (routing, retries, DLQ, checkpoints, compaction).
- `scripts/orchestrator.py`: non-blocking main orchestration loop with worker process polling.
