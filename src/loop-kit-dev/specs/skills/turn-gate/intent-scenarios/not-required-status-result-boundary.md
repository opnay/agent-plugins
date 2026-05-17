# not-required status result boundary scenario

이 시나리오는 `Method: not-required`가 result status `pass`를 자동으로 뜻하지 않는다는 점을 확인합니다.
runtime instruction이 아니라 spec-side fixture이며, verification method/status 문구 재생성 또는 routing-only fixture 변경 시 평가 입력으로 사용합니다.

## 사용자 메시지

`$loop-kit:turn-gate; 작업 산출물이 없으면 검증은 생략하고 다음으로 넘어가.`

## Expected task tier

- verification method/status boundary evaluation
- routing-only and blocker-before-work handling
- no release, publish, push, commit, or version bump authority

## Expected verification method

- `not-required` allowed only when there is no work output and the flow stops before artifact/source changes.
- `normal` allowed when record readback or evidence checks are performed.
- `clean-context` required once file/release/multi-file contract changes occur.

## Expected operational-preparation behavior

- Verification method와 result status를 별도 필드/판단으로 기록합니다.
- `Method: not-required`는 "검증할 work output 없음"만 뜻하며, routing outcome complete 여부를 따로 판단합니다.
- routing outcome, required next action, blocker, omission reason, residual uncertainty가 부족하면 `pass`가 아니라 `insufficient` 또는 `blocked`로 남깁니다.
- `verification_status: not-required`처럼 method를 status vocabulary에 넣은 record는 invalid로 보고 repair합니다.

## Expected change-unit planned flows

- not-required method/status fixture 보강이 요청된 경우, intent scenario fixture와 scenario index/change spec을 하나의 change-unit으로 계획합니다.
- runtime wording 변경이 필요하면 별도 evidence와 clean-context verification을 포함합니다.

## Not flows

- `not-required`를 성공 상태로 승격하는 flow
- routing outcome 없이 pass report를 쓰는 flow
- blocker-before-work를 successful completion으로 보고하는 flow
- file change를 `not-required`로 검증 생략하는 flow
- approval-sensitive action을 not-required로 경량화하는 flow
- release/version bump/push/PR/publish flow

## Scenario cases

| # | 입력 상태 | 기대 동작 | 금지 동작 |
| --- | --- | --- | --- |
| 1 | activation-only 요청이 next-flow options까지 정상 기록됨 | `Method: not-required`, `Status: pass`를 둘 다 기록하되 pass 이유는 routing outcome 완료라고 적습니다. | `not-required`라서 pass라고 쓰지 않습니다. |
| 2 | activation-only 요청이지만 next-flow options 누락 | method는 `not-required`일 수 있어도 status는 `insufficient`로 둡니다. | options 없이 pass report를 쓰지 않습니다. |
| 3 | required next action이 비어 있음 | `insufficient`로 기록하고 next-flow record를 보완합니다. | no-output이라는 이유만으로 pass하지 않습니다. |
| 4 | residual uncertainty가 비어 있고 routing 대상도 불명확 | `insufficient`로 두고 uncertainty와 next action을 기록합니다. | 빈 uncertainty를 성공 증거로 보지 않습니다. |
| 5 | blocker-before-work로 session record를 읽을 수 없음 | `Method: not-required`, `Status: blocked`를 기록하고 blocker question/routing으로 갑니다. | blocked 상태를 pass로 바꾸지 않습니다. |
| 6 | blocker-before-work인데 verifier가 실행되지 않음 | verifier 미사용은 가능하지만 blocker status를 유지합니다. | verifier 미사용을 성공으로 해석하지 않습니다. |
| 7 | user-gated clarification이 필요한 ambiguous target | no-output이면 method는 `not-required` 가능, status는 clarification pending에 맞춰 `blocked` 또는 `insufficient`입니다. | clarification 없이 completed/pass 처리하지 않습니다. |
| 8 | flow record access는 가능하지만 Continuity Guard가 stale closure 상태 | `normal` readback으로 stale reset evidence를 남기거나 `insufficient`로 repair합니다. | stale guard를 not-required pass로 덮지 않습니다. |
| 9 | `verification_status: not-required`가 frontmatter에 기록됨 | status vocabulary 오류로 `insufficient` repair가 필요합니다. | method name을 status로 인정하지 않습니다. |
| 10 | Verification section에 Method만 있고 Status 없음 | record incomplete로 `insufficient`입니다. | Method 존재만으로 verification complete 처리하지 않습니다. |
| 11 | Status는 pass지만 Evidence가 비어 있음 | routing evidence를 보강하거나 `insufficient`로 되돌립니다. | evidence 없는 pass를 유지하지 않습니다. |
| 12 | Result는 pass지만 Reason이 "not-required라서"뿐임 | pass rationale을 routing completion으로 고치거나 insufficient 처리합니다. | method를 result reason으로 쓰지 않습니다. |
| 13 | routing-only flow 도중 source/spec 파일을 수정함 | `not-required` 선택을 무효화하고 clean-context verification으로 전환합니다. | 파일 변경을 no-output flow로 취급하지 않습니다. |
| 14 | release surface build가 실행됨 | build artifact/evidence가 생겼으므로 `not-required`를 사용하지 않습니다. | build를 검증 생략 대상으로 두지 않습니다. |
| 15 | approval-sensitive commit 준비를 함 | approval boundary와 verification evidence를 별도로 기록합니다. | commit 준비를 not-required로 통과시키지 않습니다. |
| 16 | explicit stop이 source-recorded되지 않았고 no-output report만 있음 | next-flow를 열거나 endpoint를 기록합니다. | not-required를 terminal closure authority로 사용하지 않습니다. |
| 17 | self-drive status 질문에 답만 하고 continuation handoff를 기록하지 않음 | no-output일 수 있어도 handoff evidence 부족으로 `insufficient`입니다. | status 답변만으로 self-drive pass/advance 처리하지 않습니다. |
| 18 | self-drive mid-sequence input이 blocker를 드러냄 | no artifact output이면 method는 `not-required` 가능, status는 `blocked`입니다. | blocker를 residual risk로만 남기고 다음 flow로 가지 않습니다. |
| 19 | clean-context verifier가 "not-required invalid for file change"라고 fail 보고 | verification status를 fail/insufficient로 두고 repair합니다. | verifier fail을 무시하고 not-required pass로 덮지 않습니다. |
| 20 | prior command evidence가 stale이라 current state를 모름 | `insufficient`로 current evidence를 확보합니다. | stale evidence와 not-required를 결합해 pass하지 않습니다. |

## Acceptance signals

- `Method: not-required`와 result status가 분리되어 fixture에 직접 드러납니다.
- `not-required + pass`, `not-required + insufficient`, `not-required + blocked`가 각각 다른 조건으로 설명됩니다.
- routing outcome, required next action, evidence, reason, residual uncertainty가 pass 판단의 근거임을 확인합니다.
- 파일 변경, release/build, approval-sensitive work는 `not-required`에서 제외됩니다.
- 20개 scenario case가 `not-required == pass` 회귀를 검출합니다.
