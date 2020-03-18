//Selection Sort
package main

import "fmt"

func main() {
	arr := []int{8, 6, 9, 4, 2}
	SelectionSort(arr)
	fmt.Println(arr)
}

func SelectionSort(arr []int) {
	var i, j, k int
	for i = 0; i < len(arr)-1; i++ {
		for j, k = i, i; j < len(arr); j++ {
			if arr[j] < arr[k] {
				k = j
			}
		}
		arr[i], arr[k] = arr[k], arr[i]
	}
}
