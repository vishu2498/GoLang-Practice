package main

import "fmt"

func main() {

	// declare variables
	var a, b int

	fmt.Print("Enter a and b: ")
	// Get inputs a and b from user via keyboard
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)

	// addition expression
	c := a + b

	// prints result
	fmt.Printf("%d + %d: %d\n", a, b, c)
}
