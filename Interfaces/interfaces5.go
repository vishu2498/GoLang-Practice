package main

import "fmt"

func main() {
	var s interface{} = 56
	assert(s)
	var i interface{} = "Steven Paul" //since string is provided to the interface and 'int' is expected by the method, the comma-ok syntax will handle this and program will not panic
	assert(i)
}

/*To resolve the panic we received which occurred due to mismatched data-type between interface value and the expected data-type from method,
we have to use comma-ok syntax.*/
func assert(i interface{}) {
	v, ok := i.(int)
	fmt.Println(v, ok)
}