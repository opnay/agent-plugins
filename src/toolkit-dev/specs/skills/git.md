## 사용자 스펙 의도

- Git commit 규칙에 branch와 remote 규칙을 더할 수 있는 `$toolkit:git` skill을 사용하고 싶다.
- commit, branch, push는 완전히 분리된 skill이나 배타적 mode가 아니라 각각 선택할 수 있고 하나의 workflow로 이어지는 기능이어야 한다.
- `SKILL.md`에는 commit, branch 생성, push와 `git push origin wip:main` 같은 refspec push 치트시트를 둔다.
- 작업 도중 발생한 예외 처리와 `codex/`, `jira/prja-000` 같은 조건부 branch prefix 규칙은 runtime reference로 라우팅한다.
- 기존 `git-committer`를 이전하거나 변경하지 않고, 먼저 독립적인 신규 skill로 만들어 동작 구조를 확인한다.
- branch를 생성하거나 기존 branch ref를 재설정한 뒤 전환하는 `git switch -C <branch> <start-point>`도 치트시트에 포함한다.
- commit subject는 120자 미만으로 제한하고, 변경 성격에 맞는 commit type을 선택한다.
- staged diff만으로 부족하면 변경 리스크에 맞는 가장 좁은 supporting check를 실행하고, failure·unavailable·skip을 구분한다.
- commit message는 skipped verification과 residual risk를 숨기지 않고 literal escape·delimiter·불필요한 blank line을 포함하지 않는다.
- commit 하나에는 하나의 related change unit만 포함하고 독립적으로 검토할 변경은 분리한다.
- Toolkit에 `git codex` external command를 두고 message file 생성·기계 검증·commit 성공 뒤 cleanup을 묶고 싶다.
- `git codex`는 관리되는 global `alias.codex`를 통해 실행하고, alias install·uninstall·doctor는 별도 reference로 둔다.
- 일반 commit workflow는 `git codex message create`로 lifecycle availability를 판단한다.
- 사용자가 설치를 요청하면 installer를 실행하고, `git codex message create`가 unavailable로 실패하면 설치 필요와 다음 권한 또는 수동 workflow 선택을 알린다.
- 메시지는 파일 수정 도구보다 생성 시 quoted heredoc으로 입력하고 자동 검증하는 경로를 우선한다.
- 메시지 파일은 Git 메타데이터 디렉터리의 `MSG-…`에 저장하고 하위 디렉터리·linked worktree에서도 Git 경로 해석을 따른다.
- 사용자가 지정한 push 형식과 위치는 preflight로 다시 찾지 않고 바로 실행하며, 잘못된 호출이나 불확실한 결과에 필요한 상태만 확인한다.
- 기존 OS-temp 메시지의 하위 호환 처리를 두지 않는다.
- `alias.codex`의 설치·제거·진단은 명시적으로 요청된 maintenance workflow에서만 수행한다.
- CLI가 명확한 성공 결과를 반환하면 같은 결과를 다시 조회해 검증하지 않는다.
- runtime `SKILL.md`는 정상 실행에 필요한 짧은 지침만 두고, 조건부 lifecycle·recovery 세부는 reference로 라우팅한다.
- branch convention과 recovery reference도 현재 조건에 필요한 상태·다음 행동만 안내하고, 정상 push의 연결정보 재탐색을 요구하지 않는다.

---

# git 스킬 스펙

## 목적

`git`은 사용자 요청 범위의 commit, branch, push를 개별 기능이자 연결 가능한 하나의 Git workflow로 수행합니다.
bundled `git-codex`는 message file lifecycle과 commit 결과 판정만 기계적으로 수행하고, skill은 scope·권한·message 의미·후속 검증을 소유합니다.
고빈도 정상 흐름과 명령은 runtime `SKILL.md`에서 짧고 직접적인 지침으로 제공하고, 조건부 lifecycle·branch convention·실패 복구는 필요한 reference만 읽도록 라우팅합니다.

## 경계

