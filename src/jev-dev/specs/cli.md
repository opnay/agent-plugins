# Jev CLI 계약

## Primitive 명령

```text
jev noul --if <question and context> [options]
jev choice --if <question and context> --conditions <a,b,...> [options]
jev score --if <question and context> --level <description> --level <description> [options]
```

- `--if`는 질문과 필요한 문맥을 함께 담으며 공백뿐인 입력을 허용하지 않는다.
- Noul은 예/아니오 명제의 참일 확률을 구한다. Choice는 명시한 선택지 중 하나를 고른다. Score는 순서 있는 단계에서 정도를 평가한다.
- Choice 선택지는 앞뒤 공백을 제거한 2~255개의 고유하고 비어 있지 않은 문자열이다. 선택지 내부 쉼표는 지원하지 않는다.
- Score 단계는 반복 `--level`로 순서대로 전달한 2~10개의 고유하고 비어 있지 않은 자연어 설명이다. 단계 설명에는 쉼표를 쓸 수 있다.
- 기존 `jev --if ... --conditions ...`는 `jev choice`와 같은 호환 형식이다.
- 다른 primitive의 입력 옵션, 알 수 없는 옵션, 추가 positional argument는 API 호출 전에 오류다.

## 구조화 단건 입력과 장문 모드

- `jev <noul|choice|score> --input <path>`는 UTF-8 JSON object 하나를 읽습니다. 허용 key는 `state`, `instructions`, `criteria`입니다. 알 수 없는 key와 여러 JSON 값은 오류입니다.
- `state`는 필수 문자열·객체·배열입니다. 문자열은 공백뿐이면 오류입니다. `instructions`는 선택 문자열·객체·배열·null이며 생략 시 null입니다. 객체·배열 내부 JSON 값과 숫자 표기를 보존합니다.
- `--input`과 `--if`, `--conditions`, `--level`은 함께 사용하지 않습니다. 기존 간단한 입력·출력·기준 옵션은 유지합니다.
- Choice `criteria`는 2~255개 고유한 비어 있지 않은 label의 object입니다. label 앞뒤 공백은 오류이며 값은 설명 문자열·객체·배열·null입니다.
- Score `criteria`는 낮은 단계부터 2~10개 고유한 문자열·객체·배열입니다. 공백뿐인 문자열과 null 단계는 오류입니다. 반환 `legend`는 각 단계의 원래 JSON 구조를 유지하며 요청 기준과 의미적으로 일치해야 합니다.
- Noul `criteria`는 선택 object 또는 null입니다. 허용 key는 `true`, `false`이며 값은 문자열·객체·배열·null입니다.
- JSON은 UTF-8이어야 하며 state·instructions·criteria의 중첩은 각각 최대 64단계입니다. 비유한 수와 float64 범위를 초과하는 수는 요청 전에 거부합니다. 오류에는 입력 원문을 출력하지 않습니다.
- `--long-context[=false]`는 단건·batch 호출 전용 boolean이며 기본 false입니다. true일 때만 `TypeSafe-Danger-Zone-Long-Context: true` 헤더를 보냅니다. 설정 저장·doctor에는 적용하지 않습니다.
- 기본 한도는 요청 전체 64k·state와 질문 하나 32k 토큰입니다. 장문 모드는 두 한도를 각각 262k로 확장하며, 속도·정확도 저하 가능성을 안내합니다. 자동 활성화·분할·재시도는 없습니다.

## Batch 명령

```text
jev batch --state-file <path> --input <path> [--model <id>] [--timeout <duration>] [--state-format text|json] [--long-context]
```

