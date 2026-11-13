# Judgment Kit

`judgment-kit`는 Codex 작업에 분야별 professional judgment flow를 적용하기 위한 플러그인입니다.
리서처, 제품 기획자 관점의 판단을 각각 독립 skill로 제공합니다.

이 플러그인은 특정 제품 구현 workflow나 커밋 절차를 소유하지 않습니다.
작업을 시작하기 전 문제, 범위, 설계 기준, 검증 기준을 더 정확히 잡는 데 목적이 있습니다.

`pro-researcher`는 제품, 디자인, 엔지니어링, 품질 관리 판단 전에 research question, 근거, 사실/가정, source quality, 불확실성, decision input을 정리해야 할 때 사용합니다.
`pro-planner`는 제품, 서비스, 기능 요청에서 넓은 요청을 사용자 문제, 제품 방식 후보, 기능 영역, 디자인 시스템 브리프 같은 부가 기획 표면, MVP 범위, 요구사항, 우선순위, acceptance criteria, tradeoff, designer/engineer handoff로 분해해야 할 때 사용합니다.

산출물의 검증 계획·실행·근거 판정은 `$quality:verify`를 사용합니다. 별도 플러그인이며 이 플러그인의 필수 의존성은 아닙니다.
