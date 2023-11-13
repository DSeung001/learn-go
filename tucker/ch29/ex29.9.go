package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello World")
	})

	err := http.ListenAndServeTLS(":3000", "localhost.crt", "localhost.key", nil)
	if err != nil {
		log.Fatal(err)
	}
}