- `--state-file`은 공통 UTF-8 state 파일입니다. `--state-format text`가 기본이며 비어 있지 않은 원문을 trim하지 않고 전송합니다. `--state-format json`은 문자열·객체·배열 JSON 하나를 원형대로 전달하며 구조화 단건과 동일하게 검증합니다. 자동 형식 추론은 없습니다.
- `--input`은 한 줄에 질문 하나를 둔 JSONL 파일이다. 공백 줄은 무시하며 질문이 하나 이상이어야 한다.
- 각 줄은 고유하고 비어 있지 않은 `id`, `type`과 `question` 또는 `instructions` 중 정확히 하나를 포함합니다. `question`은 기존 비어 있지 않은 문자열이며 `instructions`는 문자열·객체·배열·null입니다. Key는 정확한 snake_case만 허용하고 알 수 없거나 대소문자가 다른 field는 오류다.
- `id`, `conditions`, `levels` 값의 앞뒤 공백은 자동 정리하지 않고 오류로 처리한다. 검증된 원문을 request와 output에 유지한다.
- `type`은 `noul`, `choice`, `score` 중 하나다.
- Choice는 `conditions` 문자열 배열에 2~255개의 고유하고 비어 있지 않은 label을 둔다. JSON 문자열이므로 쉼표를 허용한다.
- Score는 `levels` 문자열 배열에 2~10개의 고유하고 비어 있지 않은 설명을 낮은 단계부터 둔다.
- 각 row는 기존 `conditions`·`levels` 대신 구조화 단건과 동일한 `criteria`를 사용할 수 있습니다. `criteria`와 기존 기준 field는 함께 사용하지 않습니다. Noul에는 `conditions`·`levels`, Choice에는 `levels`, Score에는 `conditions`를 허용하지 않습니다.
- 전체 state와 JSONL을 API 호출 전에 검증한다. 한 줄이라도 잘못되면 요청하지 않는다.
- 모든 질문은 하나의 SystemOne 요청으로 전송한다. 자동 chunk, retry, resume, 부분 성공은 제공하지 않는다.
- `model`과 `timeout`은 batch 전체에 공통 적용한다. threshold, min_score, json, pick은 batch 입력·옵션이 아니며 저장값도 batch에 적용하지 않는다.

## Usage 도움말

- `jev --help`는 전체 탐색용이다. `Evaluation`, `Configuration`, `Maintenance`, `Common evaluation options`, `Primitive options`, `Config keys`, `Output`, `Exit codes`로 구분한다.
- `jev noul --help`, `jev choice --help`, `jev score --help`는 해당 primitive의 형식, 전용·공통 옵션, 기본·JSON 출력, 결과 기준, 종료 코드를 설명한다.
- `jev batch --help`는 state·JSONL schema, 공통 옵션, 원자적 실패, JSONL 출력, 종료 코드를 설명한다.
- `jev config --help`는 config 호출 형식, 지원 key·기본값·적용 범위, 인증 우선순위를 설명한다.
- `jev install --help`, `jev uninstall --help`, `jev doctor --help`는 해당 유지보수 명령의 옵션, 대상, 동작·보존 범위를 설명한다.
- 모든 help는 설정파일·토큰을 읽거나 API를 호출하지 않고 종료 코드 0으로 stdout에 출력한다.
- 옵션·positional argument 오류는 가능한 경우 해당 subcommand help를 가리킨다.

## 요청 계약

- 질문 하나를 `POST https://api.typesafe.ai/v1/systemone`으로 보낸다.
- 간단한 단건의 `state`는 `--if` 원문이고, 구조화 단건은 파일의 state·instructions·criteria를 사용합니다. question id는 `result`, model은 실효 설정값입니다.
- 간단한 Noul question은 type `noul`, instruction `Answer whether the proposition or question in state is true.`이며 criteria를 보내지 않는다.
- 간단한 Choice question은 type `choice`, instruction `Select the option that best answers the question in state.`이며 선택지 이름을 key, `null`을 값으로 한 criteria를 보낸다.
- 간단한 Score question은 type `score`, instruction `Rate the state against the ordered criteria.`이며 `--level` 순서의 문자열 배열을 criteria로 보낸다.
- Batch는 선택한 형식의 공통 state와 각 row의 id·question/instructions·criteria를 사용합니다. 기존 conditions·levels는 이전 criteria 형태로 변환합니다.
- Batch의 모든 row는 하나의 `questions` map에 들어가며 API 요청은 한 번이다.
- 임의 endpoint 옵션, 자동 재시도, API 오류 본문 출력은 제공하지 않는다. HTTP redirect는 따라가지 않는다.
- timeout·취소·HTTP 오류·잘못된 응답은 실행 실패다.

## 결과와 출력

- Noul answer는 0~1의 참일 확률이다. 정도를 나타내는 점수가 아니다.
- Choice answer는 요청 선택지에 속하는 API choice이며 최고 확률이어야 한다. 동률이면 API 선택을 유지한다.
- Score answer는 0부터 `단계 수 - 1`까지의 확률 가중 점수이며 단계 사이 값이 가능하다.
- probabilities, confidence, legend, model은 API 값을 유지한다. confidence를 재계산하지 않는다.
- 필요한 응답 필드가 없거나 타입·범위·선택지·단계와 맞지 않으면 오류다.
- 기본 출력은 타입별 answer와 줄바꿈 하나다.
- `--json`은 primitive별 객체를 출력한다.
  - Noul: `answer`, `model`.
  - Choice: `answer`, `probabilities`, `confidence`, `model`.
  - Score: `answer`, `legend`, `probabilities`, `confidence`, `model`.