- 포함:
  - repository 규칙, working tree, current branch와 failure recovery에 필요한 Git 상태 확인
  - task-scoped staging, commit message, commit 생성과 post-commit 확인
  - branch 생성·전환, 명시적으로 허용된 force-create와 start point 확인
  - current-branch push, upstream 설정, 명시적 refspec push
  - `alias.codex` install·uninstall·doctor의 명시적 maintenance workflow
  - installed `git-codex`의 message file create·validate·commit lifecycle
  - branch convention과 예외 복구 reference routing
- 제외:
  - 구현 readiness 판단과 제품 코드 변경
  - GitHub PR, release, publish, version bump와 hosting-service API 작업
  - 사용자 요청에 없는 commit, branch mutation, push 권한 추정
  - `reset --hard`, working tree를 버리는 강제 branch 전환·삭제, force push, history rewrite의 기본 실행
  - Git 전체 manual과 repository 전용 branch 이름의 보편화
  - plugin 설치에 따른 PATH, shell profile, Git config 자동 변경
  - `git-codex`를 통한 stage, branch, push, amend, hook bypass, message 의미 판단

## 처리하려는 작업 형태

- 현재 branch에서 task-scoped commit만 생성하는 작업
- branch를 만들거나 전환한 뒤 commit하고 push하는 작업
- exact branch와 start point를 확인한 뒤 `git switch -C`로 branch를 생성하거나 기존 branch ref를 재설정하고 전환하는 작업
- 기존 commit을 current branch 또는 명시된 remote branch로 push하는 작업
- local source와 remote destination이 다른 refspec push 작업
- `codex/`, `jira/prja-000` 같은 prefix가 repository 또는 사용자 규칙의 적용 조건인 작업
- commit 또는 push가 실패·중단된 뒤 현재 상태를 확인하고 안전하게 재개하는 작업
- installed `git codex` command로 message file을 생성·검증하고 task-scoped commit을 수행하는 작업
- 사용자가 명시적으로 요청한 `alias.codex` install·uninstall·doctor 작업

## 엔트리포인트 / 대표 표면

- 대표 표면: `toolkit/skills/git/SKILL.md`
- executable surface: `toolkit/skills/git/scripts/git-codex.go`, platform `bin/*/git-codex`
- 호출 방식: `$toolkit:git`
- passive trigger: Git commit, branch creation, branch switch, upstream push, refspec push, branch prefix, push recovery

## 핵심 처리 계약

