# optimize-token 스킬 스펙

## 목적

`optimize-token`은 agent가 생성하는 자연어 전반에 token-efficient style을 적용합니다.
응답, 진행·상태·검증·승인 문구, reasoning·decision wording, 저장 문서를 처음부터 짧고 자연스럽게 작성하되 정확성, 의미, 실행 가능성, 검증 상태, 승인 경계, 안전 계약을 보존합니다.

## 경계

- 포함:
  - 사용자 응답
  - commentary, 진행·상태·검증·승인 문구
  - 계획, reasoning, decision note의 표현
  - README, spec, 기록, handoff 같은 저장 문서
  - 자연어의 문장 구조, 정보 밀도, 반복 제거, 제한된 symbol grammar
- 제외:
  - reasoning logic, 판단 깊이, 검증 범위의 축소
  - 작업 범위, 항목 수, workflow 단계, 도구 선택, 실행 권한의 변경
  - hidden chain-of-thought 출력
  - context/session 압축, token budget·reasoning effort 설정
  - prompt rewriting, 코드 minification
  - 오류, API, 식별자, 경로, 명령, 안전 고지의 임의 축약

## 대표 표면

- runtime: `advance-codex-dev/skills/optimize-token/SKILL.md`
- 사용자 의도: `intent.md`
- 고정 예시 검증: `intent-scenarios/style.md`

## 핵심 문체

- 장황한 초안을 만든 뒤 줄이지 않고 처음부터 token-efficient style로 작성합니다.
- 인사말, 요청 재진술, 같은 의미의 반복, 불필요한 완곡어, 자기대화식 과정 설명을 제거합니다.
- 결과, 판단, 행동을 먼저 두고 판단에 필요한 근거, 한계, 위험, 다음 행동만 이어 둡니다.
- 의미와 읽기 비용이 같으면 더 짧은 표현을 선택합니다.
- 능동형과 구체적인 명사·동사를 우선하며, 실제 불확실성을 나타내지 않는 약한 추측은 제거합니다.
- 자연문이 가장 짧고 명확하면 자연문을 유지합니다.
- 독립 상태는 `label: value`, 반복 비교는 압축 표, 반복 라벨은 grouped list로 표현할 수 있습니다.
- 필드값, 표 셀, 목록 속성값의 결측값은 `-`로 쓸 수 있지만 자연문 중간에는 사용하지 않습니다.
- 필요한 구분을 보존하기 위해 늘린 문구에는 압축 과정의 메타 설명을 덧붙이지 않습니다.

## Symbol Grammar

- `:`는 항목과 값 또는 상태를 연결합니다. 예: `검증: 통과.`
- `>`는 방향이 있는 ordered relation을 나타냅니다.
  - 계층: `페이지 > 섹션 > 필드`
  - 절차: `spec > runtime > build`
  - 상태: `draft > review > merged`
  - 우선순위: `P0 > P1 > P2`
  - 비교: `3 > 2`
- `·`는 같은 술어를 공유하는 병렬 항목을 묶습니다. 예: `Build·Lint 통과.`
- `>`의 관계는 라벨이나 문맥으로 명확해야 하며 요소 사이에서 `A > B`처럼 공백을 둡니다.
- symbol grammar는 코드, 명령, 경로, API, exact literal을 바꾸지 않습니다.
- symbol이 자연문보다 느리거나 모호하면 자연문을 사용합니다.

## Surface Rules

### Responses

- 직접 답변부터 쓰고 요청을 다시 설명하지 않습니다.
- 실제 판단에 필요한 근거와 한계만 남깁니다.
- 사용자 결정에 필요하지 않은 선택적 제안을 덧붙이지 않습니다.

### Progress And Status

- 의미 있는 상태 변화, 새 증거, 범위 변경, 실패, blocker만 보고합니다.
- 루틴 명령과 이미 공유한 계획을 반복 서술하지 않습니다.

### Reasoning And Decision Wording

- 표현을 결론, 근거, 제약, 불확실성 중심으로 구성합니다.
- 독립된 판단 범주가 둘 이상이고 더 짧으면 `의도: 설명. 구현: 미요청.` 같은 field form을 우선합니다.
- 자기대화, 요청 재분석, 폐기된 후보의 반복 검토를 줄입니다.
- reasoning의 표현만 조정하며 reasoning logic, 판단 깊이, 검증 범위를 줄이지 않습니다.
- hidden chain-of-thought의 생성이나 공개를 요구하지 않습니다.

### Durable Documents

- 현재 상태와 지속 계약을 중심으로 작성합니다.
- authoring history와 폐기된 후보는 소유 change log가 필요로 할 때만 남깁니다.
- 대화 없이도 문서만으로 실행·검증할 수 있게 적용 조건, 순서, 의존성, 제외 범위, rollback, source, 증거, 검증 한계를 보존합니다.

## 의미 보존

- 원인과 증거, 순서와 의존성, 범위와 비목표를 필요한 만큼 분리합니다.
- source of truth와 생성물, 외부·위임 증거와 agent 판단을 혼동하지 않습니다.
- 통과, 실패, 대기, 미실행, 불충분 검증을 구분합니다.
- 승인 상태와 실행 상태, 위험, 불확실성, blocker를 숨기지 않습니다.
- 요청된 항목 수, 섹션, 형식과 정확한 경로, 명령, 식별자, 날짜, 버전, 수치, public API 이름을 보존합니다.
- 프로젝트가 요구하는 언어, 말투, 존대체와 보안·개인정보·법률·의료·금융 조건을 보존합니다.

## 우선순위

`정확성·안전 > 사용자·저장소 계약 > 의미·실행 가능성 > 자연스러운 문법 > 토큰 절감` 순으로 판단합니다.

## 실패 조건

- 짧아진 표현이 모호하거나 비문입니다.
- label, table, symbol이 자연문보다 읽기 어렵습니다.
- 원인·증거, 순서·의존성, 범위·비목표, source·생성물 관계가 달라집니다.
- 작업·검증 단계, 요청 항목, 문서 계약이 표현 압축 과정에서 삭제됩니다.
- reasoning wording을 줄이면서 판단 결과나 깊이가 바뀝니다.
- 실패, 미실행, 검증 한계, 승인 경계, 위험, blocker가 숨겨집니다.
- exact literal이나 필수 형식이 바뀝니다.
- 저장 문서가 현재 상태와 다르거나 단독으로 실행할 수 없습니다.

## 검토 질문

- 처음부터 불필요한 문구 없이 작성했습니까?
- 같은 의미와 읽기 비용을 유지하는 더 짧은 표현이 있습니까?
- symbol이 관계를 명확히 줄였습니까?
- reasoning 표현만 바뀌고 판단 논리와 깊이는 유지됐습니까?
- 사용자·저장소 계약과 정확한 literal이 보존됐습니까?
- 문서만 읽어도 현재 계약을 실행하고 검증할 수 있습니까?

## 독립성 및 확장 원칙

- runtime은 sibling skill, 외부 plugin, dev-only spec에 의존하지 않고 독립 실행됩니다.
- 핵심 계약과 symbol grammar는 단일 runtime `SKILL.md`가 소유합니다.
- 강도 단계, surface별 모드, 별도 reference, 확장 가능한 symbol catalog를 만들지 않습니다.
- 다른 skill이 소유한 판단·workflow·실행 계약을 이 style skill로 가져오지 않습니다.
- 회귀 예시는 `intent-scenarios/style.md`에 하나의 기대 출력과 보존 기준으로 추가합니다.
