## 사용자 스펙 의도

- advance-codex에 새로운 스킬 `subagent-work`를 추가하고 싶다.
- 메인 에이전트는 기본적으로 오케스트레이션을 기반으로 작업하고, subagent에게 실제 작업을 위임해서 작업하는 방식이어야 한다.
- Codex의 subagent는 한 번 spawn하면 이어서 대화하며 작업하거나, subagent가 또 다른 subagent를 spawn할 수 있다.
- 특정 spec이나 작업을 진행할 때 처음에 subagent를 spawn하고, 결정하거나 사용자와 대화가 필요한 것은 메인 에이전트가 맡고, 관련 작업 준비가 완료되면 subagent가 작업과 검증을 진행하고, 사용자가 해당 작업의 완료를 알리는 시점(commit-ready, 저장 등)이 되면 subagent를 dispose해서 메인 에이전트의 context window를 줄이고 싶다.
  - `subagent-work`는 어느 정도로 강한 운영 규칙이어야 할까요?
    - 엄격한 생애주기 [선택]
  - `subagent-gate`와의 관계는 어떻게 둘까요?
    - 독립 완결 [선택]
  - worker subagent는 언제 close/dispose하는 규칙이 좋을까요?
    - 작업 단위 기준 [선택]

---

# subagent-work 스킬 스펙

## 목적

`subagent-work`는 하나의 reviewable work unit 동안 worker subagent를 생성, 운영, 동기화, 검증, 종료하는 엄격한 생애주기 스킬입니다.
메인 thread는 사용자 대화, 범위 결정, 승인 경계, 통합 판단을 소유하고, worker subagent는 준비된 작업 단위의 구현과 1차 검증을 수행합니다.
작업 단위가 완료되면 worker를 닫고, 다음 작업 단위는 새 worker와 새 compact handoff로 시작해 메인 thread context window를 절약합니다.

## 경계

- 포함:
  - worker subagent를 사용할 reviewable work unit 정의
  - spawn 전 work packet 작성
  - worker ownership, write scope, output contract, verification expectation 지정
  - main-owned decision, user routing, approval boundary 분리
  - worker와의 sync cadence와 escalation condition 정의
  - worker 결과 통합 전 main thread 검토 기준
  - 다음 worker를 위한 compact handoff summary 작성
  - work unit 완료 후 worker close/dispose 기준
- 제외:
  - custom agent TOML definition 설계
  - one-off handoff packet만 필요한 좁은 위임
  - 일반 프로젝트 실행 workflow 전체
  - user-gated approval 대체
  - destructive, external, commit, push, PR, publish, release, version bump 실행
  - worker를 여러 reviewable work unit에 걸쳐 무기한 유지하는 장기 운영

## 처리하려는 작업 형태

- 하나의 plugin, skill, document, fixture, script, test, 또는 config change unit을 worker에게 맡기려는 경우
- 메인 thread가 사용자와 계속 대화해야 하지만 실제 구현 세부 맥락은 worker에게 격리하고 싶은 경우
- worker가 한 번 이상의 follow-up message를 받으며 구현과 검증을 진행해야 하는 경우
- 작업 완료 후 worker를 닫고 다음 단위에는 새 worker를 쓰고 싶은 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `advance-codex-dev/skills/subagent-work/SKILL.md`
- 호출 방식: 메인 thread context window를 아끼기 위해 하나의 reviewable work unit을 worker subagent에게 맡기려 할 때 호출한다.

## 핵심 처리 계약

- 먼저 work unit을 정한다. work unit은 reviewable하거나 commit-sized인 단위여야 한다.
- worker가 맡을 구현/검증 책임과 main thread가 맡을 사용자/승인/통합 책임을 분리한다.
- spawn 전 packet에는 goal, assigned scope, editable paths, non-goals, approval limits, expected checks, return format, sync points, close criteria를 포함한다.
- worker는 user-gated approval, destructive action, external action, commit, push, PR, publish, release, version bump를 승인하거나 실행하지 않는다.
- worker가 scope breach, approval need, ambiguous requirement, conflicting edits, verification failure, missing dependency를 만나면 멈추고 main thread로 escalates한다.
- 메인 thread는 worker 결과를 그대로 최종 결과로 취급하지 않고 changed paths, assumptions, validation, residual risk를 검토한다.
- work unit이 끝나면 worker를 close/dispose하고, 이어지는 작업에는 compact handoff summary를 사용해 새 worker를 연다.
- worker가 또 다른 subagent를 spawn할 수 있는 환경이라도 nested delegation은 명시된 assigned scope 안에서만 허용하며, approval-sensitive action을 하위 subagent로 우회하지 않는다.

