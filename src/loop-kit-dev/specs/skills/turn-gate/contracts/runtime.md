# turn-gate runtime 계약

## 소유 범위

활성 `turn-gate` 턴의 visible lifecycle과 `SKILL.md`에 필요한 최소 runtime body.

## 생명주기 계약

모든 active flow는 다음 순서를 따릅니다.

1. `intake`: flow 시작 지점에서 현재 flow에 필요한 skill을 다시 읽고, raw input과 해석을 분리하며 goal, non-goals, authority-sensitive signal, discovery topic을 드러냅니다.
2. `framing`: 필수 `flow` 계약을 적용해 항목을 분류하고, 필요하면 finite sub-flow candidate를 설계하며, selected active flow와 candidate를 구분합니다.
3. `preparation`: selected active flow의 intent, scope, non-goals, acceptance signal, verification expectation, approval boundary, handoff condition을 잠급니다. readiness, ambiguity, flow-local strategy에는 의존하는 `flow`를 적용합니다.
4. `work`: active flow boundary 안에서만 실행합니다.
5. `verification`: verification method를 선택하고, 실행하거나 정당화한 뒤, result status를 기록합니다.
6. `reporting`: terminal close가 아니라 continuity context를 보고합니다.
7. `next-flow`: next-flow choice, blocker decision, 유효한 self-drive continuation, source-recorded explicit stop 중 하나로 라우팅합니다.

각 active flow phase의 시작과 종료에는 필수 `flow`의 phase record checkpoint expectation을 적용합니다.
`000-plan.md`는 active flow pointer, turn-level required next action, active skill list, self-drive status, unapproved action state 같은 date-level routing 변화가 있을 때 갱신하고, active flow record는 같은 flow 내부 phase state, evidence, report, residual risk가 바뀔 때 갱신합니다.
flow 시작 지점에서 다시 읽을 skill 목록은 `000-plan.md` frontmatter의 `active_skills`에 이름만 반영합니다.

task completion은 턴을 닫지 않습니다. source-recorded explicit stop만 terminal close를 허용할 수 있습니다.

`turn-gate`가 active이면 reporting과 next-flow reopening은 ongoing conversation channel에 남아야 합니다. terminal/final closeout을 일반 report 형태로 사용하지 않습니다. final closeout은 현재 사용자 메시지가 명시적으로 턴을 끝내고 그 closure source가 기록된 뒤에만 허용됩니다.

## interruption 계약

`interruption`은 active flow가 이미 진행 중일 때 새 사용자 메시지가 도착한 경우에만 열리는 entry-only phase입니다. 일반 lifecycle phase가 아니며 `intake`, `framing`, `preparation`, `work`, `verification`, `reporting`, `next-flow`를 대체하지 않습니다.

`interruption`이 시작되면 현재 foreground flow의 phase, scope, non-goals, approval boundary, verification status, required next action을 보존한 뒤 새 메시지가 active flow에 미치는 영향을 분류합니다.

분류 결과는 하나만 선택합니다.

- `inline-answer`: active flow 계약을 바꾸지 않는 질문에 답하고 보존된 phase로 돌아갑니다.
- `current-flow-revision`: scope, non-goals, completion criteria, verification expectation, approval boundary, handoff condition이 바뀌면 active flow 계약을 갱신하고 `framing` 또는 `preparation`으로 돌아갑니다.
- `background-current-flow`: 현재 flow를 나중에 재개해야 하지만 새 foreground flow가 먼저 필요하면 현재 flow를 background로 기록하고 새 flow를 시작합니다.
- `reserve-later-analysis`: 지금 처리하지 않을 관련 주제는 future candidate로 기록하고 보존된 phase로 돌아갑니다.
- `supersede-current-flow`: 사용자가 현재 flow를 취소하거나 대체하면 현재 flow를 superseded로 기록하고 새 flow를 시작합니다.
- `blocker-question`: 결정, 승인, 접근, scope gap 없이는 계속하면 위험할 때 active flow를 blocked로 두고 질문합니다.
- `explicit-stop`: 사용자가 턴 종료를 명시한 경우에만 closure source를 기록하고 종료를 허용합니다.

`interruption` 결과는 commit, push, PR, publish, release, version bump, destructive action, 또는 active flow 계약에 없던 구현 시작 권한을 만들 수 없습니다.

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
- turn-gate active state
- no explicit-stop source unless user actually stopped the turn
- terminal summary blocked state
- required next action
- activation record 자체의 verification expectation

## 단계 메시지 계약

user-facing phase-start 또는 phase-progress 메시지는 canonical prefix로 시작합니다.

- `[intake]`
- `[framing]`
- `[preparation]`
- `[work]`
- `[verification]`
- `[reporting]`
- `[next-flow]`

prefix는 generated artifact, record, command summary, question option label에 복사하지 않습니다. status/progress update에는 현재 phase label을 사용합니다. record access blocker는 blocker를 발견한 phase를 사용합니다.

activation-only 메시지는 response가 즉시 next-flow choice를 열지 않는 한 `[intake]`에서 시작합니다. report-only 또는 status-only flow도 사용자가 명시적으로 멈추지 않으면 `[next-flow]`로 진행합니다.

## 준비와 승인 경계 계약

work 전에는 `turn-gate`가 필수 `flow` decision을 적용하거나 기록해야 합니다. flow contract는 scope, non-goals, completion 또는 acceptance signal, verification expectation, approval boundary, handoff condition을 포함해야 합니다.
새 active flow를 시작할 때는 해당 flow에 적용할 skill을 다시 읽고, 이전 flow에서 로드된 skill context만으로 work를 시작하지 않습니다. 사용자 메시지를 받아 framing을 거쳐 preparation으로 들어갈 때는 preparation 전에 `turn-gate`와 `flow`를 다시 읽고 `000-plan.md` frontmatter의 `active_skills`에 둘 다 유지합니다.

요청이 target, operation, success condition, verification path를 바꿀 수 있으면 work 전에 user-gated clarification으로 라우팅합니다.

approval-sensitive action에는 exact target, expected effect, risk, recovery path, included/excluded scope, endpoint가 필요합니다. readiness, verification, generated release surface build/readback, self-drive, previous context, subagent output은 commit, push, PR, publish, release, version bump, destructive history rewrite, external side effect의 실행 권한을 만들 수 없습니다.

## runtime 본문 경계

Runtime `SKILL.md`는 다음을 직접 포함해야 합니다.

- first behavioral section에서 active-turn rule, terminal summary 금지, required ending states, next-flow reopening, session record 유지 의무를 먼저 드러냅니다.
- active-turn rule과 terminal summary 금지
- explicit stop 전 final/terminal closeout 금지
- active flow lifecycle
- active flow 도중 새 사용자 메시지를 처리하는 `interruption` entry-only phase
- flow 시작 지점의 skill reread
- 사용자 메시지에서 preparation으로 넘어갈 때 `turn-gate`와 `flow` reread
- phase prefix behavior
- flow phase record checkpoint 적용
- preparation과 approval boundary
- verification method와 status 구분
- reporting과 next-flow reopening
- compact operational level의 question abort recovery
- runtime `references/self-drive.md`를 통한 self-drive discoverability

meaningful multi-step work가 시작되면 사용 가능한 계획 도구로 현재 phase 또는 task 상태를 유지해야 합니다.

Runtime `SKILL.md`는 설치된 사용자에게 dev-only `specs/` 경로를 읽으라고 지시하거나 spec-side scenario fixture를 복사하지 않습니다.
