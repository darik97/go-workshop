package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"workshop/internal/api"
	"workshop/internal/api/mocks"
)

func TestHandler_Hello(t *testing.T) {
	tests := []struct {
		joke     *api.JokeResponse
		name     string
		err      error
		codeWant int
		bodyWant string
	}{
		{
			name:     "simple test",
			joke:     &api.JokeResponse{Joke: "test joke\n"},
			err:      nil,
			codeWant: 200,
			bodyWant: "test joke\n",
		},
	}
	for _, tt := range tests {
		apiMock := &mocks.Client{}
		apiMock.On("GetJoke").Return(tt.joke, tt.err)

		h := NewHandler(apiMock)

		req, _ := http.NewRequest("GET", "/hello", nil)
		rr := httptest.NewRecorder()

		h.Hello(rr, req)

		gotRaw, _ := ioutil.ReadAll(rr.Body)
		got := string(gotRaw)

		if got != tt.bodyWant {
			t.Errorf("wrong response body %s want %s", got, tt.bodyWant)
		}

		if status := rr.Result().StatusCode; status != tt.codeWant {
			t.Errorf("wrong response body %d want %d", status, tt.codeWant)
		}
	}
}
