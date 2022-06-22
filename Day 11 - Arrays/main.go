package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readInt(query string) int {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(query)
		temp, _ := reader.ReadString('\n')
		result, err := strconv.Atoi(strings.TrimSpace(temp))

		if err != nil {
			fmt.Println("Valor invalido! Tente novamente.")
			continue
		}

		return result
	}
}

func main() {
	// Arrays have fixed size!

	arr := [5]int{1, 2, 3, 4, 5}

	fmt.Println("Please, type 5 numbers, one at a time:")
	for i := 0; i < len(arr); i++ {
		arr[i] = readInt(fmt.Sprintf("%vº>  ", i+1))
	}

	fmt.Printf("The five numbers you typed: %v\n", arr)

	// Sorting...
	sort.Ints(arr[:])

	fmt.Printf("The five numbers you typed sorted: %v\n", arr)
}
