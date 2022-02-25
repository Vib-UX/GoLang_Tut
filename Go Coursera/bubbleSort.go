package main

import (
	"fmt"
)

func swap(slc []int, j int) {
	temp := slc[j]
	slc[j] = slc[j+1]
	slc[j+1] = temp
}
func bubbleSort(slc []int) {
	var n int = len(slc)
	for i := 0; i < n-1; i = i + 1 {
		for j := 0; j < n-i-1; j = j + 1 {
			if slc[j] > slc[j+1] {
				swap(slc, j)
			}
		}
	}
}

func main() {
	slc := make([]int, 0)
	var numberOfElements int
	fmt.Println("Enter the number of elements you want to enter: ")
	fmt.Scan(&numberOfElements)

	for n := 0; n < numberOfElements; n = n + 1 {
		var input int
		fmt.Scan(&input)
		slc = append(slc, input)
	}
	bubbleSort(slc)
	fmt.Println("Sorted slice : ", slc)
}
