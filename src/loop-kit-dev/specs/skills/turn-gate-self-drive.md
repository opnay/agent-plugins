## 사용자 스펙 의도

- `self-drive`를 `turn-gate` 내부 mode가 아니라 별도 skill로 분리하고 싶다.
- 그래도 `self-drive`는 독립된 loop gate가 아니라 `turn-gate`를 직접 참고하는 overlay여야 한다.
- `turn-gate` 스킬 본문에서는 self-drive packet/answer/pause/recovery 계약을 제거하고 싶다.
- self-drive 동안 각 flow 종료의 질문 라우팅 파트에서 `turn-gate`와 `turn-gate-self-drive` skill을 다시 읽는 과정을 추가하고 싶다.
- 이 reread는 context cleanup 또는 compaction이 실제로 발생했는지 감지해서 실행하는 조건부 동작이 아니어야 한다.
- assistant는 context cleanup 시점을 안정적으로 알 수 없으므로, flow 종료 후 다음 질문을 만들기 전마다 무조건 reread하는 deterministic refresh 절차여야 한다.
- 목적은 context cleanup 여부를 판별하는 것이 아니라, 긴 self-drive 이후 다음 질문/라우팅 전에 base loop contract와 self-drive overlay contract를 다시 적재하는 것이다.

---

# turn-gate-self-drive 스킬 스펙

## 목적

`turn-gate-self-drive`는 `loop-kit-dev`에 속한 별도 plugin skill이며, 같은 플러그인에 포함된 `turn-gate` skill의 turn continuity 계약을 기반으로 동작하는 self-drive question-routing overlay입니다.
turn continuity, flow record, Continuity Guard, next-flow reopening은 같은 플러그인의 `turn-gate` skill 계약을 따르고, 이 skill은 bounded decision을 subagent question packet으로 라우팅하는 계약을 소유합니다.

## 경계

- 포함:
  - self-drive question packet 작성
  - subagent answer contract
  - user intervention 중 stale answer 처리
  - context gap recovery
  - self-drive flow 종료 후 질문 라우팅 직전마다 무조건 수행하는 context-refresh reread
  - approval boundary에서 같은 플러그인의 `turn-gate` skill이 제공하는 user-gated question routing으로 복귀
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

- 대표 plugin skill: `loop-kit-dev/skills/turn-gate-self-drive/SKILL.md`
- base skill in the same plugin: `loop-kit-dev/skills/turn-gate/SKILL.md`
- 관련 상위 라우팅: `loop-kit-dev-guide`

## 핵심 처리 계약

- 같은 플러그인의 `turn-gate` skill을 base loop contract로 적용한다.
- self-drive는 `turn-gate`를 대체하지 않고, bounded decision routing만 subagent packet으로 위임하는 overlay로 동작한다.
- 모든 self-drive packet은 현재 `Continuity Guard`를 포함해야 한다.
- subagent answer는 terminal summary 대신 계속 이어질 next action을 제시해야 한다.
- 사용자 중간 개입은 pending/stale subagent answer보다 우선한다.
- explicit approval boundary가 있으면 autonomous routing을 일시 중지하고 같은 플러그인의 `turn-gate` skill이 제공하는 user-gated question routing으로 복귀한다.
- self-drive가 한 flow의 result reporting을 끝내고 다음 flow 질문, scope 질문, self-drive subagent packet 질문, 또는 user-gated approval 질문으로 넘어가기 직전에는 같은 플러그인의 `turn-gate` skill과 `turn-gate-self-drive` skill을 항상 다시 읽는다.
- 이 reread는 context cleanup 또는 compaction 발생 여부를 감지해서 조건부로 실행하는 절차가 아니다. self-drive는 cleanup 시점을 안정적으로 알 수 없으므로, flow-end question-routing boundary마다 무조건 다시 읽는 deterministic refresh 절차로 다룬다.
- reread의 목적은 긴 self-drive 수행 중 누적된 요약/압축/부분 문맥 drift를 줄이고, 다음 질문을 만들기 전에 base loop contract와 self-drive overlay contract를 다시 적재하는 것이다. reread 후 현재 `Continuity Guard`, active flow record, pending next-flow candidates, approval boundary를 재확인해야 한다.
- reread 이후에도 base turn continuity와 user-gated approval boundary는 `turn-gate` 계약을 따른다. self-drive는 reread를 이유로 approval, destructive, irreversible, external-action, safety 결정을 자동 위임하지 않는다.

## 검토 질문

- base `turn-gate` 계약을 적용했는가?
- subagent에게 보낼 decision이 충분히 bounded한가?
- packet과 answer에 `Continuity Guard`와 next action이 포함됐는가?
- user intervention이 stale answer보다 우선했는가?
- approval boundary에서 turn stop이 아니라 같은 플러그인의 `turn-gate` skill이 제공하는 user-gated question routing으로 복귀했는가?
- flow-end question-routing boundary마다 context cleanup 감지 여부와 무관하게 `turn-gate`와 `turn-gate-self-drive` skill을 다시 읽고, `Continuity Guard`와 pending next-flow candidates를 재확인했는가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 아니오.
- 그렇다면 왜 필요한가 / 아니라면 어떤 plugin context를 허용하는가: 이 skill은 `loop-kit-dev` 안에서 같은 플러그인의 `turn-gate` skill을 base contract로 전제하는 overlay다. 단, self-drive packet/answer/pause/recovery 책임은 이 스펙만 읽어도 이해 가능해야 한다.

## 확장 원칙

- self-drive packet, answer, recovery, stale-answer 규칙은 이 skill의 책임으로 유지한다.
- flow-end deterministic reread/context-refresh 규칙은 self-drive overlay의 책임으로 유지하되, reread 대상인 base continuity 계약은 같은 플러그인의 `turn-gate` skill이 소유한다.
- turn continuity 자체와 user-gated flow record 규칙은 같은 플러그인의 `turn-gate` skill 책임으로 유지한다.
