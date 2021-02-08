package main

import "fmt"

func main() {
	a := 3

	switch a {
	case 1, 2:
		fmt.Println("a is 1 or 2")
	default:
		fmt.Println("default")
	}

	switch {
	case a == 1:
		fmt.Println("a is 1")
	default:
		fmt.Println("default")
	}
}
