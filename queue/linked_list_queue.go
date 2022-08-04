package queue

type Node struct {
	value string
}

type LinkedListQueue struct {
	head *Node
	end  *Node
}

func (l *LinkedListQueue) Enqueue(item string) {
	if l.head == nil && l.end == nil {
		newNode := &Node{value: item}
		l.head = newNode
		l.end = newNode
	}
}

func (l *LinkedListQueue) Dequeue() string {
	if l.head != nil {
		value := l.head.value
		l.head = nil
		return value
	}

	return ""
}
