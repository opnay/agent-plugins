# turn-gate verification sub-spec

## 목적

이 문서는 `work` 이후 risk-based verification method와 non-pass handling 계약을 소유합니다.
result reporting과 next-flow reopening 세부 계약은 `records/question-routing.md`가 소유합니다.

## Risk-Based Verification Method

- `work` 뒤에는 결과 보고 전에 명시적 검증 판단을 둔다.
- verification method는 flow의 task tier, verification expectation, work 중 이미 확보한 evidence, 변경 위험도에 맞게 선택한다.
- verification method는 `clean-context`, `normal`, `not-required` 중 하나로 기록한다.
- verification result status는 `pass`, `fail`, `blocked`, `insufficient` 중 하나로 통합한다. method 이름을 result status처럼 쓰지 않는다.
- `clean-context`는 bounded read-only verifier subagent가 독립 context에서 검증하는 방법이다.
- `normal`은 main agent가 같은 context에서 command/check, source readback, evidence checklist, 논리 반례 검토를 수행하고 그 근거를 기록하는 방법이다.
- `not-required`는 별도 검증 동작이 필요하지 않은 경우에만 쓰는 method 판단이다. 이 경우에도 omission reason, already-known evidence 또는 no-output rationale, residual uncertainty를 기록한다.

## Method Selection

- 다음 경우에는 `clean-context`를 기본값으로 둔다.
  - 파일 수정이 있다.
  - release surface, manifest, template, scenario fixture, build output을 바꾼다.
  - 여러 파일 사이 계약이나 plugin usage surface가 바뀐다.
  - 이전 command/check 실패 이력이 있다.
  - 사용자가 검증, QA, review, commit-readiness 판단을 요청했다.
  - destructive, irreversible, external, commit, push, PR, publish, release, version bump 같은 approval-sensitive action을 준비하거나 실행했다.
- 다음 경우에는 `normal`을 사용할 수 있다.
  - no-edit research 또는 read-only inspection이다.
  - 단일 설명, 단일 상태 확인, 범위 확인처럼 command/source readback evidence로 충분하다.
  - work 중 이미 충분한 command/check evidence를 확보했고, 변경 후 미검증 경로가 없다.
  - clean-context를 쓰지 않는 이유와 남은 불확실성을 reporting 또는 flow record에 남길 수 있다.
- 다음 경우에는 `not-required`를 사용할 수 있다.
  - flow가 work 실행 전에 blocker나 user-gated question으로 멈췄다.
  - activation-only, next-flow selection, 또는 scope-routing처럼 검증할 work output이 없다.
  - 사용자가 요청한 결과가 별도 검증보다 질문 routing 자체이며, 검증 생략 이유를 기록할 수 있다.
- `not-required`는 파일 변경, release surface, 다중 파일 계약, 실패 이력, 사용자 요청 검증, approval-sensitive action에는 사용할 수 없다.

## Clean-Context Verifier Contract

- `turn-gate`가 활성화된 동안 읽기 전용 bounded verifier subagent 실행은 clean-context verification 계약의 일부로 사전 허용된 것으로 취급한다.
- 이 사전 허용은 edit permission 없음, scope 확장 없음, destructive/external action 없음, commit/push/PR/publish approval 없음 조건에만 적용된다.
- verifier subagent가 이 조건을 넘어서야 한다면 즉시 중단하고 user-gated question-routing으로 돌아온다.
- clean context는 full conversation history를 그대로 fork하지 않은 독립 검증 context를 뜻한다. 검증 subagent에는 필요한 파일 경로, 사용자 의도, 변경 요약, 검사 명령, pass/fail 기준만 bounded packet으로 전달한다.
- 검증 요청에는 subagent identity 또는 request id, 검증 대상, 기대된 사용자 의도, 변경 파일 또는 산출물, 실행해야 할 command/check, 보고해야 할 pass/fail 기준, edit permission 없음, stop condition을 명시한다.
- work 중 이미 동일한 command/check가 실행됐고 그 결과가 충분히 기록된 경우, verifier packet은 같은 check를 자동 재실행하라고 요구하지 않는다. verifier는 먼저 기록된 evidence를 검토하고, stale output, 불완전한 evidence, 실패 의심, 변경 후 미검증 경로가 있을 때만 재실행을 요청하거나 blocker로 보고한다.
- 검증 subagent는 active flow 작업을 구현하거나 scope를 확장하지 않고, 검증과 finding 보고만 수행한다.

