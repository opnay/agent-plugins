# Native Before Dependency

Bad:

```ts
import leftPad from "left-pad";

const id = leftPad(String(value), 6, "0");
```

Lean:

```ts
const id = String(value).padStart(6, "0");
```

Why: standard library behavior is clear, tested by the runtime, and avoids a package edge.
