package route

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"shortUrl.com/db"
	"shortUrl.com/model"
	"shortUrl.com/utils"
)

const (
	Port       = 4000
	PublicPath = "/public"
)

func Setting() {
	staticHandler := http.FileServer(http.Dir("." + PublicPath))
	http.Handle(PublicPath+"/", http.StripPrefix(PublicPath, staticHandler))

	http.HandleFunc("/", indexHandler)
	// url, url/1 같은 형식을 urlhandler로 처리
	http.HandleFunc("/url", urlHandler)
	http.HandleFunc("/url/", urlHandler)

	log.Printf("Listening on localhost:%d", Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", Port), nil))
}

// indexHandler : /public/index.html을 보여줌
func indexHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "."+PublicPath)
		return
	default:
		fmt.Fprintf(w, "Sorry, only GET methods are supported.")
	}
}

func urlHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUrlHandler(w)
	case http.MethodPost:
		postUrlHandler(w, r)
	case http.MethodPatch:
		patchUrlHandler(w, r)
	case http.MethodDelete:
		deleteUrlHandler(w, r)
	default:
		fmt.Fprintf(w, "Sorry, only GET, POST, PATCH, DELETE methods are supported.")
	}
}

// getUrlHandler : DB에 저장된 url 정보를 json 형식으로 반환
func getUrlHandler(w http.ResponseWriter) {
	// json 으로 변환
	jsonUrls, err := json.Marshal(db.GetUrlList())
	utils.HandleErr(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonUrls)
	utils.HandleErr(err)
}

// postUrlHandler : DB에 url 정보를 저장
func postUrlHandler(w http.ResponseWriter, r *http.Request) {
	utils.HandleErr(r.ParseForm())
	postData := r.PostForm

	url := model.Url{
		AliasURL: postData.Get("aliasUrl"),
		FullURL:  postData.Get("fullUrl"),
	}

	db.InsertUrl(url)
	w.WriteHeader(http.StatusOK)
}

// patchUrlHandler : DB에 저장된 url 정보를 업데이트
func patchUrlHandler(w http.ResponseWriter, r *http.Request) {
	var url model.Url
	body, err := io.ReadAll(r.Body)
	utils.HandleErr(err)

	// body가 json 형식으로 바꿀 수 있어서 json으로 바꾸고 이를 구조체에 디코딩 해서 저장
	if err := json.Unmarshal(body, &url); err != nil {
		http.Error(w, "Failed to decode JSON data", http.StatusBadRequest)
		return
	}

	db.PatchUrl(url, utils.GetThirdIndexUrl(w, r))
	w.WriteHeader(http.StatusOK)
}

// deleteUrlHandler : DB에 저장된 url 정보를 삭제
func deleteUrlHandler(w http.ResponseWriter, r *http.Request) {
	db.DeleteUrl(utils.GetThirdIndexUrl(w, r))
	w.WriteHeader(http.StatusOK)
}
