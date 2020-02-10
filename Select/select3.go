//Wait for the program to execute completely
//Random case in 'select' statement
package main

import (
	"fmt"
	"time"
)

func main() {
	result1:=make(chan string)
	result2:=make(chan string)
	result3:=make(chan string)

	go printmsg1(result1)
	go printmsg2(result2)
	go printmsg2(result3)

	time.Sleep(3 * time.Second) //telling the main() function to sleep while the GoRoutines execute and return data to channels
	//In the case if the channel receives data from all the sources at the same time, random case will be chosen.
	select {
	case v1:=<-result1:
		fmt.Println(v1)
	case v2:=<-result2:
		fmt.Println(v2)
	case v3:=<-result3:
		fmt.Println(v3)
	}
}

//Note that the time.Sleep() has been removed from both the GoRoutines.
func printmsg1(ch chan string) {
	ch<-"1st"
}

func printmsg2(ch chan string)  {
	ch<-"2nd"
}

func printmsg3(ch chan string)  {
	ch<-"3rd"
}
