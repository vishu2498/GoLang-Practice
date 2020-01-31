package main

import "fmt"

func main() {
	//Initializing variables
	//Variables can't be declared multiple times in same scope
	//Unused variables are not tolerated by GoLang

	var var1 int //declaring variable
	var1 = 94 //defining variable
	fmt.Println(var1)

	var var2 int = 65 //declaring and defining variable together
	fmt.Println(var2)

	var var3 int
	var3 = 15
	var3 = 18 //this value will be printed since this is the latest defined (actually changing the value and not its copy)
	fmt.Println(var3)

	//Short-declaration operator (:=) is used when we want to declare and define the variable without providing its data type
	//Using short-declaration operator
	var4:=23
	var5:="vishu"
	var6:=63.24
	fmt.Println(var4,var5)
	fmt.Printf("%v, %T",var4,var4) //getting the value and its data type (%v is value and &T is data type)
	fmt.Println()
	fmt.Printf("%v, %T",var5,var5)
	fmt.Println()
	fmt.Printf("%v, %T",var6,var6)
	fmt.Println()
	fmt.Println(var7) //showing that GoLang also initializes variables. In this case, since data type is int, value is stored as '0'.
	fmt.Println(var8) //default value will be nothing
	fmt.Println(var9) //default value will be false
	fmt.Println(var10)
	fmt.Println(var11)

	//If a value is defined in global scope then it can be redefined in local scope
	var var11 string = "there"
	fmt.Println(var11)
}

var var7 int //this is package-level or global-level variable
//we can't use (:=) when declaring and defining global-level variable

//We can define the package-level variables in a block instead of writing 'var' keyword everytime while declaring the variables
//Also, we have the ability to define multiple variable blocks
var (
	var8 string
	var9 bool
	var10 bool = true
)
var (
	var11 string = "hello"
)

//Lowercase variables are only available to the package-level
//Uppercase variable are available to all the packages and anyone could use them (globally visible)
var Var12 string = "Hello from My Package"

//It is good to have longer name of variable if it will be used multiple times
//If variables are acronyms, then keep their name in Uppercase (eg. HTML,HTTP,URL)