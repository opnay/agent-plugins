# Toolkit 005: plain Git runtime

## 변경사항 요약

- Git runtime을 정상 실행 지침과 조건부 reference routing으로 압축합니다.

## 변경 상세

- 목적: 설치된 skill이 설명보다 명령·조건·다음 행동을 먼저 전달합니다.
- 범위: `SKILL.md`는 normal commit, branch, push, alias routing만 유지합니다. message-file lifecycle 세부는 runtime reference가 소유합니다.
- 비목표: Git 권한, message validation, branch·push recovery 계약을 완화하거나 제거하지 않습니다.
- 관련 표면: git skill spec, runtime `SKILL.md`, `references/message-lifecycle.md`.

## 검증 기준

- normal workflow가 `SKILL.md`만으로 실행 가능하고 조건부 detail이 discoverable합니다.
- spec과 runtime의 scope·reference routing이 일치합니다.
- runtime frontmatter와 bundled reference links를 검증합니다.
- `git diff --check`를 통과합니다.
