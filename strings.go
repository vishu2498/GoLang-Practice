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

	//Iterating over the string since the inside representation of a string is byte
	var var10 string
	var10="hello"
	for key, value := range var10 {
		fmt.Println(key,value) //will print the the index of each character of string with its ASCII value
	}
	fmt.Println()

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
	fmt.Println()
	
	//Characters defined explicitly are actually of 'int32' data-type.
	var7:='v'
	fmt.Printf("%v, %T", var7,var7) //output shows that when a character is printed directly, its ASCII value is shown since it is of 'int32' data-type
	fmt.Println()
	//To get the string representation of any character, '%c' operator is used.
	fmt.Printf("%c",var7)
	fmt.Println()
	//It can also be type-casted to string.
	fmt.Println(string(var7))
	//Since character is of 'int32' data-type integers values can be operated on it to get another simultaneous character in alphabets.
	fmt.Println('h'+2) //will only show the ASCII value
	fmt.Println(string('h'+2))
	fmt.Println()

	//Printing all alphabets from character
	for var9:='a';var9<'a'+26 ; var9++ {
		fmt.Print(string(var9))
	}
}
