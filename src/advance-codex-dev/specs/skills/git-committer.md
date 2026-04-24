## 사용자 스펙 의도

- 직접 만든 변경을 task-scoped commit으로 깔끔하게 마무리하고 싶다.
- staged scope 검토와 수동 검증을 거친 뒤 commit하고 싶다.
- commit message 형식과 body 품질을 같이 통제하고 싶다.

---

# git-committer 스킬 스펙

## 목적

`git-committer`는 직접 만든 변경을 검토 가능한 task-scoped commit으로 정리하기 위한 finalization 스킬입니다.
핵심은 staged scope 검토, 수동 검증, 명확한 commit message, 최종 확인을 한 흐름으로 묶는 것입니다.

## 경계

- 포함:
  - commit 범위 분리와 staged diff 검토
  - manual CI 또는 relevant verification 실행
  - colon-separated commit message와 body 작성 규율
  - 최종 commit 확인
- 제외:
  - unrelated change cleanup
  - interactive git tutoring 전반
  - implementation 자체의 설계

## 처리하려는 작업 형태

- 변경을 commit-ready 상태로 마무리해야 하는 경우
- mixed change를 task-scoped unit으로 나눠야 하는 경우
- commit message quality와 final verification이 중요한 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `advance-codex-dev/skills/git-committer/SKILL.md`
- 관련 상위 라우팅: `advance-codex-dev-guide`, `workflow-kit/commit-readiness-gate`

## 핵심 처리 계약

- staged scope를 먼저 검토하고 혼합된 변경은 분리한다.
- commit 직전 relevant verification을 다시 실행하고 결과를 기준으로 message를 작성한다.
- 첫 줄은 `type: detailed subject` 형식을 따르고 body는 변경 내용과 검증을 구체적으로 담는다.
- commit 이후에는 최신 로그나 commit metadata로 형식과 범위를 다시 확인한다.

## 검토 질문

- staged scope가 정말 하나의 task-scoped change unit으로 잠겨 있는가?
- commit 직전 relevant verification을 다시 돌렸는가?
- commit message의 subject와 body가 변경 범위와 검증 내용을 충분히 설명하는가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 예.
- 그렇다면 왜 필요한가 / 아니라면 어떤 sibling context를 허용하는가: commit finalization 규율은 hidden sibling context 없이도 재사용 가능해야 하며, 다른 skill은 readiness 판단까지만 맡고 실제 finalization 규칙은 이 스펙이 소유한다.

## 확장 원칙

- 새 reference는 message format, command usage, verification example처럼 finalization 품질에 직접 기여할 때만 추가한다.
- commit workflow 단계가 바뀌면 reference와 spec을 함께 갱신한다.