## 생애주기 규칙

- `Prepare`: work unit, main-owned decisions, worker-owned execution, approval boundary를 정한다.
- `Spawn`: worker에게 독립적으로 실행 가능한 packet을 보낸다.
- `Operate`: worker는 구현과 1차 검증을 진행하고, main thread는 사용자 대화와 non-overlapping orchestration만 수행한다.
- `Sync`: worker는 정해진 checkpoint 또는 blocker에서 changed paths, decisions, validation, risk를 반환한다.
- `Integrate`: main thread는 결과를 읽고 필요한 추가 지시, 재작업, 사용자 질문, 또는 다음 단위 분리를 결정한다.
- `Close`: work unit이 reviewable 상태가 되면 worker를 close/dispose한다.
- `Handoff`: 다음 work unit이 있으면 완료 요약, 남은 위험, 다음 scope만 담은 compact handoff를 작성한다.

## Worker Packet 규칙

- worker packet은 메인 thread의 전체 대화가 아니라 worker에게 필요한 최소 정보만 포함한다.
- packet에는 `Return when`, `Stop if`, `Close plan`, `Main-owned decisions`, `Assigned work unit`, `Editable scope`, `Do not touch`, `Validation`, `Output`을 포함한다.
- 편집 가능성이 있으면 worker에게 같은 저장소에서 혼자 작업하는 것이 아니며, 관련 없는 변경을 되돌리거나 덮어쓰면 안 된다고 명시한다.
- worker가 여러 파일을 만질 수 있으면 ownership과 write scope를 먼저 제한한다.

## Close / Dispose 규칙

- worker는 하나의 reviewable work unit을 넘기면 닫는다.
- 다음 경우에는 worker를 닫거나 더 이상 진행시키지 않는다.
  - work unit 구현과 1차 검증이 끝났다.
  - main thread가 commit-ready, 저장, 완료, 또는 다음 단위 전환 신호를 받았다.
  - user-gated approval이 필요하다.
  - 요구사항이나 범위가 바뀌어 새 packet이 필요하다.
  - worker context가 누적되어 compact handoff보다 비싸거나 흐려졌다.
- 같은 worker를 다음 work unit으로 계속 쓰는 것은 예외이며, work unit 경계가 실제로 바뀌지 않은 경우에만 허용한다.

## 검토 질문

- 이 작업은 하나의 reviewable work unit으로 자를 수 있는가?
- worker가 구현과 1차 검증을 맡고, main thread가 사용자 대화와 승인 경계를 맡는가?
- packet이 worker 관점의 최소 정보만 담는가?
- worker가 멈춰야 하는 approval, ambiguity, conflict, failure 조건이 명시됐는가?
- worker 결과를 main thread가 통합 검토할 수 있는 output contract가 있는가?
- 다음 작업 단위로 넘어가기 전에 worker를 닫고 compact handoff를 남기는가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 예.
- 그렇다면 왜 필요한가 / 아니라면 어떤 sibling context를 허용하는가: `subagent-work`는 worker 생애주기 운영을 단독으로 수행할 수 있어야 한다. sibling skill이 one-off handoff나 custom agent definition을 더 세밀하게 도울 수는 있지만, 이 skill의 runtime은 spawn부터 close/dispose까지 필요한 기본 packet과 판단 기준을 자체적으로 제공해야 한다.

## 확장 원칙

- 새 규칙은 worker lifecycle, main/worker 책임 분리, context 절약, close/dispose, compact handoff 중 하나를 더 명확하게 만들 때만 추가한다.
- 일반 프로젝트 관리, 도메인 구현 방식, custom agent file definition, release/commit execution workflow는 이 skill로 흡수하지 않는다.
