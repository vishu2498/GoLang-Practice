package main

import (
	"fmt"
)

//Variadic functions are the functions that can have multiple arguments which are declared at runtime.
//Arguments of these variadic functions are known as Variadic Parameters.
//Only one variadic parameter is allowed in a variadic function.
func main() {
	f1("hello","there","how","are","you")
	fmt.Println()
	f2("here we go:",1,2,3,4,5,6,7,8,9)
}

func f1(msgs ...string) { //'...' literal makes any function as variadic function
	//So, all the arguments must be of the data-type defined in the function argument and this function will store all the arguments in a slice.
	fmt.Println(msgs)
}

//We can have one or multiple parameters before variadic parameter
//Also, the variadic parameter compulsorily has to be the last parameter in the variadic function.
func f2(msg string,values ...int) {
	fmt.Println(msg,values)
}