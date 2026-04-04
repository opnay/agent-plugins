# Event Schema Reference

## 1. Base event envelope

All bus events follow this shape:

```json
{
  "id": "uuid",
  "seq": 1,
  "ts": "2023-02-24T03:22:04.551440+00:00",
  "from": "main|orchestrator|alpha|beta|gamma|system",
  "to": "main|orchestrator|alpha|beta|gamma",
  "type": "event.type",
  "payload": {}
}
```

- `id`: unique event identifier.
- `seq`: global monotonically increasing sequence.
- `ts`: ISO-8601 timestamp.
- `from`/`to`: sender and logical recipient.
- `type`: event category.
- `payload`: event-specific body.

## 2. Common event types

### route.request

Dispatch request for orchestrator.

```json
{
  "from": "main",
  "to": "main",
  "type": "route.request",
  "payload": {
    "text": "@all summarize current status",
    "team": ["alpha", "beta", "gamma"],
    "mentions": ["all"]
  }
}
```

### chat.message

Worker-targeted collaboration message.

```json
{
  "from": "orchestrator",
  "to": "alpha",
  "type": "chat.message",
  "payload": {
    "text": "summarize current status",
    "origin": "main",
    "hops": 0,
    "team": ["alpha", "beta", "gamma"]
  }
}
```

### chat.reply

Worker reply back to origin.

```json
{
  "from": "alpha",
  "to": "main",
  "type": "chat.reply",
  "payload": {
    "text": "alpha handled: summarize current status",
    "hops": 0
  }
}
```

### verify.add / verify.result

Alpha-beta calculation verification protocol.

```json
{
  "from": "alpha",
  "to": "beta",
  "type": "verify.add",
  "payload": {
    "correlation_id": "...",
    "a": 8,
    "b": 9,
    "attempt": 1
  }
}
```

```json
{
  "from": "beta",
  "to": "alpha",
  "type": "verify.result",
  "payload": {
    "correlation_id": "...",
    "sum": 17
  }
}
```

### result.timeout

Failure terminal event emitted after timeout or retry exhaustion.

```json
{
  "from": "alpha",
  "to": "main",
  "type": "result.timeout",
  "payload": {
    "correlation_id": "...",
    "reason": "beta.timeout|beta.retry_exhausted",
    "attempts": 2,
    "max_attempts": 2
  }
}
```

### worker.status

Worker lifecycle audit record.

```json
{
  "from": "alpha",
  "to": "main",
  "type": "worker.status",
  "payload": {
    "state": "started|stopped",
    "consumer": "worker:alpha"
  }
}
```

### orchestrator.*

Orchestrator audit events.

- `orchestrator.started`
- `orchestrator.dispatched`
- `orchestrator.worker_exit`
- `orchestrator.worker_terminated`
- `orchestrator.stopped`
- `orchestrator.no_target`

### system.log_compacted

Log compaction summary event written after rotation.

```json
{
  "from": "system",
  "to": "main",
  "type": "system.log_compacted",
  "payload": {
    "archived_file": ".../chat-YYYYMMDDTHHMMSSZ.jsonl.gz",
    "events_archived": 1000,
    "compressed": true,
    "previous_size_bytes": 5243901,
    "max_bytes": 5242880
  }
}
```

## 3. Storage files

- Bus: `chat.jsonl`
- DLQ: `chat.dlq.jsonl`
- Checkpoints: `chat.checkpoints.json`
- Sequence metadata: `chat.meta.json`
- Archives: `chat.archives/*.jsonl.gz`

## 4. Checkpoint contract

Checkpoint file shape:

```json
{
  "worker:alpha": {
    "last_seq": 123,
    "last_event_id": "...",
    "updated_at": "..."
  },
  "orchestrator:main": {
    "last_seq": 130,
    "last_event_id": "...",
    "updated_at": "..."
  }
}
```

Use `last_seq` as primary replay boundary and keep `last_event_id` for diagnostics.
