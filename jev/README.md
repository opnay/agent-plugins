# Jev

TypeSafe의 hosted Jev 모델을 Noul·Choice·Score 단건 또는 batch로 호출하고 구조화 입력·장문 모드를 지원하는 Go CLI와 작업 판단·reusable instruction 평가 스킬입니다.

- `$jev:jev`: Jev·CLI 사용과 관리, 질문 구성·batch·결과 해석의 기본 패턴, 작은 선택·되묻기·불확실한 근거 평가·모델 선택의 내장 유즈케이스.
- `$jev:scenario-testing`: 고정 scenario에서 expected behavior, 선택적 Jev prediction, fresh executor의 actual behavior를 비교하고 evidence를 기록.
- `jev`: `noul`, `choice`, `score`, `batch`, 결과 기준, 출력 제어, `config`, `install`, `uninstall`, `doctor`.
- 질문·문맥·기준은 TypeSafe API로 전송됩니다.

## 설치

Go 1.23 이상이 필요합니다. 플러그인 루트에서 실행합니다.

```sh
cd scripts
go run . install
```

기본 설치 위치는 `~/.local/bin/jev`입니다. 다른 위치는 `go run . install --dir /absolute/bin`, 소스 갱신 후 교체는 `go run . install --force`를 사용합니다. CLI는 셸 설정을 변경하지 않습니다.

```sh
jev install
jev install --force --dir /absolute/bin
jev uninstall
jev doctor [--provider jev]
```

- `install`: 실행 중인 바이너리를 0755 권한으로 복사합니다. 다른 일반 파일 교체에는 `--force`가 필요합니다.
- `uninstall`: Go build info로 Jev임을 식별한 일반 파일만 제거합니다. 설정·토큰·디렉터리·셸 설정은 유지합니다.
- `doctor`: Local setup·Config·Provider·Authentication으로 설치·PATH, 설정·권한, 선택 provider·어댑터 지원·모델·timeout, 키 유무·출처·API 인증을 구분합니다. 미구현 provider는 정상 설정의 Config 실패로 표시하지 않습니다. 환경변수 키 형식 오류는 Authentication 실패로 표시하고 API 점검을 건너뜁니다. 구현된 provider의 설정·키가 유효하면 어댑터로 인증을 한 번 점검하며 Jev는 `GET /v1/models`를 호출합니다. 추론 요청은 하지 않습니다.
- 결과는 `OK / FAIL / SKIP`, 문제별 해결 안내, 전체 요약으로 표시합니다. 키 원문은 숨기며, 선택 provider의 환경변수가 파일 키를 덮어쓰는지도 보여줍니다. 401·403과 네트워크·시간 초과·호출 제한·서버 오류를 구분합니다. API 조회 성공은 모델 정확성이나 추론 가능 여부를 증명하지 않습니다. 설정·설치·셸을 자동 수정하거나 요청을 재시도하지 않습니다.
- doctor 정상 결과는 stdout·종료 코드 0, 문제 결과는 빈 stdout·stderr·종료 코드 1입니다.

## 인증과 설정

설정파일은 `~/.agents/jev.toml`입니다. 파일이 없으면 내장 기본값으로 동작하며 조회·평가·doctor는 파일을 만들지 않습니다.

```sh
jev config init
jev config
jev config path
jev config set defaults.provider jev
jev config set defaults.threshold 0.8
jev config set defaults.min_score 1.5
jev config set defaults.timeout 45s
jev config set defaults.json true
jev config set defaults.pick answer,model
jev config set providers.jev.model jev-latest
jev config set providers.jev.api_key --stdin
jev config unset defaults.threshold
```

`config init`은 내장 템플릿을 0600 권한으로 생성합니다. 기존 파일은 유효성에 관계없이 덮어쓰지 않습니다. 최초 `config set`도 같은 템플릿에 요청값을 반영합니다. 미구현 provider 예시는 주석으로 제공하며 토큰·활성 placeholder는 포함하지 않습니다.

```toml
[defaults]
provider = "jev"
timeout = "30s"
json = false
pick = []

# threshold = 0.8
# min_score = 1.5

[providers.jev]
model = "jev-latest"
# 인증: TYPESAFE_API_KEY 또는 providers.jev.api_key
```

