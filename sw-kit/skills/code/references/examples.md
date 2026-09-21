# Examples

These examples show principles, not language-specific rules. Better code is not always shorter.

## 1. Unclear Name

Bad:

```js
function fix(x) {
  return x.trim().toLowerCase();
}
```

Better:

```js
function normalizeEmailAddress(rawEmail) {
  return rawEmail.trim().toLowerCase();
}
```

Problem: the original hides the domain. Failure: the same function may be reused for display names or payment IDs with different rules. Investigate existing email normalization, framework validators, and domain modules. Direct logic is acceptable if email normalization is truly this local contract. Tradeoff: the name is longer but prevents wrong reuse.

## 2. Excessive Function Splitting

Bad:

```js
function total(items) {
  return add(taxes(subtotal(items)));
}
```

Better:

```js
function invoiceTotal(items) {
  const subtotal = items.reduce((sum, item) => sum + item.price, 0);
  const tax = subtotal * TAX_RATE;
  return subtotal + tax;
}
```

Problem: tiny helpers hide simple flow. Failure: maintainers chase call chains to verify a simple formula. Investigate existing invoice domain rules before writing either version. Direct local code is clearer if the rule is local. Tradeoff: a few more lines keep the calculation visible.

## 3. Premature Abstraction

Bad:

```ts
interface StringTransformer { transform(value: string): string }
class UsernameTransformer implements StringTransformer { transform(v) { return v.trim(); } }
```

Better:

```ts
function normalizeUsername(rawUsername: string): string {
  return rawUsername.trim();
}
```

Problem: an interface predicts future variants. Failure: callers depend on an abstraction with no stable concept. Investigate domain naming rules and existing account modules. Direct implementation wins until multiple real contracts exist. Tradeoff: future expansion may require refactoring later, but current code stays honest.

## 4. Wrong DRY

Bad:

```js
const normalizeIdentifier = (value) => value.replace(/\s+/g, "").toUpperCase();
```

Better:

```js
const normalizeCouponCode = (value) => value.replace(/\s+/g, "").toUpperCase();
const normalizePaymentReference = (value) => value.trim();
```

Problem: similar string operations hide different business rules. Failure: payment references lose meaningful whitespace or casing. Investigate existing payment and promotion domain modules. Choose separate direct implementations because meanings differ. Tradeoff: duplication preserves independent change.

## 5. Swallowed Error

Bad:

```js
try { await saveOrder(order); } catch { return { ok: true }; }
```

Better:

```js
try {
  await saveOrder(order);
  return { ok: true };
} catch (error) {
  throw new OrderSaveError(order.id, { cause: error });
}
```

Problem: failure is reported as success. Failure: users think an order exists when it was not saved. Investigate repository error conventions and transaction helpers. Reuse the existing error model if compatible. Tradeoff: callers must handle a real failure.

## 6. I/O Mixed With Core Logic

Bad:

```js
async function discountedTotal(path) {
  const cart = JSON.parse(await fs.readFile(path, "utf8"));
  return cart.items.reduce((sum, item) => sum + item.price * 0.9, 0);
}
```

Better:

```js
function discountedTotal(cart) {
  return cart.items.reduce((sum, item) => sum + item.price * 0.9, 0);
}
```

Problem: file parsing and business logic are inseparable. Failure: tests need filesystem setup and parse errors obscure calculation bugs. Investigate existing cart parsers and schema validation. Direct pure logic is appropriate if the discount rule is local. Tradeoff: caller must own loading and validation.

## 7. Test Coupled to Implementation

Bad:

```js
expect(service.cache.size).toBe(1);
```

Better:

```js
expect(await service.loadUser("u1")).toEqual(user);
expect(fetchUser).toHaveBeenCalledTimes(1);
```

Problem: the test locks private structure. Failure: a safe cache implementation change breaks tests. Investigate public behavior and existing test helpers. Reuse repository mocking patterns. Tradeoff: less internal precision, more refactor tolerance.

## 8. Unrelated Large Refactor

Bad: a one-line validation fix also renames modules, moves files, and reformats unrelated code.

Better: change the validation path and add the regression test near it.

Problem: review and rollback become hard. Failure: unrelated move introduces regressions hidden inside a bug fix. Investigate whether the refactor is required for safety. Choose narrow change unless the broader edit is necessary. Tradeoff: known cleanup may remain for a follow-up.

## 9. Reimplementing Standard Library

Bad:

```js
function uuidLike() {
  return Date.now().toString(16) + Math.random().toString(16).slice(2);
}
```

Better:

```js
const id = crypto.randomUUID();
```

Problem: custom IDs may collide and may not meet standards. Failure: duplicate identifiers corrupt data. Investigate standard library and supported runtime versions. Use standard support when available. Tradeoff: runtime compatibility must be checked.

## 10. Wrong Utility Reuse

Bad:

```js
const normalized = StringUtils.normalizeDisplayName(paymentIdentifier);
```

Better:

```js
const normalized = normalizePaymentIdentifier(paymentIdentifier);
```

Problem: display-name rules are not payment-ID rules. Failure: payment reconciliation breaks. Investigate `StringUtils` implementation, tests, and payment domain modules. Direct or domain-owned code is better when contracts differ. Tradeoff: more code, less hidden coupling.

## 11. Large Dependency for Small Logic

Bad: add a new package to left-pad a short label.

Better:

```js
label.padStart(8, "0");
```

Problem: dependency cost exceeds logic cost. Failure: supply chain, license, or bundle size risk for trivial behavior. Investigate standard library support first. Use direct standard API. Tradeoff: less external configurability.

## 12. Dangerous Crypto or Protocol Code

Bad:

```js
const passwordHash = sha256(password + salt);
```

Better: use the repository-approved password hashing library or platform API with appropriate cost settings.

Problem: naive hashing is insecure. Failure: leaked hashes are easier to crack. Investigate security libraries, framework auth, and existing credential code. Reuse vetted implementations. Tradeoff: dependency/configuration cost is justified by security.

## 13. Wrapper Without Purpose

Bad:

```ts
class AppLogger { info(message: string) { frameworkLogger.info(message); } }
```

Better: use `frameworkLogger` directly, unless the wrapper enforces policy, context, redaction, or observability.

Problem: wrapper adds no responsibility. Failure: every logger feature requires duplicate wrapper changes. Investigate repository logging policy. Direct framework API is acceptable when it is the adopted surface. Tradeoff: fewer indirection points.

## 14. Ignored Domain Module

Bad: reimplement subscription eligibility in a controller with new conditionals.

Better:

```js
if (!subscriptionPolicy.canRenew(account, plan)) {
  return forbidden();
}
```

Problem: duplicate business rules drift. Failure: one path allows renewals the other rejects. Investigate existing domain modules and tests. Reuse the domain policy if its contract matches. Tradeoff: controller depends on domain policy, which is intended ownership.

## 15. Forced Installed Dependency

Bad: use a broad date library through adapters for one ISO date parse when the standard library and repository convention already cover it.

Better:

```js
const parsed = new Date(isoTimestamp);
if (Number.isNaN(parsed.getTime())) throw new InvalidTimestampError(isoTimestamp);
```

Problem: an installed dependency can still complicate code. Failure: adapters hide invalid input or add bundle cost. Investigate standard behavior, time zone requirements, and existing date utilities. Direct standard support is fine for a narrow timestamp contract. Tradeoff: richer date features remain unavailable unless needed.
