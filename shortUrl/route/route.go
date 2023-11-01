package route

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"shortUrl.com/db"
	"shortUrl.com/utils"
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
	switch r.Method {
	case http.MethodGet:
		getUrl(w, r)
	case http.MethodPost:
		postUrl(w, r)
	case http.MethodPatch:
		patchUrl(w, r)
	case http.MethodDelete:
		deleteUrl(w, r)
	default:
		fmt.Fprintf(w, "Sorry, only GET, POST, PATCH, DELETE methods are supported.")
	}
}

// db에서 url 데이터들 가져오기
func getUrl(w http.ResponseWriter, r *http.Request) {
	// json 으로 변환
	jsonUrls, err := json.Marshal(db.GetUrlList())
	utils.HandleErr(err)

	// http 성공 코드 및 json 데이터를 반환
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonUrls)
	utils.HandleErr(err)
}

func postUrl(w http.ResponseWriter, r *http.Request) {

}

func patchUrl(w http.ResponseWriter, r *http.Request) {

}

func deleteUrl(w http.ResponseWriter, r *http.Request) {

}
