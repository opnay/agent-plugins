# 마켓플레이스

## 배포 계약

마켓플레이스는 저장소 루트의 플러그인 폴더를 참조합니다.
플러그인 복사본이나 별도 빌드를 만들지 않습니다.

| 소유 표면 | 책임 |
| --- | --- |
| `.agents/plugins/marketplace.json` | 공통 플러그인 목록·순서와 Codex 설치 정책 |
| `.claude-plugin/marketplace.json` | Claude용 목록·소유자와 저장소 기준 상대 경로 |
| `<plugin>/.codex-plugin/plugin.json` | 공통 이름·버전·메타데이터와 Codex interface |
| `<plugin>/.claude-plugin/plugin.json` | Claude용 메타데이터와 같은 `skills/` 경로 |
| `<plugin>/skills/` | 두 환경이 공유하는 실행 지침·스크립트·리소스 |

Claude 목록은 공통 목록의 이름·순서·대상 폴더를 유지합니다.
이름·버전·author·license·keywords·공통 description은 원본 매니페스트와 맞춥니다.
Claude의 `displayName`은 Codex의 `interface.displayName`을 사용합니다.
Codex 전용 `interface`와 marketplace `policy`는 Claude JSON에 넣지 않습니다.
플러그인 버전은 각 매니페스트에 두며 marketplace 항목에 중복하지 않습니다.
이름·버전·목록을 변경할 때는 두 플랫폼의 해당 파일을 함께 갱신합니다.
이 계약은 배포 형식을 소유하며 개별 skill의 책임이나 실행 계약을 변경하지 않습니다.

## 등록과 설치

마켓 등록은 설치 가능한 플러그인 목록을 Claude에 알려줍니다.
각 플러그인은 별도 설치로 선택합니다. 마켓을 등록해도 8개 플러그인이 모두 설치되지는 않습니다.
같은 이름의 다른 마켓이 등록되어 있다면 먼저 source를 확인합니다.

저장소 루트에서 사용하는 CLI로 로컬 마켓을 등록합니다.

```sh
codex plugin marketplace add .
```

Codex에서는 `/plugins`에서 플러그인을 선택하고, 목록 갱신에는 `codex plugin marketplace upgrade`를 사용합니다.
GitHub에서 등록하려면 `codex plugin marketplace add opnay/agent-plugins`를 사용합니다.

Claude Code 명령은 다음과 같습니다.

```sh
claude plugin marketplace add .
```

필요한 플러그인만 설치합니다. 예를 들어 SW Kit은 다음과 같습니다.

```sh
claude plugin install sw-kit@opnay-plugins
```

새 세션에서 `/sw-kit:spec` 또는 `/sw-kit:code`를 호출합니다.
설치 상태는 `claude plugin marketplace list --json`과 `claude plugin list --json`으로 확인합니다.

로컬 경로에서 등록한 마켓의 상대 경로 플러그인은 원본 파일을 사용합니다.
수정 후에는 새 세션을 시작하거나 `/reload-plugins`를 실행합니다.
GitHub에서 등록한 마켓은 배포된 저장소와 플러그인 캐시를 사용하므로 로컬 미커밋 변경은 반영되지 않습니다.

Claude용 파일이 `main`에 배포된 뒤에는 GitHub로 등록할 수 있습니다.

```sh
claude plugin marketplace add opnay/agent-plugins
```

로컬 등록과 GitHub 등록은 같은 `opnay-plugins` 이름을 사용하므로 동시에 서로 다른 마켓으로 취급하지 않습니다.

## 스킬 사용과 호환성

- Claude의 사용자 호출은 `/<plugin>:<skill>`입니다. 예: `/advance-codex:optimize-token`.
- 공유 본문의 `$<plugin>:<skill>`은 스킬 식별자이며 Claude의 슬래시 명령 표기와 구분합니다. 에이전트는 사용 가능한 Skill 도구의 해당 이름으로 호출합니다.
- `$<plugin>:<skill>/<resource>`는 실제 절대 파일 경로가 아닙니다. 소유 스킬을 로드해 확인한 설치 디렉터리를 기준으로 리소스를 읽습니다. Codex 캐시 경로를 Claude에서 사용하지 않습니다.
- `advance-codex`의 `skill-creator`와 `plugin-creator`는 canonical 기본 스킬을 확장합니다. 이 마켓은 기본 스킬을 제공하거나 설치하지 않습니다. 기본 스킬이 없으면 해당 creator 계약의 선행 조건이 충족되지 않습니다. `optimize-token`과 `asd-ste100`은 이 의존성을 갖지 않습니다.
- `advance-subagent`의 모델 선택은 적용 중인 지침이 소유합니다. Claude에서는 현재 Agent 도구와 모델 프로필을 따르고 Codex 전용 도구 인자를 그대로 전달하지 않습니다.
- `app-extensions`는 Figma 연결을 제공하지 않습니다. 호스트에서 사용 가능한 Figma 기능이 필요합니다.
- `toolkit`과 `agent-wiki`의 실행 파일은 각 스킬의 설치·권한 계약을 따릅니다. 플러그인 설치는 CLI 설치나 Git alias 변경을 자동 실행하지 않습니다.

JSON 검증은 패키징 형식과 경로를 확인합니다. 실제 모델의 스킬 선택, 외부 앱 연결, CLI 실행까지 보장하지 않습니다.

## 검증

저장소 루트에서 목록과 각 플러그인을 검증합니다.

```sh
claude plugin validate . --strict
claude plugin validate ./sw-kit --strict
```

공통 목록의 모든 플러그인에 두 번째 검증을 적용합니다.
추가로 JSON 파싱, 이름·버전·목록 순서·대상 폴더·공유 스킬 경로의 일치와 `git diff --check`를 확인합니다.
기존 Codex JSON과 `skills/` 변경 여부도 확인합니다.
공개 디렉터리 제출용 license·icon 등의 검사는 로컬 로딩 검증과 구분하며 제출은 별도 범위입니다.

공식 기준: [마켓 생성](https://code.claude.com/docs/en/plugin-marketplaces), [마켓 스키마](https://code.claude.com/docs/en/plugins/marketplace-reference), [플러그인 스키마](https://code.claude.com/docs/en/plugins-reference).
