//Mine
//Everywhere the optimization is done, '//changed' is written.
/*In this optimized program of 'interfaces6.go', the 'error' data-type in all the methods having '(int,error)' return type
of the given interfaces is removed. It was present previously since the variables of the method were holding the data from
the fmt.Println() function which has the return type of (n int, err error). So, removal of fmt.Println() function gave the
variables the option to just hold the output of the operations performed from the struct values.

Hence, since those methods are only returning values, we need to print those values from the main() function.
So, when those methods are called in main() function, the values returned are stored in 'print' variables and these 'print' variables are printed.
*/
package main

import "fmt"

func main() {
	var var1 interface1 = calculate{value1:56,value2:61}
	//Values provided by this 'calculate' struct will be stored to var1 interface which will pass it on to Add() function
	print1:=var1.Add() //changed
	fmt.Println(print1) //added

	var var2 interface2 = calculate{value1:23,value2:91}
	//Values provided by this 'calculate' struct will be stored to var2 interface which will pass it on to Multiply() function
	print2:=var2.Multiply() //changed
	fmt.Println(print2) //added
}

//Implementing multiple interfaces
type interface1 interface {
	Add() int //changed
}

type interface2 interface {
	Multiply() int //changed
}

//A type can implement multiple interfaces
//Since now the struct is defining its own fields, we have to directly provide its values from the variable which is using the interface in the main() function
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
