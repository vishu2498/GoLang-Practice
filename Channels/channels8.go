package main

import (
	"fmt"
	"sync"
)
//When we want to pass multiple values to channels, we don't assign same or different multiple channel receiving end variables.
//Instead, we use 'for' loop for that.
func main() {
	ch:=make(chan int,50)
	var wg=sync.WaitGroup{}
	wg.Add(2)
	go func(ch <-chan int) {
		//This 'for' loop gets multiple values from that channel as receiving end with the 'range' as the number of values sent through channel
		//While using 'for' loop for channels, the temporary variable assigned for 'for' loop actually gets the channel values instead of their indexes
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch<-42
		ch<-27
		ch<-65
		ch<-96
		ch<-453
		ch<-9821
		close(ch)
		/*If we don't mention this function here, the 'for' loop will continue to execute and will not know when to stop receiving values.
		So it will cause a deadlock.
		close() is used to close the channel at the point of time where it is mentioned in the goroutine.
		Also, if we mention close() before or in between sending values to the channel, it will cause a 'panic' saying that
		we passed value to a closed channel. The panic will be: "panic: send on closed channel". And this 'panic' is not recoverable.
		 */
		wg.Done()
	}(ch)
	wg.Wait()
}
