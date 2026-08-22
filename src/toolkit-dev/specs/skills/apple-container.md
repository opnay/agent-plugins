## 사용자 스펙 의도

- advance-codex에 스킬하나 추가할거야.
  - docker gui 가 엔터프라이즈 라이선스 제한으로 사용안할 예정이야.
  - 대체제인 podman이 있어 그걸로 사용하다가 올해 apple container라는 오픈소스가 공개되었어.
  - 그 때문에 이제 docker cli 대신 container cli를 사용할 예정이라. 관련 조사가 필요해. 확인좀해줘.
- toolkit이란 플러그인 만들어서 cli 계열 스킬들 몰아넣어버릴까?
- 그럼 설계해보자.
  - SKILL.md: container cli cheatsheet를 넣고, 하단에 references 폴더 routing 정보
  - references/*.md: 설치, 삭제 가이드, 일부 이론 wiki
- k8s plugin 설명도 넣어두자. kubectl 같이 서드파티 cli 사용해야하는건 제외하거나 보충 설명용으로만 제한하자

---

# apple-container 스킬 스펙

## 목적

Apple Silicon macOS에서 Apple `container` CLI를 로컬 Linux container의 기본 직접 실행 도구로 사용합니다.
고빈도 core 명령과 bundled `k8s` plugin lifecycle을 빠르게 선택하되 Docker CLI 비호환 영역, experimental plugin 상태, 설치 상태, 서비스 lifecycle, 데이터 삭제 위험을 먼저 구분합니다.

## 경계

- 포함:
  - `container` CLI 설치·버전·서비스 상태 확인
  - OCI image build, pull, push와 container, network, volume, registry 관리
  - bundled `container k8s` plugin 탐지와 local single-node cluster create, start, list, load-image, write-config, delete lifecycle
  - Docker CLI 작업을 직접 대응 가능한 `container` 명령으로 옮기는 판단
  - 설치, 갱신, 제거, 개념, Kubernetes plugin 설명을 runtime reference로 라우팅
- 제외:
  - Docker Desktop 설치·운용
  - Docker Engine API, Compose, Testcontainers, devcontainer를 지원한다고 가정하는 자동 변환
  - Podman, Docker 또는 third-party compatibility shim의 자동 설치·전환
  - Linux·CI·production orchestration의 기존 runtime 계약 변경
  - Apple Containerization Swift API 개발
  - `kubectl`, Helm, Kustomize 등 third-party Kubernetes CLI 설치·실행·인증·workload 관리
  - third-party CLI를 `container k8s` lifecycle의 필수 전제로 두는 guidance

## 처리하려는 작업 형태

- 이미 설치된 `container`로 image를 빌드·실행·검사·배포하는 작업
- container, image, network, volume, registry, system 상태를 관리하는 작업
- `docker` 명령을 `container`로 옮길 수 있는지 확인하는 작업
- `container` 설치·갱신·제거 및 데이터 보존 범위를 안내하거나 수행하는 작업
- Apple container의 OCI, VM isolation, Docker·Podman 차이를 설명하는 작업
- bundled `container k8s` plugin으로 local development cluster를 생성·재시작·조회·삭제하거나 local image를 cluster에 전달하고 kubeconfig를 갱신하는 작업
- third-party Kubernetes CLI의 이름과 역할을 kubeconfig handoff 또는 plugin 한계 설명에 필요한 만큼만 언급하는 작업

## 엔트리포인트 / 대표 표면

- 대표 표면: `skills/apple-container/SKILL.md`
- 호출 방식: `$toolkit:apple-container`
- passive trigger: Apple container, container CLI, container k8s, Apple container Kubernetes plugin, Docker Desktop 대체, Docker CLI migration, Podman migration, macOS OCI container

## 핵심 처리 계약

1. 실행 작업이면 host architecture, macOS version, CLI version, system status를 필요한 범위에서 확인합니다.
2. 현재 설치된 CLI의 `container help <command>`를 정적 cheatsheet보다 우선합니다.
3. repository와 요청에서 Compose, Docker socket, Testcontainers, devcontainer, buildx 같은 Docker-specific dependency를 확인합니다.
4. 직접 지원되는 범위만 `container` 명령으로 실행하고 비호환 영역은 자동 번역하지 않습니다.
5. 기존 `Dockerfile`, OCI image reference, registry contract는 변경 필요성이 입증되지 않으면 보존합니다.
6. 설치·갱신·제거·개념 질문은 필요한 runtime reference만 읽도록 라우팅합니다. downgrade처럼 제거 결정이 선행되는 경우 installation reference가 uninstallation reference로 명시적으로 연결합니다.
7. 자원 삭제, prune, user data 제거는 정확한 대상과 복구 가능성을 확인한 뒤 수행합니다.
8. `container help`의 plugin 목록과 `container k8s --help`로 `k8s` plugin availability를 확인하고 plugin subcommand는 `container k8s <subcommand> --help`를 source of truth로 사용합니다.
9. `container k8s`가 생성·갱신·삭제하는 cluster container와 kubeconfig entry를 side effect로 취급하고 정확한 cluster name과 kubeconfig target을 확인합니다.
10. third-party Kubernetes CLI는 설치하거나 실행하지 않습니다. plugin이 만든 kubeconfig의 소비자나 workload 관리 경계를 설명할 때만 보충적으로 언급합니다.

## Cheatsheet 소유권

- `SKILL.md`는 system, container, image, registry, network, volume과 bundled `k8s` plugin의 고빈도 명령을 짧은 예시로 제공합니다.
- 모든 flag와 subcommand를 복제하지 않고, 현재 CLI help를 확인하는 명령을 포함합니다.
- 설치, 갱신, 제거 절차와 장문의 개념 설명을 중복하지 않습니다.
- 명령 예시는 Apple `container` 1.2.2에서 확인하되 특정 version을 영구 최신값으로 단정하지 않습니다.

## Reference Routing

- `references/installation.md`: CLI가 없거나 초기 설치, kernel setup, update, downgrade가 필요할 때 읽습니다. downgrade는 기존 설치 제거 범위를 정하기 위해 uninstallation reference로 이어질 수 있습니다.
- `references/uninstallation.md`: tool 제거, user data 보존 또는 완전 삭제를 판단할 때 읽습니다.
- `references/concepts.md`: OCI, per-container VM, Docker·Podman 차이, compatibility boundary를 설명하거나 판단할 때 읽습니다.
- `references/kubernetes.md`: experimental `container k8s` plugin의 model, lifecycle, kubeconfig side effect, third-party CLI 경계를 설명하거나 판단할 때 읽습니다.
- `SKILL.md` 하단이 routing 정보를 소유하며 reference끼리 같은 내용을 반복하지 않습니다.

## 호환성 판단

- Docker Desktop의 구독 조건과 Docker CLI·Dockerfile의 라이선스·format을 같은 문제로 취급하지 않습니다.
- Apple `container`가 OCI image와 Dockerfile/Containerfile build를 지원한다는 사실과 Docker Engine API 호환성을 구분합니다.
- Compose나 Docker API consumer가 발견되면 기본 지원을 가정하지 않고 현재 plugin·socket·tool contract를 확인합니다.
- cross-platform repository에서는 local macOS command preference와 CI·Linux command contract를 분리합니다.
- `container k8s`는 local single-node development cluster용 experimental plugin으로 다루고 production Kubernetes나 일반 Kubernetes client workflow로 확장하지 않습니다.
- plugin 자체 lifecycle command와 third-party client command를 분리하며, `kubectl` 등은 보충 설명 외 runtime 실행 계약에 포함하지 않습니다.

## 검토 질문

- CLI 작업의 host가 Apple Silicon macOS인가?
- 설치된 `container` version과 system status를 확인했는가?
- 현재 command help가 cheatsheet 예시와 일치하는가?
- 요청이나 repository가 Docker Compose/API socket에 의존하는가?
- Dockerfile, OCI image, registry contract를 불필요하게 변경하지 않았는가?
- 삭제 또는 prune의 정확한 대상과 복구 가능성을 확인했는가?
- `k8s` plugin이 현재 설치본에서 발견되고 plugin 전용 help를 확인했는가?
- cluster name과 `~/.kube/config` 또는 alternate kubeconfig target의 변경 범위를 확인했는가?
- third-party Kubernetes CLI를 설치·실행하거나 필수 prerequisite로 만들지 않았는가?
- 현재 작업에 필요한 reference만 읽었는가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 예
- 이유: 설치 후 sibling skill이나 dev-only spec 없이 `SKILL.md`와 bundled references만으로 동작해야 합니다.

## 확장 원칙

- 고빈도 명령이 늘어나도 full manual을 복제하지 않고 현재 help 확인과 authoritative link를 유지합니다.
- 반복되는 별도 작업 모드가 생길 때만 reference를 추가합니다.
- third-party Compose나 Docker API shim은 검증된 독립 책임이 생기기 전까지 기본 계약에 포함하지 않습니다.
- Kubernetes workload management가 필요해져도 `apple-container` 안에 third-party client workflow를 흡수하지 않고 별도 소유 표면을 검토합니다.
