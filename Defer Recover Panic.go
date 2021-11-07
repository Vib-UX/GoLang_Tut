package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Defer keyword generally works in LIFO order
	defer fmt.Println("start")
	defer fmt.Println("middle")
	defer fmt.Println("end")

	// Output := end-->middle-->start

	// Defer example
	a := "start"
	defer fmt.Println(a)
	a = "end"

	// Output is start although execution of defer allows us to print the statement after the main but value wont be changed

	// Panic : Occur when program cannot continue at all

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello Go"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
	// Note: Panic happens after all the defer statements executed

	// Panic with fxn
	panicker()
	fmt.Println("Main End")
}

func panicker() {
	fmt.Println("about to panic")
	// Lets create an anonymous function
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
		}

	}()
	panic("something bad happened")
	fmt.Println("End")
}
