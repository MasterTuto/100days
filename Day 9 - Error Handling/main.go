package main

import (
	"errors"
	"fmt"
	"math"
)

var ErrDivideByZeroVar = errors.New("DivisionByZeroError: Cannot divide by zero!")

type ErrNegativeSqrtStruct struct {
	value float64
}

func (err ErrNegativeSqrtStruct) Error() string {
	return fmt.Sprintf("Cannot get squared of negative value %v", err.value)
}

//////////////////////////
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivideByZeroVar
	}

	return a / b, nil
}

func SquareRoot(x float64) (float64, error) {
	if x < 0 {
		return 1, ErrNegativeSqrtStruct{value: x}
	}

	return math.Sqrt(x), nil
}

func testDivision() {
	show := func(a, b float64) {
		result, err := Divide(a, b)
		if err != nil {
			fmt.Printf("Divide(%v, %v): An error ocurred, description below,\n\t%v\n", a, b, err)
		} else {
			fmt.Printf("Divide(%v, %v) = %v\n", a, b, result)
		}
	}

	show(5, 2)
	show(9, 3)
	show(5, 0)
	fmt.Print("\n\n")
}

func testSqrt() {
	show := func(x float64) {
		result, err := SquareRoot(x)
		if err != nil {
			fmt.Printf("SquareRoot(%v): An error ocurred, description below,\n\t%v\n", x, err)
		} else {
			fmt.Printf("SquareRoot(%v) = %v\n", x, result)
		}
	}

	show(10)
	show(0)
	show(-10)
	fmt.Print("\n\n")
}

func main() {
	testDivision()
	testSqrt()
}
