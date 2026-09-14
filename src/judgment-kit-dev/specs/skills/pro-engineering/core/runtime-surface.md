# Runtime Surface Contract

## 책임

이 문서는 `pro-engineering` runtime의 독립 실행 가능성, trigger metadata, dev/release 검증 기준을 소유합니다.

## 독립성 원칙

- `pro-engineering`은 독립 실행 가능한 runtime skill이어야 합니다.
- runtime 본문은 sibling skill 이름이나 dev-only spec·source 경로를 읽으라고 지시하지 않습니다.
- 다른 `judgment-kit` skill과 함께 사용할 수 있지만, 문제 해결과 코드 작성 판단 자체는 이 skill 본문만으로 수행 가능해야 합니다.
- runtime 본문은 특정 언어·프레임워크 레시피, 고정 아키텍처, 제품 정책 결정 절차로 좁아지지 않습니다.

## Description Trigger Metadata

이 skill은 passive skill로 선택될 수 있어야 합니다.
frontmatter는 `name`, `description`만 사용합니다.
`description` 끝에는 `#` 없는 쉼표 구분 plain token 목록을 둡니다.
권장 token tail:

`engineering judgment, problem solving, root cause analysis, technical reasoning, code quality, implementation discipline`

## 검증 기준

- dev runtime skill이 `skills/pro-engineering/SKILL.md`에 존재해야 합니다.
- release build 후 root `judgment-kit/skills/pro-engineering/SKILL.md`가 존재해야 합니다.
- plugin spec, README, manifest prompt가 `pro-engineering`의 역할과 사용 기준을 언급해야 합니다.
- runtime skill 본문은 dev-only `specs/` 경로나 `src/judgment-kit-dev` 경로를 실행 지시로 포함하지 않아야 합니다.
- runtime skill은 문제 해결 중심이며 특정 언어/프레임워크 레시피로 좁아지지 않아야 합니다.
- runtime skill은 완결 산출물과 모듈·재사용 경계의 허용 조건을 포함하고, 모든 구현에 단계나 package 분리를 강제하지 않아야 합니다.
- runtime skill은 소비자 간 정책 차이를 실제 중복으로 단정하거나 미확정 제품 정책을 공유 모듈에 고정하지 않아야 합니다.
- runtime skill은 도메인·조정·기반을 책임 렌즈로 구분하되 고정 계층으로 강제하지 않고, 단일 소유자 규칙은 분리 비용보다 현재의 유지보수·검증 이익이 클 때만 추출해야 합니다.
- runtime skill은 도메인 책임에 외부 전송·저장·화면 상태가 새거나, 조정 책임이 제품 정책을 정하거나, 기반 책임이 제품·화면 소유권에 의존하는 역방향 결합을 피해야 합니다.
- runtime skill은 제품 계약이나 외부 표현을 아는 어댑터·저장소를 저수준이라는 이유만으로 기반 책임에 합치지 않아야 합니다.
