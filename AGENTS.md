# AGENTS.md

이 문서는 이 저장소의 운영 규칙입니다.

---

## Part 1. 저장소 및 마켓플레이스 규칙

## 저장소 역할

이 저장소는 사용자가 직접 만드는 Codex 플러그인을 관리합니다.
로컬 플러그인 마켓플레이스 루트이자 하네스 엔지니어링 작업 공간입니다.
제품화 저장소가 아니며, 운영 가이드와 품질 기준을 함께 소유합니다.

## 저장소 레이아웃

- 공개 설치용 release 플러그인: 저장소 루트 `./<plugin-name>`
- 개발 원본 플러그인: `./src/<plugin-name>-dev`
- specs 위치: `./src/<plugin-name>-dev/specs/`
- 기본 편집 대상: `./src/<plugin-name>-dev`
- release 플러그인 갱신: build command 산출물만 사용합니다.
- `./plugins/<plugin-name>` 경로는 만들거나 사용하지 않습니다.

## 브랜치 운영

- `main`: 공개 release 브랜치
- `next`: 개발 브랜치
- 일반 플러그인 수정은 `next`의 `src/<plugin-name>-dev`에 적용합니다.
- `main`에는 검증된 내용을 release로 승격할 때만 반영합니다.
- `next`에서도 dev 플러그인을 설치/사용할 수 있어야 하므로 `<plugin-name>-dev` 이름과 `src/<plugin-name>-dev` 구조를 유지합니다.
- 플러그인 변경마다 build command로 루트 release surface를 갱신합니다.

## 버전 승격

- 마지막 `main` merge 이후 특정 dev 플러그인을 처음 수정하면 먼저 version bump 여부를 확인합니다.
- bump 종류는 자동 단정하지 않습니다. 사용자에게 patch/minor/major 또는 구체 version을 묻습니다.
- 사용자가 version 유지 또는 bump를 명시하면 그 결정을 따릅니다.
- 같은 플러그인의 이후 변경은 추가 bump 없이 build만 수행합니다.
- 공개 release 단계에서는 `next`를 `main`으로 merge하고, release surface를 루트 폴더로 생성/갱신합니다.

## 필수 플러그인 구조

개발 원본 플러그인은 다음을 포함해야 합니다.

- `./src/<plugin-name>-dev/.codex-plugin/plugin.json`
- `./src/<plugin-name>-dev/README.md`
- `./src/<plugin-name>-dev/specs/plugin.md`
- 선택: `specs/skills/`, `skills/`, `assets/`, `scripts/`, `.mcp.json`, `.app.json`

플러그인 폴더 이름과 `plugin.json`의 `"name"` 값은 일치해야 합니다.
공개 release 플러그인은 `<plugin-name>`, 개발 원본은 `<plugin-name>-dev`를 사용합니다.

## SDD 규칙

- SDD 상세 규칙은 `docs/SDD.md`가 소유합니다.
- 플러그인 작업은 spec을 먼저 확인하거나 갱신한 뒤 구현 표면을 맞춥니다.
- 플러그인 표면이 바뀌면 `README.md`, `specs/plugin.md`, 관련 skill spec, `plugin.json`을 함께 점검합니다.
- change spec을 정식 규칙으로 승격하라는 요청은 repo-level 규칙, `docs/SDD.md`, 또는 소유 spec에 반영합니다.
- folder-based skill spec의 사용자 스펙 의도는 `intent.md`가 소유합니다. `spec.md`와 child spec에는 반복하지 않습니다.
- 스펙 문서는 `advance-codex:optimize-token extreme` 규격으로 작성합니다.
- skill spec이 바뀌면 기존 runtime skill 본문을 부분 패치하지 않고, 현재 spec 기준으로 `SKILL.md`를 처음부터 재작성합니다.
- clean-context subagent는 runtime skill 작성 대행자가 아닙니다. spec/runtime 정합성의 독립 read-only 검증에 사용합니다.

## 마켓플레이스

