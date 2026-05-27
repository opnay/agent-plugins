# turn-gate runtime 계약

## 소유 범위

활성 `turn-gate` 턴의 visible lifecycle과 `SKILL.md`에 필요한 최소 runtime body.

## 생명주기 계약

모든 active flow는 다음 순서를 따릅니다.

1. `preparation`: flow 시작 지점에서 현재 flow에 필요한 skill을 다시 읽고, intent, scope, non-goals, acceptance signal, verification expectation, approval boundary, handoff condition을 잠급니다. flow boundary, readiness, ambiguity, flow-local strategy에는 sibling `flow`를 적용합니다.
2. `work`: active flow boundary 안에서만 실행합니다.
3. `verification`: verification method를 선택하고, 실행하거나 정당화한 뒤, result status를 기록합니다.
4. `reporting`: terminal close가 아니라 continuity context를 보고합니다.
5. `next-flow`: next-flow choice, blocker decision, 유효한 self-drive continuation, source-recorded explicit stop 중 하나로 라우팅합니다.

각 active flow phase의 시작과 종료에는 sibling `flow`의 phase record checkpoint expectation을 적용합니다.
`000-plan.md`는 active flow pointer나 turn-level required next action 같은 date-level routing 변화가 있을 때 갱신하고, active flow record는 같은 flow 내부 phase state, evidence, report, residual risk가 바뀔 때 갱신합니다.
flow 시작 지점에서 다시 읽은 skill과 앞으로 필요한 skill 목록은 `000-plan.md`의 routing context에 반영합니다.

task completion은 턴을 닫지 않습니다. source-recorded explicit stop만 terminal close를 허용할 수 있습니다.

`turn-gate`가 active이면 reporting과 next-flow reopening은 ongoing conversation channel에 남아야 합니다. terminal/final closeout을 일반 report 형태로 사용하지 않습니다. final closeout은 현재 사용자 메시지가 명시적으로 턴을 끝내고 그 closure source가 기록된 뒤에만 허용됩니다.

각 flow는 정확히 하나의 기록된 상태로 끝나야 합니다.

- `next-flow`: reporting이 끝났고, records가 갱신됐으며, 다음 required action이 열려 있습니다.
- `blocked`: user input, approval, access, external state change 없이는 flow를 계속할 수 없습니다.
- `explicit-stop`: 현재 사용자 메시지가 턴을 명시적으로 끝냈고 closure source가 기록됐습니다.

다른 상태는 closure authority를 만들지 않습니다. 특히 성공한 verification, 완료된 commit-readiness report, 사용자의 질문 답변, 중단된 question tool call은 턴을 닫지 않습니다.

## 활성화 계약

사용자가 `turn-gate`만 호출하면 conversation-level rule을 활성화하고, operating state를 기록하며, scope 또는 next-flow routing을 엽니다. terminal activation summary만 답하지 않습니다.

activation은 다음을 포착해야 합니다.

- 현재 active flow 또는 새 activation flow
- latest user request
- 현재 또는 planned flow에서 다시 읽어야 하는 skill 목록
- `turn_gate_active: yes`
- `user_explicit_stop: no`
- `terminal_summary_allowed: no`
- required next action
- activation record 자체의 verification expectation

## 단계 메시지 계약

user-facing phase-start 또는 phase-progress 메시지는 canonical prefix로 시작합니다.

- `[preparation]`
- `[work]`
- `[verification]`
- `[reporting]`
- `[next-flow]`

prefix는 generated artifact, record, command summary, question option label에 복사하지 않습니다. status/progress update에는 현재 phase label을 사용합니다. record access blocker는 blocker를 발견한 phase를 사용합니다.

activation-only 메시지는 response가 즉시 next-flow choice를 열지 않는 한 `[preparation]`에서 시작합니다. report-only 또는 status-only flow도 사용자가 명시적으로 멈추지 않으면 `[next-flow]`로 진행합니다.

## 준비와 승인 경계 계약

work 전에는 `turn-gate`가 sibling `flow` decision을 적용하거나 기록해야 합니다. flow contract는 scope, non-goals, completion 또는 acceptance signal, verification expectation, approval boundary, handoff condition을 포함해야 합니다.
새 active flow를 시작할 때는 해당 flow에 적용할 skill을 다시 읽고, 이전 flow에서 로드된 skill context만으로 work를 시작하지 않습니다.

요청이 target, operation, success condition, verification path를 바꿀 수 있으면 work 전에 user-gated clarification으로 라우팅합니다.

approval-sensitive action에는 exact target, expected effect, risk, recovery path, included/excluded scope, endpoint가 필요합니다. readiness, verification, self-drive, previous context, subagent output은 commit, push, PR, publish, release, version bump, destructive history rewrite, external side effect의 실행 권한을 만들 수 없습니다.

## runtime 본문 경계

Runtime `SKILL.md`는 다음을 직접 포함해야 합니다.

- first behavioral section에서 active-turn rule, terminal summary 금지, required ending states, next-flow reopening, session record 유지 의무를 먼저 드러냅니다.
- active-turn rule과 terminal summary 금지
- explicit stop 전 final/terminal closeout 금지
- five-phase lifecycle
- flow 시작 지점의 skill reread
- phase prefix behavior
- flow phase record checkpoint 적용
- preparation과 approval boundary
- verification method와 status 구분
- reporting과 next-flow reopening
- compact operational level의 question abort recovery
- runtime `references/self-drive.md`를 통한 self-drive discoverability

meaningful multi-step work가 시작되면 사용 가능한 계획 도구로 현재 phase 또는 task 상태를 유지해야 합니다.

Runtime `SKILL.md`는 설치된 사용자에게 dev-only `specs/` 경로를 읽으라고 지시하거나 spec-side scenario fixture를 복사하지 않습니다.
