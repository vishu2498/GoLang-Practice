package main

import "fmt"

func main() {
	fmt.Println(PalindromeIt(725))
	fmt.Println(IsThisPalindrome(725))
}

func PalindromeIt(value int) int {
	var remainder int
	sum:=0
	for value>0 {
		remainder=value%10
		sum=sum*10+remainder
		value=value/10
	}
	return sum
}

func IsThisPalindrome(value int) bool {
	var finalresult bool
	var remainder int
	sum:=0
	for value>0 {
		remainder=value%10
		sum=sum*10+remainder
		value=value/10
	}

	if sum==value {
		finalresult=true
	} else {
		finalresult=false
	}
	return finalresult
}
