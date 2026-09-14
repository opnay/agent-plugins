#!/usr/bin/env node
import process from "node:process";
import { releasePlugin } from "./plugin-release-lib.mjs";

function usage(exitCode = 0) {
  const out = exitCode === 0 ? console.log : console.error;
  out(`Usage:
  pnpm release:plugin <plugin-name> --version <version> [--force]
  pnpm release:plugin <plugin-name> --bump <patch|minor|major> [--force]

Release bumps the dev plugin version and then builds the root release surface.
Run it only when a plugin is modified for the first time on next after the last main merge.

Examples:
  pnpm release:plugin rpg-kit --version 0.1.1
  pnpm release:plugin rpg-kit --bump patch --force
`);
  process.exit(exitCode);
}

function parseArgs(argv) {
  argv = argv.filter((arg) => arg !== "--");
  if (argv.includes("--help") || argv.includes("-h")) usage(0);
  const pluginName = argv[0];
  if (!pluginName || pluginName.startsWith("-")) usage(1);

  const versionIndex = argv.indexOf("--version");
  const version = versionIndex >= 0 ? argv[versionIndex + 1] : null;
  const bumpIndex = argv.indexOf("--bump");
  const bump = bumpIndex >= 0 ? argv[bumpIndex + 1] : null;

  if (versionIndex >= 0 && (!version || version.startsWith("-"))) {
    throw new Error("--version requires a value");
  }
  if (bumpIndex >= 0 && (!bump || bump.startsWith("-"))) {
    throw new Error("--bump requires patch, minor, or major");
  }
  if (version && bump) {
    throw new Error("Use either --version or --bump, not both");
  }
  if (!version && !bump) {
    throw new Error("Release requires --version or --bump");
  }

  return {
    pluginName,
    version,
    bump,
    force: argv.includes("--force"),
  };
}

try {
  const result = releasePlugin(parseArgs(process.argv.slice(2)));
  console.log(JSON.stringify(result, null, 2));
} catch (error) {
  console.error(error instanceof Error ? error.message : String(error));
  process.exit(1);
}
