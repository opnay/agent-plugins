# SW Kit test value

## 변경사항 요약

- `code` 본문은 테스트 작성 gate, junk patterns, 유지 예외, regression 확인을 소유한다.
- 조건부 reference는 집중 감사와 전체 subsystem campaign을 소유한다.
- `maintenance`의 테스트 부채 조사는 동일한 `code` 기준을 사용한다.

## 변경 상세

### 목적과 범위

테스트가 독립적으로 검출하는 실패와 계약을 기준으로 작성·보존·통합을 판단한다. 일상 기준은 본문에 두어 테스트 작성마다 감사 절차를 읽지 않는다. 기존 네 skill 책임과 public identifier를 유지한다.

### 기준과 적용

- 작성 gate, 전체 junk categories, 유지 예외, deletion evidence, keeper, R/F/C/D ledger, 보존 review·mutation 기준을 유지한다.
- 실행 명령·tool·PR 흐름은 대상 저장소의 검증·권한·Git 정책을 따른다.
- 관련 표면: code·maintenance spec/runtime, code references, plugin spec·README·manifest.

### 비목표와 호환성

새 skill, runner, executable test, 설치, 버전 변경, commit·push·release는 포함하지 않는다. 기존 호출은 유지하며 감사·campaign reference는 해당 작업에서만 로드한다. 테스트 수·LOC·속도 개선을 보장하지 않는다.

### 검증 기준

- `quick_validate.py`로 실제 code·maintenance skill을 확인한다.
- JSON 파싱, resource·plugin identifier·marketplace 경로, diff hygiene를 확인한다.
- clean-context read-only verifier로 본문 계약과 조건부 절차·plugin 경계를 각각 확인한다.
- fresh agent가 가상 코드 snapshot의 작성·감사 요청에 적용한 판단과 reference 로드를 확인한다. 실제 repository 감사·테스트 실행·파일 수정 권한은 제공하지 않는다.

### 로컬 검증 결과

- code·maintenance `quick_validate.py`, manifest·marketplace JSON, UI metadata·resource·identifier·runtime language 검증, `git diff --check`: 통과.
- clean-context 작성 기준 검토: 기존 decision guide의 테스트 편의 wrapper 표현을 발견해 실제 dependency·lifecycle 경계로 보정했다. 감사·campaign·maintenance·plugin surface 검토는 불일치 없음.
- fresh authoring trial: 테스트 전용 helper export·같은 helper로 만든 expected 제안은 보류, 정확한 소스 문자열 검사는 동작 검증으로 수정 권고, 구별되는 경계값은 기존 table 확장으로 유지했다. `code/SKILL.md`만 읽었고 파일 수정·테스트 실행은 하지 않았다.
- fresh audit trial: 중복 aggregate와 identifier 결합 source check는 삭제 후보, 구별되는 입력 표와 shipped config asset 계약은 유지로 판정했다. 본문·`test-audit.md`만 읽고 campaign은 읽지 않았다. 실제 삭제·테스트 실행은 하지 않았고 필요한 검증을 미실행으로 보고했다.
- 현재 spec/runtime을 비교한 추가 clean-context 검토: 핵심 기준의 유의미한 누락·충돌 없음. 보정된 wrapper guide도 실제 경계와 test-only 편의를 구분함을 확인했다.
- 실제 코드베이스의 테스트 감소·실행 속도·대규모 campaign 실행은 검증하지 않았다. 설치·commit·push·release는 수행하지 않았다.
