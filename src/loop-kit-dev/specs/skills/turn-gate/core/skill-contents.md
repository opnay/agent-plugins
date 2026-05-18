# turn-gate skill-contents sub-spec

## 목적

이 문서는 `turn-gate` runtime skill folder가 설치 후 사용자와 agent에게 제공해야 하는 실행 표면의 내용 계약을 소유합니다.

runtime skill folder는 다음 세 가지 표면으로 구성됩니다.

- `SKILL.md`: activation 직후 읽히는 최상위 실행 지시
- `references/`: `SKILL.md`에 모두 펼치면 너무 길어지는 세부 운영 계약
- `templates/`: session record를 새로 만들 때 쓰는 runtime 시작점

`spec.md`는 skill 전체 index와 sibling spec map을 소유합니다.
이 문서는 runtime skill folder의 내용 구성과 보존 기준만 소유합니다.

## 소유

- `SKILL.md` body의 필수 top-level section과 우선순위
- `SKILL.md`에 직접 남아야 하는 turn continuity, phase flow, verification, reporting, next-flow 계약
- `SKILL.md`가 local `references/`와 `templates/`를 어떻게 discoverable하게 만드는지
- skill folder 재생성 시 runtime `references/`와 `templates/`에서 보존해야 하는 운영 계약
- runtime body에 넣지 않아야 하는 dev-only fixture, spec-side 평가 절차, internal gate model 노출 제한

## 비소유

- phase 전환의 전체 세부 계약: `core/runtime-flow.md`, `phases/*`, `gates/*`가 소유합니다.
- session record의 전체 운영 계약: `records/session-records.md`가 소유합니다.
- template의 정확한 field 형식: `templates/plan.md`, `templates/flow.md`, `templates/self-drive.md`가 소유합니다.
- self-drive 세부 판단: `modes/self-drive.md`와 runtime `references/self-drive.md`가 소유합니다.
- Stop/SessionStart hook의 세부 input/output 및 packaging: plugin-level hook spec과 bundled implementation이 소유합니다.

## Runtime Surface 계약

- `loop-kit-dev/skills/turn-gate/SKILL.md`는 spec 요약본이 아니라 설치 후 실제로 읽히는 운영 표면입니다.
- fresh runtime reader가 dev-only spec을 열지 않아도 즉시 행동할 수 있어야 합니다.
- 문장을 다듬을 수는 있지만 필수 단계, 금지 규칙, 승인 경계, 검증 경계, next-flow reopening 규칙을 일반론으로 뭉개면 안 됩니다.
- `SKILL.md`는 `references/`와 `templates/`가 설치 후 존재하는 local runtime 표면임을 안내해야 합니다.
- `SKILL.md`는 dev-only `specs/`, `intent-scenarios/`, change history를 실행 중 읽으라고 지시하면 안 됩니다.
- skill folder 재생성은 `SKILL.md`만 새로 쓰는 작업이 아닙니다. 함께 설치되는 `references/`와 `templates/`의 판단 기준, 소유권 경계, recovery state, approval limit도 보존해야 합니다.

## SKILL.md 구조 계약

`SKILL.md`는 다음 top-level section 이름과 순서를 유지합니다.

1. `## Important`
2. `## Purpose`
3. `## Phase Messages`
4. `## Operating Cycle`
5. `## Session Records`
6. `## Optional Hooks`
7. `## Common Misclassifications`

위 section은 서로 흡수하거나 합치지 않습니다.
예를 들어 `Session Records`, `Optional Hooks`, `Common Misclassifications`의 내용을 `Important`나 `Operating Cycle` 안에만 녹여 없애면 안 됩니다.

`## Important`는 `Purpose`보다 먼저 위치하며, 최소한 다음 내용을 직접 포함합니다.

- `turn-gate` activation은 conversation-level first-class operating rule입니다.
- explicit stop 없이 terminal summary로 끝내면 안 됩니다.
- reporting 뒤에는 explicit stop 또는 recorded self-drive continuation이 없는 한 `request_user_input`으로 next-flow를 다시 열어야 합니다.
- session records는 turn continuity 표면이며, 사용자가 all-file/no-record write 금지를 명시하지 않는 한 유지해야 합니다.
- trusted bundled Stop hook은 terminal closure backstop일 수 있지만, conversation-level rule, reporting, record refresh, next-flow reopening을 대체하지 않습니다.
- trusted bundled SessionStart hook context는 startup/restart 힌트일 뿐이며 approval, continuation, terminal closure authority가 아닙니다.

