# Question Routing

Use this reference for question routing and question recovery after `flow skill: handoff`.

## Open Routing

When `turn-gate` is active, `flow skill: handoff` is not terminal closure.
After every handoff, immediately enter `<gate:next-flow>`: run `skill reconfigure`, keep routing open unless an explicit stop is source-recorded, then update `000-plan.md`.
Do this before terminal-looking reporting, final responses, or flow-only closeout.

<gate:next-flow>

`next-flow gate` paths:

- identify the full session active skill list
- reread each active skill body
- accept the refreshed list as the active skill set
- `request_user_input` answer or user message
- optional prepared self-drive gate
- optional `000-self-drive.md` update when self-drive is active
- update `000-plan.md`
- the same interview flow as `flow: deep-interview`
- reenter `flow skill: interview` with clarified input

</gate:next-flow>

Blocker decisions, approval decisions, and explicit stop are handled by global routing and approval boundaries, not as separate main graph nodes.

Do not treat final-looking wording, status-only answers, compression, verification pass, flow reporting, or a successful `flow skill: handoff` as turn closure.
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

Use `request_user_input` when it is available and the choices are narrow. Label the visible choice surface as `질문 도구: 다음 플로우 선택` when describing the graph.
Prefer two or three mutually exclusive choices.
When the tool is unavailable, ask an active plain-text question and record the required next action.
Always update `000-plan.md` after the question answer, fallback answer, abort state, or pending question state is known.

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
