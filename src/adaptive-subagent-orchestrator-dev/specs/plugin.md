# Adaptive Subagent Orchestrator Dev 플러그인 스펙

## 사용자 스펙 의도

- 사용자는 `adaptive-subagent-orchestrator`라는 Codex skill을 만들거나 업데이트하길 원합니다.
- 이 skill은 사용자가 서브에이전트나 병렬 처리를 직접 언급하지 않아도, 복잡한 소프트웨어 엔지니어링 요청에서 독립 작업 흐름이 두 개 이상이고 병렬 위임 이점이 있을 때 최소한의 서브에이전트를 자동 생성하고 조율해야 합니다.
- skill 활성화 자체를 명시적인 위임 평가 및 생성 지시로 취급해야 하며, 위임 조건이 충족되지 않으면 메인 에이전트가 직접 처리해야 합니다.
- instruction-only skill로 만들고, `scripts/`와 `assets/`는 생성하지 않아야 합니다.
- `agents/openai.yaml`에는 `allow_implicit_invocation: true`와 명시적 spawn 지시가 있는 `default_prompt`가 있어야 합니다.
- 사용자가 준 `$HOME/.agents/skills` 위치 지시는 이 저장소의 마켓플레이스 등록 요구에 맞춰 dev plugin source와 release plugin surface로 보정해야 합니다.

---

## 플러그인 목적

`adaptive-subagent-orchestrator-dev`는 software-engineering 작업의 병렬 위임 가능성을 판단하고, 가치가 있을 때만 bounded subagent를 최소 수로 생성하도록 안내하는 플러그인입니다.

핵심 책임은 task decomposition, delegation gate, execution mode selection, subagent task contract, evidence validation, final integration을 단일 skill로 제공하는 것입니다.

## 플러그인 경계와 비목표

- 포함:
  - 암묵적 또는 명시적 subagent orchestration skill 제공
  - 독립 작업 흐름 판단
  - DIRECT, PARALLEL_READ, PARALLEL_WRITE 선택 기준
  - subagent task packet과 반환 형식
  - evidence 검증과 최종 통합 규칙
  - implicit invocation metadata 제공
- 제외:
  - 모든 개발 요청에 대한 무조건 delegation
  - 외부 설정, 사용자 설정, Codex 설정 변경
  - 커스텀 subagent role, 모델명, 추론 강도, 런타임 설정 하드코딩
  - deterministic script나 asset 제공
  - sibling plugin 또는 hidden context에 의존하는 사용 방식

## 처리하려는 작업 형태

- 다중 모듈, 패키지, 서비스, 계층을 각각 조사하는 작업
- 보안, 정확성, 성능, 테스트 같은 독립 관점의 cross-cutting review
- 독립 테스트 묶음별 실패 원인 조사와 최소 수정
- 플랫폼, 런타임, 설계 대안, 기술 대안 비교
- 대규모 migration 영향 범위 분석
- 서로 다른 로그, 코드 영역, 실행 경로를 나눠 조사하는 작업
- disjoint write scope가 명확하고 공통 계약이 확정된 제한적 병렬 구현

## 대표 표면

- 대표 스펙: `adaptive-subagent-orchestrator-dev/specs/plugin.md`
- skill 상세 스펙 위치: `adaptive-subagent-orchestrator-dev/specs/skills/adaptive-subagent-orchestrator.md`
- runtime skill: `adaptive-subagent-orchestrator-dev/skills/adaptive-subagent-orchestrator/SKILL.md`

## 내장 skill 체계

- `adaptive-subagent-orchestrator`:
  - 복잡한 engineering task를 독립 workstream으로 평가하고, 위임 조건이 통과할 때만 최소한의 bounded subagent를 생성하도록 안내합니다.
  - spec: `adaptive-subagent-orchestrator-dev/specs/skills/adaptive-subagent-orchestrator.md`

## SDD 운영 원칙

- plugin spec은 플러그인 경계와 bundle 사용 표면을 소유합니다.
- skill의 판단 계약, task contract, 반환 형식, 예시는 skill spec과 runtime references가 소유합니다.
- skill spec이 바뀌면 runtime skill folder를 현재 spec 기준으로 다시 작성합니다.
- release surface는 build command 산출물만 사용하며 `specs/`와 `changes/`를 포함하지 않습니다.

## 현재 구조 메모

- 이 플러그인은 단일 instruction-only skill bundle입니다.
- runtime surface는 `SKILL.md`, `agents/openai.yaml`, `references/`만 포함합니다.
- scripts, assets, MCP, app integration은 제공하지 않습니다.
