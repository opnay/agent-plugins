# turn-gate self-drive sidecar template sub-spec

## 목적

이 문서는 `turn-gate`의 optional `skills/turn-gate/templates/self-drive-template.md` 형식 계약을 소유합니다.

`000-self-drive.md`는 `.agents/sessions/{YYYYMMDD}/` 아래에서 self-drive가 active일 때만 생성되는 000-level 보조 record입니다.
`000-plan.md`와 같은 디렉터리 위계에 있지만, `000-plan.md`를 대체하지 않습니다.

## 소유

- self-drive sequence objective
- planned flow list
- active flow index as a 0-based machine field
- current flow label with human-readable number/name/file or slug
- allowed autonomous actions
- prohibited autonomous actions
- approval-sensitive checkpoints
- endpoint
- blocker return conditions
- progress ledger

## 비소유

- 날짜 단위 전체 flow index
- completed flow summaries
- current active flow의 canonical scope/evidence/report
- terminal closure authority
- commit/push/PR/publish/release/version-bump approval

위 항목은 각각 `000-plan.md`, `001+` flow record, approval boundary 계약이 소유합니다.

## 필수 구조

`self-drive-template.md`는 다음 구조를 유지합니다.

1. YAML frontmatter
2. `Sequence Contract`
3. `Autonomous Boundary`
4. `Progress Ledger`
5. `User-Gated Return Conditions`
6. `Residual Risk`

## Frontmatter 규격

frontmatter에는 다음 필드를 둡니다.

- `sequence_objective`
- `active_flow_index`
- `planned_flow_count`
- `endpoint`
- `status`
- `last_updated_flow`
- `required_next_action`

## 중복 방지 규칙

- `000-plan.md`에는 self-drive active 여부와 `000-self-drive.md` pointer만 둡니다.
- `000-self-drive.md`는 sequence-level state를 소유하고, 각 flow의 상세 scope/evidence/report를 반복하지 않습니다.
- `active_flow_index`는 0-based machine field로 쓰고, `Current flow`에는 사람이 읽는 번호, 이름, 파일 또는 slug를 같이 써서 off-by-one ambiguity를 피합니다.
- `001+` flow record는 자기 flow의 local progress, next handoff, blocker return condition만 남기고 전체 planned flow list를 반복하지 않습니다.
- self-drive가 active가 아니면 `000-self-drive.md`를 만들지 않습니다.

## 검토 질문

- `000-plan.md`가 self-drive pointer만 소유하고 sequence detail을 반복하지 않는가?
- `000-self-drive.md`가 sequence-level state를 충분히 담고 있는가?
- active flow record와 `000-self-drive.md`의 active flow index/current flow label이 서로 맞는가?
- approval-sensitive checkpoint가 execution approval로 오해되지 않게 기록됐는가?
- self-drive 종료 후 terminal close가 아니라 endpoint, blocker decision, commit-readiness handoff, next-flow reopening 중 하나로 이어지는가?
