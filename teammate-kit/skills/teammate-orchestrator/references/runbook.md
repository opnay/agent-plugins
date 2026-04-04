# Runbook

## 1. Startup

1. Prepare runtime files.

```bash
python3 scripts/team_bus.py reset --bus ./.tmp/chat.jsonl
: > ./.tmp/chat.dlq.jsonl
: > ./.tmp/chat.checkpoints.json
```

2. Seed initial request.

```bash
python3 scripts/team_bus.py send \
  --bus ./.tmp/chat.jsonl \
  --from main \
  --to main \
  --type route.request \
  --payload '{"text":"@all summarize status","team":["alpha","beta","gamma"]}'
```

3. Start orchestrator.

```bash
python3 scripts/orchestrator.py \
  --bus ./.tmp/chat.jsonl \
  --dlq ./.tmp/chat.dlq.jsonl \
  --checkpoint ./.tmp/chat.checkpoints.json \
  --duration 60 \
  --poll 0.2 \
  --watch-mode auto \
  --spawn-workers \
  --workers alpha,beta,gamma \
  --worker-duration 45 \
  --worker-poll 0.2 \
  --max-hops 8 \
  --pending-timeout 5 \
  --max-attempts 3 \
  --retry-base 0.5
```

4. Confirm summary field `wait_backend` (`watchfiles` or `poll`).

## 2. Watch backend setup

1. Install watch backend in the runtime environment.

```bash
python3 -m pip install watchfiles
```

2. Keep `--watch-mode auto` as baseline for production-like runs.
3. If troubleshooting watcher behavior, run once with `--watch-mode poll` for baseline comparison.
4. Verify orchestrator output includes expected `wait_backend`.

## 3. Operating checks

1. Inspect bus timeline.

```bash
python3 scripts/team_bus.py list --bus ./.tmp/chat.jsonl
```

2. Inspect checkpoint progress.

```bash
python3 scripts/team_bus.py checkpoint-list --checkpoint ./.tmp/chat.checkpoints.json
```

3. Inspect DLQ failures.

```bash
python3 scripts/team_bus.py dlq-list --dlq ./.tmp/chat.dlq.jsonl
```

## 4. Failure handling

### A. Retry exhausted

Symptoms:
- `result.timeout` with `reason=beta.retry_exhausted`
- matching DLQ entry

Actions:
1. Confirm worker health (`worker.status`, `orchestrator.worker_exit`).
2. Fix root cause (worker unavailable, wrong team routing, invalid message).
3. Re-dispatch task from original payload via new `route.request`.

### B. Mention loop risk

Symptoms:
- repeated chat forwarding
- `guard.max_hops`

Actions:
1. Lower `--max-hops`.
2. Remove cyclic mention chains.
3. Prefer structured `mentions` over free-form multi-mention text.

### C. Unknown target

Symptoms:
- `guard.unknown_target`
- `chat.unknown_target` in DLQ

Actions:
1. Align `team` list with intended recipients.
2. Validate mention spelling.
3. Re-send request.

## 5. Compaction and retention

1. Compact when size threshold is exceeded.

```bash
python3 scripts/team_bus.py rotate \
  --bus ./.tmp/chat.jsonl \
  --max-bytes 5242880 \
  --archive-dir ./.tmp/chat.archives
```

2. Verify new bus contains `system.log_compacted` summary.
3. Keep archives for audit/replay and enforce external retention policy.

## 6. Recovery

1. Do not delete checkpoints if replay protection is needed.
2. Restart orchestrator with same bus/checkpoint paths.
3. Verify consumer checkpoint moves forward.
4. If forced replay is required, back up then clear checkpoint file.

## 7. Residual operational risks

- If watch backend is unavailable, orchestration latency falls back to `--poll`.
- Rotation is manual unless scheduled externally.
- Event source authenticity is not cryptographically enforced.
- Checkpoint is single-stream and does not support partition balancing.
