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
  - 독립적으로 검증 가능한 완결 산출물과 도메인·조정·기반 책임을 포함한 모듈 경계 판단
  - 기존 패턴 우선, 작은 변경, 명시적 계약, 실패 모드 분리
  - 위험도에 맞는 검증과 residual risk 보고
- 제외:
  - 특정 언어, 프레임워크, 라이브러리의 세부 구현 레시피
  - 프로젝트별 코딩 컨벤션 대체
  - tool selection, escalation, approval policy
  - commit finalization 절차
  - subagent handoff/lifecycle 절차
  - 제품 요구사항이나 도메인 정책 결정
  - 모든 구현을 별도 단계, package, shared library로 만드는 일반 규칙
  - 도메인·조정·기반 책임을 고정된 계층, 폴더, 파일 구조로 강제하는 아키텍처 규칙

## 대표 표면

- 대표 runtime 표면: `judgment-kit-dev/skills/pro-engineering/SKILL.md`
- 사용자 스펙 의도: `judgment-kit-dev/specs/skills/pro-engineering/intent.md`
- skill spec index: `judgment-kit-dev/specs/skills/pro-engineering/spec.md`
- sub-spec directory: `judgment-kit-dev/specs/skills/pro-engineering/`

## 상세 계약 구조

- `intent.md`: 사용자 스펙 의도 기록
- `core/problem-solving.md`: 문제 정의, 증거 수집, 원인 후보, 가설 검증 루프
- `core/scope-control.md`: 조사와 범위 설정은 넓게, 구현은 좁게 진행하고 큰 범위를 완결 산출물로 나누는 breadth-to-focus 규칙
- `core/engineering-judgment.md`: 기술 판단, 소유권, 도메인·조정·기반 책임, 모듈 경계, 재사용 조건, 허용 조건 우선 계약, 추상화, 실패 모드, 리스크 배수 기준
- `core/code-discipline.md`: 코드 작성 방식, 완결 산출물의 일관된 상태, 소유자 가까이의 단순 구현 후 개선, 변경 범위, 제어 흐름 명확성, 사용자 변경 보존
- `core/verification-reporting.md`: 검증 범위 선택, residual risk, 보고 형식
- `core/runtime-surface.md`: runtime 독립성, description trigger metadata, dev/release surface 검증

## 확장 원칙

- 새로운 child spec은 실제 runtime 판단 책임이 커져 index를 읽기 어렵게 만들 때만 추가합니다.
- 사용자 의도는 `intent.md`에만 둡니다.
- child spec은 자기 세부 계약만 소유하고 상위 의도를 반복하지 않습니다.
- runtime skill은 folderized spec 전체를 짧고 실행 가능한 지시로 압축해야 합니다.
