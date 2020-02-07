package main

import (
	"fmt"
)

func main() {
	//Two-Dimensional Array
	arr1:=[3][3]int{[3]int{1,2,5}, [3]int{5,6,1}, [3]int{4,9,7}}
	fmt.Println(arr1)
	fmt.Println()

	//Copying two-dimensional arrays
	arr7:=[3][3]int{
		0:{0,1,5},
		1:{2,3,4},
		2:{5,7,8},
	}
	fmt.Println(arr7)
	arr8:=arr7
	fmt.Println(arr8)
	fmt.Println()

	//Refined way of declaring and defining two-dimenstional array
	var arr2 [3][3]int
	arr2[0]=[3]int{4,6,1}
	arr2[1]=[3]int{4,6,1}
	arr2[2]=[3]int{4,6,1}
	fmt.Println(arr2)
	fmt.Println()

	//If an array is defined equal to another array, it doesn't point to that array. Rather it creates copy of it and saves it inside itself
	arr3:=[3]int{2,5,4}
	arr4:=arr3
	arr4[1]=8
	fmt.Println(arr3) //Output shows that even if the value is changed in arr4, the original arr3 array has no effect of it
	fmt.Println(arr4) //It will show the changed array

	//To make an array directly point to another array, we need to make use of '&' operator

	arr5:=[3]int{2,4,9}
	arr6:=&arr5
	arr6[1]=5
	fmt.Println(arr5) //Output shows that the value of arr5 is changed too
	fmt.Println(arr6) //Output shows that arr6 is now actually pointing to arr5
}
