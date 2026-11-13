# SW Kit migration

## 변경사항 요약

- `sw-kit`에 `spec`, `engineering`, `code`, `maintenance` skill을 제공한다.
- 기존 `$judgment-kit:pro-engineering`, `$judgment-kit:pro-code-keeper`, `$code-quality:code-quality` 공개 식별자를 제거한다.
- `judgment-kit`은 research, product planning, quality management 판단에 집중한다.

## 변경 상세

### 목적

소프트웨어 행동 계약, 기술 방향, source-level 품질, 코드베이스 수명 관리를 명확한 책임으로 나눈다.

### 포함 변경

- `spec`은 요구를 기술 중립 행동 계약으로 정리한다.
- `engineering`은 기술·데이터·운영 방향을 선택한다.
- `code`는 구현·리팩터링·테스트·리뷰를 수행한다.
- `maintenance`는 dependency, deprecation, drift, debt를 관리한다.

### 비목표

- 제품 기획, 디자인, Git workflow, 배포 실행을 이 플러그인에 넣지 않는다.

### 호환성 및 마이그레이션

이 변경은 breaking migration이다. 기존 호출은 각각 `$sw-kit:engineering`, `$sw-kit:maintenance`, `$sw-kit:code`로 바꿔야 한다. 이전 plugin/skill 복사본은 유지하지 않는다.

### 검증 기준

- manifest와 marketplace JSON이 유효하다.
- 네 runtime skill이 각 spec과 일치하고 독립적으로 읽힌다.
- 기존 public identifier가 live surface에 남지 않는다.
