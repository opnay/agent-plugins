# flow 스킬 스펙

## 기준

- 기준 문서: `intent.md`
- 보존 문서: `intent.md`, `intent-legacy.md`, `intent-scenarios/*`, `templates/*`
- 이 스펙 트리는 `intent.md`의 그래프를 실행 계약으로 풀어 씁니다.

## 목적

`flow`는 모든 사용자 메시지를 같은 경로로 처리합니다.

```text
메시지 인터뷰 -> 플로우 설계 -> 메인 플로우 -> 메인 플로우 회고 -> handoff condition
```

질문, 상태 확인, 설명 요청, 작업 요청은 예외 경로를 만들지 않습니다.
메시지 인터뷰가 충분히 잠긴 실행 요약을 만들면 사용자 질문 없이 플로우 설계로 진행할 수 있습니다.

## 소유 범위

- 메시지 해석과 잠긴 실행 요약
- 진행할 플로우 구성
- 메인 플로우 생애주기
- 메인 플로우 회고
- handoff condition
- `000-plan.md`, flow record, `000-review.md`의 갱신 시점 의미

## 비소유 범위

- 질문 도구 실행 방식
- 활성 턴 유지
- 다음 플로우 질문 라우팅
- self-drive 진행 제어
- 커밋, 푸시, 풀 리퀘스트, 릴리스, 버전 변경 실행 권한

## 문서 맵

- `message-interview.md`: 메시지 인터뷰
- `flow-design.md`: 플로우 설계
- `main-flow.md`: 메인 플로우 생애주기
- `main-flow-review.md`: 메인 플로우 회고
- `handoff-condition.md`: handoff condition
- `records.md`: 기록 표면과 갱신 시점
- `templates/*`: 기록 템플릿 계약. 이번 구조 재작성에서도 유지합니다.

## 핵심 계약

- 모든 사용자 메시지는 메시지 인터뷰로 들어갑니다.
- 메시지 인터뷰는 초기 의도 스냅샷, alignment risk, high-leverage 질문, 답변 반영, 압력 테스트를 거쳐 locked execution brief를 만듭니다.
- 플로우 설계는 locked execution brief에서 항목 분류, 플로우 분해, flow별 계약 작성, 진행 순서 정리를 수행합니다.
- 플로우 설계 결과는 하나 이상의 메인 플로우가 될 수 있습니다.
- 선택된 메인 플로우는 `intake -> framing -> preparation -> work -> verification -> reporting` 순서로 진행합니다.
- 다음 플로우가 있으면 `reporting`에서 다음 `intake`로 라우팅합니다.
- 메인 플로우 그룹 이후에는 `000-review.md`를 갱신하고, 그 뒤 handoff condition을 산출합니다.
- `000-review.md`는 active routing이나 handoff authority로 쓰지 않습니다.

## 검토 질문

- 모든 사용자 메시지가 메시지 인터뷰와 플로우 설계를 거쳤는가?
- 질문 없이 진행한 경우에도 locked execution brief가 남았는가?
- 플로우 설계가 active flow, parent flow, sub-flow candidate, phase, handoff를 구분했는가?
- 각 flow 계약에 scope, non-goals, completion criteria, verification expectation, approval boundary, handoff condition이 있는가?
- 메인 플로우가 `intake -> framing -> preparation -> work -> verification -> reporting` 순서를 유지했는가?
- 다음 플로우가 reporting에서 다음 intake로 연결되는가?
- 메인 플로우 회고가 handoff condition 직전에 수행되는가?
- 기록 갱신 시점이 `intent.md`의 그래프 노드와 일치하는가?
