package main

import (
	"fmt"
	"sync"
)
//If we send more values to the channel as compared to the number of goroutines that are gonna receive it, we will get deadlock.
/*For eg. If we provide 2 values to the channel and only have one receiving end, then only the 1st value will be received and processed.
2nd value will notice that it has no receiving end. So, it will stuck the goroutine at that moment and program will enter into deadlock.
*/
//This problem is solved by Buffered Channels.
/*By default channels are unbuffered, meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value.
Buffered channels accept a limited number of values without a corresponding receiver for those values.*/
//Buffered channels are designed so that when sender and receiver channels needs some time to process the data, they do their work but also don't want to cause delay in the program and block each other.
func main() {
	ch:=make(chan int,50) //This '50' value is telling the channel that it can hold 50 integers at once (buffered channel)
	var wg=sync.WaitGroup{}
	wg.Add(2)
	go func(ch <-chan int) {
		i:=<-ch
		fmt.Println(i)
		i=<-ch //If we didn't write this another line that is providing channel value to the variable again, we would only get '42' value printed.
		/*So, even if Buffered channel removes the problem of holding only one value to the channel, stopping deadlock and removing panic,
		it doesn't guarantee processing all the values passed to the channel on the receiving end.*/
		fmt.Println(i)
		wg.Done()
	}(ch)
	go func() {
		ch<-42
		ch<-27
		wg.Done()
	}()
	wg.Wait()
}