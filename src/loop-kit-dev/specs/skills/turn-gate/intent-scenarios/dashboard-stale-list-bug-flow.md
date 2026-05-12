# 대시보드 목록 갱신 버그 Flow Boundary 시나리오

## 목적

이 시나리오는 버그 수정 요청에서 `원인 파악`, `수정`, `검증`을 서로 다른 planned flow로 쪼개지 않는지 확인합니다.
이 intent scenario 자체는 실제 버그 수정을 실행하지 않고, 사용자 메시지를 해석해 후속 실행 후보와 검증 기대를 정리하는 Flow 0에서 끝납니다.
실제 수정은 사용자가 이어가기로 했을 때 하나 이상의 change-unit flow로 진행합니다.

## 사용자 메시지

```text
대시보드에서 필터를 바꿔도 목록이 예전 결과로 남아 있어. 고쳐줘.
```

## 사용자 메시지의 의미

- 직접 요청된 변경: 필터 변경 뒤 목록이 최신 결과로 갱신되도록 수정.
- 후속 실행 후보:
  - stale list 원인 확인과 최소 수정
  - 필요할 경우 cache invalidation, query key, state reset, fetch trigger 중 실제 원인에 해당하는 변경
  - 재발 방지를 위한 targeted regression fixture 또는 test
- 사용자가 아직 말하지 않은 것:
  - 대상 대시보드 화면 또는 route
  - 필터 종류
  - 데이터 fetch 방식
  - 기존 test/fixture 존재 여부

대상 화면과 필터가 명확하지 않으면 active question-routing으로 target을 잠급니다.
저장소에서 대상이 명확하고 변경 위험이 낮으면 scope를 추론하고 Flow 0에서 후속 실행 후보를 기록할 수 있습니다.

## 기대하는 Operational-Preparation Flow

- Flow type: `operational-preparation`
- 목적: 버그 증상을 해석하고, target ambiguity, 예상 원인 범위, 후속 실행 후보, 검증 기대를 정리한다.
- 소유 산출물: active flow record의 scope/non-goal, target ambiguity 판단, 후속 실행 후보, verification expectation.
- 완료 기준:
  - 이 intent scenario가 Flow 0에서 끝난다는 점이 드러난다.
  - `원인 파악`, `수정`, `검증`을 planned flow로 나열하지 않는다.
  - 후속 실행 후보가 실제 산출물 변경 단위로 기록된다.
  - targeted regression fixture/test가 필요하면 후속 change-unit 후보 또는 같은 change-unit의 산출물로 기록된다.
  - 별도 release, commit, push, PR은 승인되지 않은 handoff로 남긴다.

### 기대하는 Operational-Preparation 출력 예시

```text
Flow 0: 대시보드 목록 갱신 버그 범위 확인
- Flow type: operational-preparation
- Owns: active flow record, scope/non-goal, target ambiguity, next execution candidates, verification expectation
- Result: 이 intent scenario는 Flow 0에서 종료
- Follow-up candidates:
  - dashboard-stale-list-fix
  - dashboard-stale-list-regression-coverage, 필요한 경우에만
- Question-routing: 대상 dashboard route나 filter가 불명확하면 target lock
```

## Flow 0의 판단 결과

- 이 시나리오 자체의 flow: `dashboard-stale-list-bug-intent`
  - Flow type: `operational-preparation`
  - 소유 산출물: scope/non-goal, target ambiguity 판단, 후속 실행 후보, verification expectation.
  - 종료 조건: 후속 실행 후보가 어떤 산출물 변경 단위로 이어질지 기록한다.
- 후속 실행 후보: `dashboard-stale-list-fix`
  - 후보 type: `change-unit`
  - 예상 산출물: filter 변경에 맞춘 query key, cache invalidation, state reset, fetch trigger, 또는 stale result 원인에 해당하는 최소 코드 변경.
  - 근거: 버그의 실제 원인과 수정은 하나의 검토 가능한 변경 단위로 묶는 것이 자연스럽다.
  - 예상 검증: 필터 변경 후 목록 갱신 path, loading/error state, stale data 재현 여부 확인.
- 조건부 후속 실행 후보: `dashboard-stale-list-regression-coverage`
  - 후보 type: `change-unit`
  - 예상 산출물: targeted regression test, fixture, mock response, 또는 existing test 보강.
  - 조건: 기존 검증 경로가 없거나 버그가 재발하기 쉬운 shared data-fetching boundary에 있을 때만 분리한다.

## 허용되는 압축

수정 범위가 작고 regression coverage를 같은 변경 단위 안에 넣는 것이 자연스럽다면, 후속 실행 후보는 하나로 압축할 수 있습니다.

```text
Flow 0: 대시보드 목록 갱신 버그 범위 확인
```

Flow 0은 intent 판단에서 끝납니다.
실제 실행으로 이어질 때만 `dashboard-stale-list-fix` change-unit flow를 시작합니다.

## Flow가 아닌 항목

- `원인 분석`
- `버그 수정 작업`
- `검증`
- `최종 QA`
- `commit-readiness reporting`

이 항목들은 별도 artifact를 만들지 않는 한 후속 change-unit flow 내부 phase 또는 reporting/handoff로 남아야 합니다.

## 평가 관점

- 버그 요청을 phase list로 쪼개지 않는다.
- intent scenario 자체는 Flow 0에서 끝낸다.
- 후속 실행 후보는 실제 산출물 변경 단위로 기록한다.
- regression fixture/test는 필요할 때만 후속 후보로 기록한다.
- target ambiguity가 있으면 active question-routing과 pending question state를 기록한다.
- release, commit, push, PR은 initial agreement 또는 user-gated approval 없이 포함하지 않는다.

## 수용 신호

Fresh executor는 이 intent scenario를 Flow 0에서 끝내고, 실제 실행이 필요하면 `dashboard-stale-list-fix`를 후속 change-unit 후보로 기록해야 합니다.
검증 보강이 필요할 때만 `dashboard-stale-list-regression-coverage`를 조건부 후속 후보로 둡니다.
