## 사용자 스펙 의도

- Notion-backed agent memory를 `advance-codex-dev` 안에 `notion-memory` skill로 통합하고 싶다.
- 공개 skill 이름은 `notion-memory`로 유지하고, `$advance-codex-dev:notion-memory setup`과 `$advance-codex:notion-memory setup`을 대표 호출로 삼고 싶다.
- 개인 Notion DB URL, DB ID, data source, view 같은 workspace-specific 값은 plugin에 하드코딩하지 않고 setup 때 사용자에게 받아 `~/.agents/configs/notion-memory.toml`에 저장하고 싶다.
- setup 요청은 literal command뿐 아니라 “셋업하자”, “스킬을 사용하기 위해 준비해줘”, “Notion 메모리 기록 쓸 수 있게 준비해줘” 같은 의미론적 요청도 같은 흐름으로 처리하고 싶다.
- 기록 대상은 작업 히스토리, 결정, 후속 작업, 검증, 재사용 가능한 작업 지식으로 제한하고 싶다.
- 설정은 `general`과 `notion` 영역으로 나누고 `general.timezone`, `notion.db_url` 같은 이름으로 다루고 싶다.
- 기록 문서는 한국어, 서울 기준, 일관된 존댓말, 간결하지만 다음 에이전트가 실행 가능한 형식으로 작성하고 싶다.
- Notion memory의 문서 속성 값은 고정 계약으로 유지하되, 본문 형식은 기본 템플릿 강제가 아니라 작성 전 기록 목적에 맞게 직접 정할 수 있는 자유 형식으로 두고 싶다.
  - 본문 형식은 누가 언제 정하도록 고정할까요?
    - 에이전트 결정 [무응답으로 권장값 적용]: 기록 전에 에이전트가 기록 목적에 맞는 형식을 정하고, 애매할 때만 사용자에게 묻는다.

---

# notion-memory 스킬 스펙

## 목적

`notion-memory`는 Codex 작업 메모리를 Notion DB에 연결해 setup, schema 준비, workspace rule 준비, 기록, 검증을 수행하는 skill입니다.
핵심은 agent work memory를 재사용 가능하게 남기는 것이며, 일반 Notion workspace 자동화가 아닙니다.

## 경계

- 포함:
  - Notion-backed agent work memory setup
  - config 파일 생성과 갱신 안내
  - config 파일 생성/검증 스크립트 제공
  - memory DB schema 확인, 추가, 수동 fallback 안내
  - workspace rule 안내 또는 사용자 승인 후 적용
  - 작업 히스토리, 결정, 후속 작업, 검증, 재사용 가능한 작업 지식 기록
  - relation/link, KST `기록일`, title, 고정 property, body structure 선택 계약
  - memory record 문서 작성 규약
  - allow-list first, block-list security review
- 제외:
  - 일반 Notion workspace 정리
  - 사용자 승인 없는 schema 파괴 변경
  - credential, token, Notion auth state 저장
  - 모든 대화의 자동 저장
  - unverified claim의 사실화
  - commit, push, PR, release
  - plugin에 개인 DB URL/ID 하드코딩

## 처리하려는 작업 형태

- 사용자가 Notion 메모리 skill을 쓰기 위해 준비해 달라고 요청한 경우
- 사용자가 작업 히스토리, 결정, 후속 작업, 검증 결과, 재사용 가능한 작업 지식을 Notion memory DB에 남기려는 경우
- 사용자가 Notion memory DB schema, config, workspace rule, setup verification을 점검하려는 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `advance-codex-dev/skills/notion-memory/SKILL.md`
- dev 호출: `$advance-codex-dev:notion-memory setup`
- release 호출: `$advance-codex:notion-memory setup`
- 의미론적 setup 요청:
  - `$advance-codex:notion-memory 셋업하자`
  - `$advance-codex:notion-memory 스킬을 사용하기 위해 준비해줘`
  - “Notion 메모리 기록 쓸 수 있게 준비해줘”

## 핵심 처리 계약

- setup 요청은 `setup`, `setup config`, `setup schema`, `setup workspace-rule`, `setup verify` command family로 라우팅한다.
- `setup`은 `setup config > setup schema > setup workspace-rule > setup verify` 순서로 처리한다.
- 기록 요청은 먼저 scope를 map한다:
  - user scope: 사용자가 기록하라고 한 내용
  - Notion scope: config에 연결된 memory DB
  - local scope: config, workspace rule, 임시 notes
  - excluded scope: unrelated DB, unrelated project, commit/push/PR/release
- allow-list를 먼저 적용한다:
  - confirmed work history
  - decisions
  - follow-up tasks
  - verification results
  - reusable work knowledge
  - explicit user preferences
- block-list를 좁은 방어선으로 적용한다:
  - credentials, tokens, auth state
  - 불필요한 personal data
  - speculation, unsupported interpretation
  - unrelated conversation
  - private Notion DB IDs copied into plugin files
