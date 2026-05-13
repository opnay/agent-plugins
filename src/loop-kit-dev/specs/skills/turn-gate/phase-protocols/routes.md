# turn-gate phase protocol routes spec

## 목적

이 문서는 `turn-gate`의 implicit default operating state와 phase protocol routing을 소유합니다.
상세 protocol 계약은 같은 폴더의 개별 protocol spec이 소유합니다.

## 핵심 계약

- `turn-gate`의 기본 동작은 별도 mode label을 선택하지 않는 implicit default operating state다.
- 이 기본 상태는 일반적인 `preparation -> work -> verification -> reporting -> next-flow` 흐름을 소유한다.
- `deep-interview`, `review-loop`, `ralph-loop`, `autopilot`, `commit-readiness-gate`는 mode가 아니라 phase protocol이다.
- phase protocol은 현재 operating state 안에서 phase를 어떻게 수행할지 정하는 보조 계약이며, operating state 선택 자체를 대체하지 않는다.
- commit execution, push, PR 같은 external action은 approval-sensitive execution step 또는 handoff workflow로 다룬다.

## Operating State 선택 기준

implicit default operating state는 모든 일반 turn-gated work에 적용한다.
단일 flow 실행, 질문 기반 scope lock, 일반 파일 수정, 조사, 검증, 보고, next-flow reopening을 포함한다.

## Protocol 선택 기준

- `deep-interview`: requirement discovery와 scope lock이 work 진입을 막는 경우. 상세 계약은 `deep-interview.md`가 소유한다.
- `review-loop`: review, QA, self-review finding이 현재 병목인 경우. 상세 계약은 `review-loop.md`가 소유한다.
- `ralph-loop`: 하나의 작은 fix-verify-reassess cycle이 적절한 경우. 상세 계약은 `ralph-loop.md`가 소유한다.
- `autopilot`: locked scope 안에서 broad end-to-end execution이 필요한 경우. 상세 계약은 `autopilot.md`가 소유한다.
- `commit-readiness-gate`: 현재 change unit의 commit readiness 판단이 필요한 경우. 상세 계약은 `commit-readiness-gate.md`가 소유한다.

여러 phase protocol이 겹쳐 보이면 earliest blocker를 기준으로 고른다.
일반 우선순위는 `deep-interview -> review-loop -> ralph-loop -> autopilot -> commit-readiness-gate`다.

## Local References

- local `references/`는 runtime에서 읽을 수 있는 absorbed operational contract로 유지한다.
- runtime reference 파일명은 mode 이름과 일대일 대응할 필요가 없다.
- phase protocol의 일반 의미가 바뀌면 관련 workflow skill spec 또는 해당 phase protocol spec에서 정리한다.
- 관련 workflow skill spec과 `turn-gate` references의 문구가 어긋나면 같은 변경 단위에서 함께 갱신한다.
- 새로운 phase protocol은 implicit default state의 phase routing으로 표현할 수 없을 때만 추가한다.
- operating state set이나 mandatory tool rule이 바뀌면 `loop-kit-dev` plugin spec, manifest prompt, `turn-gate`, `turn-gate/references/`를 함께 점검한다.

## Provenance Note

- `deep-interview` 원본 skill source는 `https://github.com/Yeachan-Heo/oh-my-codex/blob/main/skills/deep-interview/SKILL.md`다.
- 원본은 ambiguity gating, OMX tooling, artifact handoff까지 포함한 더 큰 workflow다.
- `loop-kit-dev`의 runtime reference는 full workflow를 그대로 복제한 것이 아니라, `turn-gate`가 phase protocol 선택과 적용에 필요한 boundary만 흡수한 derived reference다.

## 검토 질문

- current flow가 implicit default state 안에서 어떤 phase protocol이 필요한지 좁혀졌는가?
- deep-interview/review-loop 같은 phase protocol을 mode로 기록하지 않았는가?
- protocol 세부 계약을 `routes.md`에 중복하지 않고 개별 protocol spec으로 보냈는가?
- mode 또는 protocol 선택 전 meaning resolution 또는 approval boundary가 필요한데 건너뛰지 않았는가?
- external action을 approval-sensitive execution step 또는 user-gated handoff로 다뤘는가?
- implicit default state와 필요한 phase protocol의 local reference를 읽고 적용했는가?
