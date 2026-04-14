# React 18 And 19 Breaking-Change Notes

Use this guide when a React architecture task also includes a framework upgrade or a behavior change that may invalidate the current component, hook, context, rendering, SSR, or test setup.

Focus on changes that materially affect structure and migration planning, not on every new feature.

## How To Use This Reference

1. Confirm the current runtime target:
   - client-only SPA
   - SSR or hydration
   - component library
   - test harness
2. Confirm the current React baseline:
   - React 17 to 18
   - React 18.2 to 18.3
   - React 18.x to 19
3. Check only the changes that affect the ownership boundary being discussed:
   - root API
   - effect behavior
   - Suspense or hydration
   - context or refs
   - test environment
   - TypeScript surface
4. Prefer architectural fixes over compatibility shims:
   - fix effect idempotence before disabling Strict Mode
   - fix stale assumptions about sync DOM timing before scattering `flushSync`
   - replace removed legacy APIs instead of wrapping them

## React 18

### Highest-Risk Changes

- `createRoot` and `hydrateRoot` replace legacy root APIs.
  - `ReactDOM.render` and `ReactDOM.hydrate` are deprecated and keep the app in React 17 behavior mode until migrated.
  - Architecture impact:
    - upgrade root entrypoints first before reasoning about concurrent behavior
    - verify test harnesses and app boot code together
- Automatic batching now applies to more async boundaries once you use `createRoot`.
  - Updates inside promises, timeouts, and native event handlers now batch by default.
  - Architecture impact:
    - code that relied on intermediate synchronous DOM reads may break
    - ownership bugs that were hidden by extra renders may surface
    - use `flushSync` only for real DOM timing requirements
- Strict Mode in development now simulates unmount and remount on first mount.
  - Effects and layout effects are created, destroyed, and created again.
  - Architecture impact:
    - non-idempotent effects, implicit subscriptions, and leaked imperative setup become visible
    - weak effect ownership is usually the real bug
- Hydration and Suspense behavior became stricter.
  - Hydration text mismatches are treated as errors instead of warnings and can fall back to client rendering up to the nearest `Suspense` boundary.
  - Suspended trees are discarded and retried instead of being committed half-finished.
  - Layout effects inside a re-suspending boundary are cleaned up and recreated.
  - Architecture impact:
    - SSR markup consistency matters more
    - layout measurement logic inside Suspense boundaries must tolerate teardown and re-setup
- Server rendering APIs changed.
  - `renderToNodeStream` is deprecated in favor of `renderToPipeableStream`.
  - `renderToString` and `renderToStaticMarkup` keep working but only have limited Suspense support.
  - Architecture impact:
    - SSR adapters and harnesses should be reviewed together with app upgrade work

### Secondary But Real Upgrade Friction

- React 18 drops Internet Explorer support and assumes modern JS features such as `Promise`, `Symbol`, and `Object.assign`.
- Test environments using `createRoot` may need explicit `act` environment support.
- TypeScript definitions got stricter.
  - A common break is that `children` must be declared explicitly in props.

### Architecture Review Questions For React 18

- Does any effect assume it runs exactly once on mount?
- Does any code depend on observing intermediate DOM between two state updates?
- Does SSR rely on React patching hydration mismatches instead of fixing markup?
- Does a store or CSS-in-JS integration need React 18 library APIs such as `useSyncExternalStore` or `useInsertionEffect`?
- Does the test setup still assume legacy root behavior?

## React 18.3 Preflight For React 19

- React 18.3 is intended as a warning-focused step before React 19.
- It is effectively React 18.2 plus warnings for APIs and patterns removed or tightened in React 19.
- Use it to surface migration work early before changing the major version.

## React 19

### Upgrade Gate

- Upgrade to React 18.3 first when possible.
  - The official React 19 upgrade guide recommends this path to surface deprecations before the major upgrade.
- The modern JSX transform is required.
  - Architecture impact:
    - older toolchains or legacy transpiler settings must be fixed before runtime debugging

### Deprecations That Still Have Surface Area

These are the cases most similar to `forwardRef`:
React 19 signals that new code should move away from them, but the symbol, package, or docs surface still exists for compatibility or migration.

- `forwardRef`
  - The React API reference marks `forwardRef` as deprecated in React 19 and says it is no longer necessary for new function components because `ref` can be passed as a prop.
  - But it still has a live API reference page and remains callable for compatibility.
  - Migration meaning:
    - stop introducing new `forwardRef` wrappers in React 19 code
    - treat existing usage as migration debt, not as an immediate runtime break
