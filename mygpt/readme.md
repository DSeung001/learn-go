# AI 스토리텔러

## 1. 전체 프로젝트의 특징 및 시연
- **Go + Ollama 기반 실시간 AI 스토리텔러**
- WebSocket을 통한 실시간 LLM 응답 스트리밍
- ChatGPT 스타일 UI, 대화 로그, 저장/불러오기, 마크다운 렌더링 등 UX 강화
- Ollama llama3.2 모델(영어 최적화) + 구글 번역 API로 한글 번역 제공
- 키워드 입력 → 영어 스토리 생성 → 한글 번역 → 요약(영/한) → 저장/팝업/리스트

**Ollama 모델 설치**
```bash
# llama2 모델 다운로드
ollama pull llama3.2
```

**저장 구조:**
  ```json
  {
    "title": "제목",
    "keyword": "키워드",
    "englishStory": "영어 스토리",
    "koreanStory": "한글 번역",
    "englishSummary": "영문 요약",
    "koreanSummary": "한글 요약"
  }
  ```

**시연 흐름**
1. 사용자가 키워드 입력
2. AI가 영어 스토리 생성, 한글 번역 동시 제공
3. 대화 로그/번역본 실시간 표시
4. 저장 시 md 파일로 보관, 사이드 리스트/팝업으로 확인

---

## 2. 전체 프로세스의 흐름
1. **프론트엔드** (`static/app.js`, `index.html`, `style.css`)
    - 키워드 입력, WebSocket 연결, 대화 로그/번역/요약 표시, 저장/불러오기, 팝업/사이드바 UI
2. **WebSocket 서버** (`hub.go`)
    - 클라이언트 연결 관리, 메시지 브로드캐스트, Ollama 핸들러 호출
3. **Ollama 연동/번역/요약** (`ollama.go`)
    - 키워드 영어 번역 → Ollama로 스토리 생성(제목은 **로 추출) → 한글 번역 → 요약(영/한)
4. **스토리 저장/불러오기 API** (`main.go`)
    - json 파일 저장, 리스트, 개별 조회

---

## 3. 채팅 입력 시 동작 순서
1. **사용자 키워드 입력** (프론트엔드 input)
2. **WebSocket으로 서버에 메시지 전송**
3. **서버(hub.go)에서 메시지를 받으면면 handleOllama 호출**
4. **handleOllama(ollama.go)에서 처리**
    - (1) 키워드 영어 번역 (구글 번역 API)
    - (2) Ollama llama3.2로 영어 스토리 생성 (title은 **로 추출), (이때 스토리는 스트리밍 응답을 받아 실시간으로 클라이언트에 전송송)
    - (3) 영어 스토리 → 한글 번역 (구글 번역 API)
    - (4) Ollama로 영문 요약 생성
    - (5) 영문 요약 → 한글 요약 (구글 번역 API)
    - (6) 결과를 `{ title, englishStory, koreanStory, englishSummary, koreanSummary }`로 WebSocket으로 전송
5. **프론트엔드(app.js)에서 실시간으로 결과 표시**
    - 실시간으로 영문 스토리를 받으면서 표시하다가
    - 채팅방에 제목, 영문/한글 스토리, 영문/한글 요약된 json 데이터가 오면 이 값으로 UI 업데이트트
    - 저장 시 동일 구조로 저장, 사이드바/팝업에서 상세 확인 가능

---

## 4. 주요 파일별 역할
- **WebSocket/허브**: `hub.go`
- **Ollama/번역/요약**: `ollama.go`
- **API/서버**: `main.go`
- **프론트엔드**: `static/app.js`, `static/index.html`, `static/style.css`
- **저장 파일**: `static/storage/`

---

## 5. 기타(저장/불러오기, 마크다운 렌더링 등)
- **스토리 저장/불러오기**: `main.go`의 `/save-story`, `/list-stories`, `/story/` 핸들러
- **UI/UX**: ChatGPT 스타일, 팝업 닫기, 자동 스크롤 등 사용자 경험 강화

---

## 📁 코드 위치 요약
- **WebSocket/허브**: `hub.go`
- **Ollama/번역/요약**: `ollama.go`
- **API/서버**: `main.go`
- **프론트엔드**: `static/app.js`, `static/index.html`, `static/style.css`
- **저장 파일**: `static/storage/` 

## ※ 후기
- 스트리밍 응답을 지원해 실시간 AI 서비스에 적합하다.
- 프롬프트 엔지니어링이 중요하며, 영어에 최적화된 모델이 많다.
- 다국어 지원이 약한 경우 외부 번역 API와 조합해 서비스 품질을 높일 수 있다.
- Ollama는 단순히 모델을 띄우는 것뿐 아니라, 실시간 대화형 AI 서비스의 백엔드 엔진으로도 충분히 활용 가능하다.
- 웹소캣 + HTTP 통신만 할 줄 알면 쉽게 서비스 가능 