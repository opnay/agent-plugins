# Response Optimization

Make responses compact without weakening meaning, safety, verification, required format, language, or tone.

## Priority

1. Correctness
2. User-requested or repository-required format
3. Safety and approval clarity
4. Verification status
5. Concision

Concision never outranks the first four items. A shorter but more ambiguous answer fails.

## Preserve

- Direct answer or completed result
- Required headings and output sections
- Paths, commands, symbols, identifiers, API names, dates, versions, and numbers
- Verification commands, results, failures, skipped checks, and residual risk
- Approval-sensitive details: target, effect, risk, recovery path
- Safety-critical exceptions, caveats, and constraints
- Material judgment calls and required next action
- Active language, tone, formality, honorifics, and local style

## Remove Or Shorten

- Generic openings and closings
- Restatements of the request
- Long setup before the answer
- Vague reassurance and low-value hedging
- Repeated conditions, evidence, and caveats
- Process narration that does not affect the user's next decision
- Multiple sentences making the same point

## Compression Strength

Use the weakest strength that satisfies the request.

- `light`: Default for broad requests such as "shorter", "concise", "no fluff", "짧게", "간결하게", or "군더더기 없이". Remove greetings, repetition, hedging, and self-description while keeping complete natural sentences.
- `standard`: Use for long answers, key points, summaries, "핵심만", or "요약". Cut secondary background; keep result, evidence, next action, and material risk.
- `extreme`: Use only for labels, status-only output, one-line output, maximum compression, or structured sources such as status tables, numeric lists, or fixed headings. Keep safety details, verification, exact names, paths, commands, and numbers.

Generic shorter-response requests do not permit `extreme`. Start with `light`; increase only when output shape or source structure justifies it.

## High-Impact Step-Down

Step down one level for high-impact or failure-sensitive content.

- From `extreme` to `standard` for failed or skipped verification, approval-sensitive action, destructive work, security, privacy, legal, medical, financial, release, publish, commit, push, or version bump.
- From `standard` to `light` when the user needs explanation, teaching, ordered procedure, or context for a safe decision.

After the sensitive detail is clear, return to compact wording.

## Ordering

Prefer:

1. Result
2. Reason or evidence
3. Next action or risk

Example:

- Before: "확인해 보니 이 문제는 설정 파일의 경로가 맞지 않아서 발생했을 가능성이 높아 보입니다."
- After: "설정 파일 경로가 맞지 않습니다. `configPath`가 실제 파일 위치와 다릅니다."

## Grammar

- Merge duplicate sentences before deleting grammar words.
- Keep particles, articles, prepositions, verbs, honorific markers, and other required function words.
- Use fragments only for labels, headings, status values, or requested terse notes.
- Keep explanation sentences grammatical, even under terse headings or status labels.
- Do not make sentences sound broken to save tokens.

## Language And Register

- Preserve the active language and register.
- Keep Korean honorific style when the user or repository requires it.
- Keep another active language's natural professional register.
- Follow user-specified formality, regional wording, or domain style.
- Do not translate or switch languages only to save tokens.
- Do not weaken meaning, safety boundaries, or verification status because a language is harder to shorten.

## Exceptions

Use enough detail for destructive operations, commit, push, PR, release, publish, version bump, security, privacy, legal, medical, financial, failed or insufficient verification, ordered procedures, and user-requested teaching.

## Final Check

- Does the response answer the request?
- Are required sections intact?
- Are verification and residual risk visible?
- Are exact names, paths, commands, dates, versions, and numbers intact?
- Is language and register preserved?
- Is the strength appropriate?
- Is the response grammatical and unambiguous?
