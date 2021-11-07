package main

import "fmt"

func main() {
	greeting := "Hello"
	// Simple Function Pass by reference
	sayMessage(&greeting)

	// Returning multiple values from the function
	sum, err := multiReturn(1, 2, 3, 4, 5)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(sum)

	// Anonymous Function
	func() {
		fmt.Println("Hello Go!")
	}()

	// Error Handling with functions as a type
	var divide func(float64, float64) (float64, error)
	divide = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("Cannot divide it by zero")
		}
		return a / b, nil
	}

	d, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

	// Method Calling in go
	g := greeter{
		greeting: "Salut",
		name:     "Ram",
	}
	g.greet()

}

type greeter struct {
	greeting string
	name     string
}

// Method can be of any types
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

// Function declaration with variadic parameters
func multiReturn(values ...int) (int, error) {
	result := 0
	for _, i := range values {
		result += i
	}
	return result, nil
}

// Parameterized function
func sayMessage(msg *string) {
	fmt.Println(*msg)
}
