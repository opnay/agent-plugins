## 사용자 스펙 의도

- toolkit이란 플러그인 만들어서 cli 계열 스킬들 몰아넣어버릴까?
- 그럼 설계해보자.
  - SKILL.md: container cli cheatsheet를 넣고, 하단에 references 폴더 routing 정보
  - references/*.md: 설치, 삭제 가이드, 일부 이론 wiki
- k8s plugin 설명도 넣어두자. kubectl 같이 서드파티 cli 사용해야하는건 제외하거나 보충 설명용으로만 제한하자
- 등록도 해야지 개발 설치 안할거?

---

# Toolkit Dev 플러그인 스펙

## 플러그인 목적

`toolkit-dev`는 Codex가 로컬 개발 CLI를 직접 다룰 때 필요한 도구별 실행 계약을 제공합니다.
설치 여부와 실행 환경을 확인하고, 현재 도구의 명령과 호환성 경계를 적용하며, 실패 시 원인을 좁히고 안전한 다음 행동을 선택하게 합니다.

## 플러그인 경계와 비목표

- 포함:
  - CLI 자체가 주된 작업 대상인 설치, 환경 탐지, 실행, 유지보수, 문제 진단
  - 도구별 명령 문법, 상태 lifecycle, 호환성, 파괴적 작업 경계
  - 설치 후에도 독립적으로 이해 가능한 bounded skill과 조건부 reference
- 제외:
  - CLI를 사용한다는 공통점만 있는 느슨한 skill 집합
  - 일반 shell 명령 치트시트와 운영체제 전반의 사용법
  - 특정 제품 기능 구현, Codex 산출물 제작, 판단 role, subagent orchestration
  - connector나 별도 plugin이 이미 소유하는 앱·서비스 작업

## 처리하려는 작업 형태

- 특정 로컬 개발 CLI로 자원이나 서비스를 확인하고 운용하는 작업
- 도구의 설치·갱신·제거와 데이터 보존 범위를 결정하는 작업
- 기존 CLI workflow를 다른 도구로 옮길 수 있는지 호환성 경계를 판단하는 작업
- 반복되는 명령과 실패 경로를 도구별 skill로 재사용하는 작업

## 대표 표면

- 대표 스펙: `src/toolkit-dev/specs/plugin.md`
- skill 상세 스펙 위치: `src/toolkit-dev/specs/skills/*.md`
- 선택 기준: CLI가 단순 실행 수단이 아니라 작업의 주된 관리 대상인가

## 내장 skill 체계

- `apple-container`: Apple Silicon macOS에서 Apple `container` CLI와 bundled `k8s` plugin을 확인하고 사용하며 설치·제거·개념·Kubernetes plugin reference로 라우팅합니다.
  - spec: `src/toolkit-dev/specs/skills/apple-container.md`

## Plugin Usage 계약

- manifest와 README는 실제 제공되는 CLI skill만 노출합니다.
- 새로운 CLI skill은 도구 자체의 환경, lifecycle, 호환성, 실패 복구 계약이 독립적으로 필요할 때만 추가합니다.
- plugin 공통 선택 기준은 plugin spec, README, manifest가 소유하며 개별 skill에 반복하지 않습니다.
- 개별 skill은 sibling context 없이 독립 실행할 수 있어야 합니다.

## SDD 운영 원칙

- skill 계약을 먼저 spec에 고정하고 runtime skill folder 전체를 현재 spec 기준으로 작성합니다.
- runtime `SKILL.md`는 고빈도 행동과 routing을 소유하고 조건부 세부 절차는 `references/`가 소유합니다.
- source-only spec과 change 기록은 release surface에 포함하지 않습니다.
- skill을 추가하면 README, plugin spec, manifest description·prompt를 함께 점검합니다.

## 현재 구조 메모

- 초기 version은 `0.1.0`입니다.
- 첫 runtime surface는 `apple-container` 하나이며 미래 CLI skill을 미리 약속하지 않습니다.
- marketplace는 기존 순서를 유지한 채 `./toolkit` release surface를 가리키는 `toolkit` 항목을 마지막에 등록합니다.
- local development 확인은 repository marketplace 등록 후 `toolkit@opnay-plugins` 설치와 새 thread pickup으로 검증합니다.
