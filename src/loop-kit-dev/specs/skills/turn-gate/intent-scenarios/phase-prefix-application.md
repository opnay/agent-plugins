# phase prefix application scenario

이 시나리오는 phase prefix가 필요한 사용자-facing phase/progress message에는 붙고, artifact body나 기록 파일처럼 prefix를 전파하면 안 되는 표면에는 붙지 않는지 확인합니다.
runtime instruction이 아니라 spec-side fixture이며, turn-gate phase prefix 문구를 재생성하거나 self-drive/status/reporting 동작을 바꾸는 경우 평가 입력으로 사용합니다.

## Scenario Contract

- Expected task tier: depends on requested work; this fixture evaluates response formatting and routing.
- Expected verification method: `normal` for no-edit prefix routing checks, `clean-context` if runtime/spec/scenario files are changed.
- Primary risk: prefix underuse on phase-start/progress messages, or prefix overuse inside records, artifacts, command summaries, and question choices.
- Required behavior:
  - phase-start and meaningful phase-progress user-facing messages start with the current phase prefix.
  - status/progress answers use the current active phase unless intentionally reporting flow context.
  - self-drive automatic handoff announces the next phase with a user-facing prefix.
  - artifact bodies, flow records, self-drive records, command output summaries, and question option labels are not polluted by repeated phase prefixes.
  - blocker messages use the phase where the blocker is discovered.

## Expected Classification

| Case | Input / context | Expected behavior | Forbidden behavior |
| --- | --- | --- | --- |
| 1 | `$loop-kit:turn-gate` only, no concrete task | Start scope setup with `[preparation]`. | Return an unprefixed activation summary and close. |
| 2 | Activation-only request immediately opens actual next-flow choices | Use `[next-flow]` for the choice-opening user-facing message. | Use `[preparation]` while presenting actual next-flow choices as if still only preparing. |
| 3 | Concrete task begins with scope lock | First user-facing phase message starts with `[preparation]`. | Begin work silently with no phase prefix. |
| 4 | Work starts after preparation | Work progress message starts with `[work]` or `[work/<protocol>]` if a protocol is active. | Treat a command output summary as the only prefixed text while the phase announcement is unprefixed. |
| 5 | Verification starts after edits | Verification phase message starts with `[verification]` or `[verification/<protocol>]`. | Report verification as plain prose without a phase marker. |
| 6 | Reporting begins after pass verification | Result context starts with `[reporting]`. | Use terminal-style summary without explicit stop. |
| 7 | Reporting completed and explicit stop is absent | Reopen or continue with `[next-flow]` or self-drive continuation. | Close the turn because reporting finished. |
| 8 | User asks for status during active work | Answer starts with current active phase, usually `[work]`, then continue if self-drive remains active. | Use `[reporting]` for every status question regardless of active phase. |
| 9 | User asks for a flow-context summary rather than live progress | Use `[reporting]` only when the answer intentionally summarizes flow context. | Force `[work]` when the response is actually a report. |
| 10 | Self-drive finishes one flow and moves to the next | User-facing handoff starts with `[next-flow]` or the next phase prefix. | Silently advance the sidecar without a user-facing phase/progress prefix. |
| 11 | Self-drive updates `000-self-drive.md` | Record content stays normal markdown/frontmatter without phase prefixes on every line. | Prefix every ledger or frontmatter line with `[next-flow]`. |
| 12 | Flow record is refreshed before reporting | Flow record fields and markdown headings are not prefixed. | Add `[reporting]` to flow record headings or every evidence line. |
| 13 | Generated artifact body is the requested output | Artifact text is not polluted by operational phase prefixes unless the artifact itself needs that text. | Add `[work]` or `[reporting]` to each artifact paragraph. |
| 14 | Command output is summarized inside a phase update | The phase message may be prefixed, but raw command output or every command summary line is not mechanically prefixed. | Prefix every command output line. |
| 15 | Question tool options are shown after reporting | The message opening the question uses `[next-flow]` when appropriate; option labels stay clean. | Put `[next-flow]` inside every option label. |
| 16 | Session record access blocker occurs before result report | Blocker message uses `[reporting]`. | Use `[next-flow]` before reporting has started. |
| 17 | Session record access blocker occurs while reopening next flow | Blocker message uses `[next-flow]`. | Use stale active phase from earlier work. |
| 18 | Report-only evaluation with no edits | Gather evidence and report with current phase labels, then continue to `[next-flow]` if no explicit stop. | Treat no-edit work as reason to omit all phase prefixes and close. |
| 19 | General explanatory sentence inside an already prefixed message | Do not add another prefix to every following sentence. | Prefix every sentence or bullet mechanically. |
| 20 | Phase protocol is active for the phase | Use slash suffix, such as `[verification/review-loop]`, and do not print literal optional notation. | Print `[verification(/review-loop)]` or omit the suffix when the protocol is active. |

## Acceptance Signals

- Fresh executor can identify prefix-required user-facing messages and prefix-excluded artifact/record surfaces separately.
- Self-drive status and automatic next-flow handoff are covered as prefix-required user-facing progress messages.
- Prefix overuse inside `000-self-drive.md`, flow records, generated artifacts, command output, and question option labels is explicitly forbidden.
- Phase protocol suffix examples preserve slash syntax and avoid literal optional notation.
