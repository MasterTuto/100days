package main

import (
	"fmt"

	"brenocs.dev/100days.day22/heap"
)

func main() {
	h := heap.New[int]()

	h.Insert(4)
	h.Insert(7)
	h.Insert(1)
	h.Insert(3)
	h.Insert(2)
	h.Insert(5)
	h.Insert(6)
	h.Insert(8)
	h.Insert(3)

	fmt.Println(h.Size())

	fmt.Println("Heap: ", h)

	for h.Size() > 0 {
		v, _ := h.Take()
		fmt.Println("Minium: ", v)
	}

	fmt.Println("Heap: ", h)

	fmt.Println("Size: ", h.Size())

}
