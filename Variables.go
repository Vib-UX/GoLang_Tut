/* 1. Variables cant be declared 2 times in the same scope
   2. Go is strongly typed you can't declare the variable which is not used in the application
*/

package main

import (
	"fmt"
	"strconv"
)

// At package level go declaration we have to use full var declaration, ":" declaration wont work
var p int = 1000

// Shadowing of var i
var i int = 10

func main() {
	// Declaration of variable of type int
	var i int
	i = 27 // local variable gets precedence over global scope
	fmt.Println("Value of i = ", i)

	// Simple Declaration of variable
	var j int = 32
	fmt.Println("Value of j = ", j)

	// Smart Go Compiler figures out the type
	k := 12
	fmt.Println(k)

	// Conversion of int variable type to float :
	var temp int = 21
	fmt.Printf("%v, %T\n", temp, temp)
	var cnvTemp float32
	cnvTemp = float32(temp)
	fmt.Printf("%v, %T\n", cnvTemp, cnvTemp)

	// Conversion of int to string
	var samp1 int = 48
	var cnvSamp1 string = strconv.Itoa(samp1) // Integer to ascii
	fmt.Printf("%v, %T\n", cnvSamp1, cnvSamp1)

}
