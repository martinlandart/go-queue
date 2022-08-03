package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubQueue struct {
	item         string
	enqueueCalls []string
}

func (s *StubQueue) Dequeue() string {
	return s.item
}

func (s *StubQueue) Enqueue(item string) {
	s.enqueueCalls = append(s.enqueueCalls, item)
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
			server := NewQueueServer(&queue)

			request := NewDequeueRequest()
			response := httptest.NewRecorder()

			server.ServeHTTP(response, request)

			assertStatusCode(t, response.Code, test.expectedHTTPStatus)
			assertResponseBody(t, response.Body.String(), test.dequeuedItem)
		})
	}
}

func TestEnqueue(t *testing.T) {
	t.Run("returns 400 on empty body", func(t *testing.T) {
		queue := StubQueue{}
		server := NewQueueServer(&queue)

		request := NewEnqueueRequest("")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, http.StatusBadRequest)
	})

	t.Run("enqueues item", func(t *testing.T) {
		queue := StubQueue{}
		server := NewQueueServer(&queue)
		item := "item"

		request := NewEnqueueRequest(item)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, http.StatusAccepted)

		if len(queue.enqueueCalls) != 1 {
			t.Errorf("got %d calls to Enqueue, want %d", len(queue.enqueueCalls), 1)
		}

		if queue.enqueueCalls[0] != item {
			t.Errorf("did not store correct item, got %q want %q", queue.enqueueCalls[0], item)
		}
	})
}

func NewDequeueRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodPost, "/dequeue", nil)
	return request
}

func NewEnqueueRequest(item string) *http.Request {
	bodyReader := bytes.NewReader([]byte(item))
	request, _ := http.NewRequest(http.MethodPost, "/enqueue", bodyReader)

	return request
}

func assertStatusCode(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("expected status code %v got %v", want, got)
	}
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
