# Agent Wiki 001: Multiple Root 계획

## 변경사항 요약

- 단일 `root` 대신 사용자 정의 named roots를 지원합니다.
- root는 백업·공개·동기화 경계를, root 내부 폴더는 지식·조사·프로젝트 같은 의미 분류를 소유합니다.
- reader·researcher는 필요한 여러 root를 읽을 수 있고, writer는 기록마다 정확히 하나의 root를 선택합니다.
- 각 root는 독립된 `index.md`, 연결성, 파일 수, 깊이 검증 단위입니다.
- 기존 단일 `root` 설정과 `set <directory>`·`path` 호출을 호환합니다.

## 변경 목적

현재 단일 root는 검색 시작점, 쓰기 대상, 문서 연결 기준, 백업·공개 단위를 함께 소유합니다. 한 root 안에서 일반 지식, 개인정보, 작업 기록을 폴더로만 나누면 상위 root의 백업·공개 과정에서 민감 정보가 함께 포함될 수 있습니다.

이번 변경은 정보를 주제뿐 아니라 공개·백업·보존 경계에 따라 물리적으로 분리할 수 있게 합니다. 사용자는 같은 상위 폴더 아래 여러 root를 배치하거나 민감 root를 별도 경로에 둘 수 있습니다. 공통 basedir는 배치 편의이며 안전 경계로 간주하지 않습니다.

## 릴리즈 노트 기준

- 독자에게 보이는 변화: 여러 위키 root를 이름으로 등록하고 기본 root를 지정할 수 있습니다.
- 호환성 영향: 기존 단일 `root` 설정과 인자 없는 `path`는 계속 동작합니다.
- migration 필요 여부: 강제 migration은 없습니다. 기존 설정은 단일 implicit root로 해석하고 설정을 갱신할 때 새 형식으로 정규화할 수 있습니다.
- 운영자/사용자가 알아야 할 검증 결과: 이 문서는 구현 계획입니다. CLI·skill·release 검증 결과는 후속 구현 후 기록합니다.

## 기록 근거

- 관련 커밋: 후속 구현 시 기록합니다.
- 관련 세션 기록: 2026-09-18 다중 저장 경계 설계 요청.

## 변경 상세

### Named roots와 저장 경계

- 목적: 공개·백업·동기화 정책이 다른 정보를 물리적으로 분리합니다.
- 범위:
  - root 이름은 사용자가 정의합니다. `knowledge`, `personality`, `work`는 예시이며 고정 분류가 아닙니다.
  - root는 최종 절대 경로를 저장합니다. 여러 root가 같은 부모 폴더를 가질 수 있지만 동일 경로나 상하 중첩 경로를 가질 수 없습니다.
  - 민감 root를 공개 저장소나 동기화 대상 밖에 둘 수 있습니다. 같은 basedir 전체를 백업하면 모든 하위 root가 포함될 수 있음을 문서화합니다.
  - root 내부 폴더는 지식, 조사, 프로젝트, 작업 기록 등 의미 분류를 자유롭게 표현합니다.
- 관련 표면: plugin spec, README, CLI 설정 모델과 검증, reader·researcher·writer spec/runtime.
- 검증: canonical path 중복·중첩 거부, 형제 root 허용, 존재하지 않거나 파일인 경로 거부, 실패 시 기존 설정 보존.

예상 설정 형식:

```toml
version = 2
default = "knowledge"

[roots.knowledge]
path = "/Users/example/.agents/wiki"
description = "일반 지식과 공유 가능한 조사 결과"

[roots.personality]
path = "/Users/example/.agents/personality"
description = "개인정보, 선호, 개인 작업 기록"

[roots.work]
path = "/Users/example/Documents/agent-work"
description = "프로젝트별 작업 과정과 의사결정 기록"
```

`description`은 root 선택을 돕는 선택 필드입니다. 누락되거나 의미가 모호하면 에이전트가 민감한 기록의 대상을 추측하지 않습니다.

### 탐색 범위와 쓰기 대상

- 목적: 다중 root 탐색과 기록 권한을 분리합니다.
- 범위:
  - reader·researcher는 요청, 현재 작업, root 이름과 설명에 관련된 하나 이상의 root를 읽을 수 있습니다. 모든 root를 항상 검색하지 않습니다.
  - 읽은 정보는 root 이름과 상대 문서 경로를 함께 유지합니다. 검색 가능 여부가 쓰기·복사·이동 권한을 만들지 않습니다.
  - writer는 root를 먼저 결정한 뒤 그 안의 폴더와 문서를 결정합니다. 하나의 기록은 정확히 하나의 root에 둡니다.
  - writer의 선택 순서는 사용자 지정 root, 현재 작업의 명확한 root, 이름·설명에 정확히 맞는 root, 비민감하고 모호하지 않은 경우의 default root입니다.
  - 개인정보나 서로 다른 공개 경계가 섞였거나 대상이 모호하면 자동 분류·이동하지 않습니다. 더 엄격한 root에 둘지 문서를 분리할지 사용자에게 확인합니다.
  - 동일 문서를 여러 root에 자동 복제하거나 기존 문서를 root 사이에서 자동 이동하지 않습니다.
