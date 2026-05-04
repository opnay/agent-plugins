# turn-gate 스킬 스펙

## 목적

`loop-kit-dev`의 `turn-gate`는 turn continuity를 유지하면서 current-phase work에 맞는 내부 loop mode를 선택해 실행하는 메인 loop controller입니다.
이 skill이 활성화되면 `turn-gate` 메인 플로우는 대화 응답 자체의 1급 제어 규칙이 되며, 사용자의 explicit stop 전까지 결과 보고를 terminal response로 닫지 않습니다.

## 경계

- 포함:
  - turn-level continuity 유지
  - `analysis -> plan -> work -> verification -> result reporting -> next-flow question-routing response` 구조 유지
  - 사용자 메시지의 operation 의미 해독
  - current-phase work의 internal mode selection
  - 결과 보고 뒤 explicit choice 기반 next-flow reopening
  - session record와 Continuity Guard 유지
- 제외:
  - broad workflow taxonomy 자체의 소유
  - 여러 direct loop skill을 사용자에게 노출하는 일
  - domain-specific implementation detail 자체
  - commit execution, push, PR 같은 외부 작업의 세부 실행 계약

## 처리하려는 작업 형태

- 사용자가 턴을 종료하자고 요청하기 전까지 한 턴 안에서 여러 phase를 이어가야 하는 작업
- requirement discovery, autonomous execution, refinement, review-driven correction, readiness checking 같은 current-phase work가 번갈아 나타나는 작업
- 결과 보고 뒤 clean stop이 아니라 다음 플로우 선택이 기본이어야 하는 작업
- 사용자 지시어가 여러 구조 단위를 가리켜 작업 전에 target 또는 operation을 잠가야 하는 작업

## 엔트리포인트 / 대표 표면

- 대표 표면: `loop-kit-dev/skills/turn-gate/SKILL.md`
- 기본 skill spec index: `loop-kit-dev/specs/skills/turn-gate/spec.md`
- 사용자 스펙 의도: `loop-kit-dev/specs/skills/turn-gate/intent.md`
- sub-spec directory: `loop-kit-dev/specs/skills/turn-gate/`
- local runtime references: `loop-kit-dev/skills/turn-gate/references/*.md`
- 관련 상위 라우팅: `loop-kit-dev-guide`

## 상세 계약 구조

`turn-gate`의 사용자 스펙 의도와 세부 계약은 이 index가 직접 모두 소유하지 않고, 아래 sibling spec 문서로 분리합니다.
이 파일은 top-level ownership, routing overview, sibling contract map을 소유합니다.

- `intent.md`: `turn-gate`의 사용자 스펙 의도 기록
- `lifecycle.md`: activation, response ending states, in-turn user intervention, explicit stop 처리
- `session-records.md`: `000-plan.md`, `001+` flow records, Continuity Guard, templates, `Next Flow Options`
- `meaning-resolution.md`: operation/target ambiguity, provenance/intent block target locking, user-gated clarification
- `verification.md`: mandatory clean-context subagent verification, result reporting, next-flow reopening

## 핵심 처리 계약

- `loop-kit-dev/skills/turn-gate/SKILL.md`는 이 스펙의 단순 요약본이 아니라 runtime에서 읽는 운영 표면이다.
- skill body는 대화 응답 자체를 제어하는 conversation-level first-class rule을 `## Important` 섹션으로 앞부분에서 명시해야 한다.
- `## Important` 섹션은 `Purpose`보다 먼저 위치해야 하며, session-level activation, terminal summary 금지, required ending states, `request_user_input` 기반 next-flow reopening, session record 유지 의무를 포함해야 한다.
- skill body에는 `Core Loop` 또는 이에 준하는 단계별 실행 섹션이 있어야 하며, 최소한 analysis, plan, work, verification, result reporting, question-routing reopening을 각각 구분해 설명해야 한다.
- skill body에는 internal mode selection과 local `references/` 읽기 규칙이 직접 남아 있어야 한다.
- skill body에는 terminal summary 금지, next-flow reopening, Continuity Guard 확인, user-gated question routing, explicit turn-end option 기록 규칙이 직접 남아 있어야 한다.
- skill body를 짧게 다듬더라도 위 단계와 금지 규칙을 한 문단으로 뭉개지 말고, 실행 중 빠르게 확인 가능한 형태로 유지한다.

## Internal Mode Selection

