# manager decomposition spec

## 목적

이 문서는 `manager`의 orchestration fit 판단, no-use/no-spawn 보고, 실무형 job 분해, 병렬/순차 실행 판단을 소유합니다.

## Fit 판단

- 먼저 worktree orchestration이 필요한지 판단합니다.
- 작은 단일 수정, 순수 read-only 답변, commit-sized slice로 나눌 수 없는 작업에는 `manager`를 사용하지 않습니다.
- 부적합하면 `Orchestration fit: no`, `Spawn plan: none`, caller-local handling, verification expectation, residual risk를 보고합니다.
- 작업이 `Workflow > Jobs > Runs` 계획 문서로 표현되고 각 job의 write scope, dependency, acceptance, handoff가 정의될 수 있을 때만 orchestration을 엽니다.

## 실무형 분해 기준

- `workstream`: 기능, 버그, 문서, 검증, release-surface 같은 실제 업무 줄기로 나눕니다.
- `write scope`: 각 job이 수정할 수 있는 module, 화면, API, 문서 표면, generated artifact 범위를 정합니다.
- `dependency`: 먼저 끝나야 하는 contract, source change, setup, verification gate를 `needs`로 둡니다.
- `parallel blockers`: 같은 파일, shared contract, migration, generated output, secret surface처럼 병렬 실행을 막는 이유를 표시합니다.
- `acceptance`: 각 job이 어떤 검증, commit, rebase, report로 끝나는지 고정합니다.

## 하위 판단 예시

- 원인 가설, 요구사항, risk, blocker는 `workstream`을 나눌 때 참고합니다.
- package, module, layer, frontend/backend/docs/tests/infra는 `write scope`를 정할 때 참고합니다.
- schema, API, service, client, UI, test contract는 `dependency`와 `parallel blockers`를 정할 때 참고합니다.
- lint, typecheck, test, build, smoke는 `acceptance`를 정할 때 참고합니다.
- generated release surface, generated client, generated migration output은 보통 병렬 writer job이 아니라 source change 이후 integration/build job으로 둡니다.
- token, credential, secret handling은 실제 secret 접근, secret rotation, 외부 인증 서비스 호출, 로그/fixture secret 노출을 승인 없는 subagent job에 포함하지 않습니다.
- 순수 조사, triage, hypothesis elimination은 commit-sized fix job이 확인되기 전에는 manager worktree job으로 만들지 않습니다.

## 병렬/순차 판단

- disjoint slice만 병렬화합니다.
- shared file, shared contract, generated output, migration, schema/API contract가 겹치면 순차화합니다.
- generated-output update는 repository policy가 다르게 정하지 않는 한 source change 이후 순차 integration/build 단계로 둡니다.
- job acceptance가 아직 보이지 않는 큰 refactor는 먼저 caller-local planning으로 좁히고, 실패하면 no-use/no-spawn으로 보고합니다.
