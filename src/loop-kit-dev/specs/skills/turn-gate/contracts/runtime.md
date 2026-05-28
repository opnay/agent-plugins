# turn-gate runtime 계약

## 소유 범위

활성 `turn-gate` 턴의 visible lifecycle과 `SKILL.md`에 필요한 최소 runtime body.

## 생명주기 계약

모든 active flow는 `intake -> framing -> preparation -> work -> verification -> reporting -> next-flow` 순서를 따릅니다.

`turn-gate`는 각 phase에서 `flow` decision을 적용합니다.
`flow`는 taxonomy, readiness, discovery, ambiguity, contract impact, phase checkpoint expectation, verification expectation, flow-local strategy, handoff를 소유합니다.
`turn-gate`는 그 결과를 active turn 안에서 기록하고 라우팅합니다.

phase별 책임은 다음처럼 압축합니다.

- `intake`: 필요한 skill을 다시 읽고 `flow` intake 결과를 기록합니다.
- `framing`: `flow` classification으로 selected active flow와 candidate를 구분합니다.
- `preparation`: `flow` readiness/ambiguity 결과로 work 전 계약을 잠급니다.
- `work`: active flow boundary 안에서만 실행합니다.
- `verification`: method와 result status를 분리해 기록합니다.
- `reporting`: terminal close가 아니라 continuity context를 보고합니다.
- `next-flow`: next action, blocker, self-drive continuation, explicit stop 중 하나로 라우팅합니다.

task completion은 턴을 닫지 않습니다.
source-recorded explicit stop만 terminal close를 허용할 수 있습니다.

`turn-gate`가 active이면 reporting과 next-flow reopening은 ongoing conversation channel에 남아야 합니다.
terminal/final closeout은 현재 사용자 메시지가 명시적으로 턴을 끝내고 그 closure source가 기록된 뒤에만 허용됩니다.

## 기록 계약

각 phase 시작과 종료에는 `flow` phase record checkpoint expectation을 적용합니다.
`000-plan.md`는 active flow pointer, turn-level required next action, active skill list, self-drive status, unapproved action state 같은 date-level routing 변화가 있을 때 갱신합니다.
active flow record는 같은 flow 내부 phase state, evidence, report, residual risk가 바뀔 때 갱신합니다.
flow 시작 지점에서 다시 읽을 skill 목록은 `000-plan.md` frontmatter의 `active_skills`에 이름만 반영합니다.

## 하위 계약 연결

- 날짜 해석과 기록 기반 날짜 충돌은 `contracts/date-authority.md`가 소유합니다.
- active flow 도중 새 사용자 메시지의 entry-only routing은 `contracts/interruption.md`가 소유합니다.
- next-flow reopening, post-flow continue, question abort recovery는 `contracts/question-routing.md`가 소유합니다.
- verification method/status와 non-pass routing은 `contracts/verification.md`가 소유합니다.
- session record shape와 recovery는 `contracts/session-records.md`가 소유합니다.
- self-drive overlay는 `contracts/self-drive.md`와 runtime `references/self-drive.md`가 소유합니다.

## 활성화 계약

사용자가 `turn-gate`만 호출하면 conversation-level rule을 활성화하고, operating state를 기록하며, scope 또는 next-flow routing을 엽니다.
terminal activation summary만 답하지 않습니다.

activation은 다음을 포착해야 합니다.

- current active flow 또는 activation flow
- latest user request
- 현재 또는 planned flow에서 다시 읽어야 하는 skill 목록
- turn-gate active state
- no explicit-stop source unless user actually stopped the turn
- terminal summary blocked state
- required next action
- activation record 자체의 verification expectation

## 단계 메시지 계약

user-facing phase-start 또는 meaningful progress 메시지는 canonical prefix로 시작합니다.

- `[intake]`
- `[framing]`
- `[preparation]`
- `[work]`
- `[verification]`
- `[reporting]`
- `[next-flow]`

prefix는 generated artifact, record, command summary, question option label에 복사하지 않습니다.

## 준비와 승인 경계 계약

work 전에는 `turn-gate`가 필수 `flow` decision을 적용하거나 기록해야 합니다.
flow contract는 scope, non-goals, completion 또는 acceptance signal, verification expectation, approval boundary, handoff condition을 포함해야 하며, 상세 판단은 `flow`가 소유합니다.

새 active flow를 시작할 때는 해당 flow에 적용할 skill을 다시 읽고, 이전 flow에서 로드된 skill context만으로 work를 시작하지 않습니다.
사용자 메시지를 받아 framing을 거쳐 preparation으로 들어갈 때는 preparation 전에 `turn-gate`와 `flow`를 다시 읽고 `000-plan.md` frontmatter의 `active_skills`에 둘 다 유지합니다.

approval-sensitive action에는 exact target, expected effect, risk, recovery path, included/excluded scope, endpoint가 필요합니다.
readiness, verification, generated release surface build/readback, self-drive, previous context, subagent output은 commit, push, PR, publish, release, version bump, destructive history rewrite, external side effect의 실행 권한을 만들 수 없습니다.

## runtime 본문 경계

Runtime `SKILL.md`는 다음을 직접 포함해야 합니다.

- active-turn rule, terminal summary 금지, required ending states
- `flow` dependency와 turn-gate owned responsibilities
- compact lifecycle
- relative date 기본값과 기록 기반 충돌의 clarification 경계
- interruption entry-only phase의 적용 방식
- flow 시작 지점의 skill reread
- phase prefix behavior
- preparation과 approval boundary
- verification method와 status 구분
- reporting과 next-flow reopening
- compact question abort recovery
- runtime `references/self-drive.md` discoverability

상세 decision table은 runtime references나 owning contract로 내립니다.
Runtime `SKILL.md`는 설치된 사용자에게 dev-only `specs/` 경로를 읽으라고 지시하거나 spec-side scenario fixture를 복사하지 않습니다.
