package main

import "net/http"

type Queue interface {
	Dequeue() string
}

type QueueServer struct {
	queue Queue
}

func (q *QueueServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res := q.queue.Dequeue()

	if res == "" {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.Write([]byte(res))
		w.WriteHeader(http.StatusOK)
	}
}
