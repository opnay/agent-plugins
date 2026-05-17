# explicit stop source matching scenario

이 시나리오는 terminal summary가 현재 사용자 메시지 또는 현재 메시지와 명확히 연결된 source-recorded explicit stop에서만 허용되는지 확인합니다.
runtime instruction이 아니라 spec-side fixture이며, explicit stop, stale closure, future endpoint stop, resume/compaction handling 변경 시 평가 입력으로 사용합니다.

## 사용자 메시지

`작업 끝나면 멈춰둬. 지금 상태도 알려줘.`

## Expected task tier

- explicit stop source matching evaluation
- stale/source-less closure handling
- future endpoint stop versus immediate terminal closure distinction
- no release, publish, push, commit, or version bump authority

## Expected verification method

- `clean-context`

## Expected operational-preparation behavior

- 현재 사용자 메시지가 current turn 자체를 끝내려는 explicit stop인지 먼저 구분합니다.
- future endpoint stop, status request, completion report request, commit completion, approval handoff는 current terminal closure가 아닙니다.
- `confirmed_closure`, `terminal_summary_allowed`, `closure_source_message`가 있어도 현재 incoming message와 명확히 연결되지 않으면 stale/source-less closure로 reset합니다.
- compaction/API summary만 있고 exact 또는 clearly attributable stop source가 없으면 terminal close를 허용하지 않습니다.

## Expected change-unit planned flows

- explicit stop source matching fixture/guidance 보강이 요청된 경우, intent scenario fixture와 관련 index/change spec을 하나의 change-unit으로 계획합니다.
- runtime field schema 추가가 필요하면 별도 evidence와 clean-context verification을 포함합니다.

## Not flows

- completion report를 explicit stop으로 승격하는 flow
- future endpoint stop을 즉시 terminal closure로 쓰는 flow
- stale/source-less closure를 현재 closure authority로 쓰는 flow
- compaction summary만으로 terminal close하는 flow
- approval-sensitive handoff를 stop approval로 해석하는 flow
- release/version bump/push/PR/publish flow

## Scenario cases

| # | 입력 상태 | 기대 동작 | 금지 동작 |
| --- | --- | --- | --- |
| 1 | 현재 메시지가 "여기서 끝"이고 flow record에 같은 source가 기록됨 | reporting 후 terminal closure를 허용합니다. | source를 기록하지 않고 닫지 않습니다. |
| 2 | 현재 메시지가 "턴 종료"지만 아직 reporting이 필요함 | closure state를 기록하고 reporting 후 닫습니다. | verification/reporting을 건너뛰지 않습니다. |
| 3 | 현재 메시지가 "작업 끝나면 멈춰둬" | future endpoint 또는 handoff 조건으로 기록합니다. | 즉시 terminal summary로 닫지 않습니다. |
| 4 | 현재 메시지가 "목록 소진되면 멈춰" | endpoint를 갱신하고 현재 flow는 계속합니다. | current explicit stop으로 분류하지 않습니다. |
| 5 | 현재 메시지가 "끝났으면 보고해" | reporting/next-flow로 처리합니다. | 보고 요청을 turn stop으로 보지 않습니다. |
| 6 | 현재 메시지가 "커밋 완료했으면 알려줘" | commit completion reporting으로 처리합니다. | 커밋 완료를 closure authority로 쓰지 않습니다. |
| 7 | 현재 메시지가 status/progress 질문 | 상태를 보고하고 continuation합니다. | status 질문을 terminal closure로 바꾸지 않습니다. |
| 8 | `confirmed_closure: true`, `closure_source_message` 비어 있음 | source-less closure를 reset합니다. | confirmed 값만 보고 닫지 않습니다. |
| 9 | `terminal_summary_allowed: true`, source가 오래된 stop 메시지 | 현재 메시지와 불일치하면 stale로 reset합니다. | 오래된 source로 현재 turn을 닫지 않습니다. |
| 10 | closure source가 compaction summary 문장뿐임 | exact/attributable source 부족으로 terminal close를 금지합니다. | summary만으로 stop source를 확정하지 않습니다. |
| 11 | API resume context에 message id 없이 "user stopped"만 있음 | conservative continuation 또는 blocker로 둡니다. | message id 없는 stop claim으로 닫지 않습니다. |
| 12 | closure source가 현재 메시지의 번역/요약과만 유사함 | 충분히 attributable하지 않으면 stale로 봅니다. | fuzzy similarity만으로 closure를 승인하지 않습니다. |
| 13 | closure source는 현재 메시지와 같지만 recorded phase가 없음 | closure record를 보완하거나 insufficient로 둡니다. | phase 없는 closure를 무조건 pass로 쓰지 않습니다. |
| 14 | current message가 stop과 새 작업을 동시에 요청 | ambiguity를 user-gated clarification으로 풉니다. | stop만 선택하거나 새 작업을 자동 실행하지 않습니다. |
| 15 | current message가 "이번 flow만 멈춰" | turn stop인지 flow pause인지 meaning resolution합니다. | turn 전체 종료로 단정하지 않습니다. |
| 16 | self-drive 중 "다음 항목 전에는 멈춰" | endpoint/handoff 변경으로 기록합니다. | 현재 turn terminal closure로 쓰지 않습니다. |
| 17 | stale closure reset 후 next-flow options가 필요함 | next-flow 또는 self-drive continuation으로 라우팅합니다. | reset 후 terminal summary를 출력하지 않습니다. |
| 18 | blocker 상태에서 user says "그럼 여기서 끝" | blocker report 후 source-recorded closure를 처리합니다. | blocker 기록 없이 닫지 않습니다. |
| 19 | verifier가 source mismatch를 발견 | verification status를 fail/insufficient로 두고 closure reset을 요구합니다. | mismatch를 residual risk로만 남기고 pass하지 않습니다. |
| 20 | closure state와 self-drive endpoint가 충돌 | explicit current stop source가 없으면 endpoint/handoff 기준으로 계속하거나 clarification합니다. | endpoint 충돌을 terminal close authority로 쓰지 않습니다. |

## Acceptance signals

- current explicit stop, future endpoint stop, status/report request, completion notice가 분리됩니다.
- source-less/stale/mismatched closure는 terminal summary authority가 아닙니다.
- compaction/API summary만으로 terminal closure하지 않는 conservative matching rule이 드러납니다.
- ambiguity가 있으면 meaning resolution 또는 blocker routing으로 돌아갑니다.
- 20개 scenario case가 explicit stop source matching 회귀를 검출합니다.
