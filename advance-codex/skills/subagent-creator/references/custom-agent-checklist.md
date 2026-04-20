# Custom Agent Checklist

Use this checklist before finalizing a custom agent file.

## 1. Role Boundary

- Is the recurring role narrow enough to justify one custom agent?
- Is the role clearly different from `default`, `worker`, or `explorer`?
- Would one custom agent own one stable job instead of several unrelated jobs?

## 2. File Completeness

- `name` is present.
- `description` is present.
- `developer_instructions` is present.
- Optional fields are included only when they materially improve the role.

## 3. Description Quality

- The `description` says when to use the agent.
- The wording is human-facing rather than internal-only.
- The role is distinguishable from other nearby agents.

## 4. Instruction Quality

- `developer_instructions` keep the role narrow.
- The instructions reduce drift into adjacent work.
- The instructions reflect the actual tool and sandbox expectations.
- The instructions say what the agent should prioritize.
- The instructions say what the agent should avoid widening into.
- The instructions define what kind of output the agent should return.
- The instructions are operational, not just descriptive.

## 5. Scope And Reuse

- Personal scope is used for broadly reusable private agents.
- Project scope is used for repo-specific or team-shared agents.
- The chosen scope matches the intended reuse pattern.

## 6. Invocation

- The agent can be called clearly by name from normal work.
- If nickname candidates are used, they are unique and readable.
