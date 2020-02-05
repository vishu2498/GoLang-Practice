//This program will give a panic.

package main

import "fmt"

func main() {
	var s interface{} = "Steven Paul"
	assert(s)
}

//In case if a value is provided to the interface whose data-type doesn't match the data-type that is defined and fetched from the method, method will return a panic.
func assert(i interface{}) {
	s := i.(int) //we are expecting the interface to provide the value of 'int' data-type but instead we are providing a string to the interface
	fmt.Println(s)
}
/*The output will be:-
panic: interface conversion: interface {} is string, not int */
//To solve this, solution is provided in next file.