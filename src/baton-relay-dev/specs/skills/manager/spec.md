# manager 스킬 스펙

## 목적

`manager`는 모든 작업 요청을 worktree-safe task slice로 구조분해하고, 각 slice를 fresh subagent에게 맡기며, commit-required job은 commit/rebase 완료 후 prepared commit으로, no-commit job은 evidence handoff로 회수하도록 관리하는 스킬입니다.

## 경계

- 포함:
  - Markdown workflow plan 작성
  - 실무형 작업 분해 기준 선택
  - task slice와 worktree write scope 정의
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

## 대표 표면

- runtime skill: `baton-relay-dev/skills/manager/SKILL.md`
- workflow plan template: `baton-relay-dev/skills/manager/templates/workflow-plan.md`
- 사용자 의도와 전체 플로우 그래프: `baton-relay-dev/specs/skills/manager/intent.md`

## Sub-Spec Map

- `planning.md`: Markdown workflow plan의 `Workflow > Jobs > Runs` todo 문서 계약과 runtime template 위치를 소유합니다.
- `decomposition.md`: plan-first decomposition, 최소 단일 job 구성, workstream/write scope/dependency/parallel blockers/acceptance 기준, 병렬/순차 판단을 소유합니다.
- `lifecycle.md`: dispatch packet과 subagent 작업 사이클을 소유합니다.
- `handoff.md`: commit/rebase 또는 evidence handoff gate, prepared commit 회수, integration verification, cleanup을 소유합니다.
- `failure.md`: no commit, uncommitted-only, missing verification, unresolved conflict, out-of-scope commit, repeated failure 처리를 소유합니다.

## 처리하려는 작업 형태

- 단일 slice 또는 여러 독립 worktree에서 진행할 수 있는 작업
- 문제 가설, 코드 영역, workflow 단계, 계약 표면, 검증 경로별로 분리 가능한 작업
- 각 slice가 commit 또는 evidence handoff로 회수될 수 있는 작업
- subagent가 rebase conflict를 자기 worktree에서 해결할 수 있는 작업
- 메인 에이전트가 통합 순서와 검증을 책임져야 하는 작업

## 핵심 처리 계약

- 메인 에이전트는 manager로 동작합니다.
- `manager`는 작업 dispatch 전에 `Workflow > Jobs > Runs` 구조의 Markdown todo 계획 문서를 작성합니다.
- 모든 사용자 작업은 최소 하나의 job으로 계획합니다.
- 실제 task slice 작업은 subagent가 수행합니다.
- 하나의 subagent는 하나의 task slice와 하나의 worktree만 담당합니다.
- 다음 task에는 새 subagent를 만듭니다.
- manager workflow는 subagent local commit 권한을 자동으로 만들지 않습니다.
- subagent 완료는 작업 수정, 검증, git commit, 메인 기준 branch로 rebase 완료를 모두 포함합니다.
- 메인 에이전트는 미커밋 변경을 가져오지 않습니다.
- 메인 에이전트는 rebase 완료된 prepared commit만 회수합니다.
- 기준 HEAD가 subagent rebase 이후 변했으면 통합 전에 rebase 재요청 또는 통합 순서 재설계를 합니다.
- worktree cleanup은 commit 회수와 통합 검증 후에만 수행합니다.

## 검토 질문

- 이 작업은 `Workflow > Jobs > Runs` 계획 문서로 표현할 수 있는가?
- 각 job은 workstream, write scope, dependency, parallel blockers, acceptance 기준으로 설명되는가?
- subagent가 자기 worktree 안에서만 완료할 수 있는가?
- 각 slice의 write scope가 충돌 없이 분리되는가?
- rebase target branch와 handoff HEAD를 확인할 수 있는가?
- prepared commit을 가져온 뒤 어떤 검증이 필요한가?
- worktree cleanup 전에 회수해야 할 evidence가 남아 있는가?
- 이 요청을 최소 하나의 subagent job으로 만들려면 어떤 write scope, verification, handoff 조건이 필요한가?
- 바로 dispatch할 수 없다면 어떤 승인, 입력, 또는 권한이 blocked 상태를 해소하는가?
- subagent local commit 권한이 dispatch 전에 확인됐는가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 예.
- 이유: `manager`는 worktree orchestration의 핵심 절차를 독립적으로 제공해야 합니다. 다른 subagent role skill이나 commit skill이 있더라도 lifecycle, handoff, cleanup 계약은 숨은 sibling context에 의존하지 않습니다.

## 확장 원칙

- 새 decomposition axis는 실제 orchestration 판단을 바꿀 때만 추가합니다.
- integration 방식은 repo policy에 맞춰 확장할 수 있지만 prepared commit만 회수한다는 원칙은 유지합니다.
- approval-sensitive action은 별도 스킬이나 사용자 승인으로 넘깁니다.
