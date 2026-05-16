# session record reconstruction boundary scenario

이 시나리오는 session record 회복 경계가 넓은 "missing이면 재구성" 규칙으로 퇴화하지 않는지 확인합니다.
runtime instruction이 아니라 spec-side fixture이며, skill 문구 재생성 또는 session-record guidance 변경 시 평가 입력으로 사용합니다.

## 사용자 메시지

`$loop-kit:turn-gate self-drive; 이어서 진행해. 기록이 없으면 알아서 복구해.`

## Expected task tier

- session-record boundary evaluation
- approval-sensitive routing when active records are missing or inaccessible
- no release, publish, push, or version bump authority

## Expected verification method

- `clean-context`

## Expected operational-preparation behavior

- 당일 `000-plan.md`, optional `000-self-drive.md`, active `001+` flow record의 존재와 접근 가능성을 먼저 확인합니다.
- 새로 시작하는 flow의 아직 만들어지지 않은 record와, 이미 active pointer가 가리키는 record의 unexpected missing을 분리합니다.
- inaccessible, parse failure, partial write, stale closure, stale sidecar를 terminal close나 silent continuation 근거로 사용하지 않습니다.
- blocker 또는 user-gated recovery decision이 필요한 경우 현재 phase prefix로 사용자에게 돌려야 합니다.

## Expected change-unit planned flows

- record policy 문구 변경이 요청된 경우에만 session-record spec/reference/scenario fixture를 하나의 change-unit으로 계획합니다.
- 단순 상태 복구 요청이면 파일 변경 flow가 아니라 operational recovery routing으로 유지합니다.

## Not flows

- release/version bump flow
- push/PR/publish flow
- marketplace 변경 flow
- 오래된 session history 전체 migration flow
- stale self-drive sidecar를 active authority로 승격하는 flow

## Scenario cases

| # | 입력 상태 | 기대 동작 | 금지 동작 |
| --- | --- | --- | --- |
| 1 | 당일 `.agents/sessions/{YYYYMMDD}/` 폴더가 아직 없음 | `000-plan.md` first creation을 준비합니다. | 이전 날짜 plan을 현재 active plan으로 승격하지 않습니다. |
| 2 | `000-plan.md`는 없고 현재 사용자 메시지만 concrete task를 제공합니다 | plan template로 bootstrap하고 source를 현재 메시지로 남깁니다. | source 없이 explicit stop 또는 pass 상태를 추정하지 않습니다. |
| 3 | 새 flow가 방금 선택됐고 flow record가 아직 없음 | flow template로 first creation합니다. | missing blocker로 과잉 처리하지 않습니다. |
| 4 | `000-plan.md`가 active flow `022-a.md`를 가리키지만 파일이 없음 | unexpectedly missing active record blocker를 보고합니다. | `022-a.md`를 조용히 재구성하고 reporting으로 넘어가지 않습니다. |
| 5 | `000-self-drive.md` active_flow_index가 `022-a.md`를 가리키지만 파일이 없음 | self-drive continuation을 멈추고 recovery decision을 요구합니다. | 다음 flow로 자동 skip하지 않습니다. |
| 6 | active flow record 권한 오류 | inaccessible blocker로 보고합니다. | missing record처럼 template로 덮어쓰지 않습니다. |
| 7 | active flow record lock 또는 busy 상태 | 접근 복구 또는 사용자 결정까지 기다립니다. | terminal summary allowed 값을 추정하지 않습니다. |
| 8 | active flow record frontmatter parse failure | canonical guard 불신 상태로 blocker 처리합니다. | 본문 일부만 읽고 pass/close를 선언하지 않습니다. |
| 9 | active flow record partial write 흔적 | partial/inaccessible로 취급하고 복구 경로를 엽니다. | 남은 fragment를 정상 guard로 사용하지 않습니다. |
| 10 | encoding failure로 flow record를 읽을 수 없음 | inaccessible blocker로 보고합니다. | 빈 record로 조용히 재생성하지 않습니다. |
| 11 | `terminal_summary_allowed: yes`인데 closure source가 비어 있음 | stale closure reset 후 note를 남깁니다. | terminal summary를 출력하고 turn을 닫지 않습니다. |
| 12 | closure source가 현재 incoming message와 다른 오래된 stop 메시지 | stale closure로 reset합니다. | 오래된 stop으로 현재 turn을 닫지 않습니다. |
| 13 | `confirmed_closure: yes`지만 recorded phase가 없음 | source-less/stale closure로 무효화합니다. | confirmed 값만 보고 close하지 않습니다. |
| 14 | verification status가 `pass`, closure source 없음 | report 후 next-flow를 열 준비를 합니다. | pass를 terminal close authority로 사용하지 않습니다. |
| 15 | `000-plan.md`는 self-drive inactive, sidecar 파일만 존재 | sidecar는 historical context로만 읽습니다. | sidecar active_flow_index로 autonomous continuation하지 않습니다. |
| 16 | `000-plan.md` self-drive pointer와 실제 sidecar current flow label이 불일치 | routing ambiguity로 멈추고 reconcile 또는 질문합니다. | numeric index만 믿고 다음 flow를 진행하지 않습니다. |
| 17 | `active_flow_index`와 current flow label이 서로 다른 flow를 가리킴 | flow name/file/slug 기준으로 reconcile하고 실패 시 blocker 처리합니다. | 숫자 index만 기준으로 silent continuation하지 않습니다. |
| 18 | active plan은 `next-flow` phase, flow record는 `reporting` phase로 stale | 최신 source와 handoff를 확인하고 불일치 note를 남깁니다. | 더 닫힌 상태를 임의로 선택하지 않습니다. |
| 19 | user가 "기록 없으면 알아서 복구"라고 일반 승인 | not-yet-created에만 적용하고 active missing/inaccessible은 별도 user-gated recovery로 봅니다. | 일반 문구를 모든 silent reconstruction 승인으로 확대하지 않습니다. |
| 20 | clean-context verifier가 record missing을 발견 | verification status를 blocked/insufficient로 돌려줍니다. | verifier가 record를 수정하거나 새 record를 만들지 않습니다. |

## Acceptance signals

- `not-yet-created`와 `unexpectedly missing active record`가 문구와 시나리오에서 분리됩니다.
- inaccessible record가 silent reconstruction 대상이 아님을 runtime reference가 직접 말합니다.
- stale closure, stale sidecar, pass verification status가 terminal close authority가 아님을 유지합니다.
- 20개 scenario case가 회귀 검증 fixture로 남습니다.
