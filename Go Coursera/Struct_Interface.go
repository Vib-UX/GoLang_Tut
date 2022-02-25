package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}
type Snake struct{}
type Bird struct{}

func (c Cow) Eat() {
	fmt.Println("grass")
}
func (c Cow) Move() {
	fmt.Println("walk")
}
func (c Cow) Speak() {
	fmt.Println("moo")
}

func (s Snake) Eat() {
	fmt.Println("mice")
}
func (s Snake) Move() {
	fmt.Println("slither")
}
func (s Snake) Speak() {
	fmt.Println("hss")
}

func (b Bird) Eat() {
	fmt.Println("worms")
}
func (b Bird) Move() {
	fmt.Println("fly")
}
func (b Bird) Speak() {
	fmt.Println("peep")
}

func main() {

	var animalMap = make(map[string]Animal)
	animalMap["cow"] = Cow{}
	animalMap["snake"] = Snake{}
	animalMap["bird"] = Bird{}
	for true {
		fmt.Println("> ")
		reader := bufio.NewReader(os.Stdin)
		var scanner = bufio.NewScanner(reader)
		scanner.Scan()
		input := scanner.Text()
		strings.ToLower(input)
		if input == "exit" {
			break
		}
		fields := strings.Split(input, " ")
		if fields[0] == "newanimal" {

			var sample Animal
			switch fields[2] {
			case "cow":
				sample = Cow{}
			case "snake":
				sample = Snake{}
			case "bird":
				sample = Bird{}
			}
			animalMap[fields[1]] = sample
			fmt.Println("Created it!")
		} else if fields[0] == "query" {
			if sample, ok := animalMap[fields[1]]; ok {

				switch fields[2] {
				case "eat":
					sample.Eat()
				case "move":
					sample.Move()
				case "speak":
					sample.Speak()
				default:
					fmt.Println("Attribute not defined")
				}
			} else {
				fmt.Println("Invalid animal input")
			}
		} else {
			fmt.Println("Invalid Request, enter again")
		}
	}
}
