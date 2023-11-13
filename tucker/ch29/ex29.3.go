package main

import (
	"fmt"
	"net/http"
)

func main() {
	// mux : Multiplexer로 여러 입력 중에 하나를 선택해서 반환하는 디지털 장치
	// ServerMux라고도 부르며 라우터라고도 부름
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello World")
	})
	mux.HandleFunc("/bar", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello Bar")
	})
	http.ListenAndServe(":3000", mux)
}
