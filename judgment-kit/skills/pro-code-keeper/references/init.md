# Init

Use this reference at the start of every `pro-code-keeper` task.

## Scope Map

Lock the required scope before acting.

- User goal: what must be true when the work is done.
- Current evidence: files, tests, logs, UI state, dependency graph, or docs inspected.
- Owning boundary: caller, callee, domain layer, adapter, UI, storage, config, or test.
- Required behavior: the contract that must stay or change.
- Non-goals: behavior, cleanup, abstractions, or releases not requested.
- Verification signal: the smallest check that proves the contract.

If evidence is weak, inspect the real path before deciding. Ask only when the missing answer cannot be inferred safely from local context.

## Lean Meaning

Lean means smallest safe surface after adequate investigation.

- Investigation can be broad.
- Edits should be narrow.
- Review can be broad.
- Findings should be narrow and actionable.

## Allow-List First

Define what is allowed before adding block-list exceptions.

- Valid inputs, states, transitions, outputs, and side effects.
- Owner of each rule.
- Failure mode for invalid or unknown values.
- Narrow block-list safeguards only when they protect the allow-list contract.

## Evidence Order

Prefer evidence in this order:

1. Current code and call flow.
2. Tests, fixtures, logs, schemas, generated types, runtime config.
3. Local docs and specs.
4. User explanation.
5. General knowledge.

User explanation is a valuable clue, not a substitute for the current code path.
