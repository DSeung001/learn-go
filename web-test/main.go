package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	staticHandler := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public", staticHandler))

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/public/img/test-img.jpg">`)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public")
}
