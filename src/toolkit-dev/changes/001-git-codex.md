# Toolkit 001: git-codex 계획

## 변경사항 요약

- Toolkit에 Git 외부 서브커맨드용 단일 실행 파일 `git-codex`를 포함합니다.
- `git codex --version`, `git codex message create`, `git codex message validate <file>`, `git codex commit <file>`을 제공합니다.
- commit 성공 시 메시지 파일을 삭제하고, commit 실패 시 재개를 위해 보존합니다.
- Git hook·plugin hook·Git alias 없이 명시적 명령으로만 실행합니다.
- PATH 노출은 별도 설치 과정으로 제공하고 plugin 설치만으로 자동 등록하지 않습니다.

## 변경 목적

현재 `$toolkit:git`은 메시지 파일 할당, 작성·readback, commit, cleanup을 개별 shell 단계로 수행합니다. 반복되는 파일 lifecycle과 commit 결과 판정을 실행 프로그램으로 옮겨 exact-path cleanup을 안정화하고 모델의 명령 조합 부담을 줄입니다.

`git-codex`는 Git mutation 권한이나 메시지 의미 판단을 확장하지 않습니다. `$toolkit:git`이 stage 범위, repository convention, 메시지 의미, commit·push 권한을 판단하고 프로그램은 승인된 commit의 기계적 실행과 임시파일 lifecycle만 소유합니다.

## 릴리즈 노트 기준

- 독자에게 보이는 변화: bundled installer로 PATH에 별도 설치한 뒤 `git codex ...` 형식으로 메시지 파일 생성·검증과 commit·cleanup을 실행할 수 있습니다.
- 호환성 영향: 수동 message-file workflow는 fallback으로 유지하되 프로그램과 같은 lifecycle을 적용해 검증 실패 시 정리하고 commit 실패 시 보존합니다.
- migration 필요 여부: `git codex` 사용자는 `git-codex`를 PATH의 안정된 위치에 명시적으로 설치해야 합니다.
- 운영자/사용자가 알아야 할 검증 결과: 이 문서는 구현 계획입니다. 실행 파일·설치·runtime 검증 결과는 후속 구현 후 기록합니다.

## 기록 근거

- 관련 커밋: 후속 구현 시 기록합니다.
- 관련 세션 기록: 2026-09-18 Toolkit Git 메시지 lifecycle 실행 프로그램 설계 요청.

## 변경 상세

### 실행 프로그램과 Git 서브커맨드

- 목적: hook이나 shell alias에 의존하지 않는 명시적 Toolkit Git 명령을 제공합니다.
- 범위:
  - 실행 파일 이름은 `git-codex`입니다. Git은 PATH의 `git-codex`를 `git codex ...`로 호출합니다.
  - 프로그램 원본은 `toolkit:git`이 소유하는 `toolkit/skills/git/scripts/`에 둡니다.
  - 프로그램은 Git repository 안에서 실행하며 현재 process의 working directory와 Git이 확인한 repository를 사용합니다.
  - `git codex --version`은 `toolkit-git-codex <version>` 형식의 식별자를 출력하며 version은 Toolkit plugin manifest version과 일치합니다.
  - `$toolkit:git`은 `alias.codex` 존재 여부를 확인하고, alias가 있으면 어떤 external command가 실행될지 추측하지 않고 `git codex` 사용을 차단합니다.
  - alias가 없더라도 Git의 external-command 검색 순서에서 실제 선택되는 canonical executable을 확인합니다. `GIT_EXEC_PATH`, `git --exec-path`, PATH를 포함한 dispatch 결과가 installer `--check`가 반환한 `~/.local/bin/git-codex`와 같은 파일이고 `git codex --version` 검증까지 통과할 때만 실행 프로그램을 사용합니다.
  - 첫 범위는 `--help`, `--version`, `message create`, `message validate`, `commit`입니다. namespace는 미래 확장을 허용하지만 다른 subcommand를 미리 약속하지 않습니다.
- 관련 표면: plugin spec, git skill spec/runtime, Toolkit README·manifest, 실행 파일과 테스트.
- 검증: PATH에 임시 설치한 `git-codex`를 `git codex --help`와 각 subcommand로 호출하고, `alias.codex` 충돌과 잘못된 version 식별자를 차단해야 합니다.

### `git codex message create`

- 목적: 예측 가능한 prefix와 안전한 권한으로 commit message 파일을 할당합니다.
- 범위:
  - macOS·Linux의 운영체제 임시 디렉터리에 `toolkit-git-message.*` 형식의 고유 파일을 생성합니다.
  - 다른 process가 먼저 점유할 수 없는 원자적 생성 방식을 사용하고 파일 mode는 `0600`으로 제한합니다.
  - 성공 시 stdout에는 후속 명령에 사용할 절대 경로만 출력합니다.
  - 생성 실패 시 경로를 추측하거나 다른 위치를 사용하지 않고 nonzero로 종료합니다.
