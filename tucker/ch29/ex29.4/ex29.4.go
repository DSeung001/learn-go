package main

import "net/http"

func main() {
	//http.Handle("/", http.FileServer(http.Dir("static")))
	//http.ListenAndServe(":3000", nil)

	// static/gopher.jpg로 보여주고 싶으면 아래 코드하면 됨
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":3000", nil)
}
