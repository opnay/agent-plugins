# AGENTS.md

---
## Part 1. 저장소 및 마켓플레이스 규칙
이 섹션은 저장소 구조, 플러그인 배치, 마켓플레이스 메타데이터 변경에 적용됩니다.
---

## 저장소의 역할

이 저장소는 사용자가 직접 만드는 플러그인을 관리하기 위해 존재합니다.
로컬 플러그인 마켓플레이스의 루트이자 하네스 엔지니어링 작업 공간으로 사용합니다.
제품화 저장소가 아닙니다.
이 파일은 저장소 레이아웃과 하네스 품질 기준을 함께 다루는 운영 가이드로 취급합니다.

## 저장소 레이아웃

- 공개 설치용 release 플러그인은 저장소 루트 바로 아래에 둡니다.
- 개발 원본 플러그인은 `./src/<plugin-name>-dev` 아래에 둡니다.
- specs는 개발 원본인 `./src/<plugin-name>-dev/specs/` 안에서만 관리합니다.
- 일반 개발 변경의 기본 편집 대상은 `./src/<plugin-name>-dev`입니다.
- 저장소 루트의 release 플러그인은 수동 편집하지 않고 build command 산출물로 갱신합니다.
- 나중에 저장소 구조를 의도적으로 바꾸지 않는 한, 이 저장소에서 `./plugins/<plugin-name>` 경로를 만들거나 사용하지 않습니다.

## 브랜치 운영 규칙

- `main` 브랜치는 공개 release 브랜치입니다.
- `next` 브랜치는 개발 브랜치입니다.
- 일반 플러그인 수정은 `next`에서 `src/<plugin-name>-dev`에 적용합니다.
- `main`에는 `next`에서 검증된 개발 내용을 release로 승격할 때만 반영합니다.
- `next` 브랜치를 사용하는 동안에는 dev 플러그인 자체를 설치/사용할 수 있어야 하므로, 개발 표면은 `<plugin-name>-dev` 이름과 `src/<plugin-name>-dev` 구조를 유지합니다.
- build command는 매 plugin 변경마다 루트 `./<plugin-name>` release surface를 갱신합니다.

## 버전 승격 규칙

- `next`에서 마지막 `main` merge 이후 특정 dev 플러그인을 처음 수정할 때, 먼저 `next`의 dev 플러그인 version bump 대상입니다.
- version bump 종류는 자동으로 단정하지 않고 사용자에게 patch/minor/major 또는 구체 version을 확인합니다.
- 사용자가 version 유지 또는 bump를 명시하면 그 결정에 따라 `src/<plugin-name>-dev/.codex-plugin/plugin.json` version을 유지하거나 올립니다.
- 예: `turn-gate`를 수정했다면 `next`의 `src/loop-kit-dev`와 필요 시 upstream `src/workflow-kit-dev`를 수정합니다.
- 이때 마지막 `main` merge 이후 첫 `loop-kit-dev` 수정이라면, 수정 전에 사용자에게 `loop-kit-dev` patch/minor/major bump 선택지를 엽니다.
- 같은 플러그인의 이후 변경은 추가 version bump 없이 build만 수행합니다.
- `next` 브랜치를 marketplace로 등록해 개발 버전을 사용하는 사람은 bump된 dev plugin version을 설치/업데이트합니다.
- 실제 공개 release 단계가 되면 `next`를 `main`으로 merge하고, release surface를 루트 폴더로 생성/갱신합니다.
- build는 매 plugin 변경마다 수행하고, release(version bump)는 마지막 `main` merge 이후 해당 plugin의 첫 수정 때 수행합니다.

## 필수 플러그인 구조

각 개발 원본 플러그인은 다음을 포함해야 합니다.

- `./src/<plugin-name>-dev/.codex-plugin/plugin.json`
- `./src/<plugin-name>-dev/README.md`
- `./src/<plugin-name>-dev/specs/plugin.md`
- 선택 사항: `./src/<plugin-name>-dev/specs/skills/`
- 선택 사항: `./src/<plugin-name>-dev/skills/`
- 선택 사항: `./src/<plugin-name>-dev/assets/`
- 선택 사항: `./src/<plugin-name>-dev/scripts/`
- 선택 사항: `./src/<plugin-name>-dev/.mcp.json`
- 선택 사항: `./src/<plugin-name>-dev/.app.json`

