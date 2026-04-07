# Hierarchy And State Rules

## Hierarchy Decision Questions

Ask these before adjusting spacing or polish:

1. What is the single primary task on this screen or section?
2. What should the user notice first?
3. Which CTA is primary, and which ones are supporting or optional?
4. Does the grouping match the order in which the user needs to decide or act?
5. Are important elements visually stronger for a reason, or only because they happen to occupy more space?

## Hierarchy Smells

- `CTA ambiguity`
  - the primary action and secondary actions have similar prominence
- `hierarchy flattening`
  - headings, body content, summaries, and actions all have similar visual weight
- `dense but directionless`
  - the screen contains many elements but gives no clear reading order
- `grouping without decision value`
  - elements are visually clustered, but the grouping does not help the user choose or understand what matters

## State Communication Questions

Ask these when reviewing stateful UI:

1. Can the user tell within a moment whether the screen is loading, empty, failed, complete, blocked, or ready?
2. If an action is unavailable, is the reason visible?
3. Are empty, error, and loading states visually and semantically distinct?
4. Does the UI explain what the user should do next after seeing the state?
5. Does state change shift attention toward the most important new information?

## State Communication Smells

- `state invisibility`
  - the current state exists in logic but is weak or unclear in the UI
- `same treatment for different states`
  - loading, empty, and error are rendered with nearly the same treatment
- `silent disablement`
  - controls are disabled without enough explanation
- `status without action`
  - the UI reports a problem or result but does not help the user decide the next step
- `color-only signaling`
  - state differences rely mainly on color

## Cross-Check

- If the state changes, re-check CTA priority.
- If a warning or blocking state appears, it should not be visually weaker than unrelated supporting content.
- If success or completion changes the next recommended action, the hierarchy should reflect that change.
