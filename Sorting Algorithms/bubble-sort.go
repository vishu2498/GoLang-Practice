//Bubble Sort
package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 3, 9, 4, 8, 6, 7, 2, 1}
	BubbleSort(arr)
	fmt.Println(arr)

	arr1 := []int{1, 2, 3, 4, 5, 6}
	BubbleSortFlag(arr1)
	fmt.Println(arr1)
}

func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < (len(arr) - 1 - i); j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

//Checking if list is already sorted
func BubbleSortFlag(arr []int) {
	var flag int
	for i := 0; i < len(arr)-1; i++ {
		flag = 0
		for j := 0; j < (len(arr) - 1 - i); j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = 1
			}
		}
		if flag == 0 {
			panic("list already sorted")
		}
	}
}