package main

import (
	"fmt"
	"math"
)

const c4 int64 = 5
func main() {
	//Constants in which we define data-type are known as Typed Constants.
	//Constants that are defined in another packages (eg. math) and require runtime processing can't be determined
	//However if we use already present constants that will not change and don't require runtime processing, those can be determined
	const c1 int = 56
	const c2 int16 = 51
	fmt.Println(c1)
	fmt.Println(c2+int16(c1))
	const c3 float64 = math.Pi //Even if we didn't specify type, math.Pi is always float64
	fmt.Println(c3)
	const c4 int = 6 //Even if any constant is defined at the package level, we can change it at the function level
	//And, its data type can also be changed
	fmt.Println(c4)
	fmt.Println()

	//All arithmetic operations on constants are allowed
	const c5 int64 = 5
	const c6 int64 = 2
	fmt.Println(c5+c6)
	fmt.Println(c5-c6)
	fmt.Println(c5*c6)
	fmt.Println(c5/c6)
	fmt.Println(c5%c6)

	//Constants of all data types can be defined
	const c7 string = "hello"
	const c8 float64 = 56.29
	const c9 bool = false

	//It is also possible to perform operations between variables and constants
	//In case if the type of constant and variable is different, we have to perform type-casting
	var1:=56
	const c10 int = 64
	fmt.Println(var1+c10)
	fmt.Println()


}
