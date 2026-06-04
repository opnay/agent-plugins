# flow 메시지 인터뷰 스펙

## 기준 그래프

```text
메시지
-> 초기 의도 스냅샷
-> alignment risk 식별
-> high-leverage 질문 하나
-> 답변 반영
-> 예시/반례/비목표/tradeoff 압력 테스트
-> locked execution brief
```

압력 테스트가 실패하면 같은 alignment risk로 돌아갑니다.

## 계약

- 메시지 인터뷰는 사용자 메시지에서 초기 의도 스냅샷을 만듭니다.
- 초기 의도 스냅샷은 원하는 결과, 대상, 범위, 제약을 드러냅니다.
- alignment risk는 locked execution brief를 실행 입력으로 쓰기 어렵게 만드는 가장 큰 불확실성입니다.
- high-leverage 질문은 하나의 alignment risk를 좁히기 위해 한 번에 하나만 사용합니다.
- 답변 반영 뒤 예시, 반례, 비목표, tradeoff로 압력 테스트합니다.
- 압력 테스트가 실패하면 같은 alignment risk를 다시 좁힙니다.
- locked execution brief는 목적, 대상, 범위, 비목표, 완료 기준, 검증 기대, 승인 경계, 근거, 해소된 alignment risk, 남은 모호성을 현재 확정 상태로 남깁니다.
- 메시지 인터뷰가 충분히 잠긴 brief를 만들면 사용자 질문 없이 플로우 설계로 진행합니다.

## 산출

- 초기 의도 스냅샷
- alignment risk
- high-leverage 질문 또는 질문 없음 판단
- 답변 반영 결과
- 압력 테스트 결과
- locked execution brief
- `000-plan.md` 갱신 필요 여부
