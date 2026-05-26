# flow discovery preparation 계약

## 소유 범위

요구사항 발견과 scope lock이 flow contract 형성의 병목일 때 필요한 질문 주제.

## 계약

- discovery는 flow contract를 만들기 위한 flow-local preparation 전략이다.
- 사용자 의도, 포함 범위, 비목표, 성공 기준, 검증 기대가 부족하면 discovery가 필요하다.
- 같은 요청에서 여러 유효한 산출물이나 sub-flow candidate가 나올 수 있으면 discovery가 필요하다.
- 질문 답변에 따라 결과물, flow decomposition, verification path, approval-sensitive checkpoint가 달라지면 work로 들어가지 않는다.
- discovery 결과는 flow decision 또는 parent flow의 sub-flow candidate output으로 이어져야 한다.
- bounded choices로 잠글 수 있는 질문 주제를 산출하되, `request_user_input` 사용 방식은 question-routing이 소유한다.
- 위험 작업 승인은 discovery 질문만으로 대체하지 않는다.

## 검토 기준

- 질문 답변에 따라 flow contract 또는 sub-flow 후보가 달라지는가?
- discovery 결과가 scope, non-goals, completion criteria, verification expectation으로 반영되는가?
- 질문 주제 산출과 질문 도구 실행 권한을 분리했는가?
