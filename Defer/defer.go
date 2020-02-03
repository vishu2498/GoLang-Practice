package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)
//All GoLang programs follow the top-down approach of executing the tasks

//'defer' keyword executes any function passed with it after the function after the function finishes its last statement but before it actually returns
//Deferring doesn't move its defined function to the end of the current function but it actually executes it at the end of the current function.
//Defer is also mainly used for closing resources/files.
func main() {
	f1()
	fmt.Println()
	f2()
	fmt.Println()
	f3()
	fmt.Println()
	f4()
}

func f1()  {
	fmt.Println("1st")
	defer fmt.Println("2nd") //this will execute at the end of the current function
	fmt.Println("3rd")
}

//Deferred functions actually execute in LIFO order.
//So, the last function that gets deferred will be the first one to be called.
func f2()  {
//Execution of below functions will occur in reverse order
	defer fmt.Println("1st")
	defer fmt.Println("2nd")
	defer fmt.Println("3rd")
}

//Using 'defer' for file closing
func f3() {
	file,err:=os.Open("file.txt")
	if err!=nil {
		log.Fatal("error")
	}

	defer file.Close() //will close the at the end of this function

	contents, err := ioutil.ReadFile("file.txt")
	if err!=nil {
		log.Fatal("error")
	}
	fmt.Println(string(contents))
}

//'defer' keyword causes the function to take the argument at the time the 'defer' is called not at the time the function is executed
func f4() {
	var1:="hello" //this will execute
	defer fmt.Println(var1)
	var1="there"
}

