## 사용자 스펙 의도

- Figma처럼 외부 app plugin이 제공하는 기능에 사용자 정의 guidance를 덧붙이는 skill을 만들고 싶다.
- 외부 도구별 extension skill을 `app-extensions` plugin이 소유하게 하고 싶다.
- 첫 skill 이름은 `figma-ext`로 한다.

---

# App Extensions Dev 플러그인 스펙

## 플러그인 목적

`app-extensions-dev`는 외부 app plugin이 제공하는 capability에 사용자·팀의 추가 workflow guidance를 적용하는 instruction-only companion plugin입니다.
upstream plugin의 manifest, tool schema, 인증, prerequisite를 복제하거나 변경하지 않고, upstream에 없는 지속적인 판단 기준과 작업 규율만 delta로 소유합니다.

## 플러그인 경계와 비목표

- 포함:
  - 외부 app capability를 사용할 때 적용할 사용자·팀의 추가 판단과 workflow 계약
  - 대상 app, 적용 조건, upstream 소유권, capability 부재 시 실패 조건이 분명한 bounded extension skill
  - upstream 변경과 독립적으로 검토할 수 있는 delta-only guidance
- 제외:
  - upstream plugin manifest의 상속, 병합, monkey patch
  - 외부 app의 tool schema, 인증, 연결 설정, 공식 prerequisite 복제
  - 외부 도구라는 공통점만 있는 일반 사용법 모음
  - generic product design judgment, frontend implementation, CLI operation, Codex surface 제작
  - 실제로 제공하지 않는 MCP server, app connection, hook, script 선언

## 처리하려는 작업 형태

- 외부 app plugin을 사용하면서 upstream guidance에 없는 로컬 규칙을 함께 적용하는 작업
- app capability의 대상, 좌표, 상태, side effect, handoff 기준을 일관되게 해석하는 작업
- upstream capability가 없거나 필요한 근거가 부족할 때 추정 실행을 멈추고 한계를 보고하는 작업

## 대표 표면

- 대표 스펙: `src/app-extensions-dev/specs/plugin.md`
- skill 상세 스펙 위치: `src/app-extensions-dev/specs/skills/*.md`
- 선택 기준: 외부 app 자체의 기능이 아니라 그 기능에 덧붙일 지속적인 extension contract가 필요한가

## 내장 skill 체계

- `figma-ext`: Figma 작업에서 target design element, hierarchy, absolute·relative coordinates, viewport·responsive context, flow-first layout translation을 추가로 통제합니다.
  - spec: `src/app-extensions-dev/specs/skills/figma-ext.md`

## Plugin Usage 계약

- manifest와 README는 실제 제공되는 extension skill만 노출합니다.
- 각 skill은 대상 app과 extension delta를 description에서 식별할 수 있어야 합니다.
- upstream plugin은 capability와 tool-use prerequisite를 계속 소유하며, extension skill은 이를 복제하거나 우회하지 않습니다.
- 필요한 upstream capability가 없으면 extension skill은 설치나 연결을 추정하지 않고 불가 상태를 보고합니다.
- 새 skill은 외부 app이라는 이유만으로 추가하지 않고, 독립적인 extension contract와 반복되는 change pressure가 있을 때만 추가합니다.

## SDD 운영 원칙

- plugin boundary와 usage routing은 plugin spec, README, manifest가 소유합니다.
- app별 판단과 실행 보정은 각 skill spec과 runtime skill이 소유합니다.
- skill spec을 바꾸면 runtime skill folder를 현재 spec 기준으로 처음부터 재작성합니다.
- source-only spec과 change 기록은 release surface에 포함하지 않습니다.

## 현재 구조 메모

- 초기 version은 `0.1.0`입니다.
- 첫 runtime surface는 `figma-ext` 하나이며 미래 app extension을 미리 약속하지 않습니다.
- 초기 plugin은 skills-only이며 `.app.json`, `.mcp.json`, hooks, scripts를 제공하지 않습니다.
- release surface는 build command로 생성하고 repository marketplace의 `app-extensions` 항목은 root `./app-extensions`를 가리킵니다.
- local install은 configured `opnay-plugins` marketplace에서 release plugin을 설치하며 새 thread에서 runtime pickup을 확인합니다.