- 단일 진실 공급원: `./.agents/plugins/marketplace.json`
- 공개 release 플러그인은 모두 marketplace 항목을 가져야 합니다.
- `source.path`는 저장소 루트의 release 플러그인 폴더를 가리켜야 합니다.
- 등록 순서는 `plugins` 배열을 기준으로 유지합니다.
- 이 저장소 로컬 플러그인을 `./plugins/<plugin-name>`로 등록하지 않습니다.
- 공개 release 항목 추가나 갱신은 release 승격이 요청된 경우에만 수행합니다.
- 모든 marketplace 항목은 `policy.installation`, `policy.authentication`, `category`를 포함해야 합니다.

## 플러그인 변경 Workflow

1. 개발 원본은 `./src/<plugin-name>-dev`에 생성하거나 이동합니다.
2. README, specs, skills, manifest를 dev source에서 먼저 수정합니다.
3. skill spec을 수정했다면 runtime skill을 현재 spec 기준으로 처음부터 재작성합니다.
4. skill spec 변경은 `src/<plugin-name>-dev/changes/<version>.md`에 기록합니다.
5. 관련 skill spec, plugin spec, upstream/downstream plugin surface를 함께 점검합니다.
6. `pnpm build:plugin <plugin-name> [--force]`로 루트 release surface를 갱신합니다.
7. JSON 변경은 파싱 검증합니다.

Version bump가 필요한 release 승격은 다음 중 하나를 사용합니다.

- `pnpm release:plugin <plugin-name> --bump <patch|minor|major> [--force]`
- `pnpm release:plugin <plugin-name> --version <version> [--force]`

## Release 승격 Workflow

1. `next`에서 release 범위를 확인합니다.
2. dev plugin version bump 반영 여부를 확인합니다.
3. 마지막 `main` merge 이후 첫 수정인 플러그인은 사용자에게 bump 종류를 확인합니다.
4. 이후 같은 플러그인 수정은 추가 bump 없이 build만 수행합니다.
5. `next`를 `main`으로 merge할 release 단위를 확정합니다.
6. build command로 `src/<plugin-name>-dev`를 루트 release surface로 변환합니다.
7. release root에 `specs/`가 없는지, manifest와 marketplace가 맞는지 검증합니다.
8. 검증된 release 변경만 `main`에 반영합니다.

## 플러그인 설계 기준

- 플러그인은 하나의 일관된 번들로 독립적으로 이해 가능해야 합니다.
- `advance-codex`는 사용자가 직접 만들고 유지할 수 있는 Codex 기능을 깊게 관리하기 위한 플러그인입니다.
- `advance-codex`의 대표 표면은 skill, tool-use guidance, plugin, subagent입니다.
- 명시적 의도 없이 `advance-codex`를 무관한 워크플로나 cross-plugin 유틸리티 범위로 넓히지 않습니다.
- 여러 사용자 지향 skill이 있으면 `plugin.json`, `defaultPrompt`, `README.md`, `specs/plugin.md`에서 각 skill의 역할과 시작점을 분명히 드러냅니다.
- 플러그인 경계를 먼저 정의하고, 그 안의 skill을 정의합니다.
- 느슨한 skill 묶음에서 출발해 나중에 플러그인 모양을 억지로 맞추지 않습니다.
- 플러그인은 자기 목적을 설명하기 위해 무관한 외부 플러그인 구조에 기대지 않습니다.

## Skill 경계와 독립성

- 모든 skill에 독립 실행 가능성을 일괄 강제하지 않습니다.
- 각 skill spec은 독립성을 강제하는지, sibling context를 허용하는지, 이유를 명시합니다.
- skill 책임은 분명해야 하며, sibling skill이나 숨은 사용 문맥을 전제로 끌어오지 않도록 spec에서 경계를 설명합니다.
- plugin-scoped skill 변경은 개별 skill 수정이면서 plugin surface 수정일 수 있습니다.
- 플러그인 소속 skill을 수정할 때는 plugin 안의 역할, sibling 관계, 관련 spec을 함께 확인합니다.
- 플러그인 관점의 영향이 있으면 skill 단독 수정으로 축소하지 않습니다.
- Plugin을 통해 설치된 skill은 `$<plugin>:<skill>` 식별자를 사용합니다. 예: `$advance-codex:skill-creator`
- skill이 에이전트 행동을 보정할 때는 금지 조건을 계속 늘리기보다 허용되는 행동, 판단 기준, 실패 조건을 먼저 정의합니다.
- blocklist 방식의 차단 조건은 허용 계약을 보완하는 좁은 방어선으로만 둡니다.

