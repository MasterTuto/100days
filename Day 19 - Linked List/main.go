package main

import (
	"fmt"

	"brenocs.dev/100days.day19/linkedlist"
)

func main() {
	sameValuesLinkedList := linkedlist.New(6, 6, 6, 6, 6, 6, 6, 6, 6)
	fmt.Println("Size of sameValuesLinkedList:", sameValuesLinkedList.Length())
	fmt.Println("Removing all 6s from sameValuesLinkedList...")
	sameValuesLinkedList.AddToBeginning(1)
	sameValuesLinkedList.AddToEnd(7)

	fmt.Printf("sameValuesLinkedList contains %d: %v\n", 0, sameValuesLinkedList.Contains(0))
	fmt.Printf("sameValuesLinkedList contains %d: %v\n", 1, sameValuesLinkedList.Contains(1))
	fmt.Printf("sameValuesLinkedList contains %d: %v\n", 6, sameValuesLinkedList.Contains(6))
	fmt.Printf("sameValuesLinkedList contains %d: %v\n", 7, sameValuesLinkedList.Contains(7))
	fmt.Printf("sameValuesLinkedList contains %d: %v\n", 8, sameValuesLinkedList.Contains(8))

	sameValuesLinkedList.RemoveAll(6)
	fmt.Println("Resulted LinkedList:", sameValuesLinkedList)
	fmt.Println("Size of it: ", sameValuesLinkedList.Length())
	fmt.Println("First: ", sameValuesLinkedList.First())
	fmt.Println("Last: ", sameValuesLinkedList.Last())

	sameValuesLinkedList.Purge()

	fmt.Println("Resulted LinkedList:", sameValuesLinkedList)
	fmt.Println("Size of it: ", sameValuesLinkedList.Length())

	newList := linkedlist.New(1, 2, 3, 4, 5, 6)
	mappedLinkedList := linkedlist.Mapper[int, int](*newList).Map(func(i int) int { return i * i })
	fmt.Println("Mapped:", mappedLinkedList)

	filteredLinkedList := newList.Filter(func(i int) bool { return i%2 == 0 })
	fmt.Println("Filtered:", filteredLinkedList)

	reducedLinkedList := linkedlist.Reducer[int, int](*newList).Reduce(func(v, acc int) int { return v + acc }, 0)
	fmt.Println("Reduced:", reducedLinkedList)

}
