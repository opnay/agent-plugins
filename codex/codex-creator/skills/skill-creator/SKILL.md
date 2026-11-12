---
name: skill-creator
description: Extension for the canonical system `skill-creator`. Use whenever you would use `skill-creator`; read the base system skill and this plugin skill together so the base workflow is paired with stricter skill boundaries, plugin-owned skill guidance, and trigger description metadata rules. skill creation, skill update, skill boundary, plugin-owned skill, description metadata, passive trigger
---

# Skill Creator

Read the canonical system `skill-creator` first and apply this extension alongside it. Own skill boundaries, plugin-owned skill guidance, description metadata, and comparative behavioral validation. Keep the base scaffold with the canonical skill and domain work with its owning skill.

## Establish Ownership

Define the skill's job and meaningful non-goals before its workflow. Define a plugin's boundary before its bundled skills; keep plugin-wide routing in the manifest, README, and plugin usage guidance.

Make the runtime contract understandable without development specs, unpublished paths, or hidden sibling context. Sibling links may offer navigation but must not supply missing execution rules. Keep unrelated runtime tool selection, sequencing, and escalation policy with the owning plugin or system guidance.

## Write Selection Metadata

Treat `description` as selection guidance read before the body. Begin with the capability and when it applies.

For passive skills that need selection without a named invocation, append narrow, comma-separated plain tokens drawn from likely user phrases, artifact names, workflows, and synonyms. Use no `#` prefixes or broad unrelated terms. Direct-call skills and already clear, narrow descriptions do not need forced token tails. Write the body for runtime readers and keep plugin usage guidance on its owning surface.

## Compare With And Without The Skill

Use comparative validation for new skills, substantive behavior, boundary, trigger, or tool-procedure changes, and requested skill tests. For wording-only changes, assess existing-case impact before choosing coverage. A design or review request authorizes a validation plan, not test execution.

1. Prepare ordinary cases and cases where absence of the skill could plausibly cause failure, with at least one of each. Scale coverage and repetitions to changed contracts and risk; the minimum covers categories, not confidence.
2. Before execution, freeze each case's ID, category, tested contract, request, initial fixture, permitted and forbidden actions, and observable success criteria. For failure-hypothesis cases, record the predicted baseline failure and its rationale separately from observations.
3. Prepare independent fresh executions with the same model, settings, tools, permissions, request, initial fixture, and execution budget. Use the native loading procedure below. Keep shared instructions and resources unchanged; disclose unmatched conditions or unverified isolation rather than claiming a valid control.
4. Give executors only the task input, fixture, permitted operations, and evidence format. Hide criteria, failure hypotheses, and the other execution's results. Reset isolated fixtures for each run. Simulate live effects or mark them unrun when additional authority is needed; preserve transmission, cost, and permission boundaries.
5. Judge actual artifacts and tool evidence against the frozen criteria. Classify each execution before interpreting the pair. Record with-skill violations regardless of baseline performance.
6. When evidence attributes a violation or regression to instructions, revise the skill and rerun original cases plus affected ordinary cases. Preserve criteria. Version invalid fixture or criterion corrections, explain the change, and rerun both conditions.

### Load Through Native Discovery

- Use the current skill folder specified by the user. For a skill under revision, use its repository working files rather than an installed cache. Freeze the source path and content hash before execution; keep that source unchanged during runs.
- For each Codex condition, start a separate fresh process in an isolated workspace with new `HOME` and `CODEX_HOME` directories and empty user configuration. Provide required authentication only within existing authority. Match shared environment, tools, canonical skills, and other required instructions. A fresh subagent context alone does not establish this isolation.
- In the with-skill workspace, create a temporary `.agents/skills/<skill-folder>` symlink to the frozen target folder, including its dedicated resources. Keep native skill discovery enabled and permit discovery and body reads. Let the executor select and read the skill; do not paste its body into the prompt or read it on the executor's behalf.
- In the without-skill workspace, omit the target link. Check parent, user, administrator, plugin, and same-name skill routes for target exposure. Preserve shared canonical skills and distinguish them from the target extension by actual path.
- Record metadata discovery, selection, and reads of the target `SKILL.md` and needed resources separately. A symlink alone does not prove use. Automatic-selection cases use ordinary requests without invocation hints; judge them separately from explicit-call cases. Preserve non-selection as an observation instead of replacing it with manual body injection.
- After execution, remove only links created for that run; preserve source folders and existing configuration. If symlinks, native loading, or observable body-read evidence are unavailable, report native-loading validation incomplete.

This skill owns the comparison contract. Evaluation tools such as `$jevx:scenario-testing` are optional, not required plugins, hidden prerequisites, or reasons to impose a fixed large case count.

### Interpret Evidence

| Without skill | With skill | Interpretation |
|---|---|---|
| Failure | Success | Improvement evidence for this case |
| Success | Success | Contract met; added benefit unestablished |
| Failure | Failure | Remedy unestablished |
| Success | Failure | Regression candidate; investigate cause |

Preserve successful baselines. Do not select cases or alter criteria after execution to manufacture baseline failures.

Separate environment errors, invalid fixtures, insufficient observation, loading contamination, and unrun cases from success or failure. Exclude invalid comparisons from effect claims and state why. When independent execution or isolation is unavailable, report structural review and missing behavioral validation.

Test completion and benefit are separate claims. A single run supports only its observed case, not stability or general performance. Never average away a critical authority violation.

Direct body injection supplies supplementary evidence about behavior after loading; it does not replace native discovery, selection, and loading through the symlink. When changing the description, test automatic selection on ordinary requests and false selection on out-of-scope requests. Neither explicit invocation nor body injection establishes automatic trigger behavior.

## Review And Report

Check ownership, meaningful non-goals, plugin boundaries, runtime readability, discoverable resources, and description selection. Check passive-token relevance and absence of `#`; avoid unnecessary tails on direct-call or narrow skills.

When this extension materially affects the result, report skill boundaries, non-goals, plugin ownership, description trigger behavior, unrelated runtime tool guidance avoided or separated, validation performed, and residual risk.

For comparative validation, include:

- Target skill, version, source path and content hash; frozen cases and criteria.
- Temporary links and isolation settings; discovery, selection, body-read, and baseline non-exposure evidence; cleanup outcome.
- Both conditions' results and evidence locations; invalid or unrun cases, violations, regressions, revisions, reruns, and limits.

Distinguish plans from observed execution. Claim comparison completion only with both case categories, valid paired executions, and observable evidence. State unmet conditions and unverified behavior explicitly.
