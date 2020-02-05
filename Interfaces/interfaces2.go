package main

import "fmt"

func main() {
	myInt:=IntCounter(0) //IntCounter requires an initialized value. This is similar to 'var myInt IntCounter = 0'.
	var inc Incrementer = &myInt //'&' operator is used since the Increment() method inside 'Incrementer' interface has a pointer receiver.
	for i:=0;i<10 ;i++  {
		fmt.Println(inc.Increment())
	}
}

type Incrementer interface {
	Increment() int
}
//This shows that it is not compulsory to use structs to implement interfaces.
//Any custom-type can be used that will be the receiver argument for the method declared by the interface.
type IntCounter int

func (ic *IntCounter) Increment() int {
	//In the case where we want to manipulate the underlying data from the custom data-type, pointer is used for receiver argument.
	*ic++
	return int(*ic) //type-casted since even if IntCounter is used as 'pointer int' in the method, it is actually 'int'
}