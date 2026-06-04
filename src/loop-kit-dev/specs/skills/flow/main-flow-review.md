# flow 메인 플로우 회고 스펙

## 기준 그래프

```text
메인 플로우 그룹 -> 메인 플로우 회고 -> handoff condition
```

## 계약

- 메인 플로우 그룹이 끝나면 handoff condition 전에 항상 회고를 수행합니다.
- 회고 결과는 `000-review.md`에 남깁니다.
- 회고 finding이 없으면 no-finding 결과를 짧게 기록합니다.
- 회고는 다음 active routing을 결정하지 않습니다.
- 회고는 handoff authority가 아닙니다.
- 회고는 raw log, 검증 권한, 커밋 권한, 종료 권한을 소유하지 않습니다.

## 산출

- `000-review.md` 갱신
- finding 또는 no-finding 결과
- handoff condition으로 넘길 residual risk
