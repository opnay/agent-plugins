import fs from "node:fs";
import path from "node:path";

const repoRoot = process.cwd();

export function buildPlugin({ pluginName, force }) {
  const { releaseName, devName, sourceDir, targetDir, sourceManifestPath, targetManifestPath } =
    resolvePluginPaths(pluginName);

  const sourceManifest = readSourceManifest(sourceManifestPath, devName);
  if (fs.existsSync(targetDir) && !force) {
    throw new Error(`Refusing to overwrite existing release directory without --force: ${releaseName}`);
  }

  fs.rmSync(targetDir, { recursive: true, force: true });
  copyRuntimeFiles(sourceDir, targetDir, devName, releaseName);
  writeJson(targetManifestPath, buildReleaseManifest(sourceManifest, devName, releaseName, sourceManifest.version));
  assertNoSpecs(targetDir);

  return {
    plugin: releaseName,
    source: path.relative(repoRoot, sourceDir),
    output: path.relative(repoRoot, targetDir),
    version: sourceManifest.version,
    files: listFiles(targetDir).map((filePath) => path.relative(repoRoot, filePath)),
  };
}

export function releasePlugin({ pluginName, version, bump, force }) {
  const { releaseName, devName, targetDir, sourceManifestPath } = resolvePluginPaths(pluginName);
  const sourceManifest = readSourceManifest(sourceManifestPath, devName);
  if (fs.existsSync(targetDir) && !force) {
    throw new Error(`Refusing to overwrite existing release directory without --force: ${releaseName}`);
  }

  const resolvedVersion = version ?? bumpVersion(sourceManifest.version, bump);

  writeJson(sourceManifestPath, {
    ...sourceManifest,
    version: resolvedVersion,
  });

  return buildPlugin({ pluginName, force });
}

export function bumpVersion(currentVersion, bump) {
  if (!["patch", "minor", "major"].includes(bump)) {
    throw new Error("--bump must be one of: patch, minor, major");
  }

  const match = /^(\d+)\.(\d+)\.(\d+)$/.exec(currentVersion);
  if (!match) {
    throw new Error(`Cannot bump non-semver version: ${currentVersion}`);
  }

  const major = Number(match[1]);
  const minor = Number(match[2]);
  const patch = Number(match[3]);

  if (bump === "major") return `${major + 1}.0.0`;
  if (bump === "minor") return `${major}.${minor + 1}.0`;
  return `${major}.${minor}.${patch + 1}`;
}

function resolvePluginPaths(pluginName) {
  if (!pluginName || pluginName.startsWith("-")) {
    throw new Error("Missing plugin name");
  }

  const releaseName = pluginName;
  const devName = `${pluginName}-dev`;
  const sourceDir = path.join(repoRoot, "src", devName);
  const targetDir = path.join(repoRoot, releaseName);
  const sourceManifestPath = path.join(sourceDir, ".codex-plugin/plugin.json");
  const targetManifestPath = path.join(targetDir, ".codex-plugin/plugin.json");

  return { releaseName, devName, sourceDir, targetDir, sourceManifestPath, targetManifestPath };
}

function readSourceManifest(sourceManifestPath, devName) {
  if (!fs.existsSync(sourceManifestPath)) {
    throw new Error(`Missing source dev manifest: ${path.relative(repoRoot, sourceManifestPath)}`);
  }

  const sourceManifest = readJson(sourceManifestPath);
  if (sourceManifest.name !== devName) {
    throw new Error(`Expected source manifest name ${devName}, got ${sourceManifest.name}`);
  }
  if (!sourceManifest.version) {
    throw new Error(`Missing source manifest version: ${path.relative(repoRoot, sourceManifestPath)}`);
  }

  return sourceManifest;
}

function readJson(filePath) {
  return JSON.parse(fs.readFileSync(filePath, "utf8"));
}

function writeJson(filePath, value) {
  fs.mkdirSync(path.dirname(filePath), { recursive: true });
  fs.writeFileSync(filePath, `${JSON.stringify(value, null, 2)}\n`);
}

function listFiles(dir) {
  const result = [];
  for (const entry of fs.readdirSync(dir, { withFileTypes: true })) {
    const fullPath = path.join(dir, entry.name);
    if (entry.isDirectory()) {
      result.push(...listFiles(fullPath));
    } else if (entry.isFile()) {
      result.push(fullPath);
    }
  }
  return result;
}

