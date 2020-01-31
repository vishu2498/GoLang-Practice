package main

import "fmt"

//iota is counter that we can use while creating enumerated constants
const c1 = iota //It's default value is '0' and data-type is int

//Constants can also be declared in blocks
//iota values are scoped inside const blocks
const (
	c2 = iota
	c3 = iota
	c5 = iota
)
const (
	c6 = iota
	c7
	c8
)
func main() {
	fmt.Printf("%v, %T", c1,c1)
	//Output will be progressing numbers since iota simplifies the increasing sequence
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(c5)
	fmt.Println()

	//Even if the value of iota is defined in the const block first time, all the proceeding constants will have the incrementing values
	fmt.Println(c6)
	fmt.Println(c7)
	fmt.Println(c8)
}
