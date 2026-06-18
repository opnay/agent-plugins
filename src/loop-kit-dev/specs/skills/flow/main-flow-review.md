# flow 메인 플로우 회고 스펙

## 기준 그래프

```text
메인 플로우 그룹 -> 메인 플로우 회고 -> handoff condition
```

## 계약

- 메인 플로우 그룹이 끝나면 handoff condition 전에 회고를 수행합니다.
- 회고 결과는 `000-review.md`에 남깁니다.
- 회고 finding이 없으면 no-finding 결과를 짧게 기록합니다.
- 이 갱신은 회고 수행 여부를 복구 가능하게 남기기 위한 필수 기록입니다.
- `000-review.md`는 active routing이나 handoff authority로 쓰지 않습니다.

## 산출

- `000-review.md` 갱신
- finding 또는 no-finding 결과
- handoff condition으로 넘길 남은 위험
