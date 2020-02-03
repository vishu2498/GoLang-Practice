package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	runtime1()
	fmt.Println("returning from main()")
}
func runtime1() {
	//this function shows how to recover from runtime panics
	defer rec()
	slice:=[]int{64,98,34}
	fmt.Println(slice[5]) //notice that slice[5] doesn't exist
	fmt.Println("slice[5] doesn't exist")
}
func rec()  { //this is the recovery function for runtime1()
	if r := recover(); r != nil {
		fmt.Println("Recovered", r)
		debug.PrintStack() //this function is used to get the stack trace after recovery
		//it is similar to exception.printStackTrace() of JAVA
	}
}