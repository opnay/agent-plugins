# pro-code-keeper Skill Spec

## 목적

`pro-code-keeper`는 Codex가 코드 작성, 디버깅, 리팩터링, 단순화 검토에서 lean senior developer judgment를 적용하게 하는 guidance skill입니다.
핵심 초점은 실제 코드 flow와 계약을 이해한 뒤 가장 작고 안전한 변경, 삭제, 재사용, 표준 기능 사용, 불필요한 추상화 방지를 선택하는 것입니다.
runtime은 `SKILL.md` router와 task-specific `references/`로 구성되어 필요한 분기만 로드합니다.

## 경계

- 포함:
  - 실제 파일, 호출자, 소유 경계, 변경 계약 확인
  - 작은 안전 변경, 삭제 우선, 기존 패턴 재사용 판단
  - standard library, platform/native 기능, 이미 설치된 dependency 우선 판단
  - 불필요한 abstraction, wrapper, config, factory, hook, extension point 식별
  - overengineering review, simplification review, dependency reduction findings
  - `lean:` comment와 lean debt ledger 판단
  - legacy `ponytail:` marker의 debt ledger 입력 처리
  - 위험도에 맞는 최소 검증 선택
- 제외:
  - 조사를 줄이거나 원인 분석을 생략하는 방식의 축소
  - 보안, 접근성, validation, data-loss 방지, 명시 요청 기능 축소
  - 특정 언어/프레임워크 구현 레시피
  - 새 dependency 도입 중심 설계
  - commit, push, PR, release 절차
  - review만 요청된 작업에서 직접 코드 수정

## 처리하려는 작업 형태

- 코드 변경에서 가장 작은 안전 구현 단위를 정해야 하는 작업
- 리팩터링 요청에서 삭제, 재사용, 표준 기능 사용, 추상화 제거 가능성을 검토해야 하는 작업
- repository overengineering audit, removable code review, dependency reduction review
- `lean:` comment, deferred simplification, future expansion point를 목록화해야 하는 작업
- bug fix 또는 root-cause fix에서 증상 패치보다 소유 원인을 좁혀야 하는 작업

## 대표 표면

- 대표 runtime 표면: `judgment-kit-dev/skills/pro-code-keeper/SKILL.md`
- runtime references: `judgment-kit-dev/skills/pro-code-keeper/references/*.md`
- runtime examples: `judgment-kit-dev/skills/pro-code-keeper/examples/*.md`
- 사용자 스펙 의도: `judgment-kit-dev/specs/skills/pro-code-keeper/intent.md`
- skill spec: `judgment-kit-dev/specs/skills/pro-code-keeper/spec.md`

## Runtime Folder 구조

```text
pro-code-keeper/
  SKILL.md
  references/
    init.md
    implement.md
    code-review.md
    repo-audit.md
    root-cause-fix.md
    dependency-check.md
    refactor-shrink.md
    debt-ledger.md
    impact-scoreboard.md
    output-style.md
    safety-boundaries.md
  examples/
    native-before-dependency.md
    reuse-before-rewrite.md
    delete-before-add.md
```

## 핵심 처리 계약

- lean은 code 양과 불필요한 구조를 줄이는 것이며, 문제 이해와 code-flow 조사를 줄이는 것이 아닙니다.
- 코드 작성 전 관련 파일, 호출 flow, caller contract, 소유 경계를 확인합니다.
- 변경은 필요한 기능을 완성하는 가장 작은 surface에 둡니다.
- decision ladder는 no-build, existing helper, standard library, platform/native feature, installed dependency, one clear line, smallest working code 순서로 적용합니다.
- 삭제와 재사용을 선호하되, edge case 안전성이 갈리면 더 안전한 선택을 우선합니다.
- validation, security, accessibility, user-requested features, data-loss prevention, concurrency/time/hardware uncertainty를 단순화 대상으로 삼지 않습니다.
- review 요청에서는 findings를 보고하고 직접 수정하지 않습니다.
- review findings는 `delete`, `stdlib`, `native`, `yagni`, `shrink`로 분류합니다.
- `lean:` comment는 현재 단순화, 한계, upgrade trigger를 함께 기록할 때만 사용합니다.
- 검증은 risk에 맞추고, trivial change에는 과한 test harness를 추가하지 않습니다.
- allow-list style contract를 먼저 정의하고, block-list는 좁은 방어선으로만 둡니다.
- argument 또는 사용자 표현은 semantic branch로 해석합니다. 공개 skill 호출 이름은 `pro-code-keeper`만 사용합니다.

