## 사용자 스펙 의도

- Figma 작업에서는 page, section, section part, component, asset 등 대상 design element를 먼저 식별하고 싶다.
- import 전에 target design element를 확정하고 싶다.
- 좌표를 absolute와 parent-relative 관점에서 함께 해석하고 싶다.
- mobile, tablet, desktop, responsive 등 page와 frame size를 고려하고 싶다.
- 구현에서는 document flow를 우선하고 position은 static, sticky, fixed, absolute 순서로 검토하며 layout은 grid와 flex를 먼저 사용하고 싶다.
- 이 guidance를 `app-extensions`의 `figma-ext` skill로 사용하고 싶다.

---

# figma-ext 스킬 스펙

## 목적

Figma capability를 사용하는 작업에 target identification, hierarchy, coordinate context, viewport·responsive interpretation, flow-first implementation rules를 추가합니다.
upstream Figma plugin이 소유하는 tool-use contract를 유지하면서, 디자인 구조를 읽거나 Figma와 code 사이를 변환할 때 필요한 로컬 handoff 기준을 일관되게 적용합니다.

## 경계

- 포함:
  - page, section, section part, component, asset 단위의 target design element 식별
  - import, inspection, write, design-to-code, code-to-design 전에 target과 parent hierarchy 확인
  - canvas absolute coordinates와 parent-relative coordinates의 동시 해석
  - frame·page size, mobile·tablet·desktop·responsive variant 확인
  - document flow와 grid·flex를 우선하는 layout translation
  - static, sticky, fixed, absolute 순서의 position 판단
  - design evidence와 implementation inference 구분
- 제외:
  - Figma plugin manifest, MCP tool, app connection, 인증 구성
  - upstream Figma skill의 prerequisite, tool schema, API syntax 복제 또는 우회
  - 일반 시각 디자인, 브랜딩, 색상, tone, product design quality 판단
  - Figma와 무관한 generic frontend 구현이나 code-quality 판단
  - 근거 없는 responsive variant, coordinate, overlay behavior 생성

## 처리하려는 작업 형태

- Figma node, frame, page를 읽고 구조와 배치를 설명하는 작업
- Figma의 page·section·component·asset을 code layout으로 옮기는 작업
- code나 설명을 Figma의 특정 design element에 반영하는 작업
- absolute canvas position과 parent layout 관계가 함께 필요한 작업
- viewport별 design variant와 responsive behavior를 비교하는 작업

## 엔트리포인트 / 대표 표면

- 대표 표면: `skills/figma-ext/SKILL.md`
- 호출 방식: `$app-extensions:figma-ext`
- passive trigger: Figma layout, Figma import, Figma coordinates, Figma responsive, Figma design to code, Figma code to design

## 핵심 처리 계약

1. 작업 방향을 inspection, Figma-to-code, code-to-Figma, in-Figma edit 중 하나로 식별합니다.
2. read, import, write 전에 exact target design element와 parent hierarchy를 확인합니다.
3. target의 page·frame dimensions와 알려진 viewport 또는 responsive variant를 확인합니다.
4. 배치는 canvas absolute coordinates와 parent-relative coordinates·constraints·layout flow를 함께 해석합니다.
5. 구현 layout은 document flow를 보존하고 grid 또는 flex를 우선합니다.
6. position은 static, sticky, fixed, absolute 순서로 검토하며, 뒤의 방식을 선택할수록 design evidence와 이유를 명시합니다.
7. upstream Figma capability의 prerequisite와 tool-use contract를 따르며 이 skill의 guidance로 우회하지 않습니다.
8. 필요한 capability, target, hierarchy, viewport evidence가 없으면 쓰기나 구현을 추정 실행하지 않고 missing context를 보고합니다.
9. 결과에는 확인한 design facts와 implementation inference를 구분합니다.

## Target Model

- `page`: viewport family나 product surface를 소유하는 최상위 design context
- `section`: page 안에서 독립 목적과 layout zone을 가진 영역
- `section part`: section 내부에서 국소 목적과 layout 책임을 가진 구성 단위
- `component`: 재사용 contract와 variant를 가진 design element
- `asset`: layout contract보다 전달·재사용되는 visual resource가 중심인 element
- 분류가 겹치면 Figma node type만으로 단정하지 않고 현재 작업에서의 역할과 parent 관계를 함께 사용합니다.

## Coordinates And Viewports

- absolute coordinates는 canvas 또는 page에서의 위치 확인에 사용합니다.
- relative coordinates는 parent frame, auto layout, constraints, sibling flow와의 관계 확인에 사용합니다.
- 한 좌표계만으로 구현 방식을 결정하지 않습니다.
- target frame size와 mobile, tablet, desktop 또는 responsive variant의 존재를 확인합니다.
- variant가 없으면 responsive behavior를 design fact로 보고하지 않고 별도 inference로 표시합니다.

## Layout Translation

- normal document flow로 표현 가능한 element는 `static`을 유지합니다.
- scroll context 안에서 고정되는 증거가 있을 때만 `sticky`를 사용합니다.
- viewport에 고정되는 증거가 있을 때만 `fixed`를 사용합니다.
- overlay, free placement, layer composition이 필요한 근거가 있을 때만 `absolute`를 사용합니다.
- 반복되는 row·column·two-dimensional alignment는 grid 또는 flex로 먼저 모델링합니다.
- Figma canvas 좌표를 그대로 CSS absolute coordinates로 옮기지 않습니다.

## Failure Flow

- target이 불명확하면 후보를 나열하고 write 전에 정확한 node나 role을 확인합니다.
- parent hierarchy나 coordinate context를 읽을 수 없으면 배치 구현을 확정하지 않습니다.
- viewport evidence가 없으면 단일 viewport 결과와 responsive 미확인을 구분합니다.
- upstream Figma capability가 없거나 unavailable이면 연결을 임의 설치·변경하지 않고 제한을 보고합니다.

## 검토 질문

- exact target design element와 작업 방향이 확인됐는가?
- page, section, section part, component, asset 중 현재 역할이 분명한가?
- parent hierarchy와 absolute·relative coordinates를 모두 확인했는가?
- frame size와 viewport 또는 responsive variant를 확인했는가?
- grid·flex와 normal flow로 해결 가능한 layout에 absolute positioning을 사용하지 않았는가?
- sticky, fixed, absolute 선택에 design evidence가 있는가?
- upstream Figma prerequisite와 tool-use contract를 우회하지 않았는가?
- design fact와 implementation inference를 구분했는가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 예
- 이유: 설치 후 sibling skill이나 dev-only spec 없이 Figma extension guidance를 이해해야 합니다. 실제 Figma mutation은 available upstream capability에 의존하며, capability가 없으면 명시적으로 중단합니다.

## 확장 원칙

- upstream Figma manual이나 tool schema를 복제하지 않습니다.
- layout과 무관한 별도 extension contract가 반복될 때만 새로운 skill 또는 runtime reference를 검토합니다.
- 특정 framework, CSS library, design system 규칙은 이 skill의 공통 contract로 승격하지 않습니다.
