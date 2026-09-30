## 사용자 스펙 의도

- 코드의 구조와 최적화는 code 레이어가 맡아야 한다.
- 엔지니어링 방향과 source-level 구현 판단을 분리하고 싶다.
- sw-kit 플러그인 수정하자. 코드 작성시 한 함수안에서도 맥락이 바뀌는 지점에 빈줄을 남기는 방식의 readability 개선하려고
- prevent로 기본 흐름 방지 맥락이고, 다음이 함수에서 사용할 변수 초기화여서 맥락이 다른데 왜 붙어있지?
  초기화 + 초기화검증은 붙는게 맞긴한데
- 코드작성 논리를 설명하는건데 js 예제를 넣는건 일반화하기 어렵지 않아?
- 그리고 input에 대한 맥락도 같이 있어야되서 설명이 장황해지잖아
- handleImport에서 initial 함수 진입시 필수 검토 if 문 다음에 맥락 바뀌는거 아냐?
  변수 선언시 n 줄 넘어가면 그것만으로도 읽는데 피로도가 생겨 맥락이 바뀌는걸로 취급될 정도인거 같아. 이걸 뭐라 설명해야할지 모르겠네
- oneline if 절이 연속되면 그냥 붙여도 되는걸로하자.
- 매개변수로 함수실행 넣는건 되도록이면 피하는게 맞을거 같아. 이게 188L에서 그런식으로 표현되어있어.
- 테스트 작성·감사 기준을 최대한 유지하면서 sw-kit 플러그인에서 사용하고 싶다.
- SKILL.md에 스킬 사용시 자주 사용하게되는것들을 넣고, 필요에 따라 챙겨봐야하는걸 references에 들어가야되는데, test-audit.md 파일이 들어가면 너무 자주 읽게되지 않아?

---

# Code Skill Spec

## 목적

`code`는 선택된 방향 안에서 정확하고 이해하기 쉬우며 유지 가능한 production code를 구현, 수정, 리팩터링, 테스트, 리뷰한다.

## 경계

- 포함: 파일·module 구조, control flow, error handling, local reuse, abstraction, 테스트 작성·수정·리뷰·집중 감사·subsystem campaign, source-level dependency use, risk-focused review.
- 제외: system architecture와 technology/storage 선택, product requirement 결정, repository-wide lifecycle audit, Git workflow.

## 처리 계약

- 관련 코드, callers, tests, repository conventions를 확인한다.
- 현재 요구에 맞는 가장 작은 일관된 변경을 만든다. 미래 가능성만으로 abstraction을 만들지 않는다.
- 작성·수정하는 함수에서는 목적과 독해 부담을 기준으로 빈줄을 둔다. 함수 시작에 있는 진입 가드는 첫 읽기 단위로 보고, 그 뒤를 빈줄로 구분한다. 같은 검사 흐름의 연속된 한 줄 `if`는 붙여 쓸 수 있다. 준비와 즉시 검증은 보통 붙이지만, 긴 선언·표현식은 그 자체로 하나의 읽기 단위가 될 수 있다.
- 의미 있는 작업을 다른 호출의 인수 안에서 실행하면 결과에 이름을 붙여 전달한다. 중첩이 짧고 흐름이 분명한 경우에는 직접 전달할 수 있다.
- 재사용은 의미, contract, ownership, lifecycle이 맞을 때만 선택한다.
- 오류와 boundary condition을 명시적으로 처리하고 중요한 동작을 검증한다.
- 리뷰에서는 결함과 유지보수 위험을 영향 순으로 제시하며, 수정 권한이 없는 요청에는 findings만 낸다.

## 테스트 가치와 본문 계약

테스트는 관찰 가능한 동작, 실제 회귀 또는 독립 계약을 보호한다. 목표는 신뢰와 유지 비용이며 삭제량이나 속도 개선을 목표로 삼지 않는다. 이 기준은 테스트 작성·수정·리뷰에 적용하며 일반 코드 작업을 테스트 감사로 확대하지 않는다.

`SKILL.md`는 작성 gate, junk patterns, 유지 예외, 회귀 검증을 직접 소유한다. 일상적인 테스트 작업은 감사 reference를 읽지 않고 판단할 수 있어야 한다.

