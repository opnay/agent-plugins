# Toolkit 003: alias.codex maintenance

## 변경사항 요약

- `git-codex` CLI에 `install`, `uninstall`, `doctor`를 제공하고 `alias.codex` 관리를 단일 구현으로 둡니다.
- `git codex`는 managed global alias가 `$HOME/.local/bin/git-codex`를 dispatch한다는 계약으로 사용합니다.

## 변경 상세

### Alias maintenance

- 목적: normal Git workflow의 검증 부담과 alias configuration 관리를 분리합니다.
- 범위: explicit install은 bundled `git-codex` binary와 expected global alias를 설치·검증합니다. uninstall은 expected alias만 제거하고 binary는 보존합니다. doctor는 read-only로 alias·binary·version을 진단합니다.
- 보존: existing alias가 expected value와 다르면 `--force` 없는 install·uninstall은 중단하고 값을 보존합니다.
- 비목표: 일반 commit·branch·push workflow에서 alias maintenance를 자동으로 실행하지 않습니다.
- 관련 표면: `git-codex.go`, tests, git skill spec/runtime, `references/alias-codex.md`, plugin spec, README.

## 검증 기준

- skill spec, runtime, reference가 같은 managed alias contract를 설명합니다.
- CLI tests cover install, doctor, preservation, force replacement, and uninstall.
- runtime frontmatter와 bundled reference link를 검증합니다.
- `git diff --check`를 통과합니다.
