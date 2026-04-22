# designer-kit-guide 스킬 스펙

## 목적

`designer-kit-guide`는 `designer-kit` 플러그인의 엔트리포인트로서 design-only task를 적절한 design stage로 라우팅하는 스킬입니다.

## 경계

- 포함:
  - brief, critique, screen spec 중 시작 stage 선택
  - multi-stage design task의 다음 handoff 결정
- 제외:
  - 각 design artifact의 상세 작성
  - implementation or Figma execution planning

## 처리하려는 작업 형태

- design task가 아직 어떤 stage에 있는지 불명확한 경우
- brief 이후 critique로 갈지, 바로 screen spec으로 갈지 정해야 하는 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `designer-kit/skills/designer-kit-guide/SKILL.md`

## 핵심 처리 계약

- 현재 task의 병목이 direction definition, critique, screen specification 중 무엇인지 고정한다.
- routing 결과에는 선택된 stage와 그 이유, 다음 handoff를 포함해야 한다.
- specialist detail은 sibling skill에 넘기고 guide 안에 중복 정의하지 않는다.

## 독립성 원칙

- 이 스킬은 design stage routing만 소유한다.
- sibling skill 문맥 없이도 언제 어떤 stage를 먼저 써야 하는지 설명 가능해야 한다.

## 확장 원칙

- design stage가 늘어나거나 순서가 바뀌면 plugin spec과 함께 갱신한다.

