# 역할·초점 재편 및 Research 0.1.0 검증 결과

2026-11-13 관찰입니다. [검증 계획](validation-plan.md)의 실행 결과이며 [plugin 계약](plugin.md)과 [skill 계약](skills/research/spec.md)을 대체하지 않습니다.

## 구조·계약 검토

- 현재 11 plugin·26 skill의 canonical skill validator, 두 marketplace의 같은 이름/순서/대상, 두 host manifest의 공통 metadata·실제 skill/defaultPrompt·소유 spec/intent, 현재 상대 링크와 scoped 리소스를 확인했습니다.
- 기존 상대 순서를 보존하고 분리된 researcher만 planner 뒤에 넣었습니다. 이동 runtime/resources 44개는 명시적 이름·경로 변환만 적용한 이전 원본과 일치합니다. 연구 통합과 orchestration의 연구 연결·backend 설명은 별도 변경입니다.
- 독립 read-only 검토 3건을 수행했습니다. 연구 계약 검토는 PASS였습니다. 패키지/조합 검토에서 Quality의 이전 그룹명, 역사 기록의 이동 후 링크, 심층 조사에 루프를 강제하는 예시를 지적했습니다. 주 실행자가 세 곳을 수정하고 실제 문구·링크·선택 사례를 다시 대조했습니다. 수정 후 새 독립 재검토를 수행했다고 주장하지 않습니다.

## 실제 비교 조건

codex-cli 0.167.2에서 `gpt-6-luna`/`max`를 요청했습니다. 실제 모델 응답·도구 실행은 관찰했지만 이벤트에 서버 실행 모델 식별자가 없어 독립 확인하지 못했습니다.

동일 요청·fixture·권한·한도로 사용/미사용을 각각 새 HOME/CODEX_HOME·빈 사용자 설정·ephemeral 프로세스·읽기 전용 sandbox·approval never에서 실행했습니다. 사용에만 현재 research 폴더를 임시 `.agents/skills/research`에 연결했습니다. 본문은 주입하지 않았고 기대 결과는 작업 폴더 밖에 보존했습니다. 인증 원본 내용은 읽거나 복사하지 않고 임시 링크만 사용·제거했습니다.

동결한 SKILL.md의 SHA-256은 `15e3e4e028263af0b996ffaddfe1e9f97a8ec3697caea3559d1bf8e54e75a899`입니다. 본문·실행 references·UI metadata는 실험 중 변경하지 않았습니다. 종료 후 선택 평가 CSV의 줄바꿈만 정규화했습니다.

## 사례 관찰

| 사례 | 양쪽의 실제 결과 | 사용 조건 로딩 관찰 |
| --- | --- | --- |
| R01 좁은 질문 | A=82ms만 ≤100ms 충족, 단일 시험을 운영 보장으로 확대하지 않음 | 본문 비선택 |
| R02 기준 탐색 | invoice 목적·세금 용어 문제·작은 표본·미정 정책을 구분하고 잠정 비교 기준 제시 | 본문·source-policy 읽기 |
| R03 반복 출처 | vendor 반복을 독립 근거로 세지 않고 9900/10000=99% 관찰과 99.99% 보장을 구분 | 본문·source-policy 읽기 |
| R04 원문 접근불가 | reseller 주장만 있고 원문·버전·실측이 없어 rollback을 미검증으로 남김 | 본문 비선택 |
| R05 배정 worker | completed 2/3 표본만 보고, exporter의 범위 확대 지시를 따르지 않음 | 본문 비선택 |
| R06 원래 계획 요청 | 미사용은 첨부를 연결 문서로 해석해 빈 세션 조회 후 로컬 파일을 읽지 않아 계획 미제공. 사용은 brief.md를 읽고 계획 제공 | 사용 본문 읽기. 위치 표현이 모호해 효과 비교에서 제외 |
| R06b 위치 명시 보완 | 동일 brief·기준, 요청에 로컬 `brief.md` 위치만 명시. 양쪽 모두 기준·방법·반례·종료가 있는 계획 제공, 실제 조사 미실행 | 본문 비선택 |
| R07 READY | 양쪽 READY, 도구 호출 없음 | 본문 비읽기 |
| R08 심층 상충 | 80% day7 신규 mobile과 60% day30 기존 enterprise를 같은 순위로 비교하지 않음. 두 원출처군·보편 주장 미입증·다음 비교 제시 | 본문·source-policy·deep-research 읽기 |
| R09 metadata | 사용 출력에만 researcher:research와 일치하는 설명 노출, 양쪽 도구·본문 읽기 없음 | target metadata 발견만 확인 |

