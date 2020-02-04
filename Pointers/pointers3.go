package main

import "fmt"

func main() {
	//Values of the array are considered intrinsic to the variable to which the array is defined
	//So, in concept of array, variable defined for the array holds both the values of array and its size
	arr1:=[3]int{5,3,2}
	arr2:=arr1 //arr2 will get the copy of arr1
	arr2[1]=9 //changing value from arr2 will have no effect on arr1
	fmt.Println(arr1,arr2) //the value of arr1 didn't change despite changes done from arr2

	//However, the slice doesn't contain the data itself.
	//It contains a pointer to the first element that the slice is pointing to on the underlying array.
	//The internal representation of the slice actually has a pointer to the array.
	slice1:=[]int{6,2,5}
	slice2:=slice1 //slice2 will get the pointer to the array that slice1 is pointing to rather than the data of the slice1
	//So, they are pointing to the same array
	slice2[1]=8
	fmt.Println(slice1,slice2) //values of both slice1 & slice2 got changed

	//Similar behavior of slices is with maps. They too don't contain the actual data but the pointer to the data.
	map1:= map[string]int{
		"1st" : 64,
		"2nd" : 21,
	}
	map2:=map1
	map2["2nd"]=61 //this will also change the value on map1
	fmt.Println(map1,map2) //both the values of the maps got changed
}