- `--pick <fields>`는 JSON 최상위 필드만 선택한다. 한 필드도 객체로 출력한다. `--pick ''`는 필드 제한을 해제한다.
- pick은 최종 JSON 모드에서만 허용한다. primitive가 제공하지 않는 필드, 알 수 없는 필드, 빈 항목, 중복 항목은 API 호출 전에 오류다.
- CLI·설정·JSON이 소유하는 key는 `snake_case`다. Choice 라벨은 자유 문자열이며 기계적 후속 처리에는 `snake_case`를 권장한다. Score 단계는 자연어 설명을 유지한다.
- Batch 성공 출력은 입력 row마다 JSON object 하나인 JSONL이다. 입력 순서와 `id`, `type`을 유지하고 primitive별 전체 결과 field를 포함한다.
- Batch 응답은 model과 모든 question id·type·결과를 검증한다. 누락·추가 answer, 잘못된 type·범위·criteria 결과는 전체 오류다.

## 결과 기준과 종료 코드

- `threshold` 미지정은 필터 없음이다. Choice에서는 `probabilities[answer]`, Noul에서는 answer가 threshold 이상이면 통과한다.
- Score는 `min_score` 미지정 시 필터가 없다. 지정 값은 0부터 현재 요청의 `단계 수 - 1`까지이며 answer가 그 값 이상이면 통과한다.
- confidence는 threshold나 min_score 검사에 쓰지 않는다.
- 종료 코드 `0`: 성공 결과는 stdout.
- 종료 코드 `2`: 결과 기준 미달, stdout 없음·stderr에 결과와 기준.
- 종료 코드 `1`: 입력·설정·인증·통신·응답·출력 오류, stderr에 이유.
- JSON 여부나 pick은 실패 시 stdout을 비우는 계약을 바꾸지 않는다. 출력 장치의 부분 쓰기 실패는 되돌릴 수 없다.
- Batch는 전체 JSONL을 먼저 encode한 뒤 stdout에 쓰고 성공 시 종료 코드 0을 반환한다. 쓰기 전 입력·파일·설정·API·응답 오류는 stdout 없이 종료 코드 1이며 종료 코드 2를 사용하지 않는다. 출력 장치의 부분 쓰기 실패는 되돌릴 수 없다.

## 설정

- 경로: `~/.agents/jev.toml`. `[defaults]`는 provider·timeout·threshold·min_score·json·pick, `[providers.<name>]`은 model·api_key를 소유합니다.
- 내장 기본값: provider `jev`, timeout `30s`, json false, pick 제한 없음, threshold·min_score 미적용. Jev 모델 기본값은 `jev-latest`이며 다른 provider에는 모델 기본값이 없습니다.
- provider 선택: 명시적 `--provider` > defaults.provider > 내장 jev. model: `--model` > 선택 provider의 model > 해당 provider 내장값. 공통 옵션: 명시적 flag > defaults > 내장값.
- 인증: 선택한 provider의 비어 있지 않은 환경변수 > 해당 provider의 api_key. jev는 TYPESAFE_API_KEY, openrouter는 OPENROUTER_API_KEY, openai는 OPENAI_API_KEY입니다. 다른 provider의 키·모델로 대체하지 않습니다.
- 단건·batch·doctor는 `--provider`를 지원합니다. 현재 실행 구현은 Jev뿐이며 OpenRouter·OpenAI는 설정만 저장할 수 있습니다. 미구현 provider 선택은 네트워크 호출 없이 명확한 오류입니다.
- 설정 전체의 TOML·지원 key·타입·값을 검증하되 미선택 provider의 model·api_key 부재는 허용합니다. threshold는 유한한 0~1, min_score는 유한한 0 이상, timeout은 양수 duration, model은 지정 시 비어 있지 않은 문자열입니다. pick은 지원 field의 고유 목록입니다. provider 이름은 jev·openrouter·openai만 허용합니다.
- Batch는 defaults의 provider·timeout과 해당 provider의 model·인증만 사용하며 threshold·min_score·json·pick은 무시합니다. long_context는 요청 전용이며 저장하지 않습니다.
- 구형 최상위 api_key·model·공통 key나 단수 default 섹션은 오류입니다. defaults와 providers.jev로 옮길 위치를 안내하며 자동 변환·저장하지 않습니다. 구형·신형 혼용도 추측하지 않습니다.

