# flow 메시지 인터뷰 스펙

## 기준 그래프

```text
entry skill reconfigure
-> deep-interview skill
-> locked execution brief
```

## 계약

- 메시지 인터뷰는 flow entry skill reconfigure가 끝난 뒤 반드시 `deep-interview` skill을 적용합니다.
- `flow`는 `deep-interview`의 질문, 압력 테스트, locked brief 산출 계약을 재구현하거나 반복 설명하지 않습니다.
- 메시지 인터뷰의 산출은 `deep-interview`가 만든 locked execution brief입니다.
- `deep-interview`가 충분히 잠긴 brief를 만들면 `flow`는 사용자 질문 없이 플로우 설계로 진행합니다.

## 산출

- locked execution brief
- `000-plan.md` 갱신 필요 여부
