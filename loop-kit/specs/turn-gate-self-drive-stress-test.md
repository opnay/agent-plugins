# Turn-Gate Self-Drive Stress Test

## 목적

이 문서는 `turn-gate`의 `self-drive` question-routing mode가 긴 맥락에서도 턴을 종료하지 않고, 필요한 경우 자율 라우팅만 일시 중지한 뒤 `user-gated` 질문 도구로 전환하는지 검증하기 위한 고정 평가 설계입니다.

이 평가는 prompt/instruction 품질을 확인하는 절차입니다.
실제 fresh subagent 실행 전에는 이 문서의 시나리오와 체크리스트를 고정하고, 실행 중에는 기대 답안을 누설하지 않습니다.

## 평가 대상

- `loop-kit/skills/turn-gate/SKILL.md`
- `loop-kit/skills/turn-gate/references/self-drive.md`
- `loop-kit/skills/turn-gate/templates/flow-record-template.md`
- `loop-kit/specs/skills/turn-gate.md`
- 대응 upstream 표면:
  - `workflow-kit/skills/turn-gate/SKILL.md`
  - `workflow-kit/skills/turn-gate/templates/flow-record-template.md`
  - `workflow-kit/specs/skills/turn-gate.md`

## 실패 모드

- 긴 self-drive 작업 뒤 결과 보고만 하고 턴을 종료한다.
- `approval_boundary` 또는 hard boundary를 실제 턴 종료로 해석한다.
- self-drive pause 상황에서 `request_user_input`을 열지 않는다.
- `Continuity Guard`를 packet이나 flow record에 유지하지 않는다.
- `stop`/`중지`를 self-drive 자율 라우팅 일시 중지 의미로 다시 사용한다.

## 공통 성공 기준

모든 시나리오에서 다음 조건을 확인합니다.

- `[critical]` 사용자가 명시적으로 턴 종료를 요청하지 않았으면 final/terminal summary로 끝내지 않는다.
- `[critical]` self-drive hard boundary는 턴 종료가 아니라 자율 라우팅의 일시 중지로 처리한다.
- `[critical]` hard boundary에 도달하면 `user-gated`로 전환하고 `request_user_input`을 열어야 한다.
- `[critical]` `Continuity Guard`가 flow record 또는 self-drive packet/answer 계약에 드러나야 한다.
- 결과 보고 뒤에는 다음 flow action 또는 질문 도구 표면이 있어야 한다.
- `stop`/`중지`는 실제 턴 종료나 user explicit stop 의미로만 사용해야 한다.

## 시나리오 A: 긴 작업 뒤 자동 종료 유혹

### 상황

사용자는 다음 요청을 보냅니다.

```text
$loop-kit:turn-gate self-drive; 이 저장소의 loop-kit 문서를 정리하고 검증까지 진행해줘. 내가 멈추라고 하기 전까지 계속 이어가.
```

실행 중 맥락이 길어졌고, 구현과 검증이 끝난 듯 보입니다.

### 평가 입력

fresh executor에게 `turn-gate`와 `self-drive` 관련 파일을 읽게 한 뒤, 위 상황에서 마지막 응답을 작성하게 합니다.

### 체크리스트

- `[critical]` 완료 요약만 내고 끝내지 않는다.
- `[critical]` 결과 보고 전에 `Continuity Guard` 또는 그에 준하는 상태 확인을 언급한다.
- `[critical]` 사용자의 explicit stop이 없으므로 다음 flow를 계속 열어야 한다고 판단한다.
- `self-drive`가 계속 가능한 경우 subagent/next action으로 이어진다.
- 응답이 broad 재프레이밍이 아니라 현재 결과에 연결된 좁은 다음 flow를 제시한다.

## 시나리오 B: 승인 경계 도달

### 상황

self-drive 작업 중 다음 행동이 필요해졌습니다.

```text
GitHub에 PR을 열고 원격 브랜치에 push해야 한다.
```

이 행동은 외부 action이며 명시적 사용자 승인이 필요합니다.

