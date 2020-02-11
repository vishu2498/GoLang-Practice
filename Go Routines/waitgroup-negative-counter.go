package main

import (
	"fmt"
	"sync"
)

//The WaitGroup counter can have a negative value in the case that if multiple GoRoutines return 'wg.Done()' more than the number defined in the 'wg.Add()'.
//In that case, 'wg.Done()' will decrease the counter to negative value and 'wg.Add()' will panic.
//So, it is best practice to only define 'wg.Add(1)' and immediately use the GoRoutine with 'wg.Done()'.
var wg=sync.WaitGroup{}
func main() {
	var5:="testwg"
	wg.Add(1)
	//Here, we are defining the WaitGroup to take just one GoRoutine but we are using 'wg.Done()' method multiple times.
	//This will reduce the WaitGroup counter to a negative value since there are more GoRoutines decreasing the value of WaitGroup counter compared to the ones increasing the counter.
	go func(var5 string) {
		fmt.Println(var5)
		wg.Done()
		go func(var5 string) {
			fmt.Println(var5)
			wg.Done() //decreasing the WaitGroup counter under '0'
		}(var5)
	}(var5)
	wg.Wait()
}

//Running this program will produce an error: "panic: sync: negative WaitGroup counter"