# flow 메인 플로우 스펙

## 기준 그래프

```text
intake -> framing -> preparation -> work -> verification -> reporting
```

다음 flow가 있으면 `reporting -> intake`로 라우팅합니다.

## 계약

- 메인 플로우는 선택된 active flow를 단계 순서대로 진행합니다.
- 각 단계는 flow record 갱신 시점입니다.
- `intake`는 locked execution brief와 메인 플로우 입력을 확인합니다.
- `framing`은 현재 단계의 구도와 소유 경계를 확인합니다.
- `preparation`은 작업 진입 전 계약을 잠급니다.
- `work`는 계약 안에서 산출물을 만듭니다.
- 산출물은 파일 변경, 답변, 설명, 요약, 상태 보고, 검증 결과, 또는 사용자가 요청한 다른 결과일 수 있습니다.
- `verification`은 산출물과 계약의 정합성을 확인합니다.
- `reporting`은 결과, 검증, 남은 위험, 다음 intake 또는 handoff 조건을 남깁니다.

## 산출

- 단계별 flow record 갱신
- 작업 산출물
- 검증 결과
- 보고 결과
- 다음 intake 조건 또는 handoff 조건
