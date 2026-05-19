# phase is not flow

## Intent

phase 이름은 별도 검토 가능한 산출물을 소유하지 않는 한 flow가 아닙니다.

## Scenario

사용자가 planned flow로 `분석`, `작업`, `검증`, `보고`를 나열합니다.

## Expected Behavior

- 이 항목들은 flow가 아니라 하나의 active flow 내부 단계로 분류합니다.
- 실제 flow 후보는 산출물 또는 검토 가능한 변경 단위를 기준으로 다시 설계합니다.

## Forbidden Behavior

- phase checklist를 그대로 sub-flow candidates로 기록하지 않습니다.
