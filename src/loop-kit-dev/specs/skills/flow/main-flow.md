# flow 메인 플로우 스펙

## 기준 그래프

```text
intake -> framing -> preparation -> work -> verification -> reporting
```

다음 flow가 있으면 `reporting -> intake`로 라우팅합니다.

## 계약

- 메인 플로우는 선택된 active flow의 생애주기입니다.
- 각 단계는 flow record 갱신 시점입니다.
- 단계는 순서를 유지합니다.
- 단계 중 계약이 바뀌면 가장 이른 안전한 지점으로 돌아갑니다.
- 새 산출물 변경이 독립 검토 단위가 되면 플로우 설계로 돌려 새 flow 후보로 분리할 수 있습니다.

## 단계 계약

- `intake`: locked execution brief와 현재 active flow 입력을 확인합니다.
- `framing`: active flow, candidate, phase, handoff 구분과 소유권을 확인합니다.
- `preparation`: scope, non-goals, completion criteria, verification expectation, approval boundary, handoff condition을 작업 진입 가능 상태로 잠급니다.
- `work`: active flow 계약 안에서 작업하거나 답변, 설명, 요약, 상태 결과를 산출합니다.
- `verification`: 산출물이 locked execution brief와 active flow 계약에 맞는지 검증하거나 증거 부족을 기록합니다.
- `reporting`: result, verification, residual risk, 다음 intake 조건 또는 handoff 조건을 보고합니다.

## 산출

- 단계별 flow record 갱신
- 작업 결과
- 검증 결과 또는 증거 부족
- 보고 결과
- 다음 flow intake 조건
