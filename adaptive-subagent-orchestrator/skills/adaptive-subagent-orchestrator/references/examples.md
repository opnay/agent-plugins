# Examples

## Implicit Invocation: Positive

### Branch Review

Prompt:

```text
현재 브랜치를 main과 비교해서 보안 문제, 실제 버그, 테스트 누락, 성능 위험을 검토해줘.
```

Expected:

- skill: yes
- mode: PARALLEL_READ
- split by security, correctness, tests, performance
- main agent integrates only evidence-backed findings

### Multi-Module Incident

Prompt:

```text
로그인, 결제, 인벤토리 모듈을 통과하는 요청이 간헐적으로 실패해. 각 모듈의 실행 경로와 가능한 원인을 조사해줘.
```

Expected:

- skill: yes
- mode: PARALLEL_READ
- module or execution-path explorers
- main agent integrates causes before implementation

### Test Bundle Failures

Prompt:

```text
unit, integration, end-to-end 테스트 묶음이 각각 실패하고 있어. 원인을 찾고 최소 수정안을 적용해줘.
```

Expected:

- skill: yes
- initial mode: PARALLEL_READ
- investigate by test bundle
- switch to PARALLEL_WRITE only if changed files are disjoint
- use one writer if fixes overlap

### Design Options

Prompt:

```text
이 기능을 구현할 수 있는 세 가지 구조를 비교하고, 현재 코드베이스에 가장 적합한 방식을 선택해줘.
```

Expected:

- skill: yes
- mode: PARALLEL_READ
- compare by option or evaluation lens
- main agent chooses with one shared criteria set

### Migration Impact

Prompt:

```text
이 마이그레이션이 API, 배치 작업, 모바일 클라이언트, 관리자 도구에 미치는 영향을 조사해줘.
```

Expected:

- skill: yes
- mode: PARALLEL_READ
- split by impacted area
- main agent integrates shared contracts and risks

## Implicit Invocation: Negative

### README Typo

Prompt:

```text
README의 오타를 고쳐줘.
```

Expected:

- skill: no
- mode: DIRECT
- no subagents

### Local Rename

Prompt:

```text
이 함수의 변수 이름을 더 명확하게 바꿔줘.
```

Expected:

- skill: no
- mode: DIRECT

### One Type Error

Prompt:

```text
이 파일의 타입 오류 한 개를 수정해줘.
```

Expected:

- skill: no
- mode: DIRECT

### Sequential Dependency

Prompt:

```text
이 함수의 반환값을 변경하고, 그 결과를 사용해 바로 아래 함수를 수정해줘.
```

Expected:

- skill: no or DIRECT
- no subagents due to strong sequence

### User Forbids Subagents

Prompt:

```text
이 작업은 서브에이전트 없이 직접 처리해줘.
```

Expected:

- user instruction wins
- mode: DIRECT

## Boundary Cases

### Repo-Wide Lint Fix

Prompt:

```text
저장소 전체의 lint 오류를 고쳐줘.
```

Judgment:

- PARALLEL_READ if independent packages have separate causes
- limited PARALLEL_WRITE only if writable file sets are disjoint
- DIRECT if one lint config or shared type root causes most errors
- never parallelize because file count is high

### Cross-Layer Feature

Prompt:

```text
프런트엔드, API, 데이터베이스에 걸친 기능을 구현해줘.
```

Judgment:

- PARALLEL_READ until the shared interface is fixed
- main agent defines the contract
- PARALLEL_WRITE only after file ownership and dependencies are separated
- sequential implementation if the interface remains unstable

### Dependency Upgrade

Prompt:

```text
주요 의존성의 버전을 올리고 호환성 문제를 해결해줘.
```

Judgment:

- parallelize package compatibility and test impact research
- one writer handles lockfile and shared config
- main agent runs install and full tests

## Good Split

- Agent A: read-only auth execution path and logs.
- Agent B: read-only payment execution path and logs.
- Agent C: read-only inventory execution path and logs.
- Main: compare causes, decide shared contract, implement or assign one writer.

Why good: scopes differ, outputs are comparable, writes wait until evidence is integrated.

## Bad Split

- Agent A edits shared API types.
- Agent B edits frontend against guessed API types.
- Agent C edits backend against a different guessed API shape.

Why bad: shared interface is not fixed and writers depend on each other's unfinished work.

## Parallel Read Then Sequential Implementation

Use for review findings across security, correctness, tests, and performance. Let explorers report evidence. Main agent then applies one coherent patch and runs full validation.

## Safe Parallel Write

Use only when ownership is disjoint:

- Worker A: `packages/a/src/**`, `packages/a/test/**`
- Worker B: `packages/b/src/**`, `packages/b/test/**`
- Main: lockfile, root config, integration tests

## Dangerous Same-File Parallel Write

Never assign two workers to `src/api/schema.ts`, `pnpm-lock.yaml`, generated route files, or shared fixtures at the same time. Use one writer.
