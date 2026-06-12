# flow 메시지 인터뷰 스펙

## 기준 그래프

```text
entry skill reconfigure
-> deep-interview skill
-> locked execution brief
```

`deep-interview` 내부 압력 테스트가 실패하면 같은 alignment risk로 돌아갑니다.

## 계약

- 메시지 인터뷰는 flow entry skill reconfigure가 끝난 뒤 반드시 `deep-interview` skill을 적용합니다.
- `flow`는 `deep-interview`의 질문, 압력 테스트, locked brief 산출 계약을 재구현하지 않습니다.
- `deep-interview`는 초기 의도 스냅샷, alignment risk, high-leverage 질문, 답변 반영, 예시/반례/비목표/tradeoff 압력 테스트를 수행합니다.
- locked execution brief는 목적, 대상, 범위, 비목표, 완료 기준, 검증 기대, 승인 경계, 근거, 해소된 alignment risk, 남은 모호성을 현재 확정 상태로 남깁니다.
- `deep-interview`가 충분히 잠긴 brief를 만들면 `flow`는 사용자 질문 없이 플로우 설계로 진행할 수 있습니다.

## 산출

- 초기 의도 스냅샷
- alignment risk
- high-leverage 질문 또는 질문 없음 판단
- 답변 반영 결과
- 압력 테스트 결과
- locked execution brief
- `000-plan.md` 갱신 필요 여부
