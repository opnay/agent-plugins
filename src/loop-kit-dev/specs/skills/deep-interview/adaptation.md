# deep-interview 적응 계약

## 참조 원본

- 원본 스킬 이름: `deep-interview`
- 원본 저장소: `Yeachan-Heo/oh-my-codex`
- 원본 경로: `skills/deep-interview/SKILL.md`
- 원본 URL: `https://github.com/Yeachan-Heo/oh-my-codex/blob/main/skills/deep-interview/SKILL.md`

## 적응 원칙

- 원본의 intent-first clarification, one-question-per-round, pressure test, scope/non-goal/tradeoff/decision boundary 잠금은 유지합니다.
- 원본의 CLI, OMX 전용 question command, 상태 파일, artifact bridge는 가져오지 않습니다.
- 질문 도구는 현재 런타임의 `request_user_input` 또는 일반 대화 질문을 사용합니다.
- `deep-interview`는 직접 호출될 수도 있고 plugin-level 사용 표면을 통해 라우팅될 수도 있습니다.
- 요구사항 파악과 방향 잠금이 병목이면 실제 질문을 수행하고, 그 결과를 잠근 뒤 다음 workflow로 handoff합니다.

## 라우팅 성격

- plugin-level 사용 표면은 요청을 보고 현재 병목을 고릅니다.
- `deep-interview`는 workflow 선택 router가 아니라 requirements discovery와 direction evaluation을 실제로 수행하는 skill입니다.
- specialist plugin은 requirements discovery 이후 handoff 대상으로 붙습니다.

## 비목표

- 원본 skill의 OMX runtime, CLI, artifact bridge를 재현하지 않습니다.
- `deep-interview`를 planning, execution, review-loop의 상위 meta-system으로 만들지 않습니다.
- downstream workflow 선택 책임을 skill 내부로 가져오지 않습니다.

## 성공 기준

- `deep-interview`가 필요한 상황에서 advisory answer로 끝나지 않고 실제 질문 라운드를 시작합니다.
- bounded choice 질문은 `request_user_input`을 우선 사용합니다.
- specialist plugin은 requirements lock 이후에 붙습니다.
