package basic_test

import (
	"app/network/basic"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_HelloHandler(t *testing.T) {
	type input struct {
		req *http.Request
	}
	type want struct {
		statusCode  int
		contentType string
		respBody    string
	}
	tests := []struct {
		name  string
		input input
		want  want
	}{
		{
			name: "error pattern",
			input: input{
				req: httptest.NewRequest(http.MethodPost, "http://dummy.url.com", nil),
			},
			want: want{
				statusCode:  http.StatusBadRequest,
				contentType: "application/json; charset=UTF-8s",
				respBody:    `{"error": "request method should be Get"}`,
			},
		},
		{
			name: "success pattern",
			input: input{
				req: httptest.NewRequest(http.MethodGet, "http://dummy.url.com", nil),
			},
			want: want{
				statusCode:  http.StatusOK,
				contentType: "application/json; charset=UTF-8s",
				respBody:    `{"content": "Hello World"}`,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := httptest.NewRecorder()
			basic.HelloHandler(got, test.input.req)

			if test.want.statusCode != got.Code {
				t.Errorf("want %d, but %d", test.want.statusCode, got.Code)
			}

			if contentType := got.Result().Header["Content-Type"]; test.want.contentType != contentType[0] {
				t.Errorf("want %s, but %s", test.want.contentType, contentType[0])
			}

			if got := got.Body.String(); test.want.respBody != got {
				t.Errorf("want %s, but %s", test.want.respBody, got)
			}
		})
	}
}

func Test_JsonHandler(t *testing.T) {
	type input struct {
		req *http.Request
	}
	type want struct {
		statusCode  int
		contentType string
		respBody    []basic.Sample
	}
	tests := []struct {
		name  string
		input input
		want  want
	}{
		{
			name: "success pattern",
			input: input{
				req: httptest.NewRequest(http.MethodPost, "http://dummy.url.com", nil),
			},
			want: want{
				statusCode:  http.StatusOK,
				contentType: "application/json; charset=UTF-8s",
				respBody:    []basic.Sample{{A: "sample1", B: 1}, {A: "sample2", B: 2}},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := httptest.NewRecorder()
			basic.JsonStructThrowHandler(got, test.input.req)

			if test.want.statusCode != got.Code {
				t.Errorf("want %d, but %d", test.want.statusCode, got.Code)
			}

			if contentType := got.Result().Header["Content-Type"]; test.want.contentType != contentType[0] {
				t.Errorf("want %s, but %s", test.want.contentType, contentType[0])
			}

			body, err := io.ReadAll(got.Result().Body)
			if err != nil {
				t.Error(err)
			}
			defer got.Result().Body.Close()
			result := []basic.Sample{}
			json.Unmarshal(body, &result) //ポインタにする必要がある
			if !reflect.DeepEqual(result, test.want.respBody) {
				t.Errorf("want %v, but %v", result, test.want.respBody)
			}
		})
	}
}
