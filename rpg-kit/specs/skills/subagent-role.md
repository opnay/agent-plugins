## 사용자 스펙 의도

- 역할이 할당된 subagent를 spawn하는 규칙을 명시하고 싶다.
- subagent와 어떻게 동작하는지 관찰하고 배워가는 skill이 필요하다.
- subagent role을 즉흥적으로 붙이지 않고 packet과 answer contract로 다루고 싶다.
- role은 단순 직함이 아니라 전문성 조합이어야 한다.
- 예: `PM 역할의 게임 시나리오 담당자`, `엔지니어 역할의 프론트엔드 개발자`, `게임 기획 출신 프론트엔드 개발자`, `게임 PM 출신 PO`.

---

# subagent-role 스킬 스펙

## 목적

`subagent-role`은 role-assigned subagent를 사용하기 전에 role fit, packet, spawn boundary, answer contract, integration rule, learning note를 정의하는 skill입니다.

## 경계

- 포함:
  - subagent role fit 판단
  - role specialty 조합 정의
  - role packet 작성
  - spawn boundary와 승인 조건 확인
  - answer contract 정의
  - caller integration 및 verification rule 정의
  - subagent behavior learning note 작성
- 제외:
  - 모든 task에 대한 무조건 delegation
  - 승인 없이 subagent spawn policy를 우회하는 지침
  - subagent output을 caller 검증 없이 최종 답으로 사용하는 방식
  - domain-specific implementation plan 전체

## 처리하려는 작업 형태

- 여러 관점의 조사나 비교가 caller의 다음 작업과 병렬로 진행될 수 있는 작업
- 직무, 책임 영역, 출신 배경, 범용 전문성을 조합해 다른 관점의 subagent가 필요한 작업
- disjoint write scope를 가진 bounded implementation slice를 role로 나눌 수 있는 작업
- 검증, critique, exploration처럼 output contract가 명확한 sidecar task
- subagent behavior 자체를 관찰하고 다음 role 설계를 개선하려는 task

## 엔트리포인트 / 대표 표면

- 대표 표면: `rpg-kit/skills/subagent-role/SKILL.md`
- 관련 상위 라우팅: `rpg-kit-guide`

## 핵심 처리 계약

- role fit을 먼저 판단하고, subagent가 필요하지 않으면 caller가 직접 처리하도록 권한다.
- role은 loose persona가 아니라 role specialty 조합으로 정의한다.
- spawn이 적절하면 role packet을 먼저 작성한다.
- role packet에는 role name, role specialty, objective, context, ownership, non-goals, expected output, verification signal, integration rule, stop condition을 포함한다.
- coding work를 맡길 때는 write scope를 분리하고, 다른 작업자가 있을 수 있음을 명시한다.
- subagent answer는 caller가 검증하고 통합할 evidence로 취급한다.
- subagent behavior에서 얻은 관찰은 learning note로 남긴다.

## Role Specialty

- 역할 전문성은 필요한 축만 조합해 정의한다.
  - `functional_role`: PM, PO, engineer, designer, reviewer, researcher처럼 agent가 맡을 현재 직무 역할
  - `responsibility_domain`: game scenario, frontend implementation, prompt evaluation처럼 이번 task에서 책임지는 영역
  - `background_expertise`: game-planning-origin, game-PM-origin, frontend-origin처럼 판단 관점에 영향을 주는 출신 배경
  - `general_expertise`: product strategy, UX writing, implementation rigor처럼 특정 직무명보다 넓은 범용 전문성
  - `decision_style`: product sense, technical critique, verification처럼 답변에서 우선할 판단 방식
- 좋은 role specialty는 서로 다른 축을 조합해 관점을 선명하게 만든다.
- 같은 축을 반복하는 표현은 피한다.
- 예시:
  - `PM 역할의 게임 시나리오 담당자`: `functional_role=PM`, `responsibility_domain=game scenario`
  - `엔지니어 역할의 프론트엔드 개발자`: `functional_role=engineer`, `responsibility_domain=frontend implementation`
  - `게임 기획 출신 프론트엔드 개발자`: `functional_role=engineer`, `responsibility_domain=frontend implementation`, `background_expertise=game-planning-origin`
  - `게임 PM 출신 PO`: `functional_role=PO`, `background_expertise=game-PM-origin`, `general_expertise=product strategy`

