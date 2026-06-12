# Advance Codex Dev

`advance-codex-dev`는 Codex에서 할 수 있는 일을 더 깊고 안정적으로 활용하기 위한 플러그인입니다.
재사용 가능한 skill, installable plugin bundle, change finalization, reusable instruction evaluation, agent token optimization 같은 Codex 활용 체계를 설계하고 정리하는 작업을 위한 문서화와 가이드를 제공합니다.
`.agents/sessions/{YYYYMMDD}`는 session-scoped operational artifact를 두는 기본 위치로만 문서화합니다. 실제 active flow 기록, `000-plan.md`, `001-*` flow record 운영은 `loop-kit:turn-gate`가 소유합니다.

이 플러그인은 Codex 활용 방식을 더 명시적이고 유지보수 가능하게 만드는 데 목적이 있습니다.
반대로 일반적인 실행 workflow나 무관한 공용 유틸리티를 담는 용도로 넓히지 않습니다.

`skill-creator`는 재사용 가능한 Codex skill을 설계하거나 기존 skill의 경계, trigger metadata, runtime 본문을 정리해야 할 때 사용합니다.
`plugin-creator`는 설치 가능한 plugin bundle의 경계, manifest, README, plugin spec, bundled skill 관계를 정리해야 할 때 사용합니다.
`skill-scenario-testing`은 reusable instruction을 fresh subagent와 고정 시나리오로 검증하고 evidence 중심으로 분석해야 할 때 사용합니다.
`git-committer`는 작업을 task-scoped commit으로 마무리할 때 커밋 준비, 실행 권한 확인, staged 검증, 메시지, commit 실행을 묶어 처리할 때 사용합니다.
`optimize-token`은 에이전트 응답, 진행 보고, 상태 문구, 검증·승인 문구의 토큰 사용을 줄이되 정확성, 검증 결과, 승인 경계, 필수 출력 형식, 현재 상태 기준은 유지해야 할 때 사용합니다.
