package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func saveStoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "POST only", http.StatusMethodNotAllowed)
		return
	}
	type Req struct {
		Keyword        string `json:"keyword"`
		Title          string `json:"title"`
		EnglishStory   string `json:"englishStory"`
		KoreanStory    string `json:"koreanStory"`
		EnglishSummary string `json:"englishSummary"`
		KoreanSummary  string `json:"koreanSummary"`
	}
	var req Req
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}
	if _, err := os.Stat("static/storage"); os.IsNotExist(err) {
		os.MkdirAll("static/storage", 0755)
	}
	// 파일명: 제목_타임스탬프.json (특수문자/공백 치환)
	safeTitle := req.Title
	invalidChars := []string{"/", "\\", ":", "*", "?", "\"", "<", ">", "|", " "}
	for _, ch := range invalidChars {
		safeTitle = strings.ReplaceAll(safeTitle, ch, "_")
	}
	filename := filepath.Join("static", "storage", safeTitle+"_"+time.Now().Format("20060102150405")+".json")
	// 저장할 json 구조
	storyObj := map[string]string{
		"title":          req.Title,
		"keyword":        req.Keyword,
		"englishStory":   req.EnglishStory,
		"koreanStory":    req.KoreanStory,
		"englishSummary": req.EnglishSummary,
		"koreanSummary":  req.KoreanSummary,
	}
	jsonBytes, _ := json.MarshalIndent(storyObj, "", "  ")
	err := os.WriteFile(filename, jsonBytes, 0644)
	if err != nil {
		log.Println("파일 저장 실패:", err)
		http.Error(w, "파일 저장 실패", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func listStoriesHandler(w http.ResponseWriter, r *http.Request) {
	files, _ := filepath.Glob("static/storage/*.json")
	var names []string
	for _, f := range files {
		names = append(names, filepath.Base(f))
	}
	json.NewEncoder(w).Encode(names)
}

func getStoryHandler(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/story/")
	filename := filepath.Join("static", "storage", name)
	f, err := os.Open(filename)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	defer f.Close()
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	io.Copy(w, f)
}

func main() {
	hub := newHub()
	go hub.run()

	http.Handle("/", http.FileServer(http.Dir("./static")))
	// ws로 온 데이터는 ollama로 전송송
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	http.HandleFunc("/save-story", saveStoryHandler)
	http.HandleFunc("/list-stories", listStoriesHandler)
	http.HandleFunc("/story/", getStoryHandler)

	log.Println("서버 시작: http://localhost:8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("서버 시작 실패:", err)
	}
}
