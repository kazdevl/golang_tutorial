package basic_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kazdevl/golang_tutorial/network/basic"
)

func Test_Client(t *testing.T) {
	want := basic.Response{Message: "Want"}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		j, err := json.Marshal(want)
		if err != nil {
			t.Fatal(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8s")
		w.Write(j)
	}))
	defer ts.Close()
	result, err := basic.GetRequest(ts.URL)
	if err != nil {
		t.Error(err)
	}
	if want.Message != result.Message {
		t.Errorf("want %s, but %s", want.Message, result.Message)
	}
}
