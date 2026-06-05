## 사용자 스펙 의도

- 신규 플러그인 추가. 폴더 위치를 지정하면, 해당 위치에 지식 저장소 문서 폴더를 만들어 관리합니다. 스킬은 항상 사용되도록 passive 셋팅이 필요합니다. 폴더의 AGENTS.md를 만들고 관리하는 규칙을 추가합니다.
  - 새 플러그인 이름을 무엇으로 할까요?
    - `memory-vault`[직접 입력]
- 정정: 사용자는 직접 관리하는 폴더를 제공합니다. 해당 폴더 그 자체를 관리하는 스킬입니다.
- 스크립트 이름은 `memv.py`로 바꾸고, 2-depth 정도의 카테고리 인덱싱을 추가합니다. 예: `Programming - React`. 각 폴더별로 `INDEX.md`를 만들고, root `AGENTS.md`에는 1-2차 카테고리 맵을 둡니다.
- skill 이름은 플러그인 이름과 같은 `memory-vault`를 사용합니다.
- 지식 저장하는 스킬은 만들었는데, 전략을 뭘로 가져가야될까? 이 플러그인을 설치한 에이전트에게 뭔가를 시키면 자연스레 습득하는 지식들을 모아 관리하는 방식을 생각하고있거든.
- 이건 프로젝트 타겟이 아니야. 어떤 상황에서든 에이전트를 사용하다보면 기억해야되는 내용들과 지식들을 저장하는 스킬이야.
- `INDEX.md`는 정보를 나열하는 인덱싱 파일로 유지합니다. 예를 들어 `Programming/001-some-knowledge.md`, `Programming/React/001-hook-rules.md` 같은 지식 문서를 폴더별 `INDEX.md`에서 나열합니다. 지식 문서 파일명은 `<index>-<slug|lowercase|hyphen-case>.md` 형식입니다.

---

# Memory Vault Dev 플러그인 스펙

## 플러그인 목적

`memory-vault-dev`는 에이전트 사용 전반에서 반복적으로 참고할 개인/작업 지식을 하나의 vault root에 저장하고 관리하는 플러그인입니다.
핵심 책임은 에이전트가 작업 중 확인한 장기 기억 후보를 선별하고, 사용자 선호, 운영 결정, 환경 정보, 반복 workflow, 용어, 미해결 질문으로 나누어 낮은 비용으로 누적하게 만드는 것입니다.

## 플러그인 경계와 비목표

- 포함:
  - 사용자가 제공한 폴더 또는 기본 개인 vault root를 장기 메모리 저장소로 관리
  - 기본 메모리 문서와 1-2 depth 카테고리 `INDEX.md` 생성
  - 카테고리별 지식 문서 생성과 `INDEX.md` 문서 목록 갱신
  - vault root의 `AGENTS.md`에 메모리 읽기/쓰기 규칙과 카테고리 맵 관리
  - 에이전트 작업 시작 전 관련 메모리 읽기
  - 에이전트 작업 종료 전 재사용 가능한 기억 후보 선별
  - 확실한 장기 지식만 작은 단위로 기록
  - passive trigger 문구를 통한 자연어 요청 자동 선택
- 제외:
  - 특정 프로젝트 전용 지식 저장소로 한정
  - 외부 지식 베이스, 클라우드 저장소, 데이터베이스 동기화
  - 대화 전체나 작업 로그의 무차별 저장
  - 사용자 확인이 필요한 모호한 사실의 자동 확정
  - 민감 정보, 비밀값, 자격 증명 저장
  - 기존 문서 삭제 또는 기존 `AGENTS.md` 전체 재작성

## 처리하려는 작업 형태

- "이 내용을 기억해", "앞으로 이렇게 해", "내 선호로 저장해" 같은 장기 기억 요청
- 에이전트가 작업 중 반복 적용 가능한 사용자 선호, 환경, workflow, 용어, 운영 결정을 발견한 경우
- `~/Workspace/Memory-vault` 같은 개인 vault root를 초기화하거나 보완하는 요청
- 이미 만들어진 vault root 문서를 읽고 현재 작업에 필요한 기억을 적용하는 요청
  - `Agents/Prompting`, `Tools/Codex`, `Coding/TypeScript`처럼 1-2 depth 카테고리 인덱스를 추가하거나 갱신하는 요청
  - `Programming/React/hook-rules`처럼 1-2 depth 카테고리 아래 번호가 붙은 지식 문서를 만들거나 목록에 연결하는 요청

## 대표 표면

- 대표 스펙: `memory-vault-dev/specs/plugin.md`
- skill 상세 스펙 위치: `memory-vault-dev/specs/skills/memory-vault.md`
- runtime skill: `memory-vault-dev/skills/memory-vault/SKILL.md`
- helper script: `memory-vault-dev/scripts/memv.py`

## 내장 skill 체계

- `memory-vault`: 개인/작업 장기 메모리 vault를 초기화, 읽기, 분류, 갱신, 검증한다.
  - spec: `memory-vault-dev/specs/skills/memory-vault.md`

## SDD 운영 원칙

- plugin spec은 전역 개인 에이전트 메모리라는 bundle 경계와 사용 표면을 소유합니다.
- skill spec은 실제 파일 변경 allowlist, 기억 후보 판단 기준, 질문 기준, 검증 기준, passive trigger 계약을 소유합니다.
- runtime skill은 설치 후 접근 가능한 `SKILL.md`와 `scripts/memv.py`만 실행 지시로 참조합니다.
- skill 책임이 메모리 문서 구조나 `AGENTS.md` 규칙을 바꾸면 plugin spec, skill spec, README, manifest prompt를 같은 변경에서 점검합니다.

## 현재 구조 메모

- 초기 버전은 단일 skill 플러그인입니다.
- 대상 폴더 자체를 vault root로 유지하고, 하위 `memory-vault/` 폴더를 만들지 않습니다.
- 기본 운영 대상은 개인 vault root이며, 특정 프로젝트 폴더를 대상으로 삼는 것은 사용자가 명시한 경우에만 허용합니다.
- `AGENTS.md` 변경은 표식 섹션만 갱신하는 allowlist 방식으로 유지합니다.
- 카테고리 `INDEX.md`는 요약 문서가 아니라 해당 폴더의 지식 문서와 하위 카테고리 링크를 나열하는 인덱스입니다.
