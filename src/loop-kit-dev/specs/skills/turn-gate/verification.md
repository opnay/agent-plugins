# turn-gate verification sub-spec

## 목적

이 문서는 `work` 이후 clean-context subagent verification, result reporting, next-flow reopening 계약을 소유합니다.

## Clean-Context Verification

- `work` 뒤에는 결과 보고 전에 명시적 검증 단계를 둔다.
- 검증은 무조건 clean context 상태의 subagent가 수행해야 한다.
- main agent는 같은 context에서 검증을 대신 수행해 통과로 단정하지 않는다.
- main agent의 역할은 검증 요청을 구성하고, subagent 결과를 읽어 residual uncertainty와 이후 flow/phase 재설계 필요 여부를 통합 판정하는 것이다.
- clean-context subagent에는 검증 대상, 기대된 사용자 의도, 변경 파일 또는 산출물, 실행해야 할 command/check, 보고해야 할 pass/fail 기준을 명시한다.
- 검증 subagent는 active flow 작업을 구현하거나 scope를 확장하지 않고, 검증과 finding 보고만 수행한다.
- subagent 검증 결과가 실패하거나 불충분하면 result reporting 전에 active flow를 가장 이른 안전한 phase로 되돌려 보완하거나, user-gated question-routing으로 blocker를 연다.

## Result Reporting

- 결과 보고 전에는 `Continuity Guard`를 읽거나 재구성하고, 사용자가 명시적으로 종료하지 않았으면 terminal summary가 invalid임을 확인한다.
- result reporting은 terminal response가 아니라 다음 flow decision을 열기 위한 context report다.
- assistant final-answer channel이 사용되더라도 그것만으로 loop termination을 의미하지 않는다.
- explicit stop이 없다면 next-flow question이 여전히 필수다.

## Next-Flow Reopening

- 결과 보고 뒤에는 explicit choice를 주는 active question-routing mode로 다음 플로우를 다시 연다.
- 결과 보고 뒤 visible next-flow choice는 도구 기반 질문이어야 하며, 가능한 경우 `request_user_input`으로 직접 연다.
- loop continuation도 사용자에게 현재 phase, required next action, 열린 선택지를 볼 수 있게 해야 한다.
- 사용자에게 보이는 선택지가 3개 이상이라 턴 종료 선택지를 표시하지 못하는 경우에도, flow record의 `Next Flow Options`에는 별도 turn-end option을 기록한다.
- 사용자가 턴을 종료하자고 요청하지 않으면 clean stop을 기본 경로로 두지 않는다.

## 검토 질문

- work 뒤 검증을 clean-context subagent에게 맡겼는가?
- 검증 요청에 대상, 의도, 파일/산출물, checks, pass/fail 기준이 들어갔는가?
- result reporting 직전에 terminal summary 가능 여부를 확인했는가?
- result reporting 뒤 `request_user_input` 또는 active question-routing으로 다음 flow를 열었는가?
