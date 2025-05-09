package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/ollama/ollama/api"
)

// ChatRequest/ChatResponse 구조체는 ollama/api 패키지의 정의를 그대로 사용합니다.

// 구글 번역 API를 이용해 입력을 원하는 언어로 번역하는 함수
func translateText(text, sl, tl string) (string, error) {
	endpoint := "https://translate.googleapis.com/translate_a/single?client=gtx&sl=" + sl + "&tl=" + tl + "&dt=t&q=" + url.QueryEscape(text)
	resp, err := http.Get(endpoint)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var result []interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}
	if len(result) > 0 {
		if sentences, ok := result[0].([]interface{}); ok && len(sentences) > 0 {
			var translated strings.Builder
			for _, s := range sentences {
				if seg, ok := s.([]interface{}); ok && len(seg) > 0 {
					if str, ok := seg[0].(string); ok {
						translated.WriteString(str)
					}
				}
			}
			return translated.String(), nil
		}
	}
	return "", nil
}

// handleOllama는 클라이언트 메시지를 받아 Ollama 스트리밍 응답을 브로드캐스트합니다.
func handleOllama(conn *websocket.Conn, hub *Hub, userMsg []byte) {
	// 1) 입력 키워드를 영어로 번역
	englishKeyword, err := translateText(string(userMsg), "auto", "en")
	if err != nil || englishKeyword == "" {
		englishKeyword = string(userMsg) // 실패 시 원문 사용
	}

	// 2) URL 파싱 및 클라이언트 초기화
	baseURL, err := url.Parse("http://localhost:11434")
	if err != nil {
		log.Fatalf("URL 파싱 실패: %v", err)
	}
	client := api.NewClient(baseURL, http.DefaultClient)

	// 3) 타임아웃 컨텍스트 생성
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// 4) 사용자의 키워드를 받아 스토리 생성 프롬프트로 변환
	prompt := "Write a short story in English using the following keywords. Put a \"**\" at the beginning and end of the title: " + englishKeyword
	req := &api.ChatRequest{
		Model: "llama3.2",
		Messages: []api.Message{
			{Role: "user", Content: prompt},
		},
	}

	// 영어 본문만 실시간으로 보여줌
	var storyBuilder strings.Builder
	// ollama에 chat에서 반환 된 값을
	if err := client.Chat(ctx, req, func(resp api.ChatResponse) error {
		storyBuilder.WriteString(resp.Message.Content)
		// 그대로 브로드 캐스팅
		hub.broadcast <- []byte(resp.Message.Content)
		// 이러면 허브에서 run을 통해 연결된 모든 클라이언트한테 메시지 전송송
		return nil
	}); err != nil {
		log.Println("Ollama.Chat 오류:", err)
	}

	// 영어 스토리 전체를 번역
	englishStory := storyBuilder.String()
	koreanStory, err := translateText(englishStory, "en", "ko")
	if err != nil {
		koreanStory = "[번역 실패] " + err.Error()
	}

	// title: story에서 **로 감싸진 부분 추출
	title := "Story"
	if m := strings.Index(englishStory, "**"); m != -1 {
		end := strings.Index(englishStory[m+2:], "**")
		if end != -1 {
			title = strings.TrimSpace(englishStory[m+2 : m+2+end])
		}
	}

	// summary 요약 프롬프트
	summaryPrompt := "Summarize the story in 1-2 sentences:\n" + englishStory
	summaryReq := &api.ChatRequest{
		Model: "llama3.2",
		Messages: []api.Message{
			{Role: "user", Content: summaryPrompt},
		},
	}
	var summaryBuilder strings.Builder
	if err := client.Chat(ctx, summaryReq, func(resp api.ChatResponse) error {
		summaryBuilder.WriteString(resp.Message.Content)
		return nil
	}); err != nil {
		summaryBuilder.WriteString("")
	}
	englishSummary := strings.TrimSpace(summaryBuilder.String())

	// koreanSummary: englishSummary를 한국어로 번역
	koreanSummary, err := translateText(englishSummary, "en", "ko")
	if err != nil {
		log.Println("summary 번역 실패:", err)
		koreanSummary = "[번역 실패] " + err.Error()
	}

	// 마지막에만 JSON(title, story, storyKorean, englishSummary, koreanSummary) 결과 전송
	result := map[string]string{
		"title":          title,
		"englishStory":   englishStory,
		"koreanStory":    koreanStory,
		"englishSummary": englishSummary,
		"koreanSummary":  koreanSummary,
	}
	jsonBytes, _ := json.Marshal(result)
	hub.broadcast <- []byte("__STORY_JSON__" + string(jsonBytes))
	hub.broadcast <- []byte("\n\n__MESSAGE_COMPLETE__")
}
