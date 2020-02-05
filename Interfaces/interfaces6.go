//Mine
package main

import "fmt"

func main() {
	var var1 interface1 = calculate{value1:56,value2:61}
	//Values provided by this 'calculate' struct will be stored to var1 interface which will pass it on to Add() function
	var1.Add()

	var var2 interface2 = calculate{value1:23,value2:91}
	//Values provided by this 'calculate' struct will be stored to var2 interface which will pass it on to Multiply() function
	var2.Multiply()
}

//Implementing multiple interfaces
type interface1 interface {
	Add() (int,error)
}

type interface2 interface {
	Multiply() (int,error)
}

//A type can implement multiple interfaces
//Since now the struct is defining its own fields, we have to directly provide its values from the variable which is using the interface in the main() function
type calculate struct {
	value1 int
	value2 int
}

func (cal calculate) Add() (int,error){
	a,err:=fmt.Println(cal.value1+cal.value2)
	return a,err
}

func (cal calculate) Multiply() (int,error){
	b,err:=fmt.Println(cal.value1*cal.value2)
	return b,err
}
