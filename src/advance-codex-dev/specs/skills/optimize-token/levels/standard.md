# `standard` 단계 게이트

`standard`는 `light`를 상속하고, 긴 답변에서 핵심 판단 표면만 남기도록 덮어씁니다.

## 적용

- "핵심만", "요약", "더 짧게" 요청
- 긴 원문에서 결과, 근거, 다음 행동, 중요한 위험만 필요함
- `extreme` 요청 중 안전, 검증, 승인, 명확성 때문에 일부를 더 써야 함

## 덮어쓰기

- 우선 유지: 핵심 판단, 신뢰도 근거, 다음 행동/막힘 요소/사용자 결정, 중요한 위험과 남은 불확실성
- 의미 맥락 유지: workflow 순서, 의존성, 원인-증거 관계, 비교 기준, 제외 범위, 확인/미확인 범위, source of truth, evidence 출처 계층은 판단을 바꾸면 줄이지 않음
- 추가 형태: 단계나 흐름 순서는 `` `light` > `standard` > `extreme` `` 체인을 쓰고, 관련 list 항목은 의미 분류가 흐려지지 않을 때만 상위 항목과 sublist로 묶어 반복 표현을 줄임
- 추가 생략: 부차 배경, 결정에 영향 없는 예시, 의존성이 아닌 시간순 과정, 낮은 가치의 안심 표현
- 이동: 안전한 결정 설명이 필요하면 `light`, 라벨/상태/한 줄/압축 표가 명시되면 `extreme`

## 예시

- `세 파일을 수정했고 핵심 테스트는 통과했지만 외부 연동은 미실행입니다.` > `핵심 테스트 통과. 외부 연동 미실행.`
- `light, standard, extreme 순서로 읽습니다.` > `` `light` > `standard` > `extreme` 순서로 읽습니다. ``
- `생성물은 release surface이고 원본은 dev source입니다.` > `source: dev source. generated: release surface.`
- `A를 바꿨고 B는 이번 범위가 아닙니다.` > `변경: A. 제외: B.`
- list group:
  ```md
  - login failed
  - login timeout
  ```
  >
  ```md
  - login
    - failed
    - timeout
  ```
