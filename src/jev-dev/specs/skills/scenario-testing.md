## 사용자 스펙 의도

- 재사용되는 instruction을 작성자 감이 아니라 fresh executor evidence로 테스트하고 싶다.
- 시나리오와 checklist를 먼저 잠그고 분석 결과를 evidence 중심으로 보고하고 싶다.
- empirical evaluation이 불가능한 상황과 구조 리뷰만 가능한 상황을 구분하고 싶다.
- 기본 평가 세트는 40개 시나리오를 기준으로 삼고, 큰 평가 요청은 실행 가능한 batch로 나눠 진행하고 싶다.
- subagent 한도에 걸리면 완료됐거나 더 쓰지 않는 executor를 닫고 남은 batch를 이어가고 싶다.
- "아 모든 시나리오나 jev 호출 기록은 어디 문서에 남겨줘."
- "보통 엄청 디테일하게 보려고 시나리오 케이스 120개 요청해서 작업하는데, 이걸 기준으로 한다면?"
- "결국 시나리오 테스팅이란건 스킬을 읽은 에이전트가 결국 어떤 행동을 할지 예측하는거고, 그거에 대한 choice 결과가 expect와 부합한지 이걸 먼저 봐야될거 같아"
- "선택지 밖의 선택을 하는건 진짜 스킬 문장을 아에 재설계 해야되는 케이스로도 볼 수 있을거 같아. unexpected 처리를 위한 선택지 고정하는거 어때?"
- "이 스킬에서 8개 시나리오로 제한걸었던게 비용이 비싸서 였거든? 이제 5배 해서 40개 시나리오로 늘려놓자."
- "이제 더이상 codex 만의 일은 아니게되어서 플러그인 성격과 달라졌어. jev 플러그인으로 옮기는거 어때?"

---

# scenario-testing 스킬 스펙

## 목적

`scenario-testing`은 reusable agent-facing instruction을 읽은 에이전트가 어떤 행동을 할지 시나리오별로 예측하고, fresh executor의 실제 행동이 사전에 고정한 기대와 맞는지 evidence로 평가합니다.
대상은 skill, slash command, task prompt, AGENTS section, CLAUDE section, code-generation prompt처럼 반복 사용되는 텍스트 지시입니다.

## 경계

- 포함:
  - reusable instruction의 평가 시나리오·fixture 설계와 유효성 확인
  - 시나리오별 행동 Choice와 hidden expected choice 고정
  - 선택적 Jev 예측과 fresh executor 실행
  - expected·predicted·actual choice 비교
  - caller-side deterministic verdict와 원인 귀속
  - 시나리오·Jev 호출·실행·비교 evidence 기록
- 제외:
  - 원본 instruction이나 owning bundle의 경계 설계
  - 일반 implementation debugging
  - reusable tool policy 전반
  - disposable prompt polishing
  - target instruction 직접 수정 또는 patch 적용
  - Jev가 `pass`·`partial`·`fail`을 직접 판정하게 하는 방식

## 처리하려는 작업 형태

- 새로 만든 reusable instruction이 실제로 유도할 행동을 검증하는 경우
- agent failure 원인이 implementation이 아니라 instruction ambiguity라고 의심되는 경우
- 고빈도 instruction을 실제 executor 관점에서 분석하는 경우
- 수십~수백 개의 고정 시나리오를 반복 가능한 harness로 평가하는 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `jev/skills/scenario-testing/SKILL.md`
- 호출 방식: `$jev:scenario-testing`을 직접 호출하거나 manifest prompt의 안내를 따른다.

## 핵심 처리 계약

### 평가 순서

다음 순서를 유지합니다.

1. target description과 body의 정적 정합성을 확인합니다.
2. scenario catalog와 fixture를 고정하고 각 fixture의 유효성을 확인합니다.
3. 각 시나리오에 scenario-specific behavior Choice를 정의합니다.
4. caller가 `expected_choice`를 미리 고정하고 Jev·executor에게 숨깁니다.
5. 허용된 경우 Jev가 단건 또는 batch로 `predicted_choice`를 선택합니다.
6. fresh executor가 target instruction을 읽고 시나리오를 실행합니다.
7. caller가 관찰 증거로 `actual_choice`를 분류합니다.
8. expected·predicted·actual 관계와 critical check를 비교합니다.
9. caller가 deterministic verdict를 산출하고 evidence를 기록합니다.

한 evaluation pass에는 하나의 target/theme만 다룹니다. 사용자가 scenario 수를 지정하면 그 수를 우선하고, 별도 요청이 없으면 40개를 사용합니다. 명시적 smoke check는 2~3개를 사용할 수 있습니다. 큰 세트는 실행 가능한 크기로 batch 처리하되 scenario ID와 고정 계약을 유지하며, batch history 자체를 평가 결과로 기록하지 않습니다.

### Fixture 유효성

- fixture는 prompt, 초기 상태, 허용 mutation, 관찰 지점, 완료 조건을 실제로 제공해야 합니다.
- 실행 전에 fixture가 선언한 상태와 관찰 가능성을 확인합니다.
- 불충분하거나 drift한 fixture는 `invalid_fixture`로 제외하고 instruction 실패로 계산하지 않습니다.
- 실행 중 환경 자체가 실패하면 `execution_environment_error`로 분리합니다.

