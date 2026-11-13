# 에이전트 구성과 플러그인 소유권

## 구성 축

| 축 | 소유 책임 | 저장·적용 |
| --- | --- | --- |
| 직업 | 업무·결정·산출물·인계 | `agents/<직업>` plugin |
| 분야 | 직업 안의 전문 기준·제약·기법 | 직업 plugin의 분야 skill |
| 초점 | 조사·검증·실행 조정의 절차 | `specialists/<초점>` plugin |
| 숙련도 | 판단 범위·대안·예외·근거의 기대 수준 | 직업 plugin의 `assets/proficiency.md`를 역할 지침에 적용 |

`skills/`는 분야만 담는 폴더가 아닙니다. Engineer의 spec·engineering·code·maintenance처럼 직업 공통 작업도 소유합니다. 분야 skill은 고유 기준이 있을 때 추가하며 공통 작업 계약을 복제하지 않습니다. Designer의 공통 원리·프로젝트 규칙 관리도 유지합니다.

역할 지침은 직업·필요한 분야·초점·숙련도와 원본 요구·산출물·허용 행동을 명시합니다. 이 축은 필수 호출 순서가 아닙니다. 필요한 skill만 적용하며 조합마다 별도 skill을 만들지 않습니다. 숙련도는 실제 능력·권한·독립 증거를 보장하는 표시가 아닙니다.

## 현재 제공 표면

| Plugin | Skills |
| --- | --- |
| planner | plan |
| engineer | spec, engineering, code, maintenance |
| designer | design-base, product-designer, web-designer, visual-content-designer, art-director, creative-director, project-design-rules |
| researcher | research |
| quality | verify |
| orchestrator | orchestrate-workstreams |
| agent-wiki | manager, reader, researcher, writer |
| toolkit | git, apple-container |
| app-extensions | figma-ext |
| codex-creator | plugin-creator, skill-creator |
| codex-optimize | optimize-token, asd-ste100 |

Planner는 사용자 문제·제품 범위, engineer/spec은 구현 가능한 행동 계약, designer는 경험·표현을 소유합니다. 연구는 질문에 근거 있는 답을 만들고 검증은 산출물·주장의 기준 충족을 확인합니다. Orchestrator는 분업·실행·통합을, wiki manager는 구체 조건·진행·기록 연결을 소유합니다. Wiki researcher는 등록된 지식의 공백·모순·최신성 확인이며 범용 research와 동일한 skill이 아닙니다.

## 숙련도 적용

Planner·Engineer·Designer의 설치된 plugin에서 `assets/proficiency.md`를 읽고 선택한 수준과 과제에 해당하는 기대 판단을 역할 지침이나 작업 패킷에 명시합니다. 기본은 기존 skill의 전문 작업 계약입니다. 수준을 지정하지 않았다는 이유로 실행을 막지 않습니다. 낮은 수준도 필수 요구·검증·권한을 생략할 수 없습니다.

제품 디자이너의 연구 작업은 product-designer의 분야 기준, researcher/research의 조사 계약, designer의 전문가 판단 기대를 결합합니다. 연구 결과를 디자인 완료로 취급하지 않습니다. reference를 설치하거나 수준을 선언하는 것과 실제 읽고 적용한 것은 구분합니다.

## 실행과 서브에이전트

역할별 Codex 홈에는 필요한 plugin과 역할 지침을 준비합니다. 이 저장소의 운영에서는 orcat으로 실행하는 역할 에이전트를 외부 서브에이전트로 취급합니다. 호출·대기·결과·취소·기록 확인은 실행 환경의 adapter가 소유합니다. 설치된 skill은 그 기능을 새로 제공하지 않습니다.

배정 패킷에는 질문·원본 요구·자료/버전·전문 기준·초점·필요한 숙련도·권한·산출물·검증/종료 조건을 줍니다. 실행 수단이 없으면 없는 범위와 가능한 결과를 보고하며 다른 backend로 조용히 바꾸거나 실행을 꾸미지 않습니다. 조사·검증은 직접 수행할 수 있고 별도 agent를 항상 요구하지 않습니다.

현재 orcat는 구성 참조를 자동 적용하거나 plugin을 자동 설치하지 않습니다. 이 저장소는 orcat 코드·앱 세션 연결을 구현하지 않습니다. 외부 worker 기록과 앱 native subagent 화면은 별개이며 앱 연동은 완료 조건이 아닙니다.

## 작성 계약과 migration

각 분류의 `_specs/<plugin>/plugin.md`, `skills/<skill>/{spec.md,intent.md}`가 작성 계약을 소유합니다. runtime은 이 문서를 실행 전제로 삼지 않습니다. 과거 change spec·실험 기록의 이전 이름과 경로는 당시 관찰로 보존합니다.

| 이전 호출 | 현재 호출 |
| --- | --- |
| `$judgment-kit:pro-planner` | `$planner:plan` |
| `$judgment-kit:pro-researcher`, `$advance-subagent:deep-research` | `$researcher:research` |
| `$advance-subagent:orchestrate-workstreams` | `$orchestrator:orchestrate-workstreams` |
| `$sw-kit:<skill>` | `$engineer:<skill>` |
| `$design-kit:<skill>` | `$designer:<skill>` |

이전 alias는 제공하지 않습니다. Marketplace의 기존 상대 순서는 유지하고 분리된 researcher는 planner 바로 뒤에 둡니다. 사용자 설치 갱신·기존 cache 제거는 별도 실행 작업입니다.
