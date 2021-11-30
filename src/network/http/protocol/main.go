package main

import "net/http"

func main() {
	// TCP
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./")))
	// http1.1
	http.ListenAndServe(":8080", mux)
}