- 관련 표면: 세 skill spec/runtime, plugin usage guidance, README 예시.
- 검증: 단일 root 읽기, 선택적 다중 root 읽기, 명시적 쓰기 대상, 민감·모호한 기록의 확인, 무단 복제·이동 방지 시나리오.

### Root별 구조와 연결성

- 목적: 각 root가 단독 백업·공개돼도 내부 탐색 구조가 완결되게 합니다.
- 범위:
  - 각 root는 자체 `index.md`를 가집니다.
  - 루트 왕복 도달, 직접 파일 25개, root 깊이 0 기준 최대 깊이 4, 복합 폴더명 규칙을 root별로 검증합니다.
  - root 사이의 상대 Markdown 링크는 기본적으로 만들지 않습니다. 한 root만 이동·공개할 때 링크와 정보 경계가 깨지지 않아야 합니다.
  - 여러 root의 관계는 읽는 시점에 결합합니다. 영구 cross-root link나 별도 root URI 체계는 이번 범위에 포함하지 않습니다.
- 관련 표면: 공통 위키 계약, 세 skill의 연결·검증 계약.
- 검증: 각 root의 독립 왕복 도달, 한 root만 복사한 상태의 내부 링크 유효성, cross-root 상대 링크 부재.

### CLI와 설정 migration

- 목적: named roots를 관리하면서 기존 자동화와 설정을 깨지 않습니다.
- 범위:
  - `agent-wiki list`: 설정된 root 이름과 경로를 조회합니다.
  - `agent-wiki path [name]`: 이름이 있으면 해당 root, 없으면 default root 경로를 출력합니다.
  - `agent-wiki set <name> <directory>`: named root를 생성하거나 경로를 갱신합니다.
  - `agent-wiki default <name>`: 기존 named root를 default로 지정합니다.
  - 기존 `agent-wiki set <directory>`는 default root 경로를 설정하는 호출로 유지합니다.
  - 기존 `root = "..."` 설정은 단일 implicit root로 읽습니다. 첫 성공적인 설정 변경에서 새 형식으로 원자적으로 정규화할 수 있습니다.
  - 손상된 TOML, 미지원 형식, 잘못된 경로, 심볼릭 링크 설정 파일은 덮어쓰지 않습니다.
- 관련 표면: Rust CLI, CLI 테스트, README, plugin spec, runtime 경로 탐색 계약.
- 검증: legacy·v2 설정 파싱, `list`·`path`·`set`·`default`, 원자적 갱신, 경로 정규화, invalid config 보존, 삭제된 root 보고.

## 비목표

- 백업·공개·동기화 명령 자체를 구현하지 않습니다.
- `private`, `shareable` 같은 메타데이터를 실제 보안 경계로 취급하지 않습니다.
- 개인정보 자동 감지, 기존 문서 자동 분류, root 간 자동 이동·복제를 수행하지 않습니다.
- 고정된 `knowledge`·`personality`·`work` 분류나 공통 basedir를 강제하지 않습니다.
- cross-root Markdown 링크나 별도 root URI resolver를 도입하지 않습니다.
- 기존 위키 전체를 재분류하거나 구조를 일괄 변경하지 않습니다.

## 호환성 및 마이그레이션

- 기존 단일 `root` 설정은 유효한 입력으로 유지합니다.
- 인자 없는 `path`는 default root의 절대 경로 한 줄을 출력해 기존 호출 계약을 유지합니다.
- 기존 `set <directory>`는 default root를 대상으로 동작합니다.
- 기존 문서와 폴더 구조는 자동 이동하지 않습니다. 사용자가 새 root를 추가한 뒤 이동 범위와 대상 root를 명시해야 합니다.
- 새 parser는 legacy와 v2를 구분하고, 실패한 migration에서 기존 설정 바이트를 보존해야 합니다.

## 검증 기준

- plugin spec, 세 skill spec, README, manifest, runtime skill의 root 용어와 선택 계약이 일치합니다.
- skill spec 변경 후 runtime `SKILL.md`를 현재 spec 기준으로 재작성하고 각 skill에 `quick_validate.py`를 실행합니다.
- Rust 단위 테스트가 legacy·v2 설정, 다중 root 명령, 중복·중첩 경로, 원자적 설정 보존을 검증합니다.
- `cargo fmt --check`, `cargo clippy --locked`, `cargo test --locked`와 네이티브 바이너리 smoke test를 통과합니다.
- release surface에 dev-only `specs/`·`changes/`가 포함되지 않고 source/release runtime과 manifest가 일치합니다.
- 실제 설정이나 위키 문서를 변경하지 않는 임시 홈 시나리오로 검증합니다.

## 정식 규칙 승격 여부

- normative spec 또는 runtime 문서로 승격할 지속 규칙: named root 모델, root와 내부 분류의 구분, 비중첩 경로, root별 독립 연결성, skill별 탐색·쓰기 계약, legacy 호환 CLI.
- change spec에만 남길 release/migration 이력: 단일 root가 여러 책임을 함께 소유했던 배경, v2 도입 과정, 구현·검증 결과와 한계.
