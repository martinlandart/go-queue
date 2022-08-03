package main

import (
	"io/ioutil"
	"net/http"
)

type Queue interface {
	Dequeue() string
	Enqueue(item string)
}

type QueueServer struct {
	queue Queue
}

func (q QueueServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := http.NewServeMux()

	router.Handle("/enqueue", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		item, _ := ioutil.ReadAll(r.Body)

		if len(item) == 0 {
			w.WriteHeader(http.StatusBadRequest)
		}

		q.queue.Enqueue(string(item))

		w.WriteHeader(http.StatusAccepted)
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
