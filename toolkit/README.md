# Toolkit

`toolkit`는 Codex가 로컬 개발 CLI를 안전하고 일관되게 운용하도록 돕는 플러그인입니다.
CLI를 사용한다는 이유만으로 skill을 모으지 않고, 도구 자체의 설치·환경 탐지·명령 계약·호환성·복구가 반복해서 필요한 경우만 포함합니다.

## Skill 선택

### Apple Container

`$toolkit:apple-container`는 Apple Silicon macOS에서 Apple `container` CLI와 bundled experimental `k8s` plugin을 사용할 때 적용합니다.
고빈도 명령표는 `SKILL.md`, 조건부 설치·제거 절차와 개념·Kubernetes plugin 설명은 runtime `references/`가 소유합니다.

대표 요청:

- Apple `container`로 OCI 이미지를 빌드하거나 실행합니다.
- Docker Desktop 없이 로컬 Linux 컨테이너를 관리합니다.
- Docker CLI 작업을 Apple `container`로 옮길 수 있는지 판단합니다.
- Apple `container`를 설치, 갱신, 제거하거나 동작 원리를 확인합니다.
- `container k8s`로 local development cluster lifecycle과 kubeconfig handoff를 관리합니다.

### Git

`$toolkit:git`은 task-scoped commit, branch, push를 각각 선택하거나 하나의 workflow로 연결할 때 적용합니다.
고빈도 정상 흐름과 명령은 `SKILL.md`, 조건부 branch prefix와 실패·중단 복구는 runtime `references/`가 소유합니다.

대표 요청:

- task-owned related change unit만 stage하고 risk-proportional check, message 위생, 120자 미만 subject·가장 구체적인 commit type·파일 기반 메시지로 commit한 뒤 실제 저장된 full message를 검증합니다. bundled `git-codex`는 heredoc 입력·자동 validation·성공 뒤 cleanup을 제공합니다.
- exact start point에서 branch를 생성하거나 기존 branch로 전환하고, 명시적으로 요청된 경우 `git switch -C`로 branch를 force-create합니다.
- current branch에 upstream을 설정하고 push합니다.
- local source와 remote destination이 다른 refspec push를 실행·검증합니다.
- `codex/`, `jira/prja-000` 같은 policy-sensitive prefix의 owning rule을 확인합니다.
- commit 또는 push의 부분 실패 상태를 확인하고 완료된 단계를 보존한 채 재개합니다.

### 메시지 생성

skill의 alias·staged 범위 확인 뒤 quoted heredoc 입력을 우선합니다.

```sh
git codex message create --stdin <<'EOF'
fix: 변경 내용을 구체적으로 설명

- 변경 사항과 검증 근거
EOF
```

성공 시 자동 validation을 마친 `MSG-…` ID 하나를 반환합니다. 일반 repository에서는 `.git/MSG-…`에 저장하며, Git이 실제 metadata 경로를 찾아 하위 디렉터리·linked worktree에서도 동작합니다. 같은 repository·worktree에서 `git codex commit <반환된-MSG-ID>`로 사용합니다. 생성 후 파일을 수정하지 않았다면 별도 validation을 반복하지 않습니다.

빈 파일이 필요한 경우 `message create`를 사용하고 작성 후 `message validate <MSG-ID>`를 실행합니다. push 등은 종료 상태와 대상별 결과 출력이 명확하면 그대로 보고하고, 호출 오류·중단·결과 불명확에 필요한 조회만 수행합니다.

## alias.codex 관리

`git codex`는 관리되는 global `alias.codex`를 통해 `$HOME/.local/bin/git-codex`를 호출합니다. alias install·uninstall·doctor는 명시 요청에서만 [`skills/git/references/alias-codex.md`](skills/git/references/alias-codex.md)를 따릅니다.

## 경계

- 포함: 로컬 개발 CLI 자체가 작업 대상인 환경 확인, 실행, 유지보수, 호환성·실패 복구 판단
- 제외: 일반 shell 명령 모음, 특정 제품 기능 구현, connector가 소유하는 앱 작업, CLI를 우연히 사용하는 workflow
- Kubernetes 제한: `kubectl`, Helm, Kustomize 등 third-party CLI의 설치·실행·workload 관리는 소유하지 않고 plugin 경계 설명에만 사용합니다.
- Git 제한: GitHub PR·release·hosting API와 요청에 없는 commit·branch·push mutation은 소유하지 않습니다.
