## 사용자 스펙 의도

- 하나의 턴을 사용자가 턴을 종료하자고 요청할때까지 닫지 않고 유지하고 싶다.
- `turn-gate`가 현재 phase의 메인 작업을 보고 적절한 내부 loop mode를 고르길 원한다.
- 필요한 경우 requirement discovery 성격의 `deep-interview`도 `turn-gate` 안의 internal mode로 흘러가길 원한다.
- `ralph-loop`, `review-loop`, readiness gate 같은 loop는 사용자가 직접 고르지 않고 `turn-gate` 안에서 흘러가길 원한다.
- 내부 loop mode의 canonical contract는 `workflow-kit`이 SSOT로 계속 소유하길 원한다.
- 내부 loop mode는 `turn-gate` 스킬의 local `references/` 아래로 흡수돼야 하고, 실행 시 그 reference를 읽는 구조이길 원한다.
- `turn-gate`는 질문 도구와 계획 도구를 선택 사항이 아니라 필수 도구로 사용해야 한다.
- `turn-gate`로 진행한 작업은 `.agents/sessions` 아래에 기록이 남아야 한다.
- 여러 플로우를 거치는 작업의 상위 계획은 `.agents/sessions/{YYYYMMDD}/000-plan.md` 경로에 누적되길 원한다.
- 개별 플로우 기록은 `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` 형식으로 남고 싶다.
- `001+` 문서는 phase 메모가 아니라 flow 기록으로 남고 싶다.

---

# turn-gate 스킬 스펙

## 목적

`loop-kit`의 `turn-gate`는 turn continuity를 유지하면서 current-phase work에 맞는 내부 loop mode를 선택해 실행하는 메인 loop controller입니다.

## 경계

- 포함:
  - turn-level continuity 유지
  - `analysis -> plan -> work -> verification -> result reporting -> next-flow user response` 구조 유지
  - current-phase work의 internal mode selection
  - 결과 보고 뒤 explicit choice 기반 next-flow reopening
- 제외:
  - broad workflow taxonomy 자체의 소유
  - 여러 direct loop skill을 사용자에게 노출하는 일
  - domain-specific implementation detail 자체

## 처리하려는 작업 형태

- 사용자가 턴을 종료하자고 요청하기 전까지 한 턴 안에서 여러 phase를 이어가야 하는 작업
- requirement discovery, refinement, review-driven correction, readiness checking 같은 current-phase work가 번갈아 나타나는 작업
- 결과 보고 뒤 clean stop이 아니라 다음 플로우 선택이 기본이어야 하는 작업

## 엔트리포인트 / 대표 표면

- 대표 표면: `loop-kit/skills/turn-gate/SKILL.md`
- local reference surface: `loop-kit/skills/turn-gate/references/*.md`
- 관련 상위 라우팅: `loop-kit-guide`

## 핵심 처리 계약

- 각 incoming message를 같은 loop-gated turn의 현재 입력으로 취급한다.
- `analysis`, `plan`, `work`, `verification`, `result reporting`, `user-response reopening`을 응답 shape에 계속 드러낸다.
- active turn-gated task마다 `.agents/sessions/{YYYYMMDD}/000-plan.md` 상위 계획과 `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` record 체계를 유지한다.
- `000-plan.md`는 사용자 요청 종료 이후에도 더 큰 작업이 이어지면 계속 증분 갱신한다.
- completed flow가 끝날 때마다 해당 `001+` record를 갱신한다.
- 질문, 선택지 제시, scope lock, next-flow reopening에는 질문 도구 `request_user_input`를 필수로 사용한다.
- 실질적인 작업이 시작되면 계획 도구 `update_plan`을 필수로 사용하고 현재 active step 상태를 유지한다.
- `work`에 들어가기 전 current-phase work의 internal mode를 하나 선택한다.
- `loop-kit`에서는 사용자가 internal mode를 직접 호출하는 대신 `turn-gate`가 이를 선택한다.
- `turn-gate`는 선택된 internal mode에 대응하는 local `references/` 문서를 먼저 읽고 그 계약을 적용한다.
- local `references/`는 `workflow-kit` upstream spec과 동기화된 absorbed operational contract로 유지한다.
- `work` 뒤에는 결과 보고 전에 명시적 검증 단계를 둔다.
- 결과 보고 뒤에는 explicit choice를 주는 질문 도구로 다음 플로우를 다시 연다.
- 사용자가 턴을 종료하자고 요청하지 않으면 clean stop을 기본 경로로 두지 않는다.

## 내부 loop mode 선택 규칙

