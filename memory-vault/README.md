# Memory Vault

`memory-vault`는 에이전트 사용 전반에서 반복적으로 참고할 장기 기억을 개인 vault root에 저장하고 관리하는 플러그인입니다.

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

## 제공 기능

- `memory-vault`: 개인/작업 장기 메모리 vault를 초기화하고, 사용자 선호, 결정, 환경, workflow, 용어, 미해결 질문을 분류해 관리합니다.

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
└── Agents/
    ├── INDEX.md
    └── Prompting/
        └── INDEX.md
```

기본 vault root는 `~/Workspace/Memory-vault`입니다.

## 구조

```text
memory-vault/
  .codex-plugin/plugin.json
  README.md
  skills/memory-vault/SKILL.md
  scripts/memv.py
  scripts/tests/test_memv.py
```
