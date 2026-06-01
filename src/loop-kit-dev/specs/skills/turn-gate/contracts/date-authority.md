# turn-gate date authority 계약

## 소유 범위

relative date를 active turn routing에 적용할 때의 기준.

## 계약

기본 기준은 현재 시스템 날짜와 timezone입니다.
relative date가 target, verification path, reporting scope, record reconstruction을 바꾸면 절대 날짜를 기록합니다.

session record 날짜, 이전 flow 날짜, 사용자 relative date가 충돌해 결과가 바뀌면 clarification으로 라우팅합니다.
기록 날짜는 system date를 조용히 덮어쓰지 않습니다.

## 검토 기준

- 날짜 기준이 결과나 검증 경로를 바꾸는가?
- 바뀐다면 절대 날짜가 드러나는가?
- record-date와 system-date 충돌을 질문으로 라우팅했는가?
