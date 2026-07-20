# Integration Examples

## Conflicting Findings

Two explorers disagree about a cause.

1. Compare their cited code, logs, and commands.
2. Reproduce the disputed path directly.
3. Accept only the conclusion supported by stronger evidence.
4. Report unresolved uncertainty when reproduction is unavailable.

## Failed Agent

One required agent fails while others complete.

1. Keep completed evidence.
2. Check the missing bounded scope directly or issue one narrow follow-up.
3. Do not restart unrelated lanes.
4. Mark the missing scope unverified if it remains unresolved.

## Ownership Violation

A worker changes a shared file outside its packet.

- Reject that change from automatic integration.
- Inspect the shared contract and assign one owner.
- Re-run affected verification after the main agent resolves the conflict.

## Read-To-Write Transition

Explorers identify two independent fixes.

- Do not treat read access as write permission.
- Send the explicit fix scopes back through `dispatch-subagents`, even when the prior manifest kept all writes main-owned.
- Use parallel write only after contracts and file ownership are disjoint.

## Incomplete Evidence

A result claims success without test output or another required check.

- Classify it as inconclusive and preserve partial evidence under risks and uncertainties.
- Run the missing check directly when feasible.
- Do not report the whole task complete until whole-result verification passes.
