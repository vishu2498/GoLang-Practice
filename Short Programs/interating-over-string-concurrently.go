package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	slice1 := []string{"hello", "there", "you"}
	f1(slice1)
}

func f1(values []string) {
	wg.Add(len(values))
	for i := 0; i < len(values); i++ {
		go func(egvalue string) {
			fmt.Println(egvalue)
			wg.Done()
		}(values[i])
	}
	wg.Wait()
}
