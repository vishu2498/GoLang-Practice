package main

import (
	"fmt"
	"strings"
)

func main() {
	//Strings in GoLang are actually aliases of bytes
	var1 := "vishu"
	fmt.Printf("%v, %T", var1[2], var1[2]) //This will provide the ASCII (UTF-8) value of the byte in the string
	fmt.Println()
	//We have the option to have the string from the byte by typecasting
	fmt.Printf("%v, %T", string(var1[2]), var1[2])
	fmt.Println()
	//Strings are generally immutable. So we can't change the inside values of string by just changing value from byte.

	//String-Concatenation
	var2, var3 := "hello", "there"
	fmt.Println(var2 + var3)

	//Strings can be explicitly converted to slice of bytes
	var4 := "hello"
	var5 := []byte(var4)
	fmt.Println(var5) //This will show the ASCII (UTF-8) values of all the characters of the string

	//'string.Trim()' from 'strings' package cuts the string given in the 2nd argument from the 1st argument
	//And returns the remaining slice of string
	var6:="egstring"
	fmt.Println(strings.Trim(var6,"eg"))
	fmt.Println(var6) //showing that 'strings.Trim()' works on copy of data passed into arguments and not on original data
}
