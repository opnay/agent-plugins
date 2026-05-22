# Session Records

Use session records to preserve operational continuity across compaction, interruptions, status requests, and next-flow routing.

## Files

Create records under `.agents/sessions/{YYYYMMDD}/`.

- `000-plan.md`: date-level routing context, active flow pointer, request history, compact flow index, planned current or future sequence, completed summaries, explicit turn-end availability, and active date-level risks.
- `{count-pad3}-{eng-lower-slug}.md`: one active flow contract and its execution, verification, report, next-flow options, and residual risk.
- `000-self-drive.md`: optional self-drive sequence state. Create it only while self-drive is active.

Use lowercase English slugs and zero-padded counters, for example `001-runtime-authoring.md`.

## Update Rules

Update records incrementally after each phase. A flow record must not wait until the flow is complete before it exists.

Create a new flow record when the active flow boundary changes. Update the existing flow record only when the same flow continues or when repairing that record's own Continuity Guard before reporting.

Keep `000-plan.md` compact. Do not duplicate detailed scope, evidence, verification, residual risk, or self-drive sequence details there.

Keep `Flow Index` and `Completed Flow Summaries` to one compact line per flow. Do not delete completed summaries during normal continuation.

## Flow Record Minimum

Each flow record needs these sections:

- `Flow Contract`
- `Optional Risky Actions`
- `Execution Log`
- `Verification`
- `Report`
- `Next Flow Options`
- `Residual Risk`

The Continuity Guard must show:

- turn-gate active state
- question-routing mode
- user explicit stop
- terminal summary allowed
- required next action
- last refreshed phase
- confirmed closure
- closure source message
- closure recorded phase
- pending question state
- pending question id or summary
- superseded question id or summary
- verification status
- continuity note

Use `verification_status: not-started` before verification, `requested` after a verifier request but before a result, or `pass`, `fail`, `blocked`, `insufficient` after a result. `not-started` and `requested` are progress states, not successful results.

## Raw Request Handling

When preserving raw user text, keep it separate from interpretation or summary. Do not normalize, translate, soften, merge, correct, or infer missing words inside the raw request field. Add interpretation separately when needed.

## Recovery

Handle recovery cases explicitly:

- not-yet-created plan: create it from `templates/plan-template.md`
- not-yet-created flow: create it from `templates/flow-record-template.md`
- unexpectedly missing active record: route blocker recovery or ask the user
- inaccessible active record: report blocker until access is restored or the user selects recovery
- stale closure state: reset closure authority to open continuity and record the recovery
- stale self-drive sidecar: if plan says self-drive is inactive, treat sidecar state as historical
- stale routing mismatch: reconcile from the latest source or ask

Do not silently reconstruct an unexpectedly missing active record.

Read-only requests usually forbid changing the target artifact, not session records. Only skip record writes when the user forbids all writes or explicitly forbids record creation or updates.
