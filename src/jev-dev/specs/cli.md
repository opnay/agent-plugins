# Jev CLI 계약

## 판정 입력

`jev --if <질문과 문맥> --conditions <쉼표로 구분한 선택지>`

- `--if`는 공백뿐인 입력을 허용하지 않는다. 문맥은 이 값에 포함한다.
- 선택지는 앞뒤 공백을 제거한 2~255개의 고유하고 비어 있지 않은 문자열이다. 선택지 내부 쉼표는 지원하지 않는다.
- 질문 하나를 `POST https://api.typesafe.ai/v1/systemone`으로 보낸다.
- `state`는 `--if` 원문, `questions.result.type`은 `choice`, `instructions`는 `Select the option that best answers the question in state.`이다.
- `criteria`는 선택지 이름을 key, `null`을 값으로 사용한다. `model`과 Bearer API key를 전달한다.
- 임의 endpoint 옵션, 자동 재시도, API 오류 본문 출력은 제공하지 않는다. 호출 횟수와 비밀정보 노출을 제한한다.
- HTTP redirect는 따라가지 않는다. timeout·취소·HTTP 오류·잘못된 응답은 실행 실패다.

## 결과와 종료 코드

- `answer`: API `choice`. 요청 선택지에 속하며 최고 확률이어야 한다. 동률이면 API 선택을 유지한다.
- `probabilities`, `confidence`, `model`: API 값을 유지한다. confidence를 재계산하지 않는다.
- 필요한 응답 필드·선택지 확률이 없거나 값이 유효하지 않으면 오류다.
- 기본 출력은 `answer`와 줄바꿈 하나다.
- `--json`: `answer`, `probabilities`, `confidence`, `model`을 가진 JSON 객체와 줄바꿈.
- `--pick answer,probabilities`: JSON 최상위 필드만 선택한다. 한 필드도 객체로 출력한다. `--pick ''`는 필드 제한을 해제한다.
- pick은 최종 JSON 모드에서만 허용한다. 알 수 없는 필드·빈 항목·중복 항목은 실행 전에 오류다.
- threshold 미지정은 필터 없음이다. 지정 시 `probabilities[answer] >= threshold`면 통과한다. API confidence는 이 검사에 쓰지 않는다.
- 종료 코드 `0`: 성공 결과는 stdout. `2`: 임계값 미달, stdout 없음·stderr에 선택지/확률/기준. `1`: 입력·설정·인증·통신·출력 오류, stderr에 이유.
- JSON 여부나 pick은 실패 시 stdout을 비우는 계약을 바꾸지 않는다. 출력 장치 자체의 부분 쓰기 실패는 되돌릴 수 없다.

## 설정

- 경로: 사용자 홈의 `.agents/jev.toml`.
- 우선순위: 명시적 CLI 옵션 > 파일 > 내장 기본값. 인증은 비어 있지 않은 `TYPESAFE_API_KEY` > 파일 `api_key`.
- CLI flag: `--threshold`, `--model`, `--timeout`, `--json`, `--pick`. 토큰 CLI flag는 없다.
- `threshold`: 유한한 0~1 실수, 기본 미적용. 명시적 0은 저장된 양수 기준을 덮어쓴다.
- `model`: 비어 있지 않은 문자열, 기본 `jev-latest`.
- `timeout`: 양수 Go duration 문자열, 기본 `30s`.
- `json`: boolean, 기본 false. 명시적 `--json=false`가 파일보다 우선한다.
- `pick`: 문자열 배열, 기본 제한 없음. 빈 배열은 제한 없음이다.
- `api_key`: 문자열, 기본 없음. 판정 호출에만 비어 있지 않은 인증이 필요하다.
- 파일 부재는 오류가 아니다. 조회·판정은 파일을 생성하지 않는다.
- TOML 구문·타입·미지원 key 오류를 조용히 무시하지 않는다. 오류에 설정 원문을 포함하지 않는다.

## Config 커맨드

