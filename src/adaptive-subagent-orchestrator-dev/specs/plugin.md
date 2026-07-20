# Adaptive Subagent Orchestrator Dev 플러그인 스펙

## 사용자 스펙 의도

- 사용자는 `adaptive-subagent-orchestrator`라는 Codex skill을 만들거나 업데이트하길 원합니다.
- 이 skill은 사용자가 서브에이전트나 병렬 처리를 직접 언급하지 않아도, 복잡한 소프트웨어 엔지니어링 요청에서 독립 작업 흐름이 두 개 이상이고 병렬 위임 이점이 있을 때 최소한의 서브에이전트를 자동 생성하고 조율해야 합니다.
- skill 활성화 자체를 명시적인 위임 평가 및 생성 지시로 취급해야 하며, 위임 조건이 충족되지 않으면 메인 에이전트가 직접 처리해야 합니다.
- instruction-only skill로 만들고, `scripts/`와 `assets/`는 생성하지 않아야 합니다.
- 초기 단일-skill 요구에서는 `agents/openai.yaml`에 `allow_implicit_invocation: true`와 명시적 spawn 지시가 있는 `default_prompt`를 요구했습니다. 아래 세분화 결정이 이를 대체하며, entry만 implicit이고 실제 spawn 지시는 dispatcher가 소유합니다.
- 사용자가 준 `$HOME/.agents/skills` 위치 지시는 이 저장소의 마켓플레이스 등록 요구에 맞춰 dev plugin source와 release plugin surface로 보정해야 합니다.
- adaptive subagent 플러그인의 스킬을 세분화해서 오케스트레이션의 시작은 아주 가벼운 스킬로 만들려고 하는데, 어떻게 풀어볼까?
  - 세분화 축을 어느 구조로 고정할까요?
    - `3단 수명주기`를 선택했습니다: 경량 진입, 위임·생성, 결과 검증·통합.
  - 현재 스킬 식별자 `adaptive-subagent-orchestrator`는 어떻게 처리할까요?
    - `새 이름으로 교체`를 선택했습니다. 호환 shim은 유지하지 않습니다.
  - 새 경량 진입 스킬 이름은 무엇으로 고정할까요?
    - `orchestrate-subagents`를 선택했습니다.
  - 세 스킬의 암묵 호출 정책은 어떻게 둘까요?
    - `진입만 implicit`을 선택했습니다.
  - `orchestrate-subagents`의 경량성을 어떤 검증 기준으로 고정할까요?
    - `200단어·무참조`를 선택했습니다.
  - 후속 두 스킬의 명시적 직접 호출도 정식 지원할까요?
    - `입력 게이트로 지원`을 선택했습니다.
  - 암묵 호출의 경계에서는 정밀도와 포착률 중 어디에 무게를 둘까요?
    - `정밀도 우선`을 선택했습니다.

---

## 플러그인 목적

`adaptive-subagent-orchestrator-dev`는 software-engineering 작업을 저비용으로 선별하고, 병렬 위임 가치가 확인될 때만 bounded subagent를 생성하고 결과를 검증·통합하는 instruction-only 플러그인입니다.

핵심 책임은 orchestration entry, delegation dispatch, result integration을 세 개의 좁은 skill로 제공하는 것입니다.

## 플러그인 경계와 비목표

- 포함:
  - 정밀도 우선 implicit orchestration entry
  - DIRECT, PARALLEL_READ, PARALLEL_WRITE 판단과 bounded subagent 생성
  - subagent task packet과 dispatch manifest
  - evidence 검증, 충돌 해결, 전체 결과 통합
  - 입력 게이트를 갖춘 focused skill 직접 호출
- 제외:
  - 모든 개발 요청에 대한 무조건 delegation
  - 별도 planner 또는 mode별 read/write skill
  - 외부 설정, 사용자 설정, Codex 설정 변경
  - custom subagent role, 모델명, reasoning effort 하드코딩
  - deterministic script, asset, MCP, app integration
  - 숨은 sibling context에 의존하는 실행

