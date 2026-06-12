# deep-interview 질문 계약

## 계약

- 질문은 intent, scope, non-goal, tradeoff, acceptance signal을 잠그는 방향으로 진행합니다.
- 한 번에 하나의 질문을 묻습니다.
- repository fact는 먼저 로컬에서 확인하고, 확인할 수 없는 요구사항 정보만 사용자에게 묻습니다.
- 질문이 1-3개의 짧은 bounded choice로 표현 가능하면 `request_user_input`을 우선 사용합니다.
- bounded choice가 intent discovery를 망치면 일반 대화 질문을 사용합니다.
- 답변은 예시, 반례, 명시적 비목표, 또는 tradeoff로 압력 테스트합니다.
- 답변이 여전히 모호하면 같은 요구사항 불확실성을 다시 좁힙니다.
- 충분한 clarity가 있으면 execution-ready 또는 direction-ready brief를 산출합니다.

## Brief

brief는 다음을 드러냅니다.

- intent
- scope
- non-goal
- tradeoff
- acceptance signal
- decision boundary
- what may be decided autonomously
- what still needs confirmation
- recommended handoff

## Handoff

- brief를 downstream workflow나 specialist plugin handoff 입력으로 넘깁니다.
- handoff 대상의 내부 lifecycle, 설계, 라우팅 책임은 소유하지 않습니다.
- 사용자가 명시적으로 남은 모호성을 안고 진행하겠다고 선택하면 residual risk를 남깁니다.

## 검토 기준

- advisory answer로 종료하지 않고 필요한 질문 라운드를 실제로 수행했는가?
- 질문 수를 늘리는 대신 가장 큰 요구사항 불확실성을 줄였는가?
- handoff 대상이 충분히 선명한가?
