package main

import "fmt"

func main() {
	//Slices are just re-sizable arrays
	//Slice is actually a projection of the underlying array (refer to)
	//Slices can also be referred to pointer for an array
	//They are similarly declared and defined like arrays but no size is defined
	//On a 64-bit architecture, a slice requires 24 bytes of memory.
	slice1:=[]int{4,6,2,1}
	fmt.Println(slice1)
	fmt.Println()

	//Getting length and capacity of the slice
	slice2:=[]int{5,2,3,8,9}
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))
	fmt.Println()

	//Printing the slice using for-loop for every element
	slice6:=[]int{1,2,3,4,5}
	for i:=0; i<len(slice6);i++ {
		fmt.Println(slice6[i])
	}
	fmt.Println()

	//Iterating the slice via for-range-loop
	slice7:=[]int{2,4,8}
	for key,value:=range slice7{ //the 'key' and 'value' variables store the index and data of slice.
		//It is compulsory to write slice name after 'range' keyword.
		fmt.Println(key,value)
	}
	fmt.Println()

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
	//Elements will be ignored from the left-side of slice if number is written on the left-side of ':'
	//Elements will be ignored from the right-side of slice if number is written on the right-side of ':'
	var1:=slice5[:] //This will slice the all the elements of slice5
	var2:=slice5[3:] //First number is exclusive. It will ignore the first 3 elements of slice and store everything else
	var3:=slice5[:4] //Second number is inclusive. It will print the first 4 elements of slice5
	var4:=slice5[3:5] //Using inclusive & exclusive numbers. It will slice (print) the 4th and 5th elements.
	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(var3)
	fmt.Println(var4)
	//If we change an element from slice of a slice by its index, the change will reflect in the original slice too.
	var4[1]=56
	fmt.Println(var4)
	fmt.Println(slice5) //change appeared in the exact index in the original slice too
	fmt.Println()
}
