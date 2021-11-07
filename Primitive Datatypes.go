package main

import (
	"fmt"
)

// In go default initialization is 0 value

func main() {
	//Boolean Type
	var n bool = true
	fmt.Printf("%v, %T\n", n, n)

	m := 1 == 2
	fmt.Printf("%v, %T\n", m, m)

	/* Numeric Types
	int8. int16, int32, int64, (uint)
	float32, float64
	complex64, complex128
	*/

	// String types
	s := "this is a sample string "
	fmt.Printf("%v , %T\n", s, s)

	// Strings are immutable in go
	// s[2] = "u" // This will result in error

	// String concatenation
	s2 := "sample2"
	fmt.Printf("%v, %T\n", s+s2, s+s2)

	// String conversion to collection of bytes
	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b) // The output is each ascii value of the char in the string
	fmt.Println(len(s), len(b))

	//Rune
	r := 'a' // Single quotes represent rune, double quaotes is for string
	fmt.Printf("%v, %T\n", r, r)

	// More about rune refer to goapi
}
