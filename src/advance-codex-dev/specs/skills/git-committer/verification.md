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

- staged 검증을 실행할 수 없으면 이유와 residual risk를 기록하고 commit을 막습니다.
- supporting check를 실행할 수 없으면 이유와 residual risk를 기록하고 task risk가 허용할 때만 계속합니다.
- 사용자가 skip을 승인했거나 check가 task risk에 비례하지 않으면 근거와 residual risk를 명시합니다.
- skip은 pass가 아니며, commit 가능 여부 판단에 별도 risk로 남깁니다.

## Message File Verification

- commit 전에 메시지 파일의 정확한 경로가 이번 gate의 allocator에서 왔고 예상 template과 file type에 맞는지 확인합니다.
- 사용자 제공 경로, 임의 저장 경로, allocator provenance를 확인할 수 없는 경로는 message input과 cleanup에 사용하지 않습니다.
- allocation은 exit 0, 단일 반환 경로, 예상 template, regular non-symlink file을 모두 확인해야 통과합니다. nonzero exit에 경로가 없으면 cleanup 대상이 없음을 기록하고 commit을 막습니다. 경로가 반환됐다면 검증된 정확한 파일만 cleanup합니다.
- readback으로 subject, blank line, bullet body, literal `\n`, heredoc/EOF marker, 의도하지 않은 shell text를 확인합니다.
- 파일 생성 또는 readback이 실패하면 commit을 실행하지 않고, 생성된 파일이 있으면 정리합니다.
- commit 시도 후에는 성공과 실패 모두 cleanup 직전에 allocator provenance, 예상 template, regular non-symlink file을 재검증합니다. 일치할 때만 정확한 메시지 파일을 삭제하고 `! -e`와 `! -L`을 각각 확인해 file 또는 symlink directory entry가 남지 않았음을 검증합니다.
- cleanup 직전 재검증이 실패하면 파일을 삭제하지 않고 남은 정확한 경로와 residual risk를 보고합니다.
- commit 전 중단도 cleanup 대상입니다. cleanup 실패는 남은 정확한 경로와 residual risk를 별도 실패로 보고합니다.
- 강제 interruption으로 cleanup을 실행하지 못한 경우 재개 시 allocator provenance, 예상 template, file type이 검증된 정확한 경로의 cleanup과 부재 확인을 새 commit 시도보다 먼저 수행합니다.

## Post-Commit Confirmation

- 메시지 파일 cleanup을 시도하고 결과를 확인한 뒤 post-commit confirmation을 수행합니다.
- commit 후 최신 commit hash, subject, message shape를 확인합니다.
- working tree 상태를 확인해 남은 unrelated 또는 unstaged changes를 보고합니다.
