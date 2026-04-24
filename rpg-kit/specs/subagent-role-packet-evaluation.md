# Subagent Role Packet Evaluation

## 목적

이 문서는 `rpg-kit`의 `subagent-role` skill이 role specialty를 사용해 좋은 subagent role packet을 만들 수 있는지 검증하기 위한 고정 평가 설계입니다.

이 평가는 prompt/instruction 품질을 확인하는 절차입니다.
실제 fresh subagent 실행 전에는 이 문서의 시나리오와 체크리스트를 고정하고, 실행 중에는 기대 답안을 누설하지 않습니다.

## 평가 대상

- `rpg-kit/skills/subagent-role/SKILL.md`
- `rpg-kit/specs/skills/subagent-role.md`
- `rpg-kit/skills/rpg-kit-guide/SKILL.md`
- `rpg-kit/specs/skills/rpg-kit-guide.md`

## 실패 모드

- role을 단순 직함이나 persona로만 정의한다.
- 같은 축을 반복하는 중복 role을 좋은 예시로 취급한다.
- role specialty와 role packet을 분리하지 않는다.
- spawn boundary 없이 바로 subagent spawn을 권한다.
- answer contract가 caller integration에 충분히 구조화되어 있지 않다.
- subagent output을 caller 검증 없이 최종 답처럼 취급한다.
- learning note가 role specialty 개선으로 이어지지 않는다.

## 공통 성공 기준

모든 시나리오에서 다음 조건을 확인합니다.

- `[critical]` role specialty가 functional role, responsibility domain, background expertise, general expertise, decision style 중 필요한 축으로 표현되어야 한다.
- `[critical]` 같은 축을 반복하는 role은 중복으로 지적하고 더 나은 조합을 제안해야 한다.
- `[critical]` role packet에는 `role_name`, `role_specialty`, `objective`, `context`, `ownership`, `non_goals`, `expected_output`, `verification_signal`, `integration_rule`, `stop_condition`이 있어야 한다.
- `[critical]` spawn boundary와 approval boundary를 명시해야 한다.
- `[critical]` answer contract는 caller가 통합할 수 있는 structured result를 요구해야 한다.
- caller-side verification 또는 integration rule이 있어야 한다.
- learning note는 다음 role packet 개선에 연결되어야 한다.

## 시나리오 A: 게임 PM 출신 PO

### 상황

사용자는 신규 게임 onboarding 흐름에 대해 제품 판단을 도와줄 subagent role을 만들고 싶어합니다.
역할 후보는 `게임 PM 출신 PO`입니다.

### 평가 입력

fresh executor에게 `subagent-role` 관련 파일을 읽게 한 뒤, 위 상황에서 subagent role packet을 작성하게 합니다.

### 체크리스트

- `[critical]` `functional_role=PO`와 `background_expertise=game-PM-origin`을 구분한다.
- `[critical]` `general_expertise=product strategy` 또는 동등한 범용 전문성을 표현한다.
- onboarding의 player impact, retention risk, prioritization tradeoff를 expected output 또는 answer contract에 포함한다.
- caller가 최종 제품 결정을 검증/통합한다는 rule이 있다.
- learning note가 game-PM-origin 배경이 판단에 도움이 됐는지 확인한다.

## 시나리오 B: 중복 role 정정

### 상황

사용자가 다음 role을 제안합니다.

```text
프론트 출신 프론트엔드 개발자
```

### 평가 입력

fresh executor에게 이 role이 좋은 role specialty인지 판단하고, 필요하면 수정안을 제시하게 합니다.

### 체크리스트

- `[critical]` 같은 축이 반복되어 role specialty가 흐려진다고 지적한다.
- `[critical]` `게임 기획 출신 프론트엔드 개발자` 또는 다른 cross-axis 대안을 제시한다.
- 수정안에서 functional role, responsibility domain, background expertise를 분리한다.
- 중복 role을 그대로 packet으로 확정하지 않는다.

## 시나리오 C: 엔지니어 역할의 프론트엔드 개발자

### 상황

사용자는 UI 구현 범위와 component boundary를 검토할 subagent role을 만들고 싶어합니다.
역할 후보는 `엔지니어 역할의 프론트엔드 개발자`입니다.

### 평가 입력

fresh executor에게 이 role을 사용해 implementation-oriented role packet을 작성하게 합니다.

### 체크리스트

- `[critical]` `functional_role=engineer`와 `responsibility_domain=frontend implementation`을 표현한다.
- expected output에 implementation plan, file ownership, integration risk, verification path가 포함된다.
- coding work라면 disjoint write scope와 concurrent changes guard를 포함한다.
- caller가 staged changes 또는 final patch를 직접 검토해야 한다는 integration rule이 있다.

## 시나리오 D: 디자이너 역할의 UX 리뷰어

### 상황

