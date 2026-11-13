## 사용자 스펙 의도

- 구현을 위해 기술을 선택하는 레이어와 코드를 작성하는 레이어를 분리하고 싶다.
- 엔지니어링은 요구사항에 맞는 기술, 데이터 모델, 구조 방향을 정하는 데 집중하고 싶다.
- 코드 최적화와 구조 개선은 code-quality가 맡게 하고 싶다.
- 제품 기획자는 사용자 문제와 요구사항을 소유하며, 기술 설계와는 다른 역할이어야 한다.
- 소프트웨어 작업을 하나의 플러그인으로 묶되 skill 이름은 짧게 하고 싶다.

---

# Engineer 플러그인 스펙

## 목적

`engineer`는 제품 또는 사용자 요구를 실행 가능한 소프트웨어로 만들 때 필요한 행동 계약, 기술 방향, 코드 품질, 코드베이스 유지보수를 제공한다.

## 경계와 비목표

- 포함: 소프트웨어 행동 계약, 기술·데이터·운영 선택, 코드 구현과 리뷰, 코드베이스 수명 관리.
- 제외: 사용자 문제·가치·우선순위·MVP의 제품 기획, 디자인 판단, Git workflow, 배포 실행, 일반 문서 작성.
- `$planner:plan`가 만든 handoff를 사용할 수 있지만, 어느 sibling도 선행 호출을 요구하지 않는다.

## 내장 Skill

- `spec`: 사용자·제품 요구를 구현 가능한 행동 계약으로 정리한다.
- `engineering`: 데이터 모델, 구조, 기술, integration, migration, 운영 방향을 선택한다.
- `code`: 선택된 방향 안에서 코드를 구현·수정·테스트·리뷰한다. 테스트 작성 gate와 집중 감사·subsystem campaign 계약을 소유한다.
- `maintenance`: dependency, deprecation, dead code, drift, debt와 테스트 부채를 조사하고 정리 방향을 제시한다. 테스트 기준은 `code`의 bundled resource를 사용하며 수정은 요청된 source-level 작업으로 구분한다.

테스트 작성·수정·리뷰는 `code` 본문의 기준을 사용한다. 집중 삭제·통합 감사는 `code/references/test-audit.md`, 전체 subsystem 정리는 `code/references/test-campaign.md`를 읽는다. 일반 코드·maintenance 작업이 감사를 자동 실행하지 않는다.

## 사용 표면

- runtime: `agents/engineer/skills/<skill-name>/SKILL.md`
- plugin spec: `agents/_specs/engineer/plugin.md`
- skill spec: `agents/_specs/engineer/skills/<skill-name>/spec.md`

## 확장 원칙

- 새 skill은 네 책임 중 어느 것도 자연스럽게 소유하지 못하는 반복 작업일 때만 추가한다.
- product planning, Git, design, quality management의 책임을 중복하지 않는다.
- sibling은 명확한 산출물을 입력으로 받을 수 있으나 선행 실행을 요구하지 않는다. bundled resource 공유는 소유 skill을 명시하고 skill spec에서 허용 범위·이유를 설명한다.

## 직업 공통·분야·숙련도

현재 skill은 직업 공통 작업과 기존 전문 설계 책임을 유지합니다. 새로운 분야는 고유 기준·산출물·선택 조건이 있을 때만 추가합니다. 설치된 `assets/proficiency.md`는 practice/independent/expert의 직업별 판단 기대를 소유하며 역할 셋업이 읽고 선택해 명시적으로 적용합니다. 미지정이면 기존 skill 계약으로 실행하고 어느 수준도 필수 요구·권한·검증을 생략하지 않습니다. 숙련도와 분야 전문성·실제 능력·실행 권한을 구분합니다.
