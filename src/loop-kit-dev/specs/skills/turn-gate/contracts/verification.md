# turn-gate verification 계약

## 소유 범위

verification method/result routing.
`flow`가 verification expectation을 소유하고, `turn-gate`는 실제 method와 result status를 기록·라우팅합니다.

## 계약

Methods:

- `clean-context`
- `normal`
- `not-required`

Results:

- `pass`
- `fail`
- `blocked`
- `insufficient`

`not-required`는 method이지 pass가 아닙니다.
`not-started`와 `requested`는 progress state이지 success evidence가 아닙니다.

non-enum verifier 결과는 reporting 전에 위 result 중 하나로 정합화합니다.
non-pass는 success reporting, self-drive continuation, release readiness, commit-readiness, next-flow continuation보다 먼저 라우팅합니다.

## 검토 기준

- method와 result가 분리되는가?
- non-pass가 다음 진행보다 먼저 처리되는가?
- verification/build/readback이 commit, publish, release, version bump, destructive/external action 승인으로 쓰이지 않는가?
