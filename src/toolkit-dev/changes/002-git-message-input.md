# Toolkit 002: 메시지 입력과 Git 경로

## 변경사항 요약

- `message create --stdin`은 quoted heredoc 입력을 저장하고 자동 validation 후 짧은 `MSG-…` ID를 반환합니다.
- 신규 파일은 Git이 발견한 현재 repository·worktree의 metadata 디렉터리에 생성합니다.
- 명확한 push 결과는 추가 remote 조회 없이 사용합니다.

## 변경 상세

### 생성 시 입력·검증

- 목적: 별도 파일 편집과 긴 임시 경로 전달을 줄입니다.
- 범위: `--stdin`에서 EOF까지 읽고 mode `0600` 파일에 저장합니다. validation 성공 전에는 ID를 출력하지 않습니다. 실패 시 생성한 파일의 identity·안전 조건을 확인하고 정리하며, 정리 실패는 보존 경로와 함께 보고합니다.
- 보조 경로: 기존 인자 없는 `create`는 빈 파일을 할당합니다. 이후 수정·readback·명시적 validation이 필요합니다.
- 관련 표면: Go executable, skill spec/runtime, README·manifest.
- 검증: quoted heredoc의 literal 보존, 빈 입력·잘못된 message 거부, 실패 파일 정리, commit 미실행을 확인합니다.

### Git 메타데이터와 짧은 식별자

- 목적: 작업 디렉터리와 linked worktree 구조에 관계없이 Git의 경로 계약을 사용합니다.
- 범위: `git rev-parse --absolute-git-dir`로 canonical metadata 디렉터리를 찾습니다. 일반 checkout에서는 `.git/MSG-…`이며 stdout은 `MSG-…` ID만 반환합니다. `validate`·`commit`은 같은 repository·worktree에서 ID를 해석합니다.
- 경계: 다른 repository·worktree 파일, 안전하지 않은 owner·mode·symlink, Git context 부재는 거부합니다. 성공 commit만 같은 identity의 파일을 정리하고 commit 실패는 보존합니다.
- 출력: `create`는 ID를 반환합니다. 파일을 직접 편집하는 호출자는 Git directory에 상대적으로 해석하고, `validate`·`commit`에는 반환값을 그대로 전달합니다.
- 검증: root·nested directory·linked worktree·별도 Git directory·bare repository와 commit 성공·실패 lifecycle을 확인합니다.

### 결과에 비례한 확인

- 목적: 도구가 이미 제공한 명확한 결과를 중복 조회하지 않습니다.
- 범위: push의 종료 상태와 출력에서 예상 destination과 success·rejection·no-op이 명확하면 추가 status·fetch·remote-ref 조회를 생략합니다. 잘못된 호출, 중단, 불완전·상충 출력, 불명확한 결과에 필요한 확인만 수행합니다.
- 유지: mutation 전 범위·권한 판단, 내부 message 안전성 검증, 일반 commit 출력에 없는 저장된 full message 검증, 부분 성공 보존.
- 관련 표면: git skill spec/runtime, recovery reference, plugin usage.

## 비목표

- 새 CLI가 stage·push·branch mutation을 자동 실행하지 않습니다.
- 사용자 설치본, PATH, shell profile, Git config를 자동 갱신하지 않습니다.
- 버전 변경·릴리즈를 포함하지 않습니다. commit·push·사용자 설치본 갱신은 별도 명시 요청에서 수행합니다.

## 검증 기준

- Go 통합 테스트와 vet, 4개 기존 플랫폼의 executable 빌드를 수행합니다.
- spec/runtime 독립 검증, skill validator, JSON parsing, reference와 diff 검사를 수행합니다.
- 로컬 개발 산출물 검증과 사용자 설치·원격 게시 여부를 구분해 보고합니다.

## 검증 결과

- 2026-09-19 macOS arm64: Go 통합 테스트·vet, skill validator, JSON·reference·diff 검사 통과했습니다. heredoc, 10개 invalid-input 분기, CRLF 보존·단독 CR 거부, Git layout, 읽기 실패·identity 교체 cleanup, commit·hook 실패 보존, linked worktree의 nested commit을 확인했습니다.
- darwin/linux × amd64/arm64 빌드 통과했습니다. 실행 검증은 macOS arm64이며 다른 플랫폼은 cross-build 검증입니다. 사용자 설치본 갱신과 원격 게시 결과는 해당 실행 결과로 별도 확인합니다.
