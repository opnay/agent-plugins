# turn-gate verification sub-spec

## 목적

이 문서는 `work` 이후 clean-context subagent verification과 non-pass handling 계약을 소유합니다.
result reporting과 next-flow reopening 세부 계약은 `question-routing.md`가 소유합니다.

## Clean-Context Verification

- `work` 뒤에는 결과 보고 전에 명시적 검증 단계를 둔다.
- 검증은 무조건 clean context 상태의 subagent가 수행해야 한다.
- clean context는 full conversation history를 그대로 fork하지 않은 독립 검증 context를 뜻한다. 검증 subagent에는 필요한 파일 경로, 사용자 의도, 변경 요약, 검사 명령, pass/fail 기준만 bounded packet으로 전달한다.
- 검증 요청에는 subagent identity 또는 request id, 검증 대상, 기대된 사용자 의도, 변경 파일 또는 산출물, 실행해야 할 command/check, 보고해야 할 pass/fail 기준, edit permission 없음, stop condition을 명시한다.
- main agent는 같은 context에서 검증을 대신 수행해 통과로 단정하지 않는다.
- main agent의 역할은 검증 요청을 구성하고, subagent 결과를 읽어 residual uncertainty와 이후 flow/phase 재설계 필요 여부를 통합 판정하는 것이다.
- 검증 subagent는 active flow 작업을 구현하거나 scope를 확장하지 않고, 검증과 finding 보고만 수행한다.
- subagent 검증 결과는 `pass`, `fail`, `blocked`, `insufficient` 중 하나로 통합한다.
- 검증 결과가 `fail` 또는 `insufficient`이면 result reporting 전에 active flow를 가장 이른 안전한 phase로 되돌려 보완한다.
- 검증 결과가 `blocked`이거나 검증 도구/subagent를 사용할 수 없으면 통과로 간주하지 않고 user-gated question-routing으로 blocker를 연다.
- non-pass 상태 보고는 blocker report로만 가능하며, successful completion처럼 표현하지 않는다.

## 검토 질문

- work 뒤 검증을 clean-context subagent에게 맡겼는가?
- 검증 요청에 대상, 의도, 파일/산출물, checks, pass/fail 기준이 들어갔는가?
- non-pass verification을 통과로 취급하지 않고 earliest safe phase 또는 user-gated blocker로 라우팅했는가?
