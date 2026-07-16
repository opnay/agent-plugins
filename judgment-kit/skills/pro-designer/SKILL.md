---
name: pro-designer
description: Apply product design judgment to UI, UX, interface content, color, tone, expression, branding, space, composition, surface, form, and product quality when creating, reviewing, or improving screens, flows, dashboards, forms, landing pages, settings, and app interfaces. Use when a task needs design reasoning about readability, hierarchy, user flow, interface copy, color palette, tone type, brand identity, spatial model, zone architecture, density, rhythm, visual expression, card surface, depth, materiality, state clarity, accessibility, responsiveness, trust, or visual polish. design judgment, UI design, UX design, user flow, interface content, color palette, tone type, branding, brand identity, space composition, spatial model, layout density, visual expression, card design, surface design, materiality, product design, visual hierarchy, design quality, dashboard design, landing page design
---

# Pro Designer

Apply design judgment to product screens and visual interfaces. Start with the user, task, screen type, and information priority. Use `UI`, `UX`, `Content`, and `Quality` as core criteria; add `Color`, `Tone & Expression`, `Branding`, `Space & Composition`, and `Surface & Form` only when those axes affect the task.

Follow an existing design system before introducing a new pattern. Prefer user understanding and action clarity over personal taste.

## UI

Judge whether users can read information, understand structure, and predict available actions.

- Information structure: give important information, frequent actions, risky actions, and detail paths clear priority and grouping.
- Layout: use placement, alignment, density, scan path, separation, and relationship cues to reveal the screen purpose.
- Hierarchy grouping: group elements by information meaning, role, and state before styling them.
- Visual hierarchy: use typography, spacing, color, icon scale, imagery, lines, and surface strength to express importance and action priority.
- State expression: distinguish default, selected, disabled, loading, error, success, and empty states.
- Component role: size and place buttons, links, inputs, placeholders, toggles, menus, icon buttons, and auxiliary labels according to their information or action role.
- Consistency: give elements at the same hierarchy, role, and state the same component, typography, spacing, color, icon, and surface rules. Vary them only for a real semantic, interaction, or context difference that follows a repeatable rule.
- Content adaptation: handle long copy, localization, and viewport changes through content, layout, or shared responsive rules before applying one-off size or emphasis changes to individual elements.
- Accessibility: preserve contrast, readability, focus visibility, target size, and keyboard navigation.
- Affordance: make action possibility, result, and risk predictable through icons, labels, placement, state, and feedback.
- Data expression: make axes, scales, labels, rows, columns, legends, nodes, links, directions, and groups readable before decoration.
- Immediate feedback: respond visibly to clicks, input, and transitions.
- Processing state: show that delayed work is progressing and why the user is waiting.
- Result state: distinguish completion, failure, cancellation, and partial success and expose the next action.

## UX

Judge whether users can understand their goal and reach it through a predictable flow.

- Goal fit: align the screen, flow, primary actions, and secondary actions with the user's real goal.
- Use context: account for skill level, environment, task frequency, and risk.
- User flow: connect entry, understanding, choice, action, feedback, completion, and return.
- Findability: make location, available content, next path, and detail entry clear.
- Decision support: make priorities, options, consequences, message units, and confidence clear enough to choose.
- Cognitive load: avoid requiring more memory, guessing, comparison, or interpretation than the task needs.
- Error prevention and recovery: reduce mistakes and expose the cause and recovery path after failure.
- State transition: keep loading, interruption, partial failure, permission loss, and saving from breaking the flow.
- Efficiency and learnability: let new users start and repeat users move faster.
- Trust: expose cost, permission, saved state, impact, and next action when they affect confidence.

## Content

Judge whether interface language communicates meaning, risk, result, and next action in context.

