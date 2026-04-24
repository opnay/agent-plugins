# Loop Kit

`loop-kit` is a narrow loop plugin built around `turn-gate`.
Its main job is to keep one turn alive until the user asks to end the turn while keeping `analysis -> plan -> work -> verification -> result reporting -> next user response` explicit.

The default entrypoint is `loop-kit-guide`, but the main operational surface is `turn-gate`.
Inside `loop-kit`, users do not call `ralph-loop`, `review-loop`, or readiness loops directly.
Instead, `turn-gate` selects the right internal loop mode for the current phase of work, reads the needed absorbed contract from `skills/turn-gate/references/`, and keeps the turn moving.
That internal mode can be a discovery round such as `deep-interview` when the turn still needs requirement locking before later work modes.
It can also be a broad execution lane such as `autopilot` when the current phase should move from a brief request through implementation, QA, and validation.

`turn-gate` also treats the question tool `request_user_input` and the plan tool `update_plan` as mandatory operating tools.
Choice questions, scope locks, and next-flow reopening must go through `request_user_input`, and active work tracking must go through `update_plan`.

`workflow-kit` remains the SSOT for the broader workflow taxonomy and for the canonical loop-mode contracts that `loop-kit` operationalizes.
`loop-kit` is the narrower packaged surface that exposes the loop-gated runtime behavior.
