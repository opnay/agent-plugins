# manager lifecycle spec

## 목적

이 문서는 `manager`의 dispatch packet과 subagent 작업 사이클을 소유합니다.

## Dispatch Packet

각 subagent packet은 다음을 포함합니다.

- `objective`
- `worktree_path`
- `branch`
- `base_or_integration_branch`
- `write_scope`
- `non_goals`
- `verification_required`
- `commit_approval_state`
- `commit_expectation`
- `rebase_target`
- `handoff_output`
- `stop_condition`

## Subagent Lifecycle

- subagent는 task slice 시작 시 새로 생성합니다.
- 하나의 subagent는 하나의 task slice와 하나의 worktree만 담당합니다.
- subagent는 자기 worktree 밖을 수정하지 않습니다.
- subagent는 필요한 context만 읽고 worktree 안에서 작업합니다.
- subagent는 작업 후 slice verification을 수행합니다.
- local commit 권한이 확인된 경우에만 subagent는 git commit을 만듭니다.
- read-only, planning, or verification-only job은 commit 없이 완료할 수 있지만, handoff에는 commit이 없다는 사실과 산출 evidence를 명시합니다.
- 메인 에이전트가 병합 준비를 요청하면 subagent는 메인 에이전트의 현재 integration branch로 rebase합니다.
- rebase conflict는 subagent worktree 안에서 해결합니다.
- rebase 성공 후 subagent는 commit hash, rebase target HEAD, verification result, changed files, residual risk를 보고하고 종료합니다.
- 종료된 subagent는 다음 task에 재사용하지 않습니다.

## Commit Authority

- `manager` workflow는 subagent local commit 권한을 자동으로 만들지 않습니다.
- dispatch 전에 local commit이 현재 승인 경계 안에 있는지 확인합니다.
- local commit authority가 없으면 commit이 필요한 worker를 dispatch하지 않고, 해당 job을 blocked 또는 read-only/planning job으로 유지합니다.
- push, PR, publish, release, version bump, destructive work, external effect는 local commit과 별도 승인 경계입니다.
