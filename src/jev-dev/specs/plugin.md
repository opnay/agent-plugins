## 사용자 스펙 의도

- "그러면 그냥 jev를 cli환경으로 뽑아내는건 어때? 이걸 기반으로 그 다음을 생각하는거지."
- "golang으로 제작하자."
- "설정파일 관리할 수 있는 커맨드 추가하자. `jev config`"
- "`~/.agents/jev.toml`로 설정파일 만들자."
  - "옵션으로 threshold같은 기본값 지정할 수 있는 것들 추가하자. (물론 미지정시 jev cli의 기본값 사용)"
  - "typesafe ai의 api token 지정할 수 있게 해놓자."
- "install uninstall doctor 커맨드도 넣어두자."
- 최초 설치와 소스 갱신은 별도 셸 래퍼 없이 `go run . install`로 수행한다. (사용자 승인)
- "이거 내가 noul choice score 구분을 안했구나. 이것들 구분해서 추가하기 위해 커맨드를 넣어둘까 하는데, 어때?"
- "옵션들 snake_case 적도록 강제하는건? json에서 지원한다해도 golang이었나 거기서 문제있던걸로 기억하거든."
- "이제 더이상 codex 만의 일은 아니게되어서 플러그인 성격과 달라졌어. jev 플러그인으로 옮기는거 어때?"
- "시나리오 테스팅 넣기전에, jev에 batch 옵션 넣어버릴까? jsonl 파일 넣으면 일괄로 요청하는방식. status 같은게 유지되면 캐시히트율도 올라갈테니 괜찮은 방법이라 생각되는데"
- "ㅇㅋ 만들어두자."

---

# Jev 플러그인 스펙

## 플러그인 목적

TypeSafe의 hosted Jev 모델을 Go CLI에서 호출하고, 에이전트가 Noul·Choice·Score 판단과 reusable instruction scenario evaluation을 명시적으로 수행할 수 있게 한다.

## 플러그인 경계와 비목표

- 포함: Noul·Choice·Score 단건·batch API 호출, 결과 기준 필터, 출력 필드 선택, 로컬 설정 관리, CLI 설치·제거·진단, primitive 사용 스킬, reusable instruction scenario evaluation.
- 제외: 로컬 모델, 문서 자동 순회, 서로 다른 state의 자동 grouping, Jev와 무관한 generic test framework, target instruction 직접 수정, 일반 implementation debugging, 다른 플러그인 자동 연동, MCP 서버.
- Jev CLI로 전달한 질문·선택지는 외부 API로 전송한다. `scenario-testing`은 전송 권한이 없으면 prediction을 건너뛴다. 플러그인 설치는 토큰 설정이나 실행 파일의 PATH 설치를 대신하지 않는다.

## 처리하려는 작업 형태

- 질문과 필요한 문맥을 `--if`에 함께 전달하고, 판단 형태에 맞는 primitive를 고른다.
- Noul은 참일 확률, Choice는 선택된 조건, Score는 순서형 기준의 가중 점수를 받는다.
- `jev batch`는 하나의 공통 state와 JSONL 질문을 한 SystemOne 요청으로 평가한다.
- `jev config`로 기본값과 API 인증을 관리한다.
- `jev install`, `jev uninstall`, `jev doctor`로 요청된 CLI 유지보수를 수행한다.
- reusable instruction의 scenario-specific expected behavior, 선택적 Jev prediction, fresh executor actual behavior를 비교하고 evidence를 보존한다.

## 대표 표면

- CLI: [cli.md](cli.md), `jev/scripts/`의 독립 Go 모듈. 대표 명령은 `jev noul`, `jev choice`, `jev score`, `jev batch`다.
- 스킬: `$jev:jev`, [skills/jev.md](skills/jev.md), `$jev:scenario-testing`, [skills/scenario-testing.md](skills/scenario-testing.md).
- 사용 안내: `jev/README.md`, manifest의 설명과 `defaultPrompt`.
- 설치: `jev install`, 제거: `jev uninstall`, 오프라인 진단: `jev doctor`. 기본 대상은 `~/.local/bin/jev`다.
- 최초 설치는 `jev/scripts/`에서 `go run . install`로 수행한다. 소스 갱신은 `go run . install --force`를 사용한다. 셸 설정·토큰은 설치나 제거 대상이 아니다.

## 내장 skill 체계

- `jev`: primitive 선택, 단건·batch CLI 입력 구성, 타입별 결과 사용을 소유한다. 추론 실행·설정 저장·종료 코드는 CLI가 소유한다.
- `scenario-testing`: reusable instruction의 고정 scenario, behavior Choice, 선택적 Jev prediction, fresh execution, expected·predicted·actual 비교, caller verdict, evidence 기록을 소유한다.
- 업무별 scenario 정책을 `jev` primitive 사용 스킬에 흡수하지 않는다. `scenario-testing`은 `$jev:jev`의 공개 Choice 계약만 사용하며 CLI 내부 구현에 의존하지 않는다.

## SDD 운영 원칙

- CLI 계약과 skill 계약을 구분한다. CLI 변경 시 사용 안내·스킬·manifest의 노출 약속을 함께 확인한다.
- 새 primitive나 일괄 처리 기능은 실제 요구가 생기면 API 및 출력 계약부터 정의한다.
- CLI·설정·JSON의 소유 key는 `snake_case`를 사용한다. 사용자 자연어와 Choice 라벨은 Go 식별자 제약을 받지 않는다.
- scenario evaluation 계약은 별도 skill spec이 소유한다. CLI는 선택적 prediction을 위한 공통 state 다중 질문 전송까지 소유하고, scenario catalog·executor·comparison·evidence orchestration은 `scenario-testing`이 소유한다.

## 현재 구조 메모

- runtime: `jev/`, 개발 문서: `src/jev-dev/`.
- marketplace의 `jev` 항목은 `./jev`를 가리키며 기존 항목 뒤에 위치한다.
