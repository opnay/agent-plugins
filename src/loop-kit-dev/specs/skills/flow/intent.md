## 사용자 스펙 의도

- flow는 flow로써의 흐름, turn-gate는 turn에 대한 게이트 규칙으로 turn을 진행하는동안 flow를 사용해야하는 강제성을 만드는거지.
- flow는 preparation, work, verification, reporting 내부 단계가 있긴한데, 여러 flow를 쪼개야하는 방식도 있어서 flow를 회귀 스킬로 보는게 더맞을거 같아. 하나의 메시지, 동작에 대해 flow 설계를 진행, 각 플로우는 작업을 진행하는 방식이고, 그 flow가 끝나면 종료. turn-gate는 그 플로우가 끝나면 새로운 플로우가 진행될 수 있도록 질문 도구 사용
- 그러면 flow를 재귀 한다는 개념보단, sub-flow를 만드는 방식으로 하는게 낫지 않을까?
- flow도 스펙을 폴더형태로 개편, intent-scenarios폴더도 추가
