//Making a slice with 1000 elements and adding them with implementation of GoRoutine
package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	slice1 := make([]int, 1000)
	wg.Add(1)
	go f1(slice1)
	wg.Wait()
}

func f1(slice []int) (int, error) {
	var sum int
	for _, values := range slice {
		sum = sum + values
	}
	wg.Done()
	return fmt.Println(sum)
}

