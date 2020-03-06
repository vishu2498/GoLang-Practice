package main

import "fmt"

func main() {
	//Nil-channel is that channel that it explicitly given nil value.
	//So, the channel will always be in blocked state and passing any value to it will result in program deadlock.
	var ch chan string
	ch = nil
	ch <- "hello"
	fmt.Printf("%v, %T", ch, ch)

	//Also, closing a nil channel will result in a panic.
	//The panic will be: "panic: close of nil channel"
}

