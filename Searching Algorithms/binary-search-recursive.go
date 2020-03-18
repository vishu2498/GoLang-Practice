//Binary Search (Recursive-Manner)
//It compulsory for the list to be sorted to implement binary search.
package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, 6, 8, 11, 13, 27, 54}
	digit := 11
	low := 0
	high := len(arr) - 1
	if BinarySearch(arr, low, high, digit) == true {
		fmt.Println("value found")
	} else {
		fmt.Println("value not found")
	}
}

func BinarySearch(arr []int, low, high, digit int) bool {
	var mid int
	if low <= high {
		mid = (low + high) / 2
		if digit == arr[mid] {
			return true
		} else if digit < arr[mid] {
			return BinarySearch(arr, low, mid-1, digit)
		} else {
			return BinarySearch(arr, mid+1, high, digit)
		}
	}
	return false
}