- Information priority: separate what must be read now from what can wait.
- Structural copy: use page titles, section titles, and group names to support scanning and context.
- Labels: make each item's meaning and difference from nearby items clear.
- Terminology: translate jargon, abbreviations, and internal names into user language.
- CTAs: state the action intent and expected result.
- Guidance: explain only what is useful, where it is useful, without repeating known information.
- Error copy: explain the problem cause and recovery path.
- Status copy: make success, warning, loading, and saved states support the next decision.
- Factuality: reflect real system state, results, and limits without overclaiming.
- Risk copy: state consequences and reversibility for destructive, payment, permission, and other high-impact actions.
- Empty states: explain why nothing is shown and what can happen next.
- Term consistency: use one name for one concept.
- Accessible language: prefer clear, inclusive, unambiguous wording.
- Tone and voice: fit the user's situation, risk, and emotional state before adding brand flourish.

## Color

Judge whether the color system supports meaning, hierarchy, mood, and accessibility.

- Color roles: assign clear jobs to primary, secondary, accent, neutral, background, surface, border, text, and destructive colors.
- Semantic colors: keep success, warning, error, info, selected, and disabled meanings stable.
- Hierarchy and emphasis: reserve accent color for key actions and important information without competing highlights.
- Contrast and accessibility: keep text, icons, focus, and state indicators distinguishable without relying on color alone.
- Color mood: fit pastel, vivid, muted, neutral, warm, or cool palettes to the purpose, audience, tone, and emotional context.
- Color weight: balance area, brightness, and saturation so background, surface, action, status, and hierarchy roles remain clear.
- Theme adaptability: preserve meaning and hierarchy across light, dark, high-contrast, and user themes.
- Data color: keep series, thresholds, selections, legends, comparisons, and risk ranges stable and pair color with labels, shape, or position.
- Color expression: use solids, gradients, transparency, and patterns only when they clarify data, information, surface role, or tone.
- Emotional transition: keep success, warning, error, and empty states within one product mood.
- Cultural and domain context: account for regional, industry, and audience differences in color meaning.
- Scalability: keep palette and token roles stable as states, locales, and product scope grow.

## Tone & Expression

Choose a target impression and construct it through multiple expression elements.

- Tone type: choose a target such as trust-first, friendly, premium, energetic, technical, calm, or playful.
- Tone fit: match the tone to screen purpose, user risk, task frequency, and information complexity.
- Tone construction: combine color, brightness, saturation, contrast, spacing, surface, typography, icons, and copy attitude.
- Trust-first tone: use stable contrast, whitespace, readable surfaces, and predictable CTAs instead of relying on dark color.
- Energetic tone: keep saturation, motion, and strong contrast from weakening hierarchy or accessibility.
- Premium tone: use restrained saturation, space, refined surfaces, and controlled decoration without losing useful density.
- Technical tone: make axes, tables, values, states, and units readable before decoration.
- Tone boundaries: prefer functional clarity when expression weakens state meaning, accessibility, risk communication, trust, or action predictability.
- Tone consistency: keep hero, CTA, cards, tables, charts, empty, error, and success states within one product language.

## Branding

Judge whether the interface is recognizable as a specific brand without weakening product meaning.

- Brand identity: use color, logo, product name, visual motifs, and copy attitude to create recognition.
- Asset intensity: fit logos, symbols, graphics, patterns, and illustrations to screen purpose and information density.
- Brand color role: keep brand, action, status, risk, and background roles distinct.
- Brand consistency: carry one brand language across navigation, CTA, cards, forms, charts, empty, error, and success states.
- Brand and tone: keep brand identity and target tone related but separately judged.
- Brand and function: never let branding obscure clickability, state, risk, permission, cost, or result.
- Brand restraint: use repeated or oversized assets only when they add recognition without weakening structure.
- Brand scalability: preserve the brand through dark mode, high contrast, localization, long content, expansion, and new states.
- Brand distinctiveness: avoid generic template expression and preserve a memorable cue tied to product value.
- Brand truthfulness: avoid visual promises the product cannot support.

## Space & Composition

Judge how the screen works as a document, console, canvas, map, topology, grid, split view, or other navigable space.

