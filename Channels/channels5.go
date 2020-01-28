//Both the goroutines can also act as 'sending and receiving channel' goroutines
package main

import (
	"fmt"
	"sync"
)

func main() {
	ch:=make(chan int)
	var wg=sync.WaitGroup{}
	wg.Add(2)
	go func() {
		i:=<-ch //'42' is received here
		fmt.Println(i)
		ch<-27 //value '27' is sent to the channel again now that the channel is empty
		wg.Done()
	}()
	go func() {
		ch<-42 //sending '42' in the channel
		fmt.Println(<-ch) //'27' is received from the channel and printed
		//Writing '<-ch' in fmt.Println() will also print the value in the channel
		wg.Done()
	}()
	wg.Wait()
}
/* Workflow of the program:
1. 2nd goroutine is send '42' in the channel
2. 1st goroutine is receiving from the channel
3. Channel is now empty
4. 1st goroutine is sending '27' to the channel
5. 2nd goroutine is receiving from the channel
 */