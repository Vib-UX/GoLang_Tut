// Parallel and concurrent programming
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

// Declaration for Mutex locks similar to c++
var m = sync.RWMutex{}
var counter int = 0

func main() {
	var msg = "Hello"
	go func() { // Seperate Go routine // Similar to threads in c++
		fmt.Println(msg) // Because of the closure properties we will have access to MSG variable
	}()

	/*
		Although this is not a best practice because of dependancies involved which could result in race conditions
	*/

	// Alternatively we can pass the msg variable as a parameter to our anonymous function
	go func(message string) {
		fmt.Println(message)
	}(msg)

	msg = "Goodbye"

	time.Sleep(100 * time.Millisecond) // To delay the end of main execution in go scheduler

	/*
		Delaying the execution of main using time is not best practice thus
		alternatively we can use sync technique
	*/

	// Synchronization with wait group
	wg.Add(1)
	go func() {
		fmt.Println(msg)
		wg.Done()
	}()
	wg.Wait()

	// Locking mechanism with Mutex
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock() // Read Lock
		go sayHello()
		m.Lock() // Write Lock
		go increment()
	}
	wg.Wait()

	// Check how many cores available
	fmt.Println("Threads: ", runtime.GOMAXPROCS(-1))
}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
