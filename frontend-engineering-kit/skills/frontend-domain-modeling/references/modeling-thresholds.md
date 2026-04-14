# Modeling Thresholds

Use this guide when deciding whether a concept should remain plain data or become an explicit model.

## Threshold Questions

Ask:

1. Does the concept carry a real invariant or policy?
2. Does identity matter to the frontend's reasoning?
3. Is normalization or validation repeated often enough to deserve one owner?
4. Would an explicit model make the code easier to explain, not only more formal?

If the answers are mostly no, keep the data simple.

## Keep Plain Data When

- the data is mostly passed through to UI
- there is no important invariant beyond shape
- the logic is screen-local and not reused by meaning
- a wrapper object would only mirror a DTO

## Introduce A Value-Like Model When

- validation, normalization, or comparison rules matter
- the concept is defined by its contents, not identity
- repeated mistakes happen when the raw shape is passed around freely
- domain-safe helpers improve clarity

Common examples:

- money
- date range
- filter criteria
- permission set
- validated form input

## Introduce An Entity-Like Model When

- identity and lifecycle matter
- the same concept appears across several flows
- status transitions or permissions depend on the concept
- several behaviors cluster around one durable business object

Common examples:

- order with lifecycle transitions
- account with entitlement or eligibility rules
- draft object with save, publish, or archive behavior

## Prefer Policy Modules When

- the main value is a decision, not a rich data holder
- several screens ask the same yes or no question
- rules combine multiple inputs into one product decision

Common examples:

- can-submit rules
- upgrade eligibility
- allowed status transition checks
- field requirement logic

## Overmodeling Smells

- class or object wrappers that add no protected meaning
- entity vocabulary applied to one-screen-only form state
- "domain" folders full of DTO mirrors
- use-case or service objects that only forward API calls

## Undermodeling Smells

- validation duplicated across forms and submit handlers
- permissions hidden in JSX
- raw enum strings compared all over the app
- lifecycle rules copied into event handlers and reducers
