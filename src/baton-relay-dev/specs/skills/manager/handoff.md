# manager handoff spec

## 목적

이 문서는 `manager`의 commit/rebase handoff gate, prepared commit 회수, integration verification, cleanup을 소유합니다.

## Handoff Gate

메인 에이전트는 subagent handoff를 import하기 전에 다음을 확인합니다.

- subagent가 commit을 만들었는가
- commit이 assigned scope 안에 있는가
- required slice verification이 통과했는가, 또는 gap이 명시됐는가
- subagent가 요청된 integration branch로 rebase했는가
- reported rebase target HEAD가 현재 integration branch HEAD와 같은가
- 미커밋 변경을 import하지 않는가

## Integration

- gate가 통과하면 prepared commit을 cherry-pick, merge, fast-forward, 또는 repo에 맞는 비파괴 통합 방식으로 회수합니다.
- 기준 HEAD가 subagent rebase 이후 변했으면 import하지 않고 현재 HEAD 기준 rebase를 다시 요청하거나 integration 순서를 재설계합니다.
- 메인 에이전트는 subagent output을 검증 없이 최종 결과로 승격하지 않습니다.
- 통합 후 메인 에이전트는 imported slice diff를 확인합니다.
- 통합 후 slice-level narrow checks를 실행합니다.
- contract, shared code, generated surface가 바뀌면 broader checks를 실행합니다.

## Cleanup

- worktree cleanup은 prepared commit 회수와 필요한 통합 검증 이후에만 수행합니다.
- useful evidence, failure evidence, conflict state, prepared commit state가 처리되지 않았다면 cleanup하지 않습니다.
- cleanup은 push, PR, publish, release, version bump 권한을 의미하지 않습니다.
