## 사용자 스펙 의도

- subagent를 준비하는 과정에서 종료 시점 계획을 우선하고 싶다.
- 메인 에이전트의 맥락을 넘길 때는 서브 에이전트의 시점에서 필요한 항목만 정리해 전달해야 한다.

---

# subagent-gate 스킬 스펙

## 목적

`subagent-gate`는 subagent를 실제로 준비하거나 호출하기 전에 위임의 종료 시점, 전달 맥락, 책임 범위, 결과 수용 기준을 잠그는 스킬입니다.
핵심은 subagent를 "일단 보내는" 것이 아니라, 언제 돌아와야 하고 무엇을 가져와야 하며 어떤 맥락만 필요로 하는지를 먼저 결정하는 것입니다.

## 경계

- 포함:
  - subagent 호출 전 종료 시점과 회수 조건 정의
  - subagent 관점의 최소 context packet 구성
  - 위임 task, ownership, output contract, pass/fail criteria 정리
  - user-gated decision, destructive action, external action 같은 delegation 금지 경계 확인
  - 메인 에이전트가 병렬로 계속할 수 있는 non-overlapping work와 blocked condition 분리
- 제외:
  - `.codex/agents/*.toml` custom agent 정의
  - plugin packaging 또는 skill authoring 자체
  - empirical prompt evaluation의 반복 실행 설계
  - subagent에게 넘긴 뒤의 구현 상세
  - 사용자 승인이 필요한 결정을 subagent가 대신 판단하게 하는 것

## 처리하려는 작업 형태

- subagent를 spawn하기 전에 prompt packet, 종료 조건, 검증 기준을 정리해야 하는 경우
- main agent의 전체 맥락이 너무 넓어 subagent에게 필요한 정보만 다시 구성해야 하는 경우
- 위임할 작업이 critical path를 막는지, 병렬 sidecar인지, 검증 전용인지 판단해야 하는 경우
- user-gated boundary나 destructive/external action risk가 subagent prompt에 섞일 수 있는 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `advance-codex-dev/skills/subagent-gate/SKILL.md`
- 관련 상위 라우팅: `advance-codex-dev-guide`
- 인접 skill:
  - `subagent-creator`: custom agent definition을 소유한다.
  - `empirical-prompt-tuning`: fixed scenario 기반 reusable instruction evaluation을 소유한다.

## 핵심 처리 계약

- 먼저 종료 시점 계획을 작성한다.
- subagent가 무엇을 끝내면 돌아와야 하는지, 무엇을 하지 말아야 하는지, 어떤 결과 형식이 필요한지 명시한다.
- 전달 맥락은 main agent가 가진 전체 대화나 추론이 아니라 subagent 시점에서 작업 수행에 필요한 항목만 포함한다.
- context packet은 task, scope, known facts, relevant files, constraints, output contract, verification criteria, forbidden actions를 구분한다.
- user-gated, destructive, irreversible, external-action, safety decision은 subagent에게 승인 또는 최종 판단으로 위임하지 않는다.
- subagent가 병렬로 수행해도 되는 작업인지, main agent가 기다려야 하는 blocker인지 분리한다.
- handoff packet에는 main-thread blocked state를 명시 필드로 둔다.
- 요청이 `subagent-gate` 경계 밖이면 handoff packet을 억지로 만들지 않고 no-handoff routing note를 반환한다.
- no-handoff routing note는 reason, right skill or workflow, main-thread next action, main-thread blocked state, re-entry condition을 포함한다.
- no-handoff 상태에서는 subagent ownership, edit permission, pass/fail criteria를 억지로 만들지 않는다.

## 종료 시점 계획

- subagent를 준비하기 전에 다음을 먼저 정한다.
  - return trigger: 어떤 결과나 실패 조건에서 돌아와야 하는가
  - stop boundary: 어떤 상황에서 더 파고들지 말고 중단해야 하는가
  - answer contract: 최종 답변에 반드시 포함해야 하는 항목은 무엇인가
  - no-change boundary: 검증 전용 subagent라면 파일을 수정하지 않아야 하는가
  - close condition: 결과를 받은 뒤 subagent를 재사용할지 닫을지
  - main-thread blocked state: 메인 에이전트가 결과를 기다려야 하는지, non-overlapping work를 계속할 수 있는지

