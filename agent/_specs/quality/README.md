# Quality 설계 문서

`quality`는 모든 분야의 산출물에서 목적·특성·위험에 맞는 검증 기준과 절차를 도출하고 실행하는 공통 행동 패턴입니다.

- [Plugin 계약](plugin.md): 이름·경계·단일 skill 구성·역할군 및 실행 환경과의 관계.
- [Verify 계약](skills/verify/spec.md): 입력·검증 축·기준·방법·실행·판정·종료.
- [사용자 의도](skills/verify/intent.md): 분야를 닫지 않는 산출물 검증과 역할군·subagent 조합의 요구.
- [검증 계획](validation-plan.md): 설계·패키징·native 로딩·사용/미사용 비교의 기준과 대표 사례.
- [검증 결과](validation-results.md): 실제 로딩·행동 비교의 관찰과 효과·환경 한계.

실행 파일은 `agent/quality/`에서 관리하며 사용 방법은 [plugin README](../../quality/README.md)를 확인합니다. 호출 식별자는 `$quality:verify`입니다. 이 폴더는 작성·검증 계약이며 설치 대상이 아닙니다. 기존 `pro-quality-manager`의 개편은 별도 범위입니다.
