# SW Kit

`sw-kit`는 행동 계약, 기술 방향, source-level 코드 품질, 코드베이스 유지보수를 분리한 소프트웨어 작업 플러그인입니다.

- `spec`: 구현 가능한 행동 계약
- `engineering`: 기술·데이터·운영 선택
- `code`: 구현·리팩터링·테스트·리뷰와 테스트 가치 기준, 집중 감사, subsystem campaign
- `maintenance`: dependency, deprecation, drift, debt와 테스트 부채 조사

테스트 작성 gate·junk patterns·유지 예외·회귀 확인은 runtime `SKILL.md`가 소유합니다. 집중 감사와 전체 campaign은 `code/references/test-audit.md`, `code/references/test-campaign.md`에서 조건부로 읽습니다. `maintenance`는 이 bundled 계약을 참조하며 수정 책임은 `code`에 있습니다.

상세 계약은 `specs/plugin.md`와 `specs/skills/*.md`가 소유합니다.
