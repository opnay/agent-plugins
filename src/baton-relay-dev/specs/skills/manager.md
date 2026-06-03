## 사용자 스펙 의도

- 새로운 루프 스킬을 만들려고 한다. 이 스킬을 사용하면 해당 에이전트는 오케스트레이션으로서 동작한다. subagent가 모든 작업을 진행해야 한다. 어떤 작업을 받으면 그 작업을 구조분해해서 git worktree로 구분해 subagent를 구동한다. 각각의 worktree에서 작업이 완료되면 그 작업분을 커밋해서 메인 작업위치에 커밋을 가져오고, 작업하던 worktree는 정리해야 한다. 구조분해 방식은 문제를 해결하는 방식, 코드를 구분하는 방식, workflow를 구분하는 방식 등 다양한 기준점이 필요하다.
- subagent 사이클도 지정해야 한다. 작업 시작부터 작업 완료까지 subagent가 동작하고, 작업 완료가 되면 subagent는 종료한다. 다음 작업을 위해 새로 생성한다. subagent의 작업 완료는 git commit 이후, 메인 에이전트가 병합 준비를 요청하면 subagent가 메인 에이전트가 동작하는 브랜치로 rebase하고, 완료되면 메인 에이전트가 해당 커밋을 가져온다. rebase는 완료된 상태이기 때문에 conflict는 발생할 가능성이 없다.
- 플러그인 이름은 `baton-relay`로 정한다. 스킬 이름은 `manager`로 정한다.

---

# manager 스킬 스펙

## 목적

`manager`는 큰 작업을 worktree-safe task slice로 구조분해하고, 각 slice를 fresh subagent에게 맡기며, subagent가 commit과 rebase를 완료한 뒤 prepared commit만 메인 작업 위치로 회수하도록 관리하는 스킬입니다.

## 경계

- 포함:
  - 작업 구조분해 기준 선택
  - task slice와 worktree ownership 정의
  - subagent dispatch packet 작성
  - subagent lifecycle 관리
  - commit/rebase handoff gate
  - 메인 branch HEAD 일치 확인
  - prepared commit 통합 방식 선택
  - worktree cleanup 기준
  - 실패, 재시도, rebase 재요청 판단
- 제외:
  - subagent runtime 구현
  - subagent output 무검증 수용
  - approval 없는 commit, push, PR, publish, release, version bump, destructive work
  - 모든 작업의 자동 병렬화
  - 특정 언어/프레임워크 검증 전략 소유

## 처리하려는 작업 형태

- 여러 독립 worktree에서 진행할 수 있는 구현 작업
- 문제 가설, 코드 영역, workflow 단계, 계약 표면, 검증 경로별로 분리 가능한 작업
- 각 slice가 하나의 commit으로 회수될 수 있는 작업
- subagent가 rebase conflict를 자기 worktree에서 해결할 수 있는 작업
- 메인 에이전트가 통합 순서와 검증을 책임져야 하는 작업

## 엔트리포인트 / 대표 표면

- 대표 표면: `baton-relay-dev/skills/manager/SKILL.md`
- 호출 방식: 사용자가 `$baton-relay:manager` 또는 `$baton-relay-dev:manager`를 직접 호출하거나, manifest prompt가 worktree orchestration이 적합하다고 안내할 때 사용한다.

## 핵심 처리 계약

- 메인 에이전트는 manager로 동작한다.
- 실제 task slice 작업은 subagent가 수행한다.
- 하나의 subagent는 하나의 task slice와 하나의 worktree만 담당한다.
- 다음 task에는 새 subagent를 만든다.
- manager workflow는 subagent local commit 권한을 자동으로 만들지 않는다. dispatch 전에 local commit이 현재 승인 경계 안에 있는지 확인한다.
- subagent 완료는 작업 수정, 검증, git commit, 메인 기준 branch로 rebase 완료를 모두 포함한다.
- 메인 에이전트는 미커밋 변경을 가져오지 않는다.
- 메인 에이전트는 rebase 완료된 prepared commit만 회수한다.
- 기준 HEAD가 subagent rebase 이후 변했으면 통합 전에 rebase 재요청 또는 통합 순서 재설계를 한다.
- worktree cleanup은 commit 회수와 통합 검증 후에만 수행한다.

## 구조분해 기준

