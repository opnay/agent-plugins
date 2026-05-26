# Response Optimization

Use these rules to make responses compact without weakening meaning, safety, verification, required format, language, or tone.

## Priority

1. Correctness
2. User-requested or repository-required format
3. Safety and approval clarity
4. Verification status
5. Concision

Concision never outranks the first four items. A shorter answer that becomes more ambiguous is a failed optimization.

## Preserve

- Direct answer or completed result
- Required headings and required output sections
- File paths, commands, symbols, identifiers, API names, dates, versions, and numbers
- Verification commands, results, failures, skipped checks, and residual risk
- Approval-sensitive details: target, effect, risk, and recovery path
- Safety-critical exceptions, caveats, and constraints
- Material judgment calls and required next action
- Active language, tone, formality, honorifics, and local style

## Remove Or Shorten

- Generic openings and closings
- Restatements of the user's request
- Long setup before the answer
- Vague reassurance and low-value hedging
- Repeated conditions, repeated evidence, and repeated caveats
- Process narration that does not affect the user's next decision
- Multiple sentences that make the same point

## Compression Strength

Use the weakest strength that satisfies the request.

- `light`: Default for broad requests such as "shorter", "concise", "no fluff", "짧게", "간결하게", or "군더더기 없이". Remove greetings, repeated context, hedging, and self-description while keeping complete natural sentences.
- `standard`: Use when the answer is long or the user asks for key points, summary, "핵심만", or "요약". Cut secondary background and keep the result, evidence, next action, and material risk.
- `dense`: Use only when the user directly asks for labels, status-only output, one-line output, maximum compression, or when the source is already structured, such as a status table, numeric list, or fixed headings. Keep safety details, verification, exact names, paths, commands, and numbers.

Do not treat a generic shorter-response request as permission for `dense`. Start with `light`, then move stronger only when the requested output shape or source structure justifies it.

## High-Impact Step-Down

Step down one level when the content includes high-impact or failure-sensitive information.

- From `dense` to `standard` when the response includes failed verification, skipped verification, approval-sensitive action, destructive operation, security warning, privacy concern, legal caveat, medical caveat, financial caveat, release, publish, commit, push, or version bump.
- From `standard` to `light` when the user needs an explanation, teaching, ordered procedure, or enough context to make a safe decision.

After the sensitive detail is clear, return to compact wording.

## Ordering

Prefer this order:

1. Result
2. Reason or evidence
3. Next action or risk

Example:

- Before: "확인해 보니 이 문제는 설정 파일의 경로가 맞지 않아서 발생했을 가능성이 높아 보입니다."
- After: "설정 파일 경로가 맞지 않습니다. `configPath`가 실제 파일 위치와 다릅니다."

## Grammar

Shorter wording must still read naturally.

- Prefer merging duplicate sentences over deleting grammar words.
- Keep particles, articles, prepositions, verbs, honorific markers, and other required function words.
- Use fragments only for labels, headings, status values, or user-requested terse notes.
- Keep explanation sentences grammatical, even when headings or status values are terse.
- Do not make the sentence sound broken to save tokens.

## Language And Register

Response optimization preserves the active language and register.

- If the user or repository requires Korean honorific style, keep Korean honorific style.
- If another language is active, keep that language's natural professional register.
- If the user specifies formal, informal, regional, or domain-specific wording, keep that constraint.
- Do not translate or switch languages only to save tokens.
- Do not weaken meaning, safety boundaries, or verification status because a language is harder to shorten.

## Exceptions

Use enough detail when the response covers:

- Destructive operations
- Commit, push, pull request, release, publish, or version bump
- Security, privacy, legal, medical, financial, or other high-impact guidance
- Failed or insufficient verification
- Multi-step procedures where order matters
- User-requested explanation or teaching

## Final Check

- Does the response still answer the actual request?
- Did any required section disappear?
- Are verification and residual risk still visible?
- Are exact names, paths, commands, dates, versions, and numbers intact?
- Did the response preserve the active language and register?
- Is the selected compression strength appropriate?
- Is the response still grammatical and natural?
- Could a shorter sentence be misunderstood?
