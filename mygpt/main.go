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
		Keyword string `json:"keyword"`
		Title   string `json:"title"`
		Story   string `json:"story"`
	}
	var req Req
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}
	if _, err := os.Stat("static/storage"); os.IsNotExist(err) {
		os.MkdirAll("static/storage", 0755)
	}
	filename := filepath.Join("static", "storage", req.Title+"_"+time.Now().Format("20060102150405")+".md")
	filename = strings.ReplaceAll(filename, "/", "_")
	content := "# " + req.Title + "\n\n**키워드:** " + req.Keyword + "\n\n" + req.Story
	os.WriteFile(filename, []byte(content), 0644)
	w.WriteHeader(http.StatusOK)
}

func listStoriesHandler(w http.ResponseWriter, r *http.Request) {
	files, _ := filepath.Glob("static/storage/*.md")
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
	w.Header().Set("Content-Type", "text/markdown; charset=utf-8")
	io.Copy(w, f)
}

func main() {
	hub := newHub()
	go hub.run()

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	http.HandleFunc("/save-story", saveStoryHandler)
	http.HandleFunc("/list-stories", listStoriesHandler)
	http.HandleFunc("/story/", getStoryHandler)

	log.Println("서버 시작: http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("서버 시작 실패:", err)
	}
}
