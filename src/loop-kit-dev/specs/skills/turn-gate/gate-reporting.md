# turn-gate gate-reporting sub-spec

## 목적

이 문서는 reporting gate의 전환 계약을 소유합니다.

reporting gate는 flow 결과를 다음 진행을 위한 맥락으로 정리합니다.

## 소유

- 준비, 작업, 검증, 남은 불확실성, blocker, approval boundary 상태를 보고한다.
- material routing judgment call을 드러낸다.
- reporting 전에 session record와 Continuity Guard를 갱신한다.

## 비소유

- terminal summary 허용
- explicit stop 없는 clean stop
- next-flow reopening 생략

보고는 turn closure가 아닙니다.

## 검토 질문

- reporting이 terminal closure가 아니라 continuation context로 쓰이는가?
- Continuity Guard를 reporting 전에 갱신했는가?
