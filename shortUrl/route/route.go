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
	"strings"
)

const (
	Port       = 4000
	PublicPath = "/public"
)

func Setting() {
	staticHandler := http.FileServer(http.Dir("." + PublicPath))
	http.Handle(PublicPath+"/", http.StripPrefix(PublicPath, staticHandler))

	// index
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/url", urlHandler)
	http.HandleFunc("/url/", urlHandler)

	log.Printf("Listening on localhost:%d", Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", Port), nil))
}

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
	fmt.Println(r.Method)
	switch r.Method {
	case http.MethodGet:
		getUrlHandler(w, r)
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

func getUrlHandler(w http.ResponseWriter, r *http.Request) {
	// json 으로 변환
	jsonUrls, err := json.Marshal(db.GetUrlList())
	utils.HandleErr(err)

	// http 성공 코드 및 json 데이터를 반환
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonUrls)
	utils.HandleErr(err)
}

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

func patchUrlHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	utils.HandleErr(err)

	var url model.Url
	if err := json.Unmarshal(body, &url); err != nil {
		http.Error(w, "Failed to decode JSON data", http.StatusBadRequest)
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}
	id := parts[2]
	db.PatchUrl(url, id)
	w.WriteHeader(http.StatusOK)
}

func deleteUrlHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}
	id := parts[2]
	db.DeleteUrl(id)
	w.WriteHeader(http.StatusOK)
}
