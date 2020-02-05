//Mine
//Everywhere the optimization is done, '//changed' is written.
/*In this optimized program of 'interfaces7.go', the 'error' data-type in all the methods having '(int,error)' return type
of the given interfaces is removed. It was present previously since the variables of the method were holding the data from
the fmt.Println() function which has the return type of (n int, err error). So, removal of fmt.Println() function gave the
variables the option to just hold the output of the operations performed from the struct values.

Hence, since those methods are only returning values, we need to print those values from the main() function.
So, when those methods are called in main() function, the values returned are stored in 'print' variables and these 'print' variables are printed.
*/
package main

import "fmt"

func main() {
	var var1 result = calculate{value1:56,value2:61}
	//Implementing the 'result' interface will give the variable the access to all the structs that the interfaces inside parent interface are utilizing.
	//Also, the variable will have the access to all the methods that are defined in the interfaces contained by the parent interface.
	//So, both the methods will have to deal with same struct values.
	print1:=var1.Add() //changed
	print2:=var1.Multiply() //changed
	fmt.Println(print1) //added
	fmt.Println(print2) //added
}

//GoLang doesn't have inheritance
//But we can embed interfaces into other interfaces
type interface1 interface {
	Add() int //changed
}

type interface2 interface {
	Multiply() int //changed
}

type result interface {
	//Embedding both interfaces into one interface.
	//So, the variable in the calling function will implement this interface.
	interface1
	interface2
}

type calculate struct {
	value1 int
	value2 int
}

func (cal calculate) Add() int{ //changed
	a:=cal.value1+cal.value2 //changed
	return a //changed
}

func (cal calculate) Multiply() int{ //changed
	b:=cal.value1*cal.value2 //changed
	return b //changed
}