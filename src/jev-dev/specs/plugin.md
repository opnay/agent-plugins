## 사용자 스펙 의도

- "그러면 그냥 jev를 cli환경으로 뽑아내는건 어때? 이걸 기반으로 그 다음을 생각하는거지."
- "golang으로 제작하자."
- "설정파일 관리할 수 있는 커맨드 추가하자. `jev config`"
- "`~/.agents/jev.toml`로 설정파일 만들자."
  - "옵션으로 threshold같은 기본값 지정할 수 있는 것들 추가하자. (물론 미지정시 jev cli의 기본값 사용)"
  - "typesafe ai의 api token 지정할 수 있게 해놓자."
- "install uninstall doctor 커맨드도 넣어두자."
- 최초 설치와 소스 갱신은 별도 셸 래퍼 없이 `go run . install`로 수행한다. (사용자 승인)

---

# Jev 플러그인 스펙

## 플러그인 목적

TypeSafe의 hosted Jev 모델을 Go CLI에서 호출하고, 에이전트가 질문 하나와 명시적 선택지로 판정 결과를 사용할 수 있게 한다.

## 플러그인 경계와 비목표

- 포함: Choice API 호출, 선택 확률 기반 threshold, 출력 필드 선택, 로컬 설정 관리, CLI 설치·제거·진단, 사용 스킬.
- 제외: 로컬 모델, 문서 자동 순회, 일괄 평가, 특정 검증 업무 자동화, 다른 플러그인 자동 연동, MCP 서버.
- 질문·선택지는 외부 API로 전송한다. 플러그인 설치는 토큰 설정이나 실행 파일의 PATH 설치를 대신하지 않는다.

## 처리하려는 작업 형태

- 질문과 필요한 문맥을 `--if`에 함께 전달하고 `--conditions`의 선택지 중 하나를 받는다.
- `jev config`로 기본값과 API 인증을 관리한다.
- `jev install`, `jev uninstall`, `jev doctor`로 요청된 CLI 유지보수를 수행한다.

## 대표 표면

- CLI: [cli.md](cli.md), `jev/scripts/`의 독립 Go 모듈.
- 스킬: `$jev:jev`, [skills/jev.md](skills/jev.md).
- 사용 안내: `jev/README.md`, manifest의 설명과 `defaultPrompt`.
- 설치: `jev install`, 제거: `jev uninstall`, 오프라인 진단: `jev doctor`. 기본 대상은 `~/.local/bin/jev`다.
- 최초 설치는 `jev/scripts/`에서 `go run . install`로 수행한다. 소스 갱신은 `go run . install --force`를 사용한다. 셸 설정·토큰은 설치나 제거 대상이 아니다.

## 내장 skill 체계

- `jev`: CLI 입력 구성과 결과 사용을 소유한다. 추론 실행·설정 저장·종료 코드는 CLI가 소유한다.
- 특정 업무의 판단 정책이나 전체 검토 절차를 이 스킬에 흡수하지 않는다.

## SDD 운영 원칙

- CLI 계약과 skill 계약을 구분한다. CLI 변경 시 사용 안내·스킬·manifest의 노출 약속을 함께 확인한다.
- 새 primitive나 일괄 처리 기능은 실제 요구가 생기면 API 및 출력 계약부터 정의한다.

## 현재 구조 메모

- runtime: `jev/`, 개발 문서: `src/jev-dev/`.
- marketplace의 `jev` 항목은 `./jev`를 가리키며 기존 항목 뒤에 위치한다.