플러그인 폴더 이름과 `plugin.json`의 `"name"` 값은 반드시 일치해야 합니다.
공개 release 플러그인은 `<plugin-name>`을 사용하고, 개발 원본 플러그인은 `<plugin-name>-dev`를 사용합니다.

## Spec-Driven Plugin Development

이 저장소에서는 플러그인 설계와 변경을 spec driven development로 다룹니다.

- `README.md`는 플러그인의 전반적인 목적, 왜 이 플러그인이 존재하는지, 어떤 작업을 다루는지를 설명해야 합니다.
- `specs/plugin.md`는 최소한 다음을 현재 기준으로 고정해야 합니다.
  - 플러그인 목적
  - 플러그인 경계와 비목표
  - 어떤 작업 형태를 처리하려는지
  - 엔트리포인트 또는 대표 표면
  - 내장 skill 체계와 각 skill의 역할 요약
  - 새 skill을 추가하거나 기존 skill 책임을 바꿀 때 유지해야 할 확장 원칙
- 플러그인 스펙은 기본적으로 플러그인 경계, 라우팅 표면, 내장 skill 체계를 소유하고, 개별 skill의 상세 처리 계약은 가능한 한 분리해서 다룹니다.
- 개별 skill의 입력 형태, 판단 기준, 출력 계약, 가드레일처럼 skill 자체의 처리 계약이 중요해지면 `specs/skills/<skill-name>.md`로 분리하는 방식을 우선 검토합니다.
- 플러그인 스펙에서 각 skill을 언급할 때는 목적과 관계를 요약하고, 상세 기준은 대응되는 skill spec 위치로 연결하는 방식을 기본으로 합니다.
- 플러그인 작업은 spec이 먼저 있고, skill/manifest 변경은 그 spec과 일치해야 합니다.
- 플러그인 표면이 바뀌면 `README.md`, `specs/plugin.md`, 관련 skill spec, 관련 guide skill, `plugin.json`을 같은 변경 단위에서 함께 점검합니다.
- spec 없는 skill 추가를 기본 경로로 두지 않습니다.
- spec은 구현 세부보다 의도, 경계, 라우팅, 책임 배치를 먼저 고정해야 합니다.
- 파일 종류나 구현 관성을 적더라도, spec의 핵심 판단 기준은 change pressure, ownership, routing 이유를 먼저 설명하는 쪽을 우선합니다.

## 권장 스펙 양식

새 플러그인이나 skill spec을 시작할 때는 아래 템플릿을 기본 시작점으로 사용합니다.

- plugin spec starter: `docs/templates/plugin-spec.md`
- skill spec starter: `docs/templates/skill-spec.md`
- 템플릿은 현재 저장소에 남아 있는 유지 기준 spec과 동일한 방향으로 갱신합니다.

양식을 사용할 때의 기준:

- plugin spec에는 bundle 목적, 경계, 라우팅 표면, skill composition만 남기고 세부 처리 계약은 가능한 한 skill spec으로 내립니다.
- skill spec의 `사용자 스펙 의도`에는 사용자가 기대한 판단, 규칙, 예시를 정규 스펙 본문으로 옮기기 전의 입력 관점으로 적습니다.
- skill spec에는 해당 skill이 실제로 무엇을 판단하고 어떤 계약으로 동작하는지 적습니다.
- skill spec이 실제 판단을 수행하는 종류라면 `핵심 처리 계약` 뒤에 실제 문서에서 쓰는 규칙 섹션 이름을 그대로 두고, 필요할 때만 섹션 수를 줄이거나 늘립니다.
- skill spec의 `독립성 원칙`에는 독립 실행 가능성을 항상 강제하지 말고, 그 skill에서 독립성을 spec으로 강제해야 하는지 여부와 이유를 적습니다.
- guide skill은 라우팅을, narrow skill은 자기 처리 계약을 소유하게 써서 서로의 책임이 섞이지 않게 유지합니다.
- 반복되는 검토 질문, 예시, decision rule이 생기면 skill spec 안에서 소유할지 별도 reference 문서로 뺄지 의도적으로 결정합니다.

