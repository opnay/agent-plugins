# Loop Kit Dev 플러그인 스펙

## 목적

`loop-kit-dev`는 `flow`와 `turn-gate`를 함께 제공하는 loop 운영 플러그인입니다.

- `flow`: `메시지 인터뷰 -> 플로우 설계 -> 메인 플로우 -> handoff condition`
- `turn-gate`: active turn 유지, phase prefix, record, verification, next-flow routing

## Flow 계약

- 메시지 인터뷰: 사용자 메시지에서 intent snapshot, alignment risk, high-leverage question, answer pressure test, locked execution brief를 만듭니다.
- 플로우 설계: locked brief에서 active flow, parent flow, sub-flow candidate, phase, handoff를 구분하고 진행할 flow 구성을 만듭니다.
- 메인 플로우: `intake -> framing -> preparation -> work -> verification -> reporting`
- handoff condition: `reporting` 뒤 result, verification, residual risk, next intake condition, commit-readiness 같은 종료 조건을 산출합니다.
- 여러 flow가 필요하면 플로우 설계가 여러 메인 플로우 후보를 만들고, 선택된 flow가 메인 플로우 lifecycle로 들어갑니다.

## Turn Gate 계약

- active turn은 사용자의 explicit stop까지 유지합니다.
- `turn-gate`는 `flow` 판단을 적용하고, flow boundary나 handoff 의미는 `flow` output에 의존합니다.
- 사용자-facing phase 시작 또는 의미 있는 진행 메시지는 canonical prefix를 사용합니다.
  - `[intake]`
  - `[framing]`
  - `[preparation]`
  - `[work]`
  - `[verification]`
  - `[reporting]`
  - `[next-flow]`
- prefix는 artifact, record, command output, question option label의 원래 형식을 보존합니다.
- `reporting`은 다음 사용자 결정의 pre-intake 표면입니다.
- self-drive가 명시되면 prepared sequence의 다음 flow 준비를 자체적으로 진행합니다.

## 포함 범위

- flow boundary, parent flow, finite sub-flow candidate
- operational-preparation flow와 change-unit flow 구분
- message interview와 flow design의 기록 표면
- phase record checkpoint
- risk-based verification method: `clean-context`, `normal`, `not-required`
- mid-flow interruption routing
- self-drive overlay reference
- commit-readiness handoff 판단

## 제외 범위

- broad workflow taxonomy 소유
- domain-specific implementation guidance
- commit/push/PR/release/version bump 실행 권한 자동 부여
- 내부 실행 전략의 직접 사용자 엔트리포인트 확장

## 대표 표면

경로 표기 규칙: `loop-kit-dev/...`는 plugin-root-relative shorthand입니다.
개발 원본의 실제 편집 경로는 `src/loop-kit-dev/...`이고, 루트 `loop-kit/`은 build command 산출 release surface입니다.

- 대표 실행 표면: `turn-gate`
- 대표 flow 표면: `flow`
- 대표 스펙: `loop-kit-dev/specs/plugin.md`
- flow 스펙: `loop-kit-dev/specs/skills/flow/spec.md`
- turn-gate 스펙: `loop-kit-dev/specs/skills/turn-gate/spec.md`
- turn-gate runtime references: `loop-kit-dev/skills/turn-gate/references/*.md`

## 내장 Skill

- `flow`: 사용자 메시지를 해석하고 실제 진행할 flow를 설계합니다. active flow, parent flow, candidate, phase, handoff, readiness, ambiguity, output contract를 소유합니다.
- `turn-gate`: active turn을 유지하고, `flow` 판단을 적용하며, phase prefix, session record, verification routing, next-flow question routing, self-drive gate를 운영합니다.

## SDD 운영

- spec은 현재 계약을 짧게 기록합니다.
- runtime skill은 설치 후 접근 가능한 본문과 references만 실행 지시로 사용합니다.
- skill spec 변경 시 runtime skill을 현재 spec 기준으로 재작성합니다.
- root release surface는 build command 산출물로 갱신합니다.
- clean-context verifier는 read-only 검증에 사용합니다.

## 구조

```text
loop-kit-dev/
  .codex-plugin/plugin.json
  README.md
  specs/plugin.md
  specs/skills/
  skills/
    flow/
    turn-gate/
```
