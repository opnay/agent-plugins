# sub-flow candidate pending state

## Intent

sub-flow candidate 생성은 pending flow option을 만듭니다.

## Scenario

플로우 설계가 `docs-contract-update`, `runtime-skill-update`, `release-build` 세 후보를 만들었습니다.

## Expected Behavior

- 세 항목은 선택 전까지 sub-flow candidates입니다.
- next main flow는 user-gated next-flow decision 또는 prepared self-drive sequence가 선택합니다.
- candidate 목록에는 handoff condition과 unresolved approval checkpoint가 드러나야 합니다.

## Boundary Behavior

- `docs-contract-update`는 routing에서 선택될 때 active flow가 됩니다.
- 후보 전체는 각 후보의 contract와 handoff condition을 유지합니다.
