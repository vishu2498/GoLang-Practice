package main

import (
	"fmt"
)

func main() {
	//Arrays are declared by mentioning the size first (in the brackets) and then declaring its data-type
	arr1:=[3]int{51,94,37} //Initializer Syntax
	fmt.Println(arr1)

	//Different way of defining values of an array by Initializer syntax
	arr6:=[3]int{0:27,1:56, 2:98}
	fmt.Println(arr6)

	//If we don't want to declare the size of the array and just define the values and let GoLang compiler know the size of array itself:
	arr2:=[...]string{"hello","there","you"} //'...' literal tells GoLang compiler to know the size of array itself
	fmt.Println(arr2)

	//Using array with size defined but no values provided
	var arr3 [4]string
	fmt.Println(arr3)

	//Defining values of array after declaring it
	//Values of array can be defined in any order
	var arr4 [4]float64
	arr4[0]=25.31
	arr4[1]=98.24
	arr4[2]=8.59
	arr4[3]=9.58
	fmt.Println(arr4)
	fmt.Println(arr4[2]) //getting 3rd element of array

	//Getting the length of the array
	fmt.Println("Length of array: ",len(arr4))

	//Getting the capacity of array
	//Getting capacity of array is actually useless since the length and capacity will always be equal
 	var arr5 [6]string
	fmt.Println("Capacity of array: ",cap(arr5))
	
	//Copying a single value from one array to another array
	arr7:=[3]int{5,3,2}
	var arr8 [3]int
	arr8[2]=arr7[1] //Copying the 2nd value of arr7 to 3rd value of arr8
}
