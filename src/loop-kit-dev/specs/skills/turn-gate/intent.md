## 사용자 스펙 의도

- `turn-gate` 활성화는 세션 단위의 1급 응답 규칙이어야 한다.
  - 하나의 턴을 사용자가 턴을 종료하자고 요청할때까지 닫지 않고 유지하고 싶다.
  - 이 스킬을 사용한다는 것은 이 세션 동안 이 스킬을 1급 규칙으로 사용한다는 의미다.
  - 메인 플로우는 스킬 내부 체크리스트가 아니라 대화 응답 자체를 제어해야 한다.
  - 일반 목적 설명보다 높은 우선순위로 보이도록 skill body 앞부분의 `Important` 섹션에 드러나야 한다.
  - 활성화된 동안 응답은 loop continuation, question-routing, explicit user stop 처리 중 하나로 끝나야 하며, 일반적인 final summary로 턴을 닫으면 안 된다.
  - incoming message 처리는 특정 상황 목록에 갇힌 closed taxonomy가 아니라, 명시적 turn stop이 아닌 모든 사용자 입력을 gated turn continuation으로 해석하는 포괄 규칙이어야 한다.
  - 질문, 검토 요청, 상태 확인, 우선순위 변경, correction 같은 표현은 예시일 뿐이며, 예시에 없는 입력도 explicit stop이 아니면 보고 후 멈추는 근거가 될 수 없다.

- `turn-gate`가 current-phase work에 맞는 내부 loop mode를 고르길 원한다.
  - `turn-gate`의 가장 기본 flow는 `준비 -> 작업 -> 검증 -> 보고`여야 한다.
  - 이 기본 flow를 둘러싼 전환은 내부 gate로 설명되길 원한다.
  - message intake gate는 사용자 메시지를 분류하지만 실행하지 않고, flow shaping gate는 active flow와 completion criteria를 만들거나 갱신하며, task policy gate는 flow 내부 실행 정책만 소유해야 한다.
  - task policy는 flow 밖의 독립 계층이 아니며, 개별 task 완료가 flow 완료나 turn closure를 결정하면 안 된다.
  - reporting 뒤에는 source-recorded explicit stop이 없는 한 continuation gate가 next-flow reopening으로 이어져야 한다.
  - deep-interview, flow list design, 상태 파악, 수정 범위 파악, 질문 라우팅, 내부 mode 선택은 기본 flow 자체가 아니라 `준비` 안의 세부 작업이나 파생 작업이어야 한다.
  - 사용자 메시지에서 시작하는 준비는 deep-interview를 사용해 intent, scope, 성공 기준, approval boundary를 확인하고, 사용자 의도에 맞는 이후 flow list를 준비해야 한다.
  - 사용자 메시지를 받아 의도를 해석하고 flow list를 만드는 일 자체도 산출물을 가진 운영 flow가 될 수 있어야 한다. 이때 산출물은 session plan, flow record, planned flow list, scope/approval boundary다.
  - 이 운영 flow와 실제 product/code/document 변경을 소유하는 flow는 구분되어야 한다.
  - flow list는 `분석`, `작업`, `검증`, `커밋 준비` 같은 진행 phase를 나열하는 것이 아니어야 한다.
  - flow는 반드시 최종 사용자에게 직접 보이는 가치 단위일 필요도 없다. 대신 함께 이해하고 검토하고 검증하고 필요하면 커밋할 수 있는 응집된 변경 단위여야 한다.
  - 예를 들어 "로그인 페이지를 만들자"라는 요청의 직접 사용자 가치는 로그인 페이지 전체일 수 있지만, planned flow는 `로그인 UI/UX 컴포넌트 생성`, `로그인 로직 작성`, `로그인 페이지 조립`처럼 보이지 않는 준비성 작업을 포함한 커밋 가능 단위로 나눌 수 있어야 한다.
  - 비 사용자 메시지에서 시작하는 준비는 이미 준비된 flow에 대한 준비 과정이며, 필요한 수정 범위 파악, 현재 상태 파악, 대상 파일 재확인, 실행 전 조건 확인을 수행해야 한다.
  - 작업은 사용자가 요청한 실제 작업을 진행하는 단계이며, 파일 수정, 검증, 조사 같은 다양한 작업 유형을 포함할 수 있다.
  - 검증은 해당 flow에 대한 검증 단계이며, 파일 수정이 제대로 적용됐는지, 타입 오류가 없는지, 조사의 경우 다양한 관점에서 논리 비판을 수행했는지 확인해야 한다.
  - 보고는 turn 종료가 아니라 다음 flow 진행을 위한 이번 flow의 맥락 정리 단계여야 한다.
  - 계획된 flow가 소진되면 보고 단계에서 질문 도구를 사용해 사용자에게 다음 flow나 작업을 받아야 한다.
  - requirement discovery 성격의 `deep-interview`도 내부 mode로 흘러가길 원한다.
  - `autopilot`, `ralph-loop`, `review-loop`, readiness gate 같은 loop는 사용자가 직접 고르지 않고 `turn-gate` 안에서 선택되길 원한다.
  - 내부 loop mode의 canonical contract는 `workflow-kit`이 SSOT로 계속 소유하길 원한다.
  - 실행 시에는 `turn-gate` skill의 local `references/` 아래로 흡수된 contract를 읽는 구조이길 원한다.

