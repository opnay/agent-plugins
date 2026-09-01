# Toolkit Dev

`toolkit-dev`는 Codex가 로컬 개발 CLI를 안전하고 일관되게 운용하도록 돕는 플러그인입니다.
CLI를 사용한다는 이유만으로 skill을 모으지 않고, 도구 자체의 설치·환경 탐지·명령 계약·호환성·복구가 반복해서 필요한 경우만 포함합니다.

## Skill 선택

### Apple Container

`$toolkit-dev:apple-container`는 Apple Silicon macOS에서 Apple `container` CLI와 bundled experimental `k8s` plugin을 사용할 때 적용합니다.
고빈도 명령표는 `SKILL.md`, 조건부 설치·제거 절차와 개념·Kubernetes plugin 설명은 runtime `references/`가 소유합니다.

대표 요청:

- Apple `container`로 OCI 이미지를 빌드하거나 실행합니다.
- Docker Desktop 없이 로컬 Linux 컨테이너를 관리합니다.
- Docker CLI 작업을 Apple `container`로 옮길 수 있는지 판단합니다.
- Apple `container`를 설치, 갱신, 제거하거나 동작 원리를 확인합니다.
- `container k8s`로 local development cluster lifecycle과 kubeconfig handoff를 관리합니다.

### Git

`$toolkit-dev:git`은 task-scoped commit, branch, push를 각각 선택하거나 하나의 workflow로 연결할 때 적용합니다.
고빈도 정상 흐름과 명령은 `SKILL.md`, 조건부 branch prefix와 실패·중단 복구는 runtime `references/`가 소유합니다.

대표 요청:

- task-owned related change unit만 stage하고 risk-proportional check, message 위생, 120자 미만 subject·가장 구체적인 commit type·파일 기반 메시지로 commit한 뒤 실제 저장된 full message를 검증합니다.
- exact start point에서 branch를 생성하거나 기존 branch로 전환하고, 명시적으로 요청된 경우 `git switch -C`로 branch를 force-create합니다.
- current branch에 upstream을 설정하고 push합니다.
- local source와 remote destination이 다른 refspec push를 실행·검증합니다.
- `codex/`, `jira/prja-000` 같은 policy-sensitive prefix의 owning rule을 확인합니다.
- commit 또는 push의 부분 실패 상태를 확인하고 완료된 단계를 보존한 채 재개합니다.

## 경계

- 포함: 로컬 개발 CLI 자체가 작업 대상인 환경 확인, 실행, 유지보수, 호환성·실패 복구 판단
- 제외: 일반 shell 명령 모음, 특정 제품 기능 구현, connector가 소유하는 앱 작업, CLI를 우연히 사용하는 workflow
- Kubernetes 제한: `kubectl`, Helm, Kustomize 등 third-party CLI의 설치·실행·workload 관리는 소유하지 않고 plugin 경계 설명에만 사용합니다.
- Git 제한: GitHub PR·release·hosting API와 요청에 없는 commit·branch·push mutation은 소유하지 않습니다.

개발 원본 계약은 `specs/plugin.md`와 `specs/skills/`가 소유합니다.
