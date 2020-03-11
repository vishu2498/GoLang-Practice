//Merging two channels so that two channels send values to the another channel and we receive it from final channel.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	chan1 := asChan(1, 3, 5, 7)
	chan2 := asChan(2, 4, 6, 8)
	chan3 := merge(chan1, chan2)
	for v := range chan3 {
		fmt.Println(v)
	}
}

//Channels will use this function which will build the channel, input values into the channel and receive those values.
func asChan(values ...int) <-chan int {
	c := make(chan int) //building the channel
	go func() {
		for _, v := range values {
			c <- v //sending values individually to the channel
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) //making the GoRoutine sleep so that alternate channels can use the GoRoutine
		}
		close(c) //closing the channel
	}()
	return c
}

//merging the values of two channels into one and returning those values
func merge(chan1, chan2 <-chan int) <-chan int {
	newchan := make(chan int) //initializing the final channel
	go func() {
		defer close(newchan) //closing the final channel since not closing it will cause a deadlock
		/*Also, the channel closing is defined in the GoRoutine since if we define it before the GoRoutine execution,
		the channel will close when merge function ends and it may or may not hold all the values from the first 2 channels.*/

		//These 2 boolean variables are used because when the first 2 channels end up providing data and close themselves, then we get '0' values from them.
		//So, these boolean values tell that when the channels get closed, do not move to for-loop and select cases so that the values don't move to the final channel.
		chan1done, chan2done := false, false
		for !chan1done || !chan2done { //this for-loop is an infinite loop depending on the conditions given
			select {
			//Comma-Ok syntax will hold the values from the channel and a boolean value which will tell that the channel is closed or not.
			case receive, ok := <-chan1:
				if !ok {
					chan1done = true //showing channel is closed
					continue //will not execute the operation when the condition is satisfied
				}
				newchan <- receive
			case receive, ok := <-chan2:
				if !ok {
					chan2done = true //showing channel is closed
					continue //will not execute the operation when the condition is satisfied
				}
				newchan <- receive
			}
		}
	}()
	return newchan
}
