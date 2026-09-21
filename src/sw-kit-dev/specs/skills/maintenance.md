## 사용자 스펙 의도

- 코드베이스 관리 레이어는 dependency, deprecation, dead code, drift, debt를 맡아야 한다.
- 이 역할은 Git workflow를 소유하지 않아야 한다.

---

# Maintenance Skill Spec

## 목적

`maintenance`는 시간이 지나며 생기는 코드베이스의 불필요한 복잡도, 오래된 의존성, drift, debt를 발견하고 안전한 정리 방향을 제시한다.

## 경계

- 포함: dependency·deprecation 관리, dead code, unused abstraction, drift, debt ledger, maintenance audit, removal/migration plan.
- 제외: 새 기능 구현, routine bug fix, system architecture 선택, product priority, Git/PR/release workflow.

## 처리 계약

- 실제 사용처, public contract, generated/vendor 경계, migration cost를 확인한 뒤 후보를 제시한다.
- 삭제·축소·교체 판단은 behavior, compatibility, security, data integrity, operational risk를 보존하는 조건에서만 한다.
- findings는 위치, 영향, 근거, 권장 조치, 검증 또는 rollback 조건으로 기록한다.
- 큰 정리는 작은 독립 단위와 명시적인 migration path로 나눈다.
- 확실하지 않은 사용 여부나 외부 소비자는 제거 결론이 아니라 조사 항목으로 남긴다.

## 검토 질문

- 실제로 사용되거나 외부 계약인 것은 무엇인가?
- 어떤 복잡도·의존·drift가 어떤 비용을 만드는가?
- 제거 또는 migration이 보존해야 할 호환성은 무엇인가?
- 안전하게 확인할 검증과 rollback은 무엇인가?

## 독립성 원칙

독립적으로 audit과 maintenance 계획을 수행한다. source-level 수정이 필요하면 `code`에 넘길 수 있으나 호출을 전제하지 않는다.
