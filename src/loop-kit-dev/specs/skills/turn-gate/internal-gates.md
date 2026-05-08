# turn-gate internal-gates sub-spec

## 목적

이 문서는 `turn-gate` 내부의 gate 전환 계약을 소유합니다.
gate는 외부에 노출되는 별도 skill이나 phase가 아니라, 현재 입력과 flow 상태가 다음 전환 조건을 만족하는지 확인하는 내부 운영 단위입니다.

## 핵심 모델

`turn-gate`는 layer를 사용자 표면으로 노출하지 않습니다.
대신 내부 gate를 통과하며 다음 행동을 결정합니다.

- message intake gate
- flow shaping gate
- task policy gate
- verification gate
- reporting gate
- continuation gate

각 gate는 자기 전환 조건만 소유합니다.
어떤 gate도 사용자의 explicit stop 없이 terminal closure를 추론하지 않습니다.

## Message Intake Gate

message intake gate는 현재 incoming user message의 라우팅 사실을 분류합니다.

소유:

- 현재 메시지가 explicit turn stop인지 판정한다.
- explicit stop이 아니라면 continuation input으로 분류한다.
- 새 작업, correction, review, status check, approval-sensitive request, handoff result, next-task request 같은 message kind를 파악한다.
- operation/target ambiguity, scope gap, approval boundary 가능성을 표시한다.

비소유:

- active flow completion criteria 확정
- 파일 수정, 검증 실행, build, commit, push, PR, publish
- flow 완료나 turn closure 결정

명시적 stop이 아닌 메시지는 terminal close 근거가 아닙니다.

## Flow Shaping Gate

flow shaping gate는 message intake 결과를 active flow에 반영합니다.

소유:

- 새 flow를 만들지, 기존 flow를 갱신할지, reporting이나 question-routing으로 이어갈지 결정한다.
- flow kind를 `operational-preparation`, `change-unit`, user-gated handoff, reporting context, question-routing state 중 적절한 형태로 정한다.
- flow boundary, completion criteria, verification expectation, next-flow reopening 조건을 기록한다.
- 후속 `change-unit` 후보와 active execution flow를 구분한다.

비소유:

- flow 내부 command sequence 실행
- verification pass 판정
- explicit stop 없는 turn closure

flow shaping gate는 task 목록을 phase checklist나 planned flow list로 오해하지 않게 막아야 합니다.

## Task Policy Gate

task policy gate는 선택된 flow 안에서 실행 정책을 정합니다.
task policy는 flow 밖의 독립 계층이 아니라 flow 내부 정책입니다.

소유:

- 현재 flow를 달성하기 위한 구체 task sequence를 정한다.
- 필요한 local reference, target reread, command, edit, build, test, subagent handoff 조건을 정한다.
- commit execution 같은 approval-sensitive action이 flow completion criteria 안에서 허용됐는지 확인한다.

비소유:

- flow boundary 재정의
- verification pass 확정
- reporting 생략
- next-flow reopening 생략
- turn closure 결정

개별 task 완료는 flow 완료나 turn closure를 결정할 수 없습니다.
예를 들어 `git commit` 명령이 성공해도 commit flow는 commit hash 확인, working tree 상태 확인, reporting, next-flow reopening 조건을 아직 만족해야 할 수 있습니다.

## Verification Gate

verification gate는 flow work 결과가 보고 가능한 상태인지 판정합니다.

소유:

- flow의 completion criteria와 verification expectation을 기준으로 검증 packet을 구성한다.
- clean-context verifier 결과를 `pass`, `fail`, `blocked`, `insufficient` 중 하나로 통합한다.
- `fail` 또는 `insufficient`이면 earliest safe gate로 되돌린다.
- `blocked`이면 user-gated question-routing으로 blocker를 보고하게 한다.

비소유:

- 사용자 대신 approval boundary 승인
- next-flow 선택
- terminal closure

## Reporting Gate

reporting gate는 flow 결과를 다음 진행을 위한 맥락으로 정리합니다.

소유:

- 준비, 작업, 검증, 남은 불확실성, blocker, approval boundary 상태를 보고한다.
- material routing judgment call을 드러낸다.
- reporting 전에 session record와 Continuity Guard를 갱신한다.

비소유:

- terminal summary 허용
- explicit stop 없는 clean stop
- next-flow reopening 생략

보고는 turn closure가 아닙니다.

## Continuation Gate

continuation gate는 reporting 이후의 전환을 결정합니다.

소유:

- source-recorded explicit stop이 있는지 확인한다.
- explicit stop이 없으면 active question-routing, loop continuation, planned self-drive handoff, blocker decision 중 하나로 이어간다.
- `request_user_input` 사용 가능성과 fallback 필요성을 판단한다.
- session record의 `Next Flow Options`에 explicit turn-end option을 남긴다.

비소유:

- 새로운 work를 질문 없이 확장
- commit/push/PR/publish 승인 추론
- stale closure 또는 source-less closure를 terminal close 근거로 사용

continuation gate는 task 결과와 reporting 결과를 turn 종료로 해석하지 않습니다.

## Review Questions

- message intake gate가 실행이나 flow completion을 직접 결정하지 않는가?
- flow shaping gate가 active flow와 후속 후보를 분리하는가?
- task policy gate가 flow 내부 정책으로 남고 별도 planned flow처럼 쓰이지 않는가?
- task 완료가 reporting이나 continuation gate를 건너뛰지 않는가?
- reporting이 terminal closure가 아니라 continuation context로 쓰이는가?
