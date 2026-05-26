## 사용자 스펙 의도

- 이전 외부 플러그인 설명과 관련 명칭은 제거하고, `advance-codex` 플러그인에 `optimize-token` 스킬로 만들며, 상세 응답 기준은 `references/response.md`에 두고 싶다.
- 참고 출처는 spec에 남기고 싶다: https://github.com/JuliusBrussee/caveman/blob/main/plugins/caveman/skills/caveman/SKILL.md
- `optimize-token`은 한국어 존대체만 특수하게 고정하기보다, 활성 언어와 말투를 보존하고 이 저장소처럼 한국어 존대체가 요구되는 환경에서는 그 기준을 따르게 하고 싶다.
- `optimize-token`은 단순히 "압축"이라고만 하지 않고, 압축 강도와 문법 보존 기준을 둬서 어색한 문장이 나오지 않게 하고 싶다.
- 답변이나 작업을 하기 전에 생각하는 단계에서도 간결하게 판단해야 한다.
  - 예: `사용자가 설명을 원하는 것 같습니다. 자세한 답변을 제공한 후 구현해 볼 의향이 있는지 물어볼 수 있습니다.`를 `사용자가 설명을 원하며, 이후 구현해 볼 의향이 있는지 물어볼 수 있습니다.`처럼 줄인다.
- 예시 문장으로 테스트하기 위해 `optimize-token` 스펙을 folder-based로 두고 `intent-scenarios/thinking.md`에 thinking 예시를 둔다.
