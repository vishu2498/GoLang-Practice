package main

import "fmt"

func main() {
	//Constants in which we don't define data-type are known as Untyped Constants.
	const c1 = 59
	fmt.Printf("%v, %T",c1,c1)
	fmt.Println()

	//GoLang compiler is very flexible while working with untyped constants.
	//So, while performing operations with untyped constants, it can change the data-type of the constant according to the different operand.
	const c2 = 14
	var var1 int64 = 61
	fmt.Println(c2+var1) //GoLang compiler will convert c2 to int64
}