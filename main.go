package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewQueueServer(&InMemorySliceQueue{})
	log.Fatal(http.ListenAndServe(":5000", server))
}
