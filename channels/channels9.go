package main

import (
	"fmt"
	"sync"
)

func main() {
	ch:=make(chan int,50)
	var wg=sync.WaitGroup{}
	wg.Add(2)
	//We can also use comma-ok syntax in channels
	//Using comma-ok syntax avoids the problem of accidentally sending values to a closed channel.
	//This is done by manually checking this problem
	go func(ch <-chan int) {
		for {
			if i,ok:=<-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch<-42
		ch<-27
		ch<-96
		close(ch)
	}(ch)
	wg.Wait()
}
//The program will execute correctly showing all the values passed through the channel
//But it will enter in deadlock state after that