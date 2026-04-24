## 사용자 스펙 의도

- 하나의 bounded issue를 짧은 loop로 반복 개선하고 싶다.
- 가장 작은 fix로 가설을 확인하고 다음 loop 필요성을 다시 판단하고 싶다.
- broad execution이나 review triage 전체와는 분리하고 싶다.

---

# ralph-loop 스킬 스펙

## 목적

`ralph-loop`는 하나의 bounded issue를 짧은 fix-verify-reassess cycle로 반복 개선하는 workflow 스킬입니다.

## 경계

- 포함:
  - bounded issue selection
  - smallest useful fix
  - immediate verification
  - next loop decision
- 제외:
  - broad end-to-end delivery
  - review finding triage 전반
  - planning-only work

## 처리하려는 작업 형태

- UI polish, refactor stabilization, flaky issue reduction 같은 bounded refinement
- progressive improvement가 더 적합한 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `workflow-kit-dev/skills/ralph-loop/SKILL.md`
- 관련 상위 라우팅: `workflow-kit-dev-guide`

## 핵심 처리 계약

- 한 loop는 하나의 primary issue만 다룬다.
- 가장 작은 fix로 가설을 검증한다.
- 각 loop 후에 개선 여부와 다음 loop 정당성을 다시 판단한다.

## 검토 질문

- 지금 loop가 하나의 primary issue만 다루고 있는가?
- 가장 작은 fix로 가설을 검증하고 있는가?
- 다음 loop가 정말 정당한지 재평가했는가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 예.
- 그렇다면 왜 필요한가 / 아니라면 어떤 sibling context를 허용하는가: bounded loop 규칙은 hidden sibling context 없이 읽혀야 하며, broad execution이나 review handling은 별도 skill에 남겨야 한다.

## 확장 원칙

- stop condition이나 issue threshold가 바뀌면 guide skill과 함께 갱신한다.