1. 사용자 요청에서 허용된 commit, branch, push 단계를 각각 확정하고 한 단계의 권한을 다음 단계 권한으로 확장하지 않습니다.
2. repository의 `AGENTS.md`와 관련 운영 문서를 먼저 확인하고 runtime 일반 규칙보다 우선합니다. 규칙 충돌은 instruction priority와 현재 요청의 명시성을 따르며, 사용자가 같은 규칙을 명확히 override하지 않았다면 repository 기본값을 유지합니다.
3. mutation 전 확인은 해당 mutation에 필요한 working tree, current branch, start point처럼 사용자가 지정하지 않은 상태로 제한합니다. 사용자가 지정한 push remote·current branch 형식 또는 exact refspec은 preflight로 remote URL, source, destination, upstream을 다시 찾지 않습니다.
4. commit, branch, push를 배타적 mode로 분리하지 않고 요청에 필요한 단계만 순서대로 조합합니다.
5. `git codex`는 managed global `alias.codex`가 `$HOME/.local/bin/git-codex`를 dispatch한다는 설치 계약을 전제로 사용합니다. alias install·uninstall·doctor는 사용자가 명시적으로 요청할 때만 runtime `references/alias-codex.md`의 bounded workflow로 수행합니다. 기존 alias 값이 다르면 `--force` 없는 install·uninstall은 보존하고 중단합니다.
6. commit은 task-owned 범위만 stage하고 final status와 staged diff를 확인한 뒤 파일 기반 commit message로 생성합니다. staged verification이 불가능하거나 선택 범위와 다르면 commit을 차단하고, staged diff만으로 리스크를 확인할 수 없으면 가장 좁은 supporting check를 적용합니다. instruction priority를 따르고, 현재 사용자가 같은 message 규칙을 명확히 override하면 그 요청을, 그렇지 않으면 repository convention을 적용합니다. subject는 120자 미만으로 유지하며 더 엄격한 repository 제한이 있으면 그 값을 따릅니다. 별도 convention이 없으면 지원 type 중 가장 구체적인 type과 `type: detailed subject` 형식, bullet body를 사용합니다. commit lifecycle은 `message create`(stdin은 자동 validation, 수동 편집은 explicit validation) `> commit`을 따릅니다.
7. `alias.codex` maintenance workflow는 사용자 요청에서만 선택합니다. CLI command의 성공·실패 출력이 install·uninstall·doctor 결과를 소유하며, 정상 결과 뒤에 alias·binary·version을 다시 조회하지 않습니다.
8. create·validate·commit·push failure는 observed cause를 보존해 recovery workflow로 분리합니다.
9. 일반 branch 생성은 `git switch -c`를 사용합니다. `git switch -C`는 사용자가 force-create를 명시했거나 repository workflow가 같은 동작을 소유할 때만 사용하며, exact branch, start point, 기존 branch ref, working tree, 다른 worktree 사용 여부를 먼저 확인합니다. `-C` 권한을 `--force`, `--discard-changes`, branch 삭제 권한으로 확장하지 않습니다.
10. push는 사용자가 요청한 current-branch 형식 또는 `<local>:<remote>` exact refspec을 바로 실행합니다. 두 형식을 같은 의미로 취급하거나 일반 push 요청에서 refspec을 추론하지 않습니다.
11. `git push origin wip:main`은 local `wip`을 remote `main`으로 보내는 explicit refspec입니다. 사용자가 이 형식을 지정했고 repository 규칙이 허용할 때만 실행합니다.
12. mutation 결과는 종료 상태와 출력으로 판단하고 commit, branch, push 결과를 구분해 보고합니다. 예상한 대상과 성공·실패가 명확하면 같은 결과를 확인하는 추가 조회를 생략합니다. 호출 대상 오류, 불완전·상충 출력, 중단·결과 불명확에는 필요한 상태만 조회합니다.
13. prefix 조건이나 실패·중단 상태가 정상 치트시트만으로 해결되지 않으면 해당 runtime reference를 읽고, 현재 상태를 관찰하기 전에 mutation을 재시도하거나 자동 rollback하지 않습니다.
14. runtime `SKILL.md`는 명령, 선택 조건, 다음 행동을 우선하고, 일반 능력 설명·반복된 이유·조건부 세부 계약은 reference에 둡니다.
15. branch convention과 recovery reference는 branch collision 또는 실패 유형에 필요한 상태만 조회하게 하며, normal push의 remote·upstream·ref mapping preflight를 추가하지 않습니다.

## Workflow 선택 및 조합

- `commit`: scope 확인 > stage > staged 검증 > heredoc create·자동 validation > commit
- `alias maintenance`: 명시 요청 > reference의 install·uninstall·doctor
- `error`: 상태 관찰 > recovery 판단
- `branch > commit`: start point 확인 > branch 생성·전환 > commit
- `push`: 사용자 지정 current-branch 형식 또는 exact refspec 실행 > 결과 해석 > 결과 불명확·호출 오류에 필요한 확인
- `branch > commit > push`: 각 단계의 독립 권한과 검증을 보존한 채 연결
- 단계 실패 시 이미 성공한 앞 단계와 실패한 현재 단계를 분리해 보고하고, 성공한 mutation을 자동으로 되돌리지 않습니다.

## Cheatsheet 소유권

- `SKILL.md`는 workflow selection, task-scoped stage, stdin message create·commit, alias maintenance reference routing, branch 생성·전환, 명시적 force-create, current-branch push, refspec push를 제공합니다.
- 명령 바로 옆에는 current-branch와 refspec의 구분, mutation 범위, destructive option 제외처럼 실행 의미를 바꾸는 조건을 둡니다.
- 전체 Git flag와 subcommand를 복제하지 않고 설치된 Git의 `git <command> -h`와 repository 규칙을 우선합니다.
- message-file lifecycle, recovery, branch convention의 조건부 세부 규칙은 `SKILL.md`에 반복하지 않습니다.

