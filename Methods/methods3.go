package main

import "fmt"

func main() {
	value := test{} //we have to define the value with empty struct in this case
	fmt.Println(value.f1())
}

//Hard-coding a struct value from a method
//Previously, the values of struct were given from the main() function and the variables holding that struct used to call the method to perform operations.
//Now, the value is being provided from the method only.

type test struct {
	code int
}

func (t test) f1() int {
	t.code = 61 //providing struct value from method
	return t.code
}
