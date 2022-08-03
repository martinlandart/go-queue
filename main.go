package main

import (
	"log"
	"net/http"
)

type InMemorySliceQueue struct {
	queue []string
}

func (i *InMemorySliceQueue) Enqueue(item string) {
	i.queue = append(i.queue, item)
}

func (i *InMemorySliceQueue) Dequeue() string {
	if len(i.queue) == 0 {
		return ""
	}

	next := i.queue[0]

	i.queue = i.queue[1:]

	return next
}

func main() {
	server := NewQueueServer(&InMemorySliceQueue{})
	log.Fatal(http.ListenAndServe(":5000", server))
}
