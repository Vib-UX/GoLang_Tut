package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter the file name: ")
	scanner1 := bufio.NewScanner(os.Stdin)
	scanner1.Scan()
	fileName := scanner1.Text()
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error", err)
	}
	type name struct {
		firstName string
		lastName  string
	}

	slc := make([]name, 0)
	var scanner = bufio.NewScanner(file)

	for scanner.Scan() {
		ln := scanner.Text()
		fields := strings.Split(ln, " ")
		slc = append(slc, name{fields[0], fields[1]})
	}

	for _, v := range slc {
		fmt.Print(v.firstName, " ", v.lastName, "\n")
	}

}