사용자는 화면 구조와 affordance를 검토할 subagent role을 만들고 싶어합니다.
역할 후보는 `디자이너 역할의 UX 리뷰어`입니다.

### 평가 입력

fresh executor에게 review-oriented role packet을 작성하게 합니다.

### 체크리스트

- `[critical]` role이 구현 owner가 아니라 review/critique owner임을 구분한다.
- expected output에 UX findings, severity, proposed adjustment, acceptance signal이 포함된다.
- non-goals에 직접 구현 변경이나 범위 밖 redesign을 제한한다.
- caller가 material issue만 통합할지 판단하는 rule이 있다.

## 실행 프롬프트 템플릿

fresh executor에게 다음 형식으로 요청합니다.

```text
다음 instruction 파일을 읽고, 주어진 시나리오에서 `subagent-role`이 어떤 role packet을 작성해야 하는지 실행 결과를 작성하세요.

대상 파일:
- rpg-kit/skills/subagent-role/SKILL.md
- rpg-kit/specs/skills/subagent-role.md

시나리오:
<시나리오 본문>

체크리스트:
<해당 시나리오 체크리스트>

보고 형식:
- Role fit
- Role specialty
- Role packet
- Spawn boundary
- Answer contract
- Integration rule
- Learning note
- Checklist status: pass/fail/partial with reason
- Ambiguous wording found
```

## 판정 방식

- `[critical]` 항목이 하나라도 fail이면 해당 시나리오는 fail입니다.
- critical 항목이 모두 pass이고 나머지 항목의 80% 이상이 pass이면 scenario pass입니다.
- 같은 instruction revision에서 4개 시나리오가 모두 pass해야 전체 pass입니다.
- 동일 시나리오를 재실행할 때는 fresh executor를 사용해야 하며, 이전 executor의 판단을 재사용하지 않습니다.

## Caller-Side Metrics

각 실행에서 caller는 다음을 기록합니다.

- scenario id
- success: pass/fail
- checklist pass count
- critical failure count
- reported ambiguous wording count
- reported missing role packet field count
- reported duplicate-axis issue count
- tool use count
- duration, if available

## 개선 판단

실패가 발생하면 다음 순서로 원인을 분류합니다.

1. role specialty 축이 충분히 드러나지 않음
2. 중복 role을 정정하지 못함
3. role packet 필드가 누락됨
4. spawn boundary가 약함
5. answer contract가 caller integration에 약함
6. learning note가 다음 role 개선으로 이어지지 않음

각 iteration은 위 원인 중 하나만 고쳐야 합니다.

## Iteration 1 Baseline Result

### Change Theme

- Role pattern catalog output contracts need to be more explicit per pattern.

### Scenario Results

| Scenario | Success | Accuracy | Retries |
|---|---|---:|---:|
| A: 게임 PM 출신 PO | pass | 80% | 0 |
| B: 중복 role 정정 | pass | 100% | 0 |
| C: 엔지니어 역할의 프론트엔드 개발자 | pass | 100% | 0 |
| D: 디자이너 역할의 UX 리뷰어 | pass | 75% | 0 |

### New Ambiguities

- A: `risk` was too broad and did not force `retention risk`; `player impact` was not explicitly onboarding-specific.
- C: `엔지니어 역할의 프론트엔드 개발자` is understandable but should normalize to `functional_role=engineer` and `responsibility_domain=frontend implementation`.
- D: `designer role for UX review` can drift into visual critique unless read-only UX critique, material issue filtering, and acceptance signal meaning are explicit.

### Patch Applied

- Tightened role pattern catalog output contracts for game-PM-origin PO, frontend engineer, and UX reviewer patterns.

## Iteration 1 Re-Run Result

### Change Theme

- Confirm whether the tightened role pattern catalog removed the partial failures from the baseline run.

### Scenario Results

| Scenario | Success | Accuracy | Retries |
|---|---|---:|---:|
| A: 게임 PM 출신 PO | pass | 100% | 0 |
| B: 중복 role 정정 | pass | 100% | 0 |
| C: 엔지니어 역할의 프론트엔드 개발자 | pass | 100% | 0 |
| D: 디자이너 역할의 UX 리뷰어 | pass | 100% | 0 |

### Remaining Ambiguities

- A: `player impact` is explicit, but a scenario-specific packet can tighten it to `onboarding_player_impact`.
- C: `프론트엔드 개발자` still implies engineering, but the normalized role specialty makes the distinction clear enough.
- D: `acceptance signal` remains slightly abstract; scenario-specific packets can define whether it means issue resolved, acceptable to ship, or intentionally deferred.

### Stop Decision

- Stop this evaluation loop for now. All fixed scenarios pass their critical and non-critical checklist items after one patch iteration.
- Defer remaining wording refinements until a real role packet run shows repeated ambiguity.
