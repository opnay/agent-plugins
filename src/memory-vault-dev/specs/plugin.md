## 사용자 스펙 의도

- 신규 플러그인 추가. 폴더 위치를 지정하면, 해당 위치에 지식 저장소 문서 폴더를 만들어 관리합니다. 스킬은 항상 사용되도록 passive 셋팅이 필요합니다. 폴더의 AGENTS.md를 만들고 관리하는 규칙을 추가합니다.
  - 새 플러그인 이름을 무엇으로 할까요?
    - `memory-vault`[직접 입력]
- 정정: 사용자는 직접 관리하는 폴더를 제공합니다. 해당 폴더 그 자체를 관리하는 스킬입니다.
- 스크립트 이름은 `memv.py`로 바꾸고, 2-depth 정도의 카테고리 인덱싱을 추가합니다. 예: `Programming - React`. 각 폴더별로 `INDEX.md`를 만들고, root `AGENTS.md`에는 1-2차 카테고리 맵을 둡니다.
- skill 이름은 플러그인 이름과 같은 `memory-vault`를 사용합니다.

---

# Memory Vault Dev 플러그인 스펙

## 플러그인 목적

`memory-vault-dev`는 사용자가 직접 제공한 폴더를 지식 저장소 root로 관리하는 플러그인입니다.
핵심 책임은 대상 폴더 자체의 기본 지식 문서, 1-2 depth 카테고리 `INDEX.md`, `AGENTS.md` 관리 규칙을 함께 제공해, 이후 에이전트가 해당 폴더의 누적 지식을 일관되게 읽고 갱신하도록 만드는 것입니다.

## 플러그인 경계와 비목표

- 포함:
  - 사용자가 제공한 폴더 자체를 vault root로 관리
  - 1-2 depth 카테고리 폴더와 각 폴더의 `INDEX.md` 관리
  - 기본 지식 문서 생성과 기존 문서 보존
  - 같은 폴더의 `AGENTS.md` 생성 또는 표식 섹션과 카테고리 맵 갱신
  - passive trigger 문구를 통한 자연어 요청 자동 선택
- 제외:
  - 외부 지식 베이스, 클라우드 저장소, 데이터베이스 동기화
  - 사용자가 지정하지 않은 폴더 탐색 또는 대량 변경
  - 대상 폴더 아래에 별도 `memory-vault/` 하위 폴더 생성
  - 기존 문서 삭제, 기존 `AGENTS.md` 전체 재작성
  - 프로젝트별 지식 내용의 사실성 검증 자동화

## 처리하려는 작업 형태

- "이 폴더를 지식 저장소로 관리해 달라"는 요청
- "repository memory", "memory vault", "persistent notes" 같은 지속 지식 문서 요청
- 특정 폴더의 `AGENTS.md`에 지식 저장소 관리 규칙을 추가하거나 갱신하는 요청
- 이미 만들어진 vault root 문서를 읽고 다음 에이전트가 참고할 요약, 결정, 질문을 정리하는 요청
- `Programming/React`처럼 1-2 depth 카테고리 인덱스를 추가하거나 갱신하는 요청

## 대표 표면

- 대표 스펙: `memory-vault-dev/specs/plugin.md`
- skill 상세 스펙 위치: `memory-vault-dev/specs/skills/memory-vault.md`
- runtime skill: `memory-vault-dev/skills/memory-vault/SKILL.md`
- helper script: `memory-vault-dev/scripts/memv.py`

## 내장 skill 체계

- `memory-vault`: 지정 폴더 자체의 기본 지식 문서, 카테고리 `INDEX.md`, `AGENTS.md` 규칙 섹션을 생성, 갱신, 검증한다.
  - spec: `memory-vault-dev/specs/skills/memory-vault.md`

## SDD 운영 원칙

- plugin spec은 폴더-local 지식 저장소라는 bundle 경계와 사용 표면을 소유합니다.
- skill spec은 실제 파일 변경 allowlist, 질문 기준, 검증 기준, passive trigger 계약을 소유합니다.
- runtime skill은 설치 후 접근 가능한 `SKILL.md`와 `scripts/memv.py`만 실행 지시로 참조합니다.
- skill 책임이 문서 구조나 AGENTS.md 규칙을 바꾸면 plugin spec, skill spec, README, manifest prompt를 같은 변경에서 점검합니다.

## 현재 구조 메모

- 초기 버전은 단일 skill 플러그인입니다.
- 대상 폴더 자체를 vault root로 유지하고, 하위 `memory-vault/` 폴더를 만들지 않습니다.
- `AGENTS.md` 변경은 표식 섹션만 갱신하는 allowlist 방식으로 유지합니다.
