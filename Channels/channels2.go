package main

import (
	"fmt"
	"sync"
)
var wg=sync.WaitGroup{}

func main() {
	ch:=make(chan int)
	wg.Add(2)
	//When any value is passed into the channel, the copy of the value is actually passed
	go func() {
		i := <- ch
		fmt.Println(i)
		wg.Done()
	}()

	go func() {
		i := 98 //we can define a variable before and send that variable into the channel
		ch <- i //sending the value to the channel via variable
		i = 52 //this value will not be printed by receiving channel because since it has already got the channel value, it can't see any manipulation to the value afterwards.
		wg.Done()
	}()
	wg.Wait()
}