### Documentation-Only Research Artifact Packet

`.agents/researches/**/topic.md` 같은 documentation-only research artifact도 파일이 바뀌었으면 `clean-context` 기본값을 유지한다.
다만 verifier packet은 전체 source 재조사가 아니라 변경된 topic/session 파일과 이미 기록된 evidence gap 확인으로 좁힌다.

최소 packet은 다음을 포함한다.

- 변경된 research topic 파일과 active session record.
- research documentation 목적이며 source/spec/runtime 구현 변경이 아니라는 사용자 의도.
- main thread가 이미 확보한 source readback 또는 핵심 finding evidence.
- 필수 heading 존재 여부.
- conclusion이 기록된 evidence와 모순되지 않는지.
- source/spec/runtime 구현을 완료했다는 허위 claim이 없는지.
- session record의 closure state, verification status, required next action이 일관적인지.

같은 broad source search를 자동 재실행하지 않는다.
단, evidence가 stale, incomplete, suspicious 상태이거나 변경 path를 빠뜨렸으면 verifier는 재실행 요청 또는 blocker를 보고한다.

## Result Handling

- 검증 결과가 `fail` 또는 `insufficient`이면 result reporting 전에 active flow를 가장 이른 안전한 phase로 되돌려 보완한다.
- 검증 결과가 `blocked`이거나 필요한 검증 도구/subagent를 사용할 수 없으면 통과로 간주하지 않고 user-gated question-routing으로 blocker를 연다.
- non-pass 상태 보고는 blocker report로만 가능하며, successful completion처럼 표현하지 않는다.
- `normal` 또는 `not-required` method를 사용했더라도 evidence, omission reason, residual uncertainty가 불충분하면 `pass`로 통합하지 않는다.

## Non-Pass Routing Decision Table

| Verification status | 성공 보고 | 기대 라우팅 | User-gated 여부 |
| --- | --- | --- | --- |
| `pass` | 허용 | reporting 이후 next-flow reopening | 일반 next-flow 또는 self-drive continuation |
| `fail` | 금지 | 가장 이른 안전한 `work` 또는 repair phase로 복귀 | repair가 scope/approval 경계를 넘을 때만 |
| `insufficient` | 금지 | evidence 보강이 가능한 `verification` 또는 `work` phase로 복귀 | evidence 보강이 승인 경계 밖일 때만 |
| `blocked` | 금지 | user-gated blocker question-routing | 필요 |

`fail`, `insufficient`, `blocked`는 모두 successful completion, terminal summary, next-flow selection authority가 아니다.
non-pass 상태를 사용자에게 알릴 수는 있지만, 그 보고는 completion report가 아니라 repair 또는 blocker routing이어야 한다.

## 검토 질문

- verification method를 `clean-context`, `normal`, `not-required` 중 하나로 기록했는가?
- method와 result status를 섞지 않았는가?
- `clean-context`가 필요한 작업을 `normal` 또는 `not-required`로 낮추지 않았는가?
- verifier packet 또는 normal evidence가 flow 위험도와 이미 확보한 evidence에 맞게 최소 충분 범위로 구성됐는가?
- 이미 수행된 command/check를 근거 없이 반복 실행하도록 요구하지 않았는가?
- verifier subagent 요청이 읽기 전용 bounded packet이며 사전 허용 범위를 넘지 않는가?
- 검증 요청에 대상, 의도, 파일/산출물, checks, pass/fail 기준이 들어갔는가?
- `not-required`를 사용했다면 검증할 work output이 없다는 이유와 residual uncertainty를 기록했는가?
- non-pass verification을 통과로 취급하지 않고 earliest safe phase 또는 user-gated blocker로 라우팅했는가?
