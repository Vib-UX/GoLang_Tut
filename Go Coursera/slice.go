package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	// Lets create a slice
	slc := make([]int, 3)
	for true {
		fmt.Println()
		var input string
		fmt.Scan(&input)
		if input == "X" {
			break
		}
		cnvt, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}
		slc = append(slc, cnvt)
		// sort.Slice(slc, func(i, j int) bool {
		// 	return slc[i] < slc[j]
		// })
		sort.Ints(slc)
		fmt.Println("Slice in sorted order")
		// for _, v := range slc {
		// 	fmt.Printf("%d ", v)
		// }
		fmt.Println(slc)
	}
}
