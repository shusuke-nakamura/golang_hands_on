package main

import "net/http"

func main() {
	// list6-1 NotFoundサーバを動かす
	http.ListenAndServe(":8080", http.NotFoundHandler())
}
