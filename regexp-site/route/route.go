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

// url인지만 체크하면 컨텐츠가 부족한대 => url을 분석하는 걸로 바꾸자
func regexpHandler(w http.ResponseWriter, r *http.Request) {
	var result []string
	var regexpArr []*regexp.Regexp
	var patterns = []string{
		`^(https?|ftp):\/\/`,
		`^(?:https?:\/\/)?(?:ftp:\/\/)?((?:[\w-]+\.)+[\w]{2,})`,
		`[?&]([^=]+)=([^&]*)`,
	}

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

	for _, regexpObj := range regexpArr {
		// 슬라이스로 오는데 0,1 로 2개 오고 둘이 다른지
		matchs := regexpObj.FindStringSubmatch(url)[1]
		result = append(result, matchs)

	}

	w.Header().Set("Content-Type", "application/json")
	jsonResult, _ := json.Marshal(result)
	_, err := w.Write(jsonResult)
	if err != nil {
		return
	}
}
