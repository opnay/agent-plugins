# Supplemental Planning Surfaces

이 문서는 `pro-planner`의 부가 기획 표면 판단을 소유합니다.

## 목적

- 넓은 제품/기능 요청은 core feature만으로 충분하지 않습니다.
- 다음 역할이 곧바로 판단할 수 있도록 정보 구조, 화면 inventory, 디자인 시스템 브리프, 설정/운영/분석 표면을 함께 분리합니다.
- 이 문서는 디자인 실행이나 기술 설계를 소유하지 않고, handoff에 필요한 제품 결정값을 정리합니다.

## Design System Brief

- 디자인 시스템 브리프는 색, 폰트, 컴포넌트를 최종 결정하는 문서가 아닙니다.
- 제품 방식, 사용자 감정, 사용 빈도, 데이터 밀도, 신뢰/위험 수준, 상태 표현 요구를 기준으로 디자인 판단 입력을 만듭니다.
- 포함 항목:
  - `Tone`: 차분함, 경쾌함, 전문성, 신뢰감처럼 제품이 가져야 할 인상
  - `Density`: 대시보드형 고밀도, 입력 중심 저마찰, 초보자용 여유 있는 구성
  - `State language`: 정상, 빈 상태, 경고, 위험, 초과, 오류, 성공의 표현 원칙
  - `Component candidates`: 입력 폼, 거래 목록, 요약 카드, 그래프, 필터, 설정 패널처럼 필요한 UI 부품 후보
  - `Accessibility and responsiveness`: 모바일/데스크톱 우선순위, 대비, 터치 목표, 숫자 가독성

## Information Architecture

- 기능 영역 map을 사용자가 이해하는 navigation과 화면 inventory로 바꿉니다.
- 화면 inventory는 화면 배치가 아니라 필요한 화면, 목적, 핵심 정보, 상태를 정리합니다.
- navigation 후보는 사용자 빈도와 목표를 기준으로 대시보드, 기록, 통계, 설정 같은 그룹으로 나눕니다.

## Supporting Surfaces

- 설정, 온보딩, 알림, 데이터 관리, 도움말, 운영자/관리자 표면이 필요한지 판단합니다.
- core MVP에서 제외하더라도 다음 역할이 놓치지 않도록 후속 surface로 남깁니다.
- supporting surface는 기능 요구사항과 별도로 목적, trigger, 필요한 상태, 비목표를 기록합니다.

## Output Rule

- 넓은 요청의 planning output에는 필요한 경우 `Supplemental Planning Surfaces` 섹션을 둡니다.
- 이 섹션에는 design system brief, information architecture, screen inventory, supporting surfaces, deferred surfaces를 포함합니다.
- 구체 화면 배치, 시각 위계, 색상 팔레트, 컴포넌트 스타일은 designer handoff 이후의 판단으로 남깁니다.
