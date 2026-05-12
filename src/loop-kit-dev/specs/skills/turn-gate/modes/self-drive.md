# turn-gate self-drive mode spec

## 목적

`self-drive` mode는 사용자 메시지 기반 preparation이 만든 planned flow sequence를 bounded autonomous continuation으로 이어갈 때 적용하는 mode입니다.
이 mode는 별도 installed skill entrypoint가 아니라 `turn-gate` 내부 mode입니다.

## 계약

- `self-drive` mode는 준비된 planned flow sequence, scope, non-goal, acceptance signal, approval boundary, verification expectation이 충분히 기록된 뒤에만 적용한다.
- 각 planned flow는 여전히 자기 내부에서 `preparation -> work -> verification -> reporting -> next-flow` core loop를 가진다.
- self-drive는 사용자 승인 없이 destructive, irreversible, external action, commit, push, PR, publish, release, version bump를 실행하지 않는다.
- 초기 협의 범위 밖의 위험 작업이나 새 approval boundary가 나타나면 implicit default state의 user-gated question routing으로 돌아간다.
- planned flow sequence가 끝나면 terminal close가 아니라 commit-readiness reporting handoff 또는 next-flow reopening으로 이어진다.
- runtime 실행 계약은 `skills/turn-gate/references/self-drive.md`에 흡수된 reference를 사용한다.

## 검토 질문

- self-drive에 필요한 planned flow sequence와 approval boundary가 기록돼 있는가?
- user-gated로 되돌려야 할 새 위험 작업이나 범위 확장이 생기지 않았는가?
- 마지막 flow 뒤 commit execution과 commit-readiness reporting을 구분했는가?
