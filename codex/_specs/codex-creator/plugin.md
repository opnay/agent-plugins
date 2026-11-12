## 사용자 스펙 의도

- plugin과 skill 제작 책임은 `codex-creator` 플러그인이 소유합니다.

---

# Codex Creator 플러그인 스펙

## 플러그인 목적

`codex-creator`는 재사용 가능한 Codex plugin과 skill을 설계·작성·검토하는 확장 가이드를 제공합니다. plugin 경계와 개별 skill 계약을 같은 제작 번들 안에서 다룹니다.

## 플러그인 경계와 비목표

- 포함: plugin 경계·packaging·사용 안내, skill 경계·선택 metadata·비교 행동 검증.
- 제외: canonical system creator 대체·설치, 일반 제품 구현, 문체 규칙 소유, 무관한 도구 운영.
- 두 skill은 각각 canonical system 동명 skill을 확장합니다. 이 플러그인은 기본 skill을 제공하지 않습니다.

## 처리하려는 작업 형태

- 설치 가능한 plugin bundle을 만들거나 재구성합니다.
- 재사용 skill을 만들거나 경계·description·검증 계약을 검토합니다.

## 대표 표면

- 실행 플러그인: `codex/codex-creator/`
- 대표 스펙: `codex/_specs/codex-creator/plugin.md`
- 개별 계약: `skills/plugin-creator/spec.md`, `skills/skill-creator/spec.md`

## 내장 skill 체계와 선택 기준

- `$codex-creator:plugin-creator`: plugin 목적·경계, 내장 skill 구성, manifest·README의 사용 안내를 다룹니다.
  - [skill spec](skills/plugin-creator/spec.md)
- `$codex-creator:skill-creator`: 개별 skill 책임·비목표, description 선택 조건, 사용/미사용 비교 검증을 다룹니다.
  - [skill spec](skills/skill-creator/spec.md)
- plugin과 skill을 함께 만들 때 plugin 경계를 먼저 정하고 개별 skill을 작성합니다. 두 확장은 숨은 sibling context 없이 자기 계약을 설명합니다.

## SDD 운영 원칙

- plugin spec은 번들 경계와 선택 기준을 소유하며, 세부 처리 계약은 각 skill spec이 소유합니다.
- 책임·선택 기준 변경 시 README·manifest·관련 spec을 함께 확인합니다.
- runtime은 설치 후 사용할 수 있는 본문과 리소스로 실행 계약을 설명합니다.

## 확장 원칙

- Codex plugin·skill 제작에 직접 필요한 표면만 포함합니다.
- 문체·일반 실행·도메인 기능을 제작 번들에 흡수하지 않습니다.
