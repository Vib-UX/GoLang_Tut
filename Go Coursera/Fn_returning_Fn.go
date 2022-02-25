package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(acc, vel, initDisp float64) func(float64) float64 {
	fn := func(time float64) float64 {
		return 0.5*acc*(math.Pow(time, 2)) + (vel * time) + initDisp
	}
	return fn
}

func main() {
	fmt.Println("Enter the value of acceleration, velocity and initial displacement : ")
	var vlc, acc, initDisp, time float64
	fmt.Scan(&vlc, &acc, &initDisp)
	fmt.Println("Enter the value of time: ")
	fmt.Scan(&time)
	fn := GenDisplaceFn(vlc, acc, initDisp)
	fmt.Println(fn(time))

}