- `work`에 들어가기 전 current-phase work의 internal mode를 하나 선택한다.
- `loop-kit-dev`에서는 사용자가 internal mode를 직접 호출하는 대신 `turn-gate`가 이를 선택한다.
- `turn-gate`는 선택된 internal mode에 대응하는 local `references/` 문서를 먼저 읽고 그 계약을 적용한다.
- local `references/`는 `workflow-kit` upstream spec과 동기화된 absorbed operational contract로 유지한다.
- 실제 requirement discovery가 병목이면 `references/deep-interview.md`를 따른다.
- broad end-to-end delivery가 필요한 current phase이면 `references/autopilot.md`를 따른다.
- bounded issue를 작은 fix-verify-reassess cycle로 다루는 경우 `references/ralph-loop.md`를 따른다.
- review feedback이나 material finding 처리인 경우 `references/review-loop.md`를 따른다.
- 현재 변경 단위가 거의 끝났고 readiness 판단이 핵심이면 `references/commit-readiness-gate.md`를 따른다.
- 여러 mode가 겹쳐 보이면 `deep-interview -> review-loop -> ralph-loop -> autopilot -> commit-readiness-gate` 순으로 더 이른 병목을 우선한다.
- 아직 어떤 internal mode가 맞는지 확정되지 않았다면 active question-routing mode나 좁은 분석으로 먼저 mode selection을 잠근다.

## Deep-Interview 원본 관계

- 원본 skill source는 `https://github.com/Yeachan-Heo/oh-my-codex/blob/main/skills/deep-interview/SKILL.md`다.
- 원본은 ambiguity gating, OMX tooling, artifact handoff까지 포함한 더 큰 workflow다.
- `loop-kit-dev`의 `references/deep-interview.md`는 full workflow를 그대로 복제한 것이 아니라, `turn-gate`가 requirement-discovery phase에서 필요한 boundary만 흡수한 derived reference다.

## SSOT 동기화 규칙

- internal mode contract 변경은 먼저 `workflow-kit` upstream spec에서 정리한다.
- `loop-kit-dev`은 runtime orchestration 관점의 차이와 local absorbed references를 별도로 소유한다.
- upstream contract와 `turn-gate` references의 문구가 어긋나면 같은 변경 단위에서 함께 갱신한다.
- 새로운 internal mode는 기존 mode로 current-phase work를 소유할 수 없을 때만 추가한다.
- internal mode set이나 mandatory tool rule이 바뀌면 `workflow-kit` upstream spec, `loop-kit-dev` plugin spec, `loop-kit-dev-guide`, `turn-gate`, `turn-gate/references/`를 함께 갱신한다.

## 검토 질문

- 이번 응답이 turn continuity를 실제로 유지하고 있는가?
- skill body 앞부분에 `Important` 섹션이 있고 1급 규칙, terminal summary 금지, next-flow reopening이 먼저 드러나는가?
- 사용자 표현에 구조적 다의성이 있으면 internal mode 선택 전에 meaning resolution 질문을 열었는가?
- current-phase work에 맞는 internal mode를 하나로 좁혔는가?
- user-gated question routing과 계획 도구 `update_plan`를 필수 단계에서 실제로 사용했는가?
- cross-flow 작업이라면 `.agents/sessions/{YYYYMMDD}/000-plan.md`가 planned flow sequence, 각 flow의 완료 기준, 다음 flow 전환 조건을 최신 상태로 담고 있는가?
- `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md`가 현재 phase까지 증분 갱신됐는가?
- `work -> verification -> result reporting` 순서를 실제로 유지했는가?
- direct loop entrypoint를 사용자 표면으로 다시 열지 않았는가?
- 결과 보고 뒤 explicit next-flow choice를 실제로 열었는가?
- 결과 보고 직전에 `Continuity Guard`를 갱신했고 terminal summary 가능 여부를 확인했는가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 아니오.
- 그렇다면 왜 필요한가 / 아니라면 어떤 sibling context를 허용하는가: 이 skill은 `workflow-kit`의 canonical loop-mode contract와 `loop-kit-dev`의 narrow runtime packaging을 전제로 한다. 다만 turn continuity와 mode selection rule 자체는 이 index spec에서 명시적으로 읽혀야 한다.

## 확장 원칙

- `spec.md`는 top-level ownership과 routing map을 유지하고, 상세 계약이 커지면 같은 폴더 아래 sub-spec으로 내린다.
- 새 sub-spec을 추가할 때는 `spec.md`의 상세 계약 구조와 plugin spec의 skill spec 위치 설명을 함께 갱신한다.
