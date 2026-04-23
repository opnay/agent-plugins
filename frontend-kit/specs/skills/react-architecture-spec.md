## 사용자 스펙 의도

- react-architecture 스킬
- 도메인(페이지, 전역 상태 등) 축, 디자인(디자인 시스템, 카드 컴포넌트 등) 축, 시스템(리액트, 브라우저) 축을 관심사 분리를 해야한다.
- 컴포넌트, 훅, 유틸리티는 시스템 축 -> 디자인 축 -> 도메인 축 순서의 중요도를 갖는다.
- 스펙 혹은 기능 구현시엔 디자인 축 -> 도메인 축 -> 시스템 축 순서의 구현을 하며, 리팩토링 및 코드 정리시 중요도 순서로 재배치 해야한다.
- 컴포넌트 props 설계시 함수는 `on-` prefix 형태의 이벤트 리스너로 취급해야되며, 명칭의 관점은 props를 소유하게될 컴포넌트의 관점입니다.
- ex. "프로젝트 생성" 모달 컴포넌트의 confirm 버튼 클릭시 명칭으로, "모달 컴포넌트"의 관점으로 props를 설계해 `props.onProjectConfirm` -> `props.onConfirm`가 되어야한다.
- 컴포넌트의 상태 혹은 데이터 전달 방식으로 context, props driling, 외부 상태관리, 백엔드 api 커넥터 등 props로 모든걸 해결하지 말고, 내부 훅을 통해 해려하는 방식도 고려해야만 합니다.

---

# react-architecture 스킬 스펙

## 목적

`react-architecture`는 선택된 frontend structure 안에서 React-facing artifact를 `system`, `design`, `domain` 축으로 분리하고, 그 구조 중요도와 구현 순서를 정하는 스킬입니다.
핵심은 component, hook, utility, page, global state가 React/브라우저 메커니즘, 재사용 UI 계약, 제품 의미를 한 단위에 뒤섞지 않도록 만드는 것입니다.

## 경계

- 포함:
  - component, hook, utility, page, global state의 `system` / `design` / `domain` 축 분류
  - React context scope, effect discipline, rendering/update boundary의 `system` 축 판단
  - reusable UI contract와 product meaning의 `design` / `domain` 축 분리
  - component props event-listener naming contract 판단
  - props, context, broader state, backend connector, internal hook 사이의 전달 방식 판단
  - React-facing artifact의 dependency direction 판단
  - 구현 단계와 refactor 단계의 순서 분리
- 제외:
  - top-level pattern choice
  - 시각적 완성도나 디자인 품질 자체를 설계하는 일
  - business rule 자체를 설계하는 일
  - 로컬 split 기법만의 미세 리팩터링 전술

## 처리하려는 작업 형태

- React 코드에서 여러 성격의 책임이 섞여 있어 placement 기준이 필요한 경우
- component, hook, utility, page, global state의 ownership을 축 기준으로 정리해야 하는 경우
- 구현 순서와 정리 순서를 분리해서 다뤄야 하는 경우
- effect, rerender spread, hook API 설계를 domain/design 책임 분리와 함께 다뤄야 하는 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `frontend-kit/skills/react-architecture/SKILL.md`
- 관련 상위 라우팅: `frontend-kit-guide`

## 핵심 처리 계약

- 전체 pattern choice가 이미 정해졌다는 전제에서, 먼저 각 책임을 `system`, `design`, `domain` 중 하나로 분류한다.
- 분류 기준은 "왜 이 코드가 존재하는가"와 "무엇이 바뀌면 이 코드가 바뀌는가"를 기준으로 삼는다.
- `system`은 React, browser, runtime 제약 때문에 존재하는 책임이다.
- `design`은 reusable UI contract, interaction model, design-system semantics 때문에 존재하는 책임이다.
- `domain`은 page, feature, workflow, global state, business meaning 때문에 존재하는 책임이다.
- dependency direction은 `domain -> design -> system`을 기본으로 삼는다.
- `system`은 `design`이나 `domain`을 알지 않아야 한다.
- `design`은 `system`을 사용할 수 있지만 `domain`을 포함하면 안 된다.
- `domain`은 필요하면 `design`과 `system`을 조합할 수 있다.
- hook도 예외가 아니며, `domain` hook은 `design`/`system` hook을 조합할 수 있고 `design` hook은 `system` hook을 조합할 수 있지만 역방향은 허용하지 않는다.
- component, hook, utility를 정리할 때의 구조 중요도는 `system -> design -> domain`이다.
- component function props는 event listener로 보고 `on-` prefix 형태를 기본으로 삼는다.
- event prop 명칭은 그 props를 소유하는 컴포넌트의 관점에서 정하고, 부모 workflow의 의미를 기본값으로 밀어 넣지 않는다.
- 상태나 데이터 전달 문제를 props 하나로만 해결하려 하지 않고, context, prop drilling, broader external state, backend connector, internal hook을 모두 후보로 비교해야 한다.
- public prop contract를 불필요하게 넓히는 것보다 내부 hook으로 local coordination을 흡수하는 편이 더 나은지 항상 확인해야 한다.
- 새 spec 또는 기능 구현을 진행할 때의 구현 순서는 `design -> domain -> system`이다.
- refactor 및 code cleanup 시의 재배치 순서는 `system -> design -> domain`이다.
- mixed hook을 리팩토링할 때는 effect, subscription, observer, storage/URL sync 같은 `system` 책임을 먼저 분리하고, 그 다음 reusable interaction model을 `design`으로 분리한 뒤 feature workflow와 product meaning을 `domain` 경계에 남긴다.
- artifact type이나 파일 이름으로 축을 정하지 않고, 실제 존재 이유로 축을 분류해야 한다. `page`, `global state`, custom hook 같은 이름은 힌트일 뿐 기준이 아니다.

