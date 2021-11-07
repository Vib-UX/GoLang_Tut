// Control Flow if/Switch

package main

import (
	"fmt"
)

func main() {

	// Declaration of if
	if true {
		fmt.Println("The test is true")
	}

	// Usage of if with maps
	demo1 := map[string]int{
		"chocolate": 1,
		"chips":     2,
		"drinks":    4,
	}

	// Printing the value of the key if exists
	if pop, ok := demo1["chps"]; ok { // Scope of pop here is limited to the if {}
		fmt.Println(pop)
	} else {
		fmt.Println("Value does not exist in map")
	}

	// Switch with tag
	switch 2 {
	case 1: // Overlapping of tags in cases not allowed
		fmt.Println("Tag is 1")
	case 2:
		fmt.Println("Tag is 2")
	default:
		fmt.Print("This is the default case")
	}

	// Tagless Switch
	i := 10
	switch {
	case i <= 10:
		fmt.Println("less than or equal to 10")
		fallthrough // This will allow to follow the next case (logicless)
	case i <= 20:
		fmt.Println("less than or equal to 20")
	}

	// Type Switching
	var j interface{} = 1
	switch j.(type) {
	case int:
		fmt.Println("j is an int")
		break // We can use break for short circuiting
		fmt.Println("End")
	case float32:
		fmt.Println("j is a float32")
	case string:
		fmt.Println("j is a string")
	}

}
