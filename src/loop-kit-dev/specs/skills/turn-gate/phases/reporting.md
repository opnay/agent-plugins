# turn-gate reporting phase sub-spec

## 목적

이 문서는 `turn-gate` core loop의 `reporting` phase 세부 계약을 소유합니다.

## 계약

- 이 단계는 reporting gate를 통과한다.
- 결과 보고는 terminal response가 아니라 다음 flow 진행을 위한 context 정리다.
- 이번 flow에서 무엇을 준비/작업/검증했는지, 남은 불확실성과 blocker가 무엇인지 정리한다.
- residual uncertainty와 blocker가 있으면 사용자에게 보이게 포함한다.
- 계획된 flow가 소진되면 다음 flow나 작업을 받을 수 있도록 next-flow phase로 넘긴다.
- 질문 도구 사용과 next-flow reopening 세부는 `records/question-routing.md`가 소유한다.

## 검토 질문

- 보고가 terminal close가 아니라 다음 flow를 위한 continuity context로 작성됐는가?
- 남은 불확실성과 blocker를 숨기지 않았는가?
- explicit stop이 없으면 next-flow reopening으로 이어지는가?
