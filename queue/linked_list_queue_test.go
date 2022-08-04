package queue

import "testing"

func TestQueue(t *testing.T) {

	t.Run("dequeue empty queue returns empty string", func(t *testing.T) {
		queue := LinkedListQueue{}
		got := queue.Dequeue()
		assertEquals(t, got, "")
	})

	t.Run("enqueue and dequeue one item", func(t *testing.T) {
		queue := LinkedListQueue{}
		want := "item"
		queue.Enqueue(want)
		got := queue.Dequeue()
		assertEquals(t, got, want)
	})

	t.Run("enqueue one item and dequeue twice", func(t *testing.T) {
		queue := LinkedListQueue{}
		queue.Enqueue("item")
		queue.Dequeue()
		got := queue.Dequeue()
		assertEquals(t, got, "")
	})
	t.Run("enqueue two items and dequeue one item", func(t *testing.T) {
		queue := LinkedListQueue{}
		want := "item"
		queue.Enqueue(want)
		queue.Enqueue("item2")
		got := queue.Dequeue()
		assertEquals(t, got, want)
	})
	t.Run("enqueue two items and dequeue two items", func(t *testing.T) {
		queue := LinkedListQueue{}
		queue.Enqueue("item1")
		queue.Enqueue("item2")
		queue.Dequeue()
		got := queue.Dequeue()
		assertEquals(t, got, "item2")
	})
}

func assertEquals(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
