# turn-gate mode-selection sub-spec

## 목적

이 문서는 `work`에 들어가기 전 current-phase work의 internal mode를 선택하고 local runtime reference를 읽는 계약을 소유합니다.

## 핵심 계약

- `work`에 들어가기 전 current-phase work의 internal mode를 하나 선택한다.
- `loop-kit-dev`에서는 사용자가 internal mode를 직접 호출하는 대신 `turn-gate`가 이를 선택한다.
- `turn-gate`는 선택된 internal mode에 대응하는 local `references/` 문서를 먼저 읽고 그 계약을 적용한다.
- 아직 어떤 internal mode가 맞는지 확정되지 않았다면 active question-routing mode나 좁은 분석으로 mode selection을 먼저 잠근다.
- commit execution, push, PR 같은 external action은 internal mode가 아니라 user-gated handoff workflow다.

## Mode 선택 기준

- `deep-interview`: requirement discovery, unclear intent, missing scope boundaries, unresolved approval lines가 work 진입을 막는 경우
- `review-loop`: review feedback, QA finding, self-review finding 같은 material issue 처리가 현재 phase의 핵심인 경우
- `ralph-loop`: 하나의 작은 fix-verify-reassess cycle이 현재 phase의 적절한 단위인 경우
- `autopilot`: broad end-to-end delivery가 현재 phase의 핵심인 경우
- `commit-readiness-gate`: 현재 변경 단위가 거의 끝났고 readiness 판단이 핵심인 경우

여러 mode가 겹쳐 보이면 `deep-interview -> review-loop -> ralph-loop -> autopilot -> commit-readiness-gate` 순으로 더 이른 병목을 우선한다.

## Local Reference / SSOT

- local `references/`는 `workflow-kit` upstream spec과 동기화된 absorbed operational contract로 유지한다.
- internal mode contract 변경은 먼저 `workflow-kit` upstream spec에서 정리한다.
- `loop-kit-dev`은 runtime orchestration 관점의 차이와 local absorbed references를 별도로 소유한다.
- upstream contract와 `turn-gate` references의 문구가 어긋나면 같은 변경 단위에서 함께 갱신한다.
- 새로운 internal mode는 기존 mode로 current-phase work를 소유할 수 없을 때만 추가한다.
- internal mode set이나 mandatory tool rule이 바뀌면 `workflow-kit` upstream spec, `loop-kit-dev` plugin spec, `loop-kit-dev-guide`, `turn-gate`, `turn-gate/references/`를 함께 갱신한다.

## Deep-Interview 원본 관계

- 원본 skill source는 `https://github.com/Yeachan-Heo/oh-my-codex/blob/main/skills/deep-interview/SKILL.md`다.
- 원본은 ambiguity gating, OMX tooling, artifact handoff까지 포함한 더 큰 workflow다.
- `loop-kit-dev`의 `references/deep-interview.md`는 full workflow를 그대로 복제한 것이 아니라, `turn-gate`가 requirement-discovery phase에서 필요한 boundary만 흡수한 derived reference다.

## 검토 질문

- current-phase work가 하나의 internal mode로 좁혀졌는가?
- mode selection 전 meaning resolution 또는 approval boundary가 필요한데 건너뛰지 않았는가?
- external action을 internal mode처럼 취급하지 않았는가?
- selected mode의 local reference를 읽고 적용했는가?