## Core Loop Content

`SKILL.md`에는 `preparation -> work -> verification -> reporting -> next-flow` 흐름이 즉시 보이도록 남아야 합니다.
기존 section 이름 `Core Loop` 자체를 유지할 필요는 없지만, `## Operating Cycle` 아래에서 각 phase가 수행할 일을 짧고 실행 가능한 형태로 설명해야 합니다.

`SKILL.md`는 다음 규칙을 직접 설명합니다.

- preparation은 intent, scope, non-goal, acceptance signal, verification expectation, approval boundary를 정렬합니다.
- 요청 해석과 planned flow list 설계 자체가 `operational-preparation` flow가 될 수 있습니다.
- 실행용 planned flow는 검토 가능하거나 commit-sized인 `change-unit` flow여야 합니다.
- ambiguous operation 또는 target이 work surface, output shape, success criteria, verification path를 바꿀 수 있으면 work 전에 질문으로 잠급니다.
- 질문 없이 scope를 추론한 경우에도 work boundary와 non-goal을 flow record에 남깁니다.
- approval-sensitive action은 exact target, expected effect, risk, rollback/recovery path, 포함/제외 scope, stopping point가 기록돼야 합니다.
- readiness reporting은 execution authority가 아닙니다.
- individual task completion은 flow completion이나 turn closure를 결정하지 못합니다.
- reporting은 완료 요약으로 턴을 닫는 단계가 아니라 다음 flow 진행을 위한 continuity context입니다.
- reporting 뒤 explicit stop이 source-recorded되지 않으면 `next-flow` phase가 next-flow reopening으로 이어집니다.

## Phase Message Content

`SKILL.md`는 phase 시작 또는 phase 진행 상황을 알리는 사용자-facing 메시지 prefix 규칙을 직접 설명합니다.

- canonical phase label은 `preparation`, `work`, `verification`, `reporting`, `next-flow`입니다.
- notation은 `[<phase-name>(/<phase-protocol>)]`입니다.
- phase protocol 사용 시 slash suffix를 사용합니다.
- 실제 출력에는 literal parenthesis를 쓰지 않습니다.
- phase-only 예시와 phase/protocol 예시를 모두 포함합니다.
- prefix는 phase-start/progress message에 적용되는 운영 표식입니다.
- flow record, generated artifact body, command output summary, question option label 전체에 기계적으로 전파하지 않습니다.
- activation-only, mid-work status, session-record blocker, report-only evaluation처럼 여러 label이 가능해 보이는 상황의 우선순위를 설명합니다.
- self-drive continuation의 status, verification, reporting, automatic next-flow handoff 같은 사용자-facing phase/progress message에는 prefix를 쓰되, `000-self-drive.md`, flow record, 생성 산출물, question option label 안으로 오염시키지 않습니다.

## Runtime Structure Boundary

`SKILL.md`는 runtime reader에게 필요한 운영 흐름을 보여주되, spec-side 내부 모델을 다시 노출하지 않습니다.

- internal gate model을 사용자-facing 구조로 설명하지 않습니다.
- 이전 사용자 메시지 routing layer, intake layer처럼 보이는 section이나 gate 분류 절차를 만들지 않습니다.
- phase flow, work boundary, verification, reporting, next-flow, stop rule 중심으로 유지합니다.
- phase protocol은 mode가 아니라 current phase에 적용하는 local contract로 설명합니다.
- self-drive는 `turn-gate` 본체가 모두 소유하는 mode가 아니라 prepared sequence overlay로 설명하고, 세부 판단은 `references/self-drive.md`로 위임합니다.
- common misclassification은 phase-vs-flow, commit completion, self-drive status question, future endpoint stop, file-change verification default처럼 반복 실수를 줄이는 작은 decision aid로 제한합니다.

## Verification Content

`SKILL.md`에는 verification method와 result status의 분리가 직접 남아야 합니다.

