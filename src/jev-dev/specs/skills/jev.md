## 사용자 스펙 의도

- "결국 자연어 질문 -> n지선다 답변 도출인거잖아."
- "파일별로 검증하면 전체 맥락에 대한 연관성이 사라진채로 검토될거잖아."
- "그러면 그냥 jev를 cli환경으로 뽑아내는건 어때? 이걸 기반으로 그 다음을 생각하는거지."
- "사실 결과물만 있으면되서 answer만 출력한다로 하는게 어때?"
- "기본 answer는 가장 수치가 높은 값을 선택하게 하고, threshold는 임계값 미달시 stderr로 실패처리하는거지"
- "install uninstall doctor 커맨드도 넣어두자."
- 최초 설치와 소스 갱신은 별도 셸 래퍼 없이 `go run . install`로 수행한다. (사용자 승인)
- "이거 내가 noul choice score 구분을 안했구나. 이것들 구분해서 추가하기 위해 커맨드를 넣어둘까 하는데, 어때?"
- "옵션들 snake_case 적도록 강제하는건? json에서 지원한다해도 golang이었나 거기서 문제있던걸로 기억하거든."
- "시나리오 테스팅 넣기전에, jev에 batch 옵션 넣어버릴까? jsonl 파일 넣으면 일괄로 요청하는방식. status 같은게 유지되면 캐시히트율도 올라갈테니 괜찮은 방법이라 생각되는데"
- "ㅇㅋ 만들어두자."
- "jev는 jev에 대한 정보와 함께 cli 지원을 위한 스킬로 범위를 잡고 하는거지."

- "jev doctor 개선좀하자. api 키 문제가 있었는데, 그걸 doctor 과정에서 볼 수 없으니 불편하네. 이참에 아에 doctor 결과도 보기 좋게 정리좀하자."

- "offline doctor는 제거하자. jev 자체가 api 땡겨다 쓰는건데, 굳이?"

- "간단한 결정을 위한 (사용자 포함)질문은 jev를 통해 빠른 작업을 진행한다."
- "범위가 특정되지 않은 작업은 기본적인 범위와 필요한 조사, 행동 노동력을 보고 jev를 통해 모델을 선택해 subagent를 사용한다."
- "jev-use 스킬 만들어두는거어때? jev는 jev에 대한 정보와 함께 cli 지원을 위한 스킬로 범위를 잡고 하는거지."
- "ㅇㅋ 작업시작해보자. 하는김에 그 사용법도 실제로 적용해서 해보자."
- "effort 질문도 해놔"
- "불확실한 근거는 여러개의 넓은 질문과 좁은질문을 섞어서 실시한다..? 그리고 jev는 질문 갯수 꽤 크게 받을 수 있으니 그것도 포함시켜놓고"
- Jev 사용의 기본 패턴을 설명하고, 모델 선택은 대표 유즈케이스로 안내합니다.
- 기본 사용 패턴과 유즈케이스는 `jev` 스킬 안에서 제공합니다.

- 최신 Jev API의 장문 모드, 구조화 state·instructions·criteria를 CLI 단건·batch에서 사용합니다.

- 기본 설정 섹션은 `[defaults]`이며 provider의 내장 기본값은 jev입니다.
- 설정 파일 최초 생성에는 편집 가능한 기본 템플릿을 제공합니다.

---

# jev 스킬 스펙

## 목적

Jev 정보·CLI 사용과 작업에 적용하는 기본 패턴을 제공합니다. 문맥·질문·기준·후보 구성, primitive와 단건·batch 선택, 결과 해석과 후속 행동 연결을 소유합니다. 작은 선택, 되묻기 판단, 불확실한 근거 평가, 모델 선택은 이 패턴의 내장 유즈케이스입니다.

## 경계

- 포함: Jev·primitive·CLI 정보, 호출 가치 판단, 문맥·질문·기준·후보 구성, 단건·batch 선택, CLI 결과 해석과 실패 처리, 내장 유즈케이스, 요청된 설정·설치·제거·진단.
- 제외: CLI 기계적 동작 재구현, 업무별 정책·모델 정책 소유, 작업 분해·subagent 실행·결과 통합, 모든 질문의 자동 외부 전송, 사용자 취향·승인·권한 대리 결정, scenario-testing의 평가·verdict 계약.

