# advance-codex-guide 스킬 스펙

## 목적

`advance-codex-guide`는 `advance-codex` 플러그인의 첫 라우터로서, Codex 활용 개선 작업의 주된 artifact type을 분류하고 가장 좁은 built-in skill과 실행 순서를 정하는 역할을 맡습니다.

## 경계

- 포함:
  - skill, empirical evaluation, tool-use guidance, plugin, session workflow, commit workflow, custom agent 중 주된 artifact 분류
  - revision, creation, packaging 여부 판단
  - multi-artifact 작업의 실행 순서 결정
- 제외:
  - 각 artifact의 상세 구현
  - plugin 외부 workflow 선택
  - domain-specific execution guidance

## 처리하려는 작업 형태

- "어떤 creator skill을 먼저 써야 하는가"가 병목인 작업
- 여러 Codex-facing artifact가 섞여 있어 우선순위를 정해야 하는 작업
- plugin 내부 guide가 먼저 작업 형태를 분류해야 하는 작업

## 엔트리포인트 / 대표 표면

- 대표 표면: `advance-codex/skills/advance-codex-guide/SKILL.md`
- 상위 맥락: `workflow-kit`이 execution workflow를 고른 뒤, creator-oriented specialist routing이 필요할 때 사용한다

## 핵심 처리 계약

- 먼저 주된 reusable artifact를 하나로 고정한다.
- creation, revision, packaging 중 현재 task shape를 판별한다.
- plugin-level concern이면 plugin packaging을 우선하고, 하위 skill은 그 다음 순서로 둔다.
- multi-artifact 작업이면 실행 순서를 명시하고 섞지 않는다.
- 출력은 artifact 분류, 선택된 creator skill, 실행 순서, 주요 리스크를 포함해야 한다.

## 독립성 원칙

- 이 스킬은 creator artifact의 상세 설계를 소유하지 않는다.
- sibling skill의 구현 규칙을 복제하지 않고 routing 기준만 소유한다.

## 확장 원칙

- 새 artifact type이 plugin boundary 안에서 독립 concern이 될 때만 라우팅 분기를 추가한다.
- routing 분기가 바뀌면 plugin spec과 관련 skill spec을 함께 갱신한다.

