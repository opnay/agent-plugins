# Response Optimization

Compress responses without weakening meaning, safety, verification, required format, language, or tone.

## Preserve

- Direct answer or completed result
- Required headings and output sections
- Paths, commands, symbols, identifiers, API names, dates, versions, and numbers
- Verification commands, results, failures, skipped or blocked checks, and residual risk
- Approval-sensitive details: target, effect, risk, recovery path, and approval status
- Safety-critical caveats and constraints
- Material judgment calls and required next action
- Active language, tone, formality, honorifics, and local style

## Remove Or Shorten

- Generic openings and closings
- Restatements of the request
- Long setup before the answer
- Vague reassurance and low-value hedging
- Repeated conditions, evidence, and caveats
- Process narration that does not affect the user's next decision

## Level Gates

Read levels as `` `light` > `standard` > `extreme` ``.
Each level inherits the previous level and applies only its overrides.

### `light`

Use for broad requests such as "shorter", "concise", "no fluff", "짧게", "간결하게", or "군더더기 없이".

- Keep complete natural sentences.
- Remove greetings, repetition, hedging, and self-description.
- Preserve result, evidence, next action, verification, approval boundary, required format, and residual risk.

### `standard`

Inherits `light`. Use for key points, summaries, long answers, "핵심만", "요약", or partial step-down from unsafe `extreme`.

- Keep main judgment, evidence, next action, blocker, material risk, and residual uncertainty.
- Remove secondary background, nonessential examples, duplicate caveats, and chronological process detail.
- Use `>` chains for flow or stage order, such as `` `light` > `standard` > `extreme` ``.
- Step down to `light` when the user needs teaching, ordered procedure, or safe-decision context.

### `extreme`

Inherits `standard`. Use only for labels, status-only output, one-line output, maximum compression, fixed fields, compact tables, or structured sources.

- Use label/value fields, compact bullets, compact tables, or a single next-action label.
- Keep failed, skipped, blocked, or insufficient verification visible.
- Keep approval status, approval boundary, exact names, paths, commands, dates, versions, numbers, residual risk, and required next action.
- Step down for any detail where `extreme` would hide decision basis, insufficient verification, required approval, or remaining risk.

## High-Impact Step-Down

Step down for failed or skipped verification, approval-sensitive action, destructive work, security, privacy, legal, medical, financial, release, publish, commit, push, PR, or version bump.

## Final Check

- Does the response answer the request?
- Are required sections, verification, and residual risk visible?
- Are exact names, paths, commands, dates, versions, and numbers intact?
- Did the selected level pass its gates?
- Is the response grammatical and unambiguous?