- 관련 표면: `git-codex`, git runtime message workflow.
- 검증: 병렬 생성 시 경로가 겹치지 않고, 파일이 일반 파일이며, 권한·prefix·절대 경로 계약을 충족해야 합니다.

### `git codex message validate <file>`

- 목적: commit 전 메시지 파일의 안전성과 기계적으로 판정 가능한 형식을 검증합니다.
- 범위:
  - 인자는 정확히 하나의 `toolkit-git-message.*` 파일이어야 합니다.
  - 절대 경로의 canonical parent가 운영체제 임시 디렉터리와 같고 basename이 `toolkit-git-message.`로 시작해야 합니다.
  - `lstat` 기준 symlink가 아닌 일반 파일, 현재 effective user 소유, group·other 권한이 없는 mode를 확인합니다.
  - UTF-8 text, NUL과 단독 CR 부재, 비어 있지 않은 첫 subject line, 앞쪽 빈 줄 부재, body가 있으면 subject 다음 빈 줄 존재, literal `\n` 부재를 검사합니다.
  - 프로그램은 실패한 파일을 수정하거나 삭제하지 않고 진단과 nonzero status를 반환합니다. 호출한 `$toolkit:git`은 기존 계약대로 exact path를 cleanup하고 commit을 차단합니다.
  - `$toolkit:git`의 cleanup은 생성 직후 기록한 device·inode와 현재 `lstat` identity가 같고 안전 조건을 충족할 때만 실행합니다. identity가 바뀌거나 안전성을 확인할 수 없으면 파일을 보존하고 경로를 보고합니다.
  - stderr 진단은 `git-codex: <code>: <message>` 형식이며 최소한 `path_invalid`, `file_unsafe`, `message_invalid`를 구분합니다.
  - staged scope와 subject 의미, commit type 선택, repository·사용자 convention 우선순위, verification evidence의 충분성은 `$toolkit:git`이 판단합니다.
- 관련 표면: `git-codex`, git skill의 Commit Message·Commit Verification 계약.
- 검증: 유효 메시지, 빈 파일, binary/NUL, 잘못된 경로, symlink, 다른 prefix, 형식 오류를 각각 구분해야 합니다.

### `git codex commit <file>`

- 목적: 검증된 메시지로 commit하고 성공한 경우 exact message file을 정리합니다.
- 범위:
  - commit 직전에 파일 안전성과 기계적 메시지 검증을 다시 수행합니다.
  - commit 전 HEAD를 기록합니다. unborn branch는 별도 상태로 기록하고 조회 실패를 무시하지 않습니다.
  - unborn 상태가 아닌데 commit 전 HEAD를 확인할 수 없으면 commit을 실행하지 않고 status `1`과 `head_unavailable` 진단으로 종료합니다.
  - validation 직후 device·inode를 cleanup 대상 identity로 보관합니다.
  - `git commit -F <file>`만 실행합니다. stage, amend, `--no-verify`, push, branch 변경을 포함하지 않습니다.
  - `git commit` 종료 뒤 HEAD를 다시 조회합니다. 종료 성공과 새 HEAD가 함께 확인된 경우에만 commit 성공으로 확정합니다.
  - command가 실패하고 HEAD가 이전과 같으면 commit 실패로 확정합니다. HEAD가 바뀌었거나 조회할 수 없으면 결과 미확정으로 반환하고 메시지 파일을 유지합니다.
  - status `1`은 stderr의 마지막 줄에 `git-codex: commit_result: reason=<code> attempted=<true|false>`를 반환합니다. 최소 reason은 `head_unavailable`, `validation_failed`, `commit_failed`이며 호출자는 commit 실행 여부와 cleanup·fallback 가능 여부를 exit status만으로 추측하지 않습니다.
  - commit 성공을 확정하면 입력으로 받은 exact file만 삭제하고 commit hash와 cleanup 결과를 반환합니다.
  - unlink 직전 `lstat`의 device·inode가 validation 직후 identity와 같고 안전 조건을 계속 충족할 때만 삭제합니다.
  - commit 성공 후 cleanup만 실패하면 commit 성공과 남은 파일을 분리해 반환하며 자동 재시도를 유발하지 않습니다.
- 관련 표면: `git-codex`, git runtime 정상 workflow와 recovery reference.
- 검증: 성공·commit 실패·cleanup 실패를 독립 시나리오로 확인하고 성공한 commit을 cleanup 실패 때문에 반복하지 않아야 합니다.

