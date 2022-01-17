package main

import "net/http"

func main() {
	http.Handle("/file", http.FileServer(http.Dir("/contents")))
	http.ListenAndServe(":8888", nil)
}
