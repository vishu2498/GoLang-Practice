//Bubble Sort
package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 3, 9, 4, 8, 6, 7, 2, 1}
	BubbleSort(arr)
	fmt.Println(arr)
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