- 실제 requirement discovery가 병목이면 `references/deep-interview.md`를 따른다.
- bounded issue를 작은 fix-verify-reassess cycle로 다루는 경우 `references/ralph-loop.md`를 따른다.
- review feedback이나 material finding 처리인 경우 `references/review-loop.md`를 따른다.
- 현재 변경 단위가 거의 끝났고 readiness 판단이 핵심이면 `references/commit-readiness-gate.md`를 따른다.
- 아직 어떤 internal mode가 맞는지 확정되지 않았다면 질문 도구나 좁은 분석으로 먼저 mode selection을 잠근다.

## mode selection matrix

- blocker가 requirement discovery, intent ambiguity, scope boundary, approval line이라면 `deep-interview`를 고른다.
- input이 review feedback, QA finding, self-review finding이고 한 번에 하나의 material issue만 처리해야 한다면 `review-loop`를 고른다.
- blocker가 하나의 bounded improvement cycle이고 작은 fix 뒤 즉시 검증하는 흐름이 맞다면 `ralph-loop`를 고른다.
- 구현이 거의 끝났고 intended change unit의 commit readiness 판단이 핵심이면 `commit-readiness-gate`를 고른다.
- 여러 mode가 겹쳐 보이면 `deep-interview -> review-loop -> ralph-loop -> commit-readiness-gate` 순으로 더 이른 병목을 우선한다.
- 그래도 mode를 못 잠그면 질문 도구로 선택지를 좁힌 뒤 work phase로 들어간다.

## next-flow reopening 규칙

- 다음 플로우 질문은 현재 결과에 직접 연결된 좁은 선택지여야 한다.
- generic follow-up phrase나 자유형 마무리 질문으로 턴을 닫지 않는다.
- 사용자 응답은 같은 턴의 다음 메시지로 즉시 이어진다.

## session record 규칙

- 여러 플로우를 거치는 작업이면 `.agents/sessions/{YYYYMMDD}/000-plan.md`를 먼저 둔다.
- `000-plan.md`는 상위 multi-flow plan artifact로 유지하고, 요청 종료 뒤에도 같은 큰 작업이 이어지면 계속 증분한다.
- flow 기록 파일은 `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` 형식을 사용한다.
- `count-pad3`는 `001`, `002`, `003`처럼 3자리 zero-padded 숫자를 사용한다.
- slug는 영어 소문자와 `-`만 사용한다.
- flow 기본 템플릿은 `.agents/sessions/_turn-gate-flow-template.md`를 사용한다.
- `000-plan.md` 기본 템플릿은 `.agents/sessions/_turn-gate-plan-template.md`를 사용한다.
- 최소 flow 기록 항목은 user request message, task, flow scope, current mode, analysis, plan, work, verification, result report, next-flow options, residual risk다.
- `000-plan.md`는 장기 증분 계획 artifact로, `001+`는 flow 단위 운영 artifact로 취급한다.

## SSOT 동기화 규칙

- internal mode contract 변경은 먼저 `workflow-kit` upstream spec에서 정리한다.
- `loop-kit`은 runtime orchestration 관점의 차이와 local absorbed references를 별도로 소유한다.
- upstream contract와 `turn-gate` references의 문구가 어긋나면 같은 변경 단위에서 함께 갱신한다.

## 검토 질문

- 이번 응답이 turn continuity를 실제로 유지하고 있는가?
- current-phase work에 맞는 internal mode를 하나로 좁혔는가?
- 질문 도구 `request_user_input`와 계획 도구 `update_plan`를 필수 단계에서 실제로 사용했는가?
- cross-flow 작업이라면 `.agents/sessions/{YYYYMMDD}/000-plan.md`가 최신 상태인가?
- `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md`가 현재 flow까지 갱신됐는가?
- `work -> verification -> result reporting` 순서를 실제로 유지했는가?
- direct loop entrypoint를 사용자 표면으로 다시 열지 않았는가?
- 결과 보고 뒤 explicit next-flow choice를 실제로 열었는가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 아니오.
- 그렇다면 왜 필요한가 / 아니라면 어떤 sibling context를 허용하는가: 이 skill은 `workflow-kit`의 canonical loop-mode contract와 `loop-kit`의 narrow runtime packaging을 전제로 한다. 다만 turn continuity와 mode selection rule 자체는 이 스펙에서 명시적으로 읽혀야 한다.

## 확장 원칙

- 새로운 internal mode는 기존 mode로 현재 phase work를 소유할 수 없을 때만 추가한다.
- internal mode set이나 mandatory tool rule이 바뀌면 `workflow-kit` upstream spec, `loop-kit` plugin spec, `loop-kit-guide`, `turn-gate`, `turn-gate/references/`를 함께 갱신한다.