## 마켓플레이스 단일 진실 공급원

- 마켓플레이스 파일: `./.agents/plugins/marketplace.json`
- 이 저장소에 추가하는 모든 공개 release 플러그인은 해당 파일에 대응되는 항목이 있어야 합니다.
- 이 저장소에서는 마켓플레이스의 `source.path`가 항상 저장소 루트의 플러그인 폴더를 가리켜야 합니다.
- 개별 플러그인 경로와 등록 순서는 `./.agents/plugins/marketplace.json`의 `plugins` 배열을 기준으로 확인합니다.
- 이 마켓플레이스에서 저장소 로컬 플러그인을 `./plugins/<plugin-name>`로 등록하지 않습니다.

## 플러그인 변경 워크플로

1. 개발 원본 플러그인은 `./src/<plugin-name>-dev`에 생성하거나 이동합니다.
2. 일반 개발 변경은 `./src/<plugin-name>-dev`의 README, specs, skills, manifest에 먼저 적용합니다.
3. 변경한 plugin은 build command로 저장소 루트 `./<plugin-name>` release surface를 갱신합니다.
4. `.codex-plugin/plugin.json`이 존재하고 유효한 JSON인지 확인합니다.
5. specs는 `./src/<plugin-name>-dev/specs/`에서 만들거나 현재 표면에 맞게 갱신합니다.
6. release 승격이 요청된 경우에만 `./.agents/plugins/marketplace.json`에 공개 release 항목을 추가하거나 갱신합니다.
7. 모든 마켓플레이스 항목에 `policy.installation`, `policy.authentication`, `category`가 포함되도록 유지합니다.
8. 변경한 JSON 파일은 수정 후 검증합니다.

## Release 승격 워크플로

1. `next`에서 release할 플러그인 변경 범위를 확인합니다.
2. `next`에서 이미 dev plugin version bump가 반영됐는지 확인합니다.
3. 마지막 `main` merge 이후 첫 수정인 플러그인은 사용자에게 patch/minor/major 또는 목표 version을 확인하고 release command를 실행합니다.
4. 같은 플러그인의 이후 수정은 추가 version bump 없이 build command만 실행합니다.
5. `next`를 `main`으로 merge할 release 단위를 확정합니다.
6. build command로 `src/<plugin-name>-dev`를 루트 `./<plugin-name>` release surface로 변환합니다.
7. release root에 `specs/`가 없는지, manifest name/version과 marketplace entry가 맞는지 검증합니다.
8. 검증된 release 변경을 `main`에 반영합니다.

## 플러그인 의도 관련 메모

- 이 저장소는 사용자가 직접 만들고 직접 관리하는 플러그인을 유지보수하는 장소로 취급합니다.
- `advance-codex`는 사용자가 직접 만들고 유지할 수 있는 Codex 기능을 더 깊게 관리하기 위해 존재합니다.
- `advance-codex`의 대표 표면은 보통 skill, tool-use guidance, plugin, subagent입니다.
- 명시적 의도가 없는 한, `advance-codex`를 무관한 워크플로나 cross-plugin 유틸리티 범위로 넓히지 않습니다.

## 플러그인 엔트리 스킬 가이드

여러 사용자 지향 skill을 가진 플러그인이라면, 엔트리포인트 skill 하나를 두는 쪽을 우선 고려합니다.

플러그인과 skill 설계는 위에서 아래로 진행해야 합니다.

1. 먼저 플러그인 경계를 정의합니다.
2. 그다음 그 플러그인 안에 들어갈 skill을 정의합니다.

명시적이고 의도적인 마이그레이션이 아니라면, 느슨한 skill 묶음부터 시작한 뒤 나중에 그에 맞춰 플러그인 모양을 억지로 만들지 않습니다.

## 플러그인과 스킬의 독립성 판단

