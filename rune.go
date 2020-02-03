package main

import "fmt"

func main() {
	//Runes are type aliases of int32 characters
	//Rune represents UTF-32 character (unlike UTF-8 characters represented by string)
	//All UTF-8 characters are UTF-32 characters

	var1 := 'v'                      //unlike double-quotes in strings, single-quotes are used for runes
	fmt.Printf("%v ,%T", var1, var1) //This will show the ASCII (UTF-8) value of the value in the variable
	fmt.Println()
	var var2 rune = 'i' //'rune' data-type can also be used to represent them
	fmt.Println(var2)

	//In the strings package, a ReadRune() function is available which reads the data which is written in form of runes
}
