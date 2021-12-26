package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func errPanic(w http.ResponseWriter, r *http.Request) error {
	panic(123)

}

func TestErrWrapper(t *testing.T) {
	tests := []struct {
		h       appHandler
		code    int
		message string
	}{
		{errPanic, 500, ""},
	}

	for _, test := range tests {
		f := errWrapper(test.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet,
			"http://imooc.com",
			nil)
		f(response, request)
		b, _ := ioutil.ReadAll(response.Body)
		body := string(b)
		if response.Code != test.code ||
			body != test.message {
			t.Errorf("expect(%d %s)"+
				"got %d %s",
				test.code, test.message,
				response.Code, response.Body)
		}
	}
}

//4-5
