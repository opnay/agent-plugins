# turn-gate verification sub-spec

## 목적

이 문서는 `work` 이후 clean-context subagent verification과 non-pass handling 계약을 소유합니다.
result reporting과 next-flow reopening 세부 계약은 `records/question-routing.md`가 소유합니다.

## Clean-Context Verification

- `work` 뒤에는 결과 보고 전에 명시적 검증 단계를 둔다.
- 검증은 무조건 clean context 상태의 subagent가 수행해야 한다.
- clean-context verifier는 항상 필요하지만, verifier packet과 실행 check는 flow의 verification expectation, work 중 이미 확보한 evidence, 변경 위험도에 맞게 최소 충분 범위로 구성한다.
- `turn-gate`가 활성화된 동안 읽기 전용 bounded verifier subagent 실행은 clean-context verification 계약의 일부로 사전 허용된 것으로 취급한다.
- 이 사전 허용은 edit permission 없음, scope 확장 없음, destructive/external action 없음, commit/push/PR/publish approval 없음 조건에만 적용된다.
- verifier subagent가 이 조건을 넘어서야 한다면 즉시 중단하고 user-gated question-routing으로 돌아온다.
- clean context는 full conversation history를 그대로 fork하지 않은 독립 검증 context를 뜻한다. 검증 subagent에는 필요한 파일 경로, 사용자 의도, 변경 요약, 검사 명령, pass/fail 기준만 bounded packet으로 전달한다.
- 검증 요청에는 subagent identity 또는 request id, 검증 대상, 기대된 사용자 의도, 변경 파일 또는 산출물, 실행해야 할 command/check, 보고해야 할 pass/fail 기준, edit permission 없음, stop condition을 명시한다.
- work 중 이미 동일한 command/check가 실행됐고 그 결과가 충분히 기록된 경우, verifier packet은 같은 check를 자동 재실행하라고 요구하지 않는다. verifier는 먼저 기록된 evidence를 검토하고, stale output, 불완전한 evidence, 실패 의심, 변경 후 미검증 경로가 있을 때만 재실행을 요청하거나 blocker로 보고한다.
- 파일 변경 없는 질문 답변, 구조 검토, 범위 확인 같은 report-only work의 verifier packet은 불필요한 command 실행보다 source/evidence readback, 논리 반례, 사용자 의도 부합성, 누락된 위험 확인에 초점을 둔다.
- 파일 수정, release surface, 다중 파일 계약, 실패 이력, 사용자 요청 검증처럼 실행 결과가 중요한 경우에는 필요한 command/check를 packet에 포함한다.
- main agent는 같은 context에서 검증을 대신 수행해 통과로 단정하지 않는다.
- main agent의 역할은 검증 요청을 구성하고, subagent 결과를 읽어 residual uncertainty와 이후 flow/phase 재설계 필요 여부를 통합 판정하는 것이다.
- 검증 subagent는 active flow 작업을 구현하거나 scope를 확장하지 않고, 검증과 finding 보고만 수행한다.
- subagent 검증 결과는 `pass`, `fail`, `blocked`, `insufficient` 중 하나로 통합한다.
- 검증 결과가 `fail` 또는 `insufficient`이면 result reporting 전에 active flow를 가장 이른 안전한 phase로 되돌려 보완한다.
- 검증 결과가 `blocked`이거나 검증 도구/subagent를 사용할 수 없으면 통과로 간주하지 않고 user-gated question-routing으로 blocker를 연다.
- non-pass 상태 보고는 blocker report로만 가능하며, successful completion처럼 표현하지 않는다.

## 검토 질문

- work 뒤 검증을 clean-context subagent에게 맡겼는가?
- verifier packet이 flow 위험도와 이미 확보한 evidence에 맞게 최소 충분 범위로 구성됐는가?
- 이미 수행된 command/check를 근거 없이 반복 실행하도록 요구하지 않았는가?
- verifier subagent 요청이 읽기 전용 bounded packet이며 사전 허용 범위를 넘지 않는가?
- 검증 요청에 대상, 의도, 파일/산출물, checks, pass/fail 기준이 들어갔는가?
- non-pass verification을 통과로 취급하지 않고 earliest safe phase 또는 user-gated blocker로 라우팅했는가?
