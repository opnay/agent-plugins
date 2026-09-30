# Subsystem Test Campaign

Use only for an explicitly requested cleanup of one subsystem's whole test surface. Apply [SKILL.md](../SKILL.md#test-value) and [test-audit.md](test-audit.md) to every lane. Complete each stage before starting the next. A read-only audit request covers discovery and plans, not cutover or product repairs.

## 1. Baseline

Pin the upstream baseline revision. Record test and support LOC, production separately, and every in-scope test file's pass/fail result. Keep baseline failures in a separate list; they may be real bugs.

Done when every in-scope file has a baseline result. If execution is unavailable, record that limit and do not claim the campaign is fully validated.

## 2. Lanes and inventory

Partition by production owner, not filename prefix. Include every owned test file, cases at shared core boundaries, and QA/live-proof harness scenarios. Depending on the subsystem, lanes may cover accounts, commands, context, dispatch, inbound, outbound, persistence, transport, shared helpers, and harnesses.

Done when every file and owned QA scenario belongs to exactly one lane.

## 3. Read-only ledger

Assign narrow read-only lanes to separate agents when delegation is available and permitted; otherwise inspect them sequentially. Read every assigned test, including parameter tables, and its owners, entry points, callers, history, and CI routing. Mark each declaration with evidence:

- `R`: retain; name the contract and failure it catches. A rename or move stays `R`, with the move noted.
- `F`: retain the contract and repair the assertion, such as a multi-item negative that passes when only one forbidden item is absent.
- `C`: consolidate; name the owner that absorbs the assertion first, including a sibling table, stronger boundary suite, or shared package owner.
- `D`: delete; name the proof remaining or why no independent contract exists.

An `it.each` is one declaration unless rows need different marks; then classify those rows separately. Judge assertions, not test names.

Done when every declaration has a mark and an evidence line.

## 4. Layer plan

Use the ledger as input, not an edit list. A second read-only pass finds redundant layers and corrects ledger errors. Name a keeper for every contract, the files to retire, unique assertions to carry over, and private test-only seams unlocked. Prefer real transport with a fake network over a mocked collaborator implementing the behavior.

Done when every lane has that plan and deletion candidates meet the audit evidence requirements.

## 5. Cutover

With implementation authorization, edit lane by lane. Serialize shared harness and support changes through one owner. Remove obsolete injection parameters, inspection getters, reset exports, and indirection only after their production or external contract has been checked.

Register moved suites in CI routing and inventories. Update existing shrink-only line-cap baselines for the reduction; do not raise caps to bypass a gate. Record durable test-ownership rules in the subsystem's `AGENTS.md` or owning guidance, based on demonstrated findings.

Done when each plan is applied and its keepers pass.

## 6. Preservation review

Have independent read-only reviewers compare deleted coverage with keepers, preferably one per boundary group when available. Look for contracts losing their only proof and new assertions that cannot fail or exercise unreachable paths. Disclose unavailable independent review and alternative checks.

Restore each accepted gap at its owner or reject it with source evidence. For every restored contract, deliberately mutate the production owner and confirm the keeper fails for the intended reason. Stop validation before edits and restore only your mutation byte for byte; preserve unrelated user changes.

Done when every reported gap is resolved with evidence and every restored contract has caught a mutation. If a mutation cannot safely run, report the proof as incomplete.

## 7. Product defects

A retained baseline failure is a bug report. Within authorized repair scope, fix its owner as a distinct change and prove the real user flow with a control reverting only the fix and a passing candidate on the same harness. Use a separate commit only when committing is authorized. Restore control edits afterward. Record unrelated discrepancies as follow-ups.

Done when every repaired defect has failing-control and passing-candidate evidence. Unrepaired defects remain explicit findings.

## 8. Reconcile and hand off

For a long campaign, incorporate current upstream through repository Git policy; merging avoids rewriting a long multi-commit campaign when policy permits. If upstream changes a retired file, carry its new contract into the keeper rather than restoring the redundant layer. Confirm every new regression has an owner.

Rerun the full subsystem suite and required live proof on the reconciled candidate. Large diffs can truncate review file lists: supply complete boundary evidence and record maintainer compatibility decisions rather than weakening gates.

Hand off with the audit report plus baseline/final test and support LOC, production separately, lanes, retired layers, keepers, preservation gaps and mutation results, and product-defect control/candidate proof. State unavailable environments, unresolved findings, required gates, and actual publication state. Do not claim completion while required proof remains missing.
