# turn-gate skill-contents sub-spec

## 목적

이 문서는 `turn-gate` runtime `SKILL.md` 본문이 반드시 담아야 하는 실행 계약과 구성 기준을 소유합니다.
`spec.md`는 skill 전체 index와 sibling spec map을 소유하고, 이 문서는 설치 후 실제로 읽히는 skill body의 내용 계약만 소유합니다.

## 경계

- 포함:
  - `SKILL.md` body의 필수 section과 우선순위
  - runtime body에 직접 남아야 하는 turn continuity, phase flow, verification, reporting 계약
  - runtime body에 넣지 않아야 하는 dev-only fixture와 spec-side 평가 절차
- 제외:
  - 전체 phase 전환 세부 계약 자체
  - 각 internal gate의 세부 판단 규칙
  - internal gate model의 runtime-facing 노출
  - session record template의 정확한 필드 목록
  - local runtime reference 문서의 상세 내용

## Runtime Surface Role

- `loop-kit-dev/skills/turn-gate/SKILL.md`는 이 스펙의 단순 요약본이 아니라 runtime에서 읽는 운영 표면이다.
- skill body는 설치된 skill만 읽는 fresh runtime reader가 즉시 따를 수 있는 실행 지시여야 한다.
- skill body를 짧게 다듬더라도 필수 단계와 금지 규칙을 한 문단으로 뭉개지 말고, 실행 중 빠르게 확인 가능한 형태로 유지한다.

## Required Section Priority

- skill body는 대화 응답 자체를 제어하는 conversation-level first-class rule을 `## Important` 섹션으로 앞부분에서 명시해야 한다.
- `## Important` 섹션은 `Purpose`보다 먼저 위치해야 한다.
- `## Important` 섹션은 최소한 다음 내용을 포함해야 한다.
  - session-level activation
  - terminal summary 금지
  - required ending states
  - `request_user_input` 기반 next-flow reopening
  - session record 유지 의무

## Core Loop Content

- skill body에는 `Core Loop`라는 기존 섹션명을 유지할 필요가 없다.
- 대신 runtime reader가 즉시 실행할 수 있는 간결한 operating cycle을 두고, 최소한 다음 단계가 한 번에 읽히게 해야 한다.
  - preparation
  - work
  - verification
  - reporting
  - next-flow
- skill body에는 `core/runtime-flow.md`의 전체 흐름과 `phase-protocols/routes.md`의 local `references/` 읽기 규칙이 직접 남아 있어야 한다.
- skill body에는 phase 시작을 알리는 사용자-facing 메시지가 `[<phase-name>(/<phase-protocol>)]` 접두사로 시작해야 한다는 규칙이 직접 남아 있어야 한다.
- skill body의 phase prefix 규칙은 canonical phase labels `preparation`, `work`, `verification`, `reporting`, `next-flow`를 제시해야 한다.
- skill body의 phase prefix 규칙은 `(/<phase-protocol>)` segment가 optional이고, 실제 출력에서는 literal parenthesis가 아니라 slash suffix를 사용한다는 점을 설명해야 한다.
- skill body의 phase prefix 예시는 phase-only form과 phase/protocol form을 모두 포함해야 한다.
- skill body는 reporting 뒤 다음 flow를 여는 단계를 `next-flow` phase로 설명해야 한다.
- skill body는 activation-only, mid-work status, session-record blocker, report-only evaluation처럼 여러 phase label이 가능해 보이는 상황의 우선순위를 설명해야 한다.
- skill body는 이 prefix가 phase-start message에 적용되는 운영 표식이며, flow record, output artifact, command summary, question option 전체 문장에 기계적으로 붙이는 규칙이 아님을 설명해야 한다.

## Runtime Structure Boundary

- runtime body는 internal gate model을 사용자-facing 구조로 설명하지 않는다.
- runtime body에는 이전 사용자 메시지 routing 또는 intake layer처럼 보이는 섹션, 단계, gate, 분류 절차를 넣지 않는다.
- runtime body는 phase flow, work boundary, verification, reporting, next-flow, stop rule 중심으로 재구성한다.
- 개별 task 완료는 flow 완료나 turn closure를 결정할 수 없다는 규칙은 남긴다.
- reporting 뒤에는 explicit stop이 source-recorded되지 않는 한 `next-flow` phase가 next-flow reopening으로 이어져야 한다.

## Preparation Content

