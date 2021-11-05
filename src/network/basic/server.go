package basic

import (
	"encoding/json"
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

type Sample struct {
	A string `json:"a"`
	B int    `json:"b"`
}

func JsonStructThrowHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8s")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"error": "request method should be Post"}`)
		return
	}
	data := []Sample{{A: "sample1", B: 1}, {A: "sample2", B: 2}}
	msg, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error": "when struct convert to json"}`)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(msg)
}
