package route

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"regexp"
)

func Start(port int64) {
	// 라우터 생성
	router := mux.NewRouter()

	// API 매핑
	router.HandleFunc("/regexp", regexpHandler).Methods("POST")
	// index.html 매핑
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	fmt.Printf("Listening on http://localhost:%d\n", port)
	// 서버 실행
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

func regexpHandler(w http.ResponseWriter, r *http.Request) {
	var result []string
	var pattern = `^(https?:\/\/)([-\w\.]+)([-\w\/]+)+(\?[-\=\%\w]+)(\&[-\=\%\w]+)+`

	url := r.PostFormValue("url")

	// 정규식 유효성 검사
	regexpVal, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("정규식 오류:", err)
		return
	}

	// 정규식 실행
	result = regexpVal.FindStringSubmatch(url)

	// 정규식 실행 결과물을 json으로 반환
	w.Header().Set("Content-Type", "application/json")
	jsonResult, _ := json.Marshal(result)
	_, err = w.Write(jsonResult)
	if err != nil {
		return
	}
}
