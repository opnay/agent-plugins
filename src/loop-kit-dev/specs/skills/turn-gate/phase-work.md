# turn-gate phase-work sub-spec

## 목적

이 문서는 `turn-gate` core loop의 `work` phase 세부 계약을 소유합니다.

## 계약

- 이 단계는 task policy gate를 통과해 현재 flow 내부 실행 정책을 정한다.
- 사용자가 요청한 실제 작업을 진행한다.
- 작업은 파일 수정, 조사, 검증 실행, 리뷰 finding 처리, 계획 작성처럼 다양한 형태일 수 있다.
- work에 들어가기 전 current-phase work의 internal mode를 하나 선택한다.
- mode selection과 local reference 읽기 규칙은 `mode-selection.md`가 소유한다.

## 검토 질문

- current-phase work가 하나의 internal mode로 좁혀졌는가?
- 작업이 active flow의 work boundary 안에 머무르는가?
- 개별 task 완료를 flow completion이나 turn closure로 오해하지 않았는가?
