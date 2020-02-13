//Wait for the program to execute completely
package main

import (
	"fmt"
	"time"
)

/*Channels perform the Block operation. It is also called Blocking Channels.
It means that when a channel is receiving the data from the sender, it blocks the whole execution of code/program till it retrieves its data.
So, when the data is finally received (even after sleeping time), the program execution resumes.
This also means that any of the statements written after the receiving statement will not execute and code execution will get stuck till the channel receives data.
It also virtually splits the function (normally main()) into two parts: the part that runs until its time to wait for the channel to receive data, and the part that is run after.
 */

//Select case is just like switch-case but for Channels.
//So, the basic concept is that the channel will receive data from multiple sources at different times. Now, the 'select' case makes it possible to actually receive which data and process or show it.
/*The select statement is used to choose from multiple send/receive channel operations.
The select statement blocks the code until one of the send/receive operation is ready. If multiple operations are ready, one of them is chosen at random.
The syntax is similar to switch except that each of the case statement will be a channel operation.*/
//The first non-blocking channel (channel that will not block the execution of code) will be chosen (send and/or receive channel).
func main() {
	result1:=make(chan string)
	result2:=make(chan string)

	go printmsg1(result1)
	go printmsg2(result2)

	//Since we have mentioned the GoRoutine to sleep for some time, the 'select' case will wait for the value to received from the channel.
	//Here, when the GoRoutines are run parallelly, the channel will first receive the data from 'printmsg2' GoRoutine because it is having less sleep time.
	select {
	case v1:=<-result1:
		fmt.Println(v1)
	case v2:=<-result2:
		fmt.Println(v2)
	}
	//In the case where the channel is closed manually, 'select' will always choose that case despite other sources still providing data to channel.
}

func printmsg1(ch chan string) {
	time.Sleep(10*time.Second)
	ch<-"1st"
}

func printmsg2(ch chan string)  {
	time.Sleep(5*time.Second)
	ch<-"2nd"
}
