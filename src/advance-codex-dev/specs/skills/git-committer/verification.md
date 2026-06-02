# Staged Verification Contract

## Staged Verification

- commit 직전 `git status`와 `git diff --staged`로 staged scope를 확인합니다.
- staged diff가 intended change unit과 맞지 않으면 commit하지 않습니다.
- partial staging이 필요한 경우 unrelated changes를 분리한 뒤 다시 staged diff를 확인합니다.

## Supporting Checks

- staged 검증만으로 리스크를 확인할 수 없으면 변경 범위에 맞는 가장 좁은 deterministic check를 실행합니다.
- 예: docs-only는 readback/format 확인, 코드 변경은 lint/typecheck/test/build 중 실제 리스크에 맞는 조합.
- 실패한 check는 수정 후 재실행하거나 blocking issue로 보고합니다.

## Skips And Residual Risk

- 실행할 수 없는 staged 검증 또는 supporting check는 이유를 기록합니다.
- 사용자가 skip을 승인했거나 check가 task risk에 비례하지 않으면 residual risk를 명시합니다.
- skip은 pass가 아니며, commit 가능 여부 판단에 별도 risk로 남깁니다.

## Post-Commit Confirmation

- commit 후 최신 commit hash, subject, message shape를 확인합니다.
- working tree 상태를 확인해 남은 unrelated 또는 unstaged changes를 보고합니다.
