package route

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"regexp"
)

func Start(port int64) {
	router := mux.NewRouter()
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	router.HandleFunc("/regex", regexHandler).Methods("POST")

	fmt.Printf("Listening on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

func regexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("regexHandler 실행")
	url := r.PostFormValue("url")

	pattern := "https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{2,256}\\.[a-z]{2,6}\\b([-a-zA-Z0-9@:%_\\+.~#()?&//=]*)"
	regex, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("정규식 오류:", err)
		return
	}

	if regex.MatchString(url) {
		fmt.Println("정규식을 통과했습니다.")
		w.WriteHeader(http.StatusOK)

	} else {
		fmt.Println("정규식을 통과하지 못했습니다.")
		w.WriteHeader(http.StatusBadRequest)
	}
}