| 설정 | 내장 기본값 | 적용 |
|---|---|---|
| `defaults.provider` | `jev` | 단건·batch·doctor |
| `defaults.timeout` | `30s` | 추론·인증 조회 |
| `defaults.threshold` | 미적용 | Choice 선택 확률·Noul 참일 확률 |
| `defaults.min_score` | 미적용 | Score 점수 |
| `defaults.json` | `false` | 단건 출력 |
| `defaults.pick` | 제한 없음 | 단건 JSON 필드 |
| `providers.<name>.model` | Jev는 `jev-latest` | 해당 provider 모델 |
| `providers.<name>.api_key` | 없음 | 해당 provider 인증 |

Provider 선택은 `--provider` > defaults.provider > 내장 jev, 모델은 `--model` > 선택 provider의 model > 해당 provider 기본값입니다. 공통 설정은 명시적 flag > defaults > 내장값 순서입니다. 명시적 0·false·빈 pick도 저장값보다 우선합니다.

인증은 선택 provider의 비어 있지 않은 환경변수가 파일 키보다 우선합니다. Jev는 `TYPESAFE_API_KEY`, OpenRouter는 `OPENROUTER_API_KEY`, OpenAI는 `OPENAI_API_KEY`를 사용합니다. 다른 provider의 키·모델로 대체하지 않습니다. 현재 실행은 Jev만 지원하며 OpenRouter·OpenAI 설정은 저장·조회만 가능합니다. 미구현 provider 실행은 API 호출 전에 오류로 끝납니다.

모든 provider 키는 조회 시 전체 마스킹합니다. 토큰은 stdin으로만 저장하며 파일은 0600 평문입니다. set/unset은 오프라인이고 성공 시 무출력입니다. 알 수 없는 key·잘못된 타입과 symlink·디렉터리는 거부합니다.

기존 최상위 공통 key는 `[defaults]`, `model`·`api_key`는 `[providers.jev]`로 옮겨야 합니다. 구형·신형 혼용이나 `[default]` 단수 섹션은 오류이며 자동 migration하지 않습니다. 기존 파일 편집 시 주석·서식 보존은 보장하지 않습니다.

## Primitive 선택

| 명령 | 질문 형태 | 기본 출력 |
|---|---|---|
| `noul` | 명제가 참인가 | 0~1의 참일 확률 |
| `choice` | 선택지 중 무엇인가 | 선택된 라벨 |
| `score` | 순서형 단계에서 어느 정도인가 | 0~`단계 수 - 1`의 가중 점수 |

```sh
jev noul --if "이 문장은 설정 사용을 의무화하는가? 문장: ..."

jev choice --if "주장과 근거의 관계는? ..." \
  --conditions supported,conflicting,undetermined

jev score --if "이 장애의 심각도는? ..." \
  --level "기능 영향 없음" \
  --level "기능 저하, 우회 가능" \
  --level "작업 불가, 우회 불가능"
```

`--if`에는 질문과 판정에 필요한 문맥을 함께 넣습니다. Choice는 쉼표로 구분한 2~255개 선택지를 사용하며 라벨 내부 쉼표는 지원하지 않습니다. 라벨은 자유 문자열이지만 자동화 식별자에는 `snake_case`를 권장합니다. Score는 2~10개의 고유한 `--level`을 낮은 단계부터 반복하며 자연어와 쉼표를 허용합니다.

기존 `jev --if ... --conditions ...`는 `jev choice`와 동일하게 동작합니다.

## 구조화 입력

단건에서는 `--input`으로 JSON 파일을 전달합니다. 기존 `--if`, `--conditions`, `--level`과 함께 쓰지 않습니다. 출력·threshold·min-score 옵션은 그대로 사용할 수 있습니다.

```sh
jev choice --input request.json --json
```

```json
{
  "state": {"message": "결제 내역이 중복되었습니다.", "account": {"plan": "team"}},
  "instructions": {"question": "어느 팀이 처리해야 하나요?", "focus": "message의 주된 요청"},
  "criteria": {
    "billing": {"description": "결제·청구·환불", "examples": ["중복 청구"]},
    "support": {"description": "제품 사용·기술 지원"}
  }
}
```

