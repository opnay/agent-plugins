# flow review-loop execution sub-spec

## 목적

이 문서는 active flow 안에서 review feedback, QA finding, self-review finding을 처리하는 flow-local 실행 전략을 소유합니다.

## 계약

- review-loop는 현재 active flow의 scope 안에 있는 material finding을 처리할 때만 사용한다.
- finding 하나가 correctness, regression, reliability, delivery risk에 직접 연결되면 우선 처리 대상이 될 수 있다.
- 한 loop는 하나의 bounded blocking finding에 집중한다.
- 여러 review comment, QA finding, self-review finding을 한꺼번에 처리하라는 요청은 그 자체로 하나의 review-loop가 아니다. 먼저 현재 active flow 안에서 바로 처리할 bounded blocking finding 하나를 고르거나, parent flow/discovery로 finite follow-up 후보를 만든다.
- finding 처리 뒤에는 해당 finding과 직접 연결된 verification expectation을 확인한다.
- low-value note나 speculative polish는 현재 flow를 넓히지 않고 follow-up candidate로 둔다.
- finding이 새 scope, 새 approval boundary, destructive/external action을 요구하면 current flow execution을 넓히지 않고 preparation 또는 handoff로 되돌린다.

## 검토 질문

- finding이 현재 flow scope 안의 blocking issue인가?
- 한 번에 하나의 bounded finding만 처리했는가?
- 여러 finding을 active execution으로 섞지 않고 우선순위 선택 또는 follow-up 후보로 분리했는가?
- finding 처리 뒤 직접 검증했는가?