## Context Packet 규칙

- subagent에게 넘기는 맥락은 subagent 관점에서 재작성한다.
- 포함할 수 있는 항목:
  - user-visible goal
  - assigned files, directories, or responsibility boundary
  - relevant facts already verified by commands or file reads
  - current constraints, including repository rules and approval boundaries
  - expected output shape
  - pass/fail criteria or review focus
  - known non-goals and actions to avoid
- 제외할 항목:
  - main agent의 불필요한 내부 추론
  - subagent task와 무관한 conversation history
  - 아직 확인하지 않은 추측을 사실처럼 적은 내용
  - 사용자가 직접 승인해야 하는 선택을 subagent가 결정하게 만드는 문구

## No-Handoff Routing Note 규칙

- 다음 상황에서는 handoff packet 대신 no-handoff routing note를 반환한다.
  - task boundary, files, output shape, return condition이 없어 delegation-ready가 아닌 경우
  - custom agent definition처럼 sibling skill이 소유하는 경우
  - empirical evaluation처럼 다른 workflow가 소유하는 경우
  - user-gated clarification, approval, safety, destructive, external decision이 먼저 필요한 경우
- no-handoff routing note에는 다음을 포함한다.
  - reason: 왜 지금 subagent prompt가 부적절한가
  - right skill or workflow: 어디로 라우팅해야 하는가
  - main-thread next action: 메인 에이전트가 다음에 해야 할 일
  - main-thread blocked state: 라우팅/질문/승인 때문에 blocked인지, non-overlapping work를 계속할 수 있는지
  - re-entry condition: 어떤 조건이 만족되면 `subagent-gate`로 다시 handoff를 준비할 수 있는지
- no-handoff 상태에서는 아직 존재하지 않는 subagent ownership이나 pass/fail criteria를 채우지 않는다.

## 위임 판단 규칙

- critical path를 즉시 막는 작업이면 main agent가 직접 처리하거나, subagent 결과가 올 때까지 무엇이 blocked인지 명시한다.
- 병렬 sidecar 작업이면 main agent가 동시에 진행할 non-overlapping work를 함께 정한다.
- implementation subagent라면 write scope와 ownership을 명확히 적고, 다른 작업자가 같은 저장소에서 작업 중일 수 있음을 알려야 한다.
- verification subagent라면 수정 권한 없이 확인 대상, 명령, pass/fail criteria, report shape만 전달한다.
- user-gated approval은 subagent에게 넘기지 말고 main agent가 사용자에게 묻는다.

## 검토 질문

- 종료 시점이 prompt 앞부분에서 먼저 잠겨 있는가?
- subagent가 돌아와야 하는 조건과 멈춰야 하는 조건이 모두 있는가?
- 메인 에이전트가 blocked인지 non-overlapping work를 계속할 수 있는지 명시돼 있는가?
- no-handoff인 경우 reason, route, next action, blocked state, re-entry condition이 있는가?
- context packet이 subagent 관점의 필요한 항목으로 재작성됐는가?
- main agent 전체 맥락이나 추측이 불필요하게 누출되지 않았는가?
- user-gated, destructive, irreversible, external-action decision이 subagent에게 위임되지 않았는가?
- main agent가 subagent를 기다리는 동안 할 수 있는 non-overlapping work가 분리됐는가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 예.
- 그렇다면 왜 필요한가 / 아니라면 어떤 sibling context를 허용하는가: subagent handoff gate는 custom agent 정의나 empirical evaluation 없이도 호출 전 prompt packet을 판단해야 하므로 독립 실행 가능해야 한다. 다만 custom agent 파일 자체를 만드는 작업은 `subagent-creator`, instruction evaluation은 `empirical-prompt-tuning`으로 라우팅한다.

## 확장 원칙

- subagent runtime control 규칙이 추가되면 `subagent-gate`가 소유한다.
- custom agent TOML shape나 reusable role definition 규칙은 `subagent-creator`로 넘긴다.
- evaluation scenario 반복 설계는 `empirical-prompt-tuning`으로 넘긴다.