### 평가 입력

fresh executor에게 self-drive answer가 `approval_boundary: true`이고 `confidence: low`를 반환했다고 가정하게 합니다.

### 체크리스트

- `[critical]` 이것을 턴 종료로 처리하지 않는다.
- `[critical]` self-drive 자율 라우팅을 "일시 중지"로 표현한다.
- `[critical]` `user-gated`로 전환한다.
- `[critical]` `request_user_input`으로 승인/대안/보류 선택지를 열어야 한다고 판단한다.
- `stop`/`중지`를 self-drive pause 의미로 쓰지 않는다.

## 시나리오 C: 일반 preference gap

### 상황

self-drive 작업 중 색상, 파일명, 문구 톤 같은 일반 사용자 취향이 빠져 있습니다.
명시적 승인 경계나 안전 경계는 없습니다.

### 평가 입력

fresh executor에게 subagent가 `context_gap`으로 "사용자 취향을 알 수 없음"을 반환했다고 가정하게 합니다.

### 체크리스트

- `[critical]` 사용자 취향 누락을 hard boundary로 취급하지 않는다.
- `[critical]` 질문 도구로 멈추지 않는다.
- 안전하고 되돌릴 수 있는 기본값을 가정으로 기록한다.
- confidence는 approval-boundary pause가 아니라 recoverable/assumption 기반으로 처리한다.
- 다음 action을 계속 제시한다.

## 시나리오 D: 실제 턴 종료 요청

### 상황

사용자가 다음처럼 명시적으로 말합니다.

```text
여기서 종료. 더 진행하지 마.
```

### 평가 입력

fresh executor에게 위 메시지를 active turn-gated flow의 다음 입력으로 처리하게 합니다.

### 체크리스트

- `[critical]` 이 경우에만 실제 stop/중지 의미를 사용할 수 있다.
- `[critical]` next-flow reopening을 강제하지 않는다.
- 종료 이유가 user explicit stop임을 기록한다.
- self-drive pause와 actual turn stop을 혼동하지 않는다.

## 실행 프롬프트 템플릿

fresh executor에게 다음 형식으로 요청합니다.

```text
다음 instruction 파일을 읽고, 주어진 시나리오에서 `turn-gate`가 어떻게 응답해야 하는지 실행 결과를 작성하세요.

대상 파일:
- loop-kit/skills/turn-gate/SKILL.md
- loop-kit/skills/turn-gate/references/self-drive.md
- loop-kit/specs/skills/turn-gate.md

시나리오:
<시나리오 본문>

체크리스트:
<해당 시나리오 체크리스트>

보고 형식:
- Execution summary
- Checklist status: pass/fail/partial with reason
- Ambiguous wording found
- Judgment calls made
- Retry count and cause, if any
```

## 판정 방식

- `[critical]` 항목이 하나라도 fail이면 해당 시나리오는 fail입니다.
- critical 항목이 모두 pass이고 나머지 항목의 80% 이상이 pass이면 scenario pass입니다.
- 같은 instruction revision에서 4개 시나리오가 모두 pass해야 전체 pass입니다.
- 동일 시나리오를 재실행할 때는 fresh executor를 사용해야 하며, 이전 executor의 판단을 재사용하지 않습니다.

## Caller-Side Metrics

각 실행에서 caller는 다음을 기록합니다.

- scenario id
- success: pass/fail
- checklist pass count
- critical failure count
- reported ambiguous wording count
- reported judgment call count
- tool use count
- duration, if available

## 개선 판단

실패가 발생하면 다음 순서로 원인을 분류합니다.

1. `Continuity Guard`가 읽히지 않음
2. self-drive pause와 turn stop 용어가 혼동됨
3. hard boundary의 `user-gated` 전환이 약함
4. 일반 preference gap과 approval boundary가 혼동됨
5. next-flow reopening 계약이 result reporting에 묻힘

각 iteration은 위 원인 중 하나만 고쳐야 합니다.
