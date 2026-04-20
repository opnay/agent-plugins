# AGENTS.global.md

이 문서는 현재 활성화된 로컬 플러그인을 글로벌 지침에서 어떻게 강조할지 정리하는 운영용 가이드입니다.
목적은 플러그인 이름을 나열하는 것이 아니라, 작업을 받았을 때 어떤 기준으로 어떤 플러그인을 먼저 보게 할지 우선순위를 고정하는 것입니다.

## 기본 원칙

- 추측보다 질문을 통한 확정이 필요합니다.
- 단, 이는 실행에 필요한 목표, 범위, 비목표, 승인 경계가 아직 잠기지 않았을 때에 적용합니다.
- 현재 활성 플러그인이 작업과 맞으면 일반 추론보다 해당 플러그인의 guide skill을 먼저 고려합니다.
- 단, 도메인 키워드보다 먼저 `지금 막힌 일의 성격`을 봅니다.
- 플러그인 선택은 `도메인`보다 `현재 task shape`와 `현재 병목`을 기준으로 합니다.
- 여러 플러그인이 동시에 맞아 보이면, 더 좁은 specialist skill보다 먼저 해당 플러그인의 guide skill로 들어가 starting point를 고릅니다.
- sibling skill 이름이 다른 skill 본문에 하드코딩되어 있다고 해서 그것만으로 trigger 근거로 삼지 않습니다.
- skill trigger의 1차 근거는 사용자 요청, 현재 task shape, 그리고 각 skill description입니다.

## 현재 활성 플러그인

- `workflow-kit`
  - 역할: framing, definition, planning, execution mode, review loop, readiness gate 같은 workflow stage를 고릅니다.
  - 기본 진입점: `workflow-kit-guide`
  - 먼저 볼 때:
    - 작업이 아직 애매해서 무엇부터 해야 할지 불안정할 때
    - 바로 구현보다 먼저 intent, scope, tradeoff를 잠가야 할 때
    - 구현 전에 read-only planning이 먼저일 때
    - 구현 중이지만 review loop나 commit gate로 넘어가야 할 때
  - 강조 포인트:
    - 도메인 작업이어도 `정의가 먼저 필요한 상태`라면 specialist plugin보다 먼저 봅니다.
    - 특히 신규 프로젝트, 신규 시스템, repo 구조, 플랫폼 방향, acceptance boundary가 안 잠긴 요청은 우선 `workflow-kit` 가능성을 먼저 검사합니다.

- `frontend-engineering-kit`
  - 역할: frontend workflow, architecture pattern, React 구조, component boundary, domain modeling, UI implementation quality, TDD를 다룹니다.
  - 기본 진입점: `frontend-engineering-kit-guide`
  - 먼저 볼 때:
    - dominant concern이 이미 frontend-specific할 때
    - 구현 또는 구조 결정의 중심이 React, component, token, UI quality, frontend architecture일 때
    - 이미 무엇을 만들지보다는 어떻게 frontend로 설계/구현할지가 핵심일 때
  - 강조 포인트:
    - frontend라는 이유만으로 항상 먼저 고르지 않습니다.
    - 프로젝트 정의, scope lock, tradeoff alignment가 먼저면 `workflow-kit`이 앞섭니다.

- `designer-kit`
  - 역할: design brief, UI critique, pre-code screen specification 같은 design-only pre-code 작업을 다룹니다.
  - 기본 진입점: `designer-kit-guide`
  - 먼저 볼 때:
    - 아직 구현보다 디자인 방향 정리, 화면 비평, pre-code spec이 필요한 상태일 때
    - 코드나 React 구조보다 디자인 결과물 정의가 먼저일 때
  - 강조 포인트:
    - 디자인만 다루는 단계라면 `frontend-engineering-kit`보다 먼저 봅니다.

- `teammate-kit`
  - 역할: teammate-style collaboration, durable orchestration, bounded research/implementation/review role을 다룹니다.
  - 기본 진입점: `teammate-kit-guide`
  - 먼저 볼 때:
    - 작업의 핵심이 협업 모드 설계일 때
    - 한 bounded teammate role로 끝낼지, 여러 teammate orchestration이 필요한지 판단해야 할 때
  - 강조 포인트:
    - generic workflow selection의 대체재가 아닙니다.
    - raw subagent capability 언급만으로 곧바로 `teammate-kit`를 고르지 않습니다.

- `advance-codex`
  - 역할: skill, plugin, tool-use guidance, custom agent, session surface, commit surface 같은 Codex 표면을 다룹니다.
  - 기본 진입점: `advance-codex-guide`
  - 먼저 볼 때:
    - 새 skill을 만들거나 기존 skill을 고칠 때
    - plugin bundle 구조를 설계/수정할 때
    - tool-use policy를 domain workflow 밖으로 분리할 때
    - custom agent, session continuity, commit finalization surface를 다룰 때
  - 강조 포인트:
    - 제품 구현 플러그인이 아니라 Codex surface 플러그인입니다.

