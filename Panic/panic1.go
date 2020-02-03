package main

import "fmt"

//GoLang doesn't have exceptions
func main() {
	//In GoLang, when any error occurs or a situation arises where a program can't move further, that situation is called Panic.
	var1,var2:=12,0
	fmt.Println(var1/var2) //This will cause a panic
}