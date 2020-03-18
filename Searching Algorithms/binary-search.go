//Binary Search
//It compulsory for the list to be sorted to implement binary search.
package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, 6, 8, 11, 13, 27, 54}
	digit := 11
	if BinarySearch(arr, digit) == true {
		fmt.Println("value found")
	} else {
		fmt.Println("value not found")
	}
}

func BinarySearch(arr []int, digit int) bool {
	var low int
	var mid int
	var high int
	low = 0
	high = len(arr) - 1
	for low <= high {
		mid = (low + high) / 2
		if digit == arr[mid] {
			return true
		} else if digit < arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}