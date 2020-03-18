package main

import "fmt"

func main() {
	arr := []int{5, 2, 7, 8, 1, 4, 9}
	digit := 8
	fmt.Println(Search(arr, digit))
	SearchValueWithPosition(arr, digit)
}

func Search(arr []int, digit int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == digit {
			return true
		}
	}
	return false
}

func SearchValueWithPosition(arr []int, digit int) (int, error) {
	for key, value := range arr {
		if value == digit {
			return fmt.Println("Value:", value, "found at:", key)
		}
	}
	return fmt.Println("didn't get value")
}