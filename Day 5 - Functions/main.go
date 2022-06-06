package main

import (
	"fmt"
	"log"
)

func Fibonacci(x int64) int64 {
	if x < 2 {
		return x
	} else {
		return Fibonacci(x-1) + Fibonacci(x-2)
	}
}

func Add(a, b int64) int64 {
	return a + b
}

func Mult(a, b int64) (x int64) {
	x = a * b
	return
}

func AddAndMult(a, b int64) (x, y int64) {
	x = a + b
	y = a * b
	return
}

func SubAndDiv(a, b int64) (int64, int64) {
	x := a - b
	y := a / b
	return x, y
}

func Log(message string) {
	log.SetFlags(log.Lmicroseconds)
	log.Println(message)

}

func ImprimirAntesEDepois() {
	Log("Estou primeiro!")
	defer Log("Estou segundo!")
	Log("Estou terceiro!")
}

func main() {
	fmt.Println("Fibonacci de 5:", Fibonacci(5))
	fmt.Println("Fibonacci de 5:", Fibonacci(10))
	fmt.Print("\n\n")
	fmt.Println("5 + 6:", Add(5, 6))
	fmt.Println("5 * 6:", Mult(5, 6))

	fmt.Print("\n\n")

	x, y := AddAndMult(5, 6)
	fmt.Println("5 (+|*) 6:", x, y)
	x, y = SubAndDiv(5, 6)
	fmt.Println("5 (+|*) 6:", x, y)

	fmt.Print("\n\n")

	ImprimirAntesEDepois()
}
