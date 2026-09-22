# Jev

TypeSafe의 hosted Jev 모델을 Noul·Choice·Score 단건 또는 batch로 호출하는 Go CLI와 reusable instruction의 행동을 평가하는 스킬입니다.

- `$jev:jev`: primitive·단건·batch 선택, 입력 구성, 결과 사용, 요청된 설정·설치·제거·진단.
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
jev doctor
```

- `install`: 실행 중인 바이너리를 0755 권한으로 복사합니다. 다른 일반 파일 교체에는 `--force`가 필요합니다.
- `uninstall`: Go build info로 Jev임을 식별한 일반 파일만 제거합니다. 설정·토큰·디렉터리·셸 설정은 유지합니다.
- `doctor`: 설치 파일, PATH, 일반 설정 유효성·권한, 토큰 유무를 오프라인으로 점검합니다. API 인증은 확인하지 않습니다.

## 인증과 설정

설정파일은 `~/.agents/jev.toml`입니다.

```sh
jev config set api_key --stdin
jev config
jev config path
jev config set threshold 0.8
jev config set min_score 1.5
jev config set model jev-latest
jev config set timeout 45s
jev config set json true
jev config set pick answer,model
jev config unset threshold
```

실제 토큰은 stdin으로 저장합니다. 파일에 저장하지 않으려면 `TYPESAFE_API_KEY`를 사용합니다. 비어 있지 않은 환경변수가 파일의 `api_key`보다 우선합니다. 설정파일은 0600 권한의 평문입니다.

| key | 내장 기본값 | 적용 |
|---|---:|---|
| `api_key` | 없음 | 판정 인증 |
| `threshold` | 미적용 | Choice 선택 확률·Noul 참일 확률 |
| `min_score` | 미적용 | Score 점수 |
| `model` | `jev-latest` | 모든 primitive |
| `timeout` | `30s` | 모든 primitive |
| `json` | `false` | 출력 |
| `pick` | 제한 없음 | primitive별 JSON field |

명시적 CLI 옵션 > 파일 > 내장 기본값 순서로 적용합니다. `--threshold 0`, `--min-score 0`, `--json=false`, `--pick ''`도 저장값을 덮어씁니다. 조회와 판정은 설정파일을 만들지 않으며 `set`과 `unset`은 오프라인입니다.

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

단건 명령은 호출당 질문 하나를 처리하며, batch는 하나의 공통 state와 여러 질문을 한 요청으로 처리합니다. 서로 다른 state의 자동 grouping, 문서 탐색, 업무별 정책, 다른 플러그인 자동 연동은 CLI가 제공하지 않습니다. `$jev:scenario-testing`은 batch prediction과 executor 실행을 조율합니다. 연관된 문서를 단순히 개별 호출로 나누는 것만으로는 전체 관계를 검증한 결과가 되지 않습니다.

Go HTTP·JSON·flags와 MIT 라이선스의 `github.com/pelletier/go-toml/v2`를 사용합니다. 제3자 라이선스는 [THIRD_PARTY_NOTICES.md](THIRD_PARTY_NOTICES.md)에 있습니다.

외부 계약: [TypeSafe API](https://docs.typesafe.ai/api), [Noul](https://docs.typesafe.ai/primitives/noul), [Choice](https://docs.typesafe.ai/primitives/choice), [Score](https://docs.typesafe.ai/primitives/score).
