package main

import (
	"fmt"
	"math"
)

func main() {
	values := []int{1, 5, 92, 104, 21}
	fmt.Println(max(values))
	fmt.Println()

	//'math.Max()' can also be used to find bigger number (accepts float numbers)
	//But it can only take two values
	fmt.Println(math.Max(65,212))
}

//It will also work for negative values
func max(values []int) int {
	max := values[0]
	for i := 0; i < len(values); i++ { //'i' can be both '0' or '1'
		if values[i] > max { //Don't get confused that values[i] is the index. It is the value of that index.
			max = values[i]
		}
	}
	return max
}