- 모든 플러그인은 하나의 일관된 번들로 독립적으로 이해 가능해야 합니다.
- 모든 skill에 독립 실행 가능성을 일괄 강제하지 않습니다.
- 대신 각 skill spec에서 그 skill이 독립성을 강제해야 하는지, sibling context를 허용하는지, 왜 그런지를 명시합니다.
- skill이 플러그인에 속하더라도 책임은 분명해야 하며, sibling skill이나 guide가 소유해야 할 문맥을 숨은 전제로 끌어오지 않도록 spec에서 경계를 설명합니다.
- 플러그인은 자기 목적을 설명하기 위해 무관한 외부 플러그인 구조에 기대면 안 됩니다.

## 플러그인 설치 스킬 식별자

- Plugin을 통해 설치된 skill은 `$<plugin>:<skill>` 식별자를 사용합니다.
- 예시: `$advance-codex:skill-creator`

## 플러그인 소속 스킬 변경 규칙

- 플러그인에 속한 skill을 수정할 때는 해당 skill만 고립해서 보지 않습니다.
- 먼저 그 skill이 플러그인 안에서 맡는 역할과 sibling skill, `<plugin>-guide`와의 관계를 확인합니다.
- skill의 책임, 의미, 라우팅 기준이 바뀌면 관련 skill spec, guide skill, 인접 skill의 문구도 같은 변경 단위에서 함께 점검하고 필요하면 갱신합니다.
- plugin-scoped skill 변경은 개별 skill 수정이면서 동시에 plugin surface 수정일 수 있음을 전제로 작업합니다.
- 플러그인 관점의 영향이 있는데도 이를 skill 단독 수정으로 축소하지 않습니다.

## `<plugin>-guide` 스킬 규칙

플러그인을 만들 때는 해당 플러그인의 엔트리포인트 skill로 `<plugin>-guide`도 함께 만듭니다.

이 skill은 다음 역할을 해야 합니다.

- 플러그인을 효과적으로 사용하는 방법을 설명한다
- 사용자나 에이전트를 올바른 내장 skill로 라우팅한다
- 플러그인에 여러 기능이 있을 경우, 더 깊은 실행 전에 작업 유형을 분류한다

이 guide skill은 플러그인 범위에 속합니다.

- 해당 플러그인 내부에 존재해야 합니다
- 해당 플러그인의 사용성을 지원하기 위해 존재합니다
- 플러그인의 독립적인 표면 일부로서 함께 유지보수해야 합니다

다음 조건이면 엔트리포인트 skill을 사용합니다.

- 플러그인에 서로 역할이 다른 skill이 둘 이상 있다
- 사용자나 에이전트가 적절한 워크플로나 도메인 skill을 고르는 데 도움이 필요하다
- 더 좁은 skill 선택 전에 작업 형태를 분류하는 편이 유리하다

엔트리포인트 skill의 기대사항:

- 더 깊은 실행 전에 작업을 분류할 것
- 깨지기 쉬운 의존성을 하드코딩하지 않고 적절한 모드, 도메인, 워크플로로 안내할 것
- 첫 진입점으로서 독립적으로 사용 가능할 것
- 플러그인 엔트리포인트 skill에는 `<plugin>-guide` 명명 규칙을 사용할 것
- 플러그인이 의도적으로 단일 목적이라면 모든 플러그인에 강제로 둘 필요는 없음

실무적인 기준:

- 플러그인이 분명한 단일 skill 범위를 넘어선다면, 특별한 이유가 없는 한 엔트리포인트 skill을 추가합니다

## 저장소 편집 규칙

- 두 번째 레이아웃 관례를 조용히 도입하지 않습니다.
- 명시적으로 재정렬 요청이 없는 한, 기존 마켓플레이스 순서를 유지합니다.
- 저장소 루트의 release 플러그인은 직접 편집하지 않습니다.
- 루트 release 변경은 build command 산출물로만 만듭니다.
- release command를 통한 version bump는 마지막 `main` merge 이후 해당 plugin의 첫 수정 때만 수행합니다.
- 메타데이터나 경로만 손보면 되는 경우, 스캐폴드 재생성보다 작고 직접적인 수정을 선호합니다.
- 플러그인을 이동했다면 같은 변경 안에서 마켓플레이스 경로도 함께 갱신합니다.
- 스캐폴드 도구가 `./plugins/<plugin-name>`를 생성했다면 마무리 전에 `./<plugin-name>`로 옮깁니다.

