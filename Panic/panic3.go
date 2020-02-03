package main

import (
	"fmt"
)

func main() {
	//Panics happen after all the defer statements are executed
	fmt.Println("1st")
	defer fmt.Println("3rd")
	defer fmt.Println("2nd")
	panic("panic occurred") //this panic will occur after the defer statements gets executed
	//This means that in any case, 'defer' statements will always execute despite occurrence of panics
}
