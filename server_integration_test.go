package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEnqueueingAndDequeueing(t *testing.T) {
	queue := &InMemorySliceQueue{}
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

func BenchmarkQueue(b *testing.B) {
	server := NewQueueServer(&InMemorySliceQueue{})
	table := []struct {
		queueItems int
	}{
		{queueItems: 10},
		{queueItems: 100},
		{queueItems: 1000},
		{queueItems: 100000},
		{queueItems: 1000000},
	}

	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.queueItems), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for j := 0; j < v.queueItems; j++ {
					server.ServeHTTP(httptest.NewRecorder(), NewEnqueueRequest(fmt.Sprint(j)))
				}

				for j := 0; j < v.queueItems; j++ {
					server.ServeHTTP(httptest.NewRecorder(), NewDequeueRequest())
				}
			}
		})
	}
}
