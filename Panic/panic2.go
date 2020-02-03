package main

import "fmt"

func main() {
	//panic() function can be used in order to create panic situations
	//Anything written after panic() will not execute
	fmt.Println("hello")
	panic("no way")
	fmt.Println("there") //this will not execute
	//In case if any panic is going to occur when after the current panic, compiler will give 'unreachable code' error.
}