function copyRuntimeFiles(sourceDir, targetDir, devName, releaseName) {
  const runtimeRoots = [
    ".codex-plugin",
    "README.md",
    "skills",
    "assets",
    "scripts",
    ".mcp.json",
    ".app.json",
  ];

  for (const root of runtimeRoots) {
    const sourcePath = path.join(sourceDir, root);
    if (!fs.existsSync(sourcePath)) continue;
    const targetPath = path.join(targetDir, rewritePathSegment(root, devName, releaseName));
    copyPath(sourcePath, targetPath, devName, releaseName);
  }
}

function copyPath(sourcePath, targetPath, devName, releaseName) {
  const stat = fs.statSync(sourcePath);
  if (stat.isDirectory()) {
    fs.mkdirSync(targetPath, { recursive: true });
    for (const entry of fs.readdirSync(sourcePath)) {
      copyPath(
        path.join(sourcePath, entry),
        path.join(targetPath, rewritePathSegment(entry, devName, releaseName)),
        devName,
        releaseName,
      );
    }
    return;
  }

  fs.mkdirSync(path.dirname(targetPath), { recursive: true });
  const data = fs.readFileSync(sourcePath);
  if (isTextFile(sourcePath)) {
    fs.writeFileSync(
      targetPath,
      rewriteText(data.toString("utf8"), devName, releaseName, path.basename(sourcePath)),
    );
  } else {
    fs.writeFileSync(targetPath, data);
  }
}

function isTextFile(filePath) {
  return /\.(md|json|ya?ml|txt)$/.test(filePath) || path.basename(filePath) === "SKILL.md";
}

function rewritePathSegment(segment, devName, releaseName) {
  if (segment === `${devName}-guide`) return `${releaseName}-guide`;
  return segment.replaceAll(devName, releaseName);
}

function rewriteText(text, devName, releaseName, basename) {
  const releaseTitle = titleFromName(releaseName);
  const rewritten = text
    .replaceAll(`$${devName}:`, `$${releaseName}:`)
    .replaceAll(`${devName}-guide`, `${releaseName}-guide`)
    .replaceAll(devName, releaseName)
    .replaceAll(`${releaseTitle} Dev`, releaseTitle)
    .replaceAll("[DEV] ", "");
  return basename === "README.md" ? stripSourceOnlyReadmeLines(rewritten) : rewritten;
}

function stripSourceOnlyReadmeLines(text) {
  return text
    .split("\n")
    .filter((line) => !line.includes("specs/"))
    .join("\n");
}

function titleFromName(name) {
  return name
    .split("-")
    .map((part) => part.charAt(0).toUpperCase() + part.slice(1))
    .join(" ");
}

function buildReleaseManifest(sourceManifest, devName, releaseName, version) {
  const rewrittenManifest = rewriteManifestValue(sourceManifest, devName, releaseName);
  const displayName = rewrittenManifest.interface?.displayName ?? releaseName;
  return {
    ...rewrittenManifest,
    name: releaseName,
    version,
    description: stripDev(rewrittenManifest.description),
    interface: {
      ...rewrittenManifest.interface,
      displayName: stripDev(displayName),
      shortDescription: stripDev(rewrittenManifest.interface?.shortDescription ?? ""),
      longDescription: stripDev(rewrittenManifest.interface?.longDescription ?? ""),
    },
  };
}

function rewriteManifestValue(value, devName, releaseName) {
  if (typeof value === "string") return rewriteText(value, devName, releaseName, "plugin.json");
  if (Array.isArray(value)) return value.map((item) => rewriteManifestValue(item, devName, releaseName));
  if (value && typeof value === "object") {
    return Object.fromEntries(
      Object.entries(value).map(([key, item]) => [key, rewriteManifestValue(item, devName, releaseName)]),
    );
  }
  return value;
}

function stripDev(value) {
  return String(value ?? "")
    .replace(/^\[DEV\]\s*/, "")
    .replace(/\s+Dev$/, "");
}

function assertNoSpecs(targetDir) {
  const specsPath = path.join(targetDir, "specs");
  if (fs.existsSync(specsPath)) {
    throw new Error(`Release surface must not include specs/: ${specsPath}`);
  }
}
