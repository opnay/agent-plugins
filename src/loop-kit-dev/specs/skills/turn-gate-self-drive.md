## 사용자 스펙 의도

- `self-drive`를 `turn-gate` 내부 mode가 아니라 별도 skill로 분리하고 싶다.
- 그래도 `self-drive`는 독립된 loop gate가 아니라 `turn-gate`를 직접 참고하는 overlay여야 한다.
- `turn-gate` 스킬 본문에서는 self-drive packet/answer/pause/recovery 계약을 제거하고 싶다.

---

# turn-gate-self-drive 스킬 스펙

## 목적

`turn-gate-self-drive`는 `turn-gate`의 turn continuity 계약을 기반으로 동작하는 self-drive question-routing overlay입니다.
turn continuity, flow record, Continuity Guard, next-flow reopening은 `turn-gate` 계약을 따르고, 이 skill은 bounded decision을 subagent question packet으로 라우팅하는 계약을 소유합니다.

## 경계

- 포함:
  - self-drive question packet 작성
  - subagent answer contract
  - user intervention 중 stale answer 처리
  - context gap recovery
  - approval boundary에서 user-gated `turn-gate` 질문 표면 사용
- 제외:
  - `turn-gate` 기본 phase loop 재정의
  - current-phase internal mode selection 자체
  - user-gated next-flow question 자체의 운영
  - runtime/tool/safety approval 대체

## 처리하려는 작업 형태

- 사용자가 autonomous continuation 또는 self-driving progress를 원한 경우
- `turn-gate`가 governing loop이고, bounded decision을 subagent question packet으로 판단할 수 있는 경우
- mode selection, criteria, scope assumption, verification choice, next-flow decision이 bounded packet으로 전달 가능한 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `loop-kit-dev/skills/turn-gate-self-drive/SKILL.md`
- base skill: `loop-kit-dev/skills/turn-gate/SKILL.md`
- 관련 상위 라우팅: `loop-kit-dev-guide`

## 핵심 처리 계약

- `turn-gate`를 base loop contract로 적용한다.
- self-drive는 `turn-gate`를 대체하지 않고, bounded decision routing만 subagent packet으로 위임하는 overlay로 동작한다.
- 모든 self-drive packet은 현재 `Continuity Guard`를 포함해야 한다.
- subagent answer는 terminal summary 대신 계속 이어질 next action을 제시해야 한다.
- 사용자 중간 개입은 pending/stale subagent answer보다 우선한다.
- explicit approval boundary가 있으면 autonomous routing을 일시 중지하고 user-gated `turn-gate` 질문 표면을 사용한다.

## 검토 질문

- base `turn-gate` 계약을 적용했는가?
- subagent에게 보낼 decision이 충분히 bounded한가?
- packet과 answer에 `Continuity Guard`와 next action이 포함됐는가?
- user intervention이 stale answer보다 우선했는가?
- approval boundary에서 turn stop이 아니라 user-gated 질문 표면을 사용했는가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 아니오.
- 그렇다면 왜 필요한가 / 아니라면 어떤 sibling context를 허용하는가: 이 skill은 `turn-gate`를 base contract로 전제하는 overlay다. 단, self-drive packet/answer/pause/recovery 책임은 이 스펙만 읽어도 이해 가능해야 한다.

## 확장 원칙

- self-drive packet, answer, recovery, stale-answer 규칙은 이 skill의 책임으로 유지한다.
- turn continuity 자체와 user-gated flow record 규칙은 `turn-gate` 책임으로 유지한다.
