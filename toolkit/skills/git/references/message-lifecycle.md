# Commit message lifecycle

Read this reference for manual message editing, retained message files, or an unclear `git codex commit` result.

## Standard path

Use `git codex message create --stdin` for complete messages. It writes a private `MSG-…` file and validates it before printing the ID. It never commits.

Use the returned ID with `git codex commit <ID>`. The command validates the message, runs only `git commit -F <file>`, confirms a new `HEAD`, and deletes the same safe file.

## Message files

Files live in the current repository or worktree Git directory from `git rev-parse --absolute-git-dir`. Do not assume `.git` is a directory or create a nested one.

`validate` and `commit` accept only `MSG-…` files in that directory, by ID or real relative or absolute path. The directory must be discoverable, user-owned, and not group- or other-writable.

## Manual allocation

Use plain `git codex message create` only for incremental editing. Before writing, record the file device and inode. Read the completed file back, then run:

```sh
git codex message validate <returned-MSG-ID>
```

Validation checks the canonical parent, a private regular file, UTF-8, NUL, lone CR, literal `\n`, a nonempty subject, and the subject/body delimiter. It never changes or deletes the file.

If writing, readback, or validation fails, do not commit. Remove only the allocated file when its recorded identity and safety conditions still match. Preserve an uncertain file and report its path.

## Results and recovery

- `0`: commit confirmed; message deleted.
- `1`: precondition, validation, or commit failure; input remains. Read `git-codex: commit_result: reason=<code> attempted=<true|false>`. Clean `validation_failed` only when the recorded allocation identity is safe. Preserve `head_unavailable` and `commit_failed` files.
- `2`: commit confirmed but cleanup failed. Inspect `HEAD`, file identity, and the stored message; then handle only the file.
- `3`: commit outcome is unknown. Inspect `HEAD`, file identity, and the stored message before retrying or falling back.

For a clear successful result, do not add a status or stored-message query solely to verify it.