종료 상태 계약:

```text
0: commit 성공, message file 삭제 성공
1: precondition·validation 또는 commit 실패, commit 미생성, reason·attempted 반환, message file 유지
2: commit 성공, message file 삭제 실패
3: commit 실행 후 결과 미확정, message file 유지
```

상태 `2`는 부분 성공입니다. 출력은 최소한 commit hash, cleanup 상태, 남은 exact path를 구분해야 하며 호출자는 HEAD를 확인한 뒤 cleanup만 판단합니다.
상태 `3`은 commit command를 실행한 뒤 결과와 HEAD 상태를 확정할 수 없는 경우입니다. 호출자는 HEAD와 저장된 message를 다시 확인하기 전 commit이나 수동 fallback을 재실행하지 않습니다.

### 설치와 갱신

- 목적: plugin cache 경로와 분리된 안정된 Git external command를 제공합니다.
- 범위:
  - installer 원본은 `toolkit/skills/git/scripts/install-git-codex`이며 sibling `git-codex`를 `~/.local/bin/git-codex`에 복사합니다.
  - 첫 release의 지원 환경은 macOS와 Linux입니다.
  - `~/.local`과 `~/.local/bin`은 `lstat` 기준 symlink가 아닌 현재 사용자 소유 directory이며 group·other 쓰기 권한이 없어야 합니다. 없는 directory는 이 조건을 충족하도록 생성하고, 기존 directory가 조건을 위반하면 설치·교체하지 않습니다.
  - 설치 파일은 검증된 대상 directory 안의 임시 파일을 완성한 뒤 mode `0755`로 원자 교체합니다.
  - byte content가 같고 현재 사용자 소유이며 mode가 정확히 `0755`인 기존 일반 파일만 idempotent success로 처리합니다.
  - 다른 일반 파일은 기본적으로 보존하고 명시적 `--force`에서만 교체합니다. symlink와 directory는 force에서도 교체하지 않습니다.
  - `--check`는 설치 directory의 안전 조건과 설치 대상이 symlink가 아닌 현재 사용자 소유 일반 파일이고 mode `0755`이며 bundled sibling과 byte content가 같은지 검증합니다. 설치·수정 없이 성공 시 검증한 canonical target path만 stdout으로 반환합니다.
  - PATH와 shell 설정은 수정하지 않고 필요한 설정만 안내합니다.
  - 설치된 실행 파일은 plugin 원본·cache·support file에 의존하지 않습니다.
  - plugin 설치·업데이트가 사용자 PATH의 기존 실행 파일을 자동 변경하지 않습니다.
- 관련 표면: installer, README 설치·갱신 안내, plugin manifest의 실제 executable surface 설명.
- 검증: 최초 설치, 동일 파일 재설치, 충돌 보존, 명시적 교체, symlink·directory 보존, PATH가 없는 환경의 직접 실행을 확인합니다.

### `$toolkit:git` workflow 통합

- 목적: 실행 프로그램이 기계적 lifecycle을 맡되 기존 판단·권한 계약을 유지합니다.
- 범위:
  - 정상 흐름은 `stage 검증 > message create > 파일 작성·expected-message readback > message validate > final staged 검증 > commit > 저장된 full message 검증`입니다.
  - `message create` 직후 파일의 device·inode를 기록하고, commit 전 cleanup은 현재 identity와 안전 조건이 일치할 때만 수행합니다.
  - expected-message readback이나 message validation이 실패하면 exact path를 cleanup하고 commit을 차단합니다. commit command 실패 때만 재개를 위해 파일을 유지합니다.
  - 실행 프로그램을 사용할 수 없어 수동 workflow로 전환해도 같은 cleanup·보존 계약을 유지합니다.
  - `alias.codex`, Git의 external-command dispatch 결과, installer `--check`, `git codex --version`으로 사용 가능성과 identity를 확인합니다. Git이 실제 선택할 canonical executable과 `--check`가 반환한 설치 경로는 같은 파일이어야 하며, 식별자의 version은 현재 Toolkit manifest version과 정확히 일치해야 합니다. 미설치·충돌·identity·version 불일치는 자동 수정하지 않고 기존 수동 workflow를 사용하거나 설치·갱신 권한을 별도로 확인합니다.
  - status `1`은 reason과 `attempted`를 해석합니다. `validation_failed`는 identity가 유지될 때 cleanup하고, `head_unavailable`은 commit 미실행 상태로 보존하며, `commit_failed`는 재개를 위해 보존합니다. `attempted=true`이면 상태 확인 전 수동 fallback을 실행하지 않습니다.
  - commit 요청은 `git-codex` 설치, push, branch 변경 권한을 만들지 않습니다.
  - post-commit full-message 검증과 local state 확인은 실행 프로그램 밖에서 `$toolkit:git`이 계속 소유합니다.
  - `git codex commit`을 시작한 뒤 nonzero가 반환되면 status와 HEAD를 해석하기 전 수동 commit으로 fallback하지 않습니다.
