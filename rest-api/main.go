package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const port string = ":4000"

// urlDescription : url description 구조체
// `json:"url"`로 json 포멧일 시 출력되는 형태를 정의할 수 있음
// omitempty 를 추가하면 json 포멧에서 출력 안됨
type urlDescription struct {
	URL         string `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

type champion struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var champions []*champion

// jsonContentTypeMiddleware : 반환 타입의 포멧을 Content-Type application/json 으로 정의
func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func main() {
	// 기본 값 세팅
	maokai := &champion{
		1,
		"마오카이",
	}
	champions = append(champions, maokai)

	// 라우터 생성
	router := mux.NewRouter()
	router.Use(jsonContentTypeMiddleware)
	router.HandleFunc("/", documentation).Methods("GET")
	// GET 함수 연결
	router.HandleFunc("/champions", getChampions).Methods("GET")

	fmt.Printf("Listening on http://localhost%s\n", port)
	// 서버 실행
	log.Fatal(http.ListenAndServe(port, router))
}

func getChampions(rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode(champions)
}

// documentation : API 설명 출력
func documentation(rw http.ResponseWriter, r *http.Request) {
	// 앞으로 만들 API
	data := []urlDescription{
		{
			URL:         "/champions",
			Method:      "GET",
			Description: "Read champions",
		},
		{
			URL:         "/champions",
			Method:      "POST",
			Description: "Create champions",
		},
		{
			URL:         "/champions/{id}",
			Method:      "PATCH",
			Description: "Update champions",
		},
		{
			URL:         "/champions/{id}",
			Method:      "DELETE",
			Description: "Delete champions",
		},
	}
	// json 형태로 출력
	json.NewEncoder(rw).Encode(data)
}
