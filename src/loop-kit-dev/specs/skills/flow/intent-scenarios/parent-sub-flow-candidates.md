# parent sub-flow candidates

## Intent

큰 사용자 요청은 바로 change-unit flow로 실행하지 않고 parent operational-preparation flow가 finite sub-flow candidates를 만들 수 있어야 합니다.

## Scenario

사용자 요청: `로그인 페이지 만들자.`

## Expected Behavior

- 현재 flow type: `operational-preparation`
- 산출물: finite `sub-flow candidates`
- 후보 예:
  - `login-ui-components`
  - `login-logic`
  - `login-page-assembly`
- 각 후보에는 scope, non-goals, completion criteria, verification expectation, handoff condition이 있어야 합니다.

## Forbidden Behavior

- parent flow가 후보를 만들자마자 첫 후보를 자동 실행하지 않습니다.
- `analysis`, `implementation`, `verification` 같은 phase 이름만 후보로 만들지 않습니다.
