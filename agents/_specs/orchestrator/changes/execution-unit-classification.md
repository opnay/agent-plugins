# 실행 역할과 공통 기능 분류

## 변경사항 요약

- `orchestrator`의 실제 plugin을 `agents/orchestrator`, 소유 spec을 `agents/_specs/orchestrator`에 둡니다.
- 현재 안내와 두 marketplace 경로를 같은 분류에 맞춥니다.

## 변경 상세

독립 실행 주체는 agents, 실행 역할에 조합할 기능은 capabilities로 구분합니다. agents는 직업과 실행 조정 역할을 포함합니다. 연구·검증 skill이 충분한 입력으로 직접 실행할 수 있다는 기존 계약은 유지합니다.

## 비목표

플러그인/skill 이름·버전·실행 계약 변경, marketplace 재정렬, 사용자 설치·카탈로그 갱신, 메인 Codex 환경·orcat 코드·앱 내역 연동 변경은 포함하지 않습니다.

## 호환성과 migration

기존 호출과 설치 이름은 그대로입니다. 이동은 소스 경로와 작성 문서 분류를 바꾸며 설치된 cache를 재작성하지 않습니다. 과거 change/intent/검증 기록은 당시 이름·경로를 보존합니다.

## 검증 기준

세 plugin의 모든 runtime 파일 내용·권한을 이전 소스와 대조합니다. 두 marketplace의 이름·순서·경로와 host metadata, 현재 문서/리소스 링크, 소유 spec/intent를 확인합니다. 이 경로 변경만으로 새 모델 행동 비교나 실제 사용자 재설치를 필수로 하지 않습니다.
