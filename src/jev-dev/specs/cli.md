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

## Batch 명령

```text
jev batch --state-file <path> --input <path> [--model <id>] [--timeout <duration>]
```

- `--state-file`은 모든 질문이 공유할 UTF-8 state 원문이다. 공백뿐인 파일은 오류이며 내용은 trim하지 않고 전송한다.
- `--input`은 한 줄에 질문 하나를 둔 JSONL 파일이다. 공백 줄은 무시하며 질문이 하나 이상이어야 한다.
- 각 줄은 고유하고 비어 있지 않은 `id`, `type`, 비어 있지 않은 `question`을 포함한다. Key는 정확한 snake_case만 허용하고 알 수 없거나 대소문자가 다른 field는 오류다.
- `id`, `conditions`, `levels` 값의 앞뒤 공백은 자동 정리하지 않고 오류로 처리한다. 검증된 원문을 request와 output에 유지한다.
- `type`은 `noul`, `choice`, `score` 중 하나다.
- Choice는 `conditions` 문자열 배열에 2~255개의 고유하고 비어 있지 않은 label을 둔다. JSON 문자열이므로 쉼표를 허용한다.
- Score는 `levels` 문자열 배열에 2~10개의 고유하고 비어 있지 않은 설명을 낮은 단계부터 둔다.
- Noul에는 `conditions`·`levels`, Choice에는 `levels`, Score에는 `conditions`를 허용하지 않는다.
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
- `state`는 `--if` 원문이고, question id는 `result`, model은 실효 설정값이다.
- Noul question은 type `noul`, instruction `Answer whether the proposition or question in state is true.`이며 criteria를 보내지 않는다.
- Choice question은 type `choice`, instruction `Select the option that best answers the question in state.`이며 선택지 이름을 key, `null`을 값으로 한 criteria를 보낸다.
- Score question은 type `score`, instruction `Rate the state against the ordered criteria.`이며 `--level` 순서의 문자열 배열을 criteria로 보낸다.
- Batch는 `--state-file` 원문을 공통 `state`로 보내고 각 row의 `id`를 question id, `question`을 instructions로 사용한다. primitive별 criteria는 row의 `conditions` 또는 `levels`에서 만든다.
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

- 경로: 사용자 홈의 `.agents/jev.toml`.
- 우선순위: 명시적 CLI 옵션 > 파일 > 내장 기본값. 인증은 비어 있지 않은 `TYPESAFE_API_KEY` > 파일 `api_key`.
- 공통 CLI flag: `--model`, `--timeout`, `--json`, `--pick`. Choice·Noul은 `--threshold`, Score는 `--min-score`를 추가한다. 토큰 CLI flag는 없다.
- Batch는 `--model`, `--timeout`만 사용한다. 저장된 `api_key`, `model`, `timeout`은 적용하고 threshold, min_score, json, pick은 무시한다.
- `threshold`: 유한한 0~1 실수, 기본 미적용. 명시적 0은 저장된 양수 기준을 덮어쓴다.
- `min_score`: 유한한 0 이상 실수, 기본 미적용. Score 실행 시 현재 단계 범위도 검사한다.
- `model`: 비어 있지 않은 문자열, 기본 `jev-latest`.
- `timeout`: 양수 Go duration 문자열, 기본 `30s`.
- `json`: boolean, 기본 false. 명시적 `--json=false`가 파일보다 우선한다.
- `pick`: 문자열 배열, 기본 제한 없음. 빈 배열은 제한 없음이다.
- `api_key`: 문자열, 기본 없음. 판정 호출에만 비어 있지 않은 인증이 필요하다.
- 파일 부재는 오류가 아니다. 조회·판정은 파일을 생성하지 않는다.
- TOML 구문·타입·미지원 key 오류를 조용히 무시하지 않는다. 오류에 설정 원문을 포함하지 않는다.
- 저장된 threshold는 Choice·Noul에만 적용하고 저장된 min_score는 Score에만 적용한다.

## Config 커맨드

- `jev config`: 내장 기본값과 환경변수를 반영한 설정을 TOML로 출력한다. 토큰은 전체 마스킹하고 미설정 threshold·min_score는 생략한다.
- `jev config path`: 설정 경로만 출력하며 파일 존재나 유효성에 의존하지 않는다.
- `jev config set <key> <value>`: 저장값을 갱신한다. pick 값은 쉼표 목록이며 빈 문자열은 빈 배열이다.
- `jev config set api_key --stdin`: stdin에서 토큰을 읽고 앞뒤 공백을 제거해 저장한다. 토큰은 이 방식만 지원한다.
- `jev config unset <key>`: 저장값을 제거한다. 항목이 없으면 성공하며 새 파일을 만들지 않는다.
- config 명령은 네트워크나 인증을 요구하지 않는다. set/unset 성공은 무출력이다.
- 각 설정의 타입과 독립 범위는 저장 시 검사한다. pick과 primitive의 조합, min_score와 Score 단계 범위는 판정 실행 시 검사한다.
- 갱신 시 나머지 지원 설정값을 보존한다. TOML 주석·서식 보존은 보장하지 않는다.
- 파일은 일반 파일이어야 하며 symlink·디렉터리를 읽거나 덮어쓰지 않는다. 저장은 같은 디렉터리의 임시 파일을 rename하는 원자적 갱신이며 권한은 0600이다.
- 토큰은 평문 저장이다. 원문을 조회·오류·로그에 출력하지 않는다. 잘못된 TOML은 사용자가 직접 수정한다.

## 설치·제거·진단

- 공통 형식: `jev install [--force] [--dir <directory>]`, `jev uninstall [--dir <directory>]`, `jev doctor [--dir <directory>]`.
- 기본 디렉터리는 `~/.local/bin`, 정확한 대상은 그 디렉터리의 `jev`다. `--dir`는 비어 있지 않은 디렉터리이며 상대 경로도 허용한다.
- 유지보수 명령은 API를 호출하지 않는다. install/uninstall은 설정파일이나 토큰 없이 동작한다.
- install은 실행 중인 바이너리를 대상에 원자적으로 복사하고 0755 권한을 부여한다. 같은 바이너리는 성공하고 다른 일반 파일은 `--force` 없이는 보존한다. symlink·비정규 파일은 거부한다.
- uninstall은 Go build info의 main package path가 `opnay/jev`인 일반 파일만 제거한다. 대상이 없으면 성공이며 설정·토큰·설치 디렉터리·다른 파일·셸 설정은 유지한다.
- doctor는 대상 identity, PATH 선택, 일반 설정 유효성·권한, 토큰 유무를 읽는다. primitive별 pick·범위 적합성이나 API 인증은 판정하지 않는다.
- doctor는 읽기 전용이며 일반 판정·설치 뒤에 자동 호출하지 않는다.
- 유지보수 성공은 stdout과 종료 코드 0, 실패는 빈 stdout·stderr와 종료 코드 1이다. 코드 2는 판정 기준 미달 전용이다.

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
