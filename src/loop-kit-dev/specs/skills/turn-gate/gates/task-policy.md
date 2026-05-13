# turn-gate task-policy gate sub-spec

## 목적

이 문서는 task policy gate의 전환 계약을 소유합니다.

task policy gate는 선택된 flow 안에서 실행 정책을 정합니다.
task policy는 flow 밖의 독립 계층이 아니라 flow 내부 정책입니다.

## 소유

- 현재 flow를 달성하기 위한 구체 task sequence를 정한다.
- 필요한 local reference, target reread, command, edit, build, test, subagent handoff 조건을 정한다.
- commit execution 같은 approval-sensitive action이 flow completion criteria 안에서 허용됐는지 확인한다.

## 비소유

- flow boundary 재정의
- verification pass 확정
- reporting 생략
- next-flow reopening 생략
- turn closure 결정

개별 task 완료는 flow 완료나 turn closure를 결정할 수 없습니다.
예를 들어 `git commit` 명령이 성공해도 commit flow는 commit hash 확인, working tree 상태 확인, reporting, next-flow reopening 조건을 아직 만족해야 할 수 있습니다.

## 검토 질문

- task policy gate가 flow 내부 정책으로 남고 별도 planned flow처럼 쓰이지 않는가?
- task 완료가 reporting이나 next-flow phase를 건너뛰지 않는가?