## Commit Message

- subject는 `type: detailed subject` 형식을 사용하고 120자 미만으로 유지합니다. repository가 더 짧은 제한을 소유하면 더 엄격한 값을 적용합니다.
- 기본 지원 type은 다음과 같습니다.
  - `feat`: new user-facing feature
  - `fix`: bug fix
  - `refactor`: behavior-preserving code restructuring
  - `docs`: documentation-only change
  - `test`: test addition or update
  - `perf`: performance improvement
  - `style`: formatting or style-only change
  - `build`: build system or dependency change
  - `ci`: CI configuration or script change
  - `chore`: maintenance outside the above types
- staged scope에 맞는 가장 구체적인 type을 선택합니다.
- higher-priority repository 또는 user convention이 다른 type 집합을 명시하면 그 계약을 따르되, type을 임의로 새로 만들지 않습니다.
- subject는 staged scope를 구체적으로 설명하고 vague wording이나 unrelated concern을 묶지 않습니다.
- body는 실제 변경과 검증 근거를 bullet list로 설명합니다. skipped verification이나 residual risk는 body 또는 final report에 숨기지 않습니다.
- literal `\n`, 불필요한 blank line, unrelated scope, shell syntax, heredoc·EOF 같은 delimiter text를 포함하지 않습니다.
- message file 쓰기 또는 readback이 실패하거나 내용이 계약과 다르면 생성 시 기록한 identity와 현재 안전 조건이 같을 때만 exact allocated path를 cleanup하고 commit을 차단합니다.

## Git Codex Command

- `git-codex`는 Git external command로 `git codex install [--force]`, `uninstall [--force]`, `doctor`, `--version`, `message create [--stdin]`, `message validate <message>`, `commit <message>`를 제공합니다. version은 Toolkit manifest와 일치해야 합니다.
- 우선 경로는 `git codex message create --stdin <<'EOF'`입니다. quoted delimiter는 shell expansion을 막습니다. 프로그램은 EOF까지 입력을 저장하고 자동 validation을 통과한 경우에만 stdout에 짧은 `MSG-…` 식별자 하나를 출력합니다. create는 commit을 실행하지 않습니다.
- 파일은 `git rev-parse --absolute-git-dir`로 구한 현재 repository·worktree의 canonical Git 메타데이터 디렉터리에 `MSG-…`, mode `0600`으로 원자 생성합니다. 해당 directory는 현재 사용자 소유이며 group·other 쓰기 권한이 없어야 합니다. 일반 repository에서는 `.git/MSG-…`이며 하위 디렉터리, `.git` gitfile, linked worktree, 별도 Git directory, bare repository도 Git의 경로 해석을 따릅니다. 안전한 Git directory를 찾지 못하면 파일을 만들지 않고 실패합니다.
- `validate`·`commit`은 `MSG-…` 식별자를 같은 repository·worktree의 Git directory에 상대적으로 해석합니다. 현재 Git directory 안의 `MSG-…` 실제 상대·절대 파일 경로도 허용합니다. 그 밖의 파일은 거부합니다.
- 인자 없는 `message create`는 빈 파일 할당을 유지합니다. 이 보조 경로에서는 write 전 device·inode를 기록하고 파일 작성·readback·명시적 validation을 수행합니다. stdin 경로는 성공 결과가 저장·기계 검증을 보증하므로 동일한 readback·validation을 반복하지 않습니다. 파일을 이후 수정했다면 다시 검증합니다.
- stdin 읽기·쓰기·validation 실패는 commit을 차단하고 생성 당시 identity와 현재 안전 조건이 같을 때만 방금 생성한 파일을 정리합니다. 정리할 수 없으면 경로와 오류를 보고합니다. stdout에 성공 식별자를 출력하지 않습니다.
- `message validate`는 해당 canonical parent, safe private regular file, UTF-8, NUL·lone CR·literal `\\n` 부재, subject·body delimiter를 기계 검증합니다. 독립 validation 실패는 파일을 수정·삭제하지 않습니다.
- `commit`은 validation과 HEAD를 다시 확인한 뒤 `git commit -F <file>`만 실행합니다. stage, amend, push, branch mutation, hook bypass를 추가하지 않습니다. success와 new HEAD가 함께 확인된 경우에만 input file을 identity 재확인 뒤 삭제합니다.
- exit status는 `0` success and cleanup, `1` definitive precondition·validation·commit failure, `2` committed but cleanup failed, `3` attempted commit with an unknown outcome입니다. status `1`은 `reason`과 `attempted`를 machine-readable diagnostic으로 반환합니다. status `2`·`3`은 HEAD와 remaining file을 확인하기 전 commit fallback이나 retry를 실행하지 않습니다.
- `commit` command의 status `1`은 입력 파일을 보존합니다. `validation_failed`의 cleanup은 생성 때 기록한 identity를 가진 skill caller가 판단하며, command가 validation 실패 파일을 임의 삭제하지 않습니다. `create --stdin`은 자신이 생성한 파일의 실패 cleanup을 직접 소유합니다.
- 수동 workflow도 validation/readback 실패는 identity-safe cleanup, commit failure는 message file 보존이라는 같은 lifecycle을 유지합니다.

