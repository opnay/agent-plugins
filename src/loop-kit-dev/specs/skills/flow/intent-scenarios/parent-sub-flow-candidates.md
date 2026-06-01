# parent sub-flow candidates

## Intent

큰 사용자 메시지는 메시지 인터뷰와 플로우 설계를 거쳐 finite sub-flow candidates를 만들 수 있어야 합니다.

## Scenario

사용자 요청: `로그인 페이지 만들자.`

## Expected Behavior

- 메시지 인터뷰는 acceptance, scope edge, non-goal, approval boundary가 충분한지 확인합니다.
- clarity가 부족하면 high-leverage question을 산출합니다.
- clarity가 충분하면 플로우 설계가 parent operational-preparation flow로 flow 구성을 만듭니다.
- 후보 예:
  - `login-ui-components`
  - `login-logic`
  - `login-page-assembly`
- 각 후보에는 scope, non-goals, completion criteria, verification expectation, approval boundary, handoff condition이 있어야 합니다.

## Boundary Behavior

- 후보는 routing 전 pending 상태로 유지합니다.
- 후보 이름은 phase 이름보다 reviewable artifact 또는 change unit을 기준으로 둡니다.
