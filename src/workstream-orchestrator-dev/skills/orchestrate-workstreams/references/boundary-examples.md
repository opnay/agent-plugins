# Boundary Examples

| 요청 형태 | 자동 선택 | 판단 |
| --- | --- | --- |
| 시장 조사와 prototype 구현 | `DISPATCH` 조건부 | 두 lane이 현재 입력으로 독립 시작하고 별도 evidence·deliverable을 가지면 dispatch합니다. 구현이 조사 결론에 의존하면 graph dependency를 지키고 순차 실행합니다. |
| 대규모 schema-bound extraction | gate 적용 | 독립 partition이 두 개 이상이면 Terra xhigh `PROCESS_STRUCTURED`로 큰 coherent batch를 배정합니다. 단일 extraction이면 `DIRECT`입니다. |
| 완전히 분리된 implementation lane | `DISPATCH` 조건부 | shared contract가 고정되고 writable ownership이 disjoint이면 Terra xhigh `IMPLEMENT_OWNED`를 사용합니다. |
| 고품질 evidence 충돌, deterministic check 없음 | Sol 선택 가능 | 메인 에이전트가 직접 충돌을 확인하고, gate를 이미 통과한 lifecycle의 dependent node일 때만 Sol xhigh `FRONTIER_JUDGMENT` audit를 추가합니다. 단독 Sol dispatch는 만들지 않습니다. |
| 하나의 sequential root cause | `DIRECT` | 후속 작업이 공통 원인에 의존하므로 병렬화하지 않습니다. |
| 여러 출처의 pure research report | 자동 trigger 제외 | source methodology와 evidence-traceable report가 산출물인 요청은 이 skill이 자동 소유하지 않습니다. |
| ordinary software-engineering-only 요청 | 자동 trigger 제외 | 명시적 orchestration이 없으면 software-only precision routing을 중복하지 않습니다. |

Explicit subagent 요청은 자동 trigger 제외 사례도 gate에 진입시킬 수 있지만, plugin 경계를 넓히거나 spawn을 보장하지 않습니다. Sibling plugin은 참고 경계이며 runtime prerequisite가 아닙니다.
