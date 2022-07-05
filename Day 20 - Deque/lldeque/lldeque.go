package lldeque

import "fmt"

type DoublyLinkedList[T comparable] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

type Node[T comparable] struct {
	prev  *Node[T]
	next  *Node[T]
	value T
}

type IterateReturn[T comparable] struct {
	value T
	index int
}

func New[T comparable]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (dll *DoublyLinkedList[T]) InsertFront(value T) {
	newNode := &Node[T]{
		prev:  nil,
		next:  nil,
		value: value,
	}

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		newNode.next = dll.head
		dll.head.prev = newNode
		dll.head = newNode
	}

	dll.length++

}

func (dll *DoublyLinkedList[T]) InsertRear(value T) {
	newNode := &Node[T]{
		prev:  nil,
		next:  nil,
		value: value,
	}

	if dll.tail == nil {
		dll.tail = newNode
		dll.head = newNode
	} else {
		newNode.prev = dll.tail
		dll.tail.next = newNode
		dll.tail = newNode
	}

	dll.length++
}

func (dll *DoublyLinkedList[T]) DeleteFront() bool {
	if dll.head == nil {
		return false
	} else {
		dll.head = dll.head.next
		if dll.head == nil {
			dll.tail = nil
		} else {
			dll.head.prev = nil
		}

		dll.length--
		return true
	}
}

func (dll *DoublyLinkedList[T]) DeleteRear() bool {
	if dll.head == nil {
		return false
	} else {
		dll.tail = dll.tail.prev
		if dll.tail == nil {
			dll.head = nil
		} else {
			dll.tail.next = nil
		}

		dll.length--
		return true
	}
}

func (dll *DoublyLinkedList[T]) GetFront() *T {
	return &dll.head.value
}

func (dll *DoublyLinkedList[T]) GetRear() *T {
	return &dll.tail.value
}

func (dll *DoublyLinkedList[T]) IsEmpty() bool {
	return dll.head == nil
}

func (dll *DoublyLinkedList[T]) Size() int {
	return dll.length
}

func (dll *DoublyLinkedList[T]) Erase() {
	dll.head = nil
	dll.tail = nil
}

func (dll DoublyLinkedList[T]) String() string {
	result := "["

	for item := range dll.Iterate() {
		if item.index != 0 {
			result += "-> "
		}

		result += fmt.Sprint(item.value)

		if item.index != dll.Size()-1 {
			result += " <-"
		}
	}

	return result + "]"
}

func (dll DoublyLinkedList[T]) Iterate() <-chan IterateReturn[T] {
	ch := make(chan IterateReturn[T])
	go func() {
		current := dll.head
		for i := 0; current != nil; i++ {
			ch <- IterateReturn[T]{current.value, i}
			current = current.next
		}
		close(ch)
	}()
	return ch
}

func (dll DoublyLinkedList[T]) IterateBackward() <-chan IterateReturn[T] {
	ch := make(chan IterateReturn[T])
	go func() {
		current := dll.tail
		for i := dll.Size() - 1; current != nil; i-- {
			ch <- IterateReturn[T]{current.value, i}
			current = current.prev
		}
		close(ch)
	}()
	return ch
}