- `state`: 필수 문자열·객체·배열입니다. JSON 구조와 숫자 표기를 보존합니다.
- `instructions`: 선택 문자열·객체·배열·null이며 생략하면 null입니다.
- Choice `criteria`: 2~255개 label과 각 label의 설명 문자열·객체·배열·null입니다.
- Score `criteria`: 낮은 단계부터 2~10개 고유한 문자열·객체·배열입니다. 예: `[{"impact":"영향 없음"},{"impact":"우회 가능"},{"impact":"작업 불가"}]`. 결과 JSON의 `legend`에도 해당 구조가 유지됩니다.
- Noul `criteria`: 선택 `true`/`false` 설명 object 또는 null입니다. 예: `{"true":{"meaning":"명시적으로 의무화함"},"false":{"meaning":"권고하거나 언급하지 않음"}}`.

빈 문자열·잘못된 타입·미지원 key·UTF-8/JSON 오류, 비유한 수·float64 범위 밖의 수, state·instructions·criteria 각각 64단계를 넘는 중첩은 호출 전에 거부합니다. 입력 원문은 오류 메시지에 넣지 않습니다.

## 장문 모드

```sh
jev choice --input request.json --long-context
jev batch --state-file state.json --state-format json --input questions.jsonl --long-context
```

`--long-context`는 해당 요청에만 `TypeSafe-Danger-Zone-Long-Context: true`를 전달합니다. 기본값은 off이며 `--long-context=false`도 사용할 수 있습니다. 설정파일에 저장하거나 자동 활성화하지 않습니다.

