# flow handoff condition 스펙

## 기준 그래프

```text
메인 플로우 회고 -> handoff condition
```

## 계약

- handoff condition은 메인 플로우 종료 뒤 산출되는 종료 조건입니다.
- handoff condition은 result, verification, residual risk, next intake condition을 포함합니다.
- 다음 flow가 있으면 handoff condition은 다음 intake 조건을 드러냅니다.
- blocker가 남으면 blocked 또는 insufficient 상태를 드러냅니다.
- commit-readiness는 handoff 판단일 수 있지만 commit 실행 권한은 아닙니다.
- push, PR, release, version bump, destructive action은 별도 승인 없이 handoff에서 실행되지 않습니다.

## 산출

- 완료 상태
- 검증 상태
- 남은 위험
- 다음 intake 조건
- blocker 또는 insufficient evidence
- approval-sensitive action 여부
