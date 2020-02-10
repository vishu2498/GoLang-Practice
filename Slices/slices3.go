package main

import "fmt"

func main() {
	//Multi-Dimensional Slices
	slice1:=[][]int{{10},{100,200}}
	for key,value:=range slice1 {
		fmt.Println(key,value)
	}
	fmt.Println()

	//Appending to multi-dimensional slice
	//When appending is complete an entire new slice of integers, a new underlying array is allocated and then copied back into index 0 of the outer slice.
	slice1[0]=append(slice1[0],20)
	for key,value:=range slice1 {
		fmt.Println(key,value)
	}
}