package main

import "fmt"

func main() {
	// Basic for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Multiple i,j initialized for loop
	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}

	// While form of for loop
	i := 0
	for i < 5 {
		fmt.Print(i)
		i++
	}

	// Label with for loop
Loop: // This can be used to exit from both the for loops
	for j := 0; j < 3; j++ {
		for k := 0; k < 3; k++ {
			fmt.Println(j * k)
			if j*k == 2 {
				break Loop
			}
		}
	}
	// Slice of int with for
	arr := []int{1, 2, 3}
	for k, v := range arr { // for range syntax
		fmt.Println(k, v)
	}

	// Map access with for range
	map1 := map[string]int{
		"chocolate": 1,
		"chips":     2,
		"drinks":    3,
	}
	// Traversing the map
	for k, v := range map1 {
		fmt.Println(k, v)
	}

	//Access only key
	for k := range map1 {
		fmt.Println(k)
		break
	}

	// Access only values
	for _, v := range map1 {
		fmt.Println(v)
	}
}
