# Toolkit 004: push input trust

## 변경사항 요약

- 사용자가 지정한 push 형식과 위치를 push 전 다시 탐색하지 않습니다.

## 변경 상세

- 목적: 사용자가 이미 아는 작업 위치를 remote URL·upstream·source·destination 조회로 중복 확인하지 않습니다.
- 범위: current-branch push 형식과 exact `<local>:<remote>` refspec push를 바로 실행합니다.
- 보존: 잘못된 호출, nonzero·중단, 상충·불완전 출력, 결과 불명확에는 recovery에 필요한 상태만 조회합니다.
- 관련 표면: git skill spec/runtime, Toolkit plugin spec, runtime·development README.

## 검증 기준

- spec, runtime, README가 같은 push input trust 계약을 설명합니다.
- runtime frontmatter와 bundled reference link를 검증합니다.
- `git diff --check`를 통과합니다.
