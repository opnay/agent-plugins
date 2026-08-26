# Workflow Contract

## Entry Gate

- commit이 현재 사용자 요청 범위에 포함돼야 합니다. 직접 commit 요청뿐 아니라 `PR 올려놔`처럼 완료에 commit이 필요한 상위 작업 요청도 포함합니다.
- 이 skill 안에서 별도 commit-specific 승인이나 재확인 단계를 두지 않습니다.
- push, PR, release, publish, version bump 실행은 이 skill의 범위가 아니며, 해당 workflow의 commit 단계만 처리합니다.

## Commit Preparation

1. 프로젝트의 커밋 준비 단계가 끝났는지 확인합니다.
2. staged diff와 unstaged/untracked 상태를 읽어 커밋할 범위를 선택합니다.
3. 범위가 섞이면 split 또는 restage를 먼저 처리하고, 커밋 범위를 다시 확인합니다.

## Commit Execution

1. staged 검증을 실행합니다.
2. 필요한 보조 확인이 있으면 변경 범위에 맞는 가장 좁은 check를 실행합니다.
3. 실패는 고치고 재검증하거나, 고칠 수 없으면 commit을 막고 실패를 보고합니다.
4. staged 검증을 실행할 수 없으면 commit을 막습니다. supporting check를 실행할 수 없으면 이유와 residual risk를 기록하고 task risk가 허용할 때만 계속합니다.
5. 의도적 skip은 사용자 승인 또는 task risk 비례 근거와 residual risk를 기록합니다.
6. commit message를 작성하고 Message File Gate를 통과합니다.
7. commit 성공 시 최신 commit metadata와 working tree 상태를 보고합니다.

## Message File Gate

1. 신뢰된 temporary-file allocator로 전용 메시지 파일을 하나 생성하고 정확한 경로를 보존합니다. exit 0, 단일 반환 경로, 예상 template 일치를 충족해야 allocation 성공입니다. 사용자 제공 경로나 임의 저장 경로를 사용하지 않습니다. nonzero exit에 경로가 없으면 cleanup 없이 commit을 막고, 현재 allocator invocation이 예상 template과 일치하는 단일 경로를 반환했다면 그 exact path만 cleanup합니다.
2. shell multiline input이나 redirection이 아닌 파일 쓰기 도구로 commit message만 기록합니다.
3. 파일을 다시 읽어 subject, blank line, bullet body, 의도하지 않은 shell 문법이나 escape가 없는지 확인합니다.
4. final `git status`와 staged diff를 다시 확인합니다.
5. allocator가 반환한 정확한 파일 경로를 사용해 `git commit -F <file>`만 실행합니다.
6. commit 시도 성공, 실패, 실행 전 제어 가능한 중단과 관계없이 현재 allocator invocation 또는 task state provenance와 예상 template이 확인된 exact path를 `unlink`합니다. 명령이 성공하면 cleanup 완료이며, 실패하면 남은 경로와 residual risk를 보고합니다.
7. 강제 interruption으로 cleanup을 실행하지 못했다면, 재개 시 task state가 보존한 allocator provenance와 예상 template을 확인합니다. exact path의 cleanup 결과를 확인한 뒤 새 commit 시도를 시작합니다.
8. commit 결과와 cleanup 결과를 분리해 보고합니다. 어느 단계에서든 cleanup이 실패하면 남은 정확한 경로와 residual risk를 명시하며, cleanup 실패는 성공한 commit을 되돌리지 않습니다.

파일 생성, 쓰기, readback, final staged diff, `git commit -F`, cleanup 순서를 합친 단일 shell script로 실행하지 않습니다.
메시지 본문을 Bash heredoc/EOF, here-string, command substitution, stdin의 `git commit -F -`, 여러 `-m` 인자로 전달하지 않습니다.

## Guardrails

- unrelated changes를 조용히 포함하지 않습니다.
- staged 상태를 확인하지 않고 commit하지 않습니다.
- skipped verification을 green으로 보고하지 않습니다.
- 메시지 파일을 만들기 전에 경로를 추측하거나, 만든 뒤 cleanup 없이 남기지 않습니다.
- commit 완료를 push/PR/release/publish 완료로 보고하지 않습니다.
