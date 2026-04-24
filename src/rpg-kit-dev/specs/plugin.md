# RPG Kit 플러그인 스펙

## 플러그인 목적

`rpg-kit-dev`은 역할이 할당된 subagent를 어떻게 설계하고, 언제 spawn하며, 그 결과를 어떻게 읽어 다음 orchestration 규칙을 배울지 다루는 plugin입니다.
핵심 책임은 subagent role assignment를 즉흥적인 delegation이 아니라 명시적인 role packet, answer contract, integration rule, learning note로 다루게 하는 것입니다.

## 플러그인 경계와 비목표

- 포함:
  - role-assigned subagent 사용 여부 판단
  - role specialty 조합 정의
  - subagent role packet 설계
  - subagent spawn 전 승인/정책/중복 작업 경계 확인
  - subagent 답변을 caller가 통합 가능한 evidence로 읽는 규칙
  - 역할별 subagent behavior를 관찰하고 다음 role 설계에 반영하는 learning loop
- 제외:
  - 모든 작업에 subagent 사용을 강제하는 범용 delegation policy
  - subagent 결과를 검증 없이 최종 답으로 승격하는 흐름
  - runtime/tool policy가 허용하지 않는 subagent spawn 우회
  - 특정 제품 기능이나 도메인별 구현 자체

## 처리하려는 작업 형태

- 역할을 나눈 subagent가 실제로 도움이 되는지 판단해야 하는 작업
- 직무, 책임 영역, 출신 배경, 범용 전문성을 조합해 subagent 관점을 정해야 하는 작업
- subagent에게 보낼 role, scope, output contract를 명확히 해야 하는 작업
- 여러 subagent 결과를 비교하거나 caller-side decision으로 통합해야 하는 작업
- subagent behavior를 관찰해 다음 orchestration 규칙을 개선하려는 작업

## 엔트리포인트 / 대표 표면

- 대표 엔트리포인트: `rpg-kit-dev-guide`
- 대표 실행 표면: `subagent-role`
- 대표 스펙: `rpg-kit-dev/specs/plugin.md`
- skill 상세 스펙 위치: `rpg-kit-dev/specs/skills/*.md`
- 평가 설계: `rpg-kit-dev/specs/subagent-role-packet-evaluation.md`

## 내장 skill 체계

- `rpg-kit-dev-guide`: 현재 작업이 `rpg-kit-dev`의 역할 기반 subagent orchestration 범위인지 판단하고 적절한 skill로 라우팅한다.
  - spec: `rpg-kit-dev/specs/skills/rpg-kit-dev-guide.md`
- `subagent-role`: 역할이 할당된 subagent를 위한 packet, spawn rule, answer contract, learning note를 정의한다.
  - spec: `rpg-kit-dev/specs/skills/subagent-role.md`

## SDD 운영 원칙

- 먼저 plugin boundary를 정의하고, 그 다음 role-oriented skill 책임을 정의한다.
- subagent spawn은 runtime/tool policy와 user intent를 넘어서지 않는다.
- role packet은 role specialty, task ownership, write scope, expected output, integration rule, stop condition을 분리해서 적는다.
- subagent 결과는 caller가 검증하고 통합해야 하는 evidence로 취급한다.
- 새 role pattern이 반복되면 `subagent-role` spec 또는 reference로 승격할지 검토한다.

## 현재 구조 메모

- 초기 버전은 `rpg-kit-dev-guide`와 `subagent-role` 두 skill만 포함한다.
- 이 플러그인은 subagent orchestration을 배우는 킷이며, subagent runtime 자체를 제공하지 않는다.
