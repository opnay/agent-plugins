## 사용자 스펙 의도

- 직접 만든 변경을 task-scoped commit으로 깔끔하게 마무리하고 싶다.
- staged scope 검토와 수동 검증을 거친 뒤 commit하고 싶다.
- commit message 형식과 body 품질을 같이 통제하고 싶다.
- `commit-readiness-gate`는 readiness 판단을 맡고, 실제 commit finalization은 `git-committer`가 맡아야 한다.
- `git-committer`는 readiness, 검증 통과, handoff만으로 commit 실행 승인을 추정하지 않아야 한다.
- `git-committer` 스펙은 folder-based 구조로 두고, intent에는 readiness에서 commit 실행까지의 흐름도가 보여야 한다.
- `git-committer`의 흐름도는 커밋 준비, 커밋 실행 권한, 커밋 실행 흐름을 분리해서 보여야 한다.

## 전체 흐름도

```mermaid
flowchart LR
  A[사용자 요청 작업]

  subgraph G[git-committer]
    B[커밋 준비] --> C[커밋 실행 권한] --> D[커밋 실행]
  end

  A --> B
```

## 커밋 준비 흐름도

```mermaid
flowchart TD
  A[요청 작업]

  subgraph B[커밋 준비]
    C[프로젝트의 커밋 준비 단계] --> D[커밋할 범위 선택]
  end

  E[커밋 실행 권한]

  A --> C
  D --> E
```

## 커밋 실행 권한 흐름도

```mermaid
flowchart TD
  A[커밋 준비]

  subgraph B[커밋 실행 권한]
    C[사용자 커밋 실행 승인 확인]
  end

  D[커밋 실행]

  A --> C
  C --> D
```

## 커밋 실행 흐름도

```mermaid
flowchart TD
  A[커밋 실행 권한]

  subgraph B[커밋 실행]
    C[staged 검증] --> D[커밋 메시지 준비] --> E[커밋 생성]
  end

  F[결과 확인]

  A --> C
  E --> F
```
