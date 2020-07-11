package main

import "fmt"

func main() {
	slice1:=[]int{1,2,3,4}
	position:=2 //3rd element will be removed
	fmt.Println(f4(slice1,position))
}

func f4(values []int, number int) []int {
	return append(values[:number],values[number+1:]...)
}
