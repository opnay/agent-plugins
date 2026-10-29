## 사용자 스펙 의도

- skill boundary와 비책임을 더 엄격하게 잠그고 싶다.
- plugin-owned skill이 hidden sibling context에 기대지 않게 만들고 싶다.
- tool policy와 domain workflow가 섞일 때 skill boundary 안에서 분리할 기준이 필요하다.
- 스킬의 `description`은 스킬 사용 여부를 결정하는 트리거 조건이므로, passive skill은 설명 끝에 토큰 매칭에 걸릴 수 있는 plain token 목록을 쉼표로 나열해야 한다.
  - 이 목록은 literal hashtag가 아니며 `#`를 붙이지 않는다.
  - 이 방식은 사용자가 명시적으로 스킬을 요청하지 않아도 적용되어야 하는 passive skill에 사용한다.
- 스킬 테스팅할 때 유즈케이스를 만들어 스킬을 사용할 때와 안 할 때 비교하는 과정이 필요합니다.
  - 일반적인 유즈케이스와 함께, 스킬이 없을 때 실패할 만한 내용의 유즈케이스를 포함해야 합니다.
- 스킬 사용 조건은 실행 작업 폴더의 `.agents/skills`에 현재 테스트 대상 스킬 폴더를 임시 심링크하고, 에이전트가 직접 발견·선택·본문 로딩하도록 검증해야 합니다.

---

# skill-creator 스킬 스펙

## 목적

`skill-creator`는 canonical system `skill-creator` 위에 더 엄격한 skill boundary 규칙을 더하는 확장 스킬입니다.
핵심은 각 skill을 independently usable한 bounded artifact로 설계하고, plugin-owned skill인 경우 plugin usage guidance와 섞이지 않게 하는 것입니다.
일반 사례와 스킬 부재 시 실패 가능 사례의 사용/미사용 비교로 계약 충족과 스킬의 효과를 구분합니다.

## 경계

- 포함:
  - skill boundary와 비책임 명시
  - hidden sibling context 제거
  - plugin 내부 skill packaging 규칙
  - runtime-specific tool policy가 skill boundary를 흐릴 때 분리 필요성 판단
  - skill `description`의 trigger metadata 작성 규칙
  - 유즈케이스 기반 사용/미사용 비교 설계, 실행 검증과 결과 기반 개선
- 제외:
  - canonical base scaffold 전체 대체
  - plugin 사용 기준 정의
  - domain workflow 자체의 구현

## 처리하려는 작업 형태

- 새 skill을 만들거나 existing skill boundary를 재설계하는 경우
- plugin-owned skill이 sibling context에 의존하고 있는 경우
- tool-use policy가 skill boundary를 흐릴 수 있는 경우
- passive skill이 명시적 호출 없이도 필요한 상황에서 선택되도록 `description` trigger token을 설계해야 하는 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `advance-codex/skills/skill-creator/SKILL.md`
- 호출 방식: 직접 호출하거나 manifest prompt의 안내를 따른다.

## 핵심 처리 계약

- 각 skill은 목적, 소유 범위, 비목표를 명확히 가져야 한다.
- plugin 내부 skill이라도 cross-skill usage guidance는 manifest prompt, README, plugin spec에 둔다.
- main bottleneck이 tool policy면 domain skill 안에 넣지 않고 해당 plugin 또는 system tool guidance가 소유하게 한다.
- behavioral guidance는 skill reader 관점에서 작성한다.
- skill frontmatter `description`은 본문을 읽기 전에 사용되는 trigger metadata로 취급한다.
- `description`은 먼저 사람이 읽을 수 있는 기본 설명과 사용 조건을 짧게 쓰고, passive skill인 경우 마지막에 token matching을 돕는 plain token 목록을 쉼표로 나열한다.
- passive skill token 목록은 `#`가 붙은 hashtag가 아니라 `skill creation, skill update, passive trigger, description metadata` 같은 쉼표 구분 plain token 형식이어야 한다.
- 이 token 목록은 사용자가 명시적으로 skill을 호출하지 않아도 적용되어야 하는 passive skill에만 사용한다.
- 사용자가 명시적으로 호출하는 active skill이나 이미 충분히 명확한 narrow trigger skill에는 불필요한 token 목록을 붙이지 않는다.
- token 목록은 skill이 실제로 적용되어야 하는 입력 표현, artifact 이름, workflow 이름, 흔한 동의어를 포함하되 unrelated broad token으로 과도하게 넓히지 않는다.
- 신규 skill, 행동·경계·trigger·도구 절차의 실질 변경, 사용자가 요청한 skill 테스트는 아래 비교 검증 계약을 따릅니다.
- 표현만 바뀌는 수정은 기존 사례 영향 검토로 검증 범위를 정할 수 있습니다. 설계·검토 요청은 검증 계획까지 작성하며 실행 권한을 뜻하지 않습니다.

