# turn-gate runtime-flow sub-spec

## 목적

이 문서는 `turn-gate`가 하나의 턴을 어떻게 이어가는지에 대한 전체 phase 흐름과 phase 간 전환 조건을 소유합니다.

## 전체 흐름

`turn-gate`의 기본 flow는 아래 순서를 유지합니다.

1. preparation
2. work
3. verification
4. reporting

activation, incoming message classification, next-flow reopening, explicit stop handling은 이 기본 flow를 둘러싼 lifecycle guard입니다.
deep-interview alignment, flow list design, meaning resolution, current-state inspection, target reread, scope lock, approval boundary 확인은 기본적으로 `preparation` 안의 세부 작업입니다.

`turn-gate`에는 두 층의 flow가 있습니다.
첫째, `operational-preparation flow`는 사용자 메시지를 받아 intent, scope, non-goal, acceptance signal, verification expectation, approval boundary를 잠그고 planned flow list를 만드는 운영 flow입니다.
이 flow의 산출물은 `.agents/sessions/{YYYYMMDD}/000-plan.md`, active flow record, planned flow list, approval boundary note 같은 session/plan artifact입니다.
둘째, `change-unit flow`는 그 planned flow list 안에서 실제 코드, 문서, fixture, 설정, release surface 같은 검토 가능한 산출물 변경을 소유하는 실행 flow입니다.
사용자 메시지 해석과 flow list 설계는 "멈춤"이나 terminal response가 아니라 운영 flow의 preparation/work로 기록될 수 있으며, 질문을 열어 scope를 잠그는 경우에도 `active question-routing` 상태로 turn을 유지해야 합니다.

`turn-gate`에서 planned flow는 phase가 아니라 응집된 변경 단위입니다.
`분석`, `작업`, `검증`, `보고`, `커밋 준비` 같은 단계명은 flow boundary가 될 수 없고, 각 flow 내부의 core phase 또는 internal mode로 남아야 합니다.
또한 flow가 반드시 최종 사용자에게 직접 보이는 가치 단위일 필요는 없습니다.
flow는 함께 이해하고 검토하고 검증하고 필요하면 커밋할 수 있는 변경 묶음이어야 하며, 예를 들어 "로그인 페이지 만들기"는 `로그인 UI/UX 컴포넌트 생성`, `로그인 로직 작성`, `로그인 페이지 조립` 같은 planned flows로 나뉠 수 있습니다.
순수 최종 QA, 정합성 점검, 검증 결과 보고, commit-readiness reporting은 기본적으로 planned flow boundary가 아닙니다.
그것들은 마지막 변경 단위 flow의 `verification`/`reporting` 또는 별도 user-gated handoff로 남겨야 합니다.
단, 회귀 테스트 fixture, snapshot baseline, 문서, 운영자 리포트 출력, validator 진단 출력처럼 별도 검토 가능한 산출물을 만들거나 바꾸는 경우에는 그 산출물 변경이 planned flow가 될 수 있습니다.

이 문서는 각 phase의 순서와 전환을 소유하고, phase 내부의 세부 판단은 대응 child spec으로 위임합니다.

## Phase 계약

- activation:
  - `turn-gate`가 호출되면 conversation-level first-class operating rule로 활성화한다.
  - concrete task 없이 activation만 요청되면 work mode를 고르지 않고 next-flow 또는 scope selection을 연다.
  - 예: "turn-gate 켜줘", "Use turn-gate", `$loop-kit:turn-gate`만 온 경우에는 activation 완료 요약으로 닫지 않고 다음 scope 또는 next-flow 선택을 연다.
- incoming message classification:
  - 모든 incoming message를 같은 loop-gated turn의 authoritative input으로 본다.
  - 먼저 현재 메시지가 현재 turn 자체를 끝내려는 명시적 요청인지 판단한다.
  - 명시적 turn stop이 아니면 어떤 형태의 입력이든 terminal close 근거가 아니라 continuation input이다.
  - continuation input은 closed taxonomy에 맞추려고 하지 말고 현재 flow에 미치는 효과로 라우팅한다: 현재 상태를 보고하고 계속할지, analysis/plan을 수정할지, target을 다시 읽을지, 다음 flow 후보를 갱신할지, approval boundary를 열지, 검증/검토 작업으로 이어갈지를 결정한다.
  - 질문, 검토 요청, 상태 확인, correction, 우선순위 변경, 다음 작업 요청은 continuation input의 예시일 뿐이며 exhaustive list가 아니다.
  - continuation input이 target file, artifact, state를 바꾸면 이어가기 전에 해당 target을 다시 읽고 stale assumption을 재사용하지 않는다.
  - continuation input이 다음 flow 요청이라면 flow record의 next-flow 후보 중 최우선으로 등록하고 다음 safe handoff point까지 이어간다.
