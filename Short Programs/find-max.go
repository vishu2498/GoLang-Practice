package main

import "fmt"

func main() {
	values := []int{1, 5, 92, 104, 21}
	fmt.Println(max(values))
}

func max(values []int) int {
	max := values[0]
	for i := 0; i < len(values); i++ {
		if values[i] > max {
			max = values[i]
		}
	}
	return max
}
