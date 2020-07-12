package main

import "fmt"

func main() {
	fmt.Println(PalindromeItString("dark"))
	fmt.Println(IsThisPalindromeString("nitin"))
}

func PalindromeItString(value string) string {
	var result string
	var beforeResult string
	for i:=len(value)-1;i>=0;i-- {
		beforeResult=string(value[i])
		result=result+beforeResult
	}
	return result
}

func IsThisPalindromeString(value string) bool {
	var finalResult bool
	var beforeResult string
	var result string
	for i:=len(value)-1;i>=0;i-- {
		beforeResult=string(value[i])
		result=result+beforeResult
	}
	if result==value {
		finalResult=true
	} else {
		finalResult=false
	}
	return finalResult
}
