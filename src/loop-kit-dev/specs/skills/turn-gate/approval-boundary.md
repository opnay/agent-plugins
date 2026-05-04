# turn-gate approval-boundary sub-spec

## 목적

이 문서는 destructive, irreversible, external-action, commit/publish approval boundary를 meaning resolution과 분리해 소유합니다.

## 핵심 계약

- meaning resolution은 사용자의 operation/target 의미를 잠그는 절차이고, approval boundary는 위험 있는 action을 실행해도 되는지 사용자에게 허가받는 절차다.
- 의미가 잠겼더라도 destructive, irreversible, external, publish, push, PR, commit execution approval이 자동으로 주어진 것은 아니다.
- approval-sensitive action은 실행 전에 exact target, expected effect, risk, rollback or recovery 가능성, 포함/제외 scope를 확인한다.
- approval boundary는 user-gated로 유지한다. subagent, inferred intent, prior nearby wording으로 approval을 대신하지 않는다.
- approval이 필요한 상태라면 work로 진행하지 않고 active question-routing으로 target/risk/scope를 제시한다.
- 한 사용자 메시지에 target ambiguity와 approval-sensitive action이 함께 있으면 target/meaning을 먼저 잠그고, 그 다음 exact risk/scope를 제시해 action approval을 받는다.
- 하나의 질문으로 두 결정을 함께 물을 수는 있지만, target lock과 action approval이 서로 다른 결정임이 사용자에게 분명해야 한다.

## Commit / Publish Handoff

- readiness request는 commit approval이 아니다.
- commit approval은 staged/final status, intended diff, unrelated changes exclusion, commit message scope를 확인한 뒤 별도 commit workflow로 넘긴다.
- push, PR, publish 같은 external action은 branch, remote, target, external effect, risk를 확인한 뒤 matching external-action workflow로 넘긴다.
- `turn-gate`는 approval boundary와 handoff 조건을 유지하지만, commit/publish execution detail 자체는 소유하지 않는다.

## 검토 질문

- 의미 해석과 실행 승인을 혼동하지 않았는가?
- 위험 있는 action의 exact target과 risk를 사용자에게 제시했는가?
- readiness request를 commit approval로 취급하지 않았는가?
- external action을 user-gated approval 없이 실행하지 않았는가?