## 비교 행동 검증

1. 일반 사례와 skill 부재 시 실패할 가능성이 있는 사례를 각각 1개 이상 준비합니다. 최소 범주 충족과 충분한 검증을 구분하며, 사례·반복 수는 변경 계약과 위험에 맞춥니다.
2. 실행 전에 사례 ID, 유형, 검증 계약, 사용자 요청, 초기 fixture, 허용·금지 작업, 관찰 가능한 성공 기준을 고정합니다. 실패 가능 사례에는 미사용 실패 가설과 이유도 적습니다. 가설은 관찰 사실이 아닙니다.
3. 동일 모델·설정·도구·권한·요청·초기 상태·실행 예산으로 독립적인 새 실행 두 개를 준비합니다. 아래 실제 로딩 계약을 따르며, 다른 필수 지침·공통 리소스는 유지합니다. 비교 조건을 맞출 수 없거나 로딩을 확인할 수 없으면 비교 한계를 명시하고 유효한 대조 실험으로 단정하지 않습니다.
4. 실행자에게는 작업 입력, fixture, 허용 범위와 근거 출력 형식만 제공합니다. 기대 정답·실패 가설·다른 실행 결과는 숨깁니다. 상태 변경은 매번 초기화한 격리 fixture에서 수행합니다. 실서비스 변경 등 추가 권한이 필요한 실행은 모의화하거나 미실행 상태로 둡니다.
5. 평가자는 사전에 고정한 동일 기준으로 실제 산출물과 도구 기록을 판정합니다. 사용 실패·권한 위반은 미사용 결과와 무관하게 기록합니다.
6. 지침 문제로 확인된 계약 위반·퇴행은 수정 후 원래 사례와 영향받는 일반 사례를 재검증합니다. 결과에 맞추려고 성공 기준을 바꾸지 않습니다. fixture·판정 기준이 잘못되었다면 수정 이유와 새 버전을 기록하고 양쪽을 다시 실행합니다.

### 실제 스킬 로딩

- 대상은 사용자가 지정한 현재 skill 폴더입니다. 수정 중인 skill은 저장소 작업 파일을 기준으로 하며 설치 캐시로 대체하지 않습니다. 실행 전에 원본 경로·내용 해시를 고정하고 실행 중 변경하지 않습니다.
- Codex 실행은 조건마다 별도 새 프로세스·격리 작업 폴더·신규 `HOME`과 `CODEX_HOME`·빈 사용자 설정을 사용합니다. 필요한 인증은 기존 권한 안에서만 제공하고, 공통 환경·도구·canonical skill 등 필수 지침은 동일하게 유지합니다. subagent의 context 초기화만으로 이 격리를 대신하지 않습니다.
- 사용 조건의 작업 폴더에 `.agents/skills/<skill-folder>` 임시 심링크를 만들어 고정한 대상 폴더와 전용 리소스를 연결합니다. native skill discovery를 활성화하고 본문을 prompt에 직접 넣거나 대신 읽어 주지 않습니다. 허용된 discovery·본문 읽기를 실행자 권한에서 차단하지 않습니다.
- 미사용 조건에는 대상 심링크를 만들지 않으며, 상위 폴더·사용자·관리자·plugin·동명 skill 등 다른 경로로 대상이 노출되지 않는지 확인합니다. 공통 canonical skill과 대상 extension은 실제 경로로 구분합니다.
- 발견 목록의 이름·description 노출, 실제 선택, 대상 `SKILL.md` 및 필요한 리소스 읽기를 구분해 기록합니다. 심링크가 있다는 사실만으로 사용을 인정하지 않습니다. 자연어 자동 선택 사례는 호출 힌트 없이 실행하고, 명시 호출 사례와 별도로 판정합니다. 미선택도 관찰 결과로 남기며 수동 본문 주입으로 성공 처리하지 않습니다.
- 종료 후 이번 실행에서 만든 임시 심링크만 제거하고 원본 폴더·기존 설정은 보존합니다. 심링크·native 로딩이 지원되지 않거나 실제 읽기 근거가 없으면 실제 로딩 검증 미완료로 보고합니다.

