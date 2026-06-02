# turn-gate 스킬 스펙

## 목적

`turn-gate`는 active Codex turn을 명시적 종료 요청 전까지 열린 상태로 유지하는 wrapper입니다.
메인 플로우는 `flow skill` 그룹과 `next-flow gate` 그룹의 루프로 압축합니다.
`flow skill` 그룹은 `flow skill: interview -> 생략... -> flow skill: handoff`를 포함합니다.
`turn-gate`는 `flow skill` 내부를 재정의하지 않고 `flow` 스킬을 그대로 적용합니다.

## 경계

- 포함: active-turn continuity, `flow skill: handoff` 뒤 next-flow gate, 질문 도구 입력 구체화, self-drive 대체 gate, explicit-stop 라우팅.
- 보조: record, verification, interruption, date 처리는 메인 그래프 노드가 아니라 active turn을 복구하고 안전하게 라우팅하기 위한 지원 계약입니다.
- 제외: flow taxonomy, flow lifecycle, readiness/discovery/ambiguity, flow handoff 의미, workflow planner, commit/push/PR/release/version bump/destructive action 승인.

## 계약 맵

- `intent.md`: 메인 그래프와 종료 플로우 그래프
- `contracts/runtime.md`: wrapper runtime
- `contracts/flow-relationship.md`: `flow` 의존 경계
- `contracts/question-routing.md`: handoff 뒤 질문 라우팅과 질문 복구
- `contracts/self-drive.md`: 질문 도구를 대체하는 prepared sequence gate
- `contracts/session-records.md`: active turn 복구용 record 최소 계약
- `contracts/verification.md`: handoff 전 non-pass 라우팅
- `contracts/interruption.md`: active turn 중 새 메시지 라우팅
- `contracts/date-authority.md`: 날짜가 라우팅을 바꾸는 경우의 기준

## 핵심 계약

- 사용자 메시지는 `turn-gate` wrapper 안의 `flow skill` 그룹으로 진입합니다.
- `turn-gate`는 `flow skill: handoff` 이후 `next-flow gate`를 엽니다.
- `next-flow gate`는 매번 `skill reconfigure` 그룹을 거쳐 `다음 플로우 선택 -> 000-plan.md 업데이트`를 기본 경로로 처리합니다.
- `skill reconfigure` 그룹은 `flow skill: handoff`에서 시작해 세션에서 사용중인 전체 skill 목록을 식별하고, 각 skill 본문을 새로 읽고, 새 active skill set으로 수용하는 과정입니다.
- 일반 모드는 질문 도구로 `다음 플로우 선택`에 진입합니다.
- self-drive 모드는 질문 도구를 대체해 `다음 플로우 선택`에 진입하고, `000-self-drive.md 업데이트`를 거쳐 통합 `000-plan.md 업데이트`로 들어갑니다.
- 질문 도구는 `flow: deep-interview`와 같은 인터뷰 흐름으로 다음 flow 입력을 충분히 구체화합니다.
- `000-plan.md 업데이트`는 선택된 다음 flow 입력, 사용 skill, pending/answered question 상태, next action을 매번 반영합니다.
- `flow skill: interview` 재진입 뒤에는 구체화된 입력을 기준으로 flow design에 필요한 질문을 우선합니다.
- 사용자-facing phase 시작 또는 의미 있는 진행 메시지는 현재 단계 prefix를 사용합니다. `turn-gate`는 `flow`가 산출한 phase prefix를 재정의하지 않고 적용하며, `next-flow gate`에서는 `[next-flow]`를 소유합니다.
- phase prefix는 진행 표시이며, artifact 본문, record 본문, command output summary, 질문 option label에 기계적으로 복사하지 않습니다.
- self-drive는 그래프 노드가 아니라 준비된 sequence gate가 질문 도구를 대체하는 핵심 경로입니다.
- 종료 요청은 `turn-gate / 메인`의 모든 시점에서 감지하며, 종료 페이즈로 이동합니다.
- 종료 페이즈는 `작업 중이던 플로우 정리 -> explicit-stop 기록 - active turn 종료` 순서입니다.
- source-recorded explicit stop이 있을 때만 active turn을 닫습니다.
- 완료, 검증 통과, 커밋, 보고, final-looking 문구, 질문 중단은 explicit stop을 대체하지 않습니다.

## 검토 질문

- `turn-gate`가 `flow` 의미를 재정의하지 않고 wrapper로 적용하는가?
- handoff 뒤 `skill reconfigure`로 세션에서 사용중인 전체 skill 목록을 새 active skill set으로 수용하고, 질문 도구 또는 self-drive로 다음 flow를 선택한 뒤 필요한 기록을 업데이트하고 interview로 돌아오는가?
- 사용자-facing phase/progress 메시지에는 source skill이 소유한 phase prefix를 쓰고, artifact/record/command/question option에는 prefix를 전파하지 않는가?
- self-drive가 명시된 gate 없이 자동 시작되지 않는가?
- 종료 요청이 source-recorded explicit stop으로만 닫히는가?
- record, verification, interruption, date 계약이 메인 그래프 노드로 승격되지 않는가?
