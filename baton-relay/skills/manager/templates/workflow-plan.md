---
workflow: <workflow-name>
objective: <사용자 요청을 완료 상태 기준으로 한 문장으로 적는다>
integration_branch: <branch-name-or-unknown>
dispatch:
  mode: <single-job | multi-job | blocked>
  reason: <왜 이 작업 분해와 실행 모드인지>
  global_risk: <shared-contract | generated-output | secret | migration | unknown-scope | none>
jobs:
  - id: job-1
    title: <job-title>
    needs: []
    worktree: <../worktrees/<name> | pending-approval | pending-input>
    workstream: <feature | bug | docs | verification | release-surface | research | planning | other>
    write_scope: <module/screen/API/doc/generated-artifact this job may change>
    parallel_blockers: <none | shared-file | shared-contract | generated-output | migration | secret-surface>
    handoff:
      requires_commit: <yes | no>
      requires_rebase: <yes | no>
      import_uncommitted_changes: no
      report: <commit hash and rebase target HEAD when commit-required, or no-commit evidence; verification, changed files state, residual risk>
  - id: job-2
    title: <job-title>
    needs: [job-1]
    worktree: <../worktrees/<name> | pending-approval | pending-input>
    workstream: <feature | bug | docs | verification | release-surface | research | planning | other>
    write_scope: <module/screen/API/doc/generated-artifact this job may change>
    parallel_blockers: <none | shared-file | shared-contract | generated-output | migration | secret-surface>
    handoff:
      requires_commit: <yes | no>
      requires_rebase: <yes | no>
      import_uncommitted_changes: no
      report: <commit handoff or no-commit evidence required for this job>
---

# Workflow

## Mutable Fields

- Job `Status` values
- Checklist states under `Runs`, `Acceptance`, `Handoff`, `Workflow Verification`, and `Cleanup`
- Evidence text appended under completed checklist items
- `Residual Risk`

## Jobs

### Job 1. <job-title>

Status: planned

#### Runs

- [ ] <실행 가능한 todo 1>
- [ ] <실행 가능한 todo 2>
- [ ] <필요하면 verification todo>
- [ ] <subagent job이면 commit todo>
- [ ] <subagent job이면 merge-prep 대기 todo>
- [ ] <subagent job이면 rebase todo>
- [ ] <subagent job이면 handoff 보고 todo>

#### Acceptance

- [ ] <job 완료 판정 1>
- [ ] <verification이 통과했거나 gap이 명시됐다>
- [ ] <assigned scope 안에서만 변경했다>

#### Handoff

- [ ] frontmatter handoff 조건을 만족했다
- [ ] handoff report를 기록했다

---

### Job 2. <job-title>

Status: planned

#### Runs

- [ ] <실행 가능한 todo 1>
- [ ] <실행 가능한 todo 2>

#### Acceptance

- [ ] <job 완료 판정>

#### Handoff

- [ ] frontmatter handoff 조건을 만족했다
- [ ] handoff report를 기록했다

## Workflow Verification

- [ ] <통합 후 전체 검증 1>
- [ ] <통합 후 전체 검증 2>

## Cleanup

- [ ] imported commit과 evidence 확인 후 완료된 worktree 정리
- [ ] blocked worktree는 cleanup하지 않고 상태와 이유 기록

## Residual Risk

- `<none | 남은 위험과 이유>`
