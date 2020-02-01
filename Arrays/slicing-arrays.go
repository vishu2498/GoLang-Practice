package main

import "fmt"

func main() {
	//Slice operations can also be performed on arrays with no size defined i.e. using '...' literal

	//Performing Slice Operations
	arr1:=[...]int{1,2,3,4,5,6}
	var1:=arr1[:] //This will slice the all the elements of arr1
	var2:=arr1[3:] //First number is exclusive. It will ignore the first 3 elements of arr1 and store everything else
	var3:=arr1[:4] //Second number is inclusive. It will print the first 4 elements of arr1
	var4:=arr1[3:6] //Using inclusive & exclusive numbers. It will slice (print) the 4th, 5th and 6th elements.
	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(var3)
	fmt.Println(var4)
}
