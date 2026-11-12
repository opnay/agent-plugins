## 사용자 스펙 의도

- Figma 작업에서는 page, section, section part, component, asset 등 대상 design element를 먼저 식별하고 싶다.
- import 전에 target design element를 확정하고 싶다.
- 좌표를 absolute와 parent-relative 관점에서 함께 해석하고 싶다.
- mobile, tablet, desktop, responsive 등 page와 frame size를 고려하고 싶다.
- 구현에서는 document flow를 우선하고 position은 static, sticky, fixed, absolute 순서로 검토하며 layout은 grid와 flex를 먼저 사용하고 싶다.
- 이 guidance를 `app-extensions`의 `figma-ext` skill로 사용하고 싶다.
