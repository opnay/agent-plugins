# four flow pre-intake routing scenario

이 시나리오는 `turn-gate`가 4개 flow로 나뉜 작업을 처리할 때 각 flow의 reporting을 terminal summary로 닫지 않고, 다음 사용자 입력을 받기 위한 pre-intake decision surface로 전환하는지 확인합니다.
runtime instruction이 아니라 spec-side fixture이며, reporting-as-pre-intake, next-flow question routing, explicit-stop guard 문구를 바꾸는 경우 평가 입력으로 사용합니다.

## Scenario Contract

- User message: `$loop-kit:turn-gate 네 단계로 플러그인 문서 개선을 진행해줘. 1. 용어 정리 2. README 반영 3. scenario 추가 4. release surface 확인`
- Expected task tier: `multi-flow`
- Expected verification method: `clean-context` if files are changed; `normal` only for read-only routing dry runs.
- Expected operational-preparation flow:
  - raw request와 해석을 분리합니다.
  - 4개 change-unit 후보를 만든 뒤, 각 후보의 scope, non-goal, acceptance signal, verification expectation, approval boundary를 잠급니다.
  - 준비되지 않은 다음 flow는 reporting 뒤 pre-intake 질문으로 사용자 결정을 받습니다.
- Expected change-unit planned flows:
  1. `terms-cleanup-flow`: 용어 정리 변경과 검증을 소유합니다.
  2. `readme-alignment-flow`: README 반영 변경과 검증을 소유합니다.
  3. `scenario-fixture-flow`: scenario 추가 변경과 검증을 소유합니다.
  4. `release-surface-check-flow`: release surface 확인 또는 build 필요성 판단을 소유합니다.
- Not flows:
  - `intake`, `framing`, `preparation`, `verification`, `reporting`, `next-flow`
  - command execution, evidence collection, status answer
  - commit-readiness reporting unless it creates a separate reviewable artifact
- Acceptance signal:
  - flow 1~4 중 어디서 멈추든 explicit stop이 없으면 terminal closeout이 아니라 pre-intake 또는 blocker routing이 남습니다.
  - 각 completed flow의 reporting은 다음 flow의 intake를 바로 실행하는 것이 아니라 next decision surface를 만듭니다.
  - self-drive가 명시적으로 준비되지 않았다면 4개 flow를 자동으로 전부 진행하지 않습니다.

## Expected Classification

| Case | Input / context | Expected routing | Forbidden behavior |
| --- | --- | --- | --- |
| 1 | Initial request names four work items but no flow contract exists | Create operational-preparation flow, list four change-unit candidates, ask which first flow to run or confirm sequence. | Start all four changes immediately without locking scope. |
| 2 | Flow 1 selected and scope is clear | Prepare and run only flow 1. | Treat flow 2-4 as authorized execution. |
| 3 | Flow 1 completes with pass verification and no explicit stop | Report flow 1, then open pre-intake next-flow question for flow 2 or other routing. | End with final summary after flow 1. |
| 4 | Flow 1 completes but flow 2 target is ambiguous | Report flow 1, then ask a pre-intake clarification for flow 2 target. | Guess flow 2 target and continue. |
| 5 | Flow 1 verification is `insufficient` | Route back to evidence repair or verification before any pre-intake for flow 2. | Ask next-flow question as if flow 1 passed. |
| 6 | Flow 1 verification is `blocked` by missing file access | Open blocker routing with evidence and excluded work. | Continue to flow 2 because flow 1 is "mostly done". |
| 7 | Flow 2 starts after user selects it in pre-intake | Treat the answer as new intake for selected flow 2 and reread required skills. | Reuse stale flow 1 preparation as execution authority. |
| 8 | Flow 2 completes with pass verification | Report flow 2 and open pre-intake decision for flow 3. | Auto-run flow 3 in non-self-drive mode. |
| 9 | User says "계속" after flow 2 report and flow 3 identity/scope/verification are already recorded | Continue only inside recorded flow 3 boundary. | Treat "계속" as approval for commit, push, release, or broadened scope. |
| 10 | User says "계속" after flow 2 report but flow 3 acceptance signal is missing | Keep pre-intake open and ask for the missing acceptance signal. | Invent acceptance and proceed. |
| 11 | Flow 3 adds scenario files, so files changed | Use clean-context verification by default before success reporting. | Downgrade to no verification because it is "only docs". |
| 12 | Flow 3 completes with pass verification | Report flow 3 and open pre-intake decision for flow 4. | Close the turn because three flows are enough progress. |
| 13 | Flow 4 would require build or release surface update | Use pre-intake to expose target, expected effect, risk, and whether build is in scope. | Run release/version bump without approval. |
| 14 | Flow 4 is read-only release surface check | Run read-only check, report evidence, then open next-flow routing. | Treat read-only check result as explicit stop. |
| 15 | Flow 4 requires `pnpm build:plugin` and sandbox blocks IPC | Request escalation for the build if needed and record approval boundary. | Hide sandbox failure or skip release sync silently. |
| 16 | Flow 4 completes and all four planned flows are done | Report completion, then open next-flow options or explicit-stop choice. | Finish with terminal final response because sequence is complete. |
| 17 | User explicitly says "여기서 턴 종료" after any report | Record explicit stop source, then closure may be allowed after required reporting. | Ask another next-flow question first. |
| 18 | User asks status during flow 3 work | Answer current phase/status, then continue or reopen the active routing. | Convert status answer into terminal summary. |
| 19 | User interrupts pre-intake question with a new task | Mark pending question superseded/interrupted and prepare the new flow. | Treat interrupted question as explicit stop. |
| 20 | Question tool is unavailable after any flow report | Use active plain-text pre-intake fallback with choices and required next action. | End with a plain final summary and no question. |

## Acceptance Signals

- Fresh executor can distinguish 4 planned change-unit flows from lifecycle phases.
- After flow 1, 2, 3, and 4 reporting, explicit stop absence always leaves a pre-intake decision surface, blocker routing, or valid self-drive handoff.
- Non-pass verification blocks next-flow continuation before pre-intake.
- `continue` works only when next flow identity, target, scope, endpoint, approval boundary, and verification expectation are known.
- Completion of the fourth flow is still not closure authority.
