package main

import "fmt"

//Higher-Order functions are those functions that take one or more functions as arguments or return a function as result.
func main() {
	var1:= func(a,b int) int{ //defining that anonymous function here
		return a*b
	}
	f1(var1)
	fmt.Println()

	var2:=f2() //So, var2 is holding the function returned by f2() function
	fmt.Println(var2(89,14))
}

func f1(var1 func(a,b int) int)  { //an anonymous function is present as an argument inside f1() function
	fmt.Println(var1(72,23))
	//var1 is the variable provided as the argument for the f1() function.
	//var1 has the data-type as an anonymous function
}

//We can even return a function from some function
func f2() func(a,b int) int {
	//In this function, f2() function is returning an anonymous function which is taking two 'int' arguments & its return type is 'int'.
	var2:= func(a,b int) int {
		//Defining that anonymous function here.
		return a*b
	}
	return var2 //this will return the output of the anonymous function
}