## Alias Codex Maintenance

- entry: 사용자가 `alias.codex` install·uninstall·doctor를 명시적으로 요청한 경우입니다.
- exact alias value, binary installation, version dispatch, force and preservation rules는 `git-codex` CLI가 소유하고 runtime `references/alias-codex.md`는 해당 command를 라우팅합니다.
- 일반 commit·branch·push workflow는 이 maintenance workflow를 선택하지 않습니다.

## Commit Verification

- commit 전에 `git status --short`와 `git diff --staged`로 범위를 확인합니다. 이후 변경·중간 작업으로 관찰이 낡았을 때 다시 확인하고, 상태가 그대로인 명확한 결과는 반복 조회하지 않습니다.
- staged verification이 불가능하거나 intended change unit과 다르면 commit을 차단합니다.
- staged diff만으로 리스크를 확인할 수 없으면 변경 범위에 맞는 가장 좁은 deterministic check를 실행합니다.
  - docs: staged readback, formatting, whitespace
  - code: 변경 behavior에 맞는 lint, typecheck, test, build
- supporting check가 실패하면 수정 후 재실행하거나 blocking issue로 보고합니다.
- supporting check가 unavailable이면 이유와 residual risk를 기록하고 task risk가 허용할 때만 계속합니다.
- skip은 사용자 승인 또는 task risk에 비례하지 않는다는 근거가 있어야 하며, 근거와 residual risk를 보고합니다. skip은 pass가 아닙니다.

## Commit Granularity

- commit 하나에는 하나의 related change unit만 포함하고 unrelated change는 staging 전에 분리합니다.
- dependency update는 upgrade와 그 실행에 필요한 수정까지만 함께 포함합니다.
- API, database, schema처럼 독립적으로 검토 가능한 변경은 분리합니다.
- 이후 발견된 typo, 누락, index 수정은 기존 change unit에 소급해 섞지 않고 별도 commit으로 유지합니다.

## Reference Routing

- `references/branch-conventions.md`: branch name이 `codex/`, `jira/prja-000` 같은 policy-sensitive prefix와 맞거나 prefix로 branch 이름을 파생해야 할 때 읽습니다.
- `references/message-lifecycle.md`: empty allocation, message-file path·validation, commit exit status, retained file을 다룰 때 읽습니다.
- `references/recovery.md`: branch 전환 차단, detached HEAD, commit 중단, alias의 부분 실행, push 거부·인증 실패·결과 불명확·remote divergence가 발생했을 때 읽습니다.
- `references/alias-codex.md`: `alias.codex` install·uninstall·doctor를 명시적으로 요청했을 때 읽습니다.
- `SKILL.md`가 reference 선택 조건을 소유하고 reference끼리 같은 정상 workflow를 반복하지 않습니다.

## Branch Convention 판단

