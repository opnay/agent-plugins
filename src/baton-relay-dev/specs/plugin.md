## 사용자 스펙 의도

- 새로운 루프 스킬을 만들려고 한다. 이 스킬을 사용하면 해당 에이전트는 오케스트레이션으로서 동작한다. subagent가 모든 작업을 진행해야 한다. 어떤 작업을 받으면 그 작업을 구조분해해서 git worktree로 구분해 subagent를 구동한다. 각각의 worktree에서 작업이 완료되면 그 작업분을 커밋해서 메인 작업위치에 커밋을 가져오고, 작업하던 worktree는 정리해야 한다. 구조분해 방식은 문제를 해결하는 방식, 코드를 구분하는 방식, workflow를 구분하는 방식 등 다양한 기준점이 필요하다.
- subagent 사이클도 지정해야 한다. 작업 시작부터 작업 완료까지 subagent가 동작하고, 작업 완료가 되면 subagent는 종료한다. 다음 작업을 위해 새로 생성한다. subagent의 작업 완료는 git commit 이후, 메인 에이전트가 병합 준비를 요청하면 subagent가 메인 에이전트가 동작하는 브랜치로 rebase하고, 완료되면 메인 에이전트가 해당 커밋을 가져온다. rebase는 완료된 상태이기 때문에 conflict는 발생할 가능성이 없다.
- 플러그인 이름은 `baton-relay`로 정한다. 스킬 이름은 `manager`로 정한다.

---

# Baton Relay 플러그인 스펙

## 플러그인 목적

`baton-relay-dev`는 큰 작업을 git worktree 단위로 나누고 fresh subagent에게 실행을 맡긴 뒤, commit과 rebase가 완료된 작업 단위만 메인 작업 위치로 회수하는 orchestration plugin입니다.
핵심 책임은 메인 에이전트를 구현자가 아니라 manager로 세우고, subagent 작업 시작부터 완료, handoff, 통합, worktree cleanup까지 하나의 안전한 relay loop로 다루게 하는 것입니다.

## 플러그인 경계와 비목표

- 포함:
  - 작업 구조분해 기준 선택
  - worktree-safe task slice 정의
  - subagent lifecycle 정의
  - worktree 생성, branch naming, dispatch packet 설계
  - subagent commit/rebase handoff 확인
  - 메인 작업 위치로 prepared commit 회수
  - 통합 후 worktree cleanup 기준
  - 실패, 기준 HEAD 변경, rebase 재요청 판단
- 제외:
  - subagent runtime 자체 제공
  - 모든 작업에 subagent 사용 강제
  - 승인 없는 commit, push, PR, publish, release, version bump, destructive work
  - subagent 결과를 검증 없이 최종 결과로 승격
  - 특정 언어, 제품, 도메인별 구현 전략 소유

## 처리하려는 작업 형태

- 여러 독립 작업으로 나눌 수 있는 구현, 조사, 검증, 문서화 작업
- 파일 또는 module ownership을 분리할 수 있는 변경
- 문제 가설, 코드 영역, workflow 단계, API 계약, 검증 경로별로 나눌 수 있는 작업
- 병렬 또는 순차 subagent 실행이 메인 에이전트의 통합 판단을 줄이는 작업
- 각 하위 작업이 commit 단위로 회수될 수 있는 작업

## 대표 표면

- 대표 실행 표면: `manager`
- 대표 스펙: `baton-relay-dev/specs/plugin.md`
- skill 상세 스펙 위치: `baton-relay-dev/specs/skills/manager/spec.md`

## 내장 skill 체계

- `manager`: 작업을 구조분해하고, worktree별 fresh subagent를 배정하며, commit/rebase handoff와 prepared commit 통합, cleanup gate를 관리한다.
  - spec: `baton-relay-dev/specs/skills/manager/spec.md`
  - intent and flow graph: `baton-relay-dev/specs/skills/manager/intent.md`

## SDD 운영 원칙

- plugin boundary는 "worktree relay orchestration"으로 유지하고, 일반 subagent role 설계나 범용 autopilot으로 넓히지 않는다.
- subagent는 하나의 task slice와 하나의 worktree에만 묶는다.
- subagent 완료 조건은 commit과 메인 기준 branch rebase 완료를 포함한다.
- 메인 에이전트는 미커밋 변경을 가져오지 않고, prepared commit만 회수한다.
- 기준 HEAD가 rebase 이후 움직였으면 메인 에이전트는 통합 전에 rebase 재요청을 선택한다.

## 현재 구조 메모

- 초기 실행 표면은 `manager` 하나로 좁힌다.
- 이 플러그인은 worktree/subagent 운영 계약을 제공하며, subagent spawn 도구나 GitHub publish workflow를 제공하지 않는다.