### Behavior Choice 설계

- 각 시나리오는 해당 decision point에서 가능한 2~6개의 준비된 행동 label과 공통 `unexpected_action`을 사용합니다.
- label은 상호 배타적이어야 하고 `pass`, `partial`, `fail` 같은 품질 판정이 아니라 실제 행동을 표현해야 합니다.
- 각 label에 관찰 가능한 판별 기준을 둡니다.
- 모든 Choice에 `unexpected_action`을 마지막 공통 label로 포함합니다.
- `unexpected_action`은 실행 행동이 관찰됐지만 준비한 label 어느 것에도 맞지 않는 경우입니다.
- `insufficient_observation`, `invalid_fixture`, `execution_environment_error`는 행동 label이 아니라 별도 상태입니다.
- `expected_choice`는 예측·실행 전에 caller가 고정하며 Jev packet과 executor prompt에 노출하지 않습니다.

### Jev 예측

- Jev는 선택적 외부 predictive judge입니다. 사용자가 Jev 사용과 외부 전송을 허용한 경우에만 `$jev:jev`의 Choice 계약을 따릅니다.
- 유효한 scenario가 하나면 target instruction의 필요한 부분, scenario prompt, 정제된 fixture, choice label·기준만 하나의 `--if` 문맥으로 보내고 `--threshold 0 --json --pick answer,probabilities`를 사용합니다.
- 유효한 scenario가 둘 이상이고 공통 target instruction을 공유하면 정제한 target instruction을 batch state file로, scenario prompt·fixture·label 기준을 JSONL row의 `question`으로 보내고 row `id`는 scenario ID를 사용합니다.
- 단건 `--conditions`와 batch row `conditions`에는 해당 scenario의 behavior labels와 `unexpected_action`만 둡니다.
- Batch는 전체 결과를 제공하므로 threshold·pick을 사용하지 않습니다. 저장된 단건 필터·출력 설정은 batch에 적용되지 않습니다.
- expected choice, caller verdict, 로컬 절대경로, 비밀, 사용자 식별 정보, 불필요한 repository state를 전송하지 않습니다.
- 전송 권한이 없으면 `prediction_skipped_no_authority`를 기록하고 executor 평가를 계속합니다.
- Jev command가 없거나 단건·batch 실행이 불가능하면 `prediction_unavailable`을 기록하고 executor 평가를 계속합니다.
- Jev exit 1·2 또는 malformed output은 행동 판정이 아니며 prediction error로 기록합니다. 성공한 동일 질문을 재호출하지 않습니다.
- Jev는 verdict를 선택하지 않습니다. verdict는 caller가 실제 행동과 critical evidence에서 산출합니다.

### Fresh 실행과 실제 행동 분류

- empirical run마다 fresh subagent를 사용하며 author self-reread나 이전 executor 재사용을 대체물로 허용하지 않습니다.
- executor에는 target instruction, scenario, fixture, 실행 권한과 관찰 항목만 제공합니다. expected choice, predicted choice, suspected defect는 숨깁니다.
- executor는 실행 결과, 명령·도구 사용, 전후 상태, ambiguity, judgment call, retry를 구조화해 반환합니다.
- caller는 실행 후 같은 behavior Choice로 `actual_choice`를 분류합니다.
- 행동이 관찰되지 않아 분류할 수 없으면 `insufficient_observation`으로 두고 choice를 추정하지 않습니다.

### Unexpected 귀속

`actual_choice = unexpected_action`이면 자동 fail로 계산하지 않고 다음 원인 중 하나로 귀속합니다.

- `choice_set_incomplete`: 유효한 행동을 Choice가 빠뜨림. Choice를 보완하고 비교 가능성에 따라 재평가합니다.
- `skill_behavior_underspecified`: target instruction이 다른 행동을 허용하거나 유도함. instruction 재설계 후보입니다.
- `executor_deviation`: instruction과 fixture로 설명되지 않는 단발 실행 이탈. fresh executor 재실행이 필요합니다.
- `fixture_drift`: 선언한 fixture와 실행 상태가 다름. `invalid_fixture`로 처리합니다.
- `unresolved_attribution`: 현재 evidence로 원인을 구분할 수 없음. 판정을 보류하고 필요한 evidence를 적습니다.

반복된 predicted·actual unexpected 행동은 `skill_behavior_underspecified`의 강한 신호입니다. 선택지 밖 행동이 unauthorized·destructive critical violation이면 귀속 전후와 무관하게 `critical_fail`입니다.

### 비교와 caller verdict

세 choice의 관계를 다음처럼 해석합니다.

- `E=P=A`: 기대 행동이 명확하고 안정적입니다.
- `E=P`, `A` 다름: executor instability 또는 실행 조건 문제를 우선 확인합니다.
- `P=A`, `E` 다름: instruction defect의 강한 신호입니다.
- `E=A`, `P` 다름: Jev prediction error 또는 prediction packet 문제입니다.
- 모두 다름: scenario·Choice·fixture 설계를 재검토합니다.

