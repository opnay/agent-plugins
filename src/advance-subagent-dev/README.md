# Advance Subagent Dev

서브에이전트를 활용한 근거 중심 조사와 독립 workstream 위임의 심화 실행 방법을 제공하는 instruction-only 플러그인입니다.

## 목적

`advance-subagent-dev`는 다음 두 skill을 하나의 설치 단위로 제공합니다.

- `deep-research`: 여러 출처를 수집·대조하고 근거·반대 근거·불확실성·인용을 추적하는 조사 보고서
- `orchestrate-workstreams`: 독립적인 software-engineering·조사·실행 흐름의 bounded 위임, 검증, 통합

현재 두 skill의 역할, trigger, routing, 산출물 계약은 각각 유지됩니다. 두 skill은 서로의 숨은 맥락이나 실행을 요구하지 않습니다.

## Skill 선택

### Deep Research

다음과 같이 조사 자체가 산출물일 때 사용합니다.

- 시장·경쟁·문헌·기술·정책·규제·기업 실사
- 원문 또는 1차 자료 확인과 독립 출처 교차 검증
- 반대 근거, 상충 수치, 최신성 위험, 불확실성을 포함한 인용 가능한 보고서

```text
$advance-subagent-dev:deep-research 한국의 2025년 생성형 AI 규제 변화를 조사해 기업 도입 판단 보고서로 정리해 주세요.
```

날씨·현재가 같은 단순 조회, 한 문서 요약, 근거 없는 아이디어 발산, 코드 구현, 외부 변경, 지속 모니터링에는 사용하지 않습니다.

### Orchestrate Workstreams

다음과 같이 의미 있는 작업 흐름이 둘 이상이며 독립적으로 시작할 수 있을 때 사용합니다.

- 별도 모듈·서비스·실행 경로의 software-engineering 조사 또는 구현
- 조사와 prototype·문서·코드·변환이 섞인 cross-domain 작업
- correctness·security·performance·quality 같은 독립 review lens
- 분리된 ownership과 공통 성공 기준을 가진 병렬 작업

```text
$advance-subagent-dev:orchestrate-workstreams 시장 조사를 수행하고 독립적인 프로토타입 두 개를 구현한 뒤 결과를 검증·통합해 주세요.
```

명시 호출도 dispatch gate를 우회하지 않습니다. 독립 workstream, 별도 계약, 병렬 이점, shared-state 안전, 메인 에이전트의 검증·통합 가능성 중 하나라도 부족하면 `DIRECT`로 수행합니다.

## 독립성과 모델 계약

- `deep-research`는 조사 범위, 출처 정책, 증거 원장, 반증, 보고서 계약을 독립적으로 소유합니다.
- `orchestrate-workstreams`는 `DIRECT`·`DISPATCH`, graph, task/result contract, shared-state, terminal result, whole-result verification을 독립적으로 소유합니다.
- 각 skill의 역할·모델 routing은 해당 skill 계약을 따릅니다. 한 skill의 route를 다른 skill에 자동 적용하지 않습니다.
- 두 skill 모두 별도 MCP, app, hook, executable runtime dependency를 요구하지 않습니다.

## 공개 호출

release build에서는 다음 식별자를 사용합니다.

- `$advance-subagent:deep-research`
- `$advance-subagent:orchestrate-workstreams`
