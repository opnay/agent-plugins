# Codex Creator

`codex-creator`는 Codex plugin과 재사용 skill을 설계·작성·검토하는 확장 가이드입니다. plugin 경계를 먼저 정한 뒤 개별 skill의 책임과 사용 조건을 설계합니다.

- `$codex-creator:plugin-creator`: plugin 경계, 내장 skill 구성, manifest·README·plugin spec의 사용 안내를 정리합니다.
- `$codex-creator:skill-creator`: skill 책임·비목표, description 선택 조건, runtime 계약과 비교 행동 검증을 다룹니다.

두 skill은 각각 canonical system 동명 skill과 함께 사용합니다. 기본 skill은 이 플러그인에 포함되지 않습니다. 기본 skill이 없는 환경에서는 해당 creator의 선행 조건이 충족되지 않습니다.

`skill-creator`는 신규·실질 변경과 테스트 요청에서 일반 사례와 skill 부재 시 실패 가능 사례를 준비해 사용/미사용을 비교합니다. 사용 조건은 격리 작업 폴더의 `.agents/skills`에 현재 대상 skill을 임시 심링크하고 실제 발견·선택·본문 읽기를 확인합니다. 설계·검토 요청에서는 검증 계획과 실제 실행을 구분합니다.

호출 예:

- `$codex-creator:plugin-creator 두 skill의 역할과 사용 안내가 일치하는 플러그인을 설계해 주세요.`
- `$codex-creator:skill-creator 이 skill의 책임과 description 선택 조건을 검토해 주세요.`

기존 `$advance-codex:plugin-creator`, `$advance-codex:skill-creator` 호출은 위 식별자로 바꿔 주세요. `optimize-token`과 `asd-ste100`은 `advance-codex`가 제공합니다.

개발 기준은 [plugin spec](specs/plugin.md), [plugin-creator spec](specs/skills/plugin-creator.md), [skill-creator spec](specs/skills/skill-creator.md)을 따릅니다. 실행 파일은 저장소 루트의 `codex-creator/`에서 관리합니다.
