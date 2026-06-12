# flow 스킬 스펙

## 기준

- 기준 문서: `intent.md`
- 보존 문서: `intent.md`, `intent-legacy.md`, `intent-scenarios/*`, `templates/*`
- 이 스펙 트리는 `intent.md`의 그래프를 현재 계약으로 풉니다.

## 목적

`flow`는 모든 사용자 메시지를 다음 경로로 처리합니다.

```text
entry skill reconfigure -> 메시지 인터뷰 -> 플로우 설계 -> 메인 플로우 -> 메인 플로우 회고 -> handoff condition
```

skill reconfigure는 flow entry와 post-reporting continuation boundary에서 필요한 skill 본문을 다시 읽고, 메시지 인터뷰는 실행 입력을 잠그고, 플로우 설계는 진행할 흐름을 구성하며, 메인 플로우는 선택된 흐름을 실행합니다.

## 소유 범위

- 메시지 인터뷰
- flow entry와 post-reporting continuation boundary의 skill reconfigure
- 플로우 설계
- 메인 플로우
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
- `skill-reconfigure.md`: flow entry와 post-reporting continuation boundary의 active skill reread
- `flow-design.md`: 플로우 설계
- `main-flow.md`: 메인 플로우
- `main-flow-review.md`: 메인 플로우 회고
- `handoff-condition.md`: handoff condition
- `records.md`: 기록 표면
- `templates/*`: 기록 템플릿 계약

## 핵심 계약

- 모든 사용자 메시지는 entry skill reconfigure로 들어간 뒤 메시지 인터뷰로 들어갑니다.
- skill reconfigure는 현재 flow 또는 다음 main flow에 필요한 active skill 목록을 식별하고, 루프 중 잊혔거나 오래된 skill context를 source reread로 복구합니다.
- 메시지 인터뷰는 초기 의도 스냅샷, alignment risk, high-leverage 질문, 답변 반영, 압력 테스트를 거쳐 locked execution brief를 만듭니다.
- locked execution brief는 목적, 대상, 범위, 비목표, 완료 기준, 검증 기대, 승인 경계, 근거, 해소된 alignment risk, 남은 모호성을 현재 확정 상태로 남깁니다.
- 플로우 설계는 locked execution brief에서 항목 분류, 단일/다중 flow 판단, parent flow 또는 단일 active flow 식별, sub-flow candidate 추출, 후보 pending 상태 표시, flow별 계약 작성, 선택된 active flow record 작성, 진행 순서 정리를 수행합니다.
- 선택된 메인 플로우는 `intake -> framing -> preparation -> work -> verification -> reporting` 순서로 진행합니다.
- `intake`, `framing`, `preparation`, `work`, `verification`, `reporting`은 active flow 내부 phase이며, active flow의 현재 위치와 다음 행동을 드러내는 고정 단위입니다.
- `framing`은 사용자 요구사항을 requirement verification todo로 분해하고, `preparation`은 requirement verification todo와 implementation verification todo를 확정합니다.
- requirement verification은 사용자 요구사항 충족 여부를 primary axis로 검증하고, implementation verification은 구현 정합성을 supporting axis로 검증합니다.
- 다음 flow가 있으면 `reporting` 직후 skill reconfigure를 수행한 뒤 다음 `intake`로 라우팅합니다.
- 메인 플로우 그룹 이후에는 `000-review.md`를 갱신하고 handoff condition을 산출합니다.

## 검토 질문

- 메시지 인터뷰가 locked execution brief를 만들었는가?
- skill reconfigure가 flow entry와 post-reporting continuation boundary에 필요한 skill context를 복구했는가?
- locked execution brief가 목적, 대상, 범위, 비목표, 완료 기준, 검증 기대, 승인 경계, 근거, 해소된 alignment risk, 남은 모호성을 드러내는가?
- 플로우 설계가 진행할 flow 구성을 만들었는가?
- 선택된 메인 플로우가 단계 순서를 유지하는가?
- phase가 active flow의 내부 진행 단위로 유지되고, 독립 flow 후보나 handoff로 흐려지지 않는가?
- requirement verification todo가 행위 수행 여부가 아니라 결과물의 사용자 요구사항 충족 여부를 드러내는가?
- implementation verification todo가 requirement verification을 대체하지 않는가?
- reporting에서 다음 skill reconfigure, 다음 intake, 또는 handoff condition이 드러나는가?
- 메인 플로우 회고가 handoff condition 전에 수행되는가?
- 기록 갱신 시점이 `intent.md` 그래프와 일치하는가?