## 처리하려는 작업 형태

- Jev 또는 TypeSafe를 사용한 명제 확률·선택지 판정·순서형 평가, 공통 state에 대한 다중 질문, `config`, `install`, `uninstall`, `doctor` 사용 요청.
- 명시 호출 외에도 현재 작업의 제한된 판단에 Jev 기본 패턴이 유용할 때 적용합니다. 결정적 규칙·기존 증거로 답이 정해지면 직접 진행합니다. 일반 질문이나 코드 검토 전체를 자동으로 외부 API에 넘기지 않습니다.

## 엔트리포인트 / 대표 표면

- `jev/skills/jev/SKILL.md`, `$jev:jev`.
- `jev/skills/jev/references/use-cases.md`: 작업별 적용 예시이며 별도 스킬이나 실행 정책이 아닙니다.
- 실행 파일 `jev`는 PATH에 설치되어 있어야 한다. 실행 실패가 설치나 인증 문제를 나타내면 해당 설정을 안내하며 일반 호출 전에 별도 doctor·연결 확인을 요구하지 않는다.

## 핵심 처리 계약

- 예/아니오 명제는 `jev noul`, 이산 선택은 `jev choice`, 정도나 단계 평가는 `jev score`를 사용한다.
- `--if`에 질문과 판정에 필요한 문맥을 함께 담는다. Choice는 `--conditions`에 고유 선택지를 전달하며 선택지 내부 쉼표는 지원하지 않는다. Score는 순서대로 반복한 `--level`에 고유한 자연어 기준을 전달한다.
- 구조화 단건은 `jev <primitive> --input <path>`의 JSON object로 state·instructions·criteria를 전달합니다. 기존 `--if`·`--conditions`·`--level`과 혼용하지 않습니다. state는 문자열·객체·배열, instructions는 문자열·객체·배열·null입니다. Choice criteria는 label별 설명 object, Score는 순서 있는 문자열·객체·배열 단계, Noul은 true/false 설명 object입니다. Score legend는 구조를 유지합니다. 실제 입력 예시는 CLI help를 확인합니다.
- Batch는 기존 question·conditions·levels 또는 instructions·criteria를 사용합니다. question과 instructions, criteria와 conditions/levels는 각각 함께 사용하지 않습니다. `--state-format json`으로 공통 state의 JSON 구조를 유지하고 기본 text는 원문 문자열을 유지합니다.
- 단건·batch의 `--long-context`는 해당 요청만 확장 컨텍스트를 활성화합니다. 기본 64k 요청·32k 질문 한도를 각각 262k로 확장하며 속도·정확도 저하 가능성을 알립니다. 자동 활성화나 설정 저장은 하지 않습니다. JSON 오류·타입·한도는 호출 전에 검증하고 원문을 오류에 출력하지 않습니다.
- 입력이 여러 문서의 관계에 의존하면 문맥을 함께 보존한다. 입력에 없는 사실을 보충한 판정으로 보고하지 않는다.
- 결과 하나만 필요하면 기본 출력, 부가 필드가 필요하면 JSON과 해당 primitive가 제공하는 pick 필드를 사용한다.
- Noul의 answer는 참일 확률, Choice의 answer는 선택된 라벨, Score의 answer는 단계별 확률을 가중한 점수다. Score는 단계 수에 따라 범위가 달라지고 단계 사이 값이 가능하다.
- threshold는 Choice에서 선택된 라벨의 확률, Noul에서 참일 확률의 하한이다. Score는 별도 `min_score` 기준을 사용한다. confidence는 어느 기준에도 대신 쓰지 않는다.
- CLI·설정·JSON의 소유 key는 `snake_case`를 사용한다. Choice 라벨은 자유 문자열이며 기계적 후속 처리에 쓸 때 `snake_case`를 권장한다. Score 단계는 자연어 설명을 유지한다.
- 기존 `jev --if ... --conditions ...`는 Choice 호환 형식이다.
- 공통 state에 대해 둘 이상의 독립 질문을 평가하면 `jev batch --state-file <path> --input <path>`를 사용한다. JSONL row는 `id`, `type`, question/instructions와 primitive별 기존 기준 또는 criteria를 가집니다.
- Batch는 전체 파일을 먼저 검증하고 한 요청으로 전송하며 입력 순서의 전체 JSONL 결과만 출력한다. 자동 chunk·retry·resume·부분 성공은 제공하지 않는다.
- Batch에는 threshold·min_score·json·pick을 적용하지 않는다. 입력·파일·API·응답 오류는 stdout 없이 exit 1이며 exit 2는 없다.
- CLI 성공 결과로 작업을 진행한다. 별도 API 재호출이나 연결 검사로 같은 결과를 중복 검증하지 않는다.
- 종료 코드 2는 임계값 미달이며 stdout 결과가 아니다. 기준을 조용히 낮춰 재시도하지 않는다. 종료 코드 1은 실행 오류로 보고한다.
- config는 `[defaults]`의 provider·공통값과 `[providers.<name>]`의 model·api_key를 사용합니다. provider는 flag > defaults.provider > 내장 jev, model은 flag > 해당 provider 설정 > 해당 provider 기본값입니다. 인증 환경변수는 jev의 TYPESAFE_API_KEY, openrouter의 OPENROUTER_API_KEY, openai의 OPENAI_API_KEY이며 다른 provider 키를 재사용하지 않습니다. 현재 실행은 Jev만 지원하고 나머지는 설정 저장만 가능합니다.
- `config init`은 내장 템플릿으로 0600 파일을 만들며 기존 파일을 덮어쓰지 않습니다. 최초 `config set`도 같은 템플릿을 사용합니다. config 조회·평가·doctor는 파일을 만들지 않습니다. set/unset은 defaults.<key> 또는 providers.<name>.model/api_key를 사용하며 token은 stdin 전용입니다. 구형 flat 설정은 자동 변환하지 않고 이동할 섹션을 안내합니다. 모든 provider 토큰을 마스킹합니다.
- config 경로·기본값·우선순위·토큰 stdin 저장과 마스킹을 설명한다. 실제 토큰을 대화나 명령 인자에 요청하지 않는다.
- 유지보수 요청에는 `install [--force] [--dir <directory>]`, `uninstall [--dir <directory>]`, `doctor [--dir <directory>] [--provider <name>]`를 사용한다. 기본 대상은 `~/.local/bin/jev`다.
- install은 실행 중인 바이너리를 복사하며, 최신 소스를 빌드하거나 다운로드하지 않는다. 최초 설치는 플러그인의 `scripts/`에서 `go run . install`, 소스 갱신은 `go run . install --force`를 사용한다. 다른 설치 위치는 `--dir <directory>`로 지정한다.
- uninstall은 Jev로 식별되는 일반 파일만 제거하며 설정·토큰·설치 디렉터리·셸 설정을 보존한다. 없는 대상은 성공이다.
- doctor는 설치·PATH·설정·키 유무·출처와 API 인증을 점검한다. 기본은 설정·키가 유효할 때 인증된 models 조회를 한 번 수행하며 추론·redirect·retry·자동 수정은 없다. `OK / FAIL / SKIP`, 원인·해결 안내·요약을 읽고 인증 실패와 네트워크·서버 문제를 구분한다. 키 원문을 출력하지 않으며 모델 정확성·추론 가능 여부를 증명하지 않는다. 정상은 stdout·exit 0, 문제는 빈 stdout·stderr·exit 1이다.
- 질문과 선택지가 TypeSafe 외부 API로 전송됨을 밝힌다. 승인 범위 밖의 파일이나 비밀정보를 입력으로 수집하지 않는다.

