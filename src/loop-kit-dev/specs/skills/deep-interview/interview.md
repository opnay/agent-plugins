# deep-interview 질문 계약

## 계약

- 질문은 intent, scope, non-goal, tradeoff, acceptance signal을 잠그는 방향으로 진행합니다.
- 한 번에 하나의 질문을 묻습니다.
- repository fact는 먼저 로컬에서 확인하고, 확인할 수 없는 요구사항 정보만 사용자에게 묻습니다.
- 질문이 1-3개의 짧은 bounded choice로 표현 가능하면 `request_user_input`을 우선 사용합니다.
- bounded choice가 intent discovery를 망치면 일반 대화 질문을 사용합니다.
- 답변은 예시, 반례, 가정 비판, 명시적 비목표, 또는 tradeoff로 압력 테스트합니다.
- 답변이 여전히 모호하면 같은 요구사항 불확실성을 다시 좁힙니다.
- 충분한 clarity가 있으면 execution-ready 또는 direction-ready brief를 산출합니다.

## 압력 테스트

압력 테스트는 사용자의 답변을 그대로 확정하지 않고, 실행 범위와 성공 기준이 흔들리는지 확인하는 질문입니다.
답변이 실제 요구사항을 잠그는지 확인하기 위해 다음 중 하나로 되묻습니다.

- 예시: 답변이 적용되는 구체 상황을 확인합니다.
- 반례: 답변이 적용되면 안 되는 상황을 확인합니다.
- 가정 비판: 답변이 참이려면 성립해야 하는 전제를 확인합니다.
- 명시적 비목표: 이번 작업에서 하지 않을 일을 확인합니다.
- 절충점: 사용자가 받아들일 손실과 거부할 손실을 확인합니다.

압력 테스트는 새 주제로 넓히기 위한 질문이 아닙니다.
같은 요구사항 불확실성을 더 선명하게 만들기 위한 검증 질문입니다.
통과 기준은 답변이 범위, 비목표, 절충점, 수용 기준, 결정 경계 중 최소 하나를 실행에 쓸 수 있게 잠그는 것입니다.
실패 기준은 답변 후에도 같은 요구사항 불확실성이 남아 구현 방향이나 인계 대상이 달라질 수 있는 것입니다.

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
