# Token-Efficient Style Intent Scenarios

각 입력에는 하나의 기대 출력과 보존 기준을 둡니다.

| 표면 | 입력 | 기대 출력 | 핵심 보존 기준 |
| --- | --- | --- | --- |
| 응답 | `요청하신 설정 경로를 확인해 본 결과, 설정 경로가 잘못되어 있다는 사실을 확인했습니다.` | `설정 경로가 잘못됐습니다.` | 직접 결론, 자연스러운 문법 |
| 진행 상태 | `현재 Build와 Lint는 통과했고 Typecheck는 아직 대기 중이며 E2E는 실행하지 않았습니다.` | `Build·Lint 통과. Typecheck 대기. E2E 미실행.` | 세 검증 상태 분리, 병렬 술어 |
| reasoning wording | `사용자 요청은 설명입니다. 구현은 요청하지 않았습니다.` | `의도: 설명. 구현: 미요청.` | 판단 표현 압축, 범위 보존 |
| 저장 문서 | `이전에는 npm도 검토했지만 현재 운영 계약은 pnpm build:plugin advance-codex --force를 실행하는 것입니다.` | `명령: pnpm build:plugin advance-codex --force.` | 현재 계약, exact command |
| 계층 | `페이지 안에 섹션이 있고 그 안에 필드가 있습니다.` | `계층: 페이지 > 섹션 > 필드.` | ordered depth |
| 절차 | `spec을 먼저 쓰고 runtime을 작성한 다음 build합니다.` | `순서: spec > runtime > build.` | ordered sequence |
| 상태 전이 | `draft에서 review를 거쳐 merged 상태가 됩니다.` | `상태: draft > review > merged.` | ordered transition |
| 숫자 비교 | `3은 2보다 큽니다.` | `비교: 3 > 2.` | numeric comparison, context clarity |
| 승인·실행 | `커밋은 완료됐지만 push는 승인되지 않았고 PR은 생성하지 않았습니다.` | `커밋: 완료. push 승인: 미승인. PR: 미생성.` | 승인과 실행 상태 분리 |
