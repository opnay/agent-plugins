# Writing Intent Scenarios

`levels`는 전반 문구에 적용하고, 저장되는 작성물은 `writing` 보존 계약을 추가로 통과해야 합니다.

| 시나리오 | Before | light | standard | extreme |
| --- | --- | --- | --- | --- |
| 현재 상태 기록 | `처음에는 response/thinking으로 나눴다가 surface를 추가했지만 지금은 writing만 둔다.` | `저장 작성물은 writing 축이 소유합니다.` | `writing 축은 저장 작성물을 소유합니다.` | `축: writing. 대상: 저장 작성물.` |
| 검증 기록 | `테스트는 아직 실행하지 않았지만 일단 괜찮을 것 같습니다.` | `테스트는 아직 실행하지 않았습니다.` | `검증: 미실행.` | `검증: 미실행.` |
| 승인 경계 | `커밋은 하지 않았고, push와 PR도 승인되지 않았습니다.` | `커밋은 하지 않았고, push와 PR도 승인되지 않았습니다.` | `커밋 미실행. push/PR 미승인.` | `커밋: 미실행. push/PR: 미승인.` |
| 의미 맥락 | `spec을 고친 뒤 runtime을 다시 쓰고 build로 release surface를 갱신해야 합니다.` | `spec을 고친 뒤 runtime을 다시 쓰고 build로 release surface를 갱신합니다.` | `순서: spec > runtime > build.` | `순서: spec > runtime > build.` |
