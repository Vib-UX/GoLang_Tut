package main

import (
	"fmt"
	"reflect"
)

// Defining Structures
type Doctor struct {
	number int
	name   string
	spec   []string
}

// Embedding in structures
type Animal struct {
	maxSpeed int
	breed    string
}

type Lion struct {
	Animal // Composition of struct Animal in Lion
	race   string
	king   bool
}

// Tags in struct
type Vehicle struct {
	MaxSpeed int `required max: "100"`
	Type     string
}

func main() {

	// Declaration of Maps
	demo1 := map[string]int{
		"chocolate": 1,
		"waffles":   2,
		"chips":     3,
		"cerels":    21,
	}

	/*
	* In go map generally behaves like unordered_map in c++
	* Maps are referenced type
	* Though we can have only defined datatype for its key and value
	* Eg. Slice cannot be used as key though fixed length array can be used
	 */

	fmt.Println(demo1)

	// Lenght of the map
	fmt.Println(len(demo1))

	// Access the entry from the map if exists
	_, ok := demo1["cerls"] // Remeber "_" is only used as a write access variable
	fmt.Println(ok)

	// Delete an entry from the map
	delete(demo1, "chips")

	// Declaration of Structures
	aDoctor := Doctor{
		number: 1,
		name:   "Ram",
		spec: []string{
			"cardio", "gyna",
		},
	}

	fmt.Println(aDoctor)      // Accesing whole structures
	fmt.Println(aDoctor.spec) // Accessing a particular property of a struct

	// Declaration for Embedding struct
	aLion := Lion{
		Animal: Animal{maxSpeed: 134, breed: "carnivore"},
		race:   "White",
		king:   true,
	}

	fmt.Println(aLion)

	// Tags Access
	t := reflect.TypeOf(Vehicle{})
	field, _ := t.FieldByName("MaxSpeed")
	fmt.Println(field.Tag)

}