## Config 커맨드와 최초 생성

- `jev config`: defaults의 내장 기본값과 모든 설정된 provider를 TOML로 표시합니다. Jev 기본 모델도 표시합니다. 각 provider 환경변수를 반영하되 모든 key는 전체 마스킹합니다. 미구현 provider가 기본값이어도 오프라인 조회·편집은 가능합니다.
- `jev config path`: 파일을 읽지 않고 경로만 출력합니다.
- `jev config init`: go:embed 내장 `config.template.toml`로 파일을 생성합니다. 기존 파일은 유효성에 관계없이 덮어쓰지 않고 오류입니다. init·set 성공은 무출력입니다.
- `jev config set defaults.<key> <value>`, `jev config set providers.<name>.model <value>`, `jev config set providers.<name>.api_key --stdin`을 사용합니다. key 원문 인자 입력은 금지하며 stdin 토큰은 앞뒤 공백을 제거하고 64KiB 이내여야 합니다.
- `jev config unset <dotted-key>`는 저장값만 제거합니다. 항목이 없으면 성공이며 파일을 만들지 않습니다. default 값은 이후 코드의 내장값으로 해석합니다.
- 파일이 없는 최초 set은 같은 내장 템플릿에 요청값을 반영합니다. 생성 파일에는 템플릿의 안내 주석과 변경한 값이 포함됩니다. 이후 변경은 TOML 주석·서식 보존을 보장하지 않습니다.
- 템플릿에는 활성 defaults와 providers.jev.model, 선택 threshold·min_score와 미구현 provider의 주석 예시를 둡니다. token이나 활성 placeholder를 포함하지 않습니다. 내장 기본값과 템플릿의 일치 여부를 검증합니다.
- 조회·평가·doctor는 파일이나 디렉터리를 자동 생성하지 않습니다. config는 네트워크·인증을 요구하지 않습니다.
- 파일은 일반 파일이어야 하며 symlink·디렉터리는 거부합니다. 갱신은 같은 디렉터리의 임시 0600 파일로 원자적 교체합니다. 최초 생성은 원자적 no-replace로 기존 파일을 보존합니다.
- 설정·key는 평문입니다. 표시·오류·doctor는 모든 provider의 token 원문을 숨깁니다. 잘못된 TOML 원문을 오류에 포함하지 않습니다.

## 설치·제거·진단

- 공통 형식: `jev install [--force] [--dir <directory>]`, `jev uninstall [--dir <directory>]`, `jev doctor [--dir <directory>]`.
- 기본 디렉터리는 `~/.local/bin`, 정확한 대상은 그 디렉터리의 `jev`다. `--dir`는 비어 있지 않은 디렉터리이며 상대 경로도 허용한다.
- install/uninstall은 API 호출·설정파일·토큰 없이 동작한다.
- install은 실행 중인 바이너리를 대상에 원자적으로 복사하고 0755 권한을 부여한다. 같은 바이너리는 성공하고 다른 일반 파일은 `--force` 없이는 보존한다. symlink·비정규 파일은 거부한다.
- uninstall은 Go build info의 main package path가 `opnay/jev`인 일반 파일만 제거한다. 대상이 없으면 성공이며 설정·토큰·설치 디렉터리·다른 파일·셸 설정은 유지한다.
- doctor는 대상 identity, PATH 선택, 일반 설정 유효성·권한, 키 유무·출처를 점검한다. 비어 있지 않은 `TYPESAFE_API_KEY`가 파일의 `api_key`보다 우선하며 원문은 출력하지 않는다.
- 기본 doctor는 설정이 유효하고 키가 있을 때 `GET https://api.typesafe.ai/v1/models`를 Bearer 인증으로 한 번 호출한다. 저장된 timeout을 적용하며 redirect·retry·추론 요청은 하지 않는다. 설치·PATH 오류는 독립적인 API 점검을 막지 않는다.
- HTTP 200은 인증·조회 접근 성공이다. 401은 잘못되거나 만료된 키, 403은 접근 거부로 안내한다. 429·서버 오류·redirect·네트워크 오류·timeout·취소는 인증 성공이나 키 오류로 단정하지 않는다. 서버 body·원문 transport 오류는 출력하지 않는다.
- 설정 오류·키 누락은 API 점검을 SKIP하며 로컬 문제는 실패다. 별도 offline 진단은 제공하지 않는다.
- 결과는 로컬 설정·인증 섹션, `OK / FAIL / SKIP` 상태, 항목별 상세·해결 안내, 상태별 개수 요약을 포함한다. primitive별 pick·범위 적합성·모델 정확성·추론 가능 여부는 판정하지 않는다.
- doctor는 읽기 전용이며 일반 판정·설치 뒤에 자동 호출하지 않는다.
- 유지보수 성공은 stdout과 종료 코드 0, 실패는 빈 stdout·stderr와 종료 코드 1이다. 코드 2는 판정 기준 미달 전용이다.

