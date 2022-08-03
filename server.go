package main

import (
	"io/ioutil"
	"net/http"
)

func NewQueueServer(queue Queue) *QueueServer {
	return &QueueServer{queue: queue}
}

type Queue interface {
	Dequeue() string
	Enqueue(item string)
}

type QueueServer struct {
	queue Queue
}

func (q QueueServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := http.NewServeMux()

	router.Handle("/enqueue", http.HandlerFunc(q.enqueueHandler))
	router.Handle("/dequeue", http.HandlerFunc(q.dequeueHandler))

	router.ServeHTTP(w, r)
}

func (q QueueServer) enqueueHandler(w http.ResponseWriter, r *http.Request) {
	item, _ := ioutil.ReadAll(r.Body)

	if len(item) == 0 {
		w.WriteHeader(http.StatusBadRequest)
	}

	q.queue.Enqueue(string(item))

	w.WriteHeader(http.StatusAccepted)
}

func (q QueueServer) dequeueHandler(w http.ResponseWriter, r *http.Request) {
	res := q.queue.Dequeue()

	if res == "" {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.Write([]byte(res))
		w.WriteHeader(http.StatusOK)
	}
}
