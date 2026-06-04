# manager failure spec

## 목적

이 문서는 `manager`의 실패, 재시도, blocked handoff 처리 기준을 소유합니다.

## Failure Handling

- commit 전 실패: worktree를 보존하고 실패 evidence를 회수합니다. commit-required job이면 handoff 실패이고, no-commit job이면 evidence handoff 조건 충족 여부로 판단합니다.
- 미커밋 변경만 보고: import하지 않고, 같은 subagent에게 범위 내 변경을 검증, commit, rebase하도록 재요청하거나 evidence 가치에 따라 worktree 보존/폐기를 판단합니다.
- missing verification: 기본적으로 import하지 않고 추가 verification을 요청하거나 handoff를 거부합니다. verification이 불가능한 예외 import는 명시 승인과 risk 기록이 있을 때만 검토합니다.
- unresolved rebase conflict: 해당 slice를 blocked로 두고 순서 변경, scope 축소, 수동 통합 후보를 판단합니다.
- out-of-scope commit: prepared commit이라도 회수하지 않고 correction을 요청합니다.
- repeated failure: 병렬화를 중단하고 더 작은 sequential slice로 재분해합니다.

## Conflict Evidence

unresolved conflict에서는 subagent에게 다음을 보고하게 합니다.

- conflict files
- rebase state
- attempted resolution
- verification gap
- next suggested action

## Blocked State

- blocked worktree는 cleanup하지 않습니다.
- blocked state는 prepared commit import 권한이 아닙니다.
- 수동 통합, destructive cleanup, push/PR/release는 별도 승인 경계를 유지합니다.
