# pro-planner Skill Spec

## 목적

`pro-planner`는 Codex가 제품, 서비스, 기능을 정의하거나 좁힐 때 기획자 관점의 판단 기준을 적용하게 하는 guidance skill입니다.
핵심 초점은 사용자 문제, 가치, 범위, 요구사항, 우선순위, 성공 기준, tradeoff, handoff 계약을 명확히 하는 것입니다.

## 경계

- 포함:
  - 사용자 문제와 목표 사용자 정의
  - 서비스 가치, 사용 맥락, 핵심 use case 정리
  - MVP, non-goal, release slice, 우선순위 판단
  - 기능 요구사항, 정책 요구사항, 예외 상황, acceptance criteria 정리
  - 큰 사용자 흐름과 의사결정 지점 정의
  - product tradeoff와 designer/engineer handoff 계약
- 제외:
  - 화면 배치, 시각 위계, 색, 톤, 브랜딩, 컴포넌트 표현
  - 코드 구현, 원인 분석, 기술 설계, 테스트 전략
  - 비즈니스 승인, 법률 판단, 가격 정책의 최종 결정
  - 프로젝트 관리 일정표나 조직 운영 절차

## 처리하려는 작업 형태

- 새 기능이나 서비스를 만들기 전 문제와 범위를 정해야 하는 작업
- 요청이 넓거나 모호해 MVP와 non-goal을 나눠야 하는 작업
- 사용자의 요구를 검증 가능한 요구사항과 acceptance criteria로 바꿔야 하는 작업
- UX/UI 또는 구현으로 넘기기 전 빠진 정책, 상태, 예외, 성공 기준을 확인해야 하는 작업
- 여러 선택지 사이에서 가치, 위험, 비용, 학습 효과를 비교해야 하는 작업

## 대표 표면

- 대표 runtime 표면: `judgment-kit-dev/skills/pro-planner/SKILL.md`
- 사용자 스펙 의도: `judgment-kit-dev/specs/skills/pro-planner/intent.md`
- skill spec index: `judgment-kit-dev/specs/skills/pro-planner/spec.md`
- sub-spec directory: `judgment-kit-dev/specs/skills/pro-planner/core/`

## 상세 계약 구조

- `intent.md`: 사용자 스펙 의도 기록
- `core/problem-value.md`: 문제, 사용자, 가치, use case 판단
- `core/scope-requirements.md`: MVP, non-goal, 요구사항, acceptance 기준
- `core/priority-tradeoff-handoff.md`: 우선순위, tradeoff, handoff 계약

## 핵심 처리 계약

- `pro-planner`는 기획 문서 작성 대행이 아니라 product planning judgment를 제공합니다.
- 판단은 사용자 문제, 대상 사용자, 사용 맥락, 제품 목표, 제약, 검증 가능성을 기준으로 합니다.
- 요청이 넓으면 먼저 문제, 사용자, 가치, 성공 기준을 좁힌 뒤 기능 범위를 정합니다.
- 기능 목록보다 사용자가 달성해야 할 결과와 실패하면 안 되는 조건을 먼저 정합니다.
- MVP는 “작은 기능 묶음”이 아니라 핵심 가설을 검증할 수 있는 가장 작은 사용자 가치 단위로 정의합니다.
- non-goal은 나중에 할 일을 버리는 것이 아니라 이번 범위의 판단 기준을 보호하는 경계입니다.
- 요구사항은 기능, 정책, 상태, 예외, 데이터, 권한, 운영 영향을 구분해 적습니다.
- acceptance criteria는 관찰 가능해야 하며, 성공 조건과 실패 조건을 함께 가집니다.
- tradeoff는 가치, 사용성, 구현 비용, 운영 부담, 위험, 학습 효과를 기준으로 설명합니다.
- designer handoff에는 사용자 목표, 주요 흐름, 상태, 우선순위, UX 위험을 넘깁니다.
- engineer handoff에는 요구사항, 데이터/상태, 정책, 예외, acceptance criteria, non-goal을 넘깁니다.

## 독립성 원칙

`pro-planner`는 독립 실행 가능한 runtime skill이어야 합니다.
본문은 sibling skill 이름이나 dev-only spec 경로를 읽으라고 지시하지 않습니다.
다른 `judgment-kit` skill과 함께 쓰일 수는 있지만, 제품과 기능 정의 판단 자체는 이 skill 본문만으로 수행 가능해야 합니다.

## Description Trigger Metadata

이 skill은 passive skill로 선택될 수 있어야 합니다.
frontmatter `description` 끝에는 `#` 없는 쉼표 구분 plain token 목록을 둡니다.
권장 token tail:

`product planning, service planning, feature planning, product requirements, MVP scope, acceptance criteria, product strategy, user problem, prioritization, product tradeoff`

## 검증 기준

- dev runtime skill이 `skills/pro-planner/SKILL.md`에 존재해야 한다.
- release build 후 root `judgment-kit/skills/pro-planner/SKILL.md`가 존재해야 한다.
- plugin spec, README, manifest prompt가 `pro-planner`의 역할과 사용 기준을 언급해야 한다.
- runtime skill 본문은 dev-only `specs/` 경로나 `src/judgment-kit-dev` 경로를 실행 지시로 포함하지 않아야 한다.
- runtime skill은 product planning judgment 중심이며 UI 표현 기준이나 코드 구현 기준으로 좁아지지 않아야 한다.

## 확장 원칙

- 새로운 child spec은 실제 runtime 판단 책임이 커져 index를 읽기 어렵게 만들 때만 추가합니다.
- 사용자 의도는 `intent.md`에만 둡니다.
- child spec은 자기 세부 계약만 소유하고 상위 의도를 반복하지 않습니다.
- runtime skill은 folderized spec 전체를 짧고 실행 가능한 지시로 압축해야 합니다.
