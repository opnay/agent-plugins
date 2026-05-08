# turn-gate internal-gates sub-spec

## 목적

이 문서는 `turn-gate` 내부의 gate 전환 계약을 소유합니다.
gate는 외부에 노출되는 별도 skill이나 phase가 아니라, 현재 입력과 flow 상태가 다음 전환 조건을 만족하는지 확인하는 내부 운영 단위입니다.

## 핵심 모델

`turn-gate`는 layer를 사용자 표면으로 노출하지 않습니다.
대신 내부 gate를 통과하며 다음 행동을 결정합니다.

- message intake gate
- flow shaping gate
- task policy gate
- verification gate
- reporting gate
- continuation gate

각 gate는 자기 전환 조건만 소유합니다.
어떤 gate도 사용자의 explicit stop 없이 terminal closure를 추론하지 않습니다.

## Gate 상세 문서

- `gate-message-intake.md`: incoming user message의 explicit stop 여부와 continuation input 분류
- `gate-flow-shaping.md`: message intake 결과를 active flow, 후보, completion criteria로 반영
- `gate-task-policy.md`: 선택된 flow 안에서 command, edit, build, test, handoff 같은 실행 정책 결정
- `gate-verification.md`: work 결과를 pass/fail/blocked/insufficient로 통합하고 non-pass 경로 결정
- `gate-reporting.md`: flow 결과를 다음 진행을 위한 continuity context로 정리
- `gate-continuation.md`: reporting 이후 explicit stop 확인과 next-flow reopening 결정

## Review Questions

- gate 상세 계약이 각 `gate-*.md` 파일로 위임되는가?
- 각 gate가 자기 전환 조건만 소유하고 다른 gate 권한을 침범하지 않는가?
- task 완료가 reporting이나 continuation gate를 건너뛰지 않는가?
- reporting이 terminal closure가 아니라 continuation context로 쓰이는가?
