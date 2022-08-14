package main

import "net/http"

func main() {
	msg := `<html><body>
			<h1>Hello</h1>
			<p>This is GO-server!!</p>
			</body></html>`

	hh := func(w http.ResponseWriter, rq *http.Request) {
		w.Write([]byte(msg))
	}

	http.HandleFunc("/hello", hh)
	http.ListenAndServe(":8080", nil)
}
