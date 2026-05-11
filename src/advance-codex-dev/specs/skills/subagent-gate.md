## 사용자 스펙 의도

- subagent를 준비하는 과정에서 종료 시점 계획을 우선하고 싶다.
- 메인 에이전트의 맥락을 넘길 때는 서브 에이전트의 시점에서 필요한 항목만 정리해 전달해야 한다.

---

# subagent-gate 스킬 스펙

## 목적

`subagent-gate`는 subagent를 spawn하거나 기존 subagent에 message를 보내기 전에 handoff 계약을 잠그는 스킬입니다.
이 스킬은 세부 task 설명보다 먼저 subagent가 언제 멈추고 무엇을 반환해야 하는지 정리합니다.
메인 thread가 계속 진행할 수 있는 작업과 subagent 결과를 기다려야 하는 blocked state를 분리해, 위임이 병렬성을 실제로 만들도록 돕습니다.

## 경계

- 포함:
  - subagent handoff 준비
  - return point와 stop boundary 정의
  - subagent 관점의 minimal context packet 구성
  - ownership, write scope, output contract 정리
  - approval limits와 main-thread blocked state 분리
  - no-handoff decision 작성
- 제외:
  - custom agent definition 자체 설계
  - scenario-based instruction evaluation 자체 설계
  - user-gated approval 실행
  - destructive action 또는 external action 실행
  - commit, push, PR, publish, release, version bump 실행

## 처리하려는 작업 형태

- 새 subagent를 spawn하기 전 handoff prompt를 준비한다.
- 기존 subagent에 추가 message를 보내기 전 stop boundary와 output contract를 다시 정리한다.
- 여러 subagent가 동시에 움직이는 상황에서 ownership과 write scope를 분리한다.
- subagent 위임이 unsafe하거나 out of scope인 경우 no-handoff로 라우팅한다.

## 엔트리포인트 / 대표 표면

- 대표 표면: `advance-codex-dev/skills/subagent-gate/SKILL.md`
- 호출 방식: subagent를 spawn하거나 기존 subagent에 substantial message를 보내기 전에 호출한다.

## 핵심 처리 계약

- handoff는 task details보다 exit plan을 먼저 정의한다.
- exit plan에는 return when, stop if, close plan을 포함한다.
- main-thread blocked state와 병렬로 계속할 수 있는 non-overlapping work를 분리한다.
- context packet은 subagent 관점에서 필요한 goal, relevant facts, assigned scope, constraints, expected output, assumptions만 포함한다.
- 코드 또는 문서 변경을 맡길 때는 ownership과 write scope를 명시한다.
- 편집 가능성이 있으면 subagent가 같은 codebase에서 혼자 작업하는 것이 아니며 관련 없는 변경을 되돌리거나 덮어쓰면 안 된다고 알린다.
- output contract는 main thread가 바로 검토하고 통합할 수 있는 필드로 고정한다.
- user-gated, destructive, external, commit, push, PR, publish, release, version bump 결정은 subagent가 승인하거나 실행하지 않는다.
- 요청이 delegation-ready가 아니거나 이 skill 경계 밖이면 handoff packet 대신 no-handoff decision을 작성한다.

## Handoff 판단 규칙

- 다음 main-thread step이 subagent 결과에 즉시 막혀 있으면, 가능한 한 main thread가 직접 처리한다.
- subagent가 병렬로 처리할 수 있는 좁고 독립적인 작업이면 handoff를 준비한다.
- ownership이 불명확하거나 write scope가 겹치면 handoff 전에 범위를 줄인다.
- context를 많이 넘겨야만 성공하는 작업이면 no-handoff를 우선 검토한다.
- 사용자의 명시적 허가가 필요한 작업을 subagent에게 맡겨야 한다면 실행이 아니라 조사와 보고로 제한한다.

## No-Handoff 라우팅

다음 경우에는 subagent handoff를 만들지 않거나 조사 전용 handoff로 축소한다.

- 사용자 승인 자체가 핵심 작업인 경우
- custom agent definition처럼 상위 설계 판단이 아직 잠기지 않은 경우
- scenario-based instruction evaluation처럼 평가 설계와 실행 증거가 같은 thread에서 강하게 결합된 경우
- release, publish, version bump처럼 외부에 공개되거나 되돌리기 어려운 결정이 중심인 경우
- task가 너무 모호해 subagent의 return point와 stop boundary를 쓸 수 없는 경우

## 검토 질문

- subagent가 언제 돌아와야 하는지 먼저 썼는가?
- subagent가 멈춰야 하는 조건이 사용자 승인 경계와 맞는가?
- main thread가 blocked인지, 병렬로 할 일이 있는지 분리했는가?
- context packet이 subagent 관점의 최소 정보로 충분한가?
- ownership, write scope, output contract가 검토 가능한가?
- approval이 필요한 결정을 subagent에게 맡기지 않았는가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 예.
- 그렇다면 왜 필요한가 / 아니라면 어떤 sibling context를 허용하는가: subagent handoff gate는 다른 skill 이름이나 sibling context에 기대지 않고, subagent handoff 준비라는 좁은 계약만 설명해야 한다. 다른 workflow와 함께 쓰이더라도 해당 workflow는 generic category로만 언급하고, runtime은 handoff packet 작성에 필요한 지침만 포함한다.

## 확장 원칙

- 새 규칙은 subagent handoff의 return point, stop boundary, context 최소화, ownership, output contract, approval limit 중 하나를 더 명확하게 만들 때만 추가한다.
- 플러그인 전체 라우팅이나 다른 workflow의 상세 실행 절차는 이 spec에 추가하지 않는다.
