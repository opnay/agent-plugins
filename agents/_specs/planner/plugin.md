## 사용자 스펙 의도

- 제품 기획은 사용자 문제·가치·범위·우선순위·요구사항을 소유하며 기술 설계와 구분합니다.

# Planner 플러그인 스펙

## 목적과 경계

`planner`는 제품·서비스·기능 요청을 사용자 문제·제품 방식·기능 영역·MVP·요구사항·수용 기준·인계로 구체화하는 직업 plugin입니다. 코드 구현·기술 구조·디자인 표현·조사 방법론·품질 검증 절차·출시 승인은 소유하지 않습니다.

## 내장 skill과 사용 기준

단일 `$planner:plan`이 넓은 제품 요청의 분해와 제품 판단을 제공합니다. 상세 계약은 [plan spec](skills/plan/spec.md), 사용자 입력은 [intent](skills/plan/intent.md)가 소유합니다. `assets/proficiency.md`는 기획 직업의 기대 판단 수준이며 역할 셋업에서 선택해 명시적으로 제공합니다. 수준 미지정도 skill 실행을 막지 않습니다.

## 독립성과 조합

자체 skill 계약으로 실행하며 외부 skill 선행 호출을 요구하지 않습니다. 근거 조사는 researcher, 산출물 검증은 quality, 기술 행동 계약은 engineer/spec과 연결할 수 있습니다. 이들은 분야 지식·권한을 자동 공급하지 않습니다. 필요한 직업/분야/초점/숙련도와 원본 조건을 인계합니다. 현재 plan 실행 계약은 기존 제품 기획 계약을 보존합니다.

## 사용 표면과 확장

실제 plugin은 `agents/planner`, spec은 `agents/_specs/planner`입니다. 두 manifest/README/prompt는 실제 plan 하나를 노출합니다. 신규 분야는 고유 기준·산출물·선택 조건이 있을 때만 추가하며 제품 작업 단계를 기계적으로 별도 skill로 분해하지 않습니다.
