# Component API Smells

Use this guide after deciding whether the component is library-level or product-level.

## Level Detection

Library-level signals:

- little or no domain language
- expected reuse across multiple features
- structural or interaction abstraction is the main value
- constrained flexibility is more important than workflow specificity

Product-level signals:

- feature or domain language appears in the name and props
- one product workflow is the main reason the component exists
- semantic intent matters more than generalized reuse
- over-generalization would erase meaning

## Library-Level Smells

- uncontrolled growth of variants and slots
- escape-hatch props that bypass the intended contract
- product-specific behavior leaking into shared API
- too many styling or interaction knobs for one base component

## Product-Level Smells

- generic names where feature meaning should exist
- callbacks that expose low-level internal transitions instead of semantic events
- booleans that encode domain states or feature modes
- attempts to support hypothetical future reuse with broad config props

## Boolean Prop Rule

Booleans are less risky when they represent simple independent toggles such as disabled or loading.
Booleans are risky when they represent:

- mutually exclusive states
- hidden workflow modes
- domain meaning that should be modeled explicitly

Prefer variant or status props when one of several named states is intended.

## Composition Versus Config

Prefer composition when:

- structural flexibility is genuinely needed
- callers own meaningful substructure
- the extension points are stable and intentional

Prefer config or constrained props when:

- the component should allow only a known design-system surface
- the structure should remain stable
- open composition would weaken ownership or consistency

## Callback Boundary Rule

Good callbacks expose meaningful events.
Bad callbacks expose internal mechanics.

Prefer:

- `onSubmitOrder`
- `onRetryPayment`
- `onDismiss`

Be careful with:

- `onStateChange`
- `onBeforeOpen`
- `onAfterOpen`
- `onInternalSelectionChange`

## Promotion Test

Promote to shared or library level only when:

- the same responsibility repeats
- the same reason to change repeats
- the resulting API still makes sense without feature-specific context
- the component gains reuse without losing semantic clarity