## 기본 사용 패턴

- 먼저 요청·규칙·확인 가능한 증거를 확인합니다. Jev는 제한된 판단이 다음 행동을 바꿀 때 사용하며, 빠른 직접 확인으로 알 수 있는 사실을 추정시키지 않습니다.
- 목표, 확인된 사실, 중요 미지수, 제약, 판단 기준, 허용 후보를 구성합니다. 필요한 경우 `inspect_more`, `ask_user`, `insufficient`, `not_shown` 같은 실제 복구·유보 선택지를 둡니다.
- 서로 종속된 질문은 앞선 답을 확인한 뒤 후속 호출합니다. 같은 비어 있지 않은 state에서 독립적으로 답할 수 있는 질문만 batch로 묶습니다.
- Jev는 보조 판단입니다. 답·확률·confidence는 사실·권한·사용자 의사·품질의 증명이 아닙니다. 원문과 현재 정책을 확인한 뒤 결과를 행동에 반영합니다.
- CLI 부재·실행 실패·기준 미달 시 설치·설정 변경, 기준 완화, 동일 호출 반복을 자동 수행하지 않습니다. 확인된 기본값이나 자체 판단으로 진행하고, 사용자만 결정할 수 있는 내용은 묻습니다. 실제 응답·실패·대체 판단을 구분해 보고합니다.

