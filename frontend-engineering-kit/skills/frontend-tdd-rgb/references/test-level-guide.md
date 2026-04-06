# Test Level Guide

- Component test:
  - use when behavior is mostly rendering plus local interaction
- Hook or state-level test:
  - use when the logic is complex but presentation is incidental
- Integration test:
  - use when correctness depends on several units collaborating

Prefer the lowest level that still exercises the real behavior contract.
Avoid end-to-end style setups when the risk is local and isolated.
