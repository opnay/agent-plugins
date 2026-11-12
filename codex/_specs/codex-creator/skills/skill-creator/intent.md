## 사용자 스펙 의도

- skill boundary와 비책임을 더 엄격하게 잠그고 싶다.
- plugin-owned skill이 hidden sibling context에 기대지 않게 만들고 싶다.
- tool policy와 domain workflow가 섞일 때 skill boundary 안에서 분리할 기준이 필요하다.
- 스킬의 `description`은 스킬 사용 여부를 결정하는 트리거 조건이므로, passive skill은 설명 끝에 토큰 매칭에 걸릴 수 있는 plain token 목록을 쉼표로 나열해야 한다.
  - 이 목록은 literal hashtag가 아니며 `#`를 붙이지 않는다.
  - 이 방식은 사용자가 명시적으로 스킬을 요청하지 않아도 적용되어야 하는 passive skill에 사용한다.
- 스킬 테스팅할 때 유즈케이스를 만들어 스킬을 사용할 때와 안 할 때 비교하는 과정이 필요합니다.
  - 일반적인 유즈케이스와 함께, 스킬이 없을 때 실패할 만한 내용의 유즈케이스를 포함해야 합니다.
- 스킬 사용 조건은 실행 작업 폴더의 `.agents/skills`에 현재 테스트 대상 스킬 폴더를 임시 심링크하고, 에이전트가 직접 발견·선택·본문 로딩하도록 검증해야 합니다.
