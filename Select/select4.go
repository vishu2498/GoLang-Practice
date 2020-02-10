package main

func main() {
	//If a select-case doesn't have any cases, it is known as empty select.
	/*If this empty case is executed, this error will be received: "fatal error: all goroutines are asleep - deadlock!
	goroutine 1 [select (no cases)]" */
	select {

	}
}
