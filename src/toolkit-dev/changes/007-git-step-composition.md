# Toolkit 007: composable Git steps

## 변경사항 요약

- runtime `SKILL.md` 첫 문단을 spec 핵심 처리 계약 1·4에 맞춰 commit·branch·push의 결합 workflow로 되돌립니다.

## 변경 상세

- 목적: 단계를 분리된 요청과 추론 금지 목록으로 읽히게 하던 문구를 제거하고, 필요한 단계만 순서대로 조합하되 한 단계 권한을 다음 단계로 넓히지 않는 계약을 드러냅니다.
- 범위: 첫 문단만 허용 계약으로 재작성합니다. alias maintenance는 기존 명시 요청 섹션, recovery는 실패 대응 reference가 계속 소유합니다.
- 비목표: force push, `reset --hard`, history rewrite, GitHub·release 작업의 제외 경계와 branch·recovery reference의 안전 계약을 완화하지 않습니다. skill spec은 변경하지 않습니다.
- 관련 표면: `toolkit/skills/git/SKILL.md`.

## 검증 기준

- 첫 문단이 spec 핵심 처리 계약 1·2·4와 일치합니다.
- `branch-conventions.md`, `recovery.md`의 force·reset·force push 경계가 유지됩니다.
- runtime reference links를 검증합니다.
- `git diff --check`를 통과합니다.
