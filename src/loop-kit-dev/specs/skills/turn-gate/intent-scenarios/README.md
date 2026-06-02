# turn-gate intent scenarios

이 폴더는 `turn-gate` continuity, routing, verification, session record, self-drive 경계를 회귀 평가하는 spec-side fixture를 보관합니다.
시나리오는 runtime instruction이 아니며, skill 문구나 planning 동작을 바꿀 때 평가 입력으로 사용합니다.

## fixture 계약

각 시나리오는 다음 항목을 기록해야 합니다.

- 사용자 메시지
- expected task tier
- expected verification method
- 기대하는 `operational-preparation` flow 동작
- 기대하는 `change-unit` planned flows
- flow가 아니어야 하는 항목
- 수용 신호

`operational-preparation`은 요청 해석, intent/scope 정렬, approval-boundary 계획, planned-flow-list 설계가 session 또는 plan artifact를 만들 때의 운영 flow입니다.
`change-unit`은 검토 가능한 코드, 문서, fixture, config, release-surface 변경을 소유하는 실행 flow입니다.
`expected task tier`와 `expected verification method`는 runtime instruction이 아니라 fixture 평가 기준입니다.
`expected verification method`는 기본적으로 `clean-context`, `normal`, `not-required` 중 하나를 기준으로 적습니다.
조건부 method가 필요한 시나리오는 `allowed verification methods`와 각 method의 selection criteria 또는 primary/upgrade trigger를 함께 적어야 합니다.
조건부 문구만으로 "아무 method나 허용"하지 않으며, file-change, release/build, approval-sensitive, strong regression fixture는 단일 expected method를 유지합니다.
경량화 시나리오라도 approval-sensitive boundary, 파일 변경 시 검증 기본값, non-pass routing을 약화하면 안 됩니다.

## 보존 기준

`turn-gate` 고유 경계를 직접 압박하는 fixture만 보존합니다.
일반 flow 분류 예시, verification 정책 예시, 중복 status/routing 예시는 제거합니다.

## 현재 fixture

- `commit-completion-continuation-flow.md`: 커밋 완료가 explicit stop이 아니며, 보고 뒤 next-flow question-routing으로 이어지는지 확인합니다.
- `explicit-stop-source-matching.md`: current explicit stop, future endpoint stop, source-less/stale closure, compaction summary ambiguity를 구분합니다.
- `non-pass-verification-routing.md`: `fail`, `insufficient`, `blocked` verification 결과가 성공 보고, terminal summary, next-flow continuation으로 잘못 흡수되지 않는지 확인합니다.
- `not-required-status-result-boundary.md`: `Method: not-required`가 automatic pass가 아니며 `pass`/`insufficient`/`blocked` status와 분리되는지 확인합니다.
- `phase-prefix-application.md`: phase-start/progress 메시지의 prefix 누락과 record/artifact/question option 내부 prefix 과잉 적용을 확인합니다.
- `question-tool-autonomy-boundary.md`: 질문 도구 과잉 사용과 self-drive 질문 부족 사용을 구분합니다.
- `read-only-session-record-boundary.md`: target/source read-only와 workspace-wide no-write/no-record 요청을 분리해 session record 운영 기록 작성 여부를 확인합니다.
- `self-drive-reporting-auto-advance.md`: active self-drive reporting 뒤 정상 자동 전환과 user-gated 복귀 조건을 구분합니다.
- `self-drive-sequence-record.md`: self-drive 긴 planned flow sequence에서 sequence-level record와 flow-local snapshot이 분리되는지 확인합니다.
- `session-record-reconstruction-boundary.md`: session record first creation, active missing, inaccessible, stale closure, stale sidecar의 회복 경계가 silent reconstruction으로 뭉개지지 않는지 확인합니다.
- `stale-session-record-authority.md`: stale sidecar, source-less/stale closure, stale routing mismatch, inaccessible/corrupt record가 terminal closure나 autonomous continuation authority로 잘못 승격되지 않는지 확인합니다.
- `unbounded-self-drive-endpoint.md`: open-ended self-drive 요청을 finite cycle과 endpoint/repeat policy로 안전하게 기록하는지 확인합니다.