- method는 `clean-context`, `normal`, `not-required` 중 하나입니다.
- result status는 method가 아니며 `pass`, `fail`, `blocked`, `insufficient` 같은 결과 판단입니다.
- file change, release surface, multi-file contract, previous failure, user-requested verification, approval-sensitive action에는 `clean-context`가 기본값입니다.
- `clean-context`는 full-history fork가 아니라 bounded verification packet입니다.
- verifier packet은 target, expected behavior, changed surface, relevant files/commands, constraints, explicit prohibitions를 포함해야 합니다.
- verifier는 edit permission을 받지 않은 한 수정하지 않으며, scope expansion, destructive/external work, commit, push, PR, publish, release, version bump를 수행하지 않습니다.
- `not-required`는 pass status가 아니라 method 판단이며, reason과 residual uncertainty를 기록해야 합니다.
- `fail`, `blocked`, `insufficient`는 pass로 보고하지 않고 repair phase, blocker report, 또는 user-gated routing으로 돌려야 합니다.

## Session Record Content

`SKILL.md`는 session record의 세부 field를 모두 반복하지 않고, runtime reader가 기록 표면을 찾고 최소 규칙을 지킬 수 있을 만큼만 설명합니다.

- `000-plan.md`는 compact date-level recovery snapshot과 flow index입니다.
- `001+` flow record는 flow-local scope, non-goals, approval boundary, execution log, verification evidence, report, residual risk, `Continuity Guard`, `Next Flow Options`를 소유합니다.
- self-drive active sequence는 필요할 때 `000-self-drive.md` sidecar를 사용합니다.
- phase step을 별도 flow record로 바꾸지 않습니다. 하나의 flow는 보통 preparation, work, verification, reporting, next-flow를 함께 가집니다.
- target/source read-only는 session record write 금지와 같지 않습니다.
- workspace-wide no-write, no-record 요청이 있으면 session record write도 금지하고 clarification 또는 blocker로 라우팅합니다.
- 세부 record recovery와 read-only/no-record boundary는 `references/session-records.md`로 위임합니다.

## Runtime Reference Contracts

`references/`는 `SKILL.md`에서 모두 펼치면 길어지는 판단 기준을 runtime에 남기는 표면입니다.
재생성은 reference를 단순 요약본으로 낮추면 안 됩니다.

- `references/phase-protocols.md`는 phase protocol이 local phase contract이며, target/meaning/approval boundary가 work surface를 바꾸면 preparation 또는 question routing으로 돌아간다는 기준을 유지합니다.
- `references/session-records.md`는 `000-plan.md`, optional `000-self-drive.md`, `001+` flow record의 소유권 분리를 유지합니다.
- `references/session-records.md`는 `not-yet-created`, `unexpectedly missing`, `inaccessible`, `stale closure state`, `stale routing mismatch`, `stale self-drive sidecar` 같은 recovery state를 구분합니다.
- `references/session-records.md`는 stale closure, stale sidecar, missing/inaccessible record, routing mismatch가 terminal closure, successful completion, next-flow skip, autonomous continuation authority가 아니라고 설명합니다.
- `references/session-records.md`는 recovery state마다 use case, allowed action, forbidden action을 충분히 구분합니다. 표 형식일 필요는 없지만, 상태 이름만 나열하는 요약으로 낮추면 안 됩니다.
- `references/session-records.md`는 completed flow compaction, parent-plan relation, stale current-state field cleanup처럼 active recovery context를 오염시키지 않기 위한 운영 기준을 유지합니다.
- `references/session-records.md`는 target/source read-only와 workspace-wide no-write/no-record 제약을 분리합니다.
- `references/session-records.md`는 `operational-preparation` flow와 `change-unit` flow, follow-up candidate와 active execution의 차이를 유지합니다.
- `references/self-drive.md`는 self-drive를 prepared sequence overlay로 설명하고, standalone skill이나 전체 user-message taxonomy로 넓히지 않습니다.
- `references/self-drive.md`는 sequence-level state를 `000-self-drive.md`가 소유하고, `000-plan.md`는 self-drive active status와 sidecar pointer만 소유한다고 설명합니다.
- `references/self-drive.md`는 entry condition 또는 record ownership 설명에서 `active_flow_index`, `current_flow_label`, `progress_note`, blocker return conditions를 빠뜨리지 않습니다.
- `references/self-drive.md`는 `active_flow_index`가 0-based machine field이며 human-readable current flow label과 함께 reconcile되어야 한다고 설명합니다.
- `references/self-drive.md`는 mid-sequence user input priority, approval-sensitive action boundary, question-tool boundary, finite endpoint/repeat policy를 유지합니다.
- `references/self-drive.md`는 question-tool boundary와 endpoint behavior를 runtime reader가 상황별로 판단할 수 있게 유지합니다. 표 형식일 필요는 없지만, clear prepared transition, status-only input, scope/target/endpoint/order change, approval-sensitive action, blocker/record failure/current-flow ambiguity, non-self-drive reporting을 구분해야 합니다.

