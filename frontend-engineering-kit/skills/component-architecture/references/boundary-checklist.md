# Boundary Checklist

- Does the component mix rendering with domain or transport logic?
- Are state ownership and prop ownership easy to explain?
- Are effects colocated with the state they synchronize?
- Would extracting a hook improve clarity, or just move complexity sideways?
- Is the public API small enough to remain stable?
- Are reusable parts actually reused, or only hypothetically reusable?
