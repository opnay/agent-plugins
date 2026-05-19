# sub-flow candidate is not execution

## Intent

sub-flow candidate 생성은 실행 권한이 아닙니다.

## Scenario

parent flow가 `docs-contract-update`, `runtime-skill-update`, `release-build` 세 후보를 만들었습니다.

## Expected Behavior

- 세 항목은 선택 전까지 sub-flow candidates입니다.
- active flow 전환에는 user-gated next-flow decision 또는 prepared self-drive sequence가 필요합니다.
- candidate 목록에는 handoff condition과 unresolved approval checkpoint가 드러나야 합니다.

## Forbidden Behavior

- 후보를 만들었다는 이유만으로 `docs-contract-update`를 바로 실행하지 않습니다.
- 후보 전체를 하나의 active flow 안에 몰아넣지 않습니다.
