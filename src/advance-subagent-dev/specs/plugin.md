## 사용자 스펙 의도

- 서브에이전트 관련 심화된 방법을 소유하는 플러그인으로 봅니다.

---

# Advance Subagent Dev 플러그인 스펙

## 플러그인 목적

`advance-subagent-dev`는 서브에이전트를 활용한 근거 중심 조사와 독립 workstream 위임·검증·통합의 심화 실행 방법을 제공하는 instruction-only 플러그인입니다.

## 플러그인 경계와 비목표

- 포함:
  - 여러 출처의 원문, 근거, 반대 근거, 불확실성, 인용을 추적하는 독립 조사 보고서
  - software-engineering 또는 cross-domain 목표의 독립 workstream 분해와 bounded 위임
  - 역할·모델 routing, task/result contract, shared-state 안전, 결과 검증·통합
- 제외:
  - 모든 복잡한 작업의 자동 위임
  - 단순 조회, 단일 문서 요약, 근거 없는 아이디어 발산
  - sibling skill의 숨은 context나 실행을 요구하는 처리
  - MCP, app, hook, script, executable runtime dependency
  - 사용자가 요청하지 않은 외부 상태 변경, commit, push, PR, release

## 처리하려는 작업 형태

- 시장·경쟁·규제·기술·문헌·기업 사실을 여러 출처로 확인하는 조사 보고서
- 근거와 반대 근거, 최신성, 불확실성, 인용을 감사할 수 있어야 하는 팩트체크와 실사
- 조사와 prototype·문서·코드·데이터 변환이 섞인 cross-domain workstream
- 독립 모듈·서비스·실행 경로·테스트 묶음의 software-engineering 조사와 구현
- correctness·security·performance·quality 같은 독립 review lens

## 대표 표면

- 매니페스트: `.codex-plugin/plugin.json`
- 사용자 안내: `README.md`
- 플러그인 스펙: `specs/plugin.md`
- skill 상세 스펙:
  - `specs/skills/deep-research.md`
  - `specs/skills/orchestrate-workstreams.md`
- 공개 marketplace: `.agents/plugins/marketplace.json`의 `./advance-subagent`
- 개발 호출 식별자:
  - `$advance-subagent-dev:deep-research`
  - `$advance-subagent-dev:orchestrate-workstreams`
- 공개 호출 식별자:
  - `$advance-subagent:deep-research`
  - `$advance-subagent:orchestrate-workstreams`

## 내장 skill 체계

- `deep-research`:
  - 조사 범위 설정부터 원문 확보, 증거 대조, 인용·불확실성 감사, 최종 보고서까지의 실행 워크플로를 소유합니다.
  - spec: `specs/skills/deep-research.md`
- `orchestrate-workstreams`:
  - 목표 고정부터 위임 판정, 실행, 결과 정규화, 검증·통합, 최종 응답까지의 cohesive lifecycle을 소유합니다.
  - spec: `specs/skills/orchestrate-workstreams.md`

## Plugin Usage 계약

- 조사 자체가 산출물이고 출처·증거·인용 계약이 필요하면 `deep-research`를 선택합니다.
- software-engineering 또는 cross-domain 요청에서 독립 workstream이 둘 이상 명백하면 `orchestrate-workstreams`의 dispatch gate를 적용합니다.
- 명시 호출은 각 skill의 gate와 제외 조건을 우회하지 않습니다.
- 두 skill의 기존 trigger, routing, 산출물 계약은 서로 합치지 않습니다.
- 각 skill은 sibling skill 없이 독립 실행 가능해야 합니다.
- plugin manifest, README, plugin spec이 skill 선택 기준과 namespace를 소유합니다.

## SDD 운영 원칙

- plugin spec은 bundle 목적, 비목표, 사용 기준, skill composition을 소유합니다.
- 개별 조사·위임·검증 계약은 각 skill spec이 소유합니다.
- skill spec이 바뀌면 해당 runtime skill folder를 현재 spec 기준으로 처음부터 재작성합니다.
- release surface는 build command로만 만들며 `specs/`와 `changes/`를 포함하지 않습니다.

## 확장 원칙

- 새 skill은 서브에이전트 활용의 별도 사용자 산출물과 독립 처리 계약이 있을 때만 추가합니다.
- 한 skill의 routing, result field, evidence contract를 sibling에 암묵적으로 적용하지 않습니다.
- 새 모델 route는 모델명이 아니라 판단 책임과 검증 가능성으로 정의합니다.
- 실제 제공되는 MCP, app, hook, script만 runtime surface로 선언합니다.
