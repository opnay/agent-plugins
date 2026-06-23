# code-quality 시나리오 세트

## Target Instruction

- 대상 skill: `src/code-quality-dev/skills/code-quality/SKILL.md`
- 대상 release skill: `code-quality/skills/code-quality/SKILL.md`
- 목적: `code-quality`가 production code 구현, 수정, 리팩터링, 테스트, 의존성 판단, 코드 리뷰에서 의도대로 작동하는지 검증한다.

## Checklist Design

### 호출되어야 하는 시나리오 공통 체크리스트

- [critical] 기존 코드, 테스트, 호출부, 저장소 규칙을 먼저 확인한다.
- [critical] 새 구현이나 의존성 추가 전에 기존 모듈, 프레임워크/시스템 API, 표준 라이브러리, 설치된 의존성, 유틸리티를 조사한다.
- 재사용을 강제하지 않고 의미, 계약, 오류 모델, 소유권, 변경 주기, 장기 비용을 평가한다.
- 정확성, 보안, 호환성, 데이터 무결성, 운영 안정성을 코드 미관보다 우선한다.
- 변경 후 관련 테스트, 타입 검사, 린트, 빌드 등 가능한 검증을 실행하거나 못 한 이유를 보고한다.

### 리뷰 시나리오 공통 체크리스트

- [critical] 실제 결함과 유지보수 위험을 취향보다 우선한다.
- 심각도, 위치, 문제 이유, 실패 상황, 최소 수정 방향을 제시한다.
- 기존 기능 중복, 잘못된 유틸리티 재사용, 불필요한 의존성, 위험한 직접 구현, 추측성 추상화를 확인한다.
- 확실하지 않은 내용은 가정과 신뢰도를 표시한다.
- 문제가 없으면 억지로 지적하지 않는다.

### 호출되면 안 되는 시나리오 공통 체크리스트

- [critical] prose-only, 번역, generated/vendor code 직접 수정, formatting-only, style-only, naming-only 요청에는 기본적으로 skill을 적용하지 않는다.
- 사용자가 명시적으로 `$code-quality`를 호출하거나 correctness/security/compatibility/data-integrity/operational-risk 검토를 요청한 경우에만 예외적으로 적용한다.
- 코드 품질 판단이 필요한 경계 사례는 조건부로 분류한다.

## Scenario Set

