package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(1)
	go func() {
		fmt.Println("This will print afterwards")
		wg.Done()
	}()
	//NumGoroutine returns the number of goroutines that currently exist.
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()
	fmt.Println(runtime.NumGoroutine()) //However, if it used after execution of a goroutine is done, it will not count that goroutine.
}
