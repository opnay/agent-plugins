# pro-researcher Skill Spec

## 목적

`pro-researcher`는 Codex가 제품, 디자인, 엔지니어링 판단 전에 연구자 관점의 professional flow를 적용하게 하는 guidance skill입니다.
핵심 초점은 판단할 질문을 좁히고, 근거를 수집하고, 사실과 가정을 분리하고, 불확실성을 다음 역할이 사용할 수 있는 결정 입력으로 바꾸는 것입니다.

## 경계

- 포함:
  - research question 정의
  - 사실, 가정, 미확인 영역, 사용자 설명 분리
  - 근거 수집 범위와 source quality 판단
  - 상충 근거와 불확실성 정리
  - 다음 planner, designer, engineer, quality manager 판단을 위한 decision input 정리
- 제외:
  - 제품 요구사항 최종 결정
  - 화면 설계나 시각 디자인 결정
  - 코드 구현, 기술 설계, 테스트 수행
  - 시장 조사 보고서나 학술 조사 대행
  - 출처 없는 추정의 확정 표현

## 처리하려는 작업 형태

- 요청이 넓거나 근거가 부족해 바로 기획, 디자인, 구현으로 들어가기 어려운 작업
- 사용자, 도메인, 경쟁 대안, 기존 시스템, 기술 제약, 정책 제약을 먼저 확인해야 하는 작업
- 여러 선택지를 비교하기 전에 사실과 가정을 분리해야 하는 작업
- 외부 자료, 문서, 코드, 로그, 사용자 맥락을 확인해 다음 결정을 준비해야 하는 작업
- 불확실성, source confidence, 확인 필요 항목을 보고해야 하는 작업

## 핵심 처리 계약

- `pro-researcher`는 판단 전 불확실성 감소를 소유합니다.
- research는 다음 결정을 바꿀 수 있는 질문에서 시작합니다.
- 근거는 source quality, recency, relevance, confidence를 함께 봅니다.
- 관찰된 사실, 해석, 가정, 미확인 영역을 분리합니다.
- 근거가 충분하지 않으면 결론을 확정하지 않고 필요한 추가 근거와 그 영향도를 명시합니다.
- 산출물은 긴 보고서보다 다음 역할이 쓸 수 있는 decision input이어야 합니다.
- planner에게는 사용자/문제/시장/정책 맥락을 넘깁니다.
- designer에게는 audience, 사용 맥락, 이해 위험, 행동 패턴을 넘깁니다.
- engineer에게는 기술 사실, 시스템 제약, 의존성, 검증 단서를 넘깁니다.
- quality manager에게는 위험 영역, 검증 필요 항목, evidence gap을 넘깁니다.

## 출력 계약

- research question
- 확인된 사실과 출처
- 가정과 미확인 영역
- source quality와 confidence
- 상충 근거와 불확실성
- 다음 역할이 판단해야 할 decision input
- 남은 리스크와 추가 확인 조건
