# Advance Codex 플러그인 스펙

## 목적

`advance-codex`는 Codex의 기본 표면을 설계, 보강, 정리하기 위한 플러그인입니다.
핵심 역할은 먼저 지금 다루는 Codex 표면이 무엇인지 분류하고, 그다음 더 명확한 경계, 패키징 규칙, tool-use policy 분리, custom agent 가이드, session continuity 관리, change finalization 규율을 붙여 안정적인 형태로 정리하는 것입니다.

## 경계

- 포함:
  - skill 설계 및 수정
  - 재사용 가능한 instruction 또는 skill의 empirical evaluation workflow
  - tool-use guidance 설계 및 분리
  - plugin bundle 설계 및 패키징 가이드
  - custom subagent 설계 및 사용 가이드
  - session-scoped record 관리와 session continuity 보강
  - 직접 만든 변경을 task-scoped commit으로 정리하는 change finalization 보강
- 제외:
  - 일반적인 제품 구현 workflow
  - Codex 표면 보강과 무관한 도메인 실행 가이드
  - `advance-codex` 맥락 없이 공용 유틸리티를 쌓는 일

## 진입 표면

- 대표 엔트리포인트: `advance-codex-guide`
- 핵심 분기: 지금 다루는 대상이 skill, tool-use guidance, plugin, subagent, session surface, commit surface 중 무엇인지 먼저 분류한다

## 스킬 구성

- `advance-codex-guide`: 표면의 형태를 분류하고 적절한 내장 skill로 라우팅한다
- `skill-creator`: canonical `skill-creator`와 함께 읽는 확장으로서 bounded skill 설계와 plugin 내부 skill packaging 규칙을 강화한다
- `empirical-prompt-tuning`: reusable instruction, skill, AGENTS section, task prompt를 fresh subagent와 fixed scenarios로 반복 평가하고 개선하는 empirical tuning workflow를 제공한다
- `tool-use-guide`: 도메인 skill 바깥으로 재사용 가능한 tool selection, sequencing, ask-vs-infer, escalation policy를 분리한다
- `plugin-creator`: canonical `plugin-creator`와 함께 읽는 확장으로서 top-down plugin 설계, bundled skill coherence, `<plugin>-guide` 기대사항을 강화한다
- `subagent-creator`: reusable custom agent role, TOML 형태, usage guidance를 정의한다
- `session-manager`: `.agents/sessions/<uuid>/` 아래의 session record, change record, retrospective record를 통해 Codex 세션 연속성과 전달 기록을 관리한다
- `git-committer`: 직접 만든 변경을 검토, 분리, 검증하고 task-scoped commit으로 확정하는 규율을 제공한다

## 확장 원칙

- 새 skill은 기존 표면과 분명히 다른 Codex 표면이나 보강 관심사가 있을 때만 추가한다.
- plugin-level routing은 `advance-codex-guide`에 두고 creator skill들에 흩뿌리지 않는다.
- 한 번 안정적인 관심사로 분리된 tool policy를 다시 도메인 skill 안으로 밀어넣지 않는다.
- empirical evaluation workflow는 author intuition이 아니라 fresh-executor evidence를 요구하는 별도 관심사일 때만 추가한다.
- 세션처럼 Codex의 기본 동작을 안정화하는 표면은 이 플러그인 안에서 다룰 수 있다.
- commit처럼 Codex 출력의 최종 확정을 안정화하는 표면도 이 플러그인 안에서 다룰 수 있다.
- 이 플러그인은 execution workflow보다 Codex surface augmentation에 집중한다.

## 현재 의도 점검

- 현재 플러그인 표면은 일관적이어야 한다. 포함된 skill은 모두 제품 실행이 아니라 Codex 표면 보강에 관한 것이어야 한다.
- 현재의 주요 리스크는 일반 workflow 설계나 무관한 편의 기능 쪽으로 범위가 흐려지는 것이다.
