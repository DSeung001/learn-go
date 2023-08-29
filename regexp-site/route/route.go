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
	router := mux.NewRouter()
	router.HandleFunc("/regexp", regexpHandler).Methods("POST")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	fmt.Printf("Listening on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

func regexpHandler(w http.ResponseWriter, r *http.Request) {
	var result []string
	var regexpArr []*regexp.Regexp
	var regexpReason = []string{
		"", "",
	}
	var patterns = []string{
		"", "",
	}

	fmt.Println("regexpHandler 실행")
	url := r.PostFormValue("url")

	for _, pattern := range patterns {
		regexpVal, err := regexp.Compile(pattern)
		if err != nil {
			fmt.Println("정규식 오류:", err)
			return
		} else {
			regexpArr = append(regexpArr, regexpVal)
		}
	}

	for index, regexpObj := range regexpArr {
		if !regexpObj.MatchString(url) {
			result = append(result, regexpReason[index])
		}
	}

	jsonResult, _ := json.Marshal(result)
	_, err := w.Write(jsonResult)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
}
