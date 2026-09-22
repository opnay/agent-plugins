# Jev 005: Scenario Testing Migration

## 변경사항 요약

- reusable instruction scenario evaluation을 Jev plugin 책임으로 둔다.
- `$jev:scenario-testing` skill과 spec을 추가한다.
- `$advance-codex:skill-scenario-testing` runtime, skill spec, 현재 사용 표면을 제거한다.
- 공통 target을 공유하는 선택적 Jev prediction은 기존 `jev batch`를 사용한다.

## 변경 목적

Scenario testing은 Codex 기능 제작에 한정되지 않는 reusable instruction evaluation입니다. Jev prediction은 선택 사항이며, scenario catalog·fresh executor·comparison·caller verdict·evidence는 `$jev:scenario-testing`이 소유합니다.

## 릴리즈 노트 기준

- 독자에게 보이는 변화: 새 호출 식별자는 `$jev:scenario-testing`입니다.
- 호환성 영향: `$advance-codex:skill-scenario-testing`은 제거하며 alias를 제공하지 않습니다.
- migration 필요 여부: 기존 호출 문서와 자동화는 새 식별자로 변경해야 합니다.
- 운영자/사용자가 알아야 할 검증 결과: `jev batch`는 prediction 전송만 담당하며 executor evaluation이나 verdict를 소유하지 않습니다.

## 기록 근거

- 관련 커밋: 없음.
- 결정 근거: reusable instruction scenario testing을 Jev plugin이 소유하도록 이전한 사용자 승인.

## 변경 상세

### Plugin Boundary

- 목적: reusable instruction scenario evaluation을 Jev plugin에 둡니다.
- 범위: Jev와 Advance Codex의 plugin spec, README, manifest, root README.
- 관련 표면: `jev/`, `advance-codex/`, `src/jev-dev/`, `src/advance-codex-dev/`.
- 검증: 현재 사용 표면에 이전 식별자가 남지 않고 두 plugin의 설명이 실제 bundled skill과 일치해야 합니다.

### Scenario Testing Skill

- 목적: scenario-specific expected behavior, optional Jev prediction, fresh executor actual behavior를 같은 Choice로 비교합니다.
- 범위: `$jev:scenario-testing` spec과 runtime.
- 관련 표면: `src/jev-dev/specs/skills/scenario-testing.md`, `jev/skills/scenario-testing/SKILL.md`.
- 검증: 기본 40개, 명시적 smoke 2~3개, 사용자 지정 수 우선, `unexpected_action` attribution, caller verdict, durable evidence 계약이 일치해야 합니다.

### Previous Skill Removal

- 목적: 같은 workflow를 두 plugin이 동시에 소유하지 않게 합니다.
- 범위: Advance Codex의 기존 runtime skill과 skill spec 삭제.
- 관련 표면: `advance-codex/skills/skill-scenario-testing/SKILL.md`, `src/advance-codex-dev/specs/skills/skill-scenario-testing.md`.
- 검증: 두 경로가 제거되고 Advance Codex의 현재 plugin spec, README, manifest에 이전 skill이 노출되지 않아야 합니다.

## 비목표

- target instruction을 자동 수정하지 않습니다.
- Jev prediction을 executor evaluation이나 caller verdict의 필수 조건으로 만들지 않습니다.
- scenario 실행 artifact를 plugin `changes/`에 보관하지 않습니다.
- 설치, 사용자 설정, API token, commit, push, release를 변경하지 않습니다.

## 호환성 및 마이그레이션

- `$advance-codex:skill-scenario-testing` 호출자는 `$jev:scenario-testing`으로 변경해야 합니다.
- alias나 compatibility shim은 제공하지 않습니다.
- 내부 attribution key `skill_behavior_underspecified`는 기존 evidence와의 비교 가능성을 위해 유지합니다.
- Advance Codex는 scenario-testing을 더 이상 소유하지 않습니다.

## 검증 기준

- Jev와 Advance Codex의 runtime skill을 canonical validator로 검증합니다.
- 두 plugin manifest와 marketplace JSON을 파싱합니다.
- plugin folder name, manifest name, marketplace source path가 일치해야 합니다.
- 활성 spec, runtime, README, manifest에서 `$advance-codex:skill-scenario-testing` 참조가 없어야 합니다.
- skill spec과 runtime, plugin boundary와 공개 사용 표면을 clean-context read-only verifier가 독립적으로 대조합니다.
- `git diff --check`가 통과해야 합니다.

## 정식 규칙 승격 여부

- Jev의 scenario evaluation 책임은 `src/jev-dev/specs/plugin.md`와 `src/jev-dev/specs/skills/scenario-testing.md`가 소유합니다.
- 실행 계약은 `jev/skills/scenario-testing/SKILL.md`와 Jev 사용 표면이 소유합니다.
- 이전 식별자 제거와 migration 범위는 이 change spec에만 남깁니다.
