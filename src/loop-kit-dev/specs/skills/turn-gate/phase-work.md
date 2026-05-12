# turn-gate phase-work sub-spec

## 목적

이 문서는 `turn-gate` core loop의 `work` phase 세부 계약을 소유합니다.

## 계약

- 이 단계는 task policy gate를 통과해 현재 flow 내부 실행 정책을 정한다.
- 사용자가 요청한 실제 작업을 진행한다.
- 작업은 파일 수정, 조사, 검증 실행, 리뷰 finding 처리, 계획 작성처럼 다양한 형태일 수 있다.
- work에 들어가기 전 current flow가 implicit default state인지 explicit `self-drive` mode인지 확인한다.
- operating state가 정해진 뒤 phase 상황에 맞는 protocol을 선택한다.
- operating state, phase protocol selection, local reference 읽기 규칙은 `phase-protocols/routes.md`가 소유한다.

## 검토 질문

- current flow가 기본 상태 또는 `self-drive` 중 하나로 좁혀졌는가?
- phase protocol을 mode와 섞어 기록하지 않았는가?
- 작업이 active flow의 work boundary 안에 머무르는가?
- 개별 task 완료를 flow completion이나 turn closure로 오해하지 않았는가?
