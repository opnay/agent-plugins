# flow 기록 스펙

## 기준

`intent.md`는 `000-plan.md`, flow record, `000-review.md`의 갱신 시점을 그래프 노드에 붙입니다.

## 계약

- `000-plan.md`는 skill reconfigure, 메시지 인터뷰, 플로우 설계에서 갱신될 수 있습니다.
- flow record는 flow별 계약 작성 시점에 만들어지고, 메인 플로우 단계와 handoff condition에서 갱신될 수 있습니다.
- `000-review.md`는 메인 플로우 그룹 이후 handoff condition 직전에 갱신합니다.
- 기록은 실행 권한이 아닙니다.
- 템플릿 파일은 `templates/*`를 유지합니다.

## 갱신 시점

- skill reconfigure: flow entry 또는 post-reporting continuation boundary의 active skill 목록, reread 상태, blocker 여부
- 메시지 인터뷰: `deep-interview` 적용 상태, locked execution brief, `000-plan.md` 갱신 여부
- 플로우 설계: 항목 분류, flow 분해, flow별 계약, 진행 순서
- flow별 계약 작성: 선택된 active flow record 작성
- 메인 플로우: intake, framing, preparation, work, verification, reporting
- post-reporting skill reconfigure: 다음 flow가 있을 때 reporting 직후 active skill 목록 갱신
- 메인 플로우 회고: finding 또는 no-finding
- handoff condition: 완료 상태, 검증 상태, 남은 위험, 다음 intake 조건