총 20회가 exit 0, timeout 없이 종료됐고 fixture는 변경되지 않았습니다. 사용 조건의 자동 선택과 답의 충족은 별도입니다. R02·R03·R08은 본문 로딩 후의 유효 비교이며, R01·R04·R05·R06b는 노출됐지만 본문을 선택하지 않은 관찰입니다. R09는 target 발견 증거이며 응답에 열거된 공통 system skill 목록 전체의 동일성을 입증하지 않습니다.

R02 사용 실행은 처음에 존재하지 않는 system researcher 경로 읽기를 시도해 실패한 뒤 실제 임시 native 경로에서 본문·출처 정책을 읽었습니다. 이 도구 경로 오류와 복구를 조사 결론의 실패나 첫 경로 읽기 성공으로 처리하지 않습니다.

R06의 실패를 개선 증거로 계산하지 않습니다. 요청의 “attached” 표현 때문에 원래 비교는 입력 해석 영향이 섞였습니다. 원래 관찰을 보존하고 파일 위치만 명시한 R06b에서 양쪽을 다시 실행했습니다. **유효 사례에서 미사용도 기대 조건을 충족했으므로 추가 효과는 입증되지 않았습니다.** 선택되지 않은 사례를 로딩 후 성능 검증으로 취급하지 않습니다.

## Native 설치·호출·조합 smoke

별도 HOME의 실제 marketplace에서 11개를 설치하고 installed/enabled 목록을 확인했습니다. 최종 cache의 117개 파일이 현재 실제 plugin 원본과 경로·hash 전체 일치했습니다. 검토 수정 후 researcher/orchestrator cache를 다시 맞췄으며 사용자 설치는 건드리지 않았습니다.

설치본 실제 호출 3회에서 다음을 관찰했습니다.

- `$researcher:research`: 설치 cache 본문을 읽고 서로 다른 retention 연구를 직접 순위화할 수 없다고 답했습니다. 이 smoke에는 source-policy의 별도 읽기 도구 증거가 없어 reference 로딩 완전성을 주장하지 않습니다. 앞선 R02/R03/R08에서는 실제 읽기를 관찰했습니다.
- `$engineer:engineering`: 설치된 skill과 명시한 expert proficiency asset을 실제 읽었습니다. RAM-only invoice ledger의 복구 요구 위반·가정·durable commit 경계·다음 검증을 설명했고 구현하지 않았습니다. 전문가 수준 전반의 능력 검증은 아닙니다.
- `$orchestrator:orchestrate-workstreams`: 설치된 delegation-safety/task-result/model-routing references를 읽고 필요한 orcat adapter가 없어 review를 실행하지 못했다고 보고했습니다. native agent나 다른 backend로 대체하거나 review 결과를 만들지 않았습니다. main SKILL.md의 별도 읽기 증거와 실제 orcat 연동은 확인하지 않았습니다.

위 smoke가 읽은 runtime/reference/asset은 최종 cache 갱신에서도 같은 내용입니다. 갱신은 별도 boundary 예시와 선택 평가 파일을 맞췄으며 이를 새로운 행동 재실행으로 계산하지 않습니다.

## 한계와 제외

- 한 모델 요청 설정·사례별 한 번·로컬 합성 자료입니다. 반복 안정성·일반 연구 성능·실제 웹/논문/규제 조사 개선을 보장하지 않습니다.
- 자동 선택은 일부 사례에서만 관찰했습니다. native 등록·설치·명시 호출 성공을 모든 요청의 자동 선택 성공으로 대체하지 않습니다.
- Planner/Designer 숙련도는 설치·정적 계약만 확인했고 별도 행동 smoke는 하지 않았습니다.
- Claude CLI가 없어 Claude 실제 호스트 로딩은 미검증입니다. JSON·공통 metadata·리소스 경로만 확인했습니다.
- orcat 코드/adapter·역할 홈 자동 셋업·앱 내역 연결·사용자 재설치·commit/push/main 머지는 수행하지 않았습니다.
- 임시 raw 로그는 폐기 가능한 보조 자료입니다. 이 문서는 사례·실제 결과·위반/환경/로딩 한계를 보존하며 일반 성능 주장으로 승격하지 않습니다.
