# review template 계약

## 소유 범위

`flow`는 `000-review.md`를 session retrospective note 표면으로 정의합니다.
runtime template은 `skills/flow/templates/review.md`가 제공합니다.
review record는 active routing, flow log, verification authority, closure authority가 아닙니다.

## 계약

- 파일명: `.agents/sessions/{YYYYMMDD}/000-review.md`.
- 생성 규칙: session date마다 optional 파일 하나만 둡니다.
- 작성 타이밍: 메인 플로우 그룹 이후, `handoff condition` 직전에 회고 note가 있을 때만 작성합니다.
- 형식: flat tagged list.
- 항목: `[axis]` note, optional invalid example, corrected pattern, follow-up owner 또는 candidate.
- 태그 예: `[conversation]`, `[records]`, `[skills]`, `[docs]`, `[code-structure]`, `[verification]`, `[git]`, `[release]`.
- 기록 대상: 작업 중 발견한 skill contract 위반, 회고 가치가 있는 운영 규칙, 후속 owner 후보.

## 제외

- active routing state
- raw flow logs
- full verification evidence
- commit, release, closure authority
- 이미 flow record에 있는 실행 이력 반복

## 검토 기준

- 회고 note가 다음 개선 판단에 도움이 되는가?
- active turn 진행 상태와 섞이지 않는가?
- owner surface 또는 follow-up candidate가 짧게 드러나는가?
