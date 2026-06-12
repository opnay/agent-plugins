# deep-interview 스킬 스펙

## 기준

- 기준 문서: `intent.md`
- 이 스펙 트리는 `intent.md`의 사용자 의도를 현재 계약으로 풉니다.

## 목적

`deep-interview`는 질문과 압력 테스트를 통해 사용자의 실제 intent, scope, tradeoff, approval boundary, success criteria를 잠그는 loop-kit 스킬입니다.
요구사항 파악과 방향 잠금이 병목일 때 advisory answer로 끝내지 않고 실제 질문 흐름으로 들어갑니다.

## 소유 범위

- intent-first clarification
- scope, non-goal, tradeoff, decision boundary 잠금
- `request_user_input` 또는 일반 질문을 통한 requirement discovery
- execution-ready 또는 direction-ready brief 산출
- downstream workflow나 specialist plugin으로의 handoff 준비
- 원본 `deep-interview`의 intent-first clarification 철학 적응

## 비소유 범위

- generic workflow ambiguity 정리만 하는 일
- read-only planning
- implementation
- 원본 skill의 OMX runtime, CLI, artifact bridge 재현
- downstream workflow 선택과 실행 책임

## 문서 맵

- `intent.md`: 사용자 스펙 의도
- `interview.md`: 인터뷰 처리 계약
- `adaptation.md`: 원본 skill 적응 기준

## 핵심 계약

- `deep-interview`는 direct-entry workflow skill이면서 routed workflow skill입니다.
- plugin-level 사용 표면은 언제 `deep-interview`를 적용할지 선택할 수 있습니다.
- 질문은 intent, scope, non-goal, tradeoff, acceptance signal을 잠그는 방향으로 진행합니다.
- bounded choice로 잠글 수 있으면 `request_user_input`을 우선 사용합니다.
- advisory answer로 종료하지 않고 필요한 질문 라운드를 실제로 수행합니다.
- 충분한 clarity를 얻으면 execution-ready 또는 direction-ready brief를 만들고 다음 workflow로 handoff합니다.
- brief는 숨은 sibling context 없이 downstream workflow가 바로 시작할 수 있어야 합니다.

## 검토 질문

- 지금 병목이 단순 ambiguity가 아니라 실제 requirement discovery인가?
- bounded choice로 잠글 수 있는 질문인데 자유서술형으로 흐르고 있지 않은가?
- discovery 이후 handoff 대상이 충분히 선명한가?
- 원본 skill의 runtime 전제가 현재 skill에 새지 않았는가?
