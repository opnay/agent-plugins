## 사용자 스펙 의도

- `rpg-kit`이라는 새 plugin으로 role-assigned subagent orchestration을 다루고 싶다.
- `subagent-role` skill을 통해 어떤 역할의 subagent를 spawn할지 판단하고 싶다.
- subagent와 어떻게 동작하는지 관찰하고 배워가는 skill 체계를 만들고 싶다.
- role은 단순 직함이 아니라 전문성 조합으로 정의하고 싶다.

---

# rpg-kit-guide 스킬 스펙

## 목적

`rpg-kit-guide`는 현재 작업이 role-based subagent orchestration 범위인지 판단하고, 맞다면 `subagent-role`로 라우팅하는 entrypoint skill입니다.

## 경계

- 포함:
  - `rpg-kit` 적합성 판단
  - role-assigned subagent가 필요한지 판단
  - role specialty가 필요한지 판단
  - `subagent-role`로 라우팅하기 전 목표와 learning focus 정리
- 제외:
  - 구체적인 role packet 상세 작성
  - subagent spawn 실행 자체
  - 일반 workflow selection 전반

## 처리하려는 작업 형태

- subagent role 설계가 핵심인 작업
- 직무, 책임 영역, 출신 배경, 범용 전문성을 섞어 role specialty를 정의해야 하는 작업
- 역할 분담, answer contract, caller integration rule을 먼저 정해야 하는 작업
- subagent behavior를 관찰해 다음 orchestration 규칙으로 반영하려는 작업

## 엔트리포인트 / 대표 표면

- 대표 표면: `rpg-kit/skills/rpg-kit-guide/SKILL.md`
- 관련 상위 라우팅: 없음

## 핵심 처리 계약

- 먼저 사용자의 목적이 subagent role orchestration인지 판단한다.
- role-assigned subagent가 필요하지 않으면 plugin 밖에서 처리해도 된다고 명시한다.
- 필요하면 `subagent-role`로 넘기기 전에 desired role shape, specialty mix, learning focus, spawn boundary를 요약한다.
- runtime/tool policy상 subagent spawn이 명시적 허용을 요구하면 그 경계를 유지한다.

## 라우팅 규칙

- role packet, subagent scope, answer contract가 핵심이면 `subagent-role`로 보낸다.
- role specialty를 구체화해야 하면 `subagent-role`로 보낸다.
- 실제 subagent behavior를 비교하거나 학습하려는 목적이면 `subagent-role`로 보낸다.
- 단순 구현, 일반 코드 리뷰, 단발 질문이면 `rpg-kit`를 쓰지 않는다.
- 사용자가 subagent 사용을 명시하지 않았고 task가 local critical path에 가깝다면 caller가 직접 처리하는 쪽을 우선한다.

## 검토 질문

- 이 작업의 핵심이 정말 subagent role orchestration인가?
- role-assigned subagent가 caller의 immediate blocker를 대신 맡아 병목을 만들지는 않는가?
- spawn 전 role packet과 expected output이 충분히 좁은가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 아니오.
- 그렇다면 왜 필요한가 / 아니라면 어떤 sibling context를 허용하는가: guide는 `subagent-role`이라는 sibling execution surface를 전제로 한다. 다만 `rpg-kit` 적합성 판단 기준은 이 스펙만으로 이해 가능해야 한다.

## 확장 원칙

- 새 skill을 추가할 때는 먼저 `rpg-kit` plugin boundary 안의 role orchestration 책임인지 확인한다.
- guide에는 routing 판단만 두고, role packet 상세 계약은 `subagent-role`에 둔다.
