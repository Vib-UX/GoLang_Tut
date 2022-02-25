package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	map1 := map[string]string{}
	fmt.Println("Enter your name")

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	name := scanner.Text()
	fmt.Println("Enter your address")
	reader1 := bufio.NewReader(os.Stdin)
	scanner1 := bufio.NewScanner(reader1)
	scanner1.Scan()
	address := scanner1.Text()

	map1["name"] = name
	map1["address"] = address
	jsonObj, err := json.Marshal(map1)
	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(string(jsonObj))

}
