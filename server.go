package main

import "net/http"

type Queue interface {
	Dequeue() string
}

type QueueServer struct {
	queue Queue
}

func (q QueueServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := http.NewServeMux()

	router.Handle("/enqueue", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
		return
	}))

	router.Handle("/dequeue", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := q.queue.Dequeue()

		if res == "" {
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.Write([]byte(res))
			w.WriteHeader(http.StatusOK)
		}
	}))

	router.ServeHTTP(w, r)
}