## Provider 어댑터 경계

- `scripts/providers`는 provider 공통 Request·Question·Result와 Adapter 인터페이스를 소유합니다. CLI 옵션·설정·파일 경로·stdout·종료 코드는 공통 provider 계약에 포함하지 않습니다.
- Adapter는 공통 요청을 평가해 question id별 정규화 결과를 반환하며, 별도 인증 진단 메서드를 제공합니다. 인증 결과는 key 거부 여부와 진단을 반환하며 환경변수·설정 파일에 따른 key 복구 안내는 CLI가 소유합니다. 응답 wire schema는 provider 밖으로 노출하지 않습니다.
- `scripts/providers/jev`는 TypeSafe endpoint·인증 헤더·장문 헤더·HTTP 호출·Jev 응답 해석과 검증·models 인증 진단을 소유합니다. 키·timeout·transport는 생성 시 주입합니다.
- 단건과 batch는 같은 어댑터를 호출합니다. 반환 전 model·id 완전성·primitive·선택지·확률·Score legend를 검증하며 오류면 CLI는 기존 실패 계약을 유지합니다.
- 현재 연결 대상은 Jev 하나입니다. OpenRouter·OpenAI 구현·provider 선택 옵션·인증 설정 확장은 별도 작업입니다. 기존 CLI 입력·출력·기본값·인증·재시도 정책은 유지합니다.

## 배포와 검증

- Go 1.23 이상, stdlib HTTP/JSON/flags와 `github.com/pelletier/go-toml/v2`를 사용한다. 공통 통신과 primitive별 입력·응답 검증을 분리한다.
- `jev/scripts/`에서 `go run . install [--force] [--dir <directory>]`로 설치한다. 별도 셸 래퍼 없이 Go CLI가 설치 동작을 소유한다.
- 플러그인 설치는 바이너리 설치나 PATH 편집을 수행하지 않는다. 빌드 실패는 기존 바이너리를 보존한다.
- 자동 테스트는 임시 설정 경로와 가짜 HTTP transport를 사용한다. 실제 토큰·사용자 설정·유료 API를 사용하지 않는다.
- 검증 대상: primitive 라우팅·입력, 요청 schema, 응답 타입·범위, 타입별 출력·pick, 결과 기준 경계, 기존 Choice 호환성, 설정 우선순위, 유지보수 회귀.
- Batch 검증 대상: JSONL schema·중복 id·primitive criteria, 전체 선검증, 단일 요청, 공통 state·model, 응답 완전성, 입력 순서 JSONL, 원자적 실패, batch 비적용 설정.
- 최상위 usage 카테고리, subcommand별 help 내용·라우팅, 설정 독립성을 검증한다.
- 실제 API smoke test는 사용 가능한 토큰과 해당 호출 권한이 있을 때 별도로 수행한다.

## 외부 계약 근거

2026-09-22 확인: [API](https://docs.typesafe.ai/api), [Noul](https://docs.typesafe.ai/primitives/noul), [Choice](https://docs.typesafe.ai/primitives/choice), [Score](https://docs.typesafe.ai/primitives/score).
Noul의 참일 확률, Choice의 최고 확률 option, Score의 확률 가중 값과 각 타입의 confidence 차이를 적용한다. 실제 모델의 판단 정확성은 로컬 테스트가 증명하지 않는다.

- Doctor 인증 조회 근거: [TypeSafe OpenAPI](https://api.typesafe.ai/openapi.json)의 `/v1/models` GET, HTTPBearer security (2026-09-30 확인).

- 구조화 입력·Score legend 근거: [공개 OpenAPI](https://api.typesafe.ai/openapi.json), [구조화 입력](https://docs.typesafe.ai/primitives/advanced.md), [장문 모드](https://docs.typesafe.ai/models.md#long-context), 2026-10-19 확인. 이미지·Box·usage 확장은 포함하지 않습니다.
