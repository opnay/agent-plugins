# 008 Jev 기본 사용 패턴 통합

## 변경사항 요약

- `jev-use`의 기본 패턴과 활용 예시를 `$jev:jev`에 통합합니다.
- 모델 선택은 내장 유즈케이스이며 위임·모델 정책을 소유하지 않습니다.
- 공개 스킬은 `jev`, `scenario-testing` 두 개입니다.

## 변경 상세

### 목적·범위

기본 사용 패턴이 별도 오케스트레이션 스킬로 읽히는 문제를 해결합니다. `jev`는 공통 계약을, `references/use-cases.md`는 적용 예시를 제공합니다. 기존 `jev-use` runtime·skill spec을 제거하고 plugin spec·README·manifest의 시작 기준을 맞춥니다.

### 비목표

CLI·API·설정·설치 동작, `scenario-testing` 평가 계약, 다른 오케스트레이션 스킬과 모델 정책은 변경하지 않습니다.

### 호환성·마이그레이션

`$jev:jev-use` 호출은 `$jev:jev`로 바꿉니다. CLI 명령과 `$jev:scenario-testing` 호출은 유지합니다. 이전 변경 기록은 당시 이력으로 보존합니다. 설치된 캐시는 이 저장소 수정만으로 갱신되지 않습니다.

### 검증 기준

- frontmatter·JSON·링크·marketplace 경로와 현재 사용 표면의 식별자를 확인합니다.
- 기본 패턴과 예시, 프로필 또는 개별 모델·effort 선택, 근거 판단·실패 처리의 spec/runtime 정합성을 독립 검증합니다.
- CLI 안내와 `scenario-testing` 계약을 보존하고 `git diff --check`를 통과합니다.

## 정식 규칙 승격

현재 계약은 plugin·jev skill spec과 runtime이 소유하며, 통합 이력과 호출 마이그레이션은 이 문서에 남깁니다.
