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

## 경계

- 포함: 로컬 개발 CLI 자체가 작업 대상인 환경 확인, 실행, 유지보수, 호환성 판단
- 제외: 일반 shell 명령 모음, 특정 제품 기능 구현, connector가 소유하는 앱 작업, CLI를 우연히 사용하는 workflow
- Kubernetes 제한: `kubectl`, Helm, Kustomize 등 third-party CLI의 설치·실행·workload 관리는 소유하지 않고 plugin 경계 설명에만 사용합니다.

개발 원본 계약은 `specs/plugin.md`와 `specs/skills/`가 소유합니다.