## Runtime Template Contracts

`templates/`는 새 session record의 runtime 시작점입니다.
재생성은 template의 표현을 다듬을 수 있지만, machine-readable field와 기록 책임을 바꾸면 안 됩니다.

- `templates/flow-record-template.md`는 `templates/flow.md`가 정한 frontmatter field를 유지합니다.
- 특히 `turn_gate_session_id`, closure fields, pending question fields, `verification_status`, `continuity_note`, `preparation_source`, `scope_lock_status`를 제거하지 않습니다.
- `templates/flow-record-template.md`는 `# Flow Record` title과 `## Flow Contract`, `## Optional Risky Actions`, `## Execution Log`, `## Verification`, `## Report`, `## Next Flow Options`, `## Residual Risk` heading depth를 유지합니다.
- `templates/flow-record-template.md`는 raw user request와 summary/interpretation을 분리할 수 있어야 합니다.
- `templates/flow-record-template.md`는 work boundary, non-goals, acceptance signal, expected risky actions, approval boundary, user-gated checkpoints, verification expectation, material judgment calls를 기록할 자리를 유지합니다.
- `templates/flow-record-template.md`는 optional risky action section 또는 collapsed not-applicable marker를 유지합니다.
- `templates/flow-record-template.md`는 verification method와 result status를 분리합니다.
- `templates/plan-template.md`는 `# Session Plan` title과 `## Current State`, `## Flow Table`, `## Planned Flow Sequence`, `## Open Risks`, `## Turn-End Rule` heading depth를 유지합니다.
- `templates/plan-template.md`는 bounded date-level snapshot/index 역할을 유지하고, per-flow scope, evidence, verification detail을 소유하는 것처럼 보이면 안 됩니다.
- `templates/self-drive-template.md`는 `# Self-Drive Sequence` title과 `## Sequence Contract`, `## Autonomous Boundary`, `## Progress Ledger`, `## User-Gated Return Conditions`, `## Residual Risk` heading depth를 유지합니다.
- `templates/self-drive-template.md`는 `current_flow_label`, `progress_note`, sequence contract, allowed/prohibited actions, approval-sensitive checkpoints, approval notes, blocker return conditions, append-only progress ledger를 유지합니다.

## Optional Hook Content

`SKILL.md`는 optional hooks를 runtime 보조 장치로만 설명합니다.

- Stop hook은 source-recorded explicit stop 없는 terminal closure를 막는 backstop입니다.
- Stop hook이 block하면 main agent는 active flow record와 `000-plan.md`를 refresh하고 hook reason의 `required_next_action`으로 돌아갑니다.
- Stop hook은 기록을 대신 수정하지 않고, destructive/external action, commit, push, PR, publish, release, version bump, global config, plugin hook setup을 승인하지 않습니다.
- Stop hook의 detailed block/quiet-exit 조건은 runtime skill body에서 반복하지 않습니다.
- SessionStart hook은 startup/resume context를 제공할 수 있지만, current user message와 active record 재확인을 대체하지 않습니다.
- SessionStart context는 approval, continuation, terminal closure authority를 만들지 않습니다.

## 검토 질문

- `SKILL.md` 앞부분에 `Important`가 있고 first-class rule, terminal summary 금지, next-flow reopening이 먼저 드러나는가?
- `SKILL.md` top-level section skeleton이 안정적으로 유지되는가?
- `SKILL.md`가 `preparation -> work -> verification -> reporting -> next-flow` 흐름을 실행 가능한 형태로 드러내는가?
- phase prefix 규칙이 runtime body에 직접 드러나며, record/artifact/question label 오염을 막는가?
- runtime body가 internal gate model이나 broad user-message routing layer를 다시 열지 않는가?
- self-drive 세부 조건을 본체에 반복하지 않고 `references/self-drive.md`로 위임하는가?
- verification method와 result status가 분리되어 있는가?
- `references/`가 decision criteria, ownership boundary, recovery state, approval limit을 잃지 않았는가?
- `templates/`가 machine-readable field와 기록 책임을 잃지 않았는가?
- runtime body가 설치 후 존재하지 않는 dev-only spec 파일을 읽으라고 지시하지 않는가?
