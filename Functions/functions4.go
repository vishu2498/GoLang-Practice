package main

import "fmt"

func main() {
	//Anonymous functions are the functions without any name.
	//They need to be called right after the are created.
	func() {
		fmt.Println("hello")
	}()

	//Anonymous functions with loop
	for i := 1; i <= 3; i++ {
		func(i int) { //passing argument in anonymous function
			fmt.Println(i)
		}(i)
	}

	//Anonymous functions can be defined as variables so that they can be used anywhere
	func2:= func() bool{ //also giving return type to anonymous functions. This is actually 'var func2 func() bool'.
		return true
	}
	fmt.Printf("%v, %T",func2(),func2())

	//Closures are those anonymous functions that can access the variables defined outside the body of the function
	//There is no need of defining any extra code for an anonymous function to be a closure.
	var1:=12
	func() {
		fmt.Println(var1) //var1 is being accessed from outside the body of the function
	}()
	fmt.Println()

	a := appendStr()
	b := appendStr()
	fmt.Println(a("Vishu"))
	fmt.Println(b("Rishabh"))
	fmt.Println(a("bye"))
	fmt.Println(b("!"))
}

//However, each closure is bound to its surrounding variable.
//It means that when the calling variables are pointing to the functions that are returning closures, they will have their own value in the closure and will get replaced everytime a new calling variable is defined.
func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}
/*The function 'appendStr()' returns a closure.
This closure is bound to the variable 't'.
The variables 'a' and 'b' declared are closures and they are bound to their own value of 't'.
We first call 'a' with the parameter Vishu.
Now the value of a's version of 't' becomes Hello Vishu.
Now, we call 'b' with the parameter Rishabh.
Since 'b' is bound to its own variable 't', b's version of 't' has a initial value of Vishu again.
Hence after this function call, the value of b's version of 't' becomes Hello Rishabh.
 */