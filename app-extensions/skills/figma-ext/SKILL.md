---
name: figma-ext
description: Extend Figma inspection, import, design-to-code, code-to-design, and in-Figma editing with target-element identification, hierarchy, coordinate, viewport, responsive, and flow-first layout rules without replacing upstream Figma tool guidance. Figma layout, Figma import, Figma coordinates, Figma responsive, Figma design to code, Figma code to design
---

# Figma Extension

## Boundary

- Add layout and handoff judgment to tasks that already use an available Figma capability.
- Do not replace or bypass upstream Figma prerequisites, tool schemas, authentication, or mutation rules.
- Do not configure or install a Figma connection. If the required capability is unavailable, report the limitation before mutation.
- Do not own generic visual design judgment, branding, product quality, unrelated frontend implementation, or code-quality decisions.

## Workflow

1. Classify the direction as inspection, Figma-to-code, code-to-Figma, or in-Figma editing.
2. Resolve the exact target design element and its parent hierarchy before read, import, or write decisions.
3. Capture the target page or frame dimensions and any known mobile, tablet, desktop, or responsive variants.
4. Interpret canvas-absolute coordinates together with parent-relative coordinates, constraints, auto layout, and sibling flow.
5. Translate layout with normal document flow and grid or flex first.
6. Evaluate positioning in this order: static, sticky, fixed, absolute. Require design evidence when moving away from normal flow.
7. Separate observed design facts from implementation inferences and unresolved context.

## Target Model

- `page`: the top-level product surface or viewport family
- `section`: a page region with an independent purpose and layout zone
- `section part`: a local unit with a specific purpose and layout responsibility inside a section
- `component`: a reusable element with a defined contract or variants
- `asset`: a visual resource whose reuse or delivery matters more than layout behavior

Use the element's role and parent relationship, not only its Figma node type, when classifications overlap.

## Coordinates And Viewports

- Use absolute coordinates to locate an element on the canvas or page.
- Use relative coordinates to understand its parent frame, auto layout, constraints, and sibling flow.
- Do not choose an implementation from only one coordinate system.
- Confirm the target frame size and available viewport variants.
- When responsive variants are absent, label responsive behavior as an inference rather than a design fact.

## Layout Translation

- Keep `static` when normal document flow expresses the design.
- Use `sticky` only with evidence that the element remains pinned within a scroll context.
- Use `fixed` only with evidence that the element remains pinned to the viewport.
- Use `absolute` only for evidenced overlays, free placement, or layer composition.
- Model repeated rows, columns, and two-dimensional alignment with grid or flex before manual coordinates.
- Never copy Figma canvas coordinates directly into CSS absolute positioning.

## Failure Handling

- If the target is ambiguous, identify candidates and resolve the exact node or role before writing.
- If parent hierarchy or coordinate context is unavailable, do not finalize placement.
- If viewport evidence is unavailable, distinguish the single observed viewport from unverified responsive behavior.
- If the upstream Figma capability is unavailable, do not install or reconfigure it implicitly.

## Result Contract

Keep these distinctions visible when they affect the task:

- target element and parent hierarchy
- absolute and relative coordinate context
- frame size and verified viewport variants
- chosen flow, display, and position model
- observed design facts, implementation inferences, and unresolved gaps
