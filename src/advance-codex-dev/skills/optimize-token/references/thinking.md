# Thinking And Progress Wording

Use these rules to make pre-action judgment sentences and progress updates compact, actionable, and safe. Do not expose private reasoning or write long explanations of internal thought.

## Scope

This reference covers visible progress wording and compact pre-action planning notes before or during work. It guides what to write or keep brief; it does not ask you to reveal private reasoning.

- deciding what the user wants
- naming the next action
- stating why a check, edit, question, or approval is needed
- reporting short progress while work continues
- preserving uncertainty, verification needs, and approval boundaries

Final response compression is owned by `references/response.md`.

## Rules

- Convert weak hedging into direct action when the evidence is sufficient.
- Keep real uncertainty visible when it affects the next action.
- Keep the sentence about the decision, not the whole reasoning path.
- State the next action as a verb: check, edit, ask, verify, compare, rerun.
- Preserve user intent, next action, verification need, approval boundary, and scope limits.
- Keep the active language and register. If Korean honorific style is active, use it naturally.
- Make progress updates short enough to be useful while work continues.
- Do not mention hidden deliberation, private chain-of-thought, or internal reasoning logs.

## Patterns

- Intent plus action: `사용자가 설명을 원하므로, 답변 후 구현 의향을 확인합니다.`
- File check: `파일 수정 가능성이 있으므로 관련 파일을 먼저 확인합니다.`
- Verification: `변경 후 테스트로 해결 여부를 확인합니다.`
- Uncertainty: `sandbox 문제인지 코드 문제인지 불확실하므로, 로그와 재현 조건을 먼저 확인합니다.`
- Approval boundary: `외부 네트워크 접근이 필요할 수 있으므로, 실패하면 승인 후 재실행합니다.`
- Scope limit: `관련 없는 파일은 제외하고, 요청 범위의 파일만 확인하고 수정합니다.`
- Question routing: `구현 방향이 사용자 선택에 따라 달라지므로, 먼저 방향을 확인합니다.`

## Avoid

- "It seems like the user may want..." when the request is clear.
- "I think it would probably be good to..." when the next step is required.
- Long explanations of why a normal file read, edit, or test is being done.
- Progress updates that repeat the user's request without saying what changed.
- Over-compressed updates that hide failed verification, skipped checks, risk, or approval needs.

## Check

- Is the user's intent still clear?
- Is the next action explicit?
- Did real uncertainty remain visible?
- Did approval-sensitive or verification-sensitive detail survive?
- Is the sentence short without sounding broken?