- skill body는 deep-interview, flow list design, meaning resolution, current-state inspection을 `preparation`의 세부 방식으로 설명해야 한다.
- skill body는 요청 해석과 planned flow list 설계가 plan/session record를 소유하는 `operational-preparation flow`가 될 수 있다고 설명해야 한다.
- skill body는 operational-preparation 결과로 만들어지는 실행용 planned flows가 검토 가능하거나 commit-sized인 `change-unit flow`여야 한다고 설명해야 한다.
- skill body는 요청 해석 결과가 바로 실행으로 이어지지 않을 수 있고, 후속 실행 후보와 실제 실행 flow를 구분해야 한다는 일반 원칙만 설명한다.
- skill body는 preparation에서 scope가 비어 있거나 너무 넓거나 여러 결과물을 만들 수 있거나 성공 기준과 검증 경로를 바꿀 수 있으면 work 전에 질문으로 scope를 잠그도록 직접 설명해야 한다.
- skill body는 질문 없이 추론한 scope라도 work boundary와 non-goal을 flow record에 남기도록 설명해야 한다.

## Approval Content

- skill body는 preparation이 planned flow list 전체를 실행하는 데 필요한 intent, scope, non-goal, acceptance signal, verification expectation을 수집하도록 설명해야 한다.
- skill body는 `turn-gate`의 기본 loop, phase protocol selection, approval-sensitive execution boundary를 독립적으로 설명해야 한다.
- skill body는 approval-sensitive action의 exact target, expected effect, risk, rollback or recovery 가능성, 포함/제외 scope, 종료 지점이 기록돼야 한다고 설명해야 한다.
- skill body는 readiness reporting과 execution authority를 분리해 설명해야 한다.
- skill body는 self-drive가 명시적으로 요청된 prepared sequence에 대해서는 `references/self-drive.md`를 읽어야 한다는 discoverability를 제공해야 한다.
- skill body는 self-drive 세부 조건을 반복하지 않고, 해당 reference를 읽으면 그 overlay 계약이 준비된 sequence의 진행 판단을 소유한다고만 설명해야 한다.

## Verification Content

- skill body에는 clean-context verification이 full-history fork가 아니라 bounded verification packet이라는 점이 직접 남아 있어야 한다.
- skill body에는 실패, 차단, 불충분 검증을 통과로 취급하지 않는 규칙이 직접 남아 있어야 한다.
- skill body는 verifier packet에 필요한 최소 정보, edit permission 없음, scope expansion 금지, destructive/external work 금지, commit/push/PR/publish/release/version bump 금지를 설명해야 한다.

## Reporting And Next-Flow Content

- skill body에는 terminal summary 금지 규칙이 직접 남아 있어야 한다.
- skill body에는 source message에 묶인 confirmed closure 규칙이 직접 남아 있어야 한다.
- skill body에는 next-flow reopening, Continuity Guard 확인, user-gated next-flow choice, explicit turn-end option 기록 규칙이 직접 남아 있어야 한다.
- reporting은 완료 요약으로 턴을 닫는 단계가 아니라 다음 flow 진행을 위한 continuity context 정리로 설명해야 한다.

## Runtime/Spec Boundary

- runtime skill body는 설치 후 실제로 존재하는 `references/`와 `templates/`만 실행 중 읽기 대상으로 안내한다.
- dev-only spec 경로는 runtime 사용자가 읽어야 하는 실행 지시로 쓰지 않는다.
- dev-side fixture, 평가 시나리오, change history는 skill body의 행동 계약을 만드는 근거로만 사용하고, runtime body에는 일반 실행 규칙으로만 반영한다.
- `intent-scenarios/` fixture 이름이나 fixture 평가 절차는 runtime skill body에 직접 넣지 않는다.

## 검토 질문

- skill body 앞부분에 `Important` 섹션이 있고 1급 규칙, terminal summary 금지, next-flow reopening이 먼저 드러나는가?
- skill body가 `preparation -> work -> verification -> reporting -> next-flow` 흐름을 실행 중 빠르게 확인 가능한 형태로 드러내되 기존 routing 중심 구조를 반복하지 않는가?
- phase 시작을 알리는 사용자-facing 메시지가 `[<phase-name>(/<phase-protocol>)]`으로 시작해야 하고 protocol segment가 optional이라는 규칙이 runtime body에 직접 드러나는가?
- runtime body가 internal gate model이나 사용자 메시지 intake/routing layer를 사용자-facing 구조로 다시 열지 않는가?
- spec-side fixture 평가 규칙이 runtime skill body로 직접 누출되지 않았는가?
- runtime body가 설치 후 존재하지 않는 dev-only spec 파일을 읽으라고 지시하지 않는가?
