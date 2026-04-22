# teammate-kit-guide 스킬 스펙

## 목적

`teammate-kit-guide`는 협업 작업이 orchestration을 요구하는지, 아니면 하나의 bounded teammate role로 충분한지 분류하는 엔트리포인트 스킬입니다.

## 경계

- 포함:
  - orchestration vs direct role 판단
  - research, implementation, review 역할 선택
- 제외:
  - role 자체의 상세 실행
  - generic workflow planning

## 처리하려는 작업 형태

- 협업 형태를 먼저 정해야 하는 경우
- teammate-style execution의 최소 필요한 coordination 수준을 판단해야 하는 경우

## 엔트리포인트 / 대표 표면

- 대표 표면: `teammate-kit/skills/teammate-kit-guide/SKILL.md`

## 핵심 처리 계약

- 현재 작업이 durable orchestration을 필요로 하는지 먼저 판단한다.
- direct role이면 research, implementation, review 중 하나를 고른다.
- routing 결과에는 선택 이유와 예상 handoff 구조를 포함한다.

## 독립성 원칙

- 이 스킬은 collaboration mode routing만 소유한다.
- guide만 읽어도 언제 orchestration을 쓰고 언제 direct role이면 되는지 이해 가능해야 한다.

## 확장 원칙

- role 종류나 collaboration mode가 바뀌면 plugin spec과 함께 갱신한다.

