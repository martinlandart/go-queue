package main

import (
	"github.com/martinlandart/go-queue/queue"
	"log"
	"net/http"
)

func main() {
	server := NewQueueServer(&queue.InMemorySliceQueue{})
	log.Fatal(http.ListenAndServe(":5000", server))
}
