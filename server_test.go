package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubQueue struct {
	item string
}

func (s *StubQueue) Dequeue() string {
	return s.item
}

func TestDequeue(t *testing.T) {
	tests := []struct {
		name               string
		dequeuedItem       string
		expectedHTTPStatus int
	}{
		{
			name:               "dequeue empty queue returns empty response",
			dequeuedItem:       "",
			expectedHTTPStatus: http.StatusNoContent},
		{
			name:               "dequeue returns dequeued item",
			dequeuedItem:       "item",
			expectedHTTPStatus: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			queue := StubQueue{item: test.dequeuedItem}
			server := &QueueServer{&queue}

			request := MakeDequeueRequest()
			response := httptest.NewRecorder()

			server.ServeHTTP(response, request)

			assertStatusCode(t, response.Code, test.expectedHTTPStatus)

			want := test.dequeuedItem
			got := response.Body.String()
			if got != want {
				t.Errorf("expected empty response but got %q", got)
			}
		})
	}
}

func MakeDequeueRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodPost, "/dequeue", nil)
	return request
}

func assertStatusCode(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("expected status code %v got %v", want, got)
	}
}
