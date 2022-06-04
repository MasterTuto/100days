package main

import "fmt"

func main() {
	// Simple for loop, only condition
	fmt.Println("Only condition for: ")
	i := 0
	for i != 5 {
		fmt.Println("i:", i)
		i++
	}

	fmt.Println("\nFull for loop:")
	for j := 5; j < 10; j++ {
		fmt.Println("j:", j)
	}

	fmt.Println("\nOnly initialization and condition")
	for k := 9; k < 15; {
		fmt.Println("k: ", k)
		k++

	}

	// I let for...range for me to work after dealing with slices, arrays or channels

	fmt.Println("\nI let for...range for me to work after dealing with slices, arrays or channels")
}
