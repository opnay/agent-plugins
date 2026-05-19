# turn-gate work phase sub-spec

## 목적

이 문서는 `turn-gate` core loop의 `work` phase 세부 계약을 소유합니다.

## 계약

- 이 단계는 task policy gate를 통과해 현재 flow 내부 실행 정책을 정한다.
- 사용자가 요청한 실제 작업을 진행한다.
- 작업은 파일 수정, 조사, 검증 실행, 리뷰 finding 처리, 계획 작성처럼 다양한 형태일 수 있다.
- work에 들어가기 전 sibling `flow`가 산출한 flow-local strategy와 work boundary를 적용한다.
- review-loop, fix-verify-loop, broad-execution 같은 전략 판단은 `flow` skill이 소유하고, 이 phase는 active flow 안에서 그 판단을 실행 정책으로 적용한다.

## 검토 질문

- current flow에 필요한 flow-local strategy가 sibling `flow` decision으로 좁혀졌는가?
- flow-local strategy를 turn-level mode나 next-flow authority와 섞어 기록하지 않았는가?
- 작업이 active flow의 work boundary 안에 머무르는가?
- 개별 task 완료를 flow completion이나 turn closure로 오해하지 않았는가?
