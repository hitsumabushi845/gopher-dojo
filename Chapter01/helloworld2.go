package main

import "fmt"

func main() {

	// fmt.Println
	// Prints given value(s) with newline
	fmt.Println("Hello, new world!")
	fmt.Println("A", 100, true, 1.5) // A 100 true 1.5

	// fmt.Printf
	// Prints given formatted string (requires \n for newline)
	fmt.Printf("Hello, new world again!\n")
	fmt.Printf("%d-%s\n", 100, "Even")

	var price int
	fmt.Print("Price>")
	fmt.Scan(&price)
	fmt.Printf("%d yen\n", price)

}
