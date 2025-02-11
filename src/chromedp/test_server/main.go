package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/cookie/result", func(w http.ResponseWriter, r *http.Request) {
		cookies := r.Cookies()
		for i, cookie := range cookies {
			log.Printf("from %s, server received cookie %d: %v", r.RemoteAddr, i, cookie)
		}
		buf, err := json.MarshalIndent(cookies, "", "  ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, indexHTML, string(buf))
	})
	http.ListenAndServe(":8544", mux)
}

const indexHTML = `<!doctype html>
<html>
<body>
	<div id="result">%s</div>
</body>
</html>`
