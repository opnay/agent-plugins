# Debt Ledger

Use this reference when the user asks for lean comments, simplification notes, deferred improvements, future expansion points, or debt markers.

Default mode is read-only. Save or edit a ledger only when the user asks.

## Markers

Search for:

- `lean:`
- `ponytail:`

Treat both as legacy debt-marker inputs. Do not introduce new marker names unless the user asks.

## Good Marker Shape

A useful marker states:

- current simplification
- known limit
- upgrade trigger
- upgrade path

Example:

```ts
// lean: O(n) scan is fine under 1k items; switch to indexed lookup if this becomes hot.
```

Bad markers only say "temporary", "later", "cleanup", or "optimize".

## Ledger Fields

Report:

- file
- line
- marker
- current simplification
- known limit
- upgrade condition
- upgrade path
- status: `clear`, `no-trigger`, `stale`, or `unsafe`

## Edit Rule

When asked to update markers, keep comments rare. Add a marker only when a deliberately simple choice has a real limit worth tracking.