- preparation:
  - 이 flow에서 무엇을 할지, 왜 하는지, 어떤 조건에서 작업으로 넘어갈 수 있는지를 준비한다.
  - 사용자 메시지에서 시작하는 preparation은 deep-interview를 사용해 intent, scope, non-goal, success criteria, approval boundary, verification signal을 정렬한다.
  - 사용자 메시지를 해석하고 planned flow list를 설계하는 준비는 필요하면 `operational-preparation flow`로 기록한다. 이 flow는 plan/session record 산출물을 소유하며, product/code 변경 단위 flow와 같은 층에 섞어 phase처럼 나열하지 않는다.
  - 사용자 메시지의 scope가 비어 있거나 너무 넓거나 여러 결과물을 만들 수 있거나 성공 기준과 검증 경로를 바꿀 수 있으면, work로 넘어가기 전에 user-gated question-routing으로 scope를 먼저 잠근다.
  - 범위 잠금은 최소한 포함 범위, 제외 범위, 대상 파일/표면 또는 산출물, 완료 기준, 검증 신호 중 이번 flow 결과를 바꿀 항목을 다룬다.
  - scope가 충분하다고 추론하는 경우에도 그 추론한 work boundary와 non-goal을 flow record에 남겨야 하며, 추론이 틀리면 되돌리기 어려운 작업은 질문 없이 진행하지 않는다.
  - 사용자 메시지 기반 preparation은 planned flow list 전체를 진행하는 데 필요한 정보가 무엇인지 확인하고, 이후 flow들이 추가 사용자 질문 없이 self-drive로 진행될 수 있을 만큼 intent, scope, non-goal, acceptance signal, approval boundary, verification expectation을 수집한다.
  - approval boundary, destructive/irreversible/external action, commit/push/PR/publish 결정처럼 예상되는 위험 작업은 초기 preparation에서 질문해 승인/비승인 또는 handoff 경계를 별도로 표시하고, self-drive가 자동 처리하지 못하는 user-gated checkpoint로 계획한다.
  - 이후 flow 진행 중 초기 협의 범위 밖의 위험 작업이나 새 approval boundary가 나타나면, self-drive는 자동 처리하지 않고 user-gated question-routing으로 다시 질문해야 한다.
  - 사용자가 self-drive 진행을 원하거나 autonomous continuation이 적합한 경우, planned flow list의 실행은 같은 플러그인의 `turn-gate-self-drive` overlay로 넘길 수 있다.
  - self-driven planned flow sequence가 끝나면 commit execution이 아니라 commit-readiness reporting handoff로 이어진다. commit-readiness reporting 자체는 산출물 변경을 소유하지 않는 한 새 planned flow로 만들지 않는다. commit execution, push, PR, publish는 별도 user-gated handoff다.
  - 사용자 메시지 기반 deep-interview 결과는 단순 질문 답변이 아니라 이후 flow list로 변환되어야 한다.
  - flow list 변환 자체가 산출물로 남는 경우, 그 산출물은 `operational-preparation flow`의 work/reporting 결과이고, 변환 결과 목록의 각 실행 항목은 `change-unit flow`다.
  - flow list design은 phase list design이 아니다. 하나의 flow가 `preparation -> work -> verification -> reporting` 전체를 소유하며, phase나 internal mode를 독립 flow로 승격하지 않는다.
  - flow boundary는 최종 사용자에게 보이는 가치만 기준으로 잡지 않는다. 보이지 않는 UI/UX component scaffolding, domain logic, integration assembly처럼 커밋 가능한 응집 변경 단위도 flow가 될 수 있다.
  - 최종 QA, 통합 검증, 정합성 점검, readiness 보고처럼 변경 산출물 없이 확인과 보고만 수행하는 항목은 flow list에 넣지 않는다. 해당 내용은 active flow의 verification/reporting에 기록한다.
  - 비 사용자 메시지에서 시작하는 preparation은 이미 준비된 flow의 실행 전 준비이며, 필요한 수정 범위, 현재 상태, 대상 파일, stale assumption, 실행 전 조건을 확인한다.
  - operation/target ambiguity가 flow list나 작업 결과를 바꿀 수 있으면 flow list design이나 work 전에 meaning resolution으로 먼저 잠근다.
  - requested intent, requested action, current blocker, likely internal mode, approval boundary, deep-interview result 또는 current-state preparation result, planned flow list를 확인한다.
  - operation/target ambiguity는 `meaning-resolution.md`가 소유한다.
  - destructive, irreversible, external, commit/publish approval boundary는 `approval-boundary.md`가 소유한다.
  - 현재 flow의 active steps를 정하고, meaningful work가 시작되면 계획 도구를 사용해 현재 상태를 유지한다.
  - multi-flow decomposition과 session plan은 `session-records.md`가 소유한다.
