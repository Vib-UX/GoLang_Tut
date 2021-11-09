// Channels will always be working in the context of go-routine

package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {

	// // Channels are created using built-in make function
	// ch := make(chan int) // specifying data-type means channel created is strongly typed
	// wg.Add(2)

	// // Lets create some anonymous functions
	// go func() {
	// 	i := <-ch // This means "i" will be receiving data from the channel
	// 	fmt.Println(i)
	// 	wg.Done()
	// }()

	// go func() {
	// 	ch <- 42 // passing data to our channel (Value typed)
	// 	wg.Done()
	// }()
	// wg.Wait()

	// Channel with internal storage
	ch1 := make(chan int, 50)
	wg.Add(2)
	go func(ch1 <-chan int) { // Receive only channel
		/*
			If you look @ at the arrow its coming out from the channel means we will be receiving from the channel
		*/

		for j := range ch1 {
			fmt.Println(j)
		}
		wg.Done()
	}(ch1)

	go func(ch1 chan<- int) { // Send only channel arrow pointing is inwards
		ch1 <- 42
		ch1 <- 21
		close(ch1) // This is done inorder to ensure that channel is closed and range operation can be executed
		wg.Done()
	}(ch1)
	wg.Wait()

}
