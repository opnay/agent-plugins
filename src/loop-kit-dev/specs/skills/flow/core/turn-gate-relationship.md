# flow and turn-gate relationship

## 목적

이 문서는 `flow`와 `turn-gate`의 소유권 경계를 소유합니다.

## 계약

- `flow`는 flow 단위의 정의, 분해, 완료 조건을 소유합니다.
- `turn-gate`는 active turn 안에서 flow 없이 진행하지 못하게 하고, flow reporting 뒤 next-flow 질문을 엽니다.
- `turn-gate`가 session record에 flow를 기록할 때는 `flow`의 output contract를 사용합니다.
- `flow`는 다음 flow를 자동 실행하지 않습니다.
- self-drive는 준비된 flow sequence 위에서만 자동 진행할 수 있으며, sub-flow 후보 생성 자체로 self-drive 실행 권한이 생기지 않습니다.

## 검토 질문

- flow 완료와 turn 종료를 혼동하지 않았는가?
- sub-flow 후보를 `turn-gate` 질문 없이 실행하지 않았는가?
- flow output contract가 session record에 반영될 수 있는 형태인가?
