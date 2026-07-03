# Reuse Before Rewrite

Bad:

```ts
function parseUser(raw: string) {
  const data = JSON.parse(raw);
  if (!data.id) throw new Error("missing id");
  return data;
}
```

Lean:

```ts
const user = userSchema.parse(JSON.parse(raw));
```

Why: the existing schema owns validation. Rewriting it nearby splits the contract.
