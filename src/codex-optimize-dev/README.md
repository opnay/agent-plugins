# Codex Optimize

`codex-optimize`는 에이전트 문체와 기술 지시의 명확성을 개선하는 플러그인입니다.
토큰 효율적인 문체와 ASD-STE100 기반 기술 지시 작성·검토 가이드를 제공합니다.
`.agents/sessions/{YYYYMMDD}`는 session-scoped operational artifact를 두는 기본 위치로만 문서화합니다.

이 플러그인은 의미와 조건을 보존하면서 에이전트가 작성하는 문장을 개선합니다.
반대로 일반적인 실행 workflow나 무관한 공용 유틸리티를 담는 용도로 넓히지 않습니다.

`optimize-token`은 에이전트 응답, 진행·상태 문구, reasoning·decision wording, 저장 문서에 token-efficient style을 적용하되 정확성, 의미, 검증, 승인, 필수 형식, exact literal, 언어, 안전 계약을 보존할 때 사용합니다.

`asd-ste100`은 Codex 지침, skill 본문, 기술 문서와 운영 절차의 대상·행동·조건·순서를 명확히 작성하거나 검토할 때 사용합니다. 본문은 언어 공통 원칙을, [영어 reference](../../codex-optimize/skills/asd-ste100/references/english-vocabulary.md)는 승인 어휘·의미·품사, 기술 용어, 영어 용법과 예시를 소유합니다. [준수 reference](../../codex-optimize/skills/asd-ste100/references/standard-scope.md)는 공식 출처·판본·검증 한계를 설명합니다.

호출 예: `$codex-optimize:asd-ste100 이 운영 절차의 의미와 조건을 유지하면서 명확하게 다듬어 주세요.`

표현 길이·반복을 줄일 때는 `optimize-token`, 기술 지시의 명확성을 개선할 때는 `asd-ste100`을 선택합니다. 함께 적용할 수 있으며 의미·조건·명확성을 압축보다 우선합니다. 영어 전용 규칙은 영어 산문에만 적용하며, 일반 원칙 적용은 공식 STE 준수 확인이 아닙니다. 공식 준수 검토에는 해당 판본의 전체 규칙·사전과 프로젝트 용어 근거가 필요합니다.

처리 계약은 [skill spec](specs/skills/asd-ste100.md), 플러그인 사용 기준은 [plugin spec](specs/plugin.md)을 따릅니다.

기존 `advance-codex`는 `codex-optimize`로 이름이 바뀌었습니다. 호출은 `$codex-optimize:optimize-token`, `$codex-optimize:asd-ste100`을 사용해 주세요. 기존 설치와 사용자 지침은 자동 변경되지 않으므로 마켓을 갱신하고 새 이름의 플러그인을 선택해 주세요.
