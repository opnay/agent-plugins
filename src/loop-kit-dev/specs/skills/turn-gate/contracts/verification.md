# turn-gate verification 계약

## 소유 범위

verification method 선택, result status, non-pass routing.

## method 계약

- `clean-context`: full-history fork가 아니라 bounded read-only verifier packet입니다.
- `normal`: main-thread checks, readback, evidence review, logical counterexample review입니다.
- `not-required`: 별도 verification action이 필요 없다는 뜻이며 reason과 residual uncertainty를 기록해야 합니다.

method는 result status가 아닙니다.

reporting 전에 method를 선택합니다. 선택한 method와 그 이유를 모두 기록합니다.

## result status 계약

- `pass`
- `fail`
- `blocked`
- `insufficient`

`not-required`를 automatic pass로 취급하지 않습니다.

result status는 선택된 verification method의 결과를 설명해야 합니다.

- `pass`: evidence가 flow의 acceptance signal을 지지합니다.
- `fail`: evidence가 flow의 acceptance signal을 충족하지 못함을 보여줍니다.
- `blocked`: user input, approval, access, external state change 없이는 verification 또는 repair를 계속할 수 없습니다.
- `insufficient`: pass/fail을 지지하기에 evidence가 불완전하거나 약합니다.

flow record의 진행 상태는 result status와 분리할 수 있습니다. verification이 아직 실행되지 않은 새 flow는 `verification_status`를 `not-started`로 둘 수 있고, clean-context verifier를 요청했지만 결과가 오기 전에는 `requested`로 둘 수 있습니다. `not-started`와 `requested`는 성공/실패 결과가 아니라 기록용 진행 상태이며, reporting에서 성공 근거로 사용할 수 없습니다.

## clean-context 기본값

파일 변경, release surface 변경, 다중 파일 contract 변경, prior check failure, 사용자 요청 verification/review/QA/commit-readiness, approval-sensitive action에는 기본적으로 `clean-context`를 사용합니다.

verifier packet은 target, user intent, changed files 또는 artifacts, inspect할 checks/evidence, pass/fail criteria, no edit permission, no scope expansion, no destructive/external work, no commit/push/PR/publish/release/version-bump action을 포함해야 합니다.

active `turn-gate`에서 clean-context verifier subagent는 read-only bounded verification에 한해 사용할 수 있습니다. 이 허용은 edit permission, scope expansion, destructive/external work, commit/push/PR/publish/release/version-bump authority를 만들지 않습니다. verifier가 그 경계를 넘어야 하면 user-gated routing으로 돌아갑니다.

clean-context verification은 flow record가 실제 위험에 대해 `normal` 또는 `not-required`가 충분한 이유를 설명할 때만 생략할 수 있습니다. 단순 편의는 충분한 이유가 아닙니다.

documentation-only research artifact가 바뀐 경우에도 파일 변경이면 `clean-context` 기본값을 유지합니다. 다만 verifier packet은 전체 source 재조사가 아니라 변경된 artifact, active session record, 이미 기록된 evidence gap, 필수 heading, conclusion/evidence 일관성, 구현 완료 허위 claim 여부로 좁힙니다.

## non-pass 라우팅

success를 보고하기 전에 non-pass result를 라우팅합니다.

- `fail`: 가장 이른 안전한 repair/work point로 돌아갑니다.
- `insufficient`: evidence를 수집하거나 verification을 보강합니다.
- `blocked`: user-gated blocker routing을 엽니다.

non-pass state는 next-flow continuation이나 terminal close를 승인하지 않습니다.

non-pass routing은 self-drive endpoint exhaustion, release readiness, commit-readiness, next-flow continuation보다 먼저 처리해야 합니다. flow는 non-pass status를 보고할 수 있지만, successful completion이 아니라 required recovery action을 보고해야 합니다.
