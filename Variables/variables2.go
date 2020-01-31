package main

import (
	"fmt"
	"strconv"
)
func main() {
	//Type Casting
	var var13 float64 = 51.32
	var var14 int = int(var13) //information can also be lost during type casting/conversion
	fmt.Printf("%v, %T",var14,var14)
	fmt.Println()

	var15:=13
	var var16 string = strconv.Itoa(var15) //'Itoa' means integer to ASCII
	fmt.Println(var16)

	//In case if it is needed to perform operations on two unequal data-type values, then we have to type cast
	var var3 uint64 = 14
	var var4 int = 48
	fmt.Println(var4/(int(var3)))
}