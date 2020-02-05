//Mine
package main

import "fmt"

func main() {
	var var1 result = calculate{value1:56,value2:61}
	//Implementing the 'result' interface will give the variable the access to all the structs that the interfaces inside parent interface are utilizing.
	//Also, the variable will have the access to all the methods that are defined in the interfaces contained by the parent interface.
	//So, both the methods will have to deal with same struct values.
	var1.Add()
	var1.Multiply()
}

//GoLang doesn't have inheritance
//But we can embed interfaces into other interfaces
type interface1 interface {
	Add() (int,error)
}

type interface2 interface {
	Multiply() (int,error)
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

func (cal calculate) Add() (int,error){
	a,err:=fmt.Println(cal.value1+cal.value2)
	return a,err
}

func (cal calculate) Multiply() (int,error){
	b,err:=fmt.Println(cal.value1*cal.value2)
	return b,err
}