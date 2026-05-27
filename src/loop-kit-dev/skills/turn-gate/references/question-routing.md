# Question Routing

Use this reference when a flow needs user choice, clarification, blocker recovery, next-flow reopening, or recovery after an interrupted question tool call.

## When To Ask

Route through a user question when the answer can change:

- next flow selection
- scope, target, endpoint, or acceptance signal
- approval-sensitive boundary
- verification path
- blocker recovery
- current-flow identity
- whether a pending question has been answered or superseded

Ask only for the decision needed now. Do not bundle unrelated future work just because it is possible.

## Structured Tool

Use `request_user_input` when it is available and the choices are narrow. Keep options tied to the report or blocker that led to the question. Prefer two or three mutually exclusive choices.

If the tool UI cannot include an explicit stop option, still do both:

- mention in the prompt or fallback text that the user can explicitly end the turn
- record explicit turn-end in `Result.next` or temporary next options

## Plain-Text Fallback

When no structured question tool is available, keep the turn active with plain text:

1. State that the structured question tool is unavailable.
2. List the narrow choices.
3. Mark the required next action in the record.
4. Avoid terminal closeout language.

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

Do not immediately repeat the same question tool call after an abort. If the next message is ambiguous, ask a smaller clarification instead of guessing.

## Blocker Routing

A blocker question or report must say:

- what is blocked
- what evidence was collected
- what decision, access, approval, or external state change is needed
- what work is excluded until the blocker is resolved

Blocker routing keeps the turn open unless the user explicitly stops it.
