# Spec-Driven Development

이 문서는 이 저장소의 플러그인 spec-driven development 규칙을 소유합니다.
`AGENTS.md`는 운영 진입점이고, SDD 세부 계약은 이 문서를 기준으로 유지합니다.

## 목적

- 플러그인 변경은 spec을 먼저 확인하거나 갱신한 뒤 구현 표면을 맞춥니다.
- spec은 구현 세부보다 의도, 경계, 라우팅, 책임 배치를 먼저 고정합니다.
- 파일 종류나 구현 관성을 적더라도 핵심 판단 기준은 change pressure, ownership, routing 이유를 먼저 설명합니다.

## 기본 Spec 구조

각 개발 원본 플러그인은 다음 spec 표면을 사용합니다.

- `src/<plugin-name>-dev/specs/plugin.md`: 플러그인 목적, 경계, 라우팅 표면, 내장 skill 체계를 소유합니다.
- `src/<plugin-name>-dev/specs/skills/<skill-name>.md`: flat skill spec입니다. 하나의 파일이 해당 skill의 처리 계약을 소유합니다.
- `src/<plugin-name>-dev/specs/skills/<skill-name>/spec.md`: folder-based skill spec의 index입니다.
- `src/<plugin-name>-dev/specs/skills/<skill-name>/intent.md`: folder-based skill spec의 사용자 스펙 의도 기록입니다.
- `src/<plugin-name>-dev/changes/<version>.md`: version-scoped release/change 기록입니다.

## Plugin Spec

`specs/plugin.md`는 최소한 다음을 현재 기준으로 고정합니다.

- 플러그인 목적
- 플러그인 경계와 비목표
- 처리하려는 작업 형태
- 엔트리포인트 또는 대표 표면
- 내장 skill 체계와 각 skill의 역할 요약
- 새 skill을 추가하거나 기존 skill 책임을 바꿀 때 유지해야 할 확장 원칙

plugin spec은 bundle 목적, 경계, 라우팅 표면, skill composition을 소유합니다.
개별 skill의 입력 형태, 판단 기준, 출력 계약, 가드레일처럼 skill 자체의 처리 계약은 가능한 한 skill spec으로 내립니다.

## Skill Spec

개별 skill의 처리 계약이 중요해지면 `specs/skills/<skill-name>.md`로 분리하는 방식을 우선 검토합니다.

flat skill spec은 다음을 소유합니다.

- `사용자 스펙 의도`
- 목적
- 경계
- 처리하려는 작업 형태
- 엔트리포인트 또는 대표 표면
- 핵심 처리 계약
- 필요한 판단 규칙
- 검토 질문
- 독립성 원칙
- 확장 원칙

skill spec에는 해당 skill이 실제로 무엇을 판단하고 어떤 계약으로 동작하는지 적습니다.
실제 판단을 수행하는 skill이라면 `핵심 처리 계약` 뒤에 실제 문서에서 쓰는 규칙 섹션 이름을 그대로 두고, 필요할 때만 섹션 수를 줄이거나 늘립니다.

## Folder-Based Skill Spec

skill spec이 커져 하나의 파일 안에서 index와 세부 계약이 반복되면 folder-based spec으로 전환할 수 있습니다.

기본 구조:

- `specs/skills/<skill-name>/intent.md`
- `specs/skills/<skill-name>/spec.md`
- `specs/skills/<skill-name>/<responsibility>.md`

`intent.md`는 해당 folder-based skill spec의 사용자 스펙 의도 기록을 소유합니다.
`사용자 스펙 의도`를 `spec.md`와 child spec에 반복하지 않습니다.

`spec.md`는 index로서 다음만 소유합니다.

- skill 목적
- 경계
- 대표 표면
- `intent.md` 위치
- sub-spec map
- 확장 원칙

child spec은 lifecycle, routing, verification처럼 자기 세부 처리 계약만 소유합니다.
child spec은 상위 의도 기록을 반복하지 않습니다.

플러그인 스펙에서 folder-based skill을 언급할 때는 목적과 관계를 요약하고, 상세 기준은 `spec.md` 또는 해당 child spec으로 연결합니다.

## 사용자 스펙 의도

plugin spec과 flat skill spec의 `사용자 스펙 의도`는 파일 맨 위에 둡니다.
folder-based skill spec은 같은 폴더의 `intent.md`에 `사용자 스펙 의도`를 둡니다.

