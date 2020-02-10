package main

import "fmt"

func main()  {
	//If we receive the data from the channel that is empty i.e. it doesn't have any sender, the program will enter deadlock state.
	//The error will be: "fatal error: all goroutines are asleep - deadlock!"
	ch:=make(chan int)
	fmt.Printf("%v, %T", ch,ch) //this will just print the address of the channel and its data-type
	fmt.Println(<-ch) //deadlock will occur
}
