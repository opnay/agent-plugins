# SW Kit

`sw-kit`는 요구를 소프트웨어로 만들고 오래 유지하는 네 단계의 판단을 제공합니다.

- `$sw-kit:spec`: 요구를 기술 중립 행동 계약으로 정리합니다.
- `$sw-kit:engineering`: 구조, 데이터, 기술, 연동, migration, 운영 방향을 정합니다.
- `$sw-kit:code`: 선택된 방향 안에서 production code와 테스트를 구현·수정·리뷰하고, 테스트의 독립적인 검증 가치를 판단합니다.
- `$sw-kit:maintenance`: dependency, deprecation, dead code, drift, debt와 테스트 부채를 조사하고 정리 방향을 만듭니다.

테스트 작성·수정·리뷰 기준은 `code/SKILL.md` 본문에 있습니다. 기존 테스트의 삭제·통합 감사에서만 `test-audit.md`를, 전체 subsystem 정리에서만 `test-campaign.md`를 읽습니다. 테스트 개수 대신 보호하는 동작·회귀·독립 계약을 기준으로 판단합니다.

- 작성·리뷰: `$sw-kit:code`로 테스트를 작성하거나 검토합니다.
- 집중 감사: `$sw-kit:code`로 지정한 테스트의 중복과 실제 오류 검출 능력을 확인합니다. 발견 단계는 read-only입니다.
- 전체 정리: `$sw-kit:code`로 subsystem의 전체 테스트 campaign을 요청합니다.
- 유지보수 조사: `$sw-kit:maintenance`로 테스트 부채와 정리 계획을 요청합니다. 같은 `code` 기준을 사용합니다.

`$judgment-kit:pro-planner`는 사용자 문제, 가치, 범위, 우선순위와 제품 handoff를 계속 소유합니다. Git workflow는 `$toolkit:git`의 범위입니다.
