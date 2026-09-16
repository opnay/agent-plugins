# Space, Type, and Surface

## Perception and Grouping

Use proximity and similarity to suggest belonging, continuity to support paths, closure to suggest a whole, and figure/ground separation to establish focus. Check that these cues express actual relationships. A decorative enclosure can imply a grouping the content does not own.

Use alignment, axes, direction, clusters, and connection geometry to make relationships readable. In tables, charts, maps, and node-link views, preserve labels, units, rows, scales, legends, directions, and grouping before adding decoration.

## Space and Density

Identify whether the artifact behaves as a document, console, canvas, map, topology, grid, split view, or another space. Give navigation, commands, content, detail, status, and background zones clear roles.

- Fit density to information volume, comparison needs, frequency, skill, and risk. More whitespace is not an automatic improvement when it hides useful comparisons.
- Use spacing, proportion, alignment, and rhythm to create an intentional scan path. Group distance should express stronger or weaker relationships.
- Overlap, relative size, perspective, light, shadow, and negative space can establish depth. Use them coherently with interaction and layer ownership.
- Give sticky, fixed, floating, and overlay regions a reason to persist. Check occlusion and access to underlying content.
- Recompose for smaller viewports; preserve meaning, reading order, and access to primary actions rather than merely shrinking everything.

## Semantic and Visual Hierarchy

Define screen, region, group, and element importance from meaning and role. A group sets the highest persistent screen-level importance and action priority its children imply. This is a semantic constraint, not a cap on pixels, font size, or component variants.

When a child needs greater persistent emphasis, reconsider its ownership or the group's hierarchy. Do not make an isolated size exception that leaves the structure misleading.

Errors, focus, selection, and progress must remain noticeable even in a low-priority region. Judge accessible interaction targets separately from visual footprint. Never weaken essential contrast, reading, focus, target size, or keyboard access to preserve a hierarchy.

Equivalent meaning, hierarchy, role, and state should use equivalent type, spacing, color, icon, component, and surface rules. Vary them for a repeatable semantic or contextual reason.

## Typography

- Choose typeface, size, weight, line height, line length, and spacing as a system. Check actual words, numbers, scripts, and languages rather than placeholder text alone.
- Use headings, labels, values, body text, and auxiliary copy to reveal structure. Optical adjustments should preserve a repeatable rule.
- Distinguish visual prominence from semantic heading structure. A large value need not become a page heading.
- Resolve long content and localization through content, layout, wrapping, or shared responsive rules before shrinking isolated items.
- Fit images, icons, quotations, attribution, and promotional copy to their role and reading sequence.

## Surface and Form

Choose flat, raised, inset, floating, glass, solid, paper-like, or band treatments from purpose, hierarchy, and product character. Radius, border, shadow, bevel, padding, and aspect ratio should read as one intentional form.

- Use depth to distinguish same-plane, floating, pressed, and layered relationships. Make clickable, draggable, editable, selected, and disabled surfaces predictable.
- Use cards, panels, bands, plain sections, list rows, or table rows for a clear grouping, comparison, independent action, emphasis, or separation role.
- Nest a surface only when it owns independent information, action, state, or a layer. For a simple subgroup, prefer spacing, alignment, typography, rows, or dividers; avoid repeating the parent's grouping with another border, radius, and shadow.
- Full outlines express boundaries. Partial lines express selection, state, direction, affiliation, or relationships. For ordinary emphasis, consider background, type, icons, labels, or a complete outline. A one-sided line needs meaningful edge placement and coherent behavior with radius and RTL changes.
- Fit materials and tactility to the context. Preserve existing elevation, radius, border, and shadow tokens; give new rules an explicit reason.
- Balance blur, transparency, texture, and shadow against contrast, focus, density, and runtime performance.

## Verification

Inspect normal and relevant exceptional states, actual content lengths and data volumes, target sizes, responsive recomposition, and themes. Check whether the same rule still explains repeated elements. Distinguish measured layout evidence from a visual hypothesis.
