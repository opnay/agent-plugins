# Jev

TypeSafe의 hosted Jev 모델을 호출하는 Go CLI와 사용 스킬입니다. 질문 하나와 명시적 선택지를 보내고, 선택 결과와 확률을 받습니다. 로컬 모델이나 문서 자동 검증기는 아닙니다.

- `$jev:jev`: 질문·선택지 구성, CLI 결과 사용, 요청된 설정·설치·제거·진단.
- `jev`: Choice API 호출, threshold, 출력 제어, `config`, `install`, `uninstall`, `doctor`.
- 질문과 선택지는 TypeSafe API로 전송됩니다. 토큰과 전송 가능한 입력을 준비해야 합니다.

## 설치

Go 1.23 이상이 필요합니다. 이 플러그인의 루트에서 실행합니다.

```sh
cd scripts
go run . install
```

`go run`이 빌드한 임시 바이너리를 CLI의 `install`이 설치합니다. 기본 위치는 `~/.local/bin/jev`입니다. 다른 위치는 `go run . install --dir /absolute/bin`으로 지정합니다. 동일 바이너리 재설치는 변경 없이 성공합니다. 소스 갱신 후 의도적인 교체는 `go run . install --force`를 사용합니다. symlink·디렉터리는 교체하지 않습니다. 빌드 실패 시 기존 바이너리는 유지됩니다.

플러그인 설치와 CLI 설치는 별개입니다. CLI는 셸 설정을 변경하지 않으므로, 설치 디렉터리가 PATH에 포함되어야 합니다. 빌드에는 Go 모듈 다운로드가 필요할 수 있습니다. 설정파일이나 토큰은 설치 과정에서 생성하지 않습니다.

## 설치·제거·진단 커맨드

```sh
jev install
jev install --force --dir /absolute/bin
jev uninstall
jev doctor
```

세 커맨드 모두 `--dir <directory>`를 지원합니다. 기본 대상은 `~/.local/bin/jev`이며 상대 디렉터리도 허용합니다.

- `install`: 실행 중인 바이너리를 원자적으로 복사하고 0755 권한으로 설치합니다. 최신 버전을 다운로드하거나 빌드하지 않습니다. 소스 갱신 후에는 `scripts/`에서 `go run . install --force`를 실행합니다.
- `uninstall`: Go 빌드 정보로 Jev임을 식별할 수 있는 일반 파일만 제거합니다. 없는 대상은 성공이며 임의의 동명 파일·symlink·디렉터리는 보존합니다. **설정·토큰·설치 디렉터리·다른 파일·셸 설정은 삭제하지 않습니다.** 제거한 바이너리는 `scripts/`에서 `go run . install`로 복구할 수 있습니다.
- `doctor`: 설치 파일, PATH에서 선택되는 Jev, 설정과 JSON/pick 조합, 설정파일의 소유자 전용 권한, 토큰 유무를 점검합니다. 설정파일이 없으면 기본값을 사용합니다. 토큰 원문은 출력하지 않습니다.

유지보수 명령은 API를 호출하지 않습니다. doctor는 읽기 전용이며 자동 수정을 하지 않고, 토큰의 실제 인증 유효성을 확인하지 않습니다. 정상 결과는 stdout·종료 코드 0, 문제는 빈 stdout·stderr·종료 코드 1입니다. 일반 판정이나 설치 후에 doctor를 자동 실행하지 않습니다.

## 인증과 설정

설정파일: `~/.agents/jev.toml`.

```toml
api_key = "<TypeSafe API token>"
threshold = 0.8
```

위 토큰은 예시 자리표시자입니다. 실제 토큰을 명령 인자나 대화에 붙여넣지 말고, 로컬 비밀정보 저장 도구 등의 stdout을 다음 명령의 stdin으로 연결합니다.

```sh
jev config set api_key --stdin
```

파일에 저장하지 않으려면 실행 환경의 `TYPESAFE_API_KEY`를 사용합니다. 환경변수가 파일 토큰보다 우선하며 빈 환경변수는 무시합니다. 파일은 0600 권한의 **평문**입니다.

