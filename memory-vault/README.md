# Memory Vault

`memory-vault`는 사용자가 직접 제공한 지식 저장소 폴더 자체를 관리하고, 그 폴더의 `AGENTS.md`에 관리 규칙을 추가하는 플러그인입니다.

## 사용 예시

```text
이 폴더를 memory vault로 관리해 주세요: /path/to/vault
```

```text
현재 저장소의 지식 저장소 문서를 정리하고 AGENTS.md 규칙도 맞춰 주세요.
```

## 제공 기능

- `manage-memory-vault`: 사용자가 제공한 폴더를 vault root로 삼아 기본 문서와 `AGENTS.md` 규칙 섹션을 관리합니다.

## 기본 산출물

```text
<target-folder>/
├── AGENTS.md
├── README.md
├── INDEX.md
├── decisions.md
├── glossary.md
├── open-questions.md
└── Programming/
    ├── INDEX.md
    └── React/
        └── INDEX.md
```

## 구조

```text
memory-vault/
  .codex-plugin/plugin.json
  README.md
  skills/manage-memory-vault/SKILL.md
  scripts/memv.py
  scripts/tests/test_memv.py
```
