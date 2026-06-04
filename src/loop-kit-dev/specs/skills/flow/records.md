# flow 기록 스펙

## 기준

`intent.md`는 `000-plan.md`, flow record, `000-review.md`의 갱신 시점을 그래프 노드에 붙입니다.

## 계약

- `000-plan.md`는 메시지 인터뷰와 플로우 설계에서 active flow, 후보, 진행 순서, 목적 사슬, 다음 행동을 복구 가능하게 남깁니다.
- flow record는 메인 플로우 각 단계의 입력, 판단, 증거, 검증, 보고, handoff condition을 복구 가능하게 남깁니다.
- `000-review.md`는 메인 플로우 그룹 이후 handoff condition 직전에 갱신합니다.
- 기록은 실행 권한이 아닙니다.
- 기록은 commit, push, PR, release, version bump, destructive action을 승인하지 않습니다.
- 템플릿 파일은 `templates/*`를 유지합니다.

## 갱신 시점

- 메시지 인터뷰: alignment risk, 질문, 답변, pressure test, locked execution brief
- 플로우 설계: active flow, parent flow, sub-flow candidate, phase, handoff, 진행 순서
- 메인 플로우: intake, framing, preparation, work, verification, reporting
- 메인 플로우 회고: finding 또는 no-finding
- handoff condition: 완료 상태, 검증 상태, 남은 위험, 다음 intake 조건
