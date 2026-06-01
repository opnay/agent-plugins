# phase stays inside active flow

## Intent

phase 이름은 active flow 내부 단계로 유지합니다. 별도 검토 가능한 산출물을 소유하면 플로우 설계가 새 flow로 분리합니다.

## Scenario

사용자가 planned flow로 `분석`, `작업`, `검증`, `보고`를 나열합니다.

## Expected Behavior

- 플로우 설계는 이 항목들을 하나의 active flow 내부 phase로 분류합니다.
- 실제 flow 구성은 산출물, ownership, completion criteria를 기준으로 다시 설계합니다.

## Boundary Behavior

- phase checklist는 active flow 내부 단계로 기록합니다.
- next main flow identity는 reviewable artifact 또는 change unit 기준으로 정합니다.