| ID | Trigger | Category | Prompt | Expected focus |
|---|---|---|---|---|
| CQ-001 | true | explicit | `$code-quality`로 결제 금액 반올림 버그를 수정하고 회귀 테스트를 추가해줘. | 명시 호출, 금액/정밀도, 기존 결제 규칙 조사 |
| CQ-002 | true | explicit | `$code-quality` 이 PR에서 P0/P1 수준의 실제 장애나 데이터 손실로 이어질 수 있는 코드만 리뷰해줘. | 리뷰 모드, P0/P1 심각도 중심 |
| CQ-003 | true | explicit | `$code-quality` 사용자 입력으로 파일 경로를 만드는 부분을 안전하게 고쳐줘. | 보안, path traversal, 기존 path API |
| CQ-004 | true | explicit | `$code-quality` 새 로깅 래퍼를 만들기 전에 프레임워크 로거로 충분한지 판단해줘. | wrapper 필요성 판단 |
| CQ-005 | true | explicit | `$code-quality` 기존 `DateHelper`를 써서 타임존 변환을 처리해도 되는지 봐줘. | 유틸리티 계약 검사 |
| CQ-006 | true | explicit | `$code-quality` 이 캐시 구현이 동시 요청에서 안전한지 검토하고 수정해줘. | 동시성, 상태 소유권 |
| CQ-007 | true | explicit | `$code-quality` UUID 생성 때문에 새 패키지를 넣자는 변경을 검토해줘. | 표준 라이브러리, 의존성 비용 |
| CQ-008 | true | explicit | `$code-quality` 실패를 삼키는 try/catch를 저장소 오류 모델에 맞게 고쳐줘. | 오류 처리, 진단 정보 |
| CQ-009 | true | explicit | `$code-quality` 두 API 클라이언트의 중복 재시도 로직을 공통화해도 되는지 판단해줘. | DRY 적합성, retry 계약 |
| CQ-010 | true | explicit | `$code-quality` 직접 만든 JWT 검증 로직을 리뷰하고 안전한 대안을 제안해줘. | 보안 프로토콜 직접 구현 금지 |
| CQ-011 | true | implementation | 주문 취소 API에 부분 취소 기능을 추가해줘. | 도메인 규칙, API 호환성, 테스트 |
| CQ-012 | true | implementation | 사용자 프로필 저장 시 display name 정규화를 추가해줘. | 기존 도메인 모듈, 이름 규칙 |
| CQ-013 | true | implementation | CSV 업로드에서 빈 줄과 잘못된 인코딩을 처리해줘. | 경계값, 표준/라이브러리 파서 |
| CQ-014 | true | implementation | 결제 webhook 중복 수신을 idempotent하게 처리해줘. | 데이터 무결성, 재시도, 트랜잭션 |
| CQ-015 | true | implementation | API 응답에 pagination cursor를 추가해줘. | 공개 API 호환성, 기존 pagination 패턴 |
| CQ-016 | true | implementation | S3 파일 다운로드에 timeout과 취소 처리를 넣어줘. | 시스템/API 기능, 자원 정리 |
| CQ-017 | true | implementation | 사용자 검색에 이메일과 이름 필터를 추가해줘. | 쿼리 안전성, 인덱스, 입력 검증 |
| CQ-018 | true | implementation | 관리자 권한 확인을 새 엔드포인트에 추가해줘. | 인증/인가 프레임워크 패턴 |
| CQ-019 | true | implementation | 비밀번호 재설정 토큰 만료 처리를 구현해줘. | 보안, 시간 처리, 저장 데이터 |
| CQ-020 | true | implementation | 대량 알림 발송을 배치 처리로 바꿔줘. | 성능, backpressure, 재시도 |
| CQ-021 | true | implementation | URL에서 도메인을 추출하는 함수를 추가해줘. | 표준 URL API vs 직접 파싱 |
| CQ-022 | true | implementation | 업로드 파일 확장자 검증은 UI 힌트로만 쓰고, 서버 보안 경계에서는 MIME/content 검증까지 포함해 구현해줘. | MIME/content 보안, 기존 validator |
| CQ-023 | true | implementation | 주문 목록 API에 정렬 옵션을 추가해줘. | 입력 경계, 호환성, 테스트 |
| CQ-024 | true | implementation | 기존 refresh-token/session-store/cookie 모델을 먼저 확인하고 replay와 session fixation을 막는 세션 갱신 로직을 구현해줘. | 프레임워크 인증 API, 보안 |
| CQ-025 | true | implementation | 내부 이벤트 발행 실패 시 재시도 큐에 넣어줘. | 부수 효과, 실패 흐름 |
| CQ-026 | true | implementation | 사용자가 삭제되면 관련 캐시를 무효화해줘. | 상태 변경, 일관성 |
| CQ-027 | true | implementation | 통화별 소수점 자리수를 반영해 금액 포맷을 구현해줘. | 국제화/통화, 표준 API |
| CQ-028 | true | implementation | upstream 외부 API가 `429`를 반환할 때 retry-after와 백오프 정책을 반영해 처리해줘. | 외부 API 오류 모델 |
| CQ-029 | true | implementation | 요청 body 크기 제한을 적용해줘. | 프레임워크 설정, 보안 |
| CQ-030 | true | implementation | 임시 파일 생성과 삭제를 안전하게 처리해줘. | 시스템 API, cleanup |
| CQ-031 | true | bugfix | 날짜가 DST 전환일에 하루 밀리는 버그를 고쳐줘. | 시간대, 표준 라이브러리, 회귀 테스트 |
| CQ-032 | true | bugfix | 결제 실패인데 주문 상태가 paid로 남는 버그를 수정해줘. | 트랜잭션, 데이터 무결성 |
| CQ-033 | true | bugfix | 빈 검색어에서 전체 테이블을 스캔하는 문제를 고쳐줘. | 성능, 입력 경계 |
| CQ-034 | true | bugfix | 네트워크 오류를 성공 응답으로 변환하는 코드를 고쳐줘. | 오류 처리, 실패 전파 |
| CQ-035 | true | bugfix | 동시에 같은 쿠폰을 쓰면 중복 할인되는 버그를 수정해줘. | 동시성, 락/트랜잭션 |
| CQ-036 | true | bugfix | 이메일 대소문자 때문에 중복 계정이 만들어져. 수정해줘. | 정규화, 도메인 계약 |
| CQ-037 | true | bugfix | JSON 파싱 실패가 500으로만 떨어져서 원인을 알 수 없어. 고쳐줘. | 진단 가능성, 사용자 오류 구분 |
| CQ-038 | true | bugfix | 파일 업로드 실패 후 임시 파일이 남는 문제를 고쳐줘. | 자원 정리, 실패 경로 |
| CQ-039 | true | bugfix | 캐시된 권한 때문에 탈퇴 사용자가 계속 접근할 수 있어. 수정해줘. | 보안, 캐시 무효화 |
| CQ-040 | true | bugfix | 금액 합계가 floating point 때문에 틀려. 수정해줘. | 숫자 정확성, 기존 money 타입 |
| CQ-041 | true | bugfix | retry가 중복 결제를 만들 수 있어. 안전하게 고쳐줘. | idempotency, 외부 시스템 |
| CQ-042 | true | bugfix | 모바일에서 빈 문자열이 null과 다르게 처리돼 저장 실패가 나. | 입력 모델, 호환성 |
| CQ-043 | true | bugfix | 잘못된 토큰도 일부 API에서 통과해. 인증 미들웨어를 확인해줘. | 인증 경계, 프레임워크 사용 |
| CQ-044 | true | bugfix | 로그에 access token이 찍히는 문제를 고쳐줘. | 비밀 정보 노출 |
| CQ-045 | true | bugfix | 오래된 마이그레이션 데이터에서 필드가 없어 크래시가 나. | 저장 데이터 호환성 |
| CQ-046 | true | refactor | 이 거대한 함수에서 외부 I/O와 계산 로직을 분리해줘. | 부수 효과 분리, 테스트 가능성 |
| CQ-047 | true | refactor | 세 군데 있는 비슷한 validation을 공통 모듈로 빼도 되는지 보고 적용해줘. | 동일 계약 확인 |
| CQ-048 | true | refactor | `utils`에 있는 범용 함수를 도메인 모듈 근처로 옮기는 게 맞는지 봐줘. | 소유권, 변경 주기 |
| CQ-049 | true | refactor | 중첩 조건이 깊은 권한 체크 코드를 이해하기 쉽게 정리해줘. | 제어 흐름, 도메인 표현 |
| CQ-050 | true | refactor | 모든 서비스에 인터페이스를 만들자는 리팩터링을 검토해줘. | 추측성 추상화 |
| CQ-051 | true | refactor | 결제와 프로필이 같이 쓰는 `normalize` helper를 분리해줘. | 잘못된 DRY 해소 |
| CQ-052 | true | refactor | `manager`라는 이름의 클래스를 책임이 드러나게 바꿔줘. | 이름, 범위 최소화 |
| CQ-053 | true | refactor | 기존 동작 유지하면서 저장소 관례에 맞게 모듈 경계를 정리해줘. | 동작 보존, 검증 |
| CQ-054 | true | refactor | API 클라이언트마다 다른 오류 변환을 일관되게 정리해줘. | 오류 모델 통일 |
| CQ-055 | true | refactor | framework logger를 감싼 얇은 wrapper를 제거할지 판단해줘. | wrapper 책임 여부 |
| CQ-056 | true | refactor | 날짜 문자열 처리 코드를 표준 API로 바꿔줘. | 표준 라이브러리, 호환성 |
| CQ-057 | true | refactor | 테스트하기 어렵다는 이유로 production 코드를 과하게 쪼갠 부분을 정리해줘. | 테스트 가능성 vs 설계 왜곡 |
| CQ-058 | true | refactor | 공통 base class가 오히려 두 도메인을 묶고 있어. 분리해줘. | 결합도, 변경 이유 |
| CQ-059 | true | refactor | 요청과 무관한 포맷 변경 없이 이 모듈만 안전하게 리팩터링해줘. | 변경 범위 통제 |
| CQ-060 | true | refactor | deprecated 라이브러리 API 사용 부분을 새 API로 옮겨줘. | 의존성 호환성 |
| CQ-061 | true | review | 이 diff가 사용자 삭제 흐름에서 데이터 손실을 만들 수 있는지 리뷰해줘. | P0/P1 위험, 데이터 무결성 |
| CQ-062 | true | review | 이 PR에서 보안 관련 문제만 찾아줘. | 보안 중심 리뷰 |
| CQ-063 | true | review | 새 외부 패키지 추가가 과한지 리뷰해줘. | dependency risk |
| CQ-064 | true | review | 리팩터링 PR인데 동작 변경이 섞였는지 봐줘. | 변경 범위, 회귀 |
| CQ-065 | true | review | 테스트가 구현 세부사항에 너무 묶였는지 리뷰해줘. | 테스트 품질 |
| CQ-066 | true | review | API response schema 변경이 호환성을 깨는지 확인해줘. | 공개 API 호환성 |
| CQ-067 | true | review | 이 코드가 기존 도메인 모듈을 무시하고 규칙을 중복 구현했는지 봐줘. | 중복 도메인 규칙 |
| CQ-068 | true | review | 이 exception handling이 원인을 잃어버리는지 리뷰해줘. | 오류 진단 |
| CQ-069 | true | review | 동시 요청에서 race condition이 생길 수 있는지 봐줘. | 동시성 |
| CQ-070 | true | review | 새 wrapper 계층이 실제 책임이 있는지 리뷰해줘. | 추상화 근거 |
| CQ-071 | true | review | logging 변경이 개인정보를 노출할 수 있는지 봐줘. | PII/secret 노출 |
| CQ-072 | true | review | 이 변경이 성능 병목을 만들 수 있는지 확인해줘. | 성능, scale risk |
| CQ-073 | true | review | `StringUtils` 재사용이 도메인 의미에 맞는지 리뷰해줘. | utility semantic fit |
| CQ-074 | true | review | 직접 만든 CSV parser가 안전한지 리뷰해줘. | parser 직접 구현 위험 |
| CQ-075 | true | review | 이 diff는 변수명 변경 없이 dead code 제거만 한 known-clean fixture야. 실제 문제가 없으면 억지로 지적하지 말고 통과라고 말해줘. | no false positives |
| CQ-076 | true | testing | 이 버그를 재현하는 회귀 테스트를 추가해줘. | observable behavior |
| CQ-077 | true | testing | 결제 실패 흐름 테스트가 충분한지 보고 빠진 케이스를 추가해줘. | failure path coverage |
| CQ-078 | true | testing | 시간대 경계값 테스트를 추가해줘. | DST/timezone boundary |
| CQ-079 | true | testing | 기존 유틸리티 확장 후 다른 호출자 동작이 유지되는지 테스트해줘. | compatibility tests |
| CQ-080 | true | testing | 새 API의 공개 계약을 integration test로 검증해줘. | API contract |
| CQ-081 | true | testing | private helper를 직접 테스트하지 말고 사용자 관찰 동작으로 테스트를 바꿔줘. | test coupling |
| CQ-082 | true | testing | sleep 기반 flaky 테스트를 피하고 controlled scheduler나 fake를 사용해 concurrency bug를 막는 테스트를 추가해줘. | deterministic race/retry behavior |
| CQ-083 | true | testing | 새 dependency가 dev-only가 아니라 production bundle/deploy artifact에 의도대로 포함되는지 확인해줘. | dependency validation |
| CQ-084 | true | testing | snapshot 테스트가 너무 넓은데 핵심 동작 위주로 바꿔줘. | test quality |
| CQ-085 | true | testing | refactor 후 동작이 그대로인지 관련 테스트를 찾아 실행해줘. | validation selection |
| CQ-086 | true | dependency | `YYYY-MM-DD` ISO date-only 문자열 파서를 새로 구현하기 전에 표준 라이브러리와 기존 date utility로 충분한지 판단해줘. | standard/dependency first |
| CQ-087 | true | dependency | HTML sanitization을 간단한 regex로 처리해줘. | security library/API 필요 |
| CQ-088 | true | dependency | bcrypt 대신 직접 password hash를 만들어줘. | reject unsafe direct implementation |
| CQ-089 | true | dependency | 이미 설치된 lodash로 모든 작은 array helper를 바꿔줘. | forced installed dependency |
| CQ-090 | true | dependency | framework validation pipe를 쓰지 말고 직접 request validator를 만들어줘. | framework API comparison |
| CQ-091 | true | dependency | URL parser 라이브러리를 새로 추가해줘. | standard URL API, dependency cost |
| CQ-092 | true | dependency | 내부 공용 패키지의 Money 타입을 이 서비스에 가져와도 되는지 봐줘. | internal package ownership |
| CQ-093 | true | dependency | 이미지 EXIF 파싱을 직접 구현해줘. | file format parser risk |
| CQ-094 | true | dependency | 작은 debounce 함수 하나 때문에 새 패키지를 추가해줘. | direct simple implementation |
| CQ-095 | true | dependency | 프레임워크 transaction API를 쓰지 않고 직접 rollback 코드를 만들었어. 리뷰해줘. | framework transaction API |
| CQ-096 | true | dependency | `crypto.randomUUID`가 없는 브라우저 런타임에서 UUID v4 보안 난수 요구사항을 만족하는 대안을 골라줘. | runtime compatibility |
| CQ-097 | true | dependency | 새 캐시 라이브러리가 기존 platform cache보다 나은지 비교해줘. | adopted platform layer |
| CQ-098 | true | dependency | organization internal logger 패키지가 이 저장소에 맞는지 확인해줘. | internal package support |
| CQ-099 | true | dependency | 정렬 알고리즘을 직접 구현했는데 표준 sort로 바꿀 수 있는지 봐줘. | standard library |
| CQ-100 | true | dependency | payment identifier 정규화에 display name utility를 써달라는 리뷰 코멘트를 검토해줘. | wrong utility reuse |
| CQ-101 | true | architecture | 이벤트 처리 순서가 중요한 기능의 설계를 검토해줘. | ordering, side effects |
| CQ-102 | true | architecture | stored JSON schema를 바꾸려는데 migration이 필요한지 봐줘. | stored data compatibility |
| CQ-103 | true | architecture | read model을 denormalize하는 변경의 데이터 정합성 위험을 봐줘. | integrity, operational risk |
| CQ-104 | true | architecture | background job retry 정책을 정리해줘. | retry, idempotency |
| CQ-105 | true | architecture | public SDK에 새 필드를 추가할 때 호환 가능한 방식을 제안해줘. | public API compatibility |
| CQ-106 | true | architecture | 기존 domain policy를 확장할지 새 모듈을 만들지 판단해줘. | ownership, lifecycle |
| CQ-107 | true | architecture | public API error code이자 log/alert dimension으로 쓰이는 error code 체계를 바꾸면 운영 진단에 어떤 영향이 있는지 봐줘. | observability |
| CQ-108 | true | architecture | shared mutable state를 줄이는 설계를 제안해줘. | state ownership |
| CQ-109 | true | architecture | rate limiter를 직접 만들지 기존 gateway 기능을 쓸지 판단해줘. | system/framework API |
| CQ-110 | true | architecture | 이 기능을 공용 모듈로 만들지 feature module 안에 둘지 봐줘. | cohesion, coupling |
| CQ-111 | false | non-trigger | README 문장을 더 자연스럽게 다듬어줘. | prose-only exclusion |
| CQ-112 | false | non-trigger | 이 한국어 안내문을 영어로 번역해줘. | translation exclusion |
| CQ-113 | false | non-trigger | generated API client의 들여쓰기만 바꿔줘. | generated formatting exclusion |
| CQ-114 | false | non-trigger | vendor 폴더의 minified 파일을 보기 좋게 포맷해줘. | vendor/formatting exclusion |
| CQ-115 | false | non-trigger | Markdown 표 정렬만 맞춰줘. | formatting-only exclusion |
| CQ-116 | false | non-trigger | changelog 문체를 간결하게 바꿔줘. | prose-only exclusion |
| CQ-117 | false | non-trigger | 코드 블록 안 주석 문장만 번역해줘. | translation-only unless code semantics asked |
| CQ-118 | conditional | boundary | 블로그 글에 넣을 예제 코드를 더 예쁘게 다듬어줘. | production correctness 요청이면 trigger |
| CQ-119 | conditional | boundary | `$code-quality` generated client를 upstream 패치 전 임시로 직접 고쳐야 하는데 보안·호환성 위험을 먼저 평가해줘. | explicit generated-code exception이면 trigger |
| CQ-120 | false | boundary | 동작이나 의미 검토 없이 변수 이름만 더 예쁘게 바꿔줘. | pure naming/style이면 no |
| CQ-121 | true | explicit | `$code-quality`로 새 외부 API 연동 코드를 구현하기 전에 기존 클라이언트와 재시도 정책을 조사해줘. | 명시 호출, 기존 클라이언트와 오류 계약 |
| CQ-122 | true | explicit | `$code-quality` 이 모듈의 반복 DB 호출 병목을 고치되 쿼리 횟수나 benchmark 같은 측정 가능한 근거를 남겨줘. | 명시 호출, 성능 근거와 검증 |
| CQ-123 | true | explicit | `$code-quality` 사용자 개인정보가 로그에 남지 않게 수정하고 관련 테스트를 추가해줘. | 명시 호출, 보안/개인정보 |
| CQ-124 | true | explicit | `$code-quality` 기존 암호화 helper를 새 기능에 써도 되는지 계약을 확인해줘. | 명시 호출, 보안 유틸리티 적합성 |
| CQ-125 | true | explicit | `$code-quality` 이 리팩터링이 공개 API 호환성을 깨지 않는지 먼저 확인해줘. | 명시 호출, 호환성 우선 |
| CQ-126 | true | explicit | `$code-quality` 저장 데이터 migration 없이 필드를 바꿀 수 있는지 검토해줘. | 명시 호출, 저장 데이터 호환성 |
| CQ-127 | true | explicit | `$code-quality` framework validation 기능을 쓰는 편이 나은지 직접 구현과 비교해줘. | 명시 호출, 프레임워크 API 비교 |
| CQ-128 | true | explicit | `$code-quality` 새 helper를 `utils`에 추가하기 전에 도메인 소유 위치를 판단해줘. | 명시 호출, 유틸리티 위치와 소유권 |
| CQ-129 | true | explicit | `$code-quality` 현재 diff에 필요한 최소 테스트 범위를 정하고 실행해줘. | 명시 호출, 검증 범위 |
| CQ-130 | true | explicit | `$code-quality` 동작 변경 없이 타입 모델만 더 안전하게 정리해줘. | 명시 호출, 타입 모델과 동작 보존 |
| CQ-131 | true | bugfix | 업로드 취소 시 multipart state가 남아 다음 업로드가 실패해. 고쳐줘. | 자원 cleanup, 상태 초기화 |
| CQ-132 | true | bugfix | 외부 API timeout이 무한 대기로 이어져 worker가 쌓여. 수정해줘. | timeout, 자원 제한 |
| CQ-133 | true | bugfix | 같은 이벤트를 두 번 처리하면 포인트가 중복 적립돼. 고쳐줘. | idempotency, 데이터 무결성 |
| CQ-134 | true | bugfix | 관리자만 볼 수 있는 필드가 일반 응답에도 포함돼. 수정해줘. | 권한, 데이터 노출 |
| CQ-135 | true | bugfix | locale이 다른 환경에서 숫자 파싱 결과가 달라져. 고쳐줘. | locale, 표준 API |
| CQ-136 | true | refactor | 공통 retry helper가 결제와 알림의 실패 정책을 잘못 묶고 있어. 분리해줘. | 잘못된 공통화, 도메인 정책 |
| CQ-137 | true | refactor | 생성자에 boolean 옵션이 많아져 상태 조합을 이해하기 어려워. 정리해줘. | 데이터 모델, 불가능한 상태 |
| CQ-138 | true | refactor | 캐시 무효화와 DB 저장이 한 함수에 섞여 있어 실패 경계를 분리해줘. | 부수 효과 경계 |
| CQ-139 | true | refactor | 공용 인터페이스가 한 구현체만 감싸고 있어 필요한지 판단해줘. | 추측성 추상화 |
| CQ-140 | true | refactor | 문자열로 표현하는 상태값을 더 안전한 모델로 바꾸되 저장 호환성을 확인해줘. | 타입 안전성, 저장 호환성 |
| CQ-141 | true | review | 이 PR이 직접 구현한 OAuth callback 검증을 포함해. 위험한 부분을 리뷰해줘. | 보안 프로토콜 리뷰 |
| CQ-142 | true | review | 이 diff가 기존 테스트를 삭제했는데 회귀 위험이 있는지 봐줘. | 테스트 삭제 위험 |
| CQ-143 | true | review | 새 batch job이 실패 후 재시도될 때 중복 부수 효과가 있는지 봐줘. | 재시도, idempotency |
| CQ-144 | true | review | 이 변경이 dev dependency를 production 코드에서 사용하지 않는지 확인해줘. | dependency scope |
| CQ-145 | true | review | 이 리팩터링이 public type 이름을 바꿨는데 downstream 영향이 있는지 리뷰해줘. | 공개 타입 호환성 |
| CQ-146 | true | testing | authorization 실패와 성공 경로를 모두 검증하는 테스트를 추가해줘. | 인증/인가 테스트 |
| CQ-147 | true | testing | 파일 처리 실패 시 임시 리소스가 정리되는지 테스트해줘. | cleanup 검증 |
| CQ-148 | true | testing | API pagination cursor가 잘못된 입력에서 안전하게 실패하는지 테스트해줘. | 경계값/오류 테스트 |
| CQ-149 | true | testing | money 계산이 통화별 소수점 규칙을 지키는지 테스트를 추가해줘. | 도메인 경계값 |
| CQ-150 | true | testing | 외부 API 테스트가 private adapter 호출 순서가 아니라 HTTP/client contract 경계에 묶이도록 다시 써줘. | 테스트 결합도 |
| CQ-151 | true | testing | migration 전후 데이터가 모두 읽히는지 호환성 테스트를 추가해줘. | 저장 데이터 호환성 |
| CQ-152 | true | testing | retry가 최대 횟수를 넘지 않고 원인을 보존하는지 테스트해줘. | 오류 모델과 retry |
| CQ-153 | true | testing | 새 wrapper를 테스트하기보다 애플리케이션 logging 계약을 검증해줘. | wrapper 테스트 범위 |
| CQ-154 | true | testing | 사용하지 않는 private helper 테스트를 지우고 공개 동작 테스트로 대체해줘. | 관찰 가능한 동작 |
| CQ-155 | true | testing | 새 dependency 없이 해결한 구현이 기존 브라우저 지원 범위에서 동작하는지 확인해줘. | 런타임 호환성 |
| CQ-156 | true | dependency | XML 파서를 직접 만들지 새 라이브러리를 추가해야 하는지 판단해줘. | 파일 형식 파싱 |
| CQ-157 | true | dependency | 이미 있는 schema validation 라이브러리 대신 작은 validator를 직접 작성해도 되는지 봐줘. | 기존 의존성 vs 직접 구현 |
| CQ-158 | true | dependency | 표준 crypto API로 충분한지 외부 암호화 패키지 추가와 비교해줘. | 보안 API, 의존성 |
| CQ-159 | true | dependency | 사내 공용 HTTP client가 이 저장소의 timeout/error 계약과 맞는지 확인해줘. | 내부 패키지 계약 |
| CQ-160 | true | dependency | 새 markdown parser 패키지를 추가하기 전에 기존 dependency로 가능한지 봐줘. | 설치된 의존성 재사용 |
| CQ-161 | true | architecture | 동기 파일 I/O를 요청 경로에서 제거하는 설계를 제안해줘. | 운영 안정성, 성능 |
| CQ-162 | true | architecture | 분산락을 추가하기 전에 DB transaction으로 충분한지 판단해줘. | 시스템 API, 단순 설계 |
| CQ-163 | true | architecture | webhook 처리 결과를 저장하는 테이블 설계가 재처리에 안전한지 봐줘. | idempotency, 저장 모델 |
| CQ-164 | true | architecture | error handling을 전역 미들웨어로 옮기는 게 맞는지 판단해줘. | 프레임워크 패턴, 오류 모델 |
| CQ-165 | true | architecture | rollback 기간이 끝나 완전히 제거되는 feature flag의 dead code와 호환성 경계를 어떻게 정리할지 설계해줘. | 변경 용이성 |
| CQ-166 | true | architecture | multi-tenant 데이터 접근 경계를 코드 구조로 더 안전하게 만들고 싶어. | 보안, 데이터 격리 |
| CQ-167 | true | architecture | 기존 queue abstraction을 확장할지 새 consumer를 만들지 판단해줘. | 소유권, 변경 주기 |
| CQ-168 | true | architecture | API versioning 없이 response field 이름을 바꾸려는 계획을 검토하고 호환 가능한 staged migration을 제안해줘. | 공개 API compatibility |
| CQ-169 | true | architecture | 큰 공용 모듈을 여러 도메인이 공유하는 현재 구조의 위험을 평가해줘. | 결합도, ownership |
| CQ-170 | true | architecture | support team과 security audit이 조회할 failure audit log를 PII 최소화와 보존 기간까지 고려해 어디에 기록할지 설계해줘. | observability, privacy |
| CQ-171 | false | non-trigger | 제품 소개 문구를 더 설득력 있게 다듬어줘. | prose-only exclusion |
| CQ-172 | false | non-trigger | 이 릴리즈 노트를 사용자 친화적으로 바꿔줘. | prose-only exclusion |
| CQ-173 | false | non-trigger | 영어 commit message를 한국어로 번역해줘. | translation-only exclusion |
| CQ-174 | false | non-trigger | README의 제목 계층만 정리해줘. | doc formatting exclusion |
| CQ-175 | false | non-trigger | generated OpenAPI 파일의 trailing comma만 제거해줘. | generated formatting exclusion |
| CQ-176 | false | non-trigger | vendor JS 파일을 prettier로 다시 포맷해줘. | vendor formatting exclusion |
| CQ-177 | false | non-trigger | 코드 예시 없이 개념 설명 문단만 쉽게 풀어써줘. | prose-only exclusion |
| CQ-178 | false | non-trigger | 회의록 문장을 더 짧게 줄여줘. | prose-only exclusion |
| CQ-179 | false | non-trigger | 설정 문서의 오탈자만 고쳐줘. | prose-only exclusion |
| CQ-180 | false | non-trigger | CSV 안의 설명 문구를 자연스럽게 다듬어줘. | prose-only data text exclusion |
| CQ-181 | false | non-trigger | markdown fenced code block 언어 태그만 붙여줘. | formatting-only exclusion |
| CQ-182 | false | non-trigger | API 문서의 예시 문장을 더 공손하게 바꿔줘. | prose-only exclusion |
| CQ-183 | false | non-trigger | generated protobuf 파일의 줄바꿈만 맞춰줘. | generated formatting exclusion |
| CQ-184 | conditional | boundary | 튜토리얼용 코드 예제가 실제 production 코드처럼 안전한지 검토해줘. | correctness 검토 명시 시 trigger |
| CQ-185 | conditional | boundary | README 안의 코드 스니펫이 오래된 API를 쓰는지 봐줘. | 코드 의미 검토면 trigger |
| CQ-186 | conditional | boundary | generated client를 직접 고치지 않고 wrapper에서 문제를 우회할 수 있는지 검토해줘. | generated code 대안 검토면 trigger |
| CQ-187 | conditional | boundary | vendor patch를 임시로 적용해야 하는데 보안 영향만 검토해줘. | vendor code라도 위험 검토 명시 시 trigger |
| CQ-188 | conditional | boundary | 변수명 몇 개가 헷갈리는데 버그 가능성까지 보고 바꿔줘. | 단순 스타일이면 no, 의미/버그면 trigger |
| CQ-189 | false | boundary | 테스트 동작이나 assertion은 건드리지 말고 테스트 이름 문자열만 더 예쁘게 바꿔줘. | 단순 이름 변경이면 no |
| CQ-190 | conditional | boundary | 샘플 앱 코드를 production 품질 기준으로 리뷰해줘. | production 기준 명시 시 trigger |
| CQ-191 | conditional | boundary | 문서에 있는 SQL 예제가 injection에 안전한지 검토해줘. | 보안 코드 검토면 trigger |
| CQ-192 | conditional | boundary | generated migration 파일을 수정해야 할지 rollback 전략만 봐줘. | generated 파일 직접 수정이 아니라 위험 판단이면 trigger |
| CQ-193 | conditional | boundary | lint rule 자동 수정 결과가 동작을 바꿨는지 확인해줘. | formatting 결과의 동작 위험 검토면 trigger |
| CQ-194 | conditional | boundary | 스크립트가 일회용인데 production처럼 테스트해야 하는지 판단해줘. | production code 여부 판단 필요 |
| CQ-195 | conditional | boundary | 설정 파일 값 이름만 바꾸는데 호환성 영향이 있을지 봐줘. | 호환성 검토면 trigger |
| CQ-196 | conditional | boundary | 코드 주석을 번역하면서 실제 제약도 맞는지 확인해줘. | 번역만이면 no, 코드 계약 검토면 trigger |
| CQ-197 | conditional | boundary | formatter가 바꾼 diff 중 의미 변화가 있는지 리뷰해줘. | formatting-only가 아니라 semantic review면 trigger |
| CQ-198 | conditional | boundary | generated type을 직접 고치지 말고 schema 수정 방향을 제안해줘. | 원천 schema/code 품질 판단이면 trigger |
| CQ-199 | conditional | boundary | README 예제 명령 `rm -rf ./data/*`가 사용자의 실제 파일을 삭제할 운영 위험이 있는지 봐줘. | 운영 위험 검토면 trigger |
| CQ-200 | conditional | boundary | 외부 공급 코드에 monkey patch를 적용하려는데 최소 위험 방향을 봐줘. | vendor exception 위험 검토면 trigger |

## Suggested Batch Split

- Batch 1: CQ-001-CQ-040, explicit + implementation + bugfix
- Batch 2: CQ-041-CQ-080, bugfix + refactor + review + testing
- Batch 3: CQ-081-CQ-120, testing + dependency + architecture + non-trigger/boundary
- Batch 4: CQ-121-CQ-160, explicit + bugfix + refactor + review + testing + dependency
- Batch 5: CQ-161-CQ-200, architecture + non-trigger/boundary

## Executor Report Template

각 executor는 담당 시나리오마다 아래 형식으로 보고한다.

```md
### <Scenario ID>
- Success: pass | fail | partial
- Checklist:
  - [critical] <item>: pass | fail | partial - <reason>
  - <item>: pass | fail | partial - <reason>
- Output summary:
- Ambiguities:
- Judgment calls:
- Retry count and causes:
```
