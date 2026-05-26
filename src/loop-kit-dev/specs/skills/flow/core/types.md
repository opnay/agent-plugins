# flow type 계약

## 소유 범위

`flow`가 사용하는 기본 flow type.

## 계약

- `operational-preparation flow`: 사용자 메시지를 받아 의도, 범위, 비목표, 성공 신호, 검증 기대, 승인 경계를 잠그고 sub-flow 후보 또는 selected flow sequence를 만드는 flow입니다.
- `change-unit flow`: 실제 코드, 문서, fixture, 설정, release surface 같은 검토 가능한 산출물 변경을 소유하는 flow입니다.

운영 flow가 만든 sub-flow 후보는 아직 active execution flow가 아닙니다.
실행 승인이 있거나 `turn-gate`가 다음 flow로 선택하기 전까지 후속 후보로 남깁니다.

## 검토 기준

- 현재 flow가 실행 산출물을 바꾸는가, 아니면 후보와 경계를 설계하는가?
- operational-preparation flow가 만든 후보를 change-unit flow처럼 실행하지 않았는가?
