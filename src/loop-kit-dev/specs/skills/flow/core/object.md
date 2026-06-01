# flow object 계약

## 소유 범위

`flow`만 해석하고 사용할 수 있는 목적 사슬 파일 계약.
이 계약은 flow의 보조 표면이며, 독립 entrypoint나 별도 flow type이 아닙니다.

## 계약

- 목적 사슬 파일은 저장소 전체 목적에서 현재 변경 목적까지 내려오는 객체 사슬만 담습니다.
- 파일 이름은 세션 표면에서 `000-object.md`처럼 둘 수 있지만, 의미와 사용 권한은 `flow`가 소유합니다.
- `turn-gate`는 이 파일의 의미를 재정의하거나 next-flow routing, terminal closure, verification status, approval authority 근거로 사용하지 않습니다.
- 목적 사슬 항목은 객체명과 종류를 함께 드러냅니다. 예: `` `loop-kit-dev` plugin: ... ``.
- 목적 사슬은 상태, 검증 결과, continuity rule, phase log, next action을 담지 않습니다.
- 목적 사슬이 현재 flow의 scope, acceptance, verification, approval boundary, handoff를 바꾸면 `flow` intake 또는 framing으로 돌아갑니다.

## 검토 기준

- 파일 내용이 목적 사슬만 담는가?
- 항목이 저장소 전체 목적에서 현재 변경 목적까지 작은 단위로 내려오는가?
- 상태, 검증, 라우팅, continuity rule이 섞이지 않았는가?
- `turn-gate`가 파일 의미를 해석하거나 권한 근거로 사용하지 않는가?
