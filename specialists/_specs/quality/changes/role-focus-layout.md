# 직업·초점 레이아웃 변경

## 변경사항 요약

- agent/quality → specialists/quality. verify 실행 계약과 version은 유지합니다.
- Plugin spec·사용/개발 안내·두 host manifest·marketplace와 현재 참조를 맞춥니다.

## 변경 상세

직업·분야·초점·숙련도 소유권은 저장소 `docs/agent-composition.md`를 따릅니다. 실제 plugin은 `specialists/quality`, 소유 spec은 `specialists/_specs/quality`입니다. 과거 changes/intent/실험 기록의 이전 이름과 경로는 당시 근거로 보존합니다.

## 비목표

orcat/app 연결 구현, 사용자 plugin 재설치, commit·push·main 머지, 신규 분야·security skill·품질관리 전체의 기능 확장은 포함하지 않습니다.

## 호환성과 migration

새 plugin 이름·호출을 사용하며 이전 식별자 alias는 제공하지 않습니다. 새 이름의 plugin은 0.1.0부터 시작합니다. 이름이 유지된 plugin은 기존 version을 유지합니다. 사용자의 이전 cache/설정은 별도 갱신 전까지 유지됩니다. 숙련도 assets는 역할 셋업에서 명시적으로 적용하며 기존 skill 호출의 필수 선행 조건이 아닙니다.

## 검증 기준

실제 plugin/skill/spec·두 manifest·marketplace의 경로/순서/metadata·현재 문서/리소스 참조를 확인합니다. 이동 계약의 보존과 통합 연구 계약을 구분하며 독립 read-only 정합성 검토와 격리 native 설치·발견·본문 읽기·비교 실행을 수행합니다. 결과/한계는 별도 검증 기록으로 보존합니다.