- config는 `~/.agents/configs/notion-memory.toml`에 저장한다.
- config는 `general`과 `notion` section으로 나눈다.
- `general` section에는 `timezone`, `language`, `locale`, `register`를 저장한다.
- `notion` section에는 `db_url`, `db_id`, `data_source_id`, `view_id`, `schema_status`를 저장한다.
- config에는 credential, token, cookie, connector auth state를 저장하지 않는다.
- setup config 자동화가 필요하면 bundled `scripts/setup_config.py`를 사용한다.
- `scripts/setup_config.py`는 config 생성, template 출력, 검증만 소유하고 Notion DB schema 변경이나 page write는 소유하지 않는다.
- Notion connector가 있으면 connector로 fetch/update/create를 시도한다.
- connector가 없거나 schema update가 실패하면 browser UI 또는 manual fallback을 보고한다.
- schema 변경은 additive missing-property setup을 기본으로 하고, destructive 변경은 사용자 승인 없이는 하지 않는다.
- 기록 title은 `PREFIX: 짧은 의미 제목` 형식을 사용하고 날짜/시간은 넣지 않는다.
- `기록일`은 KST datetime source of truth로 사용한다.
- page property 이름, 타입, 허용 값, 의미는 고정 계약으로 유지한다.
- 본문 형식은 기록을 작성하기 전에 기록 목적에 맞게 정한다.
- 기본 본문 구조는 선택 가능한 시작점이며 필수 템플릿이 아니다.
- 본문 heading과 순서는 바꿀 수 있지만 확인된 사실, 결정, 후속 작업, 검증, 출처는 구분 가능해야 한다.
- 본문 형식 선택이 기록 해석, 사용자 의도, 계약 준수, handoff 가능성에 영향을 주면 작성 전에 사용자에게 확인한다.
- 문서는 `general.language`, `general.locale`, `general.register`를 따르고 기본값은 한국어, `ko-KR`, 존댓말이다.
- memory record는 간결하되 다음 에이전트가 실행 가능해야 한다.
- 사실, 결정, 후속 작업, 검증을 섞지 않는다.
- 장식적 문장이나 근거 없는 해석을 추가하지 않는다.
- relation은 같은 memory DB 내부 record에만 사용하고, 외부 URL은 link property에 둔다.
- Notion write 후에는 tool output만으로 끝내지 않고 필요한 property/page fetch로 확인한다.

## Setup Command 계약

- `setup`: 전체 setup flow를 수행한다.
- `setup config`: config 존재 여부를 확인하고, 없으면 사용자에게 DB URL 또는 ID를 요청한 뒤 script 또는 수동 TOML로 생성한다.
- `setup schema`: configured DB를 fetch하고 required property를 확인하거나 additive setup을 수행한다.
- `setup workspace-rule`: 현재 workspace에서 이 skill을 사용할 rule이 필요한지 확인하고, 사용자 승인 범위 안에서만 추가한다.
- `setup verify`: config, DB 접근, schema, optional test write 여부를 확인한다.

## 기록 계약

- 기록은 사용자가 요청한 recordable scope 안에서만 만든다.
- Notion page property는 고정 계약대로 채우고, 본문 형식 자유를 property schema 변경으로 해석하지 않는다.
- 기록 작성 전 본문 구조를 먼저 정한다.
- confirmed fact, decision, follow-up, verification, source를 구분한다.
- 확인되지 않은 항목은 `확인 필요` 또는 `미검증`으로 표시한다.
- 같은 분 내 여러 record가 필요하면 Notion 표시 정렬을 위해 `기록일` 분 값을 분리한다.
- reusable work knowledge는 raw log 전체가 아니라 다음 실행에 쓸 수 있는 규칙으로 압축한다.

## 문서 작성 규약

- 기본 문서 언어는 한국어다.
- 날짜, 시간, 지역 맥락은 서울 기준으로 쓴다.
- 한 문서 안에서 존댓말, 해요체, 반말을 섞지 않는다.
- 기록은 짧게 쓰되 다음 에이전트가 같은 작업을 이어갈 수 있을 만큼 실행 가능해야 한다.
- 확인된 사실, 결정, 후속 작업, 검증 상태를 분리한다.
- 기본 템플릿을 쓰지 않더라도 다음 에이전트가 사실, 결정, 후속 작업, 검증, 출처를 식별할 수 있어야 한다.
- 장식적 코멘트, 추측, 근거 없는 해석은 쓰지 않는다.
- 직접 완료 기록만 필요한 상황에서는 불필요한 진행 해설을 넣지 않는다.

## 검토 질문

- setup 요청을 command family 중 하나로 라우팅했는가?
- config에 credential이나 개인 auth state를 저장하지 않았는가?
- plugin 파일에 개인 Notion URL/ID를 복사하지 않았는가?
- 기록 내용이 allow-list에 해당하고 block-list를 통과했는가?
- 본문 형식을 작성 전에 정했고, 기본 템플릿을 필수 형식으로 강제하지 않았는가?
- 본문 형식 자유를 property schema나 허용 값 변경으로 오해하지 않았는가?
- schema 변경이 additive인지, destructive이면 사용자가 승인했는지 확인했는가?
- Notion write 후 fetch 또는 동등한 증거로 검증했는가?
- connector 실패 시 manual/browser fallback을 구체적으로 보고했는가?
- setup config에서 스크립트를 쓸 수 있으면 썼고, 스크립트 미사용 또는 실패 사유를 남겼는가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 예.
- 그렇다면 왜 필요한가 / 아니라면 어떤 sibling context를 허용하는가: `notion-memory`는 설치된 runtime skill만으로 setup과 기록을 수행해야 한다. sibling skill이나 dev-only spec 경로 없이도 config, schema, 기록, 검증 계약을 이해할 수 있어야 한다.

## 확장 원칙

- DB property나 body structure 선택 계약이 바뀌면 `notion-memory` skill spec과 runtime reference를 함께 갱신한다.
- plugin usage surface는 README, plugin spec, manifest prompt에서 갱신한다.
- Notion connector-specific 실패 대응은 runtime reference에 두고 개인 DB 식별자는 넣지 않는다.
