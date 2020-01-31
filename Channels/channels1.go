package main

import (
	"fmt"
	"sync"
)
var wg = sync.WaitGroup{}
func main() {
	//Channels are required to synchronize data between GoRoutines.
	/* In GoLang, a Channel is a medium through which a goroutine communicates with another goroutine and this communication is lock-free.
	Or in other words, a channel is a technique which allows to let one goroutine to send data to another goroutine.
	By default channel is bidirectional, means the goroutines can send or receive data through the same channel.*/
	//Channels are always strongly-typed. This means that once we declare the data type for a channel, no value of any other data type is acceptable

	/*make() function is particularly used to build channels because it runs several internal mechanisms for channel and
	allow Go runtime to properly serve the channel*/

	ch:=make(chan int) //'chan' keyword is required build a channel
	wg.Add(2)

	//'<-' operator is used to feed data into channel and retrieve data form it
	//When the '<-' operator is on the left of channel variable, we are retrieving data from the channel
	//When the '<-' operator is on the right of channel variable, we are providing data to the channel
	//We have to declare and define a variable which stores the value retrieved from the channel

	go func() { //goroutine for receiving channel
		i := <- ch //receiving channel
		fmt.Println(i)
		wg.Done()
	}()

	go func() { //goroutine for sending channel
		ch <- 65 //sending channel
		wg.Done()
	}()
	wg.Wait()
}
