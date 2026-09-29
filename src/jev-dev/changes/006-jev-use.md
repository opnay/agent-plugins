# 006 Jev 작업 활용

## 변경사항 요약

- `$jev:jev-use`가 제한된 작업 판단, 불확실한 근거 평가, subagent 모델·reasoning effort 라우팅을 소유합니다.
- `$jev:jev`는 Jev 정보와 CLI 사용·관리 계약을 유지합니다.
- 플러그인 사용 안내에 세 스킬의 시작 기준을 구분합니다.

## 변경 상세

### 목적

간단한 사용자 질문과 작업 중 작은 선택에는 Jev를 보조 판단으로 쓰고, 범위가 미정인 작업은 기본 범위·조사량·실행량을 확인한 뒤 직접 수행 또는 허용된 위임 경로를 고릅니다.

### 포함 변경

- `jev-use` skill spec과 runtime을 추가합니다. 사용자 질문 필요성, 불확실한 근거의 넓고 좁은 질문, 범위 판단, 위임 여부, 사용 가능한 모델, 모델이 지원하는 effort를 평가합니다. 공통 state의 독립 질문 수는 임의의 소수로 고정하지 않고 현재 API 한도와 결정에 필요한 범위에 맞춥니다.
- `jev` skill spec과 runtime은 Jev·CLI 정보, 입력·출력·오류·설정·유지보수에 집중합니다.
- plugin spec, 두 README, manifest의 설명과 시작 예시를 세 스킬 체계에 맞춥니다.

### 비목표

- CLI·API 변경, 모델·effort의 영구 기본값 변경, 모든 질문의 자동 Jev 호출, 사용자 권한·취향 대리 결정, `scenario-testing`의 verdict 변경은 포함하지 않습니다.

### 호환성/마이그레이션

- 기존 `jev` 명령과 `$jev:jev`, `$jev:scenario-testing` 식별자는 유지합니다. 새 작업 판단은 `$jev:jev-use`가 소유합니다.
- 플러그인 설치만으로 CLI 실행 파일·API 인증이 준비되지는 않습니다.

### 검증 기준

- 새 skill frontmatter와 명령 예시가 유효하고, 세 스킬의 책임과 시작 기준이 spec·runtime·README·manifest에서 일치해야 합니다.
- 모델·effort는 당시 허용·지원 후보에서만 고르며, Jev 결과를 권한이나 품질 증거로 취급하지 않아야 합니다.
- 넓은 질문과 좁은 질문을 같은 batch에 넣을 때 state와 질문 독립성을 확인하고, 답의 충돌·근거 부재를 원문 조사로 처리해야 합니다.
- CLI 실패와 기준 미달은 실제 응답과 구분하고, 기본값·자체 판단 또는 필요한 사용자 질문으로 이어져야 합니다.
- JSON 파싱, marketplace 경로, 참조, `git diff --check`, clean-context spec/runtime 대조를 통과해야 합니다.