- `element.ref`
  - React 19 deprecates `element.ref` in favor of `element.props.ref`.
  - The upgrade guide explicitly says access will warn now and the field will be removed from the JSX Element type in a future release.
  - Migration meaning:
    - stop reading `element.ref`
    - treat any type surface that still exposes it as temporary compatibility, not as a green light
- `react-test-renderer`
  - React 19 deprecates the package, but it remains published on npm.
  - The official warning page says it will not be maintained and may break with new React features or React internals.
  - Migration meaning:
    - do not add new test coverage on top of `react-test-renderer`
    - plan migration toward Testing Library instead of waiting for a hard removal
- `act` from `react-dom/test-utils`
  - React 19 deprecates this import path in favor of `act` from `react`.
  - The existence of a dedicated warning page is itself a sign that the old path still exists as migration surface even though it is no longer the preferred API.
  - Migration meaning:
    - treat `react-dom/test-utils` imports as stale compatibility code
    - update shared test helpers first so the old import path stops spreading
- `useFormState`
  - The React 19 release post says the Canary-era `ReactDOM.useFormState` name was renamed and deprecated in favor of `React.useActionState`.
  - This is a release-channel drift case rather than a stable-to-stable runtime break, but it belongs in the same review bucket because old examples and experimental code can survive in a codebase long after the rename.
  - Migration meaning:
    - normalize any Canary-era form action code to `useActionState`
    - be careful when copying older blog posts, demos, or gist code

### Breaking Changes With Direct Architecture Impact

- Render errors are no longer re-thrown after React catches them.
  - Uncaught render errors are reported to `window.reportError`.
  - Caught render errors are reported to `console.error`.
  - `createRoot` and `hydrateRoot` now support `onUncaughtError` and `onCaughtError`.
  - Architecture impact:
    - production error reporting that depended on rethrow behavior needs an explicit root-level error pipeline
- Function-component `propTypes` checks and `defaultProps` support are removed.
  - `propTypes` on functions are ignored.
  - `defaultProps` for function components should move to default parameters.
  - Architecture impact:
    - shared UI APIs should not depend on runtime prop validation from React
    - function defaults should live in the component signature
- Legacy Context is removed.
  - `contextTypes` and `getChildContext` must be replaced with `createContext` and `contextType` or modern context consumers.
  - Architecture impact:
    - old class-based dependency wiring must be modernized before deeper React architecture work
- String refs are removed.
  - Move to callback refs or object refs.
  - Architecture impact:
    - imperative ownership becomes more explicit and easier to reason about
- Module pattern factories and `React.createFactory` are removed.
  - Architecture impact:
    - legacy component factories should be normalized to ordinary functions and JSX before refactoring around them
- Deprecated React DOM APIs are removed or effectively dead ends.
  - `ReactDOM.render` must be replaced with `createRoot`.
  - `react-dom/test-utils` should stop being the source of `act`; import `act` from `react`.
  - Other `test-utils` helpers are removed.
  - Architecture impact:
    - app bootstrap, tests, and custom harness helpers often need coordinated changes
- `react-test-renderer/shallow` is removed.
  - React recommends `react-shallow-renderer` directly, but also recommends reconsidering shallow rendering entirely.
  - Architecture impact:
    - tests coupled to React internals should move toward user-level testing libraries

### High-Risk Deprecations To Treat As Near-Future Breaks

- `element.ref` is deprecated in favor of `element.props.ref`.
- `react-test-renderer` is deprecated and now warns.
- UMD builds are removed.
  - Script-tag setups need an ESM-based CDN or a build step.

### TypeScript Migration Friction In React 19

- Ref callback cleanup typing is stricter.
  - Implicitly returning the assigned instance from a ref callback is no longer acceptable.
- `useRef` now requires an argument.
- `ReactElement["props"]` defaults to `unknown` instead of `any`.
- Many deprecated React types were removed or relocated.

### Architecture Review Questions For React 19

- Is any shared UI component still relying on function `defaultProps` or runtime `propTypes`?
- Is any class-based legacy context still present?
- Do tests still depend on `react-dom/test-utils`, shallow rendering, or `react-test-renderer` internals?
- Does bootstrap code still call `ReactDOM.render` anywhere, including samples or fixtures?
- Does error reporting assume render-time exceptions are rethrown?
- Does the build pipeline still depend on the old JSX transform or UMD bundles?

## Official Sources

- React 18 upgrade guide:
  - https://react.dev/blog/2022/03/08/react-18-upgrade-guide
- React 18 release post:
  - https://react.dev/blog/2022/03/29/react-v18
- React 19 upgrade guide:
  - https://react.dev/blog/2024/04/25/react-19-upgrade-guide
