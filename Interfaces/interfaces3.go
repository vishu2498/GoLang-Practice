package main

import "fmt"

//Type-Assertion is used to extract the underlying value of the interface.
func main() {
	var s interface{} = 56
	assert(s)
}

func assert(i interface{}) { //the argument can be an interface
	s := i.(int) //get the underlying int value from i
	fmt.Println(s)
}

/*Workflow of the program:
1. A variable in the main() function has the data-type of an empty interface and provides a value.
2. This empty interface and its value are provided to the assert() function as argument.
3. The assert() function has an argument of an empty interface. So, it can get any type of value of any valid data-type.
4. The variable inside the assert() function takes the inputted underlying value from the interface in the data-type that the value is being provided.
5. assert() function returns the value in the 'int' data-type.
 */