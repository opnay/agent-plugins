# Thinking And Progress Wording

Compress visible pre-action judgment sentences and progress updates. Do not expose private reasoning or long internal thought.

## Preserve

- User intent and next action
- Check, edit, question, approval, or verification reason
- Real uncertainty that affects work
- Verification need, approval boundary, blocker, and scope limit
- Active language and register

## Rules

- Convert weak hedging into direct action when evidence is sufficient.
- State the next action as a verb: check, edit, ask, verify, compare, rerun.
- Keep updates short enough to be useful while work continues.
- Do not mention hidden deliberation, private chain-of-thought, or internal reasoning logs.

## Level Gates

Read levels as `` `light` > `standard` > `extreme` ``.
Each level inherits the previous level and applies only its overrides.

### `light`

Use for ordinary progress updates and pre-action judgment sentences.

- Keep a natural sentence when reason and next action both matter.
- Remove weak hedging and repeated setup.
- Preserve uncertainty, verification need, approval boundary, and scope limit.

### `standard`

Inherits `light`. Use when context is shared or the user asks for key progress only.

- Keep next action plus any uncertainty, verification need, approval boundary, or blocker that changes work.
- Drop background that does not change the user's next decision.
- Use `>` chains for flow or stage order, such as `` `light` > `standard` > `extreme` ``.
- Step down to `light` when the user needs explanation or ordered procedure.

### `extreme`

Inherits `standard`. Use only for status labels, checklists, fixed fields, or explicit status-only requests.

- Prefer `다음: ...`, `검증: ...`, `질문: ...`, `범위: ...`, or `승인 경계: ...`.
- Preserve uncertainty, skipped or failed verification, approval boundaries, blockers, and material risk.
- Step down to `standard` when a terse label would hide why input, approval, or verification is needed.

## High-Impact Step-Down

Step down for failed or skipped verification, approval-sensitive action, destructive work, security, privacy, legal, medical, financial, release, publish, commit, push, PR, or version bump.