- 관련 표면: git skill spec/runtime, recovery reference, README·manifest usage guidance.
- 검증: 설치됨·미설치·commit 실패·cleanup 부분 성공 상황에서 권한과 다음 행동이 달라지지 않아야 합니다.

## 비목표

- Git hook, plugin lifecycle hook, Git alias를 설치하거나 변경하지 않습니다.
- 실행 프로그램이 파일을 stage하거나 push·branch·amend·force 작업을 수행하지 않습니다.
- staged diff와 메시지 의미의 정합성이나 repository별 convention을 프로그램이 대신 판단하지 않습니다.
- plugin 설치 시 PATH, shell profile, 사용자 Git config를 자동 변경하지 않습니다.
- 첫 범위에서 `push`, `branch`, `merge`, `rebase`, `tag` subcommand를 추가하지 않습니다.
- 기존 commit 이후 full-message 검증과 부분 성공 recovery를 제거하지 않습니다.

## 호환성 및 마이그레이션

- `git-codex`가 없거나 설치가 허용되지 않은 환경에서는 같은 validation·cleanup·보존 계약으로 수동 message-file workflow를 유지합니다.
- `alias.codex` 충돌, 설치 경로·byte identity·version 식별 실패, unsupported platform은 mutation 전에 중단하거나 수동 workflow로 전환합니다.
- 기존 Git alias·hook·repository config는 변경하지 않습니다.
- 설치된 구버전과 plugin bundled version은 `git codex --version`과 Toolkit manifest version으로 비교하며 갱신은 명시적 installer 실행으로 수행합니다.
- 구현 언어와 패키징은 설치된 파일 하나가 macOS·Linux에서 추가 비표준 runtime 없이 실행된다는 계약을 충족해야 합니다.

## 검증 기준

- 임시 HOME·PATH·Git repository에서 실제 `git codex` 호출로 통합 테스트합니다.
- `alias.codex` 충돌, `GIT_EXEC_PATH`·Git exec path·PATH를 포함한 dispatch 해석, installer `--check`와 동일 executable 결합, 올바른 `--version`, 다른 version·shadow 실행 파일 식별 실패를 검증합니다.
- message create의 원자적 생성, 고유 경로, mode `0600`, stdout 계약을 검증합니다.
- message validate의 canonical temp parent, basename, owner·mode·symlink, UTF-8·NUL·CR·subject/body·literal escape와 비파괴 실패를 검증합니다.
- commit 성공 시 정확한 message와 새 HEAD를 확인하고 exact file 삭제를 검증합니다.
- commit 전 HEAD 조회 실패 시 commit 미실행, status `1`, `head_unavailable`, `attempted=false`를 검증합니다. validation 실패와 commit·hook·signing 실패는 reason·attempted 값, HEAD, message file disposition을 각각 검증하고, commit 실행 후 command 결과와 HEAD가 상충하거나 조회에 실패하면 status `3`과 재실행 금지를 검증합니다.
- cleanup 전 file identity 변경과 unlink 실패에서 commit 보존, status `2`, hash·남은 경로 출력을 검증합니다.
- installer의 directory 생성, 기존 directory의 symlink·owner·group/other write 차단, atomic copy, mode `0755`, byte·owner·mode-equal idempotency, `--check`, target collision·symlink·directory 보존, 명시적 교체, plugin cache 비의존을 검증합니다. `~/.local/bin`만 PATH에서 제외한 환경에서는 설치 파일을 절대 경로로 직접 실행하고 installer가 PATH를 변경하지 않았는지 확인합니다.
- git skill spec 변경 후 runtime `SKILL.md`를 현재 spec 기준으로 재작성하고 `quick_validate.py`를 실행합니다.
- plugin spec, README, manifest가 guidance-only가 아닌 실제 executable surface와 설치 경계를 정확히 설명해야 합니다.
- manifest JSON parsing, 실행 권한, bundled file 존재, runtime reference, `git diff --check`를 확인합니다.

## 정식 규칙 승격 여부

- normative spec 또는 runtime 문서로 승격할 지속 규칙: `git-codex` command·설치 계약, message file 안전성, commit·cleanup 상태, `$toolkit:git`과 프로그램의 책임 경계.
- change spec에만 남길 release/migration 이력: 수동 lifecycle에서 실행 프로그램으로 전환한 배경, 설치 도입 과정, 구현·검증 결과와 한계.
