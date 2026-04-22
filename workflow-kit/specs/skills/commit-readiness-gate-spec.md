# commit-readiness-gate 스킬 스펙

## 목적

`commit-readiness-gate`는 intended change unit이 commit 쪽으로 넘어갈 준비가 되었는지 판단하는 final gate 스킬입니다.

## 경계

- 포함:
  - final self-review
  - scoped verification
  - risk classification
  - minimum review recommendation
- 제외:
  - broad implementation
  - planning
  - commit message 작성 자체

## 처리하려는 작업 형태

- 구현이 거의 끝났고 commit-ready 상태인지 판단해야 하는 경우
- 마지막 review와 verification으로 change unit을 잠가야 하는 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `workflow-kit/skills/commit-readiness-gate/SKILL.md`
- 관련 하위/후속: `advance-codex:git-committer`

## 핵심 처리 계약

- intended change unit을 기준으로 self-review와 scoped verification을 수행한다.
- readiness checklist와 residual risk를 함께 보고한다.
- gate가 통과되고 실제 commit execution이 다음 단계라면 `advance-codex:git-committer`로 명시적으로 handoff한다.
- commit message 작성이나 git finalization은 이 gate 안에 머물지 않고 후속 skill로 넘긴다.
- user-facing wording에서 `commit-ready`는 gate 통과 상태를 가리킬 때만 쓰고, pre-gate action label과 혼용하지 않는다.

## 독립성 원칙

- 이 스킬은 implementation이나 commit finalization을 소유하지 않는다.
- gate 결과만 읽어도 commit-ready 여부를 판단할 수 있어야 한다.

## 확장 원칙

- 새 gate rule은 readiness 판단의 신뢰도를 높일 때만 추가한다.
