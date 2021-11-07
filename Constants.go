package main

// Constants can also be shadowed @ package level

import "fmt"

// Const block Enumerated Const
const (
	_ = iota
	/* 1. "_" is used as special type of variable which is write only
	   2. iota act as counter for the current const block
	*/
	a
	b
	c
)

// Example for Enum const
const (
	isAdmin = 1 << iota
	isHq
	isUsr
	isDev
)

func main() {

	// Example for iota
	fmt.Println(a, b, c) // value of a,b,c is 1,2,3

	// Typed Constant
	const myConst int = 42
	fmt.Printf("%v %T", myConst, myConst)

	// Untyped Constants
	const samp = 21
	var b int16 = 10
	fmt.Printf("%v, %T\n", samp+b, samp+b) // const samp implicitly converted to type int16

	//Enum Const example
	var roles byte = isAdmin | isDev | isHq
	fmt.Printf("Is Admin? %v\n", roles&isAdmin == isAdmin) //Bitmasking
	fmt.Printf("Is Usr? %v", roles&isUsr == isUsr)

}
