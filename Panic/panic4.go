package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("starting")
	panicker()
	fmt.Println("ending")
}
func panicker()  {
	//This function will first create a panic and then recover from it
	fmt.Println("going to panic")
	defer func() { //This will execute after the panic() at line 21
		if err:=recover(); err != nil {
			log.Println("Error:",err) //This will print the panic statement. It will not execute if the recover() fails.
			//In case if we realize that we can't recover from the panic, a panic() function can again be inserted in this recover() function.
		}
	}()
	panic("starting panic")
	fmt.Println("recovered from panic") //This will not execute since panic occurred
	//If recover is called outside the deferred function, it will not stop a panicking sequence.
}
