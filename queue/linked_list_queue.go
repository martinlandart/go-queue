package queue

type Node struct {
	value string
	next  *Node
}

type LinkedListQueue struct {
	head *Node
	end  *Node
}

func (l *LinkedListQueue) Enqueue(item string) {
	newNode := &Node{value: item}

	if l.isEmptyQueue() {
		l.head = newNode
		l.end = newNode
		return
	}

	if l.isQueueOfLengthOne() {
		l.head.next = newNode
		l.end = newNode
		return
	}

	l.end.next = newNode
	l.end = newNode
	return
}

func (l *LinkedListQueue) Dequeue() string {
	if l.isEmptyQueue() {
		return ""
	}

	if l.isQueueOfLengthOne() {
		value := l.head.value
		l.head = nil
		l.end = nil
		return value
	}

	value := l.head.value
	l.head = l.head.next
	return value
}

func (l *LinkedListQueue) isEmptyQueue() bool {
	return l.head == nil && l.end == nil
}

func (l *LinkedListQueue) isQueueOfLengthOne() bool {
	return l.head != nil && l.end != nil && l.head == l.end
}
