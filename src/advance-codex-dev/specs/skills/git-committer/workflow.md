# Workflow Contract

## Entry Gate

- 사용자가 실제 commit 실행을 요청했거나, readiness gate가 통과된 뒤 실제 commit 실행이 다음 단계로 명시되어야 합니다.
- commit 실행 승인은 readiness, verification, handoff, session record, prior context에서 추정하지 않습니다.
- commit 승인 범위는 commit에만 적용하며 push, PR, release, publish, version bump로 확장하지 않습니다.

## Commit Preparation

1. 프로젝트의 커밋 준비 단계가 끝났는지 확인합니다.
2. staged diff와 unstaged/untracked 상태를 읽어 커밋할 범위를 선택합니다.
3. 범위가 섞이면 split 또는 restage를 먼저 처리하고, 커밋 범위를 다시 확인합니다.

## Commit Execution Authority

1. 사용자에게 실제 commit 실행 승인이 있는지 확인합니다.
2. 승인 범위가 현재 선택된 staged 범위와 같은지 확인합니다.
3. 승인 범위가 다르거나 불명확하면 commit 실행으로 넘어가지 않습니다.

## Commit Execution

1. staged 검증을 실행합니다.
2. 필요한 보조 확인이 있으면 변경 범위에 맞는 가장 좁은 check를 실행합니다.
3. 실패는 고치고 재검증하거나, 고칠 수 없으면 commit을 막고 실패를 보고합니다.
4. 실행 불가 또는 의도적 skip은 이유와 residual risk를 기록합니다.
5. commit message를 작성하고 final staged diff를 다시 확인합니다.
6. commit을 생성하고 최신 commit metadata와 working tree 상태를 보고합니다.

## Guardrails

- unrelated changes를 조용히 포함하지 않습니다.
- staged 상태를 확인하지 않고 commit하지 않습니다.
- skipped verification을 green으로 보고하지 않습니다.
- commit 완료를 push/PR/release/publish 승인이나 turn closure로 취급하지 않습니다.
