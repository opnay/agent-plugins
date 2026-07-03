# Impact Scoreboard

Use this reference when many review or audit findings need ranking.

## Score

Score each candidate from 0 to 2.

- `remove`: amount of code, dependency, or concept removed
- `risk`: user or maintenance risk reduced
- `confidence`: proof that behavior is unnecessary or replacement is equivalent
- `cost`: low migration and verification cost
- `locality`: change stays inside one owner or module

Total higher scores rank first. If a high score depends on weak evidence, lower `confidence`.

## Tie Breakers

Prefer findings that:

- remove public confusion
- delete a dependency
- simplify a shared path
- reduce future branching
- have a clear test signal

Defer findings that:

- require product decisions
- cross ownership boundaries
- need broad migrations
- touch security, auth, data loss, money, accessibility, or public APIs

## Output

Use a compact table:

`rank | tag | location | target | replacement | score | verify`

Keep the table short unless the user asked for exhaustive audit output.
