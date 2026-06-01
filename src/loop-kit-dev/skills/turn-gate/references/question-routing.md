# Question Routing

Use this reference for `next turn-flow / 메시지 수신` and question recovery after `flow.end`.

## Open Routing

When `turn-gate` is active, a report is not terminal closure.
After `flow.end`, keep routing open unless an explicit stop is source-recorded.

Valid next input paths:

- user message
- self-drive interpretation
- blocker decision
- approval decision
- explicit stop

Do not treat final-looking wording, status-only reporting, compression, or a successful `flow.end` as turn closure.
Closure requires explicit stop.

## Asking

Ask only for the decision needed now when it can change:

- next flow selection
- target, scope, endpoint, or acceptance signal
- approval-sensitive boundary
- verification path
- blocker recovery
- current-flow identity
- whether a pending question has been answered or superseded

Use `request_user_input` when it is available and the choices are narrow.
Prefer two or three mutually exclusive choices.
When the tool is unavailable, ask an active plain-text question and record the required next action.

## Abort Recovery

An aborted, canceled, or interrupted `request_user_input` is not flow completion and is not explicit stop.

Record:

- `terminal_summary_blocked` in flags
- pending question state: `aborted`, `interrupted`, or `superseded`
- pending question id or compact summary when known
- no explicit-stop source unless the user actually stopped the turn

For the next user message:

- If it answers the pending question, continue from that answer.
- If it requests a new flow, mark the pending question `superseded` and prepare the new flow.
- If it asks for status, report active flow, pending question, verification state, and required next action, then reopen routing.
- If it explicitly stops the turn, record the source before closing.

Do not immediately repeat the same question tool call after an abort.
If the next message is ambiguous, ask a smaller clarification instead of guessing.
If a free-form answer does not match a visible option but clearly gives a new task, mark the pending question `superseded` and prepare that flow.
If it selects an option and adds a note, record both the selected answer and the note.

## Blocker Routing

A blocker question or report keeps the turn open unless explicit stop is recorded.
It must say:

- what is blocked
- what evidence was collected
- what decision, access, approval, or external state change is needed
- what work is excluded until the blocker is resolved
