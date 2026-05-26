# flow fix-verify-loop 계약

## 소유 범위

active flow 안에서 하나의 좁은 문제를 작은 fix-verify-reassess cycle로 처리하는 전략.

## 계약

- fix-verify-loop는 문제 하나가 명확하고 작은 수정으로 가설을 검증할 수 있을 때 사용한다.
- 한 loop는 하나의 primary issue에 집중한다.
- 현재 가설을 확인할 수 있는 가장 작은 유용한 수정 또는 확인 동작을 선호한다.
- 수정 직후 verification expectation에 맞게 검증한다.
- 다음 loop가 필요한지는 매번 재평가한다.
- success criteria, non-goal, verification expectation, approval boundary가 바뀔 정도로 커지면 현재 loop를 멈추고 flow preparation 또는 handoff로 되돌린다.

## 검토 기준

- 이 loop가 하나의 primary issue에 머무르는가?
- 수정 또는 확인 동작이 가설 검증에 충분히 작고 직접적인가?
- 다음 반복 필요성이 명확한가?
