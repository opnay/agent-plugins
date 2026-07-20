## 사용자 스펙 의도

- adaptive subagent 플러그인의 스킬을 `경량 진입 → 위임·생성 → 결과 검증·통합`으로 세분화합니다.
  - 결과 검증·통합 단계는 모든 필수 결과 대기, evidence 검증, 충돌 해결, 전체 통합을 소유합니다.
  - 후속 스킬은 implicit 호출하지 않으며 유효한 dispatch 입력이 있을 때 직접 호출할 수 있습니다.

---

# integrate-subagent-results 스킬 스펙

## 목적

`integrate-subagent-results`는 완전한 `DispatchManifest`와 active subagent results를 입력으로 모든 필수 결과를 검증하고 main agent의 최종 통합을 안내합니다.

## 경계

- 포함:
  - 필수 결과 대기와 상태 정규화
  - evidence 검증, 중복 병합, 충돌 판정
  - 제한적 follow-up과 실패 복구
  - 변경 통합, whole-result verification, final response guidance
- 제외:
  - implicit invocation
  - 초기 workstream 분해와 execution mode 결정
  - DispatchManifest 없는 새 spawn
  - read 권한의 write 권한 자동 승격
  - subagent 결론의 무검증 수용

## 처리하려는 작업 형태

- 여러 explorer의 조사 결과 통합
- 여러 worker의 disjoint 변경 검토와 통합
- completed, blocked, inconclusive 결과 혼합 처리
- 주장 또는 변경이 충돌하는 결과 검증
- PARALLEL_READ 뒤 새로운 write scope가 필요한 전환

## 엔트리포인트 / 대표 표면

- 대표 표면: `skills/integrate-subagent-results/SKILL.md`
- 결과 계약: `references/result-contract.md`
- 통합 예시: `references/examples.md`
- 호출 방식: `$adaptive-subagent-orchestrator-dev:integrate-subagent-results` 또는 완전한 `DispatchManifest` handoff
- implicit policy: `allow_implicit_invocation: false`

## 입력 게이트

- 완전한 `DispatchManifest`와 현재 session에서 식별 가능한 required agent를 요구합니다.
- manifest에는 mode, assignments, ownership, required results, main-owned work, whole-result verification, follow-up 사용 상태가 있어야 합니다.
- 입력이 없거나 agent IDs를 확인할 수 없으면 spawn하거나 추정하지 않고 missing input을 보고합니다.

## 결과 수집 계약

- 모든 required agent의 terminal result를 기다립니다.
- 각 결과를 `completed`, `blocked`, `inconclusive`로 분류합니다.
- 결과는 summary, claims, evidence, files inspected/changed, validation, risks, recommended action을 포함합니다.
- raw transcript, 긴 로그, hidden reasoning은 통합 입력으로 요구하지 않습니다.

## Evidence 검증

- subagent 결론은 evidence이며 최종 사실이 아닙니다.
- unsupported claim을 거부하고 중요한 주장은 code, test, log, diff, 직접 실행 결과로 확인합니다.
- 중복 주장은 병합하고 상충 주장은 main agent가 직접 검증하여 판정합니다.
- write result는 ownership 위반, shared file 변경, public contract drift를 확인합니다.

## 실패와 Follow-up

- 중요한 범위가 빠졌고 `follow_up_used`가 `false`일 때만 기존 agent에 좁은 follow-up을 허용하고 즉시 `true`로 기록합니다.
- agent 하나가 실패해도 전체 작업을 재시작하지 않습니다.
- 누락 범위는 main agent가 직접 확인하거나 더 좁은 재검증으로 보완하고, 확인하지 못하면 unverified scope로 보고합니다.
- 새 독립 작업이나 `PARALLEL_READ`에서 write scope가 필요하면 기존 manifest가 write를 main-owned로 표시했더라도 explicit scope로 `dispatch-subagents`를 다시 적용합니다. Dispatcher가 새 gate에서 `DIRECT`를 반환할 수 있습니다.

## 통합과 최종 응답

- main agent가 변경을 통합하고 `whole_result_verification`을 실행합니다.
- 최종 응답은 agent 수와 역할, 통합 결론, 변경 파일, 검증 결과, residual risk를 포함합니다.
- spawn이 없었던 DIRECT 작업에는 orchestration 보고 형식을 강제하지 않습니다.
- 완료를 증명하지 못한 범위는 성공으로 보고하지 않습니다.

## 검토 질문

- 모든 required agent가 terminal state인가?
- 각 claim에 검증 가능한 evidence가 있는가?
- ownership 위반과 충돌을 직접 확인했는가?
- 새 write scope가 dispatcher를 다시 거쳤는가?
- whole-result verification와 residual risk가 남김없이 보고됐는가?

## 독립성 원칙

- 이 skill은 완전한 `DispatchManifest`와 active results가 있을 때 직접 호출할 수 있습니다.
- 입력 게이트는 sibling 호출 이력을 요구하지 않고 현재 session의 명시적 artifact만 요구합니다.

## 확장 원칙

- 새 result field는 main agent의 판단이나 검증을 실제로 바꿀 때만 추가합니다.
- 초기 dispatch 판단과 agent 생성 규칙을 integrator에 추가하지 않습니다.
- result envelope와 사례는 runtime references가 소유합니다.
