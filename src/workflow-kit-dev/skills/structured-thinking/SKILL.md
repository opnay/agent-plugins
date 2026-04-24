---
name: structured-thinking
description: Stabilize difficult or ambiguous tasks before execution by surfacing missing questions, making assumptions explicit, comparing plausible approaches, and choosing the best next path. Use when Codex needs a pre-workflow framing layer before committing to a more specific workflow or execution path.
---

# Structured Thinking

## Overview

Use this as a pre-workflow framing layer.
Apply it only while the task shape is still unstable.

<entry_gate>
### Trigger

- Enter this gate when the task is ambiguous, under-structured, or still unstable enough that the next workflow cannot be chosen safely.
- Re-enter this gate when new information materially changes the task shape, assumptions, or likely next path.

### Execution

- Stabilize the task shape by identifying missing questions, risky assumptions, and the most plausible next path.
- Restate the task in simpler terms before proposing any downstream workflow or execution path.
- Ask only the smallest set of questions that would materially change the next workflow choice.
- If the ambiguity is really about user intent, scope, constraints, or evaluation criteria rather than workflow choice, exit to `deep-interview` instead of continuing here.
- If questions are not required, make the working assumptions explicit.
- Stop once the next path can be chosen safely.

### Boundaries

- Do not enter this gate when the next workflow is already clear.
- Do not skip this gate when workflow selection is still uncertain.
- Do not stay in this workflow when the right next step is an actual user interview.
- Do not ask questions that do not materially change workflow choice, scope, or risk.
- Do not answer as if assumptions were facts.
- Do not expand the task into unnecessary branches once the key ambiguity is isolated.
- Do not expose numbered internal thought traces, branches, or revision logs.
</entry_gate>

## Output Expectations

- Keep the final response conclusion-first.
- Lead with missing questions when ambiguity still matters.
- Otherwise lead with assumptions, options, and the recommended path.
- Prefer one recommended next path unless multiple paths remain genuinely competitive.
- End with a clear next step, decision, or request for missing input.

## Failure Modes

- Over-entry: activating when the next workflow is already clear.
- Over-questioning: asking questions that do not change scope, risk, or workflow choice.
- Over-decomposition: continuing to reorganize the task after the main ambiguity is isolated.
- Late exit: staying in this workflow after the next path is actionable.
- Assumption leakage: presenting assumptions as facts.
- Option sprawl: listing many branches when one recommended path is already sufficient.

<exit_gate>
### Trigger

- Enter this gate once the task shape is stable enough to choose the next workflow or execution path safely.
- Re-enter this gate when the chosen next path becomes uncertain again.

### Execution

- State the clarified task shape, key assumptions, and recommended next path.
- If the next path is `deep-interview`, hand off immediately instead of ending with a broad advisory answer.
- Exit immediately after the next workflow or execution path is clear.

### Boundaries

- Do not remain in this workflow after the next path is clear.
- Do not continue decomposition for completeness once actionability is reached.
- Do not hand off with unresolved ambiguity that still affects workflow selection.
- Do not treat this workflow as a replacement for execution planning or end-to-end delivery.
</exit_gate>

## Example Triggers

- "이 요구사항이 좀 애매한데 어디서부터 정리해야 하지?"
- "이 문제를 어떻게 쪼개서 봐야 할까?"
- "바로 구현하지 말고 먼저 쟁점이 뭔지 정리해줘."
- "가능한 접근이 몇 개 보이는데 어떤 기준으로 골라야 하지?"
