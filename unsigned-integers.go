package main

import "fmt"

func main() {
	//Unsigned integers are those integers that do not hold any negative values
	//They are denoted by 'uint'
	var var1 uint = 65
	fmt.Printf("%v, %T", var1, var1)
	fmt.Println()
	var var2 uint32 = 43
	fmt.Printf("%v, %T", var2, var2)
	fmt.Println()

	//In case if it is needed to perform operations on two unequal data-type values, then we have to type cast
	var var3 uint64 = 14
	var var4 int = 48
	fmt.Println(var4 / (int(var3)))
}
