package main

import (
	"fmt"
	"math/rand"
)

func show[T any](name string, arr []T) {
	// Slices have two functions to work with, len and cap
	// len(slice) <= cap(slice) for any slice
	fmt.Println("\n", arr)
	fmt.Printf("%v:    len: %v    cap: %v\n", name, len(arr), cap(arr))
}

func main() {
	// examples
	arr := [6]int{1, 2, 3, 4}
	slice := []int{1, 2, 3, 4, 5, 6}

	// error below
	// append(arr, 3, 4, 5, 6);
	show("slice", slice)
	show("arr", arr[:]) // arr[:] makes it a slice

	// There two ways of creating slices, which generates same elements
	slice1 := []int{5, 6, 7}
	slice2 := make([]int, 6)

	show("slice1", slice1)
	show("slice2", slice2)

	// We can append to slices, which works immutably
	slice1 = append(slice1, 15)
	slice2 = append(slice2, 10)

	show("slice1", slice1)
	show("slice2", slice2)

	iterSlice := make([]int, 1)
	fmt.Println(" ==== iterSlice ====")
	for i := 0; i < 10; i++ {
		n := rand.Intn(10)

		iterSlice = append(iterSlice, n)
		show("iterSlice", iterSlice)
	}
	fmt.Println(" ==== iterSlice ====")

	// there is the nil sice
	var nilSlice []int
	show("nilSlice", nilSlice)

	// copying iterSlice into nilSlice
	copy(iterSlice, nilSlice)
	show("nilSlice", nilSlice)
}
