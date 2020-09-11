package main

import "fmt"

func main() {
	const firstName string = "Afif"
	fmt.Print(firstName, "!\n")

	fmt.Println("john wick")
	fmt.Println("john", "wick")

	fmt.Print("john wick\n")
	fmt.Print("john ", "wick\n")
	fmt.Print("john", " ", "wick\n")
}