- 질문, 계획, 다음 플로우 선택은 user-gated 운영으로 드러나야 한다.
  - 질문, 선택지, scope lock, 다음 플로우 선택은 기본적으로 user-gated question routing으로 처리하길 원한다.
  - 질문 도구와 계획 도구는 선택 사항이 아니라 필수 도구로 사용해야 한다.
  - 다음 플로우 질문의 사용자 표시 선택지가 3개 이상이라 턴 종료 선택지를 표시하지 못하더라도, sessions flow record의 `Next Flow Options`에는 명시적인 턴 종료 선택지가 항상 남아야 한다.

- turn-gated 작업은 `.agents/sessions` 아래에 multi-flow 기록으로 남아야 한다.
  - 상위 계획은 `.agents/sessions/{YYYYMMDD}/000-plan.md`에 누적되길 원한다.
  - `000-plan.md`는 단순 현재 상태 로그가 아니라 여러 flow의 상위 계획과 흐름을 소유해야 한다.
  - `000-plan.md`의 planned flow sequence는 phase checklist가 아니라 응집된 변경 단위들의 순서여야 한다.
  - 사용자 메시지 해석과 planned flow list 설계가 필요한 경우, `000-plan.md`와 active flow record에는 이것이 `operational-preparation` flow인지, 이후 항목이 `change-unit` flow인지 드러나야 한다.
  - 개별 플로우 기록은 `.agents/sessions/{YYYYMMDD}/{count-pad3}-{eng-lower-slug}.md` 형식으로 남고 싶다.
  - `001+` 문서는 phase 메모가 아니라 flow 기록으로 남고 싶다.
  - 예를 들어 "컴포넌트 오타가 있어. 수정하자" 같은 요청은 작은 변경이라면 하나의 `컴포넌트 문구 수정` flow로 충분할 수 있고, 그 flow 안에서 점검, 수정, 검증, 보고를 수행해야 한다. 별도 flow는 검토/커밋 가능한 변경 단위가 실제로 나뉠 때만 만든다.

- 검증은 main agent의 same-context self-check가 아니라 clean-context subagent 검증이어야 한다.
  - 검증 단계는 main agent가 같은 context에서 직접 수행하지 않고, 무조건 clean context 상태의 subagent가 수행해야 한다.
  - `turn-gate` 활성 상태에서는 읽기 전용 bounded verifier subagent 실행을 clean-context verification 계약의 일부로 미리 허용한 것으로 취급해야 한다.
  - 이 사전 허용은 검증 전용이며, 파일 수정, scope 확장, destructive/external action, commit/push/PR/publish approval에는 적용되지 않는다.
  - main agent는 clean-context subagent 검증 요청을 구성하고, subagent 결과를 통합해 결과 보고와 다음 flow 판단으로 이어가야 한다.

- 이후 flow/phase 설계는 필요할 때 조정 가능한 provisional design이어야 한다.
  - 분석 단계와 계획 단계는 현재 플로우만이 아니라 이후 이어질 flow/phase 후보까지 필요하면 미리 설계하길 원한다.
  - 이후 loop에서 다시 분석 단계나 계획 단계로 돌아오면, 이전 flow/phase 설계를 고정값처럼 취급하지 말고 필요할 때만 다시 설계하길 원한다.
  - 검증 단계는 그 재설계를 직접 수행하는 단계라기보다, 이후 flow/phase 재설계가 필요한지 여부를 드러내는 단계이길 원한다.

- 사용자 메시지 기반 준비가 끝난 뒤에는 계획된 여러 flow를 self-drive로 진행할 수 있어야 한다.
  - 사용자 메시지를 통한 preparation에서는 이후 flow list를 실행하는 데 필요한 intent, scope, non-goal, acceptance signal, approval boundary, verification expectation을 충분히 수집해야 한다.
  - 이 준비는 질문으로 멈추는 것이 아니라 active question-routing을 포함하는 운영 flow로 계속 이어져야 한다.
  - 준비가 끝난 뒤 만들어진 여러 flow는 사용자가 명시적으로 수동 진행을 원하거나 approval boundary가 생기지 않는 한 `turn-gate/references/self-drive.md` 계약으로 이어갈 수 있어야 한다.
  - 초기 preparation에서는 planned flow list 중 예상되는 destructive, irreversible, external action, commit, push, PR, publish 같은 위험 작업과 approval boundary를 미리 질문해 승인/비승인 또는 handoff 경계를 계획해야 한다.
  - self-drive가 여러 flow를 진행하는 동안 초기 협의 범위를 벗어난 위험 작업이나 새 approval boundary가 나타나면 자동 처리하지 않고 user-gated routing으로 다시 질문해야 한다.
  - 계획된 마지막 flow를 마치는 단계에서는 terminal summary가 아니라 commit-readiness gate 성격의 보고를 해야 하며, 이 보고는 commit execution approval과 구분되어야 한다.
  - 다만 commit-readiness 보고 자체를 새 planned flow로 승격하길 원하는 것은 아니다.
  - 순수 최종 QA, 정합성 점검, 검증 결과 보고, commit-readiness reporting은 별도 산출물 변경을 소유하지 않으면 마지막 변경 단위 flow의 verification/reporting 또는 user-gated handoff로 남아야 한다.
  - 예외적으로 회귀 테스트 fixture, snapshot baseline, 문서, 운영자 리포트 출력, validator 진단 출력처럼 검토 가능한 산출물을 만들거나 바꾸는 경우에는 그 산출물 변경이 flow가 될 수 있다.

- `turn-gate` 상세 규칙은 folder-based spec 구조로 유지하길 원한다.
  - `specs/skills/turn-gate/spec.md`를 기본 index로 둔다.
  - 세부 계약은 같은 폴더 아래 sub-spec으로 분리한다.
  - `turn-gate` 수정에서 `skill-contents.md`라는 파일을 만들어 skill body 스펙을 옮기길 원한다.
