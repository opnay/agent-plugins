# 기존 품질 스킬 정리 연계

## 변경사항 요약

- 현재 Quality 사용 안내에서 `pro-quality-manager`와의 병행 선택 기준을 정리합니다.
- 검증 계획·실행·판정은 `$quality:verify`, 전문 역할·작업 관리·출시 승인은 각 소유자가 담당함을 명시합니다.

## 변경 상세

Judgment Kit 0.2.0에서 기존 품질 관리 skill을 제거하므로 Quality plugin spec·사용/개발 README의 현재 안내를 맞춥니다. verify runtime과 skill spec의 실행 계약은 유지합니다.

## 비목표

기존 품질 관리 계약의 일괄 이전, Quality 기능 확장, release 승인 제공, 사용자 설치 변경은 포함하지 않습니다.

## 호환성과 마이그레이션

Quality의 버전과 `$quality:verify` 호출은 유지합니다. 기존 호출의 migration과 역할 차이는 [Judgment Kit 0.2.0 변경 기록](../../../../agents/_specs/planner/changes/v0.2.0.md)이 소유합니다. Quality 0.1.0 검증 계획·결과의 기존 skill 보존 문구는 그 검증 당시의 범위입니다.

## 검증 기준

현재 안내와 실제 제공 skill을 대조하고 verify runtime·skill spec의 내용 보존, 문서 링크를 확인합니다.
