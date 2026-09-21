## 사용자 스펙 의도

- 구현을 위해 기술을 선택하는 레이어와 코드를 작성하는 레이어를 분리하고 싶다.
- 엔지니어링은 요구사항에 맞는 기술, 데이터 모델, 구조 방향을 정하는 데 집중하고 싶다.
- 코드 최적화와 구조 개선은 code-quality가 맡게 하고 싶다.
- 제품 기획자는 사용자 문제와 요구사항을 소유하며, 기술 설계와는 다른 역할이어야 한다.
- 소프트웨어 작업을 하나의 플러그인으로 묶되 skill 이름은 짧게 하고 싶다.

---

# SW Kit 플러그인 스펙

## 목적

`sw-kit`는 제품 또는 사용자 요구를 실행 가능한 소프트웨어로 만들 때 필요한 행동 계약, 기술 방향, 코드 품질, 코드베이스 유지보수를 제공한다.

## 경계와 비목표

- 포함: 소프트웨어 행동 계약, 기술·데이터·운영 선택, 코드 구현과 리뷰, 코드베이스 수명 관리.
- 제외: 사용자 문제·가치·우선순위·MVP의 제품 기획, 디자인 판단, Git workflow, 배포 실행, 일반 문서 작성.
- `$judgment-kit:pro-planner`가 만든 handoff를 사용할 수 있지만, 어느 sibling도 선행 호출을 요구하지 않는다.

## 내장 Skill

- `spec`: 사용자·제품 요구를 구현 가능한 행동 계약으로 정리한다.
- `engineering`: 데이터 모델, 구조, 기술, integration, migration, 운영 방향을 선택한다.
- `code`: 선택된 방향 안에서 코드를 구현·수정·리뷰한다.
- `maintenance`: dependency, deprecation, dead code, drift, debt를 관리한다.

## 사용 표면

- runtime: `sw-kit/skills/<skill-name>/SKILL.md`
- plugin spec: `src/sw-kit-dev/specs/plugin.md`
- skill spec: `src/sw-kit-dev/specs/skills/<skill-name>.md`

## 확장 원칙

- 새 skill은 네 책임 중 어느 것도 자연스럽게 소유하지 못하는 반복 작업일 때만 추가한다.
- product planning, Git, design, quality management의 책임을 중복하지 않는다.
- sibling은 명확한 산출물을 입력으로 받을 수 있으나 독립적으로도 실행 가능해야 한다.
