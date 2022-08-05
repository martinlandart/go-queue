package queue

import (
	"fmt"
	"testing"
)

func BenchmarkSliceQueue(b *testing.B) {
	queue := InMemorySliceQueue{}
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
					queue.Enqueue(fmt.Sprint(j))
				}

				for j := 0; j < v.queueItems; j++ {
					queue.Dequeue()
				}
			}
		})
	}
}

func BenchmarkLinkedListQueue(b *testing.B) {
	queue := LinkedListQueue{}
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
					queue.Enqueue(fmt.Sprint(j))
				}

				for j := 0; j < v.queueItems; j++ {
					queue.Dequeue()
				}
			}
		})
	}
}