## 우선순위 규칙

### 1. Stage before domain

- 먼저 지금 필요한 것이 definition인지, planning인지, execution인지, review인지, gate인지 판단합니다.
- 이 판단이 아직 안 끝났다면 `workflow-kit`을 먼저 고려합니다.
- 도메인 플러그인은 stage가 잠긴 뒤에 들어갑니다.

### 2. Design-only before frontend implementation

- 코드 구현보다 디자인 방향, critique, screen specification이 먼저면 `designer-kit`을 먼저 고려합니다.
- React, token, component, structure 얘기가 섞여 있어도 결과물이 pre-code design artifact라면 `designer-kit`이 앞섭니다.

### 3. Frontend specialist only after dominant concern is clear

- `frontend-engineering-kit`은 frontend domain 신호가 보인다는 이유만으로 바로 고르지 않습니다.
- dominant concern이 실제로 architecture, React, component, UI implementation, TDD 중 하나로 좁혀졌을 때 진입합니다.

### 4. Collaboration mode is separate from workflow mode

- 협업 구조를 결정하는 문제는 `teammate-kit`에서 다룹니다.
- generic execution workflow selection은 `workflow-kit`에서 다룹니다.
- 둘을 섞지 않습니다.

### 5. Codex surface work is separate from product work

- skill, plugin, agent, tool policy, session, commit surface 변경은 `advance-codex`를 먼저 봅니다.
- 제품 기능 구현 문제를 `advance-codex`로 보내지 않습니다.

## trigger 판단 규칙

- 사용자가 plugin 또는 skill을 직접 언급하면 그 plugin/skill을 우선합니다.
- 직접 언급이 없으면, 현재 task shape와 skill description이 가장 잘 맞는 guide skill을 고릅니다.
- guide skill은 plugin 안에서 starting point를 고르기 위한 진입점이므로, specialist skill보다 먼저 검토합니다.
- 비-guide skill은 자기 description이 직접 맞을 때만 트리거합니다.
- 다른 skill 본문에 특정 sibling skill 이름이 적혀 있어도, 그것만으로 trigger 근거로 삼지 않습니다.

## 신규 프로젝트에 대한 글로벌 판단 규칙

- 신규 프로젝트, 신규 시스템, 신규 repo, 신규 design system, 신규 platform setup 요청은 도메인 키워드만 보고 specialist plugin으로 바로 들어가지 않습니다.
- 먼저 다음을 확인합니다:
  - goal이 잠겼는지
  - in-scope / out-of-scope가 잠겼는지
  - acceptance boundary가 있는지
  - tradeoff와 non-goal이 명확한지
- 위 항목이 충분히 잠기지 않았다면, 우선 `workflow-kit`의 definition/alignment 계열을 먼저 고려합니다.
- 그 다음에야 `frontend-engineering-kit`나 `designer-kit` 같은 domain plugin으로 내려갑니다.

## 예시 라우팅

- "신규 디자인시스템 시작하려는데 React, RSC, monorepo, showcase 3종으로 가고 싶다. 어떻게 생각해?"
  - 기본 판단: frontend 키워드는 강하지만, 아직 정의와 경계가 더 중요하므로 `workflow-kit` 우선
  - 이후 정의가 잠기면 `frontend-engineering-kit`로 handoff

- "기존 React repo에서 component boundary랑 hook 구조를 어떻게 나눌지 봐줘"
  - 기본 판단: dominant concern이 이미 frontend architecture이므로 `frontend-engineering-kit` 우선

- "이 화면 방향이 맞는지 비평해줘"
  - 기본 판단: `designer-kit` 우선

- "새 plugin 만들고 guide skill도 같이 설계해줘"
  - 기본 판단: `advance-codex` 우선

- "이 작업을 여러 teammate role로 굴릴지, 한 명 역할로 끝낼지 정해줘"
  - 기본 판단: `teammate-kit` 우선

## 글로벌 가드레일

- 플러그인을 단순 키워드 매칭으로 고르지 않습니다.
- specialist plugin이 맞아 보여도, 아직 stage 판단이 안 끝났으면 `workflow-kit` 가능성을 먼저 확인합니다.
- domain specialist plugin 안에 generic workflow routing 책임을 숨기지 않습니다.
- guide skill이 맡아야 할 routing을 sibling skill 본문에 하드코딩하지 않습니다.
- 하나의 요청이 여러 플러그인에 걸쳐 보여도, 한 번에 여러 specialist skill을 동시에 여는 대신 시작점 하나와 handoff 순서를 먼저 정합니다.