- Spatial model: make the screen's spatial model clear.
- Zone architecture: separate navigation, command, content, detail, status, and background zones by size, placement, and persistence.
- Density and rhythm: fit whitespace, line height, group distance, repeated spacing, and toolbar height to information volume and use frequency.
- Information density fit: preserve comparison value on dense screens and reduce information or actions when comprehension and error prevention need lower density.
- Spatial hierarchy: use foreground, background, sticky, fixed, floating, and overlay regions to express importance and interaction priority.
- Flow and wayfinding: make the start, reading order, next action, detail entry, and return path predictable.
- Relationship geometry: use proximity, alignment, connection, direction, clustering, and axes to reveal relationships.
- Responsive recomposition: preserve spatial roles, reading order, and primary action placement across viewport changes.

## Surface & Form

Judge whether card, panel, button, and input surfaces use depth, shape, and materiality to clarify hierarchy and interaction.

- Surface model: fit flat, raised, inset, floating, glass, solid panel, paper-like, or band treatments to the purpose and hierarchy.
- Depth: make floating, pressed, same-plane, and layered relationships predictable.
- Shape language: make radius, border, divider, outline, shadow, bevel, padding, and aspect ratio fit the component role and product character and read as one intentional form.
- Materiality: use glass, paper, plastic, soft, solid, or realistic material only when it supports the product context.
- Tactility: distinguish clickable, draggable, selected, disabled, and editable surfaces.
- Hierarchy consistency: give surfaces at the same hierarchy and role the same depth and surface rules.
- Surface necessity: use a card, panel, band, section, list row, or table row only for a clear comparison, grouping, independent action, emphasis, or background-separation role.
- Nested surfaces: use an inner surface only when it owns independent information, action, state, or layer role. For a simple subgroup, prefer spacing, alignment, typography, dividers, rows, or bands; do not repeat the same grouping job with parent and child borders, radii, or shadows.
- Separation method: let spacing, alignment, typography, color, lines, and surface treatment clarify groups without duplicating the same signal.
- Boundary and emphasis: use outlines for element boundaries and partial lines for a distinct selection, state, direction, affiliation, or relationship indicator. For ordinary emphasis, prefer background, typography, icons, labels, or a complete outline. Use a one-sided line only when that edge position carries meaning and remains coherent through radius and directional or RTL changes.
- Density fit: match surface depth, decoration, spacing, and stroke strength to information density and use frequency.
- State change: vary hover, active, selected, pressed, disabled, and loading surface and depth by a predictable rule.
- Accessibility and performance: keep shadows, blur, transparency, and texture from harming contrast, focus, readability, or runtime performance.
- Design system fit: avoid conflicts with existing elevation, radius, border, and shadow tokens; introduce a new surface rule only for an explicit reason.
- Scalability: keep surface rules stable as cards, data, viewport, themes, and content length change.

## Quality

Keep all applied axes complete, trustworthy, and durable in context.

- Completeness: cover normal, loading, empty, error, disabled, permission, and success states in one design language.
- Accessibility: preserve contrast, focus, reading order, target size, and assistive technology support.
- Consistency: apply the same rules to equivalent information, actions, and states.
- Design system fit: use existing patterns, components, tokens, and interaction rules before inventing new ones.
- Adaptability: preserve information priority and action availability across viewport, input method, and content changes.
- Context fit: match density, expression, decoration, card usage, and stroke strength to the screen and user situation.
- Expression fit: tune typography, spacing, attribution, surface, and decoration to quotation, emphasis, guidance, promotion, data, or repeated-work contexts.
- Decoration restraint: use color, gradients, lines, cards, shadows, and graphics only when they support structure, reading, or interaction prediction.
- Scalability: tolerate longer copy, more data, localization, and regional variation.
- Perceived performance: keep delay, transition, and loading treatment from making the product feel frozen.
- Product trust: preserve stability in risk, error, payment, deletion, and permission moments.
- Interaction consistency: make input, save, submit, select, and cancel respond predictably.
- Recovery quality: apply consistent prevention, undo, retry, and loss-protection rules.

## Working Rule

Apply only the axes relevant to the task. Do not turn every design request into a full checklist.

When reporting design work, name the applied axis, the design basis, and any important residual risk.
