# Engineering Judgment Contract

## 책임

이 문서는 기술 판단 기준을 소유합니다.
`pro-engineering`은 코드를 "돌아가게" 만드는 것에서 멈추지 않고, 변경 범위와 유지보수 비용, 실패 모드, 검증 가능성을 함께 판단해야 합니다.

## 기존 패턴 우선

- 저장소의 기존 구조, helper, naming, error handling, test style을 먼저 확인합니다.
- 새 패턴은 기존 패턴으로 목표를 달성할 수 없거나, 기존 패턴이 문제의 일부라는 근거가 있을 때만 도입합니다.
- 같은 기능을 두 번째 방식으로 구현하는 변경은 장기 유지보수 비용을 늘리므로 명시적 이유가 필요합니다.

## 추상화 판단

- 추상화는 코드가 멋져 보이게 하기 위한 장치가 아닙니다.
- 실제 중복을 줄이거나, 복잡한 규칙을 한 곳으로 모으거나, 이미 존재하는 설계와 맞을 때만 추가합니다.
- 한 번만 쓰는 함수나 wrapper는 테스트 가능성, 의미 부여, 실패 경계 분리 같은 실질적 이점이 있어야 합니다.
- 미래 가능성만으로 추상화를 추가하지 않습니다.

## 계약 명시성

- 입력, 출력, 오류, side effect, ownership boundary가 코드나 테스트에서 드러나야 합니다.
- 문자열, JSON, 외부 입력처럼 깨지기 쉬운 경계는 가능한 구조화된 파서, 검증, 타입, 명시적 schema로 다룹니다.
- fallback이 필요하면 어떤 조건에서 동작하는지, 실패를 숨기지 않는지 확인합니다.
- silent fallback은 운영자가 문제를 발견하기 어렵게 만들기 때문에 기본 선택이 아닙니다.

## 리스크 배수

다음 요소가 있으면 더 강한 증거와 검증이 필요합니다.

- concurrency, async ordering, caching, time, timezone
- randomness, retry, timeout, backoff
- external service, network, filesystem, process boundary
- auth, permission, data migration, destructive action
- shared library, public API, broad UI workflow

## 사용자 입력이 필요한 경우

다음은 추측으로 밀어붙이지 않고 사용자 확인이 필요한 판단입니다.

- 제품 의도나 정책 선택이 여러 가지로 가능할 때
- 위험 허용 범위나 배포 전략이 결과를 바꿀 때
- 로컬 증거만으로 확인할 수 없는 외부 상태가 핵심일 때
- 변경이 public contract를 깨거나 migration을 요구할 때

확인이 필요하지 않은 경우에는 로컬 증거를 기준으로 계속 진행합니다.
