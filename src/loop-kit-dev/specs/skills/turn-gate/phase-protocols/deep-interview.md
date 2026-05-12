# turn-gate deep-interview phase protocol spec

## 목적

`deep-interview` phase protocol은 `turn-gate`의 `preparation`에서 요구사항 발견과 scope lock이 실제 병목일 때 적용하는 세부 규격입니다.
이 protocol은 mode가 아니며, implicit default state 안에서 work 진입 전 불확실성을 줄이는 역할을 합니다.

## 적합 기준

- 사용자 의도, 포함 범위, 비목표, 성공 기준, 검증 기대가 부족하다.
- 같은 요청에서 여러 유효한 산출물이 나올 수 있다.
- 결과물이나 검증 경로가 질문 답변에 따라 달라질 수 있다.
- destructive, irreversible, external, commit, push, PR, publish, release, version bump 같은 승인 경계가 아직 불명확하다.
- planned flow list 전체를 만들 만큼의 정보가 부족하다.

## 부적합 기준

- 단순 operation/target ambiguity만 잠그면 되는 경우
- 이미 scope가 충분히 잠긴 구현 작업
- 리뷰 finding 하나를 처리하는 흐름
- 최종 commit readiness 판단 또는 commit execution

## 핵심 계약

- intent, scope, non-goal, acceptance signal, verification expectation, approval boundary를 work 전에 잠근다.
- bounded choices로 잠글 수 있는 질문은 `request_user_input`을 우선한다.
- 질문 없이 추론한 scope라도 flow record에 work boundary와 non-goal을 남긴다.
- planned flow list가 필요하면 각 flow의 boundary, non-goal, verification expectation, expected risky action, user-gated checkpoint를 함께 정리한다.
- 위험 작업 승인은 silent inference로 대체하지 않는다.

## Handoff

- 충분히 잠긴 뒤에는 `phase-protocols/routes.md`로 돌아가 다음 phase protocol을 선택한다.
- scope가 여전히 넓거나 approval boundary가 불명확하면 work로 넘어가지 않고 user-gated question-routing을 유지한다.

## 검토 질문

- 질문 답변에 따라 산출물이나 검증 경로가 달라지는가?
- 전체 planned flow list를 실행할 만큼의 정보가 충분한가?
- 위험 작업과 approval boundary가 flow record에 남았는가?
