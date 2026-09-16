# App Extensions

`app-extensions`는 외부 app plugin의 capability에 사용자·팀의 추가 workflow guidance를 적용하는 companion plugin입니다.
upstream plugin을 상속하거나 수정하지 않으며, tool schema, 인증, 연결, prerequisite를 복제하지 않고 지속적인 extension delta만 소유합니다.

## Skill 선택

### Figma Ext

`$app-extensions:figma-ext`는 Figma 작업에서 target design element, parent hierarchy, absolute·relative coordinates, frame·viewport·responsive context, flow-first layout translation을 함께 판단할 때 사용합니다.

대표 요청:

- Figma page나 section을 code layout으로 옮깁니다.
- import 또는 write 전에 정확한 target node와 parent hierarchy를 확인합니다.
- canvas 좌표와 parent-relative layout 관계를 함께 해석합니다.
- mobile, tablet, desktop variant와 responsive evidence를 비교합니다.
- grid·flex와 normal flow를 우선하고 sticky, fixed, absolute의 필요성을 검토합니다.

## 경계

- 포함: upstream app capability에 덧붙이는 bounded workflow와 handoff guidance
- 제외: upstream manifest 상속·병합, app connection·MCP·인증 제공, tool schema 복제
- capability 부재: 연결을 임의 설치·변경하지 않고 unavailable 상태를 보고합니다.