## Role Pattern Catalog

- `게임 PM 출신 PO`
  - 사용 시점: 제품 판단과 게임 맥락이 함께 필요한 기능 우선순위, onboarding, economy, retention 관련 판단
  - 필수 출력: product tradeoff, player impact, retention risk, prioritization tradeoff, recommended decision
  - 통합 규칙: caller가 제품 목표, player data, product assumption에 맞춰 최종 제품 결정을 검증한다
- `PM 역할의 게임 시나리오 담당자`
  - 사용 시점: 시나리오가 제품 목표, progression, player motivation과 맞는지 봐야 할 때
  - 기대 출력: scenario intent, coherence risk, player-facing issue, revision direction
- `엔지니어 역할의 프론트엔드 개발자`
  - 사용 시점: UI 구현 범위, component boundary, state flow, testability를 나눠 봐야 할 때
  - 필수 출력: implementation plan, file ownership, integration risk, verification path
  - 통합 규칙: coding work가 있으면 disjoint write scope와 concurrent changes guard를 포함하고, caller가 staged changes 또는 final patch를 직접 검토한다
- `게임 기획 출신 프론트엔드 개발자`
  - 사용 시점: game UX나 interaction fantasy를 프론트엔드 구현으로 옮길 때
  - 기대 출력: interaction interpretation, implementation constraints, UX risk, fallback option
- `디자이너 역할의 UX 리뷰어`
  - 사용 시점: 화면 구조, affordance, visual hierarchy, empty/error states를 검토해야 할 때
  - 필수 출력: UX findings, severity, proposed adjustment, acceptance signal
  - non-goals: 직접 구현 변경, 범위 밖 redesign, material issue와 무관한 visual polish 확장을 제한한다
  - 통합 규칙: caller가 material issue만 통합할지 판단한다
- `QA 역할의 회귀 테스트 설계자`
  - 사용 시점: 변경 후 어떤 사용자 경로와 edge case를 검증해야 하는지 정해야 할 때
  - 기대 출력: test matrix, high-risk flows, regression cases, pass/fail signal

## Spawn Boundary

- user intent, runtime policy, tool policy가 subagent spawn을 허용하는지 확인한다.
- immediate blocker를 subagent에게 넘겨 caller가 기다리기만 하는 구조를 피한다.
- 중복 delegation을 피하고, 서로 다른 역할이나 disjoint write scope에만 나눈다.
- destructive, external, approval-boundary action은 subagent에게 승인된 것으로 간주하지 않는다.

## Answer Contract

- subagent에게 최종 답 대신 caller가 통합할 structured result를 요구한다.
- 결과에는 findings, changed files 또는 inspected surfaces, confidence, open questions, verification performed, integration recommendation을 포함한다.
- low confidence나 context gap은 caller가 회수할 수 있는 missing evidence와 explicit approval boundary를 구분해 적게 한다.

## Learning Contract

- role이 너무 넓었는지, output이 통합 가능했는지, caller가 다시 조사해야 했는지 기록한다.
- role specialty가 실제 답변 관점에 영향을 줬는지 기록한다.
- 반복되는 성공 role은 reusable role pattern 후보로 남긴다.
- 실패한 role은 role boundary, context packet, expected output 중 무엇이 모호했는지 분리해 기록한다.

## 검토 질문

- 이 subagent role이 caller의 critical path를 줄이는가?
- role specialty가 단순 직함보다 더 구체적인 판단 관점을 제공하는가?
- role packet이 단독으로 실행 가능한가?
- output contract가 caller integration에 충분히 구조화되어 있는가?
- subagent 결과를 검증할 방법이 있는가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 예.
- 그렇다면 왜 필요한가 / 아니라면 어떤 sibling context를 허용하는가: `subagent-role`은 guide 없이도 role packet과 answer contract를 작성할 수 있어야 한다. 다만 plugin-level routing은 `rpg-kit-guide`가 소유한다.

## 확장 원칙

- 새 role pattern은 반복적으로 유용성이 확인될 때만 spec 또는 reference로 승격한다.
- role 종류를 늘릴 때도 핵심 계약은 specialty, packet, boundary, answer, learning 다섯 축으로 유지한다.
