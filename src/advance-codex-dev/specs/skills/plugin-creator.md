## 사용자 스펙 의도

- skill 묶음이 아니라 coherent plugin boundary를 먼저 정의하고 싶다.
- multi-skill plugin이면 `<plugin>-guide`와 packaging contract를 같이 설계하고 싶다.
- canonical scaffold 위에 plugin-level 설계 규칙을 덧붙이고 싶다.

---

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

- 대표 표면: `advance-codex-dev/skills/plugin-creator/SKILL.md`
- 관련 상위 라우팅: `advance-codex-dev-guide`

## 핵심 처리 계약

- plugin boundary를 skill 목록보다 먼저 고정한다.
- 각 bundled skill이 왜 같은 plugin에 속해야 하는지 설명 가능해야 한다.
- multi-skill plugin이면 `<plugin>-guide`를 두고 cross-skill usage guidance를 그 안에 모은다.
- manifest는 실제 shipped surface와 정확히 맞아야 한다.

## 검토 질문

- plugin boundary를 skill 목록보다 먼저 잠갔는가?
- bundled skill들이 왜 같은 plugin에 속하는지 설명 가능한가?
- `<plugin>-guide`가 필요한 구조인데 빠뜨리고 있지 않은가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 아니오.
- 그렇다면 왜 필요한가 / 아니라면 어떤 sibling context를 허용하는가: canonical system `plugin-creator`를 바탕으로 쓰는 확장 스킬이므로 base workflow 전제는 허용하지만, plugin boundary와 packaging 판단 기준은 이 스펙에서 명시해야 한다.

## 확장 원칙

- plugin packaging 규칙이 바뀌면 `plugin.md`, guide skill, manifest 검토 규칙을 함께 갱신한다.
- plugin boundary와 무관한 skill-level detail은 별도 skill spec으로 분리한다.
