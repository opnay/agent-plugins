# Message Contract

## Subject

- 첫 줄은 `type: detailed subject` 형식을 사용합니다.
- subject는 120자 미만으로 유지합니다.
- subject는 실제 staged scope를 구체적으로 설명합니다.
- vague subject나 여러 unrelated concern을 묶은 subject를 피합니다.

## Type Selection

- `feat`: new user-facing feature
- `fix`: bug fix
- `refactor`: behavior-preserving code restructuring
- `docs`: documentation-only change
- `test`: test addition or update
- `perf`: performance improvement
- `style`: formatting/style-only change
- `build`: build system or dependency change
- `ci`: CI configuration or script change
- `chore`: maintenance outside the above types

## Body

- body는 항상 작성합니다.
- bullet list로 실제 변경과 검증 근거를 요약합니다.
- skipped verification이나 residual risk가 있으면 body 또는 final report에 숨기지 않습니다.
- body는 literal `\n` escape, 불필요한 blank line, unrelated scope 설명을 피합니다.
