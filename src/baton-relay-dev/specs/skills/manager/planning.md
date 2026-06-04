# manager planning spec

## 목적

이 문서는 `manager`가 작업을 dispatch하기 전에 작성해야 하는 Markdown workflow plan todo 문서 계약을 소유합니다.

## Plan Shape

계획 문서는 데이터 구조가 아니라 사람이 읽고 갱신하는 Markdown todo 문서입니다.
계층은 GitHub Actions의 `workflow > job > run` 방식을 예시로 삼되, 실제 표면은 체크리스트입니다.
frontmatter는 반복되는 workflow/job metadata를 줄이는 용도로 사용합니다.
계획 문서는 기본적으로 정적이며, 실행 중 변경 가능한 값은 mutable allowlist에 있는 본문 값뿐입니다.

필수 계층:

- `Workflow`: 전체 사용자 요청 하나
- `Jobs`: subagent에게 분배할 작업 단위
- `Runs`: 각 job 안에서 실제 수행할 todo 단계

## Workflow Section

Workflow section은 다음을 포함합니다.

- frontmatter: objective, integration branch, dispatch fit, dispatch reason, global risks
- body: workflow-level verification, cleanup, residual risk

## Job Section

각 job은 다음을 포함합니다.

- frontmatter: job id/title, static needs, worktree, workstream, write scope, parallel blockers, handoff 조건
- body: job status, runs, acceptance, handoff evidence todo

## Run Todo

- run은 실제 실행 가능한 todo여야 합니다.
- subagent run은 worktree 안에서 수행 가능한 단계여야 합니다.
- run은 조사, 구현, 검증, commit, merge-prep 대기, rebase, handoff 보고처럼 실행 순서를 드러냅니다.
- 실행 중 상태 업데이트가 가능하도록 체크박스로 씁니다.

## Mutable Allowlist

실행 중 바꿀 수 있는 값은 아래로 제한합니다.

- job `Status` 값
- `Runs`, `Acceptance`, `Handoff`, `Workflow Verification`, `Cleanup` 체크박스 상태
- 완료한 체크박스 아래에 추가하는 evidence text
- `Residual Risk`

아래 값은 정적입니다.

- 전체 frontmatter
- job id/title
- `needs`
- worktree
- workstream
- write scope
- parallel blockers
- handoff 조건

## Acceptance And Handoff

subagent job은 acceptance와 handoff를 반드시 가집니다.

- acceptance는 job 완료 판정입니다.
- handoff는 commit, rebase, verification, changed files, residual risk 보고 조건입니다.
- 미커밋 변경 import 금지는 handoff 조건에 포함합니다.

## Planning Rules

- `manager`는 plan 없이 subagent를 dispatch하지 않습니다.
- `needs`는 정적 dependency graph입니다. 실행 중 갱신하지 않습니다.
- job 시작 가능 여부는 `needs` 대상 job의 body checklist completion과 handoff evidence로 판단합니다.
- `needs: []`인 job은 동시에 시작 가능한 후보입니다.
- 같은 파일, shared contract, generated output, secret surface를 건드리는 job은 같은 병렬 그룹에 넣지 않습니다.
- generated output job은 repository policy가 다르게 정하지 않는 한 source job 이후 순차 job으로 둡니다.
- 계획 문서는 실행 중 mutable allowlist만 갱신하는 운영 표면입니다.

## Document Template

runtime template 위치: `skills/manager/templates/workflow-plan.md`

`manager`가 workflow plan을 작성할 때는 runtime template file을 기본값으로 사용합니다.
필요 없는 job section은 삭제하지 말고 `none` 또는 `not-required`로 표시해 판단 흔적을 남깁니다.
frontmatter 값은 계획 기준 metadata이며 실행 중 갱신하지 않습니다.
진행 상태는 mutable allowlist에 있는 본문 값으로만 갱신합니다.

아래 내용은 runtime template과 같은 계약을 설명하는 spec copy입니다.

```md
---
workflow: <workflow-name>
objective: <사용자 요청을 완료 상태 기준으로 한 문장으로 적는다>
integration_branch: <branch-name-or-unknown>
dispatch:
  fit: <yes | no>
  reason: <왜 이 실행 모드인지>
  global_risk: <shared-contract | generated-output | secret | migration | unknown-scope | none>
jobs:
  - id: job-1
    title: <job-title>
    needs: []
    worktree: <none | ../worktrees/<name>>
    workstream: <feature | bug | docs | verification | release-surface | other>
    write_scope: <module/screen/API/doc/generated-artifact this job may change>
    parallel_blockers: <none | shared-file | shared-contract | generated-output | migration | secret-surface>
    handoff:
      requires_commit: <yes | no>
      requires_rebase: <yes | no>
      import_uncommitted_changes: no
      report: <commit hash, rebase target HEAD, verification, changed files, residual risk>
  - id: job-2
    title: <job-title>
    needs: [job-1]
    worktree: <none | ../worktrees/<name>>
    workstream: <feature | bug | docs | verification | release-surface | other>
    write_scope: <module/screen/API/doc/generated-artifact this job may change>
    parallel_blockers: <none | shared-file | shared-contract | generated-output | migration | secret-surface>
    handoff:
      requires_commit: <yes | no>
      requires_rebase: <yes | no>
      import_uncommitted_changes: no
      report: <required handoff evidence>
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
```