```sh
jev config                         # 실효 설정, 토큰 전체 마스킹
jev config path                    # 파일 경로
jev config set threshold 0.8
jev config set model jev-latest
jev config set timeout 45s
jev config set json true
jev config set pick answer,model
jev config unset threshold
```

| 키 | 내장 기본값 | 저장 형식 |
|---|---|---|
| `api_key` | 없음 | 문자열, CLI 저장은 stdin만 지원 |
| `threshold` | 미적용 | 유한한 0~1 실수 |
| `model` | `jev-latest` | 비어 있지 않은 문자열 |
| `timeout` | `30s` | 양수 Go duration 문자열 |
| `json` | `false` | boolean |
| `pick` | 제한 없음 | 문자열 배열 |

명시적 CLI 옵션 > 파일 > 내장 기본값 순서로 적용합니다. `--threshold 0`, `--json=false`도 파일보다 우선합니다. `--pick ''`는 저장된 필드 선택을 해제합니다. nonempty pick과 JSON 모드의 조합은 판정 실행 시 검사합니다.

파일이나 항목이 없으면 기본값을 사용합니다. 조회·판정은 파일을 만들지 않으며 첫 `set`에서 생성합니다. `set/unset`은 오프라인이며 성공 시 무출력입니다. `unset`은 저장값만 제거합니다. 잘못된 TOML·타입·미지원 key는 오류이며, `jev config path`로 위치를 찾아 직접 수정합니다.

설정 갱신은 원자적 파일 교체로 다른 설정값을 보존하지만 주석·서식은 보존하지 않습니다. symlink·디렉터리를 설정파일로 사용하지 않습니다.

## 판정

```sh
jev --if "이 설정은 운영환경에서 권장하지만 필수는 아니다. 반드시 사용해야 하는가?" \
  --conditions 지지됨,충돌함,판단_불가
```

`--if`에 질문과 필요한 문맥을 함께 넣습니다. `--conditions`는 2~255개의 고유한 선택지이며 앞뒤 공백을 제거합니다. 선택지 내부 쉼표는 지원하지 않습니다.

- 기본 stdout: API가 선택한 최고 확률 답변과 줄바꿈 하나.
- `--json`: `answer`, `probabilities`, `confidence`, 실제 응답 `model`.
- `--json --pick answer,probabilities`: 지정한 최상위 필드만 출력. 한 필드도 객체를 유지합니다.
- `--threshold 0.8`: `probabilities[answer] >= 0.8`인 경우에만 결과를 출력합니다. 미지정은 필터 없음입니다.
- `--model`, `--timeout`: 해당 실행에서 설정값을 덮어씁니다.

`confidence`는 API가 분포에서 계산한 별도 값이며 threshold 비교 대상이 아닙니다. 확률이나 confidence가 실제 정답을 보장하지 않습니다. `--pick`은 CLI 출력 토큰만 줄이며 API 사용량은 줄이지 않습니다.

종료 코드:

- `0`: 성공, 결과는 stdout.
- `2`: threshold 미달, stdout 없음, stderr에 선택지·확률·기준.
- `1`: 입력·설정·인증·API·출력 오류, stderr에 이유.

JSON 모드에서도 실패는 stderr로만 보고합니다. CLI는 자동 재시도하거나 HTTP redirect를 따라가지 않습니다. API 오류 응답 원문은 비밀정보 보호를 위해 출력하지 않습니다.

## 범위와 의존성

현재는 Choice 질문 하나를 처리합니다. Score/Noul, 배치, 문서 탐색, 다른 스킬 자동 연동은 제공하지 않습니다. 서로 연관된 문서를 개별 호출로 나누면 전체 관계를 검증한 결과가 되지 않습니다.

Go HTTP·JSON·flags와 MIT 라이선스의 `github.com/pelletier/go-toml/v2`를 사용합니다. 제3자 라이선스는 [THIRD_PARTY_NOTICES.md](THIRD_PARTY_NOTICES.md)에 있습니다.

외부 계약: [TypeSafe API](https://docs.typesafe.ai/api), [Choice](https://docs.typesafe.ai/primitives/choice).
