# flow 메인 플로우 스펙

## 기준 그래프

```text
intake -> framing -> preparation -> work -> verification -> reporting
```

다음 flow가 있으면 `reporting -> post-reporting skill reconfigure -> intake`로 라우팅합니다.

## 계약

- 메인 플로우는 선택된 active flow를 단계 순서대로 진행합니다.
- `reporting` 뒤 다음 flow가 있으면 `reporting` 직후 skill reconfigure가 필요한 skill context를 복구합니다.
- post-reporting skill reconfigure가 끝난 뒤 다음 `intake`로 들어갑니다.
- 각 단계는 flow record 갱신 시점입니다.
- 각 단계는 현재 active flow의 phase이며, active flow의 현재 위치와 다음 행동을 드러냅니다.
- `intake`는 locked execution brief와 메인 플로우 입력을 확인합니다.
- `framing`은 현재 단계의 구도와 소유 경계를 확인하고, 사용자 요구사항을 산출물이 충족해야 할 requirement verification todo로 분해합니다.
- `preparation`은 작업 진입 전 계약을 잠그고, requirement verification todo와 implementation verification todo를 확정합니다.
- `work`는 계약 안에서 산출물을 만듭니다.
- 산출물은 파일 변경, 답변, 설명, 요약, 상태 보고, 검증 결과, 또는 사용자가 요청한 다른 결과일 수 있습니다.
- `verification`은 requirement verification todo와 implementation verification todo를 산출물 기준으로 확인합니다.
- `reporting`은 결과, 검증, 남은 위험, 다음 intake 또는 handoff 조건을 남깁니다.
- 사용자에게 보이는 메시지가 특정 phase에서 나오면 `[<phase-name>]` 형식의 메시지 앞 라벨을 붙입니다.
- phase 라벨은 진행 표시이며 산출물, 기록, 명령 요약, 질문 선택지 라벨에 기계적으로 전파하지 않습니다.
- phase가 별도 검토 가능한 산출물, 완료 기준, 승인 경계, handoff 조건을 갖기 시작하면 새 flow나 sub-flow candidate로 다시 분류합니다.

## 산출

- 단계별 flow record 갱신
- 작업 산출물
- 검증 결과
- 보고 결과
- 다음 skill reconfigure, intake 조건, 또는 handoff 조건

## Verification Todo

- requirement verification todo는 사용자 요구사항, 완료 기준, scope, non-goals를 산출물이 충족해야 할 세부 항목으로 나눕니다.
- requirement verification todo는 행위 수행 여부가 아니라 결과물 속 필드, 상태, 케이스, 문구, 동작, 제외 책임 같은 observable requirement를 검증합니다.
- implementation verification todo는 타입, 테스트, 빌드, lint, import path, release surface, 코드베이스 관례처럼 구현 정합성을 검증합니다.
- requirement verification todo는 primary axis이고, implementation verification todo는 supporting axis입니다.
- `verification`은 각 todo를 `pass`, `fail`, `blocked`, `insufficient`, `not-required` 중 하나로 판정합니다.
- 확인할 수 없는 항목은 성공으로 승격하지 않고 `insufficient`로 남깁니다.
- blocker나 승인 민감 작업이 남으면 `blocked` 또는 handoff condition으로 라우팅합니다.
