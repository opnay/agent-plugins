# Delete Before Add

Bad:

```ts
const enableNewCheckout = true;

if (enableNewCheckout) {
  renderCheckout();
} else {
  renderLegacyCheckout();
}
```

Lean:

```ts
renderCheckout();
```

Why: a permanent single-value flag is not a feature toggle. Delete the branch when rollback is no longer a supported behavior.
