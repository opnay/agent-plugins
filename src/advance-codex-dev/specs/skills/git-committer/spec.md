# git-committer Skill Spec

## 목적

`git-committer`는 readiness 판단이 끝난 변경을 실제 task-scoped commit으로 마무리하는 finalization skill입니다.
핵심은 커밋 준비, staged 검증, 파일 기반 commit message 전달, post-commit 확인을 하나의 좁은 실행 계약으로 묶는 것입니다.

## 경계

- 포함:
  - commit이 포함된 사용자 요청이나 상위 작업 workflow의 commit 단계 실행
  - commit 범위 분리와 staged diff 검토
  - staged 검증과 필요한 보조 확인의 skip/failure 보고
  - colon-separated commit message와 body 작성 규율
  - 메시지 파일 생성 > `git commit -F <file>` > 파일 정리 lifecycle
  - 최종 commit 생성과 metadata 확인
- 제외:
  - readiness 판단 자체
  - commit이 사용자 요청 범위에 포함되는지 결정하는 상위 작업 해석
  - unrelated change cleanup
  - interactive git tutoring 전반
  - implementation 자체의 설계
  - push, PR, release, publish, version bump 실행

## 처리하려는 작업 형태

- 사용자가 작업을 commit으로 마무리하거나 실제 commit 실행을 요청한 경우
- `PR 올려놔`처럼 완료에 commit이 필요한 상위 작업을 요청한 경우
- mixed change를 task-scoped commit 단위로 나눠야 하는 경우
- commit message quality와 staged 검증이 중요한 경우

## 대표 표면

- 대표 runtime 표면: `advance-codex-dev/skills/git-committer/SKILL.md`
- 사용자 스펙 의도: `advance-codex-dev/specs/skills/git-committer/intent.md`
- skill spec index: `advance-codex-dev/specs/skills/git-committer/spec.md`
- sub-spec directory: `advance-codex-dev/specs/skills/git-committer/`

## 상세 계약 구조

- `intent.md`: 사용자 스펙 의도와 commit flow graph
- `workflow.md`: request-scope entry, commit preparation, message file gate, commit execution
- `message.md`: commit type, subject, body, file input 작성 규칙
- `verification.md`: staged 검증, 필요한 보조 확인, message file cleanup, skip/failure, post-commit 확인

## 핵심 처리 계약

- `git-committer`는 commit이 사용자 요청 범위에 포함된 뒤의 커밋 준비와 실행을 처리합니다.
- commit을 포함하는 상위 작업 요청에는 별도 commit-specific 승인이나 재확인 단계를 추가하지 않습니다.
- push, PR, release, publish, version bump는 이 skill의 실행 범위가 아니지만, 해당 상위 workflow의 commit 단계에는 이 skill을 사용할 수 있습니다.
- 커밋 준비는 프로젝트의 커밋 준비 단계와 커밋할 범위 선택을 포함합니다.
- 커밋 실행은 staged 검증, 메시지 준비, 커밋 생성을 포함합니다.
- staged 검증이나 필요한 보조 확인은 실행 불가/스킵 이유와 residual risk를 기록할 수 있어야 합니다.
- 커밋 메시지는 `type: detailed subject` 형식과 bullet body를 사용합니다.
- commit message는 신뢰된 temporary-file allocator가 만든 전용 파일에 기록하고, 파일 내용을 확인한 뒤 `git commit -F <file>`로만 전달합니다.
- Bash heredoc/EOF, here-string, command substitution, stdin의 `git commit -F -`, 여러 `-m` 인자로 commit message를 구성하지 않습니다.
- 메시지 파일을 만든 뒤에는 commit 성공, 실패, 취소와 관계없이 정확한 파일을 정리하고 삭제 여부를 확인합니다.
- 강제 interruption으로 cleanup을 실행하지 못하면 다음 재개 시 allocator가 만든 것으로 검증된 정확한 파일만 먼저 정리하고 새 commit 시도를 시작합니다.
- 파일 생성, 내용 확인, commit, 정리 중 하나라도 계약대로 처리할 수 없으면 다음 단계로 조용히 넘어가지 않습니다.
- 커밋 후에는 최신 commit metadata와 working tree 상태를 확인해 결과를 보고합니다.

## 독립성 원칙

`git-committer`는 독립 실행 가능한 runtime skill이어야 합니다.
본문은 sibling skill 이름을 handoff 기준으로 언급할 수 있지만, 실행을 위해 dev-only spec 경로나 hidden context를 읽으라고 지시하지 않습니다.
다른 skill은 readiness 판단까지만 맡고, 실제 commit finalization 규칙은 이 skill이 소유합니다.

## 검증 기준

- dev runtime skill이 `skills/git-committer/SKILL.md`에 존재해야 한다.
- release build 후 root `advance-codex/skills/git-committer/SKILL.md`가 dev source와 맞아야 한다.
- plugin spec, README, manifest prompt가 `git-committer`의 역할과 사용 기준을 언급해야 한다.
- runtime skill 본문은 dev-only `specs/` 또는 `src/advance-codex-dev` 경로를 실행 지시로 포함하지 않아야 한다.
- runtime skill은 별도 commit-specific 승인 gate 없이 요청 범위의 commit을 실행하고, skip/failure verification reporting을 명시해야 한다.
- runtime skill과 command reference는 메시지 파일 생성 > readback > `git commit -F <file>` > 정리 순서를 명시하고 heredoc/EOF 또는 stdin 대안을 허용하지 않아야 한다.

## 확장 원칙

- 사용자 의도와 흐름도는 `intent.md`에 둡니다.
- lifecycle, message, verification 규칙은 각 child spec이 소유합니다.
- runtime skill은 folderized spec 전체를 짧고 실행 가능한 지시로 압축합니다.
- 새 reference는 command safety, message quality, verification reliability에 직접 기여할 때만 추가합니다.
