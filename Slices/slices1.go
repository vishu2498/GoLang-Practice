package main

import "fmt"

func main() {
	//Slices are just re-sizable arrays
	//Slice is actually a projection of the underlying array (refer to)
	//Slices can also be referred to pointer for an array
	//They are similarly declared and defined like arrays but no size is defined
	slice1:=[]int{4,6,2,1}
	fmt.Println(slice1)
	fmt.Println()

	//Getting length and capacity of the slice
	slice2:=[]int{5,2,3,8,9}
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	//Unlike arrays, if a slice is declared equal to another slice, then the pointer to the array is copied and now both slices point to the same array
	//So, changing values by any of the slice will change the values of the array
	slice3:=[]string{"hello","there","bye"}
	slice4:=slice3
	slice4[1]="vishu"
	fmt.Println(slice3) //Output shows that changing value from the slice4 will also change the value of slice3
	fmt.Println(slice4)
	fmt.Println()

	//Performing Slice Operations
	slice5:=[]int{1,2,3,4,5,6}
	var1:=slice5[:] //This will slice the all the elements of slice5
	var2:=slice5[3:] //First number is exclusive. It will ignore the first 3 elements of slice and store everything else
	var3:=slice5[:4] //Second number is inclusive. It will print the first 4 elements of slice5
	var4:=slice5[3:6] //Using inclusive & exclusive numbers. It will slice (print) the 4th, 5th and 6th elements.
	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(var3)
	fmt.Println(var4)
	fmt.Println()
}