`사용자 스펙 의도`는 정규 스펙 본문으로 옮기기 전의 입력 기록입니다.
사용자 요청문을 시간순으로 기록합니다.
질문 도구를 사용했다면 질문과 답변도 함께 기록합니다.

기본 형식:

```md
- <사용자 메시지>
  - <질문1>
    - <답변1>[선택여부, 추가 메시지]
  - <질문2>
    - <답변2>[선택여부, 추가 메시지]
- <사용자 메시지>
```

답변에 선택지가 있었다면 어떤 선택지가 선택됐는지 기록하고, 사용자의 추가 메시지나 free-form 입력이 있으면 같은 답변 줄에 함께 남깁니다.
질문이 없었던 사용자 메시지는 사용자 메시지만 기록하고 하위 질문 항목을 억지로 만들지 않습니다.

`사용자 스펙 의도` 아래의 정규 스펙 본문은 현재 유지돼야 하는 계약만 적고, 이번 세션의 변경 과정이나 migration 맥락을 암시하지 않습니다.
정규 스펙 본문에서는 "분리한다", "제거한다", "추가한다", "이전 질문을 바꾼다"처럼 변경 작업을 설명하는 표현보다 "소유한다", "라우팅한다", "유지한다", "위임한다"처럼 현재 동작과 책임을 설명하는 표현을 우선합니다.
예외적으로 migration 자체가 해당 스펙의 영구 처리 대상이면, 그 이유와 범위를 명시하고 현재 계약과 변경 이력을 구분합니다.

## Change Spec

version-scoped release나 migration 범위는 `src/<plugin-name>-dev/changes/<version>.md` change spec으로 고정할 수 있습니다.

change spec은 릴리즈노트 성격의 변경 기록입니다.
사용자와의 질문/답변 히스토리를 보관하지 않습니다.
`사용자 스펙 의도` 형식을 적용하지 않습니다.
change spec에 필요한 사용자 결정은 릴리즈노트에 필요한 변경 근거로만 짧게 남깁니다.

change spec은 다음을 기록합니다.

- 변경 목적
- 포함 변경
- 비목표
- 호환성/마이그레이션
- 검증 기준

change spec은 먼저 짧은 `변경사항 요약` 리스트를 보여주고, 각 변경의 목적, 범위, 관련 표면, 검증 같은 상세 내용은 별도 `변경 상세` 섹션에 둡니다.

change spec은 영구 normative plugin/skill contract를 대체하지 않습니다.
change spec에서 확정된 지속 규칙은 `specs/plugin.md`, 관련 `specs/skills/*.md`, `specs/skills/<skill-name>/spec.md`, guide skill, README 같은 소유 표면으로 승격합니다.
일회성 release history, migration 배경, 제거 범위, 검증 결과는 change spec에 남기고 일반 skill spec으로 옮기지 않습니다.
release surface에는 `changes/`를 포함하지 않습니다.

## 템플릿

새 플러그인, skill spec, change spec을 시작할 때는 아래 템플릿을 기본 시작점으로 사용합니다.

- plugin spec starter: `docs/templates/plugin-spec.md`
- skill spec starter: `docs/templates/skill-spec.md`
- change spec starter: `docs/templates/change-spec.md`

템플릿은 현재 저장소에 남아 있는 유지 기준 spec과 동일한 방향으로 갱신합니다.
folder-based skill spec을 만들 때는 `skill-spec.md`를 그대로 모든 child spec에 복제하지 말고, `intent.md`, `spec.md`, child spec의 책임을 나눕니다.

## 변경 Workflow

- 플러그인 수정 요청은 먼저 해당 dev source의 spec을 확인하고 필요한 spec 변경을 반영합니다.
- 실제 skill 본문 변경은 spec 변경 또는 spec 확인 이후에 진행합니다.
- skill 책임, guide 라우팅, plugin boundary가 바뀌면 관련 skill spec, plugin spec, guide skill, upstream/downstream plugin surface를 같은 변경 단위에서 함께 점검합니다.
- spec 없는 skill 추가를 기본 경로로 두지 않습니다.
- 반복되는 검토 질문, 예시, decision rule이 생기면 skill spec 안에서 소유할지 별도 reference 문서로 뺄지 의도적으로 결정합니다.
