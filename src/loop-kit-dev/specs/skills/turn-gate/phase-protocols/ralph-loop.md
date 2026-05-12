# turn-gate ralph-loop phase protocol spec

## 목적

`ralph-loop` phase protocol은 하나의 좁은 문제를 작은 fix-verify-reassess cycle로 개선할 때 적용하는 세부 규격입니다.
이 protocol은 mode가 아니며, bounded iteration을 다루는 phase-level 계약입니다.

## 적합 기준

- 문제 하나가 명확하고 작은 수정으로 가설을 검증할 수 있다.
- UI polish, 좁은 refactor 안정화, flaky issue 완화처럼 짧은 반복이 유효하다.
- 다음 반복 여부를 매번 재평가해야 한다.

## 부적합 기준

- broad delivery나 multi-flow planning이 필요하다.
- scope, success criteria, verification expectation이 아직 잠기지 않았다.
- review finding 여러 개를 동시에 triage해야 한다.
- 위험 작업 approval boundary가 풀리지 않았다.

## 핵심 계약

- 한 loop는 하나의 primary issue에 집중한다.
- 현재 가설을 확인할 수 있는 가장 작은 유용한 수정을 선호한다.
- 수정 직후 검증한다.
- 다음 loop가 실제로 필요한지 reassess한 뒤 이어간다.
- success criteria, non-goal, verification, expected risky action, approval boundary가 바뀔 정도로 커지면 preparation 또는 question-routing으로 돌아간다.
- destructive, irreversible, external, commit, push, PR, publish, release, version bump는 정확한 경계가 이미 승인되지 않았다면 실행하지 않는다.

## Handoff

- 한 cycle이 끝나면 verification 결과와 residual risk를 보고하거나, 정당한 경우 다음 bounded cycle을 시작한다.
- 반복이 넓어지면 새 flow 또는 user-gated question-routing으로 분리한다.

## 검토 질문

- 이 loop는 하나의 primary issue에 머물렀는가?
- 수정이 가설을 빠르게 검증할 만큼 작았는가?
- 다음 반복이 필요한 이유가 명확한가?
