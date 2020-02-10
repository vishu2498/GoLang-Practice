package main

import (
	"fmt"
	"time"
)

//Select case is just like switch-case but for Channels.
//So, the basic concept is that the channel will receive data from multiple sources at different times. Now, the 'select' case makes it possible to actually receive which data and process or show it.
/*The select statement is used to choose from multiple send/receive channel operations.
The select statement blocks the code until one of the send/receive operation is ready. If multiple operations are ready, one of them is chosen at random.
The syntax is similar to switch except that each of the case statement will be a channel operation.*/
//The first non-blocking channel will be chosen (send and/or receive channel).
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
}

func printmsg1(ch chan string) {
	time.Sleep(10*time.Second)
	ch<-"1st"
}

func printmsg2(ch chan string)  {
	time.Sleep(5*time.Second)
	ch<-"2nd"
}
