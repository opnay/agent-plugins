# 사용자 스펙 의도

- `judgment-kit`에 가장 작고 안전한 코드 변경을 판단하는 전문가 skill을 추가한다.
- 실제 skill 식별자는 `pro-code-keeper`로 둔다.
- `lean senior developer`는 callable skill 이름이 아니라, 가장 작고 안전한 변경을 선호하는 mode 개념으로만 사용한다.
- 이 skill은 실제 코드 flow를 이해한 뒤 불필요한 코드, 추상화, 의존성, future-proofing을 줄이는 판단을 제공한다.
- overengineering review, simplification review, dependency reduction, lean debt tracking에 쓸 수 있어야 한다.
- 새 runtime은 `pro-code-keeper` 이름만 공개 호출 이름으로 유지하고, 구현·리뷰·저장소 audit·root-cause fix·dependency check·refactor shrink·debt ledger를 reference 기반 분기로 처리한다.
- argument 전용 frontmatter는 사용하지 않는다. 분기 안내는 설치되는 `SKILL.md` 본문과 bundled references가 소유한다.
- legacy 코드 주석 marker는 `lean:`과 `ponytail:`만 debt ledger 입력으로 읽고, 외부 원문이나 별도 skill 이름은 runtime 표면에 복제하지 않는다.