새 테스트를 추가하거나 기존 테스트를 수정하기 전에 다음을 확인한다. 답이 없으면 테스트를 작성하지 않는다.

1. 보호하는 관찰 동작·불변식·독립 계약.
2. 테스트가 검출할 실제 회귀와 실패 이유.
3. 기존 coverage가 그 실패를 검출하지 못하는 이유. 계약은 가장 강한 책임 경계에 주된 test owner를 둔다. 다른 계층은 전달·수명주기 등 owner가 도달하지 못하는 별도 위험이 필요하다. 서로 다른 입력과 실패 조건은 유지한다. 기존 table·fixture 확장을 우선하고 중복 setup은 같은 변경에서 통합한다.
4. production caller가 필요로 하지 않는 export·flag·wrapper·injection hook 여부. 테스트는 실제 책임 경계로 옮긴다. 실제 의존성·플랫폼·수명주기 경계는 테스트 편의를 위한 설계 왜곡과 구분한다.

유지 예외가 독립 계약을 설명하지 못하면 다음 junk patterns는 작성 gate를 통과하지 못한다.

- assertion 없는 coverage probe.
- 자기 비교.
- 입력 그대로 복사 검증.
- 복사한 fixture·inventory·manifest·export 목록.
- 소스·import·문자열의 정확한 형태 검색.
- 실제 경계에서 이미 검증하는 private predicate·call shape.
- 동일 계약의 중복 호출.
- provider별 shared helper 재검증.
- test-only export·global·wrapper를 유지하기 위한 테스트.
- 테스트만 호출하는 dead production code.
- 대상 helper·renderer로 생성한 expected value.
- 검증할 동작을 직접 구현한 mock 또는 서로 다른 API를 대신하는 동일 mock.
- owner가 만들어야 할 receipt·admission·callback 순서를 공급한 fixture 또는 실제 경로가 쓰지 않는 store의 persistence 검증.
- 전달·acknowledgement 대신 선언된 capability flag만 확인.
- 다른 guard의 거부 또는 도달하지 않는 rejection처럼 엉뚱한 이유로 통과하는 negative control.
- 입력·assertion보다 넓은 동작을 주장하는 이름·fixture.

의미 보존 refactor에 실패하는 새 테스트는 책임 경계에서 재작성한다. 기존 테스트는 조사 후보이지 자동 삭제 대상이 아니다.

### 유지 예외

- public API, plugin SDK, protocol, config, migration, storage, security, platform, default, prompt bytes, generated cross-language, package, release, architecture 독립 계약.
- 순서가 관찰 동작인 call ordering과 실제 실패 가능성이 있는 regression.
- 계약의 사용자 key·byte·path 변경에는 실패하고 identifier-only refactor에는 통과하는 가장 저렴한 독립 source guard.
- baseline에서 실패한 retained test는 제품 버그 가능성으로 재현하고 owner에서 수정한다.

정적 검사·느린 실행은 삭제 이유가 아니다. 구현처럼 보여도 유일한 독립 계약 검증일 수 있으므로 제거 전에 입증한다.

### 회귀 검증

버그 regression은 수정 전 코드에서 의도한 이유로 실패하고 owner 수정 후 통과함을 확인한다. 실패를 확인하지 못하면 검증되지 않은 상태로 보고한다. 동일 버그를 통과하는 모든 계층에 복제하지 않는다. 계약별 추가 테스트는 owner가 도달하지 못하는 별도 위험으로 정당화한다.

## 조건부 감사 계약

