# turn-gate intent scenarios

이 폴더는 `turn-gate`가 의도한 flow boundary 의미를 유지하는지 확인하기 위한 intent 시나리오를 보관합니다.

이 시나리오들은 그 자체로 runtime instruction이 아니라 spec-side fixture입니다.
skill 문구를 재생성하거나 flow planning 동작을 바꾸는 경우 평가 입력으로 사용합니다.

## 시나리오 계약

각 시나리오는 다음 항목을 기록해야 합니다.

- 사용자 메시지
- 기대하는 `operational-preparation` flow 동작
- 기대하는 `change-unit` planned flows
- flow가 아니어야 하는 항목
- 수용 신호

`operational-preparation`은 사용자 메시지 intake, intent/scope 정렬, approval-boundary 계획, planned-flow-list 설계가 session 또는 plan artifact를 만들 때의 운영 flow입니다.
`change-unit`은 검토 가능한 코드, 문서, fixture, config, release-surface 변경을 소유하는 실행 flow입니다.
