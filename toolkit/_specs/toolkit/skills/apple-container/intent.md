## 사용자 스펙 의도

- advance-codex에 스킬하나 추가할거야.
  - docker gui 가 엔터프라이즈 라이선스 제한으로 사용안할 예정이야.
  - 대체제인 podman이 있어 그걸로 사용하다가 올해 apple container라는 오픈소스가 공개되었어.
  - 그 때문에 이제 docker cli 대신 container cli를 사용할 예정이라. 관련 조사가 필요해. 확인좀해줘.
- toolkit이란 플러그인 만들어서 cli 계열 스킬들 몰아넣어버릴까?
- 그럼 설계해보자.
  - SKILL.md: container cli cheatsheet를 넣고, 하단에 references 폴더 routing 정보
  - references/*.md: 설치, 삭제 가이드, 일부 이론 wiki
- k8s plugin 설명도 넣어두자. kubectl 같이 서드파티 cli 사용해야하는건 제외하거나 보충 설명용으로만 제한하자
