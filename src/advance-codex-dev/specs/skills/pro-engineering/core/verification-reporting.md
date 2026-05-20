# Verification and Reporting Contract

## 책임

이 문서는 검증 선택과 최종 보고 기준을 소유합니다.
`pro-engineering`은 "수정했다"가 아니라 어떤 계약을 어떤 증거로 만족했는지 설명해야 합니다.

## 검증 선택

- 먼저 가장 좁고 의미 있는 검증을 실행합니다.
- shared behavior, user-facing workflow, regression risk가 있으면 테스트 추가 또는 갱신을 고려합니다.
- 핵심 하네스 로직, public API, cross-module contract가 바뀌면 대표 integration 또는 end-to-end 경로를 추가로 확인합니다.
- 설정만 바뀐 경우에는 parse/load 검증과 영향을 받는 대표 경로를 확인합니다.
- fixture나 sample data가 바뀐 경우에는 그 fixture에 의존하는 시나리오를 확인합니다.

## 검증 실패 처리

- 실패 원인을 이해하지 못한 retry로 flaky 동작을 덮지 않습니다.
- failure가 infra, harness, assertion, product behavior 중 어디에 속하는지 분리합니다.
- 재현할 수 없는 실패는 어디까지 확인했고 무엇이 불확실한지 기록합니다.
- 검증이 sandbox, 권한, 네트워크, 외부 상태 때문에 막히면 그 원인과 대체 확인을 분리해서 보고합니다.

## 완료 판정

작업은 다음 기준을 만족할 때 완료로 보고할 수 있습니다.

- 원래 증상과 기대 동작이 설명되어 있다.
- 변경이 확인된 원인 또는 명시된 계약과 연결된다.
- 수행한 검증이 변경 위험도에 비례한다.
- 수행하지 못한 검증과 남은 리스크가 숨겨지지 않았다.
- 최종 상태가 현재 파일과 명령 결과 기준으로 보고된다.

## 보고 형식

보고는 현재 상태만 담아야 합니다.
이전 결정이나 폐기된 시도는 결과 이해에 필요할 때만 짧게 언급합니다.

하네스나 구현 작업처럼 변경 범위가 있는 경우 다음 항목을 포함합니다.

- `Scope handled`
- `Files changed`
- `Verification`
- `Residual risk`

작은 작업은 같은 정보를 짧은 문단으로 압축할 수 있습니다.
어떤 형식이든 검증과 불확실성은 생략하지 않습니다.
