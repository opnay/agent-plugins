# flow skill reconfigure 스펙

## 소유 범위

이 문서는 flow entry와 post-reporting continuation boundary의 skill reconfigure를 소유합니다.
skill reconfigure는 루프가 이어지는 동안 잊혔거나 오래된 skill context를 복구하기 위해 현재 flow 또는 다음 main flow에 필요한 skill 본문을 source에서 다시 읽는 절차입니다.

## 계약

- flow에 진입할 때 메시지 인터뷰보다 먼저 skill reconfigure를 수행합니다.
- `reporting -> 다음 intake`로 이어지는 main-flow loop에서는 `reporting` 직후 skill reconfigure를 다시 수행합니다.
- post-reporting skill reconfigure가 끝난 뒤 다음 `intake`로 들어갑니다.
- 현재 사용자 메시지, 선택된 다음 flow 입력, 선택된 main flow 입력, explicit skill 호출, 저장소 규칙, plugin 경계, 승인 경계를 기준으로 필요한 active skill 목록을 식별합니다.
- `flow` 자체는 항상 active skill에 포함합니다.
- flow entry에서 메시지 인터뷰로 들어갈 때는 `deep-interview`를 active skill에 포함합니다.
- `turn-gate` 같은 wrapper가 현재 턴을 소유하면 해당 wrapper skill도 active skill에 포함합니다.
- 기존에 읽은 skill context는 stale context로 보고 폐기합니다.
- 각 active skill 본문을 source에서 다시 읽습니다.
- freshly read bodies만 현재 flow의 active skill set으로 수용합니다.
- active skill 목록은 필요한 경우 `000-plan.md`에 기록합니다.
- 어떤 필수 skill 본문도 읽을 수 없으면 flow 작업으로 진입하지 않고 blocker로 라우팅합니다.

## 산출

- 현재 flow 또는 다음 main flow active skill 목록
- freshly read active skill set
- blocker 여부
- `000-plan.md` 갱신 필요 여부

## 검토 기준

- flow entry가 메시지 인터뷰 전에 필요한 skill context를 복구하는가?
- post-reporting continuation boundary가 다음 main flow 진입 전에 필요한 skill context를 복구하는가?
- 루프 중 잊힌 skill context를 이전 대화 기억이 아니라 source reread로 복구하는가?
- active skill 목록이 현재 flow에 필요한 skill과 wrapper skill을 포함하는가?
- 읽을 수 없는 필수 skill이 있을 때 작업으로 진행하지 않는가?