- work:
  - 사용자가 요청한 실제 작업을 진행한다.
  - 작업은 파일 수정, 조사, 검증 실행, 리뷰 finding 처리, 계획 작성처럼 다양한 형태일 수 있다.
  - work에 들어가기 전 current-phase work의 internal mode를 하나 선택한다.
  - mode selection과 local reference 읽기 규칙은 `mode-selection.md`가 소유한다.
- verification:
  - 현재 flow의 work 결과를 검증한다.
  - 파일 수정이라면 적용 여부, 타입 오류, 테스트/빌드/린트 같은 해당 검증 경로를 확인한다.
  - 조사나 판단 작업이라면 다양한 관점에서 논리 비판과 반례 검토를 수행한다.
  - work 뒤 reporting 전에 clean-context verification을 수행한다.
  - 검증 packet, pass/fail/blocked/insufficient 처리, non-pass return path는 `verification.md`가 소유한다.
- reporting:
  - 결과 보고는 terminal response가 아니라 다음 flow 진행을 위한 context 정리다.
  - 이번 flow에서 무엇을 준비/작업/검증했는지, 남은 불확실성과 blocker가 무엇인지 정리한다.
  - residual uncertainty와 blocker가 있으면 사용자에게 보이게 포함한다.
  - 계획된 flow가 소진되면 질문 도구를 사용해 사용자에게 다음 flow나 작업을 받는다.
  - next-flow reopening 세부는 `question-routing.md`가 소유한다.
- next-flow reopening:
  - explicit stop이 없다면 결과 보고 뒤 active question-routing으로 다음 flow를 연다.
  - `request_user_input`, fallback, visible/recorded turn-end option은 `question-routing.md`가 소유한다.
- explicit stop handling:
  - 사용자가 명시적으로 턴 종료를 요청한 경우에만 terminal summary가 가능하다.
  - "여기서 끝", "턴 종료", "이 turn은 그만", "stop the turn"처럼 현재 turn 자체를 끝내려는 의도가 분명한 입력만 explicit turn stop으로 분류한다.
  - 명시적 종료 의도가 불분명하면 종료로 추정하지 말고 continuation input으로 처리하거나 user-gated clarification을 연다.
  - flow record의 `confirmed closure`는 특정 explicit stop 사용자 메시지와 함께 기록된 경우에만 유효하다.
  - closure source message가 없거나 현재 incoming message와 맞지 않는 stale closure 기록은 terminal close 근거로 쓰지 않는다.
  - closure source message와 `Continuity Guard` 기록은 `session-records.md`와 함께 유지한다.

## 검토 질문

- 기본 flow가 `준비 -> 작업 -> 검증 -> 보고`로 한 번에 읽히는가?
- planned flow sequence가 phase checklist가 아니라 응집된 변경 단위들의 순서로 읽히는가?
- 사용자 메시지 해석과 planned flow list 설계가 필요한 경우, 그것이 운영 flow로 기록되고 실행용 change-unit flow 목록과 분리돼 있는가?
- flow boundary가 직접 사용자 가치만으로 좁혀져 보이지 않는 커밋 가능 준비 작업을 배제하지 않는가?
- 최종 QA, 정합성 점검, readiness 보고가 별도 산출물 변경 없이 planned flow로 승격되지 않았는가?
- 각 phase의 세부 판단이 적절한 child spec으로 위임되는가?
- reporting이 terminal close가 아니라 next-flow reopening으로 이어지는가?
- explicit stop 없이 흐름이 닫히는 경로가 남아 있지 않은가?
