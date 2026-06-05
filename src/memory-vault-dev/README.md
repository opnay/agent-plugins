# Memory Vault Dev

`memory-vault-dev`는 에이전트 사용 전반에서 반복적으로 참고할 장기 기억을 개인 vault root에 저장하고 관리하는 플러그인입니다.

## 사용 예시

```text
이 내용은 앞으로 기억해 주세요: 답변은 한국어 존대체로 해 주세요.
```

```text
내 에이전트 메모리 vault를 초기화해 주세요.
```

```text
Agents/Prompting 카테고리를 만들고 관련 기억을 정리해 주세요.
```

```text
Programming/React/hook-rules 지식 문서를 만들고 인덱스에 연결해 주세요.
```

## 제공 기능

- `memory-vault`: 개인/작업 장기 메모리 vault를 초기화하고, 사용자 선호, 결정, 환경, workflow, 용어, 미해결 질문, 번호가 붙은 지식 문서를 분류해 관리합니다.

## 기본 산출물

```text
<target-folder>/
├── AGENTS.md
├── README.md
├── INDEX.md
├── preferences.md
├── decisions.md
├── environment.md
├── workflows.md
├── glossary.md
├── open-questions.md
├── Agents/
│   ├── INDEX.md
│   └── Prompting/
│       └── INDEX.md
└── Programming/
    ├── INDEX.md
    └── React/
        ├── INDEX.md
        └── 001-hook-rules.md
```

기본 vault root는 `~/Workspace/Memory-vault`입니다.
카테고리 `INDEX.md`는 지식 본문을 담지 않고, 같은 폴더의 `<index>-<lowercase-hyphen-slug>.md` 문서와 하위 카테고리 링크를 나열합니다.

## 설치 산출물 구조

```text
memory-vault/
  .codex-plugin/plugin.json
  README.md
  skills/memory-vault/SKILL.md
  scripts/memv.py
```
