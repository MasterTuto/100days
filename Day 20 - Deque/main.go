// create a full main go file that tests lldeque package
// Language: go
// Path: main.go
// Compare this snippet from lldeque\lldeque.go:
// package lldeque
//
// type DoublyLinkedList[T comparable] struct {
// 	head   *Node[T]

package main

import (
	"fmt"

	"brenocs.dev/100days.day20/lldeque"
)

func main() {
	deque := lldeque.New[int]()

	deque.InsertFront(1)
	deque.InsertFront(2)
	deque.InsertFront(3)
	deque.InsertRear(4)
	deque.InsertRear(5)
	deque.InsertRear(6)

	for deque.Size() > 0 {
		item := deque.GetFront()
		fmt.Println("O item a ser retirado: ", *item)
		fmt.Println("Lista antes de retirar: ", deque)
		fmt.Println("Tamanho da lista: ", deque.Size())

		deque.DeleteFront()
	}
}
