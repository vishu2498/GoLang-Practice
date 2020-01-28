package main

import (
	"fmt"
	"sync"
)
var wg=sync.WaitGroup{}
func main() {
	ch:=make(chan int)
	//Goroutines are also allowed to occur in loops
	//Here, we are sending a value to the channel five times and receiving it five times
	for j:=0;j<5;j++ {
		wg.Add(2)
		go func() {
			i:=<-ch
			fmt.Println(i)
			wg.Done()
		}()
		go func() {
			ch<-42
			wg.Done()
		}()
		wg.Wait()
	}
}
