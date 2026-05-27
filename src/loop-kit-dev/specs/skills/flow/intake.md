# flow intake 계약

## 소유 범위

사용자 입력을 flow 후보로 해석하기 전의 입력 분석, deep interview, 목표 탐지.

## 계약

- intake는 raw user input을 해석하되 source wording과 해석을 섞지 않는다.
- intake는 사용자가 원하는 goal, 거절할 non-goal, scope edge, tradeoff, acceptance signal을 탐지한다.
- intake는 commit, push, PR, publish, release, version bump, destructive action 같은 authority-sensitive 표현을 초기에 표시한다.
- intake는 모호성이나 누락 필드가 있으면 deep interview 질문 주제를 산출한다.
- deep interview는 단계명이 아니라 intake 안의 전략이다.
- intake는 flow를 실행하지 않고, framing이 flow 후보를 설계할 수 있는 입력 계약을 만든다.

## 검토 기준

- goal과 non-goal이 분리돼 있는가?
- authority-sensitive 표현이 실행 승인으로 오해되지 않게 표시됐는가?
- 질문이 필요하다면 어떤 contract field를 잠그기 위한 질문인지 드러나는가?
- raw request와 해석이 섞이지 않았는가?
