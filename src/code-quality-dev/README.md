# Code Quality Dev

`code-quality-dev`는 생산 코드 작성, 수정, 리팩터링, 테스트, 의존성 판단, 코드 리뷰를 위한 플러그인입니다.

이 플러그인의 핵심 표면은 `code-quality` skill입니다.
이 skill은 코드 미관보다 정확성, 저장소 맥락, 안전한 변경, 검증, 적절한 재사용, 단순한 설계를 우선합니다.
기존 도메인 모듈, 프레임워크와 시스템 API, 표준 라이브러리, 설치된 의존성, 기존 유틸리티를 먼저 조사하되 재사용을 강제하지 않고 의미와 계약이 맞는지 판단합니다.

대표 호출:

- `$code-quality-dev:code-quality`로 구현, 버그 수정, 리팩터링, 테스트 추가, 의존성 검토, 코드 리뷰를 수행합니다.
- 스킬 이름을 직접 쓰지 않아도 production code 구현, 수정, 리뷰 요청에서는 암시적으로 사용할 수 있습니다.

비목표:

- 문서 전용 작성
- 번역
- 단순 포맷 변경
- 단순 이름 변경이나 스타일 변경
- 생성 코드나 외부 vendor 코드 수정

위 작업은 사용자가 `$code-quality-dev:code-quality`를 명시적으로 호출했거나, correctness/security/compatibility/data-integrity/operational-risk 같은 code-quality 검토를 직접 요청한 경우에만 다룹니다.
문서 안의 코드 스니펫, 명령, SQL, 설정 예시는 문장 다듬기가 아니라 실제 동작·보안·호환성·운영 위험 검토일 때만 범위에 포함합니다.