## Runtime 분기 계약

- `SKILL.md`: frontmatter, 최소 철학, 분기표, reference loading rule만 소유합니다.
- `init.md`: 모든 분기의 ctx scope map, 목적 잠금, 조사 범위 결정을 소유합니다.
- `implement.md`: 코드 변경의 smallest safe change 절차를 소유합니다.
- `root-cause-fix.md`: bug/debug/root cause fix의 재현, 원인 좁히기, 증상 패치 방지를 소유합니다.
- `code-review.md`: overengineering/simplification/removable-code review의 read-only finding 계약을 소유합니다.
- `repo-audit.md`: repository-wide audit의 scan 범위, 제외 경로, ranking 계약을 소유합니다.
- `dependency-check.md`: 새 dependency 또는 기존 dependency 제거 판단을 소유합니다.
- `refactor-shrink.md`: 리팩터링·축소 작업의 behavior-preserving 절차를 소유합니다.
- `debt-ledger.md`: `lean:`과 `ponytail:` marker ledger를 소유합니다.
- `impact-scoreboard.md`: audit/review finding의 reduction impact ranking을 소유합니다.
- `output-style.md`: compact final/report format을 소유합니다.
- `safety-boundaries.md`: shrink 금지, 보안/validation/accessibility/data-loss guardrail을 소유합니다.

## 출력 계약

- 코드 변경 시: changed scope, intentionally not built, expansion trigger, verification
- overengineering review 시: tag, location, target, replacement, reason
- repository audit 시: where, what, why, replacement, ranked by reduction impact
- lean debt ledger 시: file, line, current simplification, known limit, upgrade condition, upgrade path
- dependency check 시: keep/remove/add decision, native/stdlib alternative, migration cost, risk, verification
- refactor shrink 시: behavior kept, code removed, replacement, test signal, rollback trigger

## 독립성 원칙

`pro-code-keeper`는 독립 실행 가능한 runtime skill이어야 합니다.
본문은 sibling skill 이름이나 dev-only spec 경로를 읽으라고 지시하지 않습니다.
`pro-engineering`과 함께 쓰일 수 있지만, 작은 안전 변경과 불필요한 구조 축소 판단은 이 skill 본문만으로 수행 가능해야 합니다.
runtime references는 설치 후 skill folder 안에 존재하는 파일만 가리킵니다.

## Description Trigger Metadata

이 skill은 passive skill로 선택될 수 있어야 합니다.
frontmatter `description` 끝에는 `#` 없는 쉼표 구분 plain token 목록을 둡니다.
권장 token tail:

`lean senior dev, smallest safe change, overengineering review, simplify code, delete code, reduce dependencies, lean debt, lean comments`

## 검증 기준

- dev runtime skill이 `skills/pro-code-keeper/SKILL.md`에 존재해야 한다.
- dev runtime skill이 계획된 `references/`와 `examples/` 파일을 포함해야 한다.
- release build 후 root `judgment-kit/skills/pro-code-keeper/SKILL.md`가 존재해야 한다.
- release build 후 root `judgment-kit/skills/pro-code-keeper/references/`와 `examples/`가 존재해야 한다.
- plugin spec, README, manifest prompt가 `pro-code-keeper`의 역할과 사용 기준을 언급해야 한다.
- runtime skill 본문은 dev-only `specs/` 경로나 `src/judgment-kit-dev` 경로를 실행 지시로 포함하지 않아야 한다.
- runtime skill은 작은 안전 변경과 축소 판단을 소유하되, 보안·검증·명시 기능 축소를 권하지 않아야 한다.
- runtime skill frontmatter는 `name`과 `description`만 포함해야 한다.
- runtime skill은 별도 공개 skill 이름을 만들지 않아야 한다.

## 확장 원칙

- child spec은 runtime reference 중 하나가 독립 normative contract로 커질 때만 추가합니다.
- 사용자 의도는 `intent.md`에만 둡니다.
- runtime skill은 설치 후 단독으로 실행 가능한 본문만 포함합니다.
- reference는 절차, 금지사항, 출력 형식만 담고 중복 guardrail은 `safety-boundaries.md`와 `output-style.md`로 모읍니다.
