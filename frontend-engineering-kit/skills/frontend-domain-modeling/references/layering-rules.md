# Layering Rules

Use this guide when the main difficulty is deciding whether logic is domain, application, presentation, or transport work.

## Quick Test

Ask:

1. If the UI changed completely, would this logic still exist?
2. If the backend endpoint changed but the product rule stayed the same, would this logic still exist?
3. Is this logic deciding what the business means, or only how the current screen behaves?
4. Is this logic sequencing work, or defining the rule being sequenced?

## Boundary Definitions

### Domain Layer

Owns:

- core concepts and ubiquitous language
- invariants and policy rules
- domain-safe calculations
- status transitions, permissions, validation, and eligibility rules

Good signals:

- the rule should survive screen rewrites
- the same meaning appears across several flows
- product language explains the logic better than UI language

Wrong fits:

- JSX conditions
- button-disable logic that only reflects one screen
- endpoint-specific parsing
- loading or retry sequencing

### Application Layer

Owns:

- command and use-case flow
- mutation sequencing
- orchestration across domain, API, and UI boundaries
- route, modal, wizard, or submit flow coordination

Good signals:

- this code explains how work happens, not what the business rule means
- the same domain rule may be called from several orchestrations
- retries, optimistic updates, navigation, and error recovery matter

Wrong fits:

- raw transport parsing
- durable policy rules
- display-only formatting

### Presentation Layer

Owns:

- rendering structure
- event binding
- view models and display shaping
- labels, grouping, empty states, and screen-local interpretation

Good signals:

- the logic exists to make one UI easier to read or interact with
- the same domain concept may be represented differently on another screen
- removing the current screen would remove this logic too

Wrong fits:

- permissions and eligibility rules that survive UI changes
- raw API parsing
- cross-flow orchestration

### Transport Or Adapter Boundary

Owns:

- endpoint and request details
- DTOs and raw contract mapping
- pagination or protocol metadata
- external system quirks

Good signals:

- field names come from the server contract
- the code changes when the endpoint changes
- the main problem is integration rather than product meaning

Wrong fits:

- durable product language
- frontend-only workflow rules
- presentation-ready labels or grouping

## Decision Rules

- If a rule survives both screen and endpoint changes, it probably belongs in domain.
- If the code mainly coordinates timing, retries, navigation, or command order, it probably belongs in application.
- If the code mainly shapes what one screen shows, it probably belongs in presentation.
- If the code mainly interprets transport shape, keep it near API or adapter boundaries.
- Move logic downward only when doing so protects meaning or invariants, not merely because reuse might happen later.

## Cross-Layer Smells

- Domain rules hidden in JSX or component event handlers
- DTO shape treated as the product language
- View model logic promoted into domain because it became large
- Application orchestration pushed into domain as fake "use case objects"
- Transport quirks leaking into shared UI or feature language
