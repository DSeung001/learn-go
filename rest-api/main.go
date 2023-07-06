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

// http 요청에 쓰일 구조체, 속성명을 대문자로 시작해야 외부에서 포인터를 통해 값 할당이 가능
type requestBody struct {
	Id   int
	Name string
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
	// 라우터 생성
	router := mux.NewRouter()
	router.Use(jsonContentTypeMiddleware)
	router.HandleFunc("/", documentation).Methods("GET")
	// GET 함수 연결
	router.HandleFunc("/champions", getChampions).Methods("GET")
	// POST 함수 연결
	router.HandleFunc("/champions", postChampions).Methods("POST")

	fmt.Printf("Listening on http://localhost%s\n", port)
	// 서버 실행
	log.Fatal(http.ListenAndServe(port, router))
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
			Payload:     "Name:string",
		},
		{
			URL:         "/champions/{id}",
			Method:      "PATCH",
			Description: "Update champions",
			Payload:     "Name:string",
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

func getChampions(rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode(champions)
}

func postChampions(rw http.ResponseWriter, r *http.Request) {
	// request 값을 받을 구조체 변수 선언
	var requestBody requestBody
	json.NewDecoder(r.Body).Decode(&requestBody)

	// request 로 받은 Name 값을 사용하여 챔피언 추가
	newChampion := &champion{
		newID(),
		requestBody.Name,
	}
	champions = append(champions, newChampion)

	// 생성 완료 후 Http code 201 created 로 response 해더 설정
	rw.WriteHeader(http.StatusCreated)
	return
}

// newID : 새롭게 추가할 값의 아이디 가져오기
func newID() int {
	if len(champions) > 0 {
		return champions[len(champions)-1].ID + 1
	} else {
		return 1
	}
}
