# 로그인 페이지 생성 요청 Flow Boundary 시나리오

## 목적

이 시나리오는 `로그인 페이지 만들자`라는 넓은 기능 요청을 받았을 때, 곧바로 구현 flow로 들어가지 않고 먼저 Flow 0에서 scope와 후속 실행 후보를 정리하는지 확인합니다.
이 시나리오는 현재 `turn-gate` runtime/spec이 넓은 기능 요청을 어떻게 flow boundary로 나누는지 평가하는 기준을 고정합니다.

## 사용자 메시지

```text
로그인 페이지 만들자
```

## 사용자 메시지의 의미

- 실행 의도는 분명합니다.
- expected task tier: `multi-flow`
- expected verification method: `clean-context`
- 하지만 다음 항목에 따라 산출물과 검증 경로가 달라질 수 있습니다.
  - 정적 로그인 화면만 만들지
  - 실제 인증 API 연동까지 포함할지
  - route/page를 새로 만들지
  - 기존 auth provider나 API가 있는지
  - 디자인 시스템을 따를지
  - 테스트, browser 확인, build 확인 중 무엇을 완료 기준으로 볼지

따라서 이 요청은 바로 active `change-unit flow`로 시작하기보다, 먼저 Flow 0에서 scope lock과 후속 후보 정리가 필요합니다.

## 기대하는 Flow 0

- Flow type: `operational-preparation`
- 목적: 로그인 페이지 요청의 범위, target, approval boundary, verification expectation을 정리한다.
- 소유 산출물:
  - scope/non-goal
  - target ambiguity 판단
  - 질문이 필요한 지점
  - 후속 `change-unit` 후보
  - verification expectation
- 완료 기준:
  - Flow 0이 실제 구현을 시작하지 않는다.
  - 후속 후보가 active execution flow가 아니라는 점이 드러난다.
  - 범위를 바꿀 수 있는 정보는 active question-routing으로 잠근다.

## 후속 Change-Unit 후보

아래 항목들은 아직 active execution flow가 아니라 후보입니다.
사용자가 범위를 승인하거나 특정 후보를 선택한 뒤에야 별도 `change-unit flow`로 시작합니다.

1. `login-page-ui-ux`
   - 후보 type: `change-unit`
   - 예상 산출물: route/page, form layout, input/button/error 상태.
   - 예상 검증: 화면 렌더링, responsive layout, 접근성 label, 기본 interaction state.
2. `login-form-client-behavior`
   - 후보 type: `change-unit`
   - 예상 산출물: 입력 상태, 필수값/형식 검증, submit/loading/error state.
   - 예상 검증: validation branch, submit state, error display.
3. `login-auth-integration`
   - 후보 type: `change-unit`
   - 예상 산출물: 인증 API 또는 기존 auth 로직 연동, 성공/실패 처리, post-login handoff.
   - 예상 검증: 성공 응답, 실패 응답, redirect 또는 session handoff.
4. `login-page-coverage`
   - 후보 type: `change-unit`
   - 예상 산출물: UI test, validation test, story, snapshot, 또는 주요 상태 검증.
   - 조건: 기존 검증 경로가 부족하거나 변경 위험이 충분할 때만 별도 후보로 둔다.

## 질문이 필요한 지점

- 대상 app 또는 route는 어디인가?
- 범위는 정적 UI만인가, 실제 로그인 동작까지인가?
- 인증 방식은 기존 API/auth provider인가, mock인가?
- 디자인 기준은 기존 디자인 시스템인가, 새 화면 임의 설계인가?
- 완료 기준은 화면 렌더링, 성공/실패 처리, 테스트 통과 중 무엇인가?
- 검증 기대는 lint, test, build, browser 확인 중 무엇인가?

## Flow가 아닌 항목

- `분석`
- `구현`
- `검증`
- `최종 QA`
- `commit-readiness reporting`

이 항목들은 후속 `change-unit flow` 내부 phase이거나 reporting/handoff입니다.
별도 검토 가능한 artifact를 만들지 않으면 planned flow나 후보 flow로 세지 않습니다.

## 평가 관점

- Flow 0을 `operational-preparation`으로 둔다.
- 후속 후보를 active execution flow로 세지 않는다.
- scope가 넓고 결과가 여러 방향으로 갈리므로 질문이 필요한 지점을 기록한다.
- phase list를 planned flow로 만들지 않는다.
- final QA/readiness-only 항목을 별도 artifact 없이 후보로 만들지 않는다.

## 수용 신호

Fresh executor는 `로그인 페이지 만들자`를 곧바로 구현하지 않고 Flow 0에서 scope lock과 후속 후보 정리로 처리해야 합니다.
후속 후보는 사용자가 선택하거나 승인하기 전까지 active execution flow가 아닙니다.