Caller verdict 순서는 고정합니다.

1. 관찰된 unauthorized·destructive action: `critical_fail`
2. fixture 무효·drift: `invalid_fixture`
3. 실행 환경 실패: `execution_environment_error`
4. 관찰 부족: `insufficient_observation`
5. 그 밖의 critical contract 위반: `critical_fail`
6. unexpected action: 귀속 결과에 따라 `invalid_choice_set`, `fail`, `rerun_required`, `invalid_fixture`, `needs_attribution`
7. `actual_choice != expected_choice`: `fail`
8. expected 행동은 했지만 noncritical check 누락 또는 heavy inference가 있음: `partial`
9. expected 행동과 모든 check 충족: `pass`

`predicted_choice`는 verdict를 바꾸지 않고 예측 품질과 instruction clarity를 진단합니다.

### Evidence와 지표

모든 empirical evaluation은 사용자가 지정한 위치 또는 task-scoped durable artifact에 다음을 남깁니다.

Scenario 실행 artifact를 plugin `changes/` 디렉터리에 두지 않습니다. `changes/`는 release·migration 계약만 소유하며, 사용자가 명시적으로 change history 편입을 요청한 경우만 예외입니다.

- frozen scenario catalog: ID, prompt, fixture, behavior labels·기준, hidden expected choice, critical·noncritical checks
- executor evidence: 전후 상태, 실행 사실, actual choice, ambiguity, judgment, retries, `tool_uses`, `duration_ms`; unavailable metric은 `null`
- Jev call log: 정제된 request, labels, exit, stdout·stderr, parsed answer·probabilities, timestamp; secret과 미허용 원문 제외
- comparison table: expected·predicted·actual, relation, observation status, attribution, caller verdict, reason
- summary report: 분모와 제외 항목을 밝힌 지표, 발견된 instruction issue, evaluation-design issue, residual risk

큰 평가에서는 CSV 또는 JSONL과 executor용 resume 가능한 harness를 사용해야 합니다. Jev prediction은 공통 state를 공유하는 valid scenario를 `jev batch`로 한 번에 요청하며 CLI 자체의 chunk·resume를 가정하지 않습니다. Scenario catalog와 Jev call log는 scenario ID로 연결되는 단일 evidence set으로 유지합니다.

필수 집계는 다음과 같습니다.

- expected-actual accuracy의 분모는 fixture·환경·관찰이 유효하고 attribution이 확정된 scenario입니다. `invalid_choice_set`, `rerun_required`, `needs_attribution`은 제외합니다.
- predicted-actual accuracy와 expected-predicted agreement의 분모는 성공한 Jev answer가 있고 동일한 frozen Choice로 actual을 확정한 scenario입니다.
- verdict·relation·unexpected attribution 분포
- invalid fixture, insufficient observation, environment error, prediction error·skip 수
- critical failure 수와 반복 ambiguity·judgment·retry 원인

## 환경 제약

- fresh subagent dispatch가 불가능하면 `structural review only` 또는 `evaluation skipped`로 보고하며 empirical evaluation이라 부르지 않습니다.
- 외부 전송 권한이나 Jev 실행 환경이 없으면 prediction만 skip/error로 분리하고 empirical executor evaluation을 Jev 결과로 대체하지 않습니다.
- executor 한도에 걸리면 완료됐거나 더 이상 사용하지 않는 executor를 닫고 다음 batch를 이어갑니다.

## 검토 질문

- fixture와 Choice를 실행 전에 고정했는가?
- Choice가 품질 판정이 아니라 상호 배타적인 관찰 행동을 표현하는가?
- 모든 Choice에 `unexpected_action`이 있고 관찰 부족과 분리되는가?
- expected choice를 Jev와 executor에게 숨겼는가?
- Jev가 verdict가 아니라 행동을 예측했는가?
- empirical run마다 fresh executor를 사용했는가?
- actual choice가 실행 evidence에서 분류됐는가?
- unexpected 행동을 자동 fail 처리하지 않고 귀속했는가?
- caller verdict가 고정 순서로 산출됐는가?
- 모든 scenario와 Jev call이 durable evidence에 연결됐는가?

## 독립성 원칙

- 이 skill은 독립 실행 가능해야 합니다.
- `$jev:jev`는 공개 CLI 계약을 제공하는 sibling skill이며 hidden context가 아닙니다. 설치나 전송 권한 부재가 core executor evaluation을 막지 않습니다.
- 다른 sibling skill은 target 작성이나 후속 수정만 맡습니다. 평가 절차, evidence schema, verdict 규칙은 이 스펙 안에서 닫혀 있어야 합니다.

## 확장 원칙

- 반복되는 scenario·evidence schema는 runtime reference나 template으로 분리할 수 있습니다.
- 새로운 상태나 verdict는 behavior Choice와 섞지 않고 평가 lifecycle의 소유 단계에 추가합니다.
- plugin 사용 기준 변경은 `src/jev-dev/specs/plugin.md`, README, manifest prompt에 함께 반영합니다.
