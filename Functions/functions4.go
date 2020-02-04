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
}
