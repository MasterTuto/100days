// Package linkedlist provides a complete solution for linkedlist.
package linkedlist

import "fmt"

// Type LinkedList contains a head that indicates the beginning
// of itself.
type LinkedList[T comparable] struct {
	head   *Node[T]
	length int
}

// Type Node indicates a node, which stores the next node and
// its value.
type Node[T comparable] struct {
	next  *Node[T]
	value T
}

// Type Mapper is used on Map method, as Go doesn't accept
// generics on methods.
type Mapper[T comparable, V comparable] LinkedList[T]

// Type Reducer is used on Reduce function, as Go doesn't accept
// generics on methods.
type Reducer[T comparable, V any] LinkedList[T]

// Function New creates an empty LinkedList
func New[T comparable](elements ...T) *LinkedList[T] {

	ll := &LinkedList[T]{
		head: nil,
	}

	for _, element := range elements {
		ll.AddToEnd(element)
	}

	return ll
}

// Method Empty verifies whether the linkedlist is empty or not.
// It returns true if head is nil, false otherwise.
func (ll LinkedList[T]) Empty() bool {
	return ll.head == nil
}

func (ll LinkedList[T]) Length() int {
	return ll.length
}

// Method AddToEnd appends a value to the end of the linkedlist.
func (ll *LinkedList[T]) AddToEnd(value T) {
	ll.length++

	if ll.Empty() {
		ll.head = &Node[T]{
			next:  nil,
			value: value,
		}
		return
	}

	ptr := ll.head
	for ; ptr.next != nil; ptr = ptr.next {
	}

	ptr.next = &Node[T]{
		next:  nil,
		value: value,
	}
}

// Method AddToBeginning append value to the beginning of the
// linkedlist. It replaces the head.
func (ll *LinkedList[T]) AddToBeginning(value T) {
	ll.length++
	oldHead := ll.head

	newHead := &Node[T]{
		next:  oldHead,
		value: value,
	}

	ll.head = newHead
}

// Method Contains checks whether the linkedlist stores the value.
// It returns true if the value is found.
func (ll *LinkedList[T]) Contains(value T) bool {
	ptr := ll.head
	for ; ptr != nil; ptr = ptr.next {
		if ptr.value == value {
			return true
		}
	}

	return false
}

// Method Remove removes the a value from the LinkedList.
// It returns true if the value if found, false otherwise
func (ll *LinkedList[T]) Remove(value T) bool {
	var oldPtr *Node[T] = nil
	ptr := ll.head

	for ; ptr != nil; ptr = ptr.next {

		if ptr.value == value {
			if oldPtr == nil {
				ll.head = ptr.next
			} else {
				oldPtr.next = ptr.next
			}

			ll.length--
			return true
		}

		oldPtr = ptr
	}

	return false
}

// Method RemoveOnBeginning removes the value on the head
func (ll *LinkedList[T]) RemoveOnBeginning() {
	if ll.Empty() {
		return
	}
	ll.length--

	ll.head = ll.head.next
}

// Method RemoveOnEnd removes the value on the end
func (ll *LinkedList[T]) RemoveOnEnd() {
	if ll.Empty() {
		return
	}

	ll.length--
	var oldPtr *Node[T]
	ptr := ll.head
	for ; ptr.next != nil; ptr = ptr.next {
		oldPtr = ptr
	}

	oldPtr.next = nil

}

// Method RemoveAll deletes all ocorrences of that value.
// It returns true if value if found, false otherwise
func (ll *LinkedList[T]) RemoveAll(value T) bool {
	count := 0

	var oldPtr *Node[T] = nil
	ptr := ll.head

	for ; ptr != nil; ptr = ptr.next {

		if ptr.value == value {
			if oldPtr == nil {
				ll.head = ptr.next
			} else {
				oldPtr.next = ptr.next
			}

			count++
		} else {
			oldPtr = ptr
		}

	}

	ll.length -= count
	return count == 0
}

// Method TakeOnBeginning removes the first element
// and returns it.
func (ll *LinkedList[T]) TakeOnBeginning() T {
	if ll.Empty() {
		return *new(T)
	}
	ll.length--

	result := ll.head.value

	ll.head = ll.head.next

	return result
}

// Method TakeOnEnd removes the last element
// and returns it.
func (ll *LinkedList[T]) TakeOnEnd() T {
	if ll.Empty() {
		return *new(T)
	}

	ll.length--
	var oldPtr *Node[T]
	ptr := ll.head
	for ; ptr.next != nil; ptr = ptr.next {
		oldPtr = ptr
	}

	result := oldPtr.value
	oldPtr.next = nil
	return result
}

// Method First returns the first element.
func (ll *LinkedList[T]) First() T {
	return ll.head.value
}

// Method Last returns the last element.
func (ll *LinkedList[T]) Last() T {
	if ll.Empty() {
		return *new(T)
	}

	ptr := ll.head
	for ; ptr.next != nil; ptr = ptr.next {
	}

	return ptr.value
}

// Method Purge cleans the linkedlist;
func (ll *LinkedList[T]) Purge() {
	ll.length = 0
	ll.head = nil
}

// Method String returns a human readable version of that linkedlist.
func (ll *LinkedList[T]) String() string {
	s := "["

	for ptr := ll.head; ptr != nil; ptr = ptr.next {
		s += fmt.Sprintf("%v", ptr.value)
		if ptr.next != nil {
			s += " -> "
		}
	}

	s += "]"
	return s
}

// Method Iterate is used to range over the linkedlist.
// It returns a channel that sends each value, and closes when reaches end,
// which makes it usable in a for..range loop.
func (ll LinkedList[T]) Iterate() <-chan T {
	c := make(chan T)

	go func() {
		defer close(c)

		for ptr := ll.head; ptr != nil; ptr = ptr.next {
			c <- ptr.value
		}
	}()

	return c
}

// Method Filter creates a new LinkedList based on the value returned by filterFn.
// If the result is true, it appends the value to the end of the linkedlist;
func (ll LinkedList[T]) Filter(filterFn func(T) bool) *LinkedList[T] {
	newLL := New[T]()

	for v := range ll.Iterate() {
		if filterFn(v) {
			newLL.AddToEnd(v)
		}
	}

	return newLL
}

// Method Map creates a new LinkedList with new values.
// The elements of this new linkedlist are the result of mappdingFn.
func (m Mapper[T, V]) Map(mappingFn func(T) V) *LinkedList[V] {
	ll := LinkedList[T](m)
	newLL := New[V]()

	for v := range ll.Iterate() {
		newLL.AddToEnd(mappingFn(v))
	}

	return newLL
}

// Method Reduce takes a reducer function and flats the linkedlist to a single value,
// calculated by this function.
func (m Reducer[T, V]) Reduce(reduceFn func(T, V) V, initialValue V) V {
	ll := LinkedList[T](m)
	var reducedVal V = initialValue

	for v := range ll.Iterate() {
		reducedVal = reduceFn(v, reducedVal)
	}

	return reducedVal
}
