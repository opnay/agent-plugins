# turn-gate meaning-resolution sub-spec

## 목적

이 문서는 internal mode 선택이나 작업 실행 전에 사용자 지시어의 operation/target ambiguity를 잠그는 계약을 소유합니다.

## 핵심 계약

- analysis 단계에서는 internal mode 선택이나 작업 실행 전에 사용자 메시지의 operation 의미를 먼저 해독한다.
- `merge`, `absorb`, `remove`, `delete`, `split`, `route`, `phase`, `surface`, `skill`, `spec`, `contract` 또는 이에 대응되는 한국어 표현처럼 여러 구조 단위를 가리킬 수 있는 표현은 바로 하나의 작업으로 단정하지 않는다.
- `그`, `그 밑`, `그건`, `그거`, `위`, `아래`, `현재 것`처럼 주변 문맥의 여러 대상을 가리킬 수 있는 지시 표현도 해석에 따라 작업이 달라지면 meaning resolution 대상으로 본다.
- source URL, provenance note, `사용자 스펙 의도` 또는 spec intent block은 대화 맥락처럼 버릴 수 있는 텍스트가 아니라 작업 target이 될 수 있다.
- `출처`, `원본`, `의도`, `그 밑`이 provenance, intent block, normative spec body 중 무엇을 가리키는지에 따라 작업이 달라지면 먼저 target을 잠근다.
- 해석 후보에 따라 파일 범위, 삭제 여부, phase 설계, routing rule, migration 의미, commit scope가 달라지면 active question-routing으로 의미를 먼저 잠근다.

## Question Routing

- meaning resolution 질문도 user-gated이며, 구조적 선택지를 줄 수 있으면 `request_user_input`으로 잠근다.
- meaning resolution 질문은 `deep-interview`가 소유하는 requirement discovery가 아니라, 현재 지시어의 operation 또는 target을 잠그는 current-flow clarification이다.
- 질문은 넓은 freeform 질문이 아니라 "여기서 병합은 skill/spec surface를 합치는 뜻인가, `turn-gate` phase로 흡수하는 뜻인가"처럼 다의어가 가리키는 구조 단위를 직접 잠그는 형태여야 한다.
- meaning resolution이 필요한 경우 flow record의 analysis에는 literal wording, interpreted operation, operation target, alternate interpretations, impact of ambiguity를 남긴다.

## 검토 질문

- 사용자 표현의 operation과 target이 서로 다른 파일/phase/routing 결과를 만들 수 있는가?
- 해석 후보가 commit scope나 destructive action 여부를 바꾸는가?
- 질문이 구조 단위를 직접 잠그는 좁은 선택지인가?
