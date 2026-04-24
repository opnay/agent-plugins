#!/usr/bin/env node
import process from "node:process";
import { buildPlugin } from "./plugin-release-lib.mjs";

function usage(exitCode = 0) {
  const out = exitCode === 0 ? console.log : console.error;
  out(`Usage:
  pnpm build:plugin <plugin-name> [--force]

Build copies src/<plugin-name>-dev into the root release surface without changing version.
Run it after every plugin change that should refresh the root surface.

Examples:
  pnpm build:plugin rpg-kit
  pnpm build:plugin rpg-kit --force
`);
  process.exit(exitCode);
}

function parseArgs(argv) {
  argv = argv.filter((arg) => arg !== "--");
  if (argv.includes("--help") || argv.includes("-h")) usage(0);
  const pluginName = argv[0];
  if (!pluginName || pluginName.startsWith("-")) usage(1);
  return {
    pluginName,
    force: argv.includes("--force"),
  };
}

try {
  const result = buildPlugin(parseArgs(process.argv.slice(2)));
  console.log(JSON.stringify(result, null, 2));
} catch (error) {
  console.error(error instanceof Error ? error.message : String(error));
  process.exit(1);
}
