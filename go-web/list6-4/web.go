package main

import "net/http"

func main() {
	hh := func(w http.ResponseWriter, rq *http.Request) {
		w.Write([]byte("Hello, This is GO-server!!"))
	}

	http.HandleFunc("/hello", hh)
	http.ListenAndServe(":8080", nil)
}
