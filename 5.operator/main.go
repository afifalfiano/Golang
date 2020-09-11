package main

import "fmt"

func main() {
	// var value = (((2+2)*5)*4 - 2) / 2
	// var isEqual = (value == 2)

	// fmt.Printf("nilai %d (%t) \n", value, isEqual) // nilai 39 (false)

	var left = false
	var right = true

	leftAndRight := left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)

	leftOrRight := left || right
	fmt.Printf("left || right \t(%t) \n", leftOrRight)

	leftReverse := !left
	fmt.Printf("!left \t(%t) \n", leftReverse)

}
