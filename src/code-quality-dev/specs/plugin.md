## 사용자 스펙 의도

- 아래 스킬 만드는 프롬프트를 참고해서 `code-quality` 스킬을 만들고, 이 레포지토리에 맞춰 새로운 플러그인으로 생성한다. 목적은 정확하고 이해하기 쉬우며 안전하게 변경·검증할 수 있는 생산 코드를 작성, 수정, 리팩터링, 검토하게 하는 것이다. 구현 전에 기존 구현, 프레임워크와 시스템 API, 표준 라이브러리, 설치된 의존성, 기존 유틸리티를 먼저 탐색하되 재사용을 강제하지 않고 적합성을 평가한다. 문서 작성, 번역, 단순 포맷 변경, generated/vendor code에는 기본적으로 호출되지 않게 한다.

---

# Code Quality Dev 플러그인 스펙

## 플러그인 목적

`code-quality-dev`는 production code 작업에서 정확성, 이해 가능성, 변경 안전성, 테스트 가능성, 견고성, 적절한 재사용, 단순성을 판단하게 하는 플러그인입니다.
핵심 책임은 구현, 버그 수정, 리팩터링, 테스트, 의존성 판단, architecture tradeoff, 코드 리뷰를 하나의 code-quality skill 표면으로 제공하는 것입니다.

## 플러그인 경계와 비목표

- 포함:
  - production code 작성, 수정, 리팩터링, 테스트 추가
  - 기존 기능, 프레임워크/시스템 API, 표준 라이브러리, 설치된 의존성, 내부 패키지, 유틸리티 재사용 판단
  - 새 외부 의존성 추가 필요성 및 장기 비용 검토
  - 위험 중심 코드 리뷰
  - 검증 명령 탐색과 관련 검증 실행
- 제외:
  - prose-only 문서 작성
  - 번역
  - 단순 포맷, 스타일, 이름 변경
  - `$code-quality` 호출 또는 code-quality 위험 검토 요청이 없는 generated code 또는 vendor code 직접 수정
  - 특정 서적이나 단일 "클린 코드" 교리의 기계적 적용

## 처리하려는 작업 형태

- 기능 구현이나 버그 수정 전에 저장소 맥락과 기존 구현을 확인해야 하는 작업
- 리팩터링 범위와 추상화 도입 여부를 판단해야 하는 작업
- 표준 라이브러리, 프레임워크 API, 기존 의존성, 새 의존성, 직접 구현 중 하나를 선택해야 하는 작업
- 코드 리뷰에서 실제 결함과 유지보수 위험을 우선순위로 제시해야 하는 작업
- 문서 안의 code snippet, command, SQL, configuration 예시가 실제 동작, 보안, 호환성, 운영 위험을 만들 수 있는지 검토해야 하는 작업
- generated/vendor code를 직접 고치기보다 wrapper, source schema, rollback, risk review 같은 대안을 판단해야 하는 작업

## 대표 표면

- 대표 스펙: `code-quality-dev/specs/plugin.md`
- skill 상세 스펙 위치: `code-quality-dev/specs/skills/code-quality.md`
- runtime skill: `code-quality-dev/skills/code-quality/SKILL.md`
- runtime references: `code-quality-dev/skills/code-quality/references/*.md`
- 호출 검증 fixture: `code-quality-dev/skills/code-quality/evals/trigger-prompts.csv`

## 내장 skill 체계

- `code-quality`: production code 구현, 수정, 리팩터링, 테스트, 의존성 판단, 코드 리뷰를 수행한다.
  - spec: `code-quality-dev/specs/skills/code-quality.md`

## SDD 운영 원칙

- plugin spec은 plugin boundary, 대표 표면, skill composition을 소유합니다.
- `code-quality`의 세부 판단 기준은 skill spec과 runtime references가 소유합니다.
- skill spec이 바뀌면 runtime `SKILL.md`와 references를 현재 spec 기준으로 다시 점검합니다.
- release surface에는 `specs/`와 `changes/`를 포함하지 않습니다.

## 현재 구조 메모

- 이 플러그인은 하나의 skill을 가진 단일 목적 번들입니다.
- instruction-only skill이므로 deterministic automation script는 만들지 않습니다.