## Runtime Skill과 Spec 분리

- `src/<plugin-name>-dev/specs/`는 개발 원본 계약입니다. 설치되는 runtime skill 본문이 의존할 수 있는 표면이 아닙니다.
- release surface에는 `specs/`가 포함되지 않는다는 전제로 skill 본문을 작성합니다.
- skill 본문에는 설치 후 존재하지 않는 dev-only spec 경로를 실행 지시로 남기지 않습니다.
- skill을 작성하거나 재작성할 때 spec은 작성 기준으로만 사용합니다.
- 결과물인 `SKILL.md`에는 runtime에서 접근 가능한 본문, `references/`, `templates/`만 남깁니다.
- 상세 계약이 runtime에도 필요하면 본문에 간결히 포함하거나 `skills/<skill-name>/references/`로 승격합니다.
- spec 변경 후 runtime skill은 clean-context subagent에게 재생성을 맡기지 않고, 현재 spec 기준으로 처음부터 재작성합니다.

## Plugin Usage 표면

다음 표면이 plugin usage guidance를 소유합니다.

- `specs/plugin.md`: 플러그인 경계, 내장 skill 체계, 각 skill 시작 기준
- `README.md`: 사람이 읽는 사용 방법과 대표 호출 예시
- `.codex-plugin/plugin.json`: 설치 후 노출되는 설명과 `defaultPrompt`

사용 표면을 바꾸면 남은 skill spec, README, plugin spec, manifest 설명, 호출 예시, release build를 함께 점검합니다.

## 저장소 편집 규칙

- 두 번째 레이아웃 관례를 조용히 도입하지 않습니다.
- 명시적 재정렬 요청이 없으면 marketplace 순서를 유지합니다.
- 루트 release 플러그인은 직접 편집하지 않습니다.
- 루트 release 변경은 build command 산출물로만 만듭니다.
- release command를 통한 version bump는 마지막 `main` merge 이후 해당 plugin의 첫 수정 때만 수행합니다.
- 메타데이터나 경로만 손보면 스캐폴드 재생성보다 작고 직접적인 수정을 선호합니다.
- 플러그인을 이동하면 같은 변경 안에서 marketplace 경로도 갱신합니다.
- 스캐폴드 도구가 `./plugins/<plugin-name>`를 만들었다면 마무리 전에 `./<plugin-name>`로 옮깁니다.

## Clean Context

- clean context는 이전 대화, 이전 subagent 결과, main agent 결론, 임시 작업 맥락을 전제로 삼지 않습니다.
- 전달 내용은 source of truth, 대상 파일, 편집 가능 범위, 금지 범위, 검증 조건, 출력 계약으로 제한합니다.
- prior worker output, main agent의 suspected finding, git diff 기반 복구 방향, 이전 실패 서사를 전달하지 않습니다.
- 기존 subagent를 interrupt하거나 재사용하는 것은 clean context가 아닙니다. 필요하면 새 subagent를 spawn합니다.
- clean context는 spec 기준 작성 결과를 검증할 때 사용할 수 있지만, runtime skill 작성이나 재작성 자체를 맡기지 않습니다.
- spec/runtime 정합성 검증에는 가능한 한 여러 clean-context verifier를 사용하고, 각 verifier에는 좁은 spec 표면이나 책임 영역을 배정합니다.
- verifier subagent는 read-only입니다. `SKILL.md` 재작성, 수정, 빌드, 커밋 권한이 없습니다.

## 문서 맵

문서 구조가 바뀌면 이 맵을 함께 갱신합니다.

```text
.
├── AGENTS.md
└── docs/
    ├── SDD.md
    ├── release-pattern.md
    └── templates/
        ├── plugin-spec.md
        ├── skill-spec.md
        └── change-spec.md
```
