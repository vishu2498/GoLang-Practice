package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
func main() {
	for i:=0;i<10 ;i++  {
		wg.Add(2)
		go printmsg()
		go increment()
	}
	wg.Wait()
}

func printmsg()  {
	fmt.Println("Counter is:",counter)
	wg.Done()
}

func increment()  {
	counter++
	wg.Done()
}

//Running this program will result in a mess and will print the counters in a jumbled way.
//This will happen every time we run this program.
//This is happening since the GoRoutines are racing each other. There is no synchronization between them.
//This is solved by implementing mutex.