- prefix 문자열 자체에 base branch, remote destination, push 허용, cleanup 시점을 내장된 의미로 부여하지 않습니다.
- 규칙 충돌은 instruction priority와 현재 요청의 명시성을 따릅니다. 현재 사용자가 같은 branch 규칙을 명확히 override하지 않았다면 repository 규칙을 유지하고, 둘 다 없을 때만 runtime 일반 convention을 적용합니다.
- exact branch name을 사용자가 제공하면 대소문자와 separator를 보존합니다.
- ticket key나 branch suffix를 파생해야 하는데 source가 없거나 여러 후보가 있으면 임의 생성하지 않습니다.
- repository 규칙과 refspec 예시가 충돌하면 repository 규칙을 따릅니다.

## 실패와 재개 판단

- 실패한 command와 마지막으로 관찰된 local·remote 상태를 분리합니다.
- 재시도 전에는 실패 유형에 필요한 current branch, working tree, HEAD, target ref, worktree, remote ref만 다시 확인합니다.
- commit 성공 후 push 실패처럼 부분 성공이 있으면 성공한 commit을 되돌리지 않고 남은 push만 판단합니다.
- `git codex commit`이 status `2` 또는 `3`이면 HEAD, message file identity, stored message를 다시 확인하기 전 command를 재실행하지 않습니다.
- create failure는 원인을 보존해 recovery를 판단합니다.
- non-fast-forward를 force push로 자동 전환하지 않고 divergence와 허용된 해결 범위를 보고합니다.
- remote 결과가 불명확하면 같은 push를 반복하기 전에 remote ref를 조회합니다.

## 검토 질문

- commit, branch, push 중 현재 요청이 허용한 단계는 무엇인가?
- repository 규칙이 일반 workflow나 branch prefix 예시를 제한하는가?
- working tree에 unrelated staged·unstaged·untracked change가 있는가?
- commit type이 staged scope에 맞고 subject가 적용 가능한 최대 길이 미만인가?
- supporting check의 passed·failed·skipped·unavailable 상태와 residual risk를 정확히 구분했는가?
- commit message에 literal escape, delimiter text, 불필요한 blank line이 없고 skip·residual risk가 숨겨지지 않았는가?
- alias maintenance가 명시적으로 요청됐고 reference의 expected value·preservation rules를 따랐는가?
- 일반 workflow에서 `message create`를 직접 실행했고 failure 원인을 구분했는가?
- staged scope가 하나의 related change unit인가?
- CLI가 명확히 성공한 결과를 불필요하게 재조회하지 않았는가?
- branch exact name과 start point가 확인됐는가?
- `git switch -C`를 사용한다면 기존 branch ref와 다른 worktree 사용 여부를 확인하고 ref 재설정 권한을 명시적으로 확보했는가?
- push가 사용자 지정 current-branch 형식 또는 exact refspec으로 실행됐고, 실패·불확실할 때만 필요한 상태를 확인했는가?
- alias가 stage, commit, push를 묶어 요청 범위를 넓히는가?
- 실패 뒤 현재 상태를 재확인하고 부분 성공을 보존했는가?
- 종료 상태·출력으로 확정 가능한 결과와 추가 상태 조회가 필요한 결과를 구분했는가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 예
- 이유: 설치 후 기존 `git-committer`, sibling skill, dev-only spec 없이 `SKILL.md`와 bundled references만으로 commit·branch·push workflow를 판단해야 합니다.

## 확장 원칙

- commit, branch, push는 하나의 skill 안에서 결합 가능한 기능으로 유지합니다.
- 정상 흐름의 고빈도 command와 공통 안전 계약은 `SKILL.md`에 유지합니다.
- 반복되는 message-file lifecycle을 deterministic command로 옮길 때에도 scope·authority·message semantics은 skill에 남깁니다.
- 반복되는 조건부 branch policy나 failure recovery가 실제로 필요한 경우만 reference를 추가합니다.
- rebase, merge, cherry-pick, worktree 생성·이동·삭제, tag, branch deletion, force push는 검증된 별도 책임이 생기기 전까지 기본 범위에 포함하지 않습니다.
- 기존 `git-committer`와의 migration 또는 제거는 별도 사용자 결정과 change scope가 있을 때만 수행합니다.
