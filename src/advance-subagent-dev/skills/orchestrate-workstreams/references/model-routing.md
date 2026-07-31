# Model And Role Routing

## 기본 Route

기본 spawned worker:

```yaml
model: gpt-5.6-terra
reasoning_effort: xhigh
```

Task packet의 역할로 Terra worker를 구분하세요.

| 역할 | 책임 | Access |
| --- | --- | --- |
| `EXPLORE_READ` | research, discovery, logs, source·dependency mapping | read-only |
| `IMPLEMENT_OWNED` | disjoint owned surface 안의 implementation·action | write-enabled |
| `REVIEW_LENS` | correctness, security, performance, quality, counterevidence | read-only |
| `PROCESS_STRUCTURED` | schema-bound extraction, transformation, classification, test generation, repetitive mechanical work | 주로 read-only, 산출물 ownership이 분리되면 write-enabled |

`PROCESS_STRUCTURED`는 이전의 Luna형 duty를 Terra xhigh로 통합한 runtime route입니다. Terra와 Luna의 동등성을 주장하지 않으며 Luna route를 만들지 마세요.

## Frontier Judgment

`gpt-5.6-sol`, `xhigh`의 `FRONTIER_JUDGMENT`는 다음 중 하나 이상일 때만 사용하세요.

- ambiguous goal을 framing해야 합니다.
- shared contract 또는 architecture를 설계해야 합니다.
- strong evidence가 충돌합니다.
- error cost가 높고 deterministic verification이 없습니다.
- independent final frontier-level audit가 정당화됩니다.

작업이 길거나 많다는 이유로 Sol을 사용하지 마세요. Deterministic verification이 충분하면 Terra 결과를 메인 에이전트가 직접 검증하세요.

Sol worker도 dispatch gate의 예외가 아닙니다. Gate를 이미 통과한 lifecycle의 graph에 조건부 dependent audit node로 포함하고 prerequisite 완료 뒤에만 spawn하세요. 독립 workstream이 하나뿐인 새 lifecycle에서 Sol audit 한 명만 spawn하지 말고 `DIRECT`로 판단하세요.

## Cost와 Availability

- 적은 agent와 큰 coherent batch를 사용하세요.
- Strict schema, deterministic check, bounded retry, explicit stop condition을 packet에 넣으세요.
- 기계적 volume을 agent 수로 보상하지 마세요.
- 명시 선택 모델이 unavailable이면 available model로 role contract를 보존하거나 `DIRECT`로 실행하세요.
- 대체 모델, 생략된 audit, verification limit를 최종 응답에 공개하세요.
