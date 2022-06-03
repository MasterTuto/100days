package main

import "fmt"

func main() {
	// Integers
	var (
		num8  int8  = 10
		num16 int16 = 20
		num32 int32 = 30
		num64 int64 = 40
		num   int   = 50

		unum8   uint8   = 11
		unum16  uint16  = 21
		unum32  uint32  = 31
		unum64  uint64  = 41
		unumptr uintptr = 51
		unum    uint    = 61
	)

	fmt.Println("==== Integers =====")
	fmt.Println("Signed:")
	fmt.Printf("num8  = %v\tType: %T\n", num8, num8)
	fmt.Printf("num16 = %v\tType: %T\n", num16, num16)
	fmt.Printf("num32 = %v\tType: %T\n", num32, num32)
	fmt.Printf("num64 = %v\tType: %T\n", num64, num64)
	fmt.Printf("num   = %v\tType: %T\n\n", num, num)

	fmt.Println("Unsigned:")
	fmt.Printf("unum8   = %v\tType: %T\n", unum8, unum8)
	fmt.Printf("unum16  = %v\tType: %T\n", unum16, unum16)
	fmt.Printf("unum32  = %v\tType: %T\n", unum32, unum32)
	fmt.Printf("unum64  = %v\tType: %T\n", unum64, unum64)
	fmt.Printf("unumptr = %v\tType: %T\n", unumptr, unumptr)
	fmt.Printf("unum    = %v\tType: %T\n", unum, unum)
	fmt.Printf("=====================\n\n")

	// Floats
	var (
		real32 float32 = 2.13
		real64 float64 = 3.14
	)

	fmt.Println("===== Floats =======")
	fmt.Printf("real32 = %v\tType: %T\n", real32, real32)
	fmt.Printf("real64 = %v\tType: %T\n", real64, real64)
	fmt.Printf("=====================\n\n")

	// Complex
	var (
		Complex64  complex64  = 23 + 2i
		Complex128 complex128 = 45 + 23i
	)

	fmt.Println("===== Complexes ======")
	fmt.Printf("Complex64 = %v\tType: %T\n", Complex64, Complex64)
	fmt.Printf("Complex128 = %v\tType: %T\n", Complex128, Complex128)
	fmt.Printf("=====================\n\n")

	var (
		b1 byte = 'a'
		b2 byte = 64
	)
	fmt.Println("===== Byte ======")
	fmt.Printf("b1 = %v\tType: %T\n", b1, b1)
	fmt.Printf("b2 = %v\tType: %T\n", b2, b2)
	fmt.Printf("=====================\n\n")

	// rune
	var (
		r1 rune = 'A'
		r2 rune = 123
	)
	fmt.Println("===== Rune ======")
	fmt.Printf("r1 = %v\tType: %T\n", r1, r1)
	fmt.Printf("r2 = %v\tType: %T\n", r2, r2)
	fmt.Printf("=====================\n\n")

	// String
	var (
		str1 string = "Breno Carvalho"
	)
	fmt.Println("===== String ======")
	fmt.Printf("str1 = %v\tType: %T\n", str1, str1)
	fmt.Printf("=====================\n\n")

	var (
		btrue  bool = true
		bfalse bool = true
	)
	fmt.Println("===== Boolean ======")
	fmt.Printf("btrue = %v\tType: %T\n", btrue, btrue)
	fmt.Printf("bfalse = %v\tType: %T\n", bfalse, bfalse)
	fmt.Printf("=====================\n\n")

}