### 결과와 한계

| 미사용 | 사용 | 해석 |
|---|---|---|
| 실패 | 성공 | 해당 사례의 개선 근거 |
| 성공 | 성공 | 계약 충족, 개선 효과 미확인 |
| 실패 | 실패 | 해결 효과 미확인 |
| 성공 | 실패 | 퇴행 후보, 원인 조사 필요 |

- 미사용 성공을 그대로 보존합니다. 미사용 실패를 얻기 위한 사후 사례 선택·기준 변경을 하지 않습니다.
- 환경 오류, fixture 오류, 관찰 부족, skill 로딩 오염, 미실행은 성공/실패와 분리합니다. 무효 사례는 효과 판정에서 제외하고 이유를 남깁니다.
- 테스트 완료와 효과 입증은 다릅니다. 단회 결과는 관측한 사례의 근거이며 안정성이나 일반 성능 향상으로 확대하지 않습니다. 중대한 권한 위반을 평균 성공률로 상쇄하지 않습니다.
- 본문 직접 주입 실험은 로딩 후 행동의 보조 근거일 뿐, 심링크 기반 발견·선택·실제 로딩 비교를 대체하지 않습니다. description을 바꾸면 일반 사용자 요청의 자동 선택과 비대상 요청의 오선택을 별도로 확인하며, 명시 호출 결과로 자동 선택 검증을 대체하지 않습니다.
- 독립 실행이나 격리가 불가능하면 구조 검토만 수행하고 빠진 행동 검증을 보고합니다. 외부 전송이나 추가 비용·권한은 기존 승인 범위를 따릅니다.
- 비교 계약은 이 skill이 소유합니다. `$jevx:scenario-testing` 같은 평가 수단은 선택 사항이며, 설치·숨은 sibling context·고정 대규모 사례 수를 필수로 요구하지 않습니다.

## 출력과 완료 조건

- 기존 출력: skill boundary, 비목표, plugin ownership, description trigger 동작, 분리·회피한 runtime tool policy, 검증 경로, 잔여 위험.
- 비교 출력: 대상 skill·버전·원본 경로·내용 해시, 고정 사례·성공 기준, 임시 심링크와 격리 설정, 발견·선택·본문 읽기 및 baseline 비노출 근거, 사용/미사용 결과와 근거 위치, 정리 결과, 무효·미실행 목록, 계약 위반·퇴행, 수정·재검증 결과와 한계.
- 설계와 실제 실행 결과를 구분합니다. 유효한 두 사례군·양쪽 실행과 관찰 근거가 있어야 비교 테스트 완료로 보고합니다. 미충족 항목은 미완료 또는 검증 한계로 명시합니다.

## 검토 질문

- 이 skill이 hidden sibling context 없이도 이해 가능한 경계를 갖고 있는가?
- cross-skill usage guidance가 개별 skill 안으로 잘못 들어오지 않았는가?
- tool policy를 domain skill 안에 묻어 두어 skill boundary를 흐리지 않는가?
- `description`이 body를 읽기 전 trigger metadata로 충분히 작동하는가?
- passive skill이라면 `description` 끝에 `#` 없는 쉼표 구분 plain token 목록이 있는가?
- token 목록이 실제 적용 조건을 넓히되, 무관한 broad match를 만들지 않는가?
- 일반 사례와 미사용 실패 가설 사례를 모두 고정하고 독립 실행으로 비교했는가?
- 미사용의 실제 성공, 사용 실패, 비교 무효·미실행과 권한 경계를 빠짐없이 기록했는가?
- 효과 입증과 테스트 완료, 본문 적용과 trigger 선택 검증을 구분했는가?
- 현재 대상 폴더의 임시 심링크와 실제 읽기를 확인하고, 본문 직접 주입이나 설치 캐시로 실제 로딩 비교를 대신하지 않았는가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 예.
- 그렇다면 왜 필요한가 / 아니라면 어떤 sibling context를 허용하는가: 이 스킬 자체가 hidden sibling context 제거를 다루므로, skill boundary guardrail은 독립적으로 읽히게 강제하는 편이 맞다.

## 확장 원칙

- skill packaging 규칙이 바뀌면 `plugin-creator`, plugin spec, manifest prompt도 함께 점검한다.
- 새 규칙은 reusable skill quality를 높이는 경우에만 추가한다.