---
## Part 2. 하네스 엔지니어링 규칙
이 섹션은 하네스 설계, 구현, 검증, 실패 처리에 적용됩니다.
---

## 하네스 엔지니어링의 목적

이 저장소는 하네스 엔지니어링 관점으로 다룹니다.
영리한 일회성 수정이 아니라 재현 가능성, 진단 가능성, 명시적 계약, 안전한 반복 개선을 최적화합니다.

## 하네스 작업 원칙

- 외부 실서비스 의존성보다 결정론적인 fixture를 우선합니다.
- 실패를 숨기기 쉽게 만들지 말고, 디버깅하기 쉽게 만듭니다.
- 변경은 좁고, 추적 가능하며, 검증하기 쉽게 유지합니다.
- 관측 가능성은 선택적 후속 작업이 아니라 하네스의 일부로 취급합니다.
- 입력, 출력, 실패 모드에는 명시적 계약을 선호합니다.

## 하네스 변경 워크플로

1. 변경하려는 정확한 하네스 동작, 계약, 워크플로를 식별합니다.
2. 수정 전에 관련 스크립트, 설정, fixture, 문서, 테스트를 읽습니다.
3. 목표 문제를 해결하는 가장 작은 완전한 변경을 만듭니다.
4. 위험 수준에 맞는 가장 좁고 의미 있는 검증을 수행합니다.
5. 무엇을 바꿨는지, 어떻게 검증했는지, 무엇이 아직 불확실한지 보고합니다.

## 하네스 설계 규칙

- 결정론적인 입력, fixture, seed 기반 randomness를 선호합니다.
- 실행 간 숨은 상태를 피합니다.
- retry, timeout, backoff는 명시적이고 정당화된 형태로 둡니다.
- infra failure, harness failure, assertion failure를 분명히 구분합니다.
- 하네스를 재실행하거나 트리아지할 가능성이 있다면 ad hoc print보다 structured log 또는 machine-readable output을 선호합니다.
- setup과 teardown을 명시적으로 유지합니다.
- 어떤 워크플로가 결정론적으로 될 수 없다면, 불안정한 경계를 문서화하고 영향 범위를 제한합니다.

## 검증 규칙

- 설정만 바뀐 경우:
  - 파싱 또는 로드 동작을 검증하고, 그 설정이 실행에 영향을 주면 대표 경로 하나를 실행합니다.
- fixture 또는 sample data가 바뀐 경우:
  - 해당 fixture에 의존하는 시나리오만 다시 실행합니다.
- 핵심 하네스 로직이 바뀐 경우:
  - 타깃 테스트와 함께 대표 end-to-end 또는 integration 경로 하나를 추가로 실행합니다.
- 로깅 또는 리포팅이 바뀐 경우:
  - machine-readable output 형태와 운영자 가독성을 모두 검증합니다.

## 실패 처리

- 근본 원인을 이해하지 못한 채 의도되지 않은 blind retry로 flaky 동작을 덮지 않습니다.
- 상태 확인, 이벤트, assertion으로 해결할 수 있는데 wall-clock sleep에 의존하지 않습니다.
- 실패를 로컬에서 재현할 수 없다면, 정확히 어디까지가 불확실한지 기록합니다.
- 외부 시스템이 관련되어 있다면, 하네스 로직을 바꾸기 전에 어느 부분이 비결정적인지 먼저 분리합니다.

## 금지 사항

- seed나 분명한 이유 없이 randomness를 도입하지 않습니다.
- 깨진 하네스 동작을 숨기는 silent fallback을 추가하지 않습니다.
- 하네스 수정 변경에 무관한 cleanup을 섞지 않습니다.
- 테스트를 통과시키기 위해 assertion을 약화하지 않습니다.
- 호출부를 함께 갱신하거나 break를 문서화하지 않은 채 출력 계약을 함부로 바꾸지 않습니다.

## 에이전트 출력 기대사항

하네스 관련 작업을 마칠 때는 다음을 포함합니다.

- `Scope handled`
- `Files changed`
- `Verification`
- `Residual risk`
