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

}