- `jev config`: 내장 기본값과 환경변수를 반영한 설정을 TOML로 출력. 토큰은 전체 마스킹, 미설정 threshold는 생략한다.
- `jev config path`: 설정 경로만 출력. 파일의 존재나 유효성에 의존하지 않는다.
- `jev config set <key> <value>`: 저장값 갱신. pick 값은 쉼표 목록이며 빈 문자열은 빈 배열이다.
- `jev config set api_key --stdin`: stdin에서 토큰을 읽고 앞뒤 공백을 제거해 저장. 토큰은 이 방식만 지원한다.
- `jev config unset <key>`: 저장값 제거. 해당 항목이 없으면 성공하며 새 파일을 만들지 않는다.
- config 명령은 네트워크나 인증을 요구하지 않는다. set/unset 성공은 무출력이다.
- 각 설정의 타입과 값은 저장 시 검사한다. pick/json 조합은 판정 실행 시 검사하므로 여러 config 명령으로 순차 설정할 수 있다.
- 갱신 시 나머지 지원 설정값을 보존한다. TOML 주석·서식 보존은 보장하지 않는다.
- 파일은 일반 파일이어야 하며 symlink·디렉터리를 읽거나 덮어쓰지 않는다. 저장은 같은 디렉터리의 임시 파일을 rename하는 원자적 갱신이며 파일 권한은 0600이다.
- 토큰은 평문 저장이다. 원문을 조회·오류·로그에 출력하지 않는다. 잘못된 TOML은 사용자가 직접 수정한다.

## 설치·제거·진단

- 공통 형식: `jev install [--force] [--dir <directory>]`, `jev uninstall [--dir <directory>]`, `jev doctor [--dir <directory>]`.
- 기본 디렉터리는 `~/.local/bin`, 정확한 대상은 그 디렉터리의 `jev`다. `--dir`는 비어 있지 않은 디렉터리이며 상대 경로도 허용한다. 명시적 빈 값·알 수 없는 옵션·추가 positional argument는 오류다.
- 유지보수 명령은 API를 호출하지 않는다. install/uninstall은 설정파일이나 토큰 없이 동작한다.
- install은 실행 중인 바이너리를 대상에 원자적으로 복사하고 0755 권한을 부여한다. 최신 버전 다운로드나 Go 빌드는 수행하지 않는다.
- 같은 실행 파일 또는 내용·권한이 같은 대상은 성공으로 끝낸다. 다른 기존 일반 파일은 `--force` 없이는 보존한다. symlink·디렉터리·기타 비정규 파일은 force와 무관하게 거부한다.
- uninstall은 Go build info의 main package path가 `opnay/jev`인 일반 파일만 제거한다. 임의의 동명 파일, symlink·디렉터리는 제거하지 않는다. 대상이 없으면 성공이다.
- uninstall은 `jev.toml`·토큰·설치 디렉터리·다른 파일·셸 설정을 유지한다. 바이너리는 소스에서 `go run . install`로 복구할 수 있다.
- doctor는 대상이 Jev 실행 파일인지, PATH에서 선택되는 `jev`가 그 파일인지, 설정과 JSON/pick 조합이 유효한지, 파일 권한이 소유자에게만 열려 있는지, 실효 토큰이 존재하는지 점검한다.
- 설정파일 부재는 내장 기본값 사용으로 처리한다. 토큰은 유무만 출력하고 원문을 노출하지 않는다. doctor 성공은 API 인증 성공을 뜻하지 않는다.
- doctor는 읽기 전용이다. 일반 판정·설치 뒤에 자동 호출하지 않는다.
- 유지보수 성공은 stdout의 짧은 결과와 종료 코드 0, 실패는 빈 stdout·stderr의 이유와 종료 코드 1이다. doctor는 여러 점검 결과를 한 보고서로 출력한다. 코드 2는 판정 threshold 미달 전용이다.

## 배포와 검증

- Go 1.23 이상, stdlib HTTP/JSON/flags와 `github.com/pelletier/go-toml/v2` 사용. CLI·설정·API·출력 책임을 나눈다.
- `jev/scripts/`에서 `go run . install [--force] [--dir <directory>]`: Go가 실행하는 임시 바이너리의 install로 설치한다. 별도 셸 래퍼 없이 Go CLI가 설치 동작을 소유한다.
- 플러그인 설치는 바이너리 설치나 PATH 편집을 수행하지 않는다. 빌드 실패는 기존 바이너리를 보존한다.
- 자동 테스트는 임시 설정 경로와 가짜 HTTP transport를 사용한다. 실제 토큰·사용자 설정·유료 API를 사용하지 않는다.
- 검증 대상: 입력 오류, 우선순위, 미지정과 0/false 구분, config 왕복·권한·토큰 비노출, 요청 형식, API 실패, threshold 경계, 출력/종료 코드, 유지보수의 대상 보호·설정 보존·오프라인 진단.
- 실제 API smoke test는 사용 가능한 토큰과 해당 호출 권한이 있을 때 별도로 수행한다.

## 외부 계약 근거

2026-09-21 확인: [API](https://docs.typesafe.ai/api), [Choice](https://docs.typesafe.ai/primitives/choice).
HTTP schema와 최고 확률 choice·별도 confidence 의미를 적용한다. 실제 모델의 판단 정확성은 로컬 테스트가 증명하지 않는다.
