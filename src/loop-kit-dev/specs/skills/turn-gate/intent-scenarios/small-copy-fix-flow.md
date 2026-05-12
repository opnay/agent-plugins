# 작은 문구 수정 Flow Boundary 시나리오

## 목적

이 시나리오는 작은 사용자 요청을 불필요하게 여러 planned flow로 쪼개지 않는지 확인합니다.
이 intent scenario 자체는 실제 파일 수정을 실행하지 않고, 사용자 메시지를 해석해 다음 실행 단위가 무엇인지 판단하는 Flow 0에서 끝납니다.
실제 문구 수정은 사용자가 실행을 이어가기로 했을 때 후속 change-unit flow로 진행합니다.

## 사용자 메시지

```text
로그인 버튼 문구가 "로그인 하기"로 되어 있는데 "로그인"으로 바꾸자.
```

## 사용자 메시지의 의미

- 직접 요청된 변경: 로그인 버튼의 표시 문구 수정.
- 후속 실행 후보: 버튼 문구가 정의된 component 또는 locale/resource entry 수정.
- 사용자가 아직 말하지 않은 것:
  - 버튼 문구가 여러 화면에 재사용되는지
  - i18n resource를 쓰는지
  - snapshot 또는 visual baseline 갱신이 필요한지

대상 파일이나 문구 출처가 명확하면 추가 질문 없이 scope를 추론할 수 있습니다.
여러 버튼이나 locale이 같은 문구를 공유해 결과가 달라질 수 있으면 active question-routing으로 target을 잠급니다.

## 기대하는 Operational-Preparation Flow

- Flow type: `operational-preparation`
- 목적: 사용자 요청이 작은 문구 수정인지 확인하고, target ambiguity와 검증 기대를 가볍게 정리한다.
- 소유 산출물: active flow record의 scope/non-goal, 다음 실행 후보, verification expectation.
- 완료 기준:
  - 이 intent scenario가 Flow 0에서 끝난다는 점이 드러난다.
  - 후속 실행이 필요하다면 단일 change-unit 후보로 충분하다는 근거가 기록된다.
  - 수정 대상이 명확하거나, 명확하지 않다면 active question-routing으로 target lock을 연다.
  - 별도 release, commit, push, PR은 승인되지 않은 handoff로 남긴다.

### 기대하는 Operational-Preparation 출력 예시

```text
Flow 0: 로그인 버튼 문구 수정 범위 확인
- Flow type: operational-preparation
- Owns: active flow record, scope/non-goal, next execution candidate, verification expectation
- Result: 이 intent scenario는 Flow 0에서 종료. 실제 수정이 이어지면 login-button-copy-fix를 단일 change-unit 후보로 사용
- Question-routing: 같은 문구가 여러 버튼/locale에 걸쳐 있으면 target lock
```

## Flow 0의 판단 결과

- 이 시나리오 자체의 flow: `login-button-copy-intent`
  - Flow type: `operational-preparation`
  - 소유 산출물: scope/non-goal, target ambiguity 판단, 후속 실행 후보, verification expectation.
  - 종료 조건: 후속 실행 후보가 단일 change-unit으로 충분한지 판단하고 기록한다.
- 후속 실행 후보: `login-button-copy-fix`
  - 후보 type: `change-unit`
  - 예상 산출물: 로그인 버튼 문구가 정의된 component, translation resource, snapshot, 또는 관련 fixture 중 실제로 필요한 최소 변경.
  - 근거: 작은 문구 수정은 실제 실행 시 하나의 검토 가능한 변경 단위로 충분하다.
  - 예상 검증: 관련 component render, locale lookup, snapshot 또는 targeted test가 있으면 해당 경로 확인.

## 허용되는 압축

이 intent scenario에서는 아래처럼 Flow 0 하나로 끝나는 것이 기대 동작입니다.

```text
Flow 0: 로그인 버튼 문구 수정 범위 확인
```

실제 작업 실행으로 이어지는 경우에만 다음 흐름에서 `login-button-copy-fix` change-unit flow를 시작합니다.
그 경우에도 Flow 0의 active flow record에는 scope 추론과 verification expectation이 남아야 합니다.

## Flow가 아닌 항목

- `문구 분석`
- `문구 수정 작업`
- `문구 검증`
- `commit-readiness reporting`

이 항목들은 별도 artifact를 만들지 않는 한 하나의 change-unit flow 내부 phase 또는 reporting/handoff로 남아야 합니다.

## 평가 관점

- 작은 요청을 불필요한 phase flow로 쪼개지 않는다.
- intent scenario 자체는 Flow 0에서 끝낸다.
- target ambiguity가 없으면 후속 실행 후보를 단일 change-unit으로 판단한다.
- target ambiguity가 있으면 질문을 열되 active question-routing과 pending question state를 기록한다.
- 문구 수정 외 release, commit, push, PR은 initial agreement 또는 user-gated approval 없이 포함하지 않는다.
- snapshot 또는 fixture 갱신이 실제로 필요할 때만 같은 change-unit의 산출물로 포함한다.

## 수용 신호

Fresh executor는 이 intent scenario를 Flow 0에서 끝내고, 후속 실행이 필요하면 `login-button-copy-fix`를 단일 change-unit 후보로 기록해야 합니다.
단, 문구 출처가 모호한 저장소에서는 active question-routing을 열고 그 상태를 session record에 남겨야 합니다.
