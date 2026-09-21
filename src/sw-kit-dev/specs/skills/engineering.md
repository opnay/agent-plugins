## 사용자 스펙 의도

- 엔지니어링은 사용자의 요구사항에 맞는 기술 선택, 데이터 모델, 구현 방향을 잡는 데 집중해야 한다.
- 코드 파일 구성이나 로컬 리팩터링 판단은 별도 code skill이 맡아야 한다.

---

# Engineering Skill Spec

## 목적

`engineering`은 행동 계약을 지속 가능하게 구현하기 위한 기술 방향을 선택한다.

## 경계

- 포함: system boundary, data model, architecture, technology/framework/storage 선택, integration, migration, security/operations tradeoff.
- 제외: 제품 요구 우선순위, source file/module 구현, routine code review, dependency/deprecation lifecycle audit, Git workflow.

## 처리 계약

- 기능적·비기능적 제약, 기존 시스템, 운영 환경을 확인한다.
- 대안은 요구 적합성, 복잡도, 변경 비용, reliability, security, performance, 운영 가능성으로 비교한다.
- data ownership, consistency, lifecycle, failure handling, integration contract를 명시한다.
- 선택과 보류안을 근거, 영향, migration 또는 rollback 조건과 함께 기록한다.
- 미래 확장 가능성은 현재 요구, 확인된 후속 계획, 독립 소비자, 호환성·운영 제약 같은 근거가 있을 때만 선택 근거로 쓴다.
- 가능성만으로 계층, 공용 모듈, interface, 저장소, event flow, 설정 지점을 만들지 않는다. 현재 구조의 비용이나 위험이 확인됐거나 가까운 후속 요구가 같은 경계를 공유할 때 확장을 선택하고, 그렇지 않으면 확장 trigger를 남긴다.
- 확실하지 않은 사실은 가정으로 표시하고 검증 경로를 제시한다.

## 검토 질문

- 이 선택이 요구와 운영 제약을 만족하는가?
- 데이터의 소유·수명·일관성은 분명한가?
- 실패, migration, observability, rollback은 어떻게 다루는가?
- 현재 복잡도가 실제 위험과 변화에 비례하는가?
- 어떤 관찰 가능한 신호가 생기면 구조를 넓혀야 하는가?

## 독립성 원칙

상위 handoff가 없으면 기술 방향에 필요한 질문과 가정을 먼저 정리한다. code skill의 상세 구현을 요구하지 않는다.
