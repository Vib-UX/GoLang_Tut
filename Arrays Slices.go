// Collection Types Arrays and Slices
package main

import "fmt"

func main() {
	// Simple array creation
	grades := [3]int{12, 14, 15} // Syntax name:= [size]datatype {<elements>}
	fmt.Printf("Grades = %v", grades)

	// Array in go has deep copy
	a := [...]int{1, 2, 3}
	b := a
	// fmt.Println(a, b)
	b[0] = 0
	fmt.Println(b, a) // Changes in b won't be reflected in a

	// For Shallow copy
	c := &a // We use address of operator
	c[2] = 1
	fmt.Println(c, a)

	/*
		Slices are bit different then arrays,
		Its referenced type i.e. changes will be reflected
	*/

	slc := []int{9, 8, 7}
	slc1 := slc[:] // Slice all elements
	slc1[0] = 10
	fmt.Println(slc, slc1)

	/*
		Slices are dynamic in nature means it size can increase w.r.t no. of elements
	*/

	// Concatenation of two slices using append
	slc3 := append(slc[:1], slc[2:]...) // Spread operator (...) is used in the 2nd argument to flatten
	fmt.Println(slc3)

	// Declaration of Slices using make
	slc4 := make([]int, 10, 100) // name := make([]<datatype>,length,capacity)
	fmt.Println(slc4)
	/*
	 Length returns the number of elements in the slice
	 Capacity return the number of elements that canbe hold by the slice
	*/

}
