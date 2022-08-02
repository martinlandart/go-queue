package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubQueue struct {
	items []string
}

func (s StubQueue) Dequeue() string {
	if len(s.items) == 0 {
		return ""
	}
	return s.items[len(s.items)-1]
}

func TestQueue(t *testing.T) {
	t.Run("dequeue empty queue returns empty response", func(t *testing.T) {
		queue := StubQueue{}
		server := &QueueServer{queue}

		request, _ := http.NewRequest(http.MethodPost, "/dequeue", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		statusCode := response.Code
		if statusCode != http.StatusNoContent {
			t.Errorf("expected status code %v got %v", http.StatusNoContent, statusCode)
		}

		body := response.Body.String()
		if body != "" {
			t.Errorf("expected empty response but got %q", body)
		}
	})

	t.Run("dequeue one item", func(t *testing.T) {
		want := "item1"
		queue := StubQueue{items: []string{
			want,
		}}
		server := &QueueServer{queue}

		request, _ := http.NewRequest(http.MethodPost, "/dequeue", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		code := response.Code
		if code != http.StatusOK {
			t.Errorf("expected status code %v got %v", http.StatusOK, code)
		}

		got := response.Body.String()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
