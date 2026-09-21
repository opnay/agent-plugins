# Toolkit 006: focused Git references

## 변경사항 요약

- branch convention과 recovery reference를 조건별 최소 지침으로 정리합니다.

## 변경 상세

- 목적: conditional reference가 정상 workflow의 설명이나 preflight를 다시 늘리지 않습니다.
- 범위: branch policy는 이름·충돌 판단만, recovery는 실패 유형별 필요한 상태와 다음 행동만 소유합니다.
- 비목표: branch collision, partial success, push recovery의 안전 계약을 완화하지 않습니다.
- 관련 표면: git skill spec, `branch-conventions.md`, `recovery.md`.

## 검증 기준

- normal push의 remote·upstream·ref mapping preflight가 reference에 없습니다.
- branch와 recovery의 필요한 상태 조회·권한 경계가 유지됩니다.
- runtime reference links를 검증합니다.
- `git diff --check`를 통과합니다.
