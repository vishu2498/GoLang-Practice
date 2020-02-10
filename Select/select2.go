package main

import (
	"fmt"
	"time"
)

func main() {
	result1:=make(chan string)
	result2:=make(chan string)

	go printmsg1(result1)
	go printmsg2(result2)

	select {
	case v1:=<-result1:
		fmt.Println(v1)
	case v2:=<-result2:
		fmt.Println(v2)
	//'default' keyword is used when any of the channels are blocking the code and no other cases are working.
	//Here, since it will take longer than the wait time of GoLang to get the data from channel, 'default' case will execute.
	//This is generally used to prevent the 'select' statement from blocking.
	default:
		fmt.Println("took too long to respond")
	//'default' also the prevents the program moving to deadlock state. So, if the select cases are blocking the code and the program is entering deadlock state, 'default' will prevent that.
	//It will also execute when even if the data received from channel is 'nil'
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
