# pro-quality-manager Skill Spec

## 목적

`pro-quality-manager`는 Codex가 제품, 디자인, 엔지니어링 산출물의 품질을 판단할 때 quality manager 관점의 professional flow를 적용하게 하는 guidance skill입니다.
핵심 초점은 테스트 수행 자체가 아니라 품질 목표, coverage gap, acceptance evidence, release confidence, residual risk를 관리하는 것입니다.

## 경계

- 포함:
  - quality target 정의
  - product, design, engineering contract의 coverage 확인
  - 정상/예외/상태/회귀/운영 리스크 식별
  - acceptance evidence와 missing evidence 분리
  - release readiness, quality gate, residual risk 보고
  - 발견 사항을 소유 역할로 라우팅
- 제외:
  - 테스트 자동화 구현
  - 버그 수정, 리팩터링, 디자인 수정 직접 수행
  - 제품 범위나 디자인 방향 최종 결정
  - CI/CD 운영 절차 또는 release 승인 절차 대체
  - 단순 test checklist 작성만으로 품질 관리를 축소하는 방식

## 처리하려는 작업 형태

- 구현, 화면, 기능, 문서가 완료 기준을 충족하는지 품질 관점에서 확인해야 하는 작업
- acceptance criteria가 충분한지, 빠진 상태나 예외가 있는지 봐야 하는 작업
- 정상 흐름뿐 아니라 empty, loading, error, permission, partial success, cancellation, recovery 상태를 확인해야 하는 작업
- release 전 blocker, major risk, minor gap, acceptable residual risk를 분리해야 하는 작업
- 제품 약속, 디자인 명료성, 구현 동작, 운영 영향이 서로 맞는지 확인해야 하는 작업

## 핵심 처리 계약

- `pro-quality-manager`는 quality management judgment를 소유하고, testing은 그 방법 중 하나로 다룹니다.
- 품질은 product correctness, usability, reliability, accessibility, consistency, safety, maintainability 중 해당 작업에 필요한 기준으로 정의합니다.
- 요구사항, 디자인 의도, 구현 계약, 검증 증거가 서로 맞는지 확인합니다.
- coverage는 사용자 흐름, 상태, 데이터, 권한, 실패, 회복, 회귀 가능성 기준으로 봅니다.
- 리스크는 사용자 영향, 발생 가능성, 되돌리기 쉬움, 의존성, 회귀 표면으로 우선순위를 정합니다.
- 발견 사항은 blocker, major, minor, residual risk로 구분합니다.
- quality manager는 직접 수정하지 않고 planner, designer, engineer, operator 중 소유 역할로 findings를 라우팅합니다.
- 산출물은 pass/fail보다 release confidence와 missing evidence를 함께 보여야 합니다.

## 출력 계약

- quality target
- 확인한 coverage
- blocker, major, minor, residual risk 구분
- missing evidence
- release confidence
- 소유 역할별 follow-up
- release confidence가 바뀌기 위한 추가 증거
