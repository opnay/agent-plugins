## 사용자 스펙 의도

- 신규 플러그인 추가. 폴더 위치를 지정하면, 해당 위치에 지식 저장소 문서 폴더를 만들어 관리합니다. 스킬은 항상 사용되도록 passive 셋팅이 필요합니다. 폴더의 AGENTS.md를 만들고 관리하는 규칙을 추가합니다.
  - 새 플러그인 이름을 무엇으로 할까요?
    - `memory-vault`[직접 입력]
- 정정: 사용자는 직접 관리하는 폴더를 제공합니다. 해당 폴더 그 자체를 관리하는 스킬입니다.

---

# manage-memory-vault 스킬 스펙

## 목적

`manage-memory-vault`는 사용자가 직접 제공한 폴더 자체를 vault root로 관리하고, 해당 폴더의 기본 문서와 `AGENTS.md` 지식 저장소 관리 규칙을 추가하거나 갱신합니다.
이 skill은 자연어 요청에서 passive로 선택될 수 있어야 하며, 사용자가 플러그인 이름이나 skill 이름을 직접 말하지 않아도 repository memory 성격의 요청이면 적용됩니다.

## 경계

- 포함:
  - 대상 폴더 확인
  - 대상 폴더의 `README.md`, `INDEX.md`, `decisions.md`, `glossary.md`, `open-questions.md` 생성
  - 1-2 depth 카테고리 폴더의 `INDEX.md` 생성
  - `AGENTS.md`의 `Memory Vault` 표식 섹션 생성 또는 갱신
  - root `AGENTS.md`에 1-2차 카테고리 맵 갱신
  - 기존 문서가 있으면 보존하고 빠진 기본 문서만 추가
  - 변경 전 dry-run 또는 변경 예정 파일 설명
- 제외:
  - 대상 폴더 밖 파일 변경
  - 기존 vault 문서 삭제 또는 임의 재작성
  - 대상 폴더 아래 별도 `memory-vault/` 하위 폴더 생성
  - 표식 없는 기존 `AGENTS.md` 본문 재작성
  - 프로젝트 지식 내용의 자동 생성 또는 사실 확정

## 처리하려는 작업 형태

- 폴더 경로를 지정하며 해당 폴더 자체를 memory vault로 관리하라는 작업
- 현재 작업 폴더에 지식 저장소를 만들라는 작업
- `AGENTS.md`에 지식 저장소 규칙을 추가하라는 작업
- 기존 memory vault의 기본 문서 누락 여부를 확인하고 보완하는 작업
- 기존 vault 문서를 읽고 `INDEX.md`, `decisions.md`, `glossary.md`, `open-questions.md`에 맞게 지속 지식을 정리하는 작업
- `Programming/React`처럼 1-2 depth 카테고리를 만들고 각 폴더의 `INDEX.md`로 내용을 관리하는 작업

## 엔트리포인트 / 대표 표면

- 대표 표면: `skills/manage-memory-vault/SKILL.md`
- 호출 방식: 자연어 passive trigger 또는 `$memory-vault:manage-memory-vault`
- helper script: `scripts/memv.py`

## 핵심 처리 계약

- 사용자가 지정한 대상 폴더가 명확하지 않으면 먼저 질문합니다.
- 대상 폴더가 명확하면 해당 폴더만 변경 대상으로 잠급니다.
- 기본 변경은 대상 폴더의 기본 문서, 1-2 depth 카테고리 `INDEX.md`, `AGENTS.md`로 제한합니다.
- 기존 파일은 덮어쓰지 않고, `AGENTS.md`의 표식 섹션만 갱신합니다.
- 스크립트를 사용할 수 있으면 `--dry-run`으로 예정 변경을 확인한 뒤 실제 실행합니다.
- 실행 뒤에는 생성/갱신/보존된 파일과 검증 결과를 보고합니다.

## Passive Trigger 계약

- description은 skill 이름 없이도 선택될 수 있도록 한국어와 영어 트리거 표현을 포함합니다.
- passive token은 넓은 일반 문서 작업 전체가 아니라 folder-local memory, knowledge repository, AGENTS.md memory rules에 맞춥니다.
- 사용자가 "항상 사용"을 원한 의미는 관련 요청에서 자동 선택되도록 description trigger를 강화하는 것입니다.

## AGENTS.md 규칙 계약

- `AGENTS.md`에는 `<!-- memory-vault:start -->`와 `<!-- memory-vault:end -->` 표식으로 감싼 섹션을 둡니다.
- 표식 섹션은 agent가 작업 전 `INDEX.md`를 확인하고, 중요한 결정은 `decisions.md`, 용어는 `glossary.md`, 미해결 질문은 `open-questions.md`에 기록하도록 지시합니다.
- 표식 섹션은 root `AGENTS.md` 안에 1-2차 카테고리 맵을 둡니다.
- 표식 없는 기존 지침은 보존합니다.
- 기존 표식 섹션이 있으면 현재 템플릿으로 갱신할 수 있습니다.

## 검토 질문

- 사용자가 직접 제공한 vault root 폴더가 어느 경로인지 명확한가?
- 사용자가 일반 문서 정리를 원한 것인지, 지속 지식 저장소 생성을 원한 것인지 구분되는가?
- 기존 `AGENTS.md`가 있으면 표식 섹션만 변경하는가?
- 변경 대상이 대상 폴더 내부 allowlist에 머무르는가?
- 카테고리 경로가 상대 경로이며 1-2 depth를 넘지 않는가?

## 독립성 원칙

- 이 skill이 독립 실행 가능성을 spec으로 강제해야 하는가: 예.
- 이유: 설치 후 자연어 passive trigger만으로도 단독 실행되어야 하며, sibling skill 없이 대상 폴더와 helper script만으로 작업을 완료해야 합니다.

## 확장 원칙

- 새 문서 종류를 추가할 때는 plugin spec의 bundle 경계와 `AGENTS.md` 규칙 계약을 함께 점검합니다.
- 외부 동기화나 데이터베이스 기능은 별도 plugin 또는 명시적 확장 spec 없이는 추가하지 않습니다.
- 기본 문서 템플릿은 vault root와 category `INDEX.md`의 가독성과 낮은 유지 비용을 우선합니다.
