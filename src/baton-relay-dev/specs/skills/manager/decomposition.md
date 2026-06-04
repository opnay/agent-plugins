# manager decomposition spec

## 목적

이 문서는 `manager`의 plan-first decomposition, 최소 단일 job 구성, 실무형 job 분해, 병렬/순차 실행 판단을 소유합니다.

## Plan-First 판단

- `manager`가 선택된 모든 요청은 먼저 `Workflow > Jobs > Runs` 계획 문서로 표현합니다.
- 작은 단일 수정, 순수 read-only 답변, 조사, 검증, 정리 요청도 최소 하나의 subagent job으로 만듭니다.
- 하나의 작업이면 `job-1`만 둡니다.
- 여러 작업이면 write scope, dependency, acceptance, handoff가 분리되는 만큼 job을 나눕니다.
- 즉시 dispatch할 수 없으면 relay를 종료하지 않고 `blocked` job 또는 `needs-approval` job으로 계획합니다.
- blocked job은 필요한 사용자 입력, 승인, 권한, secret boundary, destructive boundary, commit authority를 acceptance 또는 handoff 조건에 명시합니다.

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
- 순수 조사, triage, hypothesis elimination은 read-only subagent job으로 만들되, 결과가 수정 작업으로 전환되면 새 write job을 추가하기 전에 계획과 승인 경계를 다시 고정합니다.

## 병렬/순차 판단

- disjoint slice만 병렬화합니다.
- shared file, shared contract, generated output, migration, schema/API contract가 겹치면 순차화합니다.
- generated-output update는 repository policy가 다르게 정하지 않는 한 source change 이후 순차 integration/build 단계로 둡니다.
- job acceptance가 아직 보이지 않는 큰 refactor는 blocked planning job으로 두고, 필요한 질문과 완료 기준을 acceptance에 적습니다.