## 내장 유즈케이스

아래 사례는 공통 패턴의 적용 예시입니다. 업무별 절차·위임·모델 정책은 현재 작업과 적용 지침이 소유합니다.

- 작은 A/B 선택: 사용자 목표·제약에 맞는 후보와 사용자 확인 필요 여부를 평가합니다.
- 되묻기 판단: 현재 정보로 진행, 추가 조사, 사용자 질문을 비교합니다. 사용자 취향·승인은 대신 정하지 않습니다.
- 불확실한 근거: 확인된 출처·발췌·주장·공백을 공통 state로 둡니다. 서로 다른 넓은 질문으로 전체 관계·대안 해석을, 좁은 질문으로 개별 주장·반례·누락 전제를 평가합니다. 같은 state에서 독립적이면 함께 batch로 묻습니다. 질문 수를 작은 고정값으로 제한하지 않되 현재 API 요청·문맥 한도와 전송 비용 안에서 행동에 필요한 만큼만 묻습니다. 근거 부재와 반증을 구분하며, 답의 충돌은 원문 재확인으로 해결합니다. 반복 질문·다수결·평균은 독립 근거를 만들지 않습니다.
- 모델 선택: 먼저 작업의 기본 범위·조사량·실행량·실패 영향과 허용된 위임 경계를 확인합니다. 필요하면 직접 수행·위임·추가 조사·사용자 확인을 Jev에 물을 수 있습니다. 모델·effort 쌍을 정한 프로필 정책이면 허용 프로필을 후보로 평가합니다. 개별 선택이 허용되면 현재 사용 가능한 모델을 먼저 평가하고, 선택한 모델의 지원·허용 effort를 후속 질문으로 평가합니다. 정해진 정책·기본값이 답을 정하면 그대로 사용합니다. 결과는 해당 작업의 권고이며 영구 설정·spawn 권한을 만들지 않습니다.

## 검토 질문

- 질문에 맞는 primitive와 필요한 문맥·기준을 사용했는가?
- 여러 질문이 같은 state를 공유할 때 batch를 사용하고, 서로 다른 state를 억지로 한 batch에 합치지 않았는가?
- 참일 확률, 선택 확률, 가중 점수, confidence와 실행 실패를 구분하는가?
- 설정·설치·외부 전송 권한을 요청 범위 안에 유지하는가?
- 공통 패턴과 업무별 예시를 구분하고, 후보·모델·effort가 현재 정책을 따르는가?
- 불확실한 근거의 넓고 좁은 질문이 독립적이며, 답의 충돌을 원문에서 확인하는가?

## 독립성 원칙

- 독립 실행 가능성을 강제한다. 스킬 본문과 번들 reference만으로 CLI·기본 패턴·유즈케이스·실패 처리를 이해할 수 있어야 합니다.
- sibling skill이나 개발 spec을 runtime 필수 문맥으로 요구하지 않는다. 플러그인 번들 목적과 설치 안내는 README·manifest가 소유한다.

## 확장 원칙

- 실제 지원 명령과 재사용 가능한 Jev 사용 패턴을 안내합니다. 업무별 사례를 실행 정책이나 오케스트레이션 책임으로 승격하지 않습니다.
