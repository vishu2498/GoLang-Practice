//Using runtime package
package main

import (
	"fmt"
	"runtime"
)
func main() {
	fmt.Println(runtime.GOROOT()) //shows the current GOROOT defined
	fmt.Println(runtime.GOARCH) //GOARCH is the running program's architecture target
	fmt.Println(runtime.GOOS) //GOOS is the running program's operating system target
	//To view possible combinations of GOOS and GOARCH, run "go tool dist list" in terminal.
	var1:=runtime.Compiler //'Compiler' constant is of string data-type and stores the name of Go compiler being used
	fmt.Println(var1)
	/*GOMAXPROCS shows the total number of CPUs available to run the program. Value can be set to reduce or increase the CPUs.
	Negative number will mean nothing.*/
	fmt.Println(runtime.GOMAXPROCS(4))

	go func() {
		fmt.Println("from anonymous")
		/*Goexit() terminates the goroutine that calls it. No other goroutine is affected.
		Goexit() runs all deferred calls before terminating the goroutine. Because Goexit() is not a panic, any recover calls in those deferred functions will return nil.
		Calling Goexit() from the main goroutine terminates that goroutine without func main returning.
		Since func main has not returned, the program continues execution of other goroutines. If all other goroutines exit, the program crashes.*/
		runtime.Goexit()
	}()
	fmt.Println("bro") //this gets printed after exiting from anonymous function

	go func() {
		fmt.Println("resumed")
		runtime.Gosched() //It allows itself and other GoRoutines to run. It doesn't suspend current GoRoutine.
		//Also, due to this function mentioned here, this anonymous function will run after the execution of next statement mentioned after GoRoutine.
	}()

	fmt.Println(runtime.NumCPU()) //NumCPU returns the number of logical CPUs usable by the current process.
	fmt.Println(runtime.NumGoroutine()) //NumGoroutine returns the number of goroutines that currently exist.
	go func() {
		runtime.StartTrace() //starts tracing for current process
		fmt.Println("for tracing")
		runtime.StopTrace() //stops tracing for current process
	}()
	fmt.Println(runtime.ReadTrace()) //reads tracing for current process
	fmt.Println(runtime.Version()) //shows GoLang version
}
