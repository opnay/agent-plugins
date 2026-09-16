## 사용자 스펙 의도

- 디자인 기본 베이스 구조와 프로젝트 특성에 맞춘 전용 디자인 패턴 규칙을 쌓고 싶다.
- 색 대비, 보색, 공간감 등 디자인 이론부터 웹·프로덕트·콘텐츠 디자이너, 아트·크리에이티브 디렉터의 기준을 다루는 `design-kit`을 원한다.
- 콘텐츠 디자이너는 시각 콘텐츠를 뜻한다.
- 패턴에는 반복을 포함한다. 동일 문구는 맥락에 따라 생략하거나 추가해 강조할 수 있다.
- 같은 맥락에서 동일한 용어, 색상, 패턴의 반복도 필요하다.

# Design Kit 플러그인 스펙

## 플러그인 목적

공통 디자인 지식을 바탕으로 프로젝트에 맞는 경험과 시각 표현을 설계하고, 디자인 방향과 재사용 가능한 규칙을 유지한다.

## 플러그인 경계와 비목표

- 포함: 디자인 원리 적용, 웹·프로덕트·시각 콘텐츠 설계, 아트·크리에이티브 디렉션, 프로젝트 디자인 규칙 관리.
- 제외: 제품 범위·사업 정책의 최종 결정, 일반 코드 구현 절차, 출시 승인, 외부 게시·배포, 도구별 조작 매뉴얼.
- 역할은 소유 결정과 산출물로 구분한다. 3층은 필수 실행 순서나 승인 계층이 아니다.
- 플러그인은 공통 지식과 기록 방법을 소유한다. 실제 프로젝트 결정값과 패턴은 대상 프로젝트가 소유한다.

## 처리하려는 작업 형태

- 디자인 원리를 적용해 화면이나 시각 산출물을 만들고 검토한다.
- 제품 사용 흐름, 웹 표현, 시각 콘텐츠를 전문 기준으로 설계한다.
- 메시지·콘셉트·시각 언어를 정하고 매체별 전개를 판단한다.
- 프로젝트의 디자인 베이스와 반복 가능한 패턴을 기록·적용·갱신한다.

## 대표 표면

- 사용 안내: `design-kit/README.md`, `design-kit/.codex-plugin/plugin.json`과 개발 README.
- 처리 계약: `specs/skills/<skill-name>.md`.
- 실행 표면: `design-kit/skills/<skill-name>/SKILL.md`, 설치되는 `references/`, `templates/`.

## 내장 skill 체계

| 영역 | Skill | 시작 기준과 소유 결정 | Spec |
| --- | --- | --- | --- |
| 기반 | `design-base` | 색채·시지각·구성·타이포·상호작용·반복의 원리 적용 | [design-base](skills/design-base.md) |
| 전문 설계 | `web-designer` | 페이지 전달, 탐색, 웹 레이아웃·반응형 | [web-designer](skills/web-designer.md) |
| 전문 설계 | `product-designer` | 과업, 흐름, 조작, 상태·회복 | [product-designer](skills/product-designer.md) |
| 전문 설계 | `visual-content-designer` | 배너·카드뉴스·에디토리얼·소셜 콘텐츠 | [visual-content-designer](skills/visual-content-designer.md) |
| 디렉션 | `art-director` | 시각 콘셉트, 이미지·타이포·색·구도 언어 | [art-director](skills/art-director.md) |
| 디렉션 | `creative-director` | 메시지, 핵심 아이디어, 매체 간 전개 | [creative-director](skills/creative-director.md) |
| 프로젝트 공통 | `project-design-rules` | 방향·베이스·전용 패턴의 기록, 채택, 재사용·갱신 | [project-design-rules](skills/project-design-rules.md) |

## 사용과 관계

- 과업 완료·상태가 중심이면 product, 페이지 전달·반응형이 중심이면 web을 선택한다. 웹으로 구현된 제품도 product 판단이 중심일 수 있다.
- UI 문구는 web/product가 소유한다. visual-content는 시각 콘텐츠의 메시지와 구성을 소유한다.
- creative는 무엇을 누구에게 어떤 아이디어로 전달할지, art는 그것을 어떤 시각 언어로 구현할지 결정한다.
- 작은 작업은 필요한 역할만 적용한다. 대안 수, 전체 감사, 디렉터 검토를 고정하지 않는다.
- 각 역할은 필요한 공통 reference를 명시적으로 읽는다. 다른 skill의 선행 실행이나 이전 대화는 전제하지 않는다.
- 디자인 역할은 결정과 근거를 만들고, project-design-rules는 지속 규칙의 적용 범위와 상태를 관리한다. 사용자 요청이 기록을 포함하면 함께 적용한다.

## SDD 운영 원칙

- flat skill spec이 처리 계약과 의도를 소유한다. 공통 이론은 design-base, 프로젝트 맥락 읽기·기록은 project-design-rules가 소유한다.
- runtime은 영문이며 설치되는 리소스만 참조한다. 다른 skill의 리소스는 `$<plugin>:<skill>/<resource-path>`, 자기 skill의 리소스는 내부 상대 경로를 사용한다. dev-only spec은 실행 입력이 아니다.
- 사용 기준 변경은 README·manifest·해당 skill spec을 함께 갱신한다.
- marketplace 항목은 `./design-kit`을 가리킨다.

## 검증 기준

- 루트 `design-kit/`에 7개 skill과 참조·template가 존재하고 JSON·skill·plugin 검증을 통과한다.
- skill 식별자 기반 참조와 내부 상대 링크의 대상이 설치되는 plugin 안에 존재하며 release에 specs·changes가 없다.
- 기반·전문 역할·디렉션·규칙 관리의 책임이 겹칠 때 시작 기준과 산출물이 구분된다.
- 반복은 생략·재제공·강조와 의미·역할의 일관성, 시각적 연결·리듬을 모두 다룬다.
- 독립 read-only 검증은 spec/runtime 정합성을 확인하고 실제 요청 시나리오로 경계를 확인한다.
