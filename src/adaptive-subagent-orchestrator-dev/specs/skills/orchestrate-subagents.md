## 사용자 스펙 의도

- adaptive subagent 플러그인의 스킬을 세분화해서 오케스트레이션의 시작은 아주 가벼운 스킬로 만들려고 하는데, 어떻게 풀어볼까?
  - 세분화 축을 어느 구조로 고정할까요?
    - `3단 수명주기`를 선택했습니다.
  - 새 경량 진입 스킬 이름은 무엇으로 고정할까요?
    - `orchestrate-subagents`를 선택했습니다.
  - 세 스킬의 암묵 호출 정책은 어떻게 둘까요?
    - `진입만 implicit`을 선택했습니다.
  - 경량성과 라우팅 기준은 어떻게 고정할까요?
    - `본문 200단어 이하`, `references 없음`, `정밀도 우선`을 선택했습니다.

---

# orchestrate-subagents 스킬 스펙

## 목적

`orchestrate-subagents`는 subagent orchestration의 유일한 implicit entry로서 요청을 저비용으로 `DIRECT` 또는 `DISPATCH`에 라우팅합니다.

## 경계

- 포함:
  - explicit subagent 요청 감지
  - 명백한 독립 workstream 두 개 이상 감지
  - 사용자 금지 조건 우선 적용
  - 최소 `RouteDecision` 생성
- 제외:
  - 전체 delegation gate
  - execution mode, agent 수, 역할, ownership 확정
  - task packet 작성과 subagent 생성
  - 결과 대기, evidence 검증, 최종 통합

## 처리하려는 작업 형태

- 사용자가 subagent, delegation, parallel agent를 명시한 요청
- 서로 독립적으로 시작 가능한 모듈, 관점, 테스트 묶음, 플랫폼, 옵션이 두 개 이상 명백한 요청
- broad engineering keyword가 있지만 독립 작업 흐름은 불명확한 요청의 저비용 DIRECT 판정

## 엔트리포인트 / 대표 표면

- 대표 표면: `skills/orchestrate-subagents/SKILL.md`
- 호출 방식: `$adaptive-subagent-orchestrator-dev:orchestrate-subagents` 또는 정밀도 우선 implicit invocation
- implicit policy: `allow_implicit_invocation: true`

## 핵심 처리 계약

- 사용자가 subagent 사용을 금지하면 항상 `DIRECT`입니다.
- 명시적 subagent 요청은 `DISPATCH`로 라우팅하되, 실제 생성 여부는 dispatcher가 결정합니다.
- implicit 요청은 의미 있고 명백한 독립 workstream이 두 개 이상일 때만 `DISPATCH`합니다.
- complexity, 많은 파일, implementation, review, testing, debugging 표현만으로는 `DISPATCH`하지 않습니다.
- 애매한 요청은 정밀도 우선으로 `DIRECT` 처리합니다.
- `DIRECT`이면 다른 orchestration skill을 읽거나 orchestration 과정을 설명하지 않습니다.

## RouteDecision 계약

`DISPATCH` handoff는 현재 대화 안의 구조화된 값으로 다음을 포함합니다.

- `route`: `DISPATCH`
- `trigger_basis`: explicit request 또는 clear independent workstreams
- `goal`: 사용자 최종 목표
- `candidate_workstreams`: 독립 후보 목록
- `user_constraints`: 금지, 범위, 권한, 완료 조건
- `shared_state_flags`: 파일, schema, config, runtime, data 공유 위험

`DIRECT`는 route와 짧은 근거만 유지하고 downstream handoff를 만들지 않습니다.

## Context 예산

- runtime `SKILL.md` body는 200단어 이하입니다.
- frontmatter description은 80단어 이하로 핵심 trigger와 제외 기준을 앞에 둡니다.
- runtime `references/`를 만들지 않습니다.
- passive token은 explicit subagent 표현과 독립 workstream 표현으로 좁힙니다.

## 검토 질문

- 사용자가 subagent를 금지했는가?
- explicit subagent 요청인가?
- implicit 요청이라면 독립 workstream 두 개 이상이 명백한가?
- generic engineering 표현만으로 과활성화하고 있지 않은가?
- handoff가 dispatcher 판단에 필요한 최소 사실만 포함하는가?

## 독립성 원칙

- 이 skill은 직접 호출과 implicit 호출에서 독립적으로 `RouteDecision`을 만들 수 있어야 합니다.
- `DISPATCH` 이후 실행은 sibling skill에 명시적으로 handoff하지만, 숨은 sibling context를 판단 입력으로 사용하지 않습니다.

## 확장 원칙

- 새 trigger는 정밀도 우선 allowlist를 넓힐 명확한 evidence가 있을 때만 추가합니다.
- 상세 gate, mode, 안전 규칙을 entry로 올리지 않습니다.
- runtime example을 추가하지 않고 trigger 시나리오는 spec과 forward test가 소유합니다.
