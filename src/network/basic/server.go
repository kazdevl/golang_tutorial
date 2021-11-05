package basic

import (
	"io"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8s")
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"error": "request method should be Get"}`)
		return
	}
	// defer r.Body.Close()
	// body, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	io.WriteString(w, `{"error": "cannot read body"}`)
	// 	return
	// }

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"content": "Hello World"}`)
}