## 컴포넌트 계약 규칙

- function props는 기본적으로 event listener contract로 취급하고 `on-` prefix 형태를 사용한다.
- event prop 명칭은 그 props를 소유하게 될 child component의 관점에서 정한다.
- 부모 feature나 page의 business meaning은 child component의 public prop 이름에 기본값처럼 밀어 넣지 않는다.
- 예를 들어 "프로젝트 생성" 흐름에서 쓰이는 modal이라도, modal boundary가 confirm action만 노출한다면 `onProjectConfirm`보다 `onConfirm`이 기본값이다.
- child component의 contract에 실제로 domain meaning이 내장되어 있을 때만 더 구체적인 이름을 허용한다.
- props name은 boundary event를 설명해야지, parent orchestration이나 use-case 전체를 설명하려고 해선 안 된다.

## 상태 및 데이터 전달 판단 규칙

- 상태나 데이터 전달 문제를 props만으로 해결하려 하지 않는다.
- 다음 후보를 모두 비교 대상으로 본다:
  - props
  - prop drilling
  - context
  - broader external state
  - backend API connector
  - internal hook
- internal hook은 local coordination, local integration, local state shaping을 public prop contract 밖으로 숨겨서 boundary를 더 작게 만들 수 있을 때 우선 후보가 될 수 있다.
- prop 수가 계속 늘어나거나, props가 단순 API보다 orchestration transport 역할을 하기 시작하면 boundary smell로 본다.
- 전달 방식을 고를 때는 access convenience보다 ownership clarity를 우선한다.
- public prop contract를 유지하는 편과 internal hook으로 local detail을 흡수하는 편 중 어느 쪽이 axis leakage를 더 줄이는지 비교해야 한다.
- context, broader state, backend connector도 자동 정답이 아니며, 각 방식이 현재 owner를 더 선명하게 만드는지로 판단한다.

## 훅 분류 및 리팩토링 규칙

- custom hook이라는 형태 자체는 축이 아니며, hook 이름이나 위치로 분류하지 않는다.
- hook의 축은 "무엇이 먼저 바뀌면 이 hook이 먼저 깨지거나 이동해야 하는가"로 판단한다.
- effect, subscription, observer, browser/runtime sync가 주된 이유라면 `system` hook이다.
- reusable interaction semantics, reusable UI behavior contract가 주된 이유라면 `design` hook이다.
- feature workflow, page orchestration, product-facing state machine이 주된 이유라면 `domain` hook이다.
- 한 hook 안에 여러 축이 섞여 있다면 리팩토링의 첫 단계는 naming 변경이 아니라 `system` 책임 분리다.
- local component 안에서만 쓰는 internal hook도 synchronization 중심이면 `system`, reusable interaction 중심이면 `design`, feature meaning 중심이면 `domain`으로 본다.
- `domain` hook은 필요하면 `design`/`system` hook을 조합할 수 있지만, `system` hook은 `domain`의 용어와 정책을 알면 안 된다.

## 검토 질문

- 이 function prop은 component boundary event인가, 아니면 parent workflow 이름이 새어 나온 것인가?
- 이 이름은 child component가 자기 contract를 설명하는 말인가?
- 지금 추가하려는 prop은 public API인가, 아니면 local coordination detail인가?
- local hook이 있으면 public prop contract를 더 작고 읽기 좋게 만들 수 있는가?
- context나 broader state가 실제 shared owner를 설명하는가, 아니면 prop 전달을 피하려는 편의적 선택인가?
- 이 hook에서 먼저 떼어내야 할 것은 feature 이름이 아니라 effect, subscription, observer 같은 `system` 책임인가?
- 이 hook이 reusable interaction contract를 소유하는가, 아니면 feature workflow를 소유하는가?

## 독립성 원칙

- 이 스킬은 top-level structure나 domain model 도입 여부를 소유하지 않는다.
- 이 스킬은 축 분류, 의존 방향, 구현 순서, 재배치 순서만으로도 독립적으로 읽히어야 한다.
- 기존 sibling skill 구성이 달라져도 이 스킬의 판단 기준은 유지 가능해야 한다.

## 확장 원칙

- 새로운 rule은 `system` / `design` / `domain` 축 분리와 의존 방향을 더 선명하게 만들 때만 추가한다.
- 특정 컴포넌트군이나 기존 sibling skill을 방어하기 위한 예외 규칙은 기본 경로로 두지 않는다.