[공식 모델 문서](https://docs.typesafe.ai/models.md#long-context)의 기본 한도는 요청 전체 64k, state와 질문 하나의 합계 32k 토큰입니다. 장문 모드는 두 한도를 각각 262k로 확장하지만 속도가 느려지고 정확도가 낮아질 수 있습니다. 토큰당 가격은 동일하며 입력이 늘면 총비용도 늘어납니다. 한도 정보는 2026-10-19 확인 기준입니다.

## 작업 중 Jev 사용

`$jev:jev`는 CLI 안내와 함께 Jev를 사용하는 기본 패턴을 제공합니다. 작은 선택이 실제 다음 행동을 바꿀 때 일반 작업에서도 적용할 수 있습니다. [내장 유즈케이스](skills/jev/references/use-cases.md)는 공통 패턴의 적용 예시입니다. 먼저 요청·규칙·확인 가능한 증거를 봅니다. 답이 이미 정해졌거나 사용자의 취향·승인·권한이 필요하면 Jev가 이를 대신 결정하지 않습니다.

- 사용자의 A/B 질문: 목표·제약을 담아 `option_a`, `option_b`, `need_user_input` 중 평가하고, 근거를 확인해 추천합니다.
- 사용자에게 되묻기 전: `proceed_with_known_default`, `inspect_more`, `ask_user` 중 필요한 행동을 평가합니다.
- 근거가 불확실할 때: 동일한 출처·발췌를 공유하는 여러 넓은 질문(전체 관계·대안 해석)과 좁은 질문(개별 주장·반례)을 batch로 평가합니다. CLI에는 고정 질문 개수 한도가 없습니다. [현재 모델 문서](https://docs.typesafe.ai/models)의 요청 전체 64k 토큰, state와 가장 긴 질문 합계 32k 토큰 한도 안에서 필요한 만큼 묻습니다. 상충하는 답은 원문을 다시 조사하며, 반복 질문을 다수결 증거로 쓰지 않습니다.
- 모델 선택 예시: 기본 범위·조사량·실행량·실패 영향과 위임 경계를 먼저 확인합니다. 프로필 정책이면 허용된 모델·effort 쌍을 함께 평가하고, 개별 선택이 허용되면 모델을 먼저 평가한 뒤 지원·허용 effort를 후속 질문으로 평가합니다. 현재 정책이 답을 정하면 그대로 사용합니다. 작업 분해·subagent 실행·결과 통합과 모델 정책은 해당 작업 흐름이 소유합니다.

Jev 결과는 권고입니다. 현재 지시와 실행 권한을 확인한 뒤 행동하며, CLI·인증·API 실패나 결과 기준 미달 시 설정을 자동 변경하거나 같은 호출을 반복하지 않습니다. 실제 Jev 응답과 자체 판단으로 진행한 경우를 구분해 보고합니다.

## Batch

같은 state에 여러 질문을 할 때 한 SystemOne 요청으로 묶습니다.

```sh
jev batch --state-file target.md --input questions.jsonl
```

```jsonl
{"id":"S001","type":"choice","question":"Which action applies?","conditions":["proceed","block","unexpected_action"]}
{"id":"S002","type":"noul","question":"Is this behavior explicitly authorized?"}
{"id":"S003","type":"score","question":"How severe is the risk?","levels":["none","recoverable","destructive"]}
```

기존 `question` 대신 `instructions` 문자열·객체·배열·null을 사용할 수 있습니다. `criteria`는 단건과 같은 형태이며 기존 `conditions`·`levels` 대신 사용합니다. `question`과 `instructions`, `criteria`와 `conditions`/`levels`는 함께 사용할 수 없습니다. 기존 `question`과 새 `criteria`의 조합도 지원합니다.

```jsonl
{"id":"S004","type":"choice","instructions":{"question":"어느 팀이 처리해야 하나요?"},"criteria":{"billing":{"description":"결제·환불"},"support":{"description":"기술 지원"}}}
{"id":"S005","type":"noul","instructions":["state에서 긴급성이 드러나나요?"],"criteria":{"true":"즉시 대응 필요","false":"일반 문의"}}
```

공통 state 파일은 기본 `--state-format text`일 때 원문 문자열로 전송합니다. 객체·배열 등 JSON 구조를 전달하려면 `--state-format json`을 명시합니다. JSON처럼 보이는 text도 자동 변환하지 않습니다.

state와 모든 JSONL row를 먼저 검증한 뒤 한 번 호출합니다. 성공하면 입력 순서대로 전체 결과 JSONL을 출력합니다. 자동 chunk·retry·resume·부분 성공은 제공하지 않습니다. Batch에는 threshold, min_score, json, pick 옵션이나 저장값을 적용하지 않습니다. stdout 쓰기 전 실패는 빈 stdout과 종료 코드 1이며, 출력 장치의 부분 쓰기는 되돌릴 수 없습니다.

## Scenario testing

`$jev:scenario-testing`은 skill, command, task prompt 같은 reusable instruction이 실제로 유도하는 행동을 평가합니다. 시나리오별 행동 Choice와 hidden expected choice를 먼저 고정하고, 허용된 경우 Jev prediction을 받은 뒤 fresh executor의 관찰 행동을 같은 Choice로 분류합니다. verdict는 Jev가 아니라 caller가 actual behavior와 critical evidence에서 결정합니다.

사용자가 수를 지정하지 않은 일반 평가는 40개 scenario를 사용하며, 명시적 smoke check는 2~3개를 사용할 수 있습니다. 공통 target을 공유하는 Jev prediction은 `jev batch`로 요청하고, executor는 stable scenario ID를 유지한 별도 batch로 실행합니다. Scenario catalog, 정제된 Jev 호출, executor evidence, expected·predicted·actual comparison은 하나의 durable evidence set으로 남깁니다.

Jev 예측은 선택 사항입니다. 외부 전송 권한이 없거나 CLI를 사용할 수 없어도 executor 평가는 계속하며, prediction skip 또는 error를 evidence에 기록합니다.

## 출력과 결과 기준

- Noul JSON: `answer`, `model`.
- Choice JSON: `answer`, `probabilities`, `confidence`, `model`.
- Score JSON: `answer`, `legend`, `probabilities`, `confidence`, `model`.
- `--json --pick answer,probabilities`: 해당 primitive가 제공하는 최상위 field만 객체로 출력합니다.

CLI·설정·JSON이 소유하는 key는 `snake_case`입니다. Noul answer는 참일 확률입니다. Score answer는 단계별 확률의 가중값이라 단계 사이에 놓일 수 있습니다. `confidence`는 분포에서 계산한 별도 값이며 결과 기준에 사용하지 않습니다.

```sh
jev noul --if "..." --threshold 0.8
jev choice --if "..." --conditions a,b,c --threshold 0.8
jev score --if "..." --level low --level medium --level high --min-score 1.5
```

종료 코드 `0`은 성공, `2`는 결과 기준 미달, `1`은 입력·설정·인증·API·응답·출력 오류입니다. stdout 쓰기 전 실패는 빈 stdout과 stderr로 보고하며, 출력 장치의 부분 쓰기는 되돌릴 수 없습니다. CLI는 자동 재시도하거나 HTTP redirect를 따라가지 않습니다.

## 범위와 의존성

단건 명령은 호출당 질문 하나를 처리하며, batch는 하나의 공통 state와 여러 질문을 한 요청으로 처리합니다. 서로 다른 state의 자동 grouping, 문서 탐색, 업무별 정책, 다른 플러그인 자동 연동은 CLI가 제공하지 않습니다. `$jev:jev`는 기본 사용 패턴과 활용 예시를 안내하고, `$jev:scenario-testing`은 batch prediction과 executor 실행을 조율합니다. 연관된 문서를 단순히 개별 호출로 나누는 것만으로는 전체 관계를 검증한 결과가 되지 않습니다.

Go HTTP·JSON·flags와 MIT 라이선스의 `github.com/pelletier/go-toml/v2`를 사용합니다. 제3자 라이선스는 [THIRD_PARTY_NOTICES.md](THIRD_PARTY_NOTICES.md)에 있습니다.

외부 계약: [TypeSafe API](https://docs.typesafe.ai/api), [Noul](https://docs.typesafe.ai/primitives/noul), [Choice](https://docs.typesafe.ai/primitives/choice), [Score](https://docs.typesafe.ai/primitives/score).

루트 `jev --help`는 명령별 한 줄 설명과 공통 판단 옵션의 적용 범위를 표시합니다. 상세 사용법과 옵션은 `jev <command> --help`에서 확인합니다.

## 요청 제한과 오류

단건·batch 모두 직렬화된 JSON 본문은 최대 32 MiB입니다. 초과하면 전송 전에 실패합니다. HTTP 오류는 상태를 유지하며, 인식 가능한 오류 body에서는 검증 위치·종류, context 초과, 크기 제한, rate limit·일시 과부하 안내를 제공합니다. 자동 재시도는 없습니다.

422 위치의 질문 id·선택지명·중첩 사용자 key는 placeholder로 표시합니다. 서버 메시지·입력값·API 키는 출력하지 않습니다. 오류 body가 64 KiB를 넘거나 읽기·형식 검증에 실패하면 HTTP 상태만 안내합니다. API 문서의 최소 개수 설명이 일치하지 않아 Choice·Score는 기존 최소 2개 제한을 유지합니다.

## CLI 내부 구조

`jev/scripts/decision`은 공통 판단 타입과 JSON 값 처리를, `providers`는 `Adapter`·인증 진단 계약을 소유합니다. `providers/jev`가 TypeSafe HTTP 호출·응답 검증·API 인증 진단을 처리합니다. `cli`는 라우팅·provider 조립·오류 보고·종료 코드를 소유하며 `cli/evaluation`은 단건·batch, `cli/configcmd`는 설정 명령, `cli/maintenance`는 설치·제거·진단을 처리합니다. `config`는 설정·우선순위·파일 저장·내장 `template.toml`, `installation`은 바이너리 설치·제거·식별을 담당합니다. 루트 `main.go`와 `go run . install` 진입점은 유지합니다.

```text
CLI 단건·batch·doctor > providers.Adapter > providers/jev > TypeSafe API
```

Jev 어댑터는 검증한 결과만 반환합니다. API·응답 오류는 기존 stderr·종료 코드 계약으로 전달합니다. 후속 OpenRouter·OpenAI decision 모델은 같은 어댑터 경계에 연결할 수 있으며, 현재 실행 가능한 provider는 Jev입니다.
