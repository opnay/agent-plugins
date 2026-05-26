# optimize-token 스킬 스펙

## 목적

`optimize-token`은 Codex의 응답과 작업 전 판단 문장을 짧고 선명하게 만들되 정확성, 안전 경계, 검증 보고, 사용자가 요구한 형식을 약화하지 않는 토큰 최적화 skill입니다.
목표는 특정 외부 플러그인이나 문체 이름을 재현하는 것이 아니라, 토큰 낭비를 줄이는 실용적인 작성 기준을 제공하는 것입니다.

## 경계

- 포함:
  - 불필요한 인사말, 완충 표현, 반복 설명, 과도한 배경 설명 줄이기
  - 답변 순서를 결론, 근거, 다음 행동 중심으로 정리하기
  - 작업 전 판단 문장을 느슨한 추측형 표현에서 실행 가능한 판단으로 압축하기
  - 짧게 쓰더라도 승인 경계, 검증 결과, residual risk, 파일 경로 같은 필수 정보를 유지하기
  - 활성 언어, 말투, 저장소별 출력 기대사항을 유지하며 응답 압축하기
- 제외:
  - 모델 context compression, 세션 요약, session record 압축
  - 코드, 오류 메시지, API 이름, 파일 경로의 임의 축약
  - 안전 경고, 법률/의료/금융 고지, approval-sensitive 설명의 축소
  - 특정 외부 플러그인명, 밈, 말투, 캐릭터화된 표현 사용
  - 내부 추론을 사용자에게 길게 노출하는 설명형 thinking 출력

## 처리하려는 작업 형태

- 사용자가 "짧게", "간결하게", "토큰 아끼게", "불필요한 말 빼고"처럼 응답 형식 최적화를 요청하는 경우
- 문서, 보고, 리뷰 결과, 구현 완료 보고처럼 내용은 유지하되 표현 밀도를 높여야 하는 경우
- 다른 skill이 만든 결과물을 사용자에게 전달하기 전에 응답 표면만 다듬어야 하는 경우
- 답변이나 작업을 시작하기 전에 사용자 의도, 다음 행동, 검증 필요성을 더 간결한 판단 문장으로 정리해야 하는 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `advance-codex-dev/skills/optimize-token/SKILL.md`
- intent 기록: `advance-codex-dev/specs/skills/optimize-token/intent.md`
- sub-spec:
  - `response.md`: 최종 사용자 응답 표면의 압축 계약
  - `intent-scenarios/thinking.md`: 작업 전 thinking 문장 최적화 예시 시나리오와 runtime 승격 근거
- 호출 방식: 사용자가 짧고 선명한 응답, 토큰 절약, 불필요한 말 제거, 응답 압축을 요청할 때 호출한다.

## 핵심 처리 계약

- `SKILL.md`는 언제 이 skill을 쓰는지와 어떤 runtime reference를 읽을지 안내합니다.
- 자세한 응답 작성 규칙은 runtime `references/response.md`가 소유합니다.
- 작업 전 간결한 판단 문장과 사용자에게 보이는 진행 보고를 간결하게 만드는 규칙은 runtime `references/thinking.md`가 소유합니다. 이 규칙은 내부 추론 공개를 요구하지 않습니다.
- spec의 `response.md`는 runtime response reference의 지속 계약을 소유합니다.
- `intent-scenarios/thinking.md`는 runtime thinking reference의 예시와 판정 기준을 검증하는 의도 시나리오를 소유합니다.
- 토큰 최적화는 의미 보존을 전제로 합니다. 짧아진 문장이 더 모호해지면 실패입니다.
- 다른 상위 지침이 요구하는 보고 항목은 삭제하지 않고 더 촘촘하게 씁니다.
- 사용자가 명시한 출력 형식은 유지합니다.

## 참고 출처 및 출처 사용 규칙

- 참고 출처: https://github.com/JuliusBrussee/caveman/blob/main/plugins/caveman/skills/caveman/SKILL.md
- 이 출처는 응답 압축 아이디어의 참고 근거로만 둡니다.
- runtime skill과 reference는 이 저장소의 `optimize-token` 책임에 맞게 독립 문서로 유지합니다.
- runtime skill에는 출처, 외부 플러그인명, 캐릭터화된 style label을 넣지 않습니다.

## 검토 질문

- 이 skill이 응답과 작업 전 판단 문장의 토큰 최적화만 다루고 있는가?
- 짧게 만들면서 정확성이나 검증 상태가 흐려지지 않았는가?
- 상위 지침의 언어, 말투, 보고 형식을 유지했는가?
- 압축 강도가 요청과 상황에 맞고, 문장이 어색해지지 않았는가?
- 단순한 "짧게" 요청을 과도하게 `dense`로 해석하지 않았는가?
- thinking 예시가 내부 추론 공개나 장황한 방법 설명으로 흐르지 않았는가?
- 외부 플러그인명이나 불필요한 출처 설명이 runtime에 남지 않았는가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 예.
- 그렇다면 왜 필요한가 / 아니라면 어떤 sibling context를 허용하는가: `optimize-token`은 토큰 최적화만 소유하는 작은 skill이므로, 다른 sibling skill의 존재를 전제로 하지 않고 단독으로 읽혀야 합니다.

## 확장 원칙

- 더 많은 응답 유형이 필요하면 spec `response.md`와 runtime `references/response.md`에 추가합니다.
- 작업 전 판단 최적화 예시가 늘어나면 `intent-scenarios/thinking.md`에 먼저 추가하고, 지속 규칙만 runtime `references/thinking.md`로 승격합니다.
- session record, prompt compression, subagent handoff 압축 같은 별도 책임은 이 skill에 흡수하지 않습니다.
- 출력 형식별 예시는 간결하게 유지하고, 반복되는 문체 예시를 과도하게 늘리지 않습니다.
