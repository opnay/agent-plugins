# flow ambiguity preparation sub-spec

## 목적

이 문서는 flow contract 형성에 영향을 주는 operation/target ambiguity를 잠그는 계약을 소유합니다.

## 계약

- 사용자 지시어의 operation 또는 target 해석에 따라 flow scope, output, verification path, handoff condition이 달라지면 ambiguity를 잠근다.
- `merge`, `absorb`, `move`, `promote`, `remove`, `delete`, `split`, `route`, `phase`, `surface`, `skill`, `spec`, `contract` 같은 표현은 여러 구조 단위를 가리킬 수 있으면 바로 실행하지 않는다.
- `그`, `그 밑`, `그건`, `위`, `아래`, `현재 것`처럼 주변 문맥의 여러 대상을 가리킬 수 있는 표현도 flow contract가 달라지면 ambiguity 대상으로 본다.
- source URL, provenance note, intent block, normative spec body가 서로 다른 target일 수 있으면 target을 잠근다.
- ambiguity resolution 결과는 interpreted operation, operation target, alternate interpretations, impact of ambiguity로 flow output 또는 record에 남길 수 있어야 한다.
- destructive, irreversible, external, commit, push, PR, publish, release, version bump 실행 승인은 이 문서가 아니라 approval boundary가 소유한다.

## 검토 질문

- 해석 후보가 flow scope나 산출물을 바꾸는가?
- 해석 후보가 approval-sensitive action 여부를 바꾸는가?
- flow contract를 잠그는 질문과 실행 승인을 혼동하지 않았는가?
