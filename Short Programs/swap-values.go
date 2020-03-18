package main

import "fmt"

func main() {
	value1 := 5
	value2 := 15
	fmt.Println(SwapValues(value1, value2))

	string1 := "hello"
	string2 := "there"
	fmt.Println(SwapString(string1, string2))
}

func SwapValues(value1, value2 int) (int, int) {
	value1, value2 = value2, value1
	return value1, value2
}

func SwapString(string1, string2 string) (string, string) {
	string1, string2 = string2, string1
	return string1, string2
}

