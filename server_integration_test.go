package main

import (
	"github.com/martinlandart/go-queue/queue"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEnqueueingAndDequeueing(t *testing.T) {
	queue := &queue.InMemorySliceQueue{}
	server := NewQueueServer(queue)

	t.Run("empty queue returns no content", func(t *testing.T) {
		request := NewDequeueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, http.StatusNoContent)
	})

	t.Run("enqueue and dequeue 3 items", func(t *testing.T) {
		want := []string{
			"item1",
			"item2",
			"item3",
		}

		for _, item := range want {
			server.ServeHTTP(httptest.NewRecorder(), NewEnqueueRequest(item))
		}

		for i := 0; i < 3; i++ {
			request := NewDequeueRequest()
			response := httptest.NewRecorder()

			server.ServeHTTP(response, request)

			assertStatusCode(t, response.Code, http.StatusOK)
			assertResponseBody(t, response.Body.String(), want[i])
		}
	})
}
