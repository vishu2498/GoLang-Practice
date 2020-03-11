//Everywhere the optimization is done, '//changed' is written.
/*In this optimized program of 'channels14.go', the performance issue is corrected from the sending channels to the select-case statements of the receiving channel.
The values '0' and 'false' were still being received after closing the channel. Hence, taking up a lot of CPU.
To rectify it, the boolean counters representing the closing of channels are removed and now when the channels stop providing the values,
the channels are converted to nil channels.*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	chan1:=asChan(1,3,5,7)
	chan2 := asChan(2, 4, 6, 8)
	chan3 := merge(chan1, chan2)
	for v := range chan3 {
		fmt.Println(v)
	}
}

func asChan(values ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range values {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}

func merge(chan1, chan2 <-chan int) <-chan int {
	newchan := make(chan int)
	go func() {
		defer close(newchan)
		//Boolean variables are removed
		for chan1!=nil || chan2!=nil { //changed (checking if the channels are now nil channels)
			select {
			/*Instead of taking a boolean counter for checking if the channel is closed or not, when the channel finally stops providing the values,
				we convert the channel into a nil channel so that the channel blocks and our select case doesn't execute since channel is already closed after providing the values.			 */
			case receive, ok := <-chan1:
				if !ok {
					chan1=nil //changed (converting channel into nil channel)
					continue
				}
				newchan <- receive
			case receive, ok := <-chan2:
				if !ok {
					chan2 = nil //changed (converting channel into nil channel)
					continue
				}
				newchan <- receive
			}
		}
	}()
	return newchan
}

