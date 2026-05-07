# 로그인 페이지 Flow Boundary 시나리오

## 목적

이 시나리오는 사용자의 한 문장 요청을 곧바로 제품 변경 flow로만 나누지 않고, 먼저 사용자 메시지를 해석해 planned flow list를 만드는 운영 flow를 보존하는지 확인합니다.
또한 planned flow가 직접 사용자 가치 단위만으로 제한되지 않고, UI/UX 컴포넌트처럼 보이지 않는 준비성 변경도 검토 가능한 change-unit flow가 될 수 있는지 확인합니다.

## 사용자 메시지

```text
로그인 페이지 만들기
```

## 사용자 메시지의 의미

- 직접 사용자 가치처럼 보이는 것: 로그인 페이지 전체.
- 실제 실행을 위해 필요한 내부 변경 단위:
  - 로그인 UI/UX 컴포넌트
  - 로그인 상태/검증/submit 로직
  - 페이지 route 또는 view 조립
- 사용자가 아직 말하지 않은 것:
  - auth provider
  - 저장 방식
  - route 경로
  - 디자인 시스템 사용 여부
  - 테스트/검증 기대 수준

이 정보가 결과를 바꿀 수 있으면 질문을 열어 scope를 잠가야 합니다.
질문을 열 때도 turn-gate는 active question-routing 상태로 유지되어야 합니다.

## 기대하는 Operational-Preparation Flow

- Flow type: `operational-preparation`
- 목적: 사용자 요청을 해석하고, scope를 질문으로 잠그거나 안전하게 추론하고, approval boundary를 식별하고, planned flow list를 설계한다.
- 소유 산출물: session plan, active flow record, planned flow list, scope/non-goal note, verification expectation, approval-boundary note.
- Question-routing: product scope, auth provider, persistence, route target, verification expectation이 결과를 바꿀 수 있으면 turn을 멈추지 말고 active question-routing을 연다.
- 완료 기준:
  - `000-plan.md` 또는 active flow record에 flow type이 드러난다.
  - 사용자 메시지가 planned flow list로 변환된 근거가 남는다.
  - 추론한 scope와 non-goal이 기록된다.
  - 추가 승인이 필요한 작업과 handoff 지점이 구분된다.

### 기대하는 Operational-Preparation 출력 예시

```text
Flow 0: 로그인 페이지 요청 해석 및 flow 설계
- Flow type: operational-preparation
- Owns: session plan, active flow record, planned flow list, scope/non-goal, verification expectation
- Result: login-ui-ux-components, login-auth-logic, login-page-assembly를 change-unit flow로 계획
- Question-routing: auth provider나 route가 결과를 바꾸면 active question-routing으로 scope lock
```

## 기대하는 Change-Unit Planned Flows

1. `login-ui-ux-components`
   - Flow type: `change-unit`
   - 소유 산출물: 로그인 페이지에 필요한 재사용 UI/UX 컴포넌트.
   - 근거: 이 자체가 직접 사용자 가치로 보이지 않을 수 있지만, 검토 가능하고 commit-sized인 변경 단위다.
   - 완료 기준: 입력, 비밀번호, submit, error/display state 같은 UI 요소가 기존 디자인 규칙에 맞게 구현되거나 기존 컴포넌트로 조립된다.
   - 검증 기대: 렌더링, responsive layout, 접근성 label, 기본 interaction state 확인.
2. `login-auth-logic`
   - Flow type: `change-unit`
   - 소유 산출물: 로그인 form state, validation, submit flow, API/client integration, error handling.
   - 완료 기준: 입력값 검증, loading/error/success state, API 호출 또는 mock boundary가 명확하다.
   - 검증 기대: validation branch, 실패 응답, 성공 응답 또는 mock path 확인.
3. `login-page-assembly`
   - Flow type: `change-unit`
   - 소유 산출물: route/page composition, UI component와 login logic wiring, page-level verification.
   - 완료 기준: route 또는 page entry가 생성되고 UI/logic이 연결되며 사용자 관점의 로그인 화면이 동작한다.
   - 검증 기대: page render, submit path, error display, navigation 또는 post-login handoff 확인.

## 허용되는 압축

작업 규모가 작거나 기존 컴포넌트와 auth helper가 이미 충분하면 다음처럼 압축할 수 있습니다.

```text
Flow 0: 로그인 페이지 요청 해석 및 flow 설계
Flow 1: 로그인 페이지 조립
```

단, 이 경우에도 `Flow 0`은 operational-preparation flow이고, `Flow 1`은 change-unit flow여야 합니다.
압축 이유는 "이미 reusable component와 auth helper가 있어 별도 변경 단위로 나눌 필요가 없다"처럼 산출물 기준으로 설명해야 합니다.

## Flow가 아닌 항목

- `analysis`
- `work`
- `verification`
- `final QA`
- `commit-readiness reporting`

이 항목들은 별도 검토 가능한 artifact를 만들지 않는 한 core phase, verification/reporting section, 또는 user-gated handoff로 남아야 합니다.

## 평가 관점

- 사용자 메시지 해석과 planned flow list 설계를 `operational-preparation` flow로 보존한다.
- 실행할 planned flows를 `change-unit` flow로 구성한다.
- UI/UX 컴포넌트 생성처럼 직접 사용자 가치가 아닌 변경도 필요한 경우 flow로 허용한다.
- 질문을 열면 active question-routing과 pending question state를 기록한다.
- scope를 추론하면 non-goal과 verification expectation을 함께 남긴다.
- final QA/readiness-only 성격의 확인은 별도 artifact 변경이 있을 때만 planned flow로 다룬다.

## 수용 신호

Fresh executor는 사용자 메시지가 planned flow list로 바뀌는 위치를 `operational-preparation` flow로 유지해야 합니다.
실행 목록은 `change-unit` flow들로 구성되어야 합니다.
평가 결과는 위 평가 관점과 연결해 설명할 수 있어야 합니다.
