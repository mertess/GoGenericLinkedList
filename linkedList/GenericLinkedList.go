package linkedList

type ILinkedList[T any] interface {
	PeekHead() *T
	PopHead()
	PeekTail() *T
	PopTail()
	GetValues() []T
	AddToHead(value T)
	AddToTail(value T)
	GetLength() int
}

type LinkedList[T any] struct {
	head, tail *node[T]
	length     int
}

type node[T any] struct {
	prev  *node[T]
	next  *node[T]
	value T
}

func (list *LinkedList[T]) PopHead() {
	if list.length > 1 {
		list.head.next.prev = nil
		list.head = list.head.next
		list.length -= 1
	} else {
		list.head = nil
		list.tail = nil
		list.length = 0
	}
}

func (list *LinkedList[T]) PeekHead() *T {
	if list.length > 0 {
		return &list.head.value
	}
	return new(T)
}

func (list *LinkedList[T]) AddToHead(value T) {
	list.addNodeToHead(&node[T]{value: value})
}

func (list *LinkedList[T]) PopTail() {
	if list.length > 1 {
		list.tail.prev.next = nil
		list.tail = list.tail.prev
		list.length -= 1
	} else {
		list.head = nil
		list.tail = nil
		list.length = 0
	}
}

func (list *LinkedList[T]) PeekTail() *T {
	if list.length > 0 {
		return &list.tail.value
	}
	return new(T)
}

func (list *LinkedList[T]) AddToTail(value T) {
	list.addNodeToTail(&node[T]{value: value})
}

func (list *LinkedList[T]) addNodeToHead(node *node[T]) {
	if list.head == nil {
		list.initList(node)
	} else {
		list.head.prev = node
		node.next = list.head
		list.head = node
	}

	list.length += 1
}

func (list *LinkedList[T]) addNodeToTail(node *node[T]) {
	if list.tail == nil {
		list.initList(node)
	} else {
		list.tail.next = node
		node.prev = list.tail
		list.tail = node
	}

	list.length += 1
}

func (list *LinkedList[T]) GetValues() []T {
	result := []T{}
	pointer := list.head
	for pointer != nil {
		result = append(result, pointer.value)
		pointer = pointer.next
	}

	return result
}

func (list *LinkedList[T]) initList(node *node[T]) {
	list.head = node
	list.head.next = list.tail
	list.tail = node
	list.tail.prev = list.head
}

func (list *LinkedList[T]) GetLength() int {
	return list.length
}