## 처리하려는 작업 형태

- 다중 모듈, 패키지, 서비스, 계층을 각각 조사하는 작업
- 보안, 정확성, 성능, 테스트 같은 독립 관점의 cross-cutting review
- 독립 테스트 묶음별 실패 원인 조사와 최소 수정
- 플랫폼, 런타임, 설계 대안, 기술 대안 비교
- 대규모 migration 영향 범위 분석
- 서로 다른 로그, 코드 영역, 실행 경로를 나눠 조사하는 작업
- disjoint write scope와 공통 계약이 확정된 제한적 병렬 구현

## 대표 표면

- 대표 스펙: `adaptive-subagent-orchestrator-dev/specs/plugin.md`
- skill spec: `adaptive-subagent-orchestrator-dev/specs/skills/*.md`
- runtime skills: `adaptive-subagent-orchestrator-dev/skills/*/SKILL.md`
- 기본 시작점: `$adaptive-subagent-orchestrator-dev:orchestrate-subagents`

## 내장 skill 체계

- `orchestrate-subagents`:
  - explicit subagent 요청 또는 명백한 독립 workstream을 저비용으로 선별하여 `DIRECT` 또는 `DISPATCH`로 라우팅합니다.
  - 유일하게 implicit invocation을 허용합니다.
  - spec: `specs/skills/orchestrate-subagents.md`
- `dispatch-subagents`:
  - 전체 delegation gate, execution mode, workstream, ownership, task packet을 확정하고 최소 bounded subagent를 실제 생성합니다.
  - raw request 또는 유효한 route handoff로 직접 호출할 수 있습니다.
  - spec: `specs/skills/dispatch-subagents.md`
- `integrate-subagent-results`:
  - dispatch manifest와 active results를 입력으로 모든 필수 결과를 기다리고 evidence를 검증하여 최종 결과를 통합합니다.
  - 완전한 dispatch 입력이 있을 때 직접 호출할 수 있습니다.
  - spec: `specs/skills/integrate-subagent-results.md`

## Lifecycle 계약

1. `orchestrate-subagents`는 `RouteDecision`을 만듭니다.
2. `DIRECT`이면 다른 orchestration skill을 읽지 않고 main agent가 직접 처리합니다.
3. `DISPATCH`이면 `dispatch-subagents`가 전체 gate를 다시 확인하고, 통과할 때만 subagent를 생성하여 `DispatchManifest`를 만듭니다.
4. 생성이 발생하면 main agent가 `integrate-subagent-results`를 적용하여 필수 결과를 검증·통합하고 lifecycle follow-up 사용 여부를 manifest에 유지합니다.
5. PARALLEL_READ 이후 write가 필요하면 write scope를 명시해 `dispatch-subagents`를 다시 거칩니다.

## Plugin Usage 계약

- README, manifest `defaultPrompt`, plugin spec이 세 skill의 시작 기준과 lifecycle 순서를 소유합니다.
- skill body는 자기 owned job과 입력·출력만 설명합니다.
- `orchestrate-subagents`의 broad keyword만으로 implicit dispatch하지 않습니다.
- focused skill은 `allow_implicit_invocation: false`를 사용하고 입력 게이트를 통과해야 합니다.

## SDD 운영 원칙

- plugin spec은 bundle 목적, skill composition, lifecycle, usage surface를 소유합니다.
- 각 skill spec은 입력, 판단, 출력, 중단 조건, 독립성 계약을 소유합니다.
- skill spec이 바뀌면 해당 runtime skill folder를 현재 spec 기준으로 다시 작성합니다.
- release surface는 build command 산출물만 사용하며 `specs/`와 `changes/`를 포함하지 않습니다.

## 현재 구조 메모

- 세 instruction-only skill이 lifecycle을 나눠 소유합니다.
- 경량 entry는 runtime reference를 갖지 않습니다.
- dispatcher와 integrator만 자기 owned job에 필요한 runtime references를 갖습니다.
- scripts, assets, MCP, app integration은 제공하지 않습니다.
