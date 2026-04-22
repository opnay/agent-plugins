# plugin-creator 스킬 스펙

## 목적

`plugin-creator`는 canonical system `plugin-creator` 위에 top-down plugin 설계 규칙을 덧붙이는 확장 스킬입니다.
핵심은 loose skill bundle이 아니라 coherent plugin boundary를 먼저 정의하고, 그 내부 skill과 `<plugin>-guide`를 그 boundary에 맞게 배치하는 것입니다.

## 경계

- 포함:
  - plugin boundary 우선 설계
  - bundled skill coherence 점검
  - `<plugin>-guide` 필요성과 역할 정의
  - manifest와 실제 표면의 정합성 검토
- 제외:
  - canonical base scaffold 전체 대체
  - 개별 skill authoring 세부
  - unrelated marketplace policy 전반

## 처리하려는 작업 형태

- 새 plugin bundle을 만들거나 기존 plugin 구조를 재정렬하는 경우
- 여러 skill을 하나의 plugin으로 묶을지 판단해야 하는 경우
- guide skill과 packaging contract를 함께 설계해야 하는 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `advance-codex/skills/plugin-creator/SKILL.md`
- 관련 상위 라우팅: `advance-codex-guide`

## 핵심 처리 계약

- plugin boundary를 skill 목록보다 먼저 고정한다.
- 각 bundled skill이 왜 같은 plugin에 속해야 하는지 설명 가능해야 한다.
- multi-skill plugin이면 `<plugin>-guide`를 두고 cross-skill usage guidance를 그 안에 모은다.
- manifest는 실제 shipped surface와 정확히 맞아야 한다.

## 독립성 원칙

- 이 스킬은 plugin-level architecture extension을 소유하며 canonical scaffold 전체를 대체하지 않는다.
- sibling creator skill의 숨은 맥락 없이도 plugin 설계 원칙을 이해할 수 있어야 한다.

## 확장 원칙

- plugin packaging 규칙이 바뀌면 `plugin-spec.md`, guide skill, manifest 검토 규칙을 함께 갱신한다.
- plugin boundary와 무관한 skill-level detail은 별도 skill spec으로 분리한다.

