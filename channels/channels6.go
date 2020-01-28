package main

import (
	"fmt"
	"sync"
)

func main() {
	ch:=make(chan int)
	var wg=sync.WaitGroup{}
	wg.Add(2)

	//We are now making one goroutine as a sending channel function & one as a receiving channel function
	//This is done by adding arguments to the goroutine i.e. accepting variables

	go func(ch <-chan int) {//By providing this argument, data is coming out of the channel
	//This makes it a receiving channel goroutine (receive only channel)
		i:=<-ch
		fmt.Println(i)
		wg.Done()
	}(ch)

	go func(ch chan<- int) {//By providing this argument, data is going into the channel
	//This makes it a sending channel goroutine (send only channel)
		ch<-42
		wg.Done()
	}(ch)
	wg.Wait()
}

/*Property of channels in Golang is that they cast the argument types from bidirectional to unidirectional
which what happening here when the 'ch' argument is being passed to the goroutine*/
/*This is actually the same program as the 1st program of 'channels' but with more specificity
i.e. telling goroutines that which one is receiving channel and which one is sending channel */