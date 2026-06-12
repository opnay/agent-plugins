# pro-engineering Skill Spec

## 목적

`pro-engineering`은 Codex가 코드 작성과 문제 해결을 할 때 엔지니어 관점의 판단을 더 명시적으로 적용하게 하는 guidance skill입니다.
핵심 초점은 문제를 정확히 정의하고, 실제 증거를 기준으로 원인 후보를 좁히고, 가장 작은 완전한 수정과 검증 가능한 결과로 마무리하는 것입니다.

## 경계

- 포함:
  - 증상과 기대 동작 분리
  - 실제 코드, 로그, 테스트, 재현 조건을 기준으로 한 사실 수집
  - 원인 후보 나열, 가설 검증, 반례 확인
  - 단순한 완전 구현 뒤 해당 턴에서 점진적으로 개선하는 코드 작성 규율
  - 기존 패턴 우선, 작은 변경, 명시적 계약, 실패 모드 분리
  - 위험도에 맞는 검증과 residual risk 보고
- 제외:
  - 특정 언어, 프레임워크, 라이브러리의 세부 구현 레시피
  - 프로젝트별 코딩 컨벤션 대체
  - tool selection, escalation, approval policy
  - commit finalization 절차
  - subagent handoff/lifecycle 절차
  - 제품 요구사항이나 도메인 정책 결정

## 처리하려는 작업 형태

- 코드 작성, 버그 수정, 리팩터링, 설계 판단에서 문제 해결 절차를 명시해야 하는 작업
- 사용자의 설명과 실제 코드/증거가 다를 수 있어 원인 분석을 먼저 해야 하는 작업
- 수정 범위, 추상화 여부, 테스트 범위, 리스크 보고 기준을 정해야 하는 작업
- 기술적 판단을 "왜 그렇게 하는지"까지 설명해야 하는 작업

## 대표 표면

- 대표 runtime 표면: `judgment-kit-dev/skills/pro-engineering/SKILL.md`
- 사용자 스펙 의도: `judgment-kit-dev/specs/skills/pro-engineering/intent.md`
- skill spec index: `judgment-kit-dev/specs/skills/pro-engineering/spec.md`
- sub-spec directory: `judgment-kit-dev/specs/skills/pro-engineering/`

## 상세 계약 구조

- `intent.md`: 사용자 스펙 의도 기록
- `core/problem-solving.md`: 문제 정의, 증거 수집, 원인 후보, 가설 검증 루프
- `core/scope-control.md`: 조사와 범위 설정은 넓게, 구현은 좁게 진행하는 breadth-to-focus 규칙
- `core/engineering-judgment.md`: 기술 판단, 소유권 경계, 허용 조건 우선 계약, 추상화, 실패 모드, 리스크 배수 기준
- `core/code-discipline.md`: 코드 작성 방식, 단순 구현 후 개선, 변경 범위, 제어 흐름 명확성, 사용자 변경 보존
- `core/verification-reporting.md`: 검증 범위 선택, residual risk, 보고 형식

## 핵심 처리 계약

- `pro-engineering`은 철학 문서가 아니라 작업 중 바로 적용할 수 있는 판단 절차입니다.
- 분석은 사용자 설명보다 실제 코드, 로그, 테스트, 재현 조건 같은 증거를 우선합니다.
- 직접 코딩 전 조사와 범위 설정에서는 관련 시스템, 사용자 흐름, 대안, 위험을 넓게 탐색합니다.
- 구현은 원인 후보를 충분히 좁힌 뒤 가장 작은 완전한 변경으로 시작합니다.
- 구현 중에는 선택한 원인과 acceptance signal에 직접 연결되는 코드 경로로 주의를 좁힙니다.
- 변경이 책임 경계를 건드리면 새 동작, 상태, 검증, 대체 동작을 소유하는 경계에 둡니다.
- 예외와 금지 목록을 늘리기 전에 허용되는 입력, 상태, 효과, 실패 조건을 먼저 정의하고, 차단 조건은 그 계약을 보완하는 좁은 방어선으로 둡니다.
- 조건문은 guard, 실패 조건, 정상 흐름을 분리해 읽기 쉽게 만들고, 중첩을 줄일 때 early return을 선호합니다.
- 개선은 같은 턴 안에서 명확성, 이름, 구조, 엣지 처리 순서로 점진적으로 수행합니다.
- 검증은 위험도와 영향 범위에 맞게 선택하고, 수행하지 못한 검증은 이유와 함께 보고합니다.

## 독립성 원칙

`pro-engineering`은 독립 실행 가능한 runtime skill이어야 합니다.
본문은 sibling skill 이름이나 dev-only spec 경로를 읽으라고 지시하지 않습니다.
다른 `judgment-kit` skill과 함께 쓰일 수는 있지만, 문제 해결과 코드 작성 판단 자체는 이 skill 본문만으로 수행 가능해야 합니다.

## Description Trigger Metadata

이 skill은 passive skill로 선택될 수 있어야 합니다.
frontmatter `description` 끝에는 `#` 없는 쉼표 구분 plain token 목록을 둡니다.
권장 token tail:

`engineering judgment, problem solving, root cause analysis, technical reasoning, code quality, implementation discipline`

## 검증 기준

- dev runtime skill이 `skills/pro-engineering/SKILL.md`에 존재해야 한다.
- release build 후 root `judgment-kit/skills/pro-engineering/SKILL.md`가 존재해야 한다.
- plugin spec, README, manifest prompt가 `pro-engineering`의 역할과 사용 기준을 언급해야 한다.
- runtime skill 본문은 dev-only `specs/` 경로나 `src/judgment-kit-dev` 경로를 실행 지시로 포함하지 않아야 한다.
- runtime skill은 문제 해결 중심이며 특정 언어/프레임워크 레시피로 좁아지지 않아야 한다.

## 확장 원칙

- 새로운 child spec은 실제 runtime 판단 책임이 커져 index를 읽기 어렵게 만들 때만 추가합니다.
- 사용자 의도는 `intent.md`에만 둡니다.
- child spec은 자기 세부 계약만 소유하고 상위 의도를 반복하지 않습니다.
- runtime skill은 folderized spec 전체를 짧고 실행 가능한 지시로 압축해야 합니다.
