## 사용자 스펙 의도

- 검색, browsing, 자료 회수, 코드·문서 조사, 구현, 변환, 검토, 검증을 모두 work로 취급하고, 조사와 실행이 섞인 cross-domain workstream orchestration을 소유합니다.
- 위임은 독립성, 명확한 계약, 병렬 이점, shared-state 안전, 메인 에이전트의 검증·통합 가능성을 모두 충족할 때만 허용합니다.
- 기본 worker는 `gpt-5.6-terra` `xhigh`이며 역할별 task packet으로 구분합니다. Luna형 작업은 `PROCESS_STRUCTURED` Terra worker로 통합하고, Luna route를 두지 않습니다.
- `gpt-5.6-sol` `xhigh`는 제한된 `FRONTIER_JUDGMENT` 조건에서만 사용합니다.
- 순수 근거 추적 조사 보고서와 일반 software-engineering-only 3단 dispatch를 자동 트리거 경계에서 분리하고, sibling plugin 없이도 독립적으로 동작합니다.
- 하나의 cohesive lifecycle skill과 progressive-disclosure references를 사용합니다.

---

# Workstream Orchestrator Dev 플러그인 스펙

## 플러그인 목적

`workstream-orchestrator-dev`는 cross-domain 또는 mixed investigation-and-action 목표를 독립 workstream으로 분해하고, 안전한 bounded subagent 위임부터 근거 검증, 충돌 해결, 전체 결과 통합까지 하나의 수명주기로 조율하는 instruction-only 플러그인입니다.

## 플러그인 경계와 비목표

- 포함:
  - 검색, browsing, source retrieval, 코드·문서 inspection, implementation, transformation, review, verification의 orchestration
  - `DIRECT` 또는 `DISPATCH` gate, workstream·dependency graph, shared-state 안전
  - 역할·모델 routing, bounded task/result contract, spawn과 terminal result normalization
  - evidence·ownership 검사, 충돌 해결, whole-result verification, 최종 응답
- 제외:
  - 모든 복잡한 작업의 자동 위임
  - 순수 조사 보고서의 source methodology, citation ledger, evidence-report contract
  - 일반 software-engineering-only 요청의 정밀 implicit routing과 3단 dispatch lifecycle
  - sibling plugin의 설치나 숨은 context를 요구하는 실행
  - MCP, app, hook, script, executable runtime dependency
  - 사용자가 요청하지 않은 외부 상태 변경, commit, push, PR, release

## 처리하려는 작업 형태

- 시장·정책·사용자 조사와 prototype·문서·코드 산출물 제작이 섞인 작업
- 자료 수집, schema-bound 처리, disjoint implementation, 독립 review lens가 하나의 목표를 지원하는 작업
- 여러 도메인 또는 artifact 유형의 결과를 공통 성공 기준으로 검증·통합해야 하는 작업
- 명시적인 subagent 요청 중 전체 dispatch gate를 통과한 작업

## 대표 표면

- 대표 스펙: `specs/plugin.md`
- skill 상세 스펙 위치: `specs/skills/orchestrate-workstreams.md`
- 사용자 안내: `README.md`
- manifest: `.codex-plugin/plugin.json`
- 호출 식별자: `$workstream-orchestrator-dev:orchestrate-workstreams`

## 내장 skill 체계

- `orchestrate-workstreams`:
  - 목표 고정부터 위임 판정, 실행, 결과 정규화, 검증·통합, 최종 응답까지 하나의 cohesive lifecycle을 소유합니다.
  - runtime references는 delegation safety, task/result contracts, model routing, boundary examples만 progressive disclosure로 제공합니다.
  - spec: `specs/skills/orchestrate-workstreams.md`

## Plugin Usage 계약

- manifest, README, plugin spec은 자동·명시 호출 기준과 인접 plugin 경계를 소유합니다.
- pure evidence-report research는 자동 호출하지 않습니다.
- 명시적 orchestration이 없는 ordinary software-engineering-only 요청은 자동 호출하지 않습니다.
- explicit subagent 요청은 skill gate에 진입하지만 spawn을 보장하지 않습니다.
- sibling plugin이 설치되어 있으면 각 소유 경계를 존중하고, 없어도 이 플러그인의 명시 호출과 runtime 계약은 독립적으로 동작합니다.

## SDD 운영 원칙

- plugin spec은 bundle 목적, 비목표, 사용 기준, 단일 skill 구성을 소유합니다.
- skill spec은 data flow, gate, 안전, routing, task/result contract, integration을 소유합니다.
- skill spec이 바뀌면 runtime skill folder를 현재 spec 기준으로 처음부터 재작성합니다.
- release surface는 build command로만 만들며 `specs/`와 `changes/`를 포함하지 않습니다.

## 확장 원칙

- lifecycle을 나누는 새 skill보다 현재 cohesive skill과 one-level reference를 우선합니다.
- 새 role은 기존 역할로 표현할 수 없고 task/result contract가 달라질 때만 추가합니다.
- 새 모델 route는 모델명이 아니라 판단 책임과 검증 가능성으로 정의합니다.
- pure research methodology나 software-only 3단 lifecycle로 경계를 넓히지 않습니다.

## 현재 구조 메모

- 하나의 instruction-only skill과 네 개의 concise runtime reference를 사용합니다.
- trigger validation을 위한 compact scenario fixture를 skill에 둘 수 있습니다.
- scripts, assets, MCP, app, hook은 제공하지 않습니다.
