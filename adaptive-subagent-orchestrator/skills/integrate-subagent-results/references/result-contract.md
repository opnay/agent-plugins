# Result Contract

Use this reference before waiting for and normalizing subagent results.

## Required Result Shape

```text
Status:
- completed | blocked | inconclusive

Summary:
- <concise conclusion>

Claims and evidence:
- <claim>: <paths, symbols, commands, key output, or other evidence>
- Mark inference and assumptions explicitly.

Files inspected:
- <main files>

Files changed:
- <files and reasons, or none>

Validation:
- <tests, checks, or repro commands with pass/fail>

Risks and uncertainties:
- <unknowns, skipped checks, regressions, or none>

Recommended action:
- <next action for the main agent>
```

## Normalization Rules

- `completed`: the assigned scope met its completion criteria with evidence.
- `blocked`: an external dependency or missing authority prevents the assigned outcome.
- `inconclusive`: investigation ran but evidence cannot support a conclusion.
- Missing validation is not success.
- A changed file outside ownership is an ownership violation, not an accepted result.

## Integration Checklist

- Wait for every ID in `required_results`.
- Compare the result with its original task packet.
- Separate claims from evidence and inference.
- Deduplicate equivalent claims.
- Resolve conflicting claims directly.
- Inspect all changed files and ownership boundaries.
- Run every feasible `whole_result_verification` item.
- Preserve failed, skipped, and unverified checks in the final report.
