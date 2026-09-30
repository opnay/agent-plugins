## 사용자 스펙 의도

- 코드베이스 관리 레이어는 dependency, deprecation, dead code, drift, debt를 맡아야 한다.
- 이 역할은 Git workflow를 소유하지 않아야 한다.

---

# Maintenance Skill Spec

## 목적

`maintenance`는 시간이 지나며 생기는 코드베이스의 불필요한 복잡도, 오래된 의존성, drift, debt를 발견하고 안전한 정리 방향을 제시한다.

## 경계

- 포함: dependency·deprecation 관리, dead code, unused abstraction, drift, debt ledger, 테스트 부채 조사, maintenance audit, removal/migration plan.
- 제외: 새 기능 구현, routine bug fix, system architecture 선택, product priority, Git/PR/release workflow.

## 처리 계약

- 실제 사용처, public contract, generated/vendor 경계, migration cost를 확인한 뒤 후보를 제시한다.
- 삭제·축소·교체 판단은 behavior, compatibility, security, data integrity, operational risk를 보존하는 조건에서만 한다.
- findings는 위치, 영향, 근거, 권장 조치, 검증 또는 rollback 조건으로 기록한다.
- 큰 정리는 작은 독립 단위와 명시적인 migration path로 나눈다.
- 확실하지 않은 사용 여부나 외부 소비자는 제거 결론이 아니라 조사 항목으로 남긴다.
- 테스트 삭제·통합 후보를 조사할 때 `$sw-kit:code`의 본문 가치 기준과 `$sw-kit:code/references/test-audit.md`를 읽는다. 전체 subsystem 테스트 부채 조사에는 `$sw-kit:code/references/test-campaign.md`의 read-only 조사·계획 단계를 적용한다. 일상 maintenance에서 이 표면을 선행 로드하지 않는다.
- 테스트 기준은 `code`가 소유하며 복제하지 않는다. 조사·권고는 read-only이고 source/test 수정은 요청된 구현으로 구분한다.

## 검토 질문

- 실제로 사용되거나 외부 계약인 것은 무엇인가?
- 어떤 복잡도·의존·drift가 어떤 비용을 만드는가?
- 제거 또는 migration이 보존해야 할 호환성은 무엇인가?
- 안전하게 확인할 검증과 rollback은 무엇인가?

## 독립성 원칙

독립적으로 audit과 maintenance 계획을 수행한다. source-level 수정이 필요하면 `code`에 넘길 수 있으나 선행 실행을 전제하지 않는다. 테스트 부채 조사에는 같은 plugin에 bundled된 `code` 본문·reference를 명시적으로 허용하며 동일 계약을 적용하기 위한 resource 의존이다.
