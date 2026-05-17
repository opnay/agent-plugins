# turn-gate intent scenarios

이 폴더는 `turn-gate`가 의도한 flow boundary 의미를 유지하는지 확인하기 위한 intent 시나리오를 보관합니다.

이 시나리오들은 그 자체로 runtime instruction이 아니라 spec-side fixture입니다.
skill 문구를 재생성하거나 flow planning 동작을 바꾸는 경우 평가 입력으로 사용합니다.

## 시나리오 계약

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
`expected verification method`는 `clean-context`, `normal`, `not-required` 중 하나를 기준으로 적습니다.
경량화 시나리오라도 approval-sensitive boundary, 파일 변경 시 검증 기본값, non-pass routing을 약화하면 안 됩니다.

## 현재 시나리오

- `commit-completion-continuation-flow.md`: 커밋 완료가 explicit stop이 아니며, 보고 뒤 next-flow question-routing으로 이어지는지 확인합니다.
- `approval-sensitive-release-verification.md`: release/version/publish 계열 요청에서 approval-sensitive boundary와 stronger verification이 경량화보다 우선하는지 확인합니다.
- `file-change-verifier-default.md`: 파일 변경이 있는 fixture 추가 요청에서 clean-context verifier 기본값이 유지되는지 확인합니다.
- `login-page-flow.md`: 넓은 기능 요청이 곧바로 구현 flow로 가지 않고 Flow 0에서 scope와 후속 실행 후보로 정리되는지 확인합니다.
- `micro-readonly-research-verification.md`: no-edit read-only 조사 요청에서 verification phase는 유지하되 evidence checklist/source readback으로 충분할 수 있는지 확인합니다.
- `not-required-routing-verification.md`: routing-only 요청에서 검증할 work output이 없을 때 `not-required` method를 status와 분리해 기록하는지 확인합니다.
- `small-copy-fix-flow.md`: 작은 문구 수정 요청을 불필요하게 여러 planned flow로 쪼개지 않는지 확인합니다.
- `dashboard-stale-list-bug-flow.md`: 버그 수정 요청에서 원인 파악, 수정, 검증을 phase별 planned flow로 쪼개지 않는지 확인합니다.
- `self-drive-sequence-record.md`: self-drive 긴 planned flow sequence에서 sequence-level record와 flow-local snapshot이 분리되는지 확인합니다.
- `self-drive-mid-sequence-status.md`: self-drive 실행 중 status/progress 질문이 terminal close나 next-flow replacement가 아니라 상태 보고 뒤 continuation으로 처리되는지 확인합니다.
- `self-drive-priority-change.md`: self-drive 실행 중 planned flow 우선순위나 scope가 바뀌면 autonomous continuation을 멈추고 updated sequence를 다시 잠그는지 확인합니다.
- `session-record-reconstruction-boundary.md`: session record first creation, active missing, inaccessible, stale closure, stale sidecar의 회복 경계가 silent reconstruction으로 뭉개지지 않는지 20개 case로 확인합니다.
- `stale-session-record-authority.md`: stale sidecar, source-less/stale closure, stale routing mismatch, inaccessible/corrupt record가 terminal closure나 autonomous continuation authority로 잘못 승격되지 않는지 20개 case로 확인합니다.
- `non-pass-verification-routing.md`: `fail`, `insufficient`, `blocked` verification 결과가 성공 보고, terminal summary, next-flow continuation으로 잘못 흡수되지 않는지 20개 case로 확인합니다.
- `read-only-session-record-boundary.md`: target/source read-only와 workspace-wide no-write/no-record 요청을 분리해 session record 운영 기록 작성 여부를 20개 case로 확인합니다.
