#!/usr/bin/env node
console.error(`generate-release.mjs has been split.

Use build-plugin.mjs to copy src/<plugin>-dev into the root release surface:
  pnpm build:plugin <plugin-name> [--force]

Use release-plugin.mjs only for the first plugin modification on next after the last main merge:
  pnpm release:plugin <plugin-name> --bump <patch|minor|major> [--force]
  pnpm release:plugin <plugin-name> --version <version> [--force]
`);
process.exit(1);
