package main

import (
	"fmt"

	"brenocs.dev/100days.day21/bst"
)

func main() {
	items := []int{25, 15, 50, 10, 22, 35, 70, 4, 12, 18, 24, 31, 44, 66, 90}

	bsearchTree := bst.New[int]()

	for _, item := range items {
		bsearchTree.Insert(item)
	}

	fmt.Println("InOrder traversal: ")
	for item := range bsearchTree.InOrder() {
		fmt.Print(item, " ")
	}
	fmt.Print("\n\n")

	fmt.Println("PreOrder traversal: ")
	for item := range bsearchTree.PreOrder() {
		fmt.Print(item, " ")
	}
	fmt.Print("\n\n")

	fmt.Println("PostOrder traversal: ")
	for item := range bsearchTree.PostOrder() {
		fmt.Print(item, " ")
	}
	fmt.Print("\n\n")

	fmt.Print("\n4 esta nele? ", bsearchTree.Contains(4), "\n\n")

	fmt.Println("Slowly removing elements...")
	for _, item := range items {
		bsearchTree.Remove(item)
		fmt.Println("Removed", item, "| remainig", bsearchTree.Count(), "items")
	}

	fmt.Println("\n4 esta nele? ", bsearchTree.Contains(4))
}
