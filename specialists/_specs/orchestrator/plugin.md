## 사용자 스펙 의도

- 분기 설계부터 fork·검증·완료·flush까지를 공통 실행 루프로 둡니다.
- 메인은 유지하며 다음 회차는 보존 자료를 받은 새 맥락에서 계획합니다.

# Orchestrator 플러그인 스펙

## 목적과 경계

`orchestrator`는 직업 간 공통 실행 조정 초점입니다. 작업 분해·분담/대안·의존성·실행 소유권·실제 결과 검증·통합·완료 판단·보존/새 맥락 계획을 소유합니다. 제품·디자인·엔지니어링·조사 고유 기준을 대신하지 않습니다.

## 내장 skill과 사용 기준

`$orchestrator:orchestrate-workstreams` 하나를 제공합니다. 상세 계약은 [skill spec](skills/orchestrate-workstreams/spec.md)이 소유합니다. 명시적 분업·서브에이전트 요청 또는 의미 있는 독립 작업/대안/새 맥락 반복에서 선택합니다. 단순 조회·일반 구현 표현·파일 수·복잡성만으로 선택하지 않습니다.

## 독립성과 조합

조사나 다른 sibling 없이 실행 가능하며 배정된 분야 계약은 명시적으로 연결합니다. 연구가 orchestration을 항상 요구하지 않습니다. 적용 중인 실행 지침·backend·권한이 worker 호출을 소유합니다. 이 저장소의 orcat 역할 실행은 외부 subagent이며 native 앱 내역 연결이나 새 실행 도구를 plugin이 제공한다고 주장하지 않습니다. 상세 실행 경계는 설치된 delegation-safety reference에서 읽습니다.

## 사용 표면과 확장

실제 plugin은 `specialists/orchestrator`, spec은 `specialists/_specs/orchestrator`입니다. 두 manifest/README/prompt는 단일 loop와 실제 references를 안내합니다. 모델·원격 실행·호스트의 설치 정책을 plugin 이름으로 추정하지 않습니다. 회차/위임/상태 계약은 기존 루프를 보존하며 분야별 루프를 복제하지 않습니다.
