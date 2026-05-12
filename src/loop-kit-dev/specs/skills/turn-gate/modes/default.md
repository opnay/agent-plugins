# turn-gate implicit default state spec

## 목적

`default`는 runtime에서 별도로 선택하거나 노출하는 mode label이 아니라, `turn-gate` skill이 활성화되면 기본으로 동작하는 implicit operating state입니다.
모든 일반 turn-gated flow는 이 기본 상태에서 진행합니다.

## 계약

- implicit default state는 `preparation -> work -> verification -> reporting -> next-flow` core loop를 그대로 수행한다.
- 사용자 메시지 기반 preparation, 기존 flow 기반 preparation, scope lock, meaning resolution, 일반 파일 수정, 조사, 검증, 보고, next-flow reopening을 포함한다.
- `deep-interview`, `review-loop`, `ralph-loop`, `autopilot`, `commit-readiness-gate` 같은 phase protocol은 기본 상태에서 phase를 수행하는 세부 규격으로 적용한다.
- phase protocol은 mode가 아니며, flow type, operating state, phase protocol 기록을 대체하지 않는다.
- external action, destructive action, commit, push, PR, publish, release, version bump는 기본 상태에서도 approval-sensitive execution으로 다루며, 필요한 approval boundary가 없으면 user-gated handoff로 남긴다.

## 검토 질문

- 일반 turn-gated flow를 기본 상태로 유지했는가?
- 현재 phase에서 필요한 protocol을 mode와 분리해 설명했는가?
- 결과 보고 뒤 source-recorded explicit stop 없이 terminal close하지 않고 next-flow로 이어졌는가?
