# 구조화 입력·장문 모드

## 변경사항 요약

- 단건 JSON 파일과 batch JSON state·instructions·criteria를 지원합니다.
- 단건·batch에 요청별 `--long-context`를 제공합니다.
- 구조화 Score legend를 보존하며 기존 간단한 명령·JSONL과 출력 계약을 유지합니다.

## 변경 상세

에이전트가 최신 Jev API의 텍스트 입력·기준을 CLI에서 사용할 수 있도록 합니다. 소유 계약은 `specs/cli.md`, 사용 계약은 `specs/skills/jev.md`입니다. 이미지·Box·usage, 설치·설정 변경, commit·push는 포함하지 않습니다.

- 단건: `--input` JSON object의 state·instructions·criteria. 기존 입력 flags와의 혼용은 오류입니다.
- Batch: 기존 JSONL과 호환하며 instructions·criteria를 지원합니다. 공통 state는 기본 text, 명시적 `--state-format json`이면 구조를 유지합니다.
- 장문: 기본 off, 활성 요청에만 공식 헤더를 전달합니다. 저장 설정은 늘리지 않습니다.
- 검증: 가짜 HTTP transport로 원형 전송·헤더·Score legend·기존 명령 호환성·잘못된 입력의 호출 전 실패를 확인합니다. 실제 모델 정확성은 로컬 검증 범위가 아닙니다.
