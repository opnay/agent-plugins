# Loop Kit Dev 플러그인 스펙

## 목적

`loop-kit-dev`는 `flow`와 `turn-gate`를 함께 제공하는 loop 운영 플러그인입니다.

- `deep-interview`: intent, scope, tradeoff, approval boundary를 질문과 압력 테스트로 잠그는 인터뷰 스킬
- `flow`: `entry skill reconfigure -> deep-interview -> 플로우 설계 -> 메인 플로우 -> 메인 플로우 회고 -> handoff condition`
- `turn-gate`: active turn 유지, flow wrapper, handoff question routing, self-drive gate, explicit stop

내장 skill 의존 방향은 `deep-interview -> flow -> turn-gate`입니다.
상위 skill은 하위 skill을 적용할 수 있지만, 하위 skill의 runtime/spec은 상위 skill 맥락을 전제로 하지 않습니다.

## Flow 계약

- skill reconfigure: flow entry와 post-reporting continuation boundary에서 필요한 skill 본문을 source에서 다시 읽어 루프 중 잊힌 skill context를 복구합니다.
- deep-interview: 메시지 인터뷰에서 locked execution brief를 만듭니다.
- 모든 사용자 메시지는 같은 flow 경로를 탑니다. 메시지 인터뷰가 충분히 잠긴 brief를 받으면 사용자 질문 없이 플로우 설계로 진행할 수 있습니다.
- 플로우 설계: locked brief에서 active flow, parent flow, sub-flow candidate, phase, handoff를 구분하고 진행할 flow 구성을 만듭니다.
- 메인 플로우: `intake -> framing -> preparation -> work -> verification -> reporting`
- 다음 flow가 있으면 `reporting` 직후 skill reconfigure를 수행하고 다음 `intake`로 라우팅합니다.
- 사용자-facing 진행 메시지: 현재 phase label을 산출하되, 기록이나 산출물 본문에는 label을 전파하지 않습니다.
- 메인 플로우 회고: 항상 `000-review.md`를 갱신하고, finding이 없으면 no-finding 결과로 짧게 남깁니다.
- handoff condition: 메인 플로우와 필요한 회고 뒤 result, verification, residual risk, next intake condition, commit-readiness 같은 종료 조건을 산출합니다.
- 여러 flow가 필요하면 플로우 설계가 여러 메인 플로우 후보를 만들고, 선택된 flow가 메인 플로우 lifecycle로 들어갑니다.

## Turn Gate 계약

- active turn은 사용자의 explicit stop까지 유지합니다.
- `turn-gate`는 `flow` 판단을 적용하고, flow boundary나 handoff 의미는 `flow` output에 의존합니다.
- `flow skill: handoff` 뒤에는 `질문 도구: 다음 플로우 선택`으로 다음 flow 입력을 고릅니다.
- `next-flow gate`에서 예약 플로우를 재검토하고 다음 flow를 선택한 뒤 `000-plan.md`를 매번 업데이트합니다. 필요한 skill reread는 flow entry 또는 post-reporting continuation boundary의 skill reconfigure가 수행합니다.
- 질문 도구는 `deep-interview`로 입력을 구체화한 뒤 다시 `flow`로 들어갑니다.
- 사용자-facing 진행 메시지는 source skill이 소유한 phase prefix로 현재 단계를 드러냅니다. `turn-gate`는 `flow` phase label을 재정의하지 않고 적용하며, `next-flow gate`에서는 `[next-flow]`를 소유합니다.
- phase prefix는 artifact, record, command summary, question option label에 전파하지 않습니다.
- self-drive가 명시되면 그래프 노드가 아니라 준비된 sequence gate가 질문 도구를 대체합니다.
- record, verification, interruption, date 처리는 메인 그래프 노드가 아니라 active turn을 복구하고 안전하게 라우팅하기 위한 지원 계약입니다.

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
- deep-interview 스펙: `loop-kit-dev/specs/skills/deep-interview/spec.md`
- turn-gate 스펙: `loop-kit-dev/specs/skills/turn-gate/spec.md`
- turn-gate runtime references: `loop-kit-dev/skills/turn-gate/references/*.md`

## 내장 Skill

- `deep-interview`: flow 메시지 인터뷰와 직접 호출 인터뷰에서 intent, scope, tradeoff, approval boundary를 잠그고 locked execution brief를 산출합니다.
- `flow`: flow entry와 post-reporting continuation boundary의 skill reconfigure, `deep-interview` 필수 적용, 실제 진행할 flow 설계를 소유합니다. active flow, parent flow, candidate, phase, handoff, readiness, ambiguity, output contract를 소유합니다.
- `turn-gate`: active turn을 유지하고, `flow` 판단을 적용하며, handoff question routing, self-drive gate, explicit stop, 지원 라우팅 계약을 운영합니다.

## SDD 운영

- spec은 계약을 짧게 기록합니다.
- runtime skill은 설치 후 접근 가능한 본문과 references만 실행 지시로 사용합니다.
- skill spec 변경 시 runtime skill을 해당 skill spec 기준으로 재작성합니다.
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
    deep-interview/
    flow/
    turn-gate/
```
