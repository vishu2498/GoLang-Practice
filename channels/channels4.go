//This program is not going to run
package main

import (
	"fmt"
	"sync"
)
func main()  {
	var wg=sync.WaitGroup{}
	ch:=make(chan int)
	//Here, we have taken out the first goroutine from the 'for' loop
	//Now, in this case, the goroutine inside the 'for' loop is pushing the value five times into the channel
	//But there is nothing to process that data inside the loop
	//So, the the goroutine mentioned outside the 'for' loop will receive the value only once but can't do that multiple times
	//Hence, this will cause deadlock.
	go func() {
		i:=<-ch
		fmt.Println(i)
		wg.Done()
	}()
	for j:=0;j<5;j++ {
		wg.Add(2)
		go func() {
			ch<-42 //This will push the data first time to the channel because it can be processed at one time
			//But the channel can only hold one data at a time and also when the data can be processed at all times simultaneously
			//So, the space is not available in the channel to transmit the value second time to the channel
			//This makes it an Unbuffered channel
			wg.Done()
		}()
		wg.Wait()
	}
}
/* Workflow of the program:
1. The goroutine inside the 'for' loop will send the data to the channel.
2. The goroutine outside the 'for' loop will receive it and display it.
3. Control moves again to the goroutine of the 'for' loop.
4. It will send the data to the channel again and will notice that this data now can't be processed and receieved since there is no function or goroutine to get the data from the channel.
5. This will happen at all times for the 'for' loop.
6. Program will have a deadlock
*/
