# flow 메시지 인터뷰 스펙

## 기준 그래프

```text
메시지 -> 초기 의도 스냅샷 -> alignment risk 식별 -> high-leverage 질문 하나 -> 답변 반영 -> 압력 테스트 -> locked execution brief
```

압력 테스트가 실패하면 같은 alignment risk로 돌아갑니다.

## 계약

- 모든 사용자 메시지는 메시지 인터뷰를 거칩니다.
- 메시지 인터뷰는 deep-interview 역할을 `flow` 내부 해석 단계로 흡수합니다.
- 초기 의도 스냅샷은 사용자가 원하는 결과, 대상, 범위, 제약, 승인 경계를 임시로 잡습니다.
- alignment risk는 실행 계약을 틀리게 만들 가능성이 가장 큰 모호성입니다.
- high-leverage 질문은 한 번에 하나만 둡니다.
- 질문은 답변에 따라 flow 계약이 달라질 때만 사용자에게 냅니다.
- 질문이 필요 없으면 사용자 질문 없이 locked execution brief로 진행합니다.
- 답변은 예시, 반례, 비목표, tradeoff로 압력 테스트합니다.
- 답변이 여전히 모호하면 같은 alignment risk를 다시 좁힙니다.

## 산출

- 초기 의도 스냅샷
- 식별된 alignment risk
- 질문 또는 질문 없음 판단
- 답변 반영 결과
- 압력 테스트 결과
- locked execution brief
- `000-plan.md` 갱신 필요 여부