- `references/test-audit.md`: 기존 테스트의 삭제·통합 집중 감사에서만 읽는다. 본문의 가치 기준을 사용하며 전체 test·production owner·entry·callers·callees·sibling·overlap·CI routing·관련 이력·필요한 dependency source/types를 조사한다.
- 발견 단계는 read-only이며 편집 전에 근거를 보고한다. 후보마다 정확한 이름·위치, 실제 검출 실패, non-test caller, 남는 더 강한 검증 또는 불필요 근거, 존재 이력, 제거 가능한 production/support, 위험·집중 검증 명령을 갖춘다. 누락이 있으면 삭제 준비가 되지 않았다.
- 범위가 넓고 delegation이 가능하면 책임 경계별 parallel discovery를 사용한다. campaign 외에는 소수의 확실한 후보를 우선한다.
- 수정은 하나의 일관된 owner boundary 단위로 수행한다. 중복 assertion을 keeper로 먼저 옮기고 불필요한 private test-only seam·dead path를 제거한다. production 순감소를 선호하되 public/external 계약과 실제 의존성 경계를 보존한다.
- 검증 중 같은 checkout의 source/test를 편집하지 않는다. 최소 owner·sibling 검증, 삭제한 source guard의 실제 executable/dry-run, formatting·diff hygiene, 저장소 changed gate를 수행한다. 최종 감사 수정 후 가능하면 독립 review를 수행하며 저장소가 요구한 review는 필수 gate로 유지한다. production/tooling과 test/support LOC, 실제 검증 범위, 미해결 항목, publication 상태를 구분해 보고한다. commit·push·PR·merge는 기존 권한과 Git 소유 workflow를 따른다.

## Subsystem Campaign 계약

`references/test-campaign.md`는 사용자가 한 subsystem의 전체 테스트 정리를 요청한 경우에만 읽는다. 집중 감사 계약과 아래 순서를 함께 적용하며 단계별 완료 조건을 충족한 뒤 진행한다.

1. 고정 baseline revision에서 test/support LOC와 모든 파일의 pass/fail을 기록한다. baseline 실패는 별도 목록에 둔다.
2. 실제 production owner별 lane으로 전체 test·shared-boundary 사례·QA/live harness를 정확히 한 번 배정한다.
3. 파일과 parameter table을 전부 읽어 declaration별 ledger를 만든다. `R` 유지, `F` 계약 유지·assertion 수리, `C` keeper로 통합, `D` 근거 있는 삭제. `it.each`는 declaration 하나이며 판정이 다른 row는 따로 기록한다.
4. 두 번째 read-only pass에서 중복 계층을 찾는다. 계약별 keeper, retired files, 옮길 assertion, 제거 가능한 test-only seam을 지정한다. mocked collaborator보다 실제 transport와 fake network를 우선한다.
5. lane별 cutover를 수행한다. shared support는 한 owner가 직렬 수정하고 CI routing·inventory를 맞춘다. 기존 shrink-only baseline은 줄어든 범위만 반영하며 발견한 소유 규칙은 해당 subsystem 운영 문서에 기록한다.
6. 가능하면 경계별 독립 read-only reviewer가 삭제 coverage와 keeper를 비교한다. 누락은 복원하거나 source 근거로 기각한다. 복원한 계약마다 owner mutation으로 실패를 확인한 뒤 자신의 mutation만 byte-for-byte 복원한다. 검증 불가·미해결 항목을 완료로 보고하지 않는다.
7. retained baseline 실패는 owner에서 수리하고 수정만 되돌린 control과 candidate를 같은 harness에서 검증한다. 다른 제품 문제는 follow-up으로 기록한다.
8. 장기 campaign은 저장소 Git 정책에 따라 현재 upstream을 반영한다. upstream이 삭제 파일에 추가한 회귀는 keeper에 옮기고 full subsystem suite와 필요한 live proof를 다시 수행한다. 최종 baseline·LOC·lane·keeper·복원·mutation·제품 수정 근거를 보고한다.

독립 reviewer나 live 환경이 없으면 제한과 대체 검증을 밝힌다. 필수 저장소 gate가 미완료인 결과는 검증 완료로 표현하지 않는다.

## 검토 질문

- 요구와 기존 contract를 보존하는가?
- 코드 흐름, 책임, 실패 처리가 이해 가능한가?
- 재사용 또는 abstraction이 실제 복잡도를 줄이는가?
- 중요한 회귀 위험을 검증했는가?

## 독립성 원칙

기술 방향이 주어지면 따른다. 주어지지 않으면 local implementation 결정에 머물고 material architecture choice는 `engineering`으로 드러낸다. 테스트 기준과 절차는 본문·bundled reference만으로 실행 가능하며 `maintenance` 선행 호출을 요구하지 않는다.
