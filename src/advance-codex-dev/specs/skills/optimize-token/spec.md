# optimize-token 스킬 스펙

## 목적

`optimize-token`은 에이전트가 쓰는 토큰을 줄이기 위해 출력·진행 보고·판단 문구·상태 문구를 짧고 선명하게 만들되 정확성, 안전 경계, 검증 보고, 요청 형식, 활성 언어와 말투를 보존하는 skill입니다.

## 경계

- 포함: 응답 압축, 작업 전 판단 문장 압축, 진행 보고 압축, 상태/검증/승인 문구 압축, 단계별 압축 강도 선택
- 제외: context/session 압축, prompt rewriting, 코드 minification, 오류/API/경로/명령어 임의 축약, 안전 고지 축소

## 대표 표면

- runtime: `advance-codex-dev/skills/optimize-token/SKILL.md`
- intent: `intent.md`
- writing 계약: `writing.md`
- 단계 gate: `levels/light.md`, `levels/standard.md`, `levels/extreme.md`
- 예시 검증: `intent-scenarios/writing.md`

## 핵심 계약

- 압축 강도는 `` `light` > `standard` > `extreme` `` 순서로 읽습니다.
- 뒤 단계는 앞 단계 규칙을 상속하고 자기 덮어쓰기 규칙만 추가합니다.
- `levels/`는 `writing`을 제외한 에이전트 출력과 작업 운영 문구 전반에 기본 적용됩니다.
- `writing.md`는 파일이나 문서에 저장되는 작성물의 별도 보존 계약만 소유합니다.
- `writing`은 전역 `levels/`의 기본 적용 대상이 아니며, `writing.md` 기준으로 별도 관리합니다.
- `levels/`는 각 단계의 진입, 보존, 덮어쓰기, 하향/상향 gate를 소유합니다.
- `intent-scenarios/`는 같은 입력에서 단계별 출력 차이가 유지되는지 확인하는 예시를 소유합니다.
- runtime은 설치 후 접근 가능한 `SKILL.md`와 `references/*.md`에 필요한 요약만 포함합니다.

## 확장 원칙

- 저장되는 작성물 규칙은 `writing.md`에 추가합니다.
- 단계별 gate 변경은 해당 `levels/<level>.md`에 추가합니다.
- writing 예시는 `intent-scenarios/writing.md`에 추가합니다.
- runtime에는 dev-only spec 경로를 실행 지시로 남기지 않습니다.
- 참고 출처와 사용자 의도 기록은 `intent.md`가 소유합니다.
