# self-drive overlay mode spec

## 목적

`self-drive` mode는 사용자 메시지 기반 preparation이 만든 planned flow sequence에 autonomous continuation을 적용하는 mode입니다.
이 mode는 별도 installed skill entrypoint가 아니라 명시적으로 적용되는 독립 overlay 계약입니다.
이 문서는 self-drive의 endpoint, stop boundary, execution authority, user-gated 복귀 조건을 소유하는 spec-side SSOT입니다.
명시적으로 self-drive가 적용된 prepared sequence에서는 `references/self-drive.md`의 runtime 계약이 turn-gate 기본 routing보다 우선 적용됩니다.
이후 진행 판단은 self-drive 계약이 소유하며, `turn-gate` 본체는 explicit stop 처리, 안전/도구 승인 한계, session record, 보고 형식처럼 self-drive가 재사용하는 최소 loop 기반만 제공합니다.

## 계약

- `self-drive` mode는 준비된 planned flow sequence, scope, non-goal, acceptance signal, approval boundary, verification expectation이 충분히 기록된 뒤에만 적용한다.
- self-drive가 적용되기 전에 session record에는 self-drive active 여부, sidecar pointer, sequence objective, planned flow list, active flow index, current flow label, allowed autonomous actions, prohibited autonomous actions, approval-sensitive checkpoints, endpoint, blocker return conditions, progress note가 드러나야 한다.
- `active_flow_index`는 0-based machine field로 취급하고, 사람이 읽는 `Current flow`에는 flow number/name/file 또는 slug를 함께 기록한다. 기존 record의 numeric index와 planned flow numbering이 충돌하거나 기준이 불명확하면 flow name/file을 확인하고, 해소할 수 없으면 autonomous continuation 전에 user-gated clarification으로 돌아간다.
- `000-plan.md`는 self-drive-specific date-level snapshot으로 self-drive active 여부와 `000-self-drive.md` pointer만 소유한다.
- sequence-level state는 optional sidecar record인 `000-self-drive.md`가 소유하고, 각 active flow record는 자기 flow의 sequence position, progress note, next handoff, blocker return condition만 flow-local snapshot으로 남긴다.
- self-drive가 적용되는 동안 prepared sequence의 진행 판단은 turn-gate 기본 routing, next-flow 질문 기본값, phase protocol 선택보다 self-drive mode 계약을 우선한다.
- 각 planned flow는 여전히 자기 내부에서 `preparation -> work -> verification -> reporting -> next-flow` core loop를 가진다.
- 각 planned flow의 reporting 전후에는 self-drive sequence record를 현재 active flow index, current flow label, progress note, next handoff, blocker 여부에 맞게 갱신한다.
- self-drive는 초기 preparation에서 exact action, target, expected effect, risk, recovery path, 포함/제외 scope, 종료 지점이 기록된 approval-sensitive action만 추가 질문 없이 실행할 수 있다.
- commit, push, PR, publish, release, version bump는 approval-sensitive execution step이다. 초기 합의에 포함되어 있고 종료 지점이 명확하면 self-drive 안에서 실행할 수 있으며, 그렇지 않으면 user-gated question routing으로 돌아간다.
- 초기 협의 범위 밖의 위험 작업, 새 approval boundary, 또는 종료 지점이 불명확한 실행이 나타나면 implicit default state의 user-gated question routing으로 돌아간다.
- planned flow sequence가 끝나면 terminal close가 아니라 기록된 종료 지점, commit-readiness reporting handoff, 또는 next-flow reopening으로 이어진다.
- runtime 실행 계약은 `skills/turn-gate/references/self-drive.md`에 흡수된 reference를 사용한다.
- 주변 spec과 `turn-gate` 본체는 self-drive의 세부 판단 조건을 반복하지 않는다.

## 검토 질문

- self-drive에 필요한 planned flow sequence와 approval boundary가 기록돼 있는가?
- self-drive active 여부와 sidecar pointer는 `000-plan.md`에 있고, sequence objective, active flow index, allowed/prohibited action, checkpoint, endpoint, blocker return condition, progress note는 `000-self-drive.md`와 active flow record의 소유권에 맞게 기록돼 있는가?
- user-gated로 되돌려야 할 새 위험 작업, 범위 확장, 또는 불명확한 종료 지점이 생기지 않았는가?
- 마지막 flow 뒤 commit-readiness reporting과 commit/push/PR 같은 execution step의 승인 근거를 구분했는가?
