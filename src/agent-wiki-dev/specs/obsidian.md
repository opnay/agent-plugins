# Obsidian CLI 조회 계약

## 경계와 대상

- Obsidian CLI는 등록된 위키의 조회·메타데이터 확인 수단입니다. root 등록·선택은 `agent-wiki` 설정 계약이 소유하며 두 CLI를 구분합니다.
- CLI가 사용 가능하고 현재 요청·허용된 맥락에서 vault 이름이나 ID를 알고 있으면 `obsidian vault=<vault> vault info=path`를 실행하고 결과의 canonical 경로를 선택한 agent-wiki root와 대조합니다. 이름이 같다는 이유만으로 root를 연결하지 않습니다.
- 명령마다 첫 번째 매개변수로 `vault=<vault>`를 지정합니다. 파일은 vault 기준의 정확한 `path=<relative-path>`를 사용하고, 검색 결과 경로도 선택 root와 영역 안인지 확인합니다. CLI 기본 활성 vault·파일에 의존하지 않습니다.
- CLI 조회가 앱을 시작할 수 있음을 고려해 현재 실행 제한을 따릅니다. 미설치·미등록·매핑 미확정·경로 불일치·명령 불가이면 설정·앱 설치·플러그인 활성화를 시도하지 않고 같은 승인 root의 파일·YAML 조회로 진행합니다. 매핑을 얻으려고 다른 vault 내용을 탐색하지 않습니다.

## 탐색·어휘·본문

- 정확 경로나 충분한 제공 근거는 재검색 없이 사용합니다. 그 외에는 root·하위 index로 관련 영역을 선택한 뒤 해당 영역의 CLI 조회를 우선 활용합니다. CLI가 인덱스 기반 범위를 대체하지 않습니다.
- `search path=<area> query='[type:benchmark] [tags:llm]' format=json limit=<n>`처럼 유형·태그 조건으로 후보를 좁힙니다. 원하는 구조화된 속성에는 `[type:...]`·`[tags:...]`를 사용합니다. 값이 여러 단어이면 Obsidian 검색 문법으로 인용하고 모든 CLI 인수는 안전하게 전달합니다.
- 검색 결과는 후보입니다. `properties path=<note> format=json`으로 속성을, `read path=<note>`로 본문·근거·시점을 확인합니다. 필요하면 `outline path=<note> format=json`·`search:context`로 확인할 부분을 좁힙니다. limit가 있는 반환을 전체 후보 목록으로 단정하지 않습니다.
- `tags`·`tag` 조회에는 YAML과 본문 태그가 함께 들어가므로 PR 참조·색상 같은 우연한 토큰도 섞일 수 있습니다. 어휘 재사용은 관련 문서의 `properties`와 본문 적합성을 기준으로 판단합니다. 전체 vault 태그 집계는 해당 범위의 어휘 검토가 필요할 때만 사용합니다.
- 메타데이터가 없는 문서도 탐색 대상입니다. 필터 무결과·frontmatter/속성 부재이면 선택 영역의 인덱스·본문 검색을 계속합니다. 범위 확대는 인덱스 누락·불완전·선택 영역 미발견 등 기존 조회 조건을 따릅니다.

## 출력 판정

- 종료 코드와 stdout 의미·자료형을 함께 확인합니다. `format=json` 요청과 exit 0만으로 정상 JSON·작업 성공을 가정하지 않습니다.
- 알려진 `No matches found.`는 해당 검색의 무결과, `No frontmatter found.`와 `Error: Property "..." not found.`는 해당 문서의 메타데이터 부재로 구분합니다. 이 메시지와 형식은 사용하는 설치본의 help·실제 출력으로 확인합니다.
- `search` JSON은 문서 경로 문자열 배열, `properties` JSON은 속성 mapping인지 확인합니다. `tags`의 객체 목록·선택 count 값은 실제 출력 형식을 확인합니다. 예상 밖 평문·JSON 자료형·경로·오류는 실패 또는 미확인으로 보고하며 빈 검색 결과로 바꾸지 않습니다.
- CLI 실패와 지식 부재를 구분하고 승인 root의 파일 조회를 이어갈 수 있으면 계속합니다. CLI 재검증이 필요해도 범위·호출을 제한하며 자동 reload·설정 변경·광범위 재탐색을 하지 않습니다.

## 역할별 적용

- reader는 선택 영역의 type·tags 후보 검색과 본문 확인을 수행합니다. 검색·태그만으로 근거 충족을 확정하지 않습니다.
- researcher는 같은 경로로 기존 조사·비교·진단을 찾아 출처·날짜·적용 조건을 확인하고 충분한 결과를 재사용합니다. reader 선행 실행은 요구하지 않습니다.
- writer는 관련 문서의 속성에서 유형·태그 어휘를 확인하고 승인된 저장 후 CLI `properties`·`read` 결과를 실제 저장 내용과 대조합니다. raw YAML의 키 중복·자료형·본문·근거 보존 검증은 유지합니다. 캐시와 파일이 다르면 실제 저장과 CLI 확인 상태를 나누어 보고하며 CLI 결과에 맞추려고 원문을 덮어쓰지 않습니다.
- 이 연계의 CLI 사용은 읽기입니다. writer의 쓰기는 기존 승인 범위의 파일·YAML 병합 계약을 유지합니다. `property:set`·create·append·move·설정 명령·eval을 자동 호출하는 권한을 부여하지 않습니다.
- manager는 허용 root·정확 경로 또는 영역·필요 type·주제/태그 단서·이미 확인한 vault 이름/경로·확인 시점·읽기/쓰기 범위를 전문 역할에 전달합니다. 새 CLI 의존을 만드는 대신 조회 가능·대체 조회·실패와 실제 저장 결과를 구분해 받습니다. 직접 CLI 실행 정책을 재구현하지 않습니다.

## Runtime·검증

- reader·researcher·writer runtime은 각자 필요한 대상·명령·출력·대체 경로 계약을 포함해 단독 실행합니다. manager는 전달·상태 판정을 소유합니다. 설치된 사용자가 이 개발 spec이나 숨은 sibling context를 읽어야 실행 가능한 구조를 만들지 않습니다.
- 검증: root/vault 매핑, 선택 영역 필터 hit·무결과·메타데이터 부재, JSON/평문·exit 0 오류, 본문 fallback, writer 저장/CLI 확인 분리, 읽기 전용 경계와 기존 type/tags·PR 표기 보존입니다.
- 제품 기능 근거: [공식 CLI](https://obsidian.md/help/cli), [공식 Search](https://obsidian.md/help/plugins/search). 현재 계약의 출력 분기는 설치본의 실제 결과로 대조합니다.
