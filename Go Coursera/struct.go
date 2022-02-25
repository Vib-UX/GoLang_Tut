package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (an *Animal) Eat() string {
	return an.food
}

func (an *Animal) Move() string {
	return an.locomotion
}

func (an *Animal) Speak() string {
	return an.noise
}

func main() {
	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	for true {
		fmt.Println("Request your query > ")
		reader := bufio.NewReader(os.Stdin)
		var scanner = bufio.NewScanner(reader)
		scanner.Scan()
		input := scanner.Text()
		strings.ToLower(input)
		if input == "exit" {
			break
		}
		fields := strings.Split(input, " ")
		if fields[0] == "cow" {
			switch fields[1] {
			case "eat":
				fmt.Println(cow.Eat())
			case "move":
				fmt.Println(cow.Move())
			case "speak":
				fmt.Println(cow.Speak())
			default:
				fmt.Println("Attribute not defined")
			}
		} else if fields[0] == "bird" {
			switch fields[1] {
			case "eat":
				fmt.Println(bird.Eat())
			case "move":
				fmt.Println(bird.Move())
			case "speak":
				fmt.Println(bird.Speak())
			default:
				fmt.Println("Attribute not defined")
			}
		} else if fields[0] == "snake" {
			switch fields[1] {
			case "eat":
				fmt.Println(snake.Eat())
			case "move":
				fmt.Println(snake.Move())
			case "speak":
				fmt.Println(snake.Speak())
			default:
				fmt.Println("Attribute not defined")
			}
		} else {
			fmt.Println("Invalid Request, enter again")
		}
	}

}
