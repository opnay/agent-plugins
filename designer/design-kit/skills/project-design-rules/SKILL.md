---
name: project-design-rules
description: Discover, establish, record, and maintain a project's design direction, foundations, and reusable patterns. Use when users ask to document design conventions, accumulate project-specific patterns, reconcile rules, or update adopted decisions and repetition behavior. project design rules, design pattern library, design conventions, project design foundations, pattern maintenance
---

# Project Design Rules

Keep project-specific design knowledge discoverable, scoped, and reusable. Own the record and its adoption status; take specialist design decisions and evidence as input.

## Find Existing Knowledge

Read and follow [project context](references/project-context.md). Identify the target project, owned design documentation, relevant instructions, tokens, components, representative artifacts, and the request's write scope.

Use adopted rules that match the task. Distinguish documented decisions from repeated implementation, experiments, and assumptions. When sources conflict, identify the difference and effect rather than silently treating either source as definitive.

## Organize the Record

Keep three kinds of knowledge distinct:

- **Direction:** audience, goal, intended impression, brand cues, expression priorities, and rationale.
- **Foundations:** color roles, typography, spacing, grids, spatial model, surfaces, state language, and motion. Link existing tokens and components; do not invent values to fill a template.
- **Patterns:** application situation, rule, reason, exceptions, representative example, evidence, and status.

For repetition-related patterns, read `$design-kit:design-base/references/repetition.md`. Record context, repeated elements, preserved meaning and role, and conditions for omission, reintroduction, emphasis, or variation. Cover terms, colors, motifs, form, layout, and interaction as relevant, not only repeated copy.

Use [the project design template](templates/project-design.md) when it helps establish a new record. Adapt to existing documentation; do not replace its organization just to match the template.

## Adopt and Evolve

Use these statuses:

- `candidate`: a proposed or experimental rule without adoption evidence.
- `adopted`: an explicit user decision or an already established project convention with evidence of adoption.
- `deprecated`: a retired rule with its replacement or retirement reason.

Give each rule its real scope: project, medium, screen family, component, or state. One observed screen is evidence, not automatic justification for a global rule. Record verified applications as evidence without confusing them with user adoption.

Classify the current work as applying, extending, or changing a rule. Preserve adopted meaning while evaluating candidates. When an authorized update changes a rule, retain its reason, affected scope, evidence, and replacement relationship. Keep unresolved conflicts visible and preserve unrelated decisions.

## Write Within Scope

For requests to establish, record, or update rules, edit the relevant entries in the owned project documents. Prefer their existing location and format. If none exists, use `docs/design.md` in the target project, after checking for existing content. Use the project's existing documentation navigation to support later discovery.

For read-only reviews or ordinary design requests without rule-writing scope, present reusable decisions as proposals. Do not create persistent documents merely because a pattern was noticed.

Keep project data out of installed plugin files, unrelated projects, and external wikis. Explain conflicts with essential readability, accessibility, truthful state, or functional meaning and propose workable alternatives.

## Verify and Report

Read back affected entries and their evidence links. Check that scope, status, examples, terminology, and related rules agree. Ensure retired and adopted rules cannot be mistaken for simultaneous defaults.

Report the exact locations read or written, rules applied or proposed, adoption evidence, changed scope, and unresolved conflicts. Say when nothing was saved. Keep a small record small; split it only when real content and retrieval needs justify that change.