- 문제 기준: 원인 가설, 요구사항, risk, blocker를 기준으로 나눈다.
- 코드 기준: package, module, layer, frontend/backend/docs/tests/infra처럼 ownership을 기준으로 나눈다.
- workflow 기준: discovery, implementation, verification, documentation, refactor처럼 단계로 나눈다.
- 계약 기준: schema, API, service, client, UI, test contract처럼 interface를 기준으로 나눈다.
- 검증 기준: 독립적으로 lint, typecheck, test, build, smoke를 수행할 수 있는 단위로 나눈다.
- 충돌 기준: 같은 파일이나 같은 contract를 동시에 바꾸는 slice는 병렬로 두지 않고 순차화한다.
- generated output 기준: generated release surface, generated client, generated migration output은 보통 병렬 writer slice가 아니라 source 변경 이후 integration/build 단계 산출물로 다룬다.
- 보안 기준: token, credential, secret handling은 실제 secret 접근, secret rotation, 외부 인증 서비스 호출, 로그/fixture secret 노출을 승인 없는 subagent 작업에 포함하지 않는다.
- read-only 기준: 순수 조사, triage, hypothesis elimination은 commit-sized fix slice가 확인되기 전에는 manager worktree slice로 만들지 않는다.

## Subagent Lifecycle

- subagent는 task slice 시작 시 새로 생성한다.
- subagent packet에는 objective, worktree path, branch, write scope, non-goals, verification, commit approval state, commit expectation, rebase target, stop condition을 포함한다.
- subagent는 자기 worktree 밖을 수정하지 않는다.
- subagent는 작업과 검증 후 commit을 만든다.
- 메인 에이전트가 병합 준비를 요청하면 subagent는 메인 에이전트의 현재 integration branch로 rebase한다.
- rebase conflict는 subagent worktree 안에서 해결한다.
- rebase 성공 후 subagent는 commit hash, rebase target HEAD, verification result, changed files, residual risk를 보고하고 종료한다.
- 종료된 subagent는 다음 task에 재사용하지 않는다.

## Handoff And Integration

- 메인 에이전트는 subagent가 보고한 rebase target HEAD와 현재 integration branch HEAD가 같은지 확인한다.
- 같으면 prepared commit을 cherry-pick, merge, fast-forward, 또는 repo에 맞는 비파괴 통합 방식으로 회수한다.
- 다르면 subagent에게 현재 HEAD 기준 rebase를 다시 요청하거나 통합 순서를 재설계한다.
- 통합 후 메인 에이전트는 해당 slice의 검증 결과와 전체 회귀 검증 필요성을 따로 판단한다.
- worktree cleanup은 통합과 필요한 검증이 끝난 뒤 수행한다.

## 실패 처리

- subagent가 commit 전 실패하면 worktree를 보존하고 실패 evidence를 회수한다.
- subagent가 미커밋 변경만 보고하면 import하지 않고, 같은 subagent에게 범위 내 변경을 검증, commit, rebase하도록 재요청하거나 evidence 가치에 따라 worktree 보존/폐기를 판단한다.
- subagent가 rebase conflict를 해결하지 못하면 해당 slice를 blocked로 두고 메인 에이전트가 순서 변경, scope 축소, 수동 통합 후보를 판단한다.
- unresolved conflict에서는 subagent에게 conflict files, rebase state, attempted resolution, verification gap을 보고하게 하고, worktree는 cleanup하지 않는다.
- subagent 결과가 scope를 벗어나면 prepared commit이라도 회수하지 않는다.
- verification이 누락된 commit은 기본적으로 import하지 않고 추가 verification을 요청하거나 handoff를 거부한다. verification이 불가능한 예외 import는 명시 승인과 risk 기록이 있을 때만 검토한다.
- 같은 실패가 반복되면 병렬화를 중단하고 더 작은 sequential slice로 재분해한다.

## 검토 질문

- 이 작업은 하나 이상의 독립 commit slice로 나뉘는가?
- subagent가 자기 worktree 안에서만 완료할 수 있는가?
- 각 slice의 write scope가 충돌 없이 분리되는가?
- rebase target branch와 handoff HEAD를 확인할 수 있는가?
- prepared commit을 가져온 뒤 어떤 검증이 필요한가?
- worktree cleanup 전에 회수해야 할 evidence가 남아 있는가?
- 이 요청은 no-use/no-spawn으로 보고해야 하는 read-only, tiny edit, 또는 commit-sized slice 미확정 상태인가?
- subagent local commit 권한이 dispatch 전에 확인됐는가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 예.
- 그렇다면 왜 필요한가 / 아니라면 어떤 sibling context를 허용하는가: `manager`는 worktree orchestration의 핵심 절차를 독립적으로 제공해야 한다. 다른 subagent role skill이나 commit skill이 있더라도 이 skill의 lifecycle, handoff, cleanup 계약은 숨은 sibling context에 의존하지 않는다.

## 확장 원칙

- 새 decomposition axis는 실제 orchestration 판단을 바꿀 때만 추가한다.
- integration 방식은 repo policy에 맞춰 확장할 수 있지만 prepared commit만 회수한다는 원칙은 유지한다.
- approval-sensitive action은 별도 스킬이나 사용자 승인으로 넘긴다.
