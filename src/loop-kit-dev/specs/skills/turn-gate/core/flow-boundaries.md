# turn-gate flow-boundaries sub-spec

## 목적

이 문서는 `turn-gate`에서 planned flow가 무엇인지, 어떤 단위가 flow가 아닌지, `operational-preparation`과 `change-unit`을 어떻게 구분하는지 소유합니다.

`core/runtime-flow.md`는 phase 순서와 전환 조건을 소유하고, 이 문서는 flow boundary와 flow taxonomy를 소유합니다.

## Flow 종류

`turn-gate`에는 두 층의 flow가 있습니다.

- `operational-preparation flow`: 사용자 메시지를 받아 intent, scope, non-goal, acceptance signal, verification expectation, approval boundary를 잠그고 planned flow list를 만드는 운영 flow입니다.
- `change-unit flow`: planned flow list 안에서 실제 코드, 문서, fixture, 설정, release surface 같은 검토 가능한 산출물 변경을 소유하는 실행 flow입니다.

`operational-preparation flow`의 산출물은 `.agents/sessions/{YYYYMMDD}/000-plan.md`, active flow record, planned flow list, approval boundary note 같은 session/plan artifact입니다.
사용자 메시지 해석과 flow list 설계는 멈춤이나 terminal response가 아니라 운영 flow의 preparation/work로 기록될 수 있습니다.
질문을 열어 scope를 잠그는 경우에도 `active question-routing` 상태로 turn을 유지합니다.

## Planned Flow Boundary

`turn-gate`에서 planned flow는 phase가 아니라 응집된 변경 단위입니다.

`분석`, `작업`, `검증`, `보고`, `커밋 준비` 같은 단계명은 flow boundary가 될 수 없습니다.
이 항목들은 각 flow 내부의 core phase 또는 phase protocol로 남아야 합니다.

flow가 반드시 최종 사용자에게 직접 보이는 가치 단위일 필요는 없습니다.
flow는 함께 이해하고 검토하고 검증하고 필요하면 커밋할 수 있는 변경 묶음이어야 합니다.

예를 들어 "로그인 페이지 만들기"는 하나의 사용자 가치처럼 보일 수 있지만, planned flows는 `로그인 UI/UX 컴포넌트 생성`, `로그인 로직 작성`, `로그인 페이지 조립`처럼 커밋 가능한 응집 단위로 나뉠 수 있습니다.

## Flow가 아닌 항목

다음 항목은 별도 산출물 변경을 소유하지 않는 한 planned flow boundary가 아닙니다.

- 순수 최종 QA
- 통합 검증
- 정합성 점검
- 검증 결과 보고
- commit-readiness reporting

이 항목들은 마지막 변경 단위 flow의 `verification`/`reporting` 또는 별도 user-gated handoff로 남깁니다.

단, 회귀 테스트 fixture, snapshot baseline, 문서, 운영자 리포트 출력, validator 진단 출력처럼 별도 검토 가능한 산출물을 만들거나 바꾸는 경우에는 그 산출물 변경이 planned flow가 될 수 있습니다.

## Follow-Up 후보와 Active Flow 구분

실제 product/code/document 변경 실행이 아직 선택되지 않은 경우, `operational-preparation flow`는 후속 `change-unit` 후보와 검증 기대를 기록하고 reporting으로 이어질 수 있습니다.

후속 후보는 active execution flow가 아닙니다.
후보를 실제로 수행하기 전에는 대상 파일 수정, fixture 갱신, release surface build, commit-readiness reporting을 같은 운영 flow의 작업으로 끌어오지 않습니다.

flow list 변환 자체가 산출물로 남는 경우, 그 산출물은 `operational-preparation flow`의 work/reporting 결과입니다.
변환 결과 목록의 각 실행 항목은 실제 실행 시 `change-unit flow`가 됩니다.

사용자가 실행이 아니라 판단, 설계, 범위 확인만 요청한 경우 변환 결과 목록은 확정 planned sequence가 아니라 후속 실행 후보로 남을 수 있습니다.
이 경우 reporting은 실제 실행 후보와 handoff 조건을 정리하고, 실행 flow를 시작하지 않습니다.

## Session Record와의 관계

`records/session-records.md`는 위 boundary를 어떻게 `.agents/sessions/{YYYYMMDD}/000-plan.md`와 개별 flow record에 기록하는지 소유합니다.

따라서 session record는 다음을 드러내야 합니다.

- flow type이 `operational-preparation`인지 `change-unit`인지
- planned flow sequence가 phase checklist가 아니라 응집된 변경 단위들의 순서인지
- 후속 실행 후보와 active execution flow가 섞이지 않았는지
- final QA, readiness, reporting 같은 확인 작업이 산출물 변경 없이 planned flow로 승격되지 않았는지

## 검토 질문

- planned flow가 phase checklist가 아니라 응집된 변경 단위인가?
- flow boundary가 직접 사용자 가치만으로 좁혀져 보이지 않는 커밋 가능 준비 작업을 배제하지 않는가?
- 사용자 메시지 해석과 planned flow list 설계가 필요한 경우 운영 flow로 기록되고, 실행용 `change-unit` 후보와 분리되는가?
- 실행이 아니라 판단, 설계, 범위 확인만 요청된 경우 후속 후보를 active execution flow로 오해하지 않는가?
- 최종 QA, 정합성 점검, readiness 보고가 별도 산출물 변경 없이 planned flow로 승격되지 않